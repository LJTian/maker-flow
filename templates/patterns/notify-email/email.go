package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ResendConfig holds the configuration for sending emails via Resend.
type ResendConfig struct {
	APIKey string
	From   string
}

// EmailRequest represents the payload for sending an email.
type EmailRequest struct {
	To      []string `json:"to"`
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html"`
}

// Client handles email sending operations.
type Client struct {
	config ResendConfig
	client *http.Client
}

// NewClient initializes a new email client.
func NewClient(config ResendConfig) *Client {
	return &Client{
		config: config,
		client: &http.Client{},
	}
}

// Send sends an email using the Resend API.
func (c *Client) Send(to []string, subject, htmlBody string) error {
	reqBody := EmailRequest{
		To:      to,
		From:    c.config.From,
		Subject: subject,
		HTML:    htmlBody,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to send email: status code %d", resp.StatusCode)
	}

	return nil
}
