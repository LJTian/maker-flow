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

// DB is the decoupled abstract database interface.
type DB interface {
	ExecQuery(ctx context.Context, sqlQuery string, params ...interface{}) (*D1Response, error)
	Close() error
}

// Config defines configuration for database initialization.
type Config struct {
	Mode        string // "local" or "d1"
	Driver      string // Driver name for database/sql in local mode (e.g. "sqlite", "postgres")
	DatabaseURL string // Connection DSN / path for local DB
	AccountID   string // Cloudflare Account ID (d1 mode)
	DatabaseID  string // Cloudflare D1 Database ID (d1 mode)
	APIToken    string // Cloudflare API Token (d1 mode)
	BaseURL     string // Optional base URL override (for testing)
}

// ConfigFromEnv reads configuration from environment variables.
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
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = path
	}
	return Config{
		Mode:        mode,
		Driver:      driver,
		DatabaseURL: dbURL,
		AccountID:   os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		DatabaseID:  os.Getenv("CLOUDFLARE_D1_DATABASE_ID"),
		APIToken:    os.Getenv("CLOUDFLARE_API_TOKEN"),
		BaseURL:     os.Getenv("CLOUDFLARE_API_BASE_URL"),
	}
}

// D1QueryResult represents a single query result.
type D1QueryResult struct {
	Results []map[string]interface{} `json:"results"`
	Success bool                     `json:"success"`
	Meta    map[string]interface{}   `json:"meta"`
}

// D1Response represents top-level query response across all drivers.
type D1Response struct {
	Result  []D1QueryResult `json:"result"`
	Success bool            `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// NewDB acts as the factory constructor, returning the DB interface.
func NewDB(cfg Config) (DB, error) {
	switch cfg.Mode {
	case "d1":
		return NewD1Driver(cfg)
	case "local":
		return NewLocalSQLDriver(cfg)
	default:
		return NewLocalSQLDriver(cfg)
	}
}

// --- D1Driver Implementation (Cloudflare D1 REST API) ---

type D1Driver struct {
	cfg        Config
	httpClient *http.Client
}

func NewD1Driver(cfg Config) (*D1Driver, error) {
	if cfg.AccountID == "" || cfg.DatabaseID == "" || cfg.APIToken == "" {
		return nil, fmt.Errorf("d1 mode requires CLOUDFLARE_ACCOUNT_ID, CLOUDFLARE_D1_DATABASE_ID, and CLOUDFLARE_API_TOKEN")
	}
	return &D1Driver{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}, nil
}

func (d *D1Driver) ExecQuery(ctx context.Context, sqlQuery string, params ...interface{}) (*D1Response, error) {
	baseURL := d.cfg.BaseURL
	if baseURL == "" {
		baseURL = "https://api.cloudflare.com/client/v4"
	}
	url := fmt.Sprintf("%s/accounts/%s/d1/database/%s/query", baseURL, d.cfg.AccountID, d.cfg.DatabaseID)

	payload := map[string]interface{}{
		"sql":    sqlQuery,
		"params": params,
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create d1 request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+d.cfg.APIToken)

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("d1 request failed: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("d1 api HTTP %d: %s", resp.StatusCode, string(respBytes))
	}

	var d1Resp D1Response
	if err := json.Unmarshal(respBytes, &d1Resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !d1Resp.Success && len(d1Resp.Errors) > 0 {
		return &d1Resp, fmt.Errorf("d1 query error: %s (code %d)", d1Resp.Errors[0].Message, d1Resp.Errors[0].Code)
	}

	return &d1Resp, nil
}

func (d *D1Driver) Close() error {
	return nil
}

// --- LocalSQLDriver Implementation (Docker / Local database/sql) ---

type LocalSQLDriver struct {
	cfg Config
	db  *sql.DB
}

func NewLocalSQLDriver(cfg Config) (*LocalSQLDriver, error) {
	var db *sql.DB
	var err error
	if cfg.Driver != "" && cfg.DatabaseURL != "" {
		db, err = sql.Open(cfg.Driver, cfg.DatabaseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to open local db (%s): %w", cfg.Driver, err)
		}
	}
	return &LocalSQLDriver{
		cfg: cfg,
		db:  db,
	}, nil
}

func NewLocalSQLDriverWithDB(db *sql.DB) *LocalSQLDriver {
	return &LocalSQLDriver{
		db: db,
	}
}

func (l *LocalSQLDriver) ExecQuery(ctx context.Context, sqlQuery string, params ...interface{}) (*D1Response, error) {
	if l.db == nil {
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

	trimmed := strings.TrimSpace(strings.ToUpper(sqlQuery))

	if strings.HasPrefix(trimmed, "SELECT") || strings.HasPrefix(trimmed, "PRAGMA") {
		rows, err := l.db.QueryContext(ctx, sqlQuery, params...)
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

	// Non-SELECT queries (INSERT, UPDATE, DELETE, CREATE, etc.)
	res, err := l.db.ExecContext(ctx, sqlQuery, params...)
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

func (l *LocalSQLDriver) Close() error {
	if l.db != nil {
		return l.db.Close()
	}
	return nil
}
