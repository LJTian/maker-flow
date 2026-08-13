package persistd1

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Config defines configuration for persistence-d1.
type Config struct {
	Mode           string // "local" or "d1"
	Driver         string // "sqlite" or "postgres" for local mode
	DatabaseURL    string // Connection string for local DB (e.g. postgres DSN or sqlite path)
	SQLitePath     string // For local mode fallback (default: "/data/app.db")
	AccountID      string // Cloudflare Account ID (for d1 mode)
	DatabaseID     string // Cloudflare D1 Database ID (for d1 mode)
	APIToken       string // Cloudflare API Token (for d1 mode)
	BaseURL        string // Optional override for API base URL (for testing)
}

// ConfigFromEnv reads config from environment variables.
func ConfigFromEnv() Config {
	mode := os.Getenv("DB_MODE")
	if mode == "" {
		mode = "local"
	}
	path := os.Getenv("SQLITE_PATH")
	if path == "" {
		path = "/data/app.db"
	}
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "sqlite"
	}
	return Config{
		Mode:        mode,
		Driver:      driver,
		DatabaseURL: os.Getenv("DATABASE_URL"),
		SQLitePath:  path,
		AccountID:   os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		DatabaseID:  os.Getenv("CLOUDFLARE_D1_DATABASE_ID"),
		APIToken:    os.Getenv("CLOUDFLARE_API_TOKEN"),
		BaseURL:     os.Getenv("CLOUDFLARE_API_BASE_URL"),
	}
}

// D1QueryResult represents a single query result from Cloudflare D1 API or local DB.
type D1QueryResult struct {
	Results []map[string]interface{} `json:"results"`
	Success bool                     `json:"success"`
	Meta    map[string]interface{}   `json:"meta"`
}

// D1Response represents the top-level response from Cloudflare D1 API or local DB.
type D1Response struct {
	Result  []D1QueryResult `json:"result"`
	Success bool            `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// Client abstracts local database and Cloudflare D1 REST operations.
type Client struct {
	cfg        Config
	httpClient *http.Client
	db         *sql.DB // Opened for local mode
}

// NewClient initializes a new Client.
func NewClient(cfg Config) (*Client, error) {
	c := &Client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	if cfg.Mode == "d1" {
		if cfg.AccountID == "" || cfg.DatabaseID == "" || cfg.APIToken == "" {
			return nil, fmt.Errorf("d1 mode requires CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_D1_DATABASE_ID, and CLOUDFLARE_API_TOKEN")
		}
	}

	return c, nil
}

// AttachDB attaches a pre-opened *sql.DB for local mode execution.
func (c *Client) AttachDB(db *sql.DB) {
	c.db = db
}

// ExecQuery executes a raw SQL query against D1 HTTP API or local DB.
func (c *Client) ExecQuery(ctx context.Context, sqlQuery string, params ...interface{}) (*D1Response, error) {
	if c.cfg.Mode == "local" && c.db != nil {
		return c.execLocal(ctx, sqlQuery, params...)
	}

	if c.cfg.Mode == "local" {
		// Mock/noop local execution when db pointer is not explicitly attached
		return &D1Response{
			Success: true,
			Result: []D1QueryResult{
				{
					Success: true,
					Results: []map[string]interface{}{},
					Meta:    map[string]interface{}{"mode": "local_mock"},
				},
			},
		}, nil
	}

	// Online D1 HTTP API execution
	baseURL := c.cfg.BaseURL
	if baseURL == "" {
		baseURL = "https://api.cloudflare.com/client/v4"
	}
	url := fmt.Sprintf("%s/accounts/%s/d1/database/%s/query", baseURL, c.cfg.AccountID, c.cfg.DatabaseID)

	payload := map[string]interface{}{
		"sql":    sqlQuery,
		"params": params,
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create d1 request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.cfg.APIToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("d1 request failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read d1 response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("d1 api HTTP %d: %s", resp.StatusCode, string(respBytes))
	}

	var d1Resp D1Response
	if err := json.Unmarshal(respBytes, &d1Resp); err != nil {
		return nil, fmt.Errorf("failed to parse d1 response: %w", err)
	}

	if !d1Resp.Success && len(d1Resp.Errors) > 0 {
		return &d1Resp, fmt.Errorf("d1 query error: %s (code %d)", d1Resp.Errors[0].Message, d1Resp.Errors[0].Code)
	}

	return &d1Resp, nil
}

// execLocal executes query against attached local *sql.DB.
func (c *Client) execLocal(ctx context.Context, sqlQuery string, params ...interface{}) (*D1Response, error) {
	trimmed := strings.TrimSpace(strings.ToUpper(sqlQuery))

	if strings.HasPrefix(trimmed, "SELECT") || strings.HasPrefix(trimmed, "PRAGMA") {
		rows, err := c.db.QueryContext(ctx, sqlQuery, params...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		cols, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		results := make([]map[string]interface{}, 0)
		for rows.Next() {
			scanArgs := make([]interface{}, len(cols))
			values := make([]interface{}, len(cols))
			for i := range values {
				scanArgs[i] = &values[i]
			}

			if err := rows.Scan(scanArgs...); err != nil {
				return nil, err
			}

			rowMap := make(map[string]interface{})
			for i, col := range cols {
				val := values[i]
				if b, ok := val.([]byte); ok {
					rowMap[col] = string(b)
				} else {
					rowMap[col] = val
				}
			}
			results = append(results, rowMap)
		}

		return &D1Response{
			Success: true,
			Result: []D1QueryResult{
				{
					Success: true,
					Results: results,
				},
			},
		}, nil
	}

	// Non-select queries (INSERT, UPDATE, DELETE, CREATE, etc.)
	res, err := c.db.ExecContext(ctx, sqlQuery, params...)
	if err != nil {
		return nil, err
	}

	lastID, _ := res.LastInsertId()
	affected, _ := res.RowsAffected()

	return &D1Response{
		Success: true,
		Result: []D1QueryResult{
			{
				Success: true,
				Results: []map[string]interface{}{},
				Meta: map[string]interface{}{
					"last_row_id":  lastID,
					"rows_written": affected,
				},
			},
		},
	}, nil
}
