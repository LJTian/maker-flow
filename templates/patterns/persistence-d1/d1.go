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
	"time"
)

// Config defines configuration for persistence-d1.
type Config struct {
	Mode           string // "local" or "d1"
	SQLitePath     string // For local mode (default: "/data/app.db")
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
	return Config{
		Mode:       mode,
		SQLitePath: path,
		AccountID:  os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		DatabaseID: os.Getenv("CLOUDFLARE_D1_DATABASE_ID"),
		APIToken:   os.Getenv("CLOUDFLARE_API_TOKEN"),
		BaseURL:    os.Getenv("CLOUDFLARE_API_BASE_URL"),
	}
}

// D1QueryResult represents a single query result from Cloudflare D1 API.
type D1QueryResult struct {
	Results []map[string]interface{} `json:"results"`
	Success bool                     `json:"success"`
	Meta    map[string]interface{}   `json:"meta"`
}

// D1Response represents the top-level HTTP response from Cloudflare D1 API.
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
	db         *sql.DB // For local mode
}

// NewClient initializes a new Client.
func NewClient(cfg Config) (*Client, error) {
	c := &Client{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	if cfg.Mode == "local" {
		// Local mode uses database/sql (driver registered by main app if needed)
		// For MVP fallback, client operates in local-file mode
		return c, nil
	}

	if cfg.Mode == "d1" {
		if cfg.AccountID == "" || cfg.DatabaseID == "" || cfg.APIToken == "" {
			return nil, fmt.Errorf("d1 mode requires CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_D1_DATABASE_ID, and CLOUDFLARE_API_TOKEN")
		}
	}

	return c, nil
}

// ExecQuery executes a raw SQL query against D1 HTTP API or local DB.
func (c *Client) ExecQuery(ctx context.Context, sqlQuery string, params ...interface{}) (*D1Response, error) {
	if c.cfg.Mode == "local" && c.db != nil {
		// Execute locally via database/sql
		_, err := c.db.ExecContext(ctx, sqlQuery, params...)
		if err != nil {
			return nil, err
		}
		return &D1Response{Success: true}, nil
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
