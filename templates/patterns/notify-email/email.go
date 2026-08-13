package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// EmailSender is the abstract email notification interface.
type EmailSender interface {
	Send(to []string, subject, htmlBody string) error
}

// Config holds email notification settings.
type Config struct {
	Mode   string // "resend", "console", or "mock"
	APIKey string // For Resend
	From   string // Sender address
}

// ConfigFromEnv reads email config from environment variables.
func ConfigFromEnv() Config {
	mode := os.Getenv("EMAIL_MODE")
	if mode == "" {
		if os.Getenv("RESEND_API_KEY") != "" {
			mode = "resend"
		} else {
			mode = "console"
		}
	}
	return Config{
		Mode:   mode,
		APIKey: os.Getenv("RESEND_API_KEY"),
		From:   os.Getenv("EMAIL_FROM"),
	}
}

// NewEmailSender creates an EmailSender based on configuration.
func NewEmailSender(cfg Config) (EmailSender, error) {
	switch cfg.Mode {
	case "resend":
		return NewResendSender(cfg)
	case "console":
		return NewConsoleSender(cfg), nil
	case "mock":
		return NewMockSender(), nil
	default:
		return NewConsoleSender(cfg), nil
	}
}

// --- ResendSender Implementation ---

type ResendSender struct {
	cfg    Config
	client *http.Client
}

func NewResendSender(cfg Config) (*ResendSender, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("resend mode requires RESEND_API_KEY")
	}
	if cfg.From == "" {
		cfg.From = "onboarding@resend.dev"
	}
	return &ResendSender{
		cfg:    cfg,
		client: &http.Client{},
	}, nil
}

type emailPayload struct {
	To      []string `json:"to"`
	From    string   `json:"from"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html"`
}

func (r *ResendSender) Send(to []string, subject, htmlBody string) error {
	payload := emailPayload{
		To:      to,
		From:    r.cfg.From,
		Subject: subject,
		HTML:    htmlBody,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+r.cfg.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("resend api error: HTTP %d", resp.StatusCode)
	}

	return nil
}

// --- ConsoleSender Implementation (Local Dev) ---

type ConsoleSender struct {
	cfg Config
}

func NewConsoleSender(cfg Config) *ConsoleSender {
	return &ConsoleSender{cfg: cfg}
}

func (c *ConsoleSender) Send(to []string, subject, htmlBody string) error {
	log.Printf("[EMAIL CONSOLE MOCK] To: %v | From: %s | Subject: %s | Body: %s", to, c.cfg.From, subject, htmlBody)
	return nil
}

// --- MockSender Implementation (Unit Testing) ---

type SentEmail struct {
	To      []string
	Subject string
	Body    string
}

type MockSender struct {
	mu     sync.Mutex
	Emails []SentEmail
}

func NewMockSender() *MockSender {
	return &MockSender{Emails: make([]SentEmail, 0)}
}

func (m *MockSender) Send(to []string, subject, htmlBody string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Emails = append(m.Emails, SentEmail{To: to, Subject: subject, Body: htmlBody})
	return nil
}
