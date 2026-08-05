package telemetry

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTelemetryClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/capture/" {
			t.Errorf("Expected path /capture/, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	err := client.Capture("user_123", "signed_up", map[string]interface{}{
		"plan": "free",
	})

	if err != nil {
		t.Fatalf("Capture failed: %v", err)
	}
}
