package storage

import (
	"context"
	"testing"
)

func TestStorageClientInitialization(t *testing.T) {
	// Simple test to ensure the initialization doesn't panic and logic runs correctly
	cfg := Config{
		Endpoint:        "https://example.com",
		AccessKeyID:     "test",
		SecretAccessKey: "test",
		Region:          "auto",
		Bucket:          "my-bucket",
	}

	client, err := NewClient(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Failed to initialize client: %v", err)
	}

	if client == nil {
		t.Fatal("Client is nil")
	}
}
