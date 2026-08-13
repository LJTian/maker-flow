package storage

import (
	"context"
	"bytes"
	"testing"
)

func TestMockStorageDriver(t *testing.T) {
	s := NewMockStorageDriver()
	ctx := context.Background()

	err := s.UploadFile(ctx, "test.txt", []byte("hello storage"), "text/plain")
	if err != nil {
		t.Fatalf("Failed to upload to mock storage: %v", err)
	}

	data, err := s.DownloadFile(ctx, "test.txt")
	if err != nil {
		t.Fatalf("Failed to download from mock storage: %v", err)
	}

	if !bytes.Equal(data, []byte("hello storage")) {
		t.Errorf("Expected 'hello storage', got '%s'", string(data))
	}
}

func TestNewStorageFactory(t *testing.T) {
	ctx := context.Background()
	cfg := Config{
		Mode: "mock",
	}

	s, err := NewStorage(ctx, cfg)
	if err != nil {
		t.Fatalf("Failed to initialize storage via factory: %v", err)
	}

	if s == nil {
		t.Fatal("Storage is nil")
	}
}
