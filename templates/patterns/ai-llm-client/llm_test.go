package llm

import (
	"context"
	"strings"
	"testing"
)

func TestMockLLMClientStream(t *testing.T) {
	mock := NewMockLLMClient("Hello AI World")
	ctx := context.Background()

	respChan, errChan := mock.GenerateStream(ctx, "system prompt", "user query")

	var sb strings.Builder
	for token := range respChan {
		sb.WriteString(token)
	}

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatalf("unexpected streaming error: %v", err)
		}
	default:
	}

	if sb.String() != "Hello AI World" {
		t.Errorf("expected 'Hello AI World', got '%s'", sb.String())
	}
}

func TestNewLLMClientFactoryMock(t *testing.T) {
	client, err := NewLLMClient(Config{Mode: "mock"})
	if err != nil {
		t.Fatalf("factory failed: %v", err)
	}
	if client == nil {
		t.Fatal("client is nil")
	}
}
