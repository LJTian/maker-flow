package telemetry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration for PostHog or compatible analytics
type Config struct {
	APIKey  string
	BaseURL string // Default: https://app.posthog.com
}

// Client handles pushing events
type Client struct {
	config Config
	client *http.Client
}

// Event represents a telemetry event
type Event struct {
	APIKey     string                 `json:"api_key"`
	Event      string                 `json:"event"`
	DistinctID string                 `json:"distinct_id"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Timestamp  time.Time              `json:"timestamp"`
}

// NewClient initializes the telemetry client
func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://app.posthog.com"
	}
	return &Client{
		config: config,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

// Capture sends a single event to PostHog
func (c *Client) Capture(distinctID, eventName string, properties map[string]interface{}) error {
	evt := Event{
		APIKey:     c.config.APIKey,
		Event:      eventName,
		DistinctID: distinctID,
		Properties: properties,
		Timestamp:  time.Now().UTC(),
	}

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/capture/", c.config.BaseURL), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("telemetry capture failed with status %d", resp.StatusCode)
	}

	return nil
}
