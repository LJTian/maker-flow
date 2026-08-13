package persistd1

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
)

type mockTripper struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (m *mockTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.fn(req)
}

func TestConfigFromEnvDefaults(t *testing.T) {
	os.Unsetenv("DB_MODE")
	os.Unsetenv("SQLITE_PATH")

	cfg := ConfigFromEnv()
	if cfg.Mode != "local" {
		t.Errorf("expected default Mode 'local', got '%s'", cfg.Mode)
	}
	if cfg.SQLitePath != "/data/app.db" {
		t.Errorf("expected default SQLitePath '/data/app.db', got '%s'", cfg.SQLitePath)
	}
}

func TestNewClientD1MissingEnvs(t *testing.T) {
	cfg := Config{
		Mode: "d1",
	}
	_, err := NewClient(cfg)
	if err == nil {
		t.Error("expected error when D1 environment variables are missing, got nil")
	}
}

func TestExecQueryMockD1API(t *testing.T) {
	cfg := Config{
		Mode:       "d1",
		AccountID:  "acc_123",
		DatabaseID: "db_123",
		APIToken:   "mock_token",
	}

	client, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	client.httpClient = &http.Client{
		Transport: &mockTripper{
			fn: func(req *http.Request) (*http.Response, error) {
				if req.Header.Get("Authorization") != "Bearer mock_token" {
					return &http.Response{
						StatusCode: http.StatusUnauthorized,
						Body:       io.NopCloser(bytes.NewBufferString("unauthorized")),
					}, nil
				}

				resp := D1Response{
					Success: true,
					Result: []D1QueryResult{
						{
							Success: true,
							Results: []map[string]interface{}{
								{"id": "1", "title": "test note"},
							},
						},
					},
				}
				bodyBytes, _ := json.Marshal(resp)
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					Body:       io.NopCloser(bytes.NewBuffer(bodyBytes)),
				}, nil
			},
		},
	}

	res, err := client.ExecQuery(context.Background(), "SELECT * FROM notes WHERE id = ?", "1")
	if err != nil {
		t.Fatalf("ExecQuery failed: %v", err)
	}

	if !res.Success {
		t.Error("expected res.Success to be true")
	}
	if len(res.Result) == 0 || len(res.Result[0].Results) == 0 {
		t.Fatal("expected query results")
	}
	if res.Result[0].Results[0]["title"] != "test note" {
		t.Errorf("expected title 'test note', got %v", res.Result[0].Results[0]["title"])
	}
}
