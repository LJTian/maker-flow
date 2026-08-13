package telemetry

import (
	"testing"
)

func TestMockTracker(t *testing.T) {
	mock := NewMockTracker()
	err := mock.Capture("user_123", "button_clicked", map[string]interface{}{"btn": "submit"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(mock.Events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(mock.Events))
	}
	if mock.Events[0].EventName != "button_clicked" {
		t.Errorf("expected event 'button_clicked', got '%s'", mock.Events[0].EventName)
	}
}

func TestConsoleTracker(t *testing.T) {
	console := NewConsoleTracker()
	err := console.Capture("dev_user", "page_view", nil)
	if err != nil {
		t.Errorf("console tracker failed: %v", err)
	}
}

func TestNewTrackerFactoryMock(t *testing.T) {
	tracker, err := NewTracker(Config{Mode: "mock"})
	if err != nil {
		t.Fatalf("factory failed: %v", err)
	}
	if tracker == nil {
		t.Fatal("tracker is nil")
	}
}
