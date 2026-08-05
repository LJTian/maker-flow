package llm

import (
	"context"
	"testing"
)

func TestLLMClientInitialization(t *testing.T) {
	cfg := Config{
		BaseURL: "https://api.openai.com/v1",
		APIKey:  "test-api-key",
		Model:   "gpt-4o",
	}

	client := NewClient(cfg)
	if client == nil {
		t.Fatal("Client is nil")
	}

	// Because GenerateStream makes network calls to the proxy, we just ensure it doesn't crash on setup
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Immediately cancel

	respChan, errChan := client.GenerateStream(ctx, "You are helpful", "Hi")
	if respChan == nil || errChan == nil {
		t.Fatal("Channels are nil")
	}
}
