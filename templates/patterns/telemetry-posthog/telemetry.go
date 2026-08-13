package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// Tracker is the abstract telemetry event capture interface.
type Tracker interface {
	Capture(distinctID, eventName string, properties map[string]interface{}) error
}

// Config holds analytics configuration.
type Config struct {
	Mode    string // "posthog", "console", or "mock"
	APIKey  string
	BaseURL string // Default: https://app.posthog.com
}

// ConfigFromEnv reads telemetry config from environment.
func ConfigFromEnv() Config {
	mode := os.Getenv("TELEMETRY_MODE")
	if mode == "" {
		if os.Getenv("POSTHOG_API_KEY") != "" {
			mode = "posthog"
		} else {
			mode = "console"
		}
	}
	baseURL := os.Getenv("POSTHOG_BASE_URL")
	if baseURL == "" {
		baseURL = "https://app.posthog.com"
	}
	return Config{
		Mode:    mode,
		APIKey:  os.Getenv("POSTHOG_API_KEY"),
		BaseURL: baseURL,
	}
}

// NewTracker is the factory constructor returning the Tracker interface.
func NewTracker(cfg Config) (Tracker, error) {
	switch cfg.Mode {
	case "posthog":
		return NewPostHogTracker(cfg)
	case "console":
		return NewConsoleTracker(), nil
	case "mock":
		return NewMockTracker(), nil
	default:
		return NewConsoleTracker(), nil
	}
}

// --- PostHogTracker Implementation ---

type PostHogTracker struct {
	config Config
	client *http.Client
}

type Event struct {
	APIKey     string                 `json:"api_key"`
	Event      string                 `json:"event"`
	DistinctID string                 `json:"distinct_id"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Timestamp  time.Time              `json:"timestamp"`
}

func NewPostHogTracker(cfg Config) (*PostHogTracker, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("posthog mode requires POSTHOG_API_KEY")
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://app.posthog.com"
	}
	return &PostHogTracker{
		config: cfg,
		client: &http.Client{Timeout: 5 * time.Second},
	}, nil
}

func (p *PostHogTracker) Capture(distinctID, eventName string, properties map[string]interface{}) error {
	evt := Event{
		APIKey:     p.config.APIKey,
		Event:      eventName,
		DistinctID: distinctID,
		Properties: properties,
		Timestamp:  time.Now().UTC(),
	}

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/capture/", p.config.BaseURL), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("telemetry capture failed with status %d", resp.StatusCode)
	}

	return nil
}

// --- ConsoleTracker Implementation (Local Dev) ---

type ConsoleTracker struct{}

func NewConsoleTracker() *ConsoleTracker {
	return &ConsoleTracker{}
}

func (c *ConsoleTracker) Capture(distinctID, eventName string, properties map[string]interface{}) error {
	log.Printf("[TELEMETRY CONSOLE MOCK] DistinctID: %s | Event: %s | Properties: %v", distinctID, eventName, properties)
	return nil
}

// --- MockTracker Implementation (Unit Testing) ---

type CapturedEvent struct {
	DistinctID string
	EventName  string
	Properties map[string]interface{}
}

type MockTracker struct {
	mu     sync.Mutex
	Events []CapturedEvent
}

func NewMockTracker() *MockTracker {
	return &MockTracker{Events: make([]CapturedEvent, 0)}
}

func (m *MockTracker) Capture(distinctID, eventName string, properties map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Events = append(m.Events, CapturedEvent{
		DistinctID: distinctID,
		EventName:  eventName,
		Properties: properties,
	})
	return nil
}
