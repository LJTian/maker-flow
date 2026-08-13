package email

import (
	"testing"
)

func TestMockSender(t *testing.T) {
	mock := NewMockSender()
	err := mock.Send([]string{"user@example.com"}, "Welcome", "<p>Hello</p>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(mock.Emails) != 1 {
		t.Fatalf("expected 1 sent email, got %d", len(mock.Emails))
	}
	if mock.Emails[0].Subject != "Welcome" {
		t.Errorf("expected subject 'Welcome', got '%s'", mock.Emails[0].Subject)
	}
}

func TestConsoleSender(t *testing.T) {
	console := NewConsoleSender(Config{From: "noreply@example.com"})
	err := console.Send([]string{"dev@example.com"}, "Test", "Body")
	if err != nil {
		t.Errorf("console sender failed: %v", err)
	}
}

func TestNewEmailSenderFactory(t *testing.T) {
	sender, err := NewEmailSender(Config{Mode: "mock"})
	if err != nil {
		t.Fatalf("factory failed: %v", err)
	}
	if sender == nil {
		t.Fatal("sender is nil")
	}
}
