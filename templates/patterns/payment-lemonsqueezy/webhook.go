package ls

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// PaymentVerifier is the decoupled webhook verification interface.
type PaymentVerifier interface {
	Verify(signature string, body []byte) bool
}

// Config defines settings for payment webhook verification.
type Config struct {
	Mode   string // "lemonsqueezy" or "mock"
	Secret string
}

// ConfigFromEnv reads config from environment variables.
func ConfigFromEnv() Config {
	mode := os.Getenv("PAYMENT_MODE")
	if mode == "" {
		if os.Getenv("LEMONSQUEEZY_WEBHOOK_SECRET") != "" {
			mode = "lemonsqueezy"
		} else {
			mode = "mock"
		}
	}
	return Config{
		Mode:   mode,
		Secret: os.Getenv("LEMONSQUEEZY_WEBHOOK_SECRET"),
	}
}

// NewPaymentVerifier factory constructor.
func NewPaymentVerifier(cfg Config) PaymentVerifier {
	switch cfg.Mode {
	case "lemonsqueezy":
		return NewLemonSqueezyVerifier(cfg.Secret)
	case "mock":
		return NewMockPaymentVerifier()
	default:
		return NewMockPaymentVerifier()
	}
}

// --- LemonSqueezyVerifier Implementation ---

type LemonSqueezyVerifier struct {
	secret string
}

func NewLemonSqueezyVerifier(secret string) *LemonSqueezyVerifier {
	return &LemonSqueezyVerifier{secret: secret}
}

func (v *LemonSqueezyVerifier) Verify(signature string, body []byte) bool {
	if signature == "" || v.secret == "" {
		return false
	}
	mac := hmac.New(sha256.New, []byte(v.secret))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedMAC))
}

// --- MockPaymentVerifier Implementation (Local Dev) ---

type MockPaymentVerifier struct{}

func NewMockPaymentVerifier() *MockPaymentVerifier {
	return &MockPaymentVerifier{}
}

func (v *MockPaymentVerifier) Verify(signature string, body []byte) bool {
	return true // Always accept webhooks in local dev / mock mode
}

// OrderEvent represents a simplified Lemon Squeezy order webhook payload
type OrderEvent struct {
	Meta struct {
		EventName  string `json:"event_name"`
		CustomData struct {
			UserID string `json:"user_id"`
		} `json:"custom_data"`
	} `json:"meta"`
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			StoreID        int    `json:"store_id"`
			Identifier     string `json:"identifier"`
			OrderNumber    int    `json:"order_number"`
			UserName       string `json:"user_name"`
			UserEmail      string `json:"user_email"`
			Currency       string `json:"currency"`
			TotalFormatted string `json:"total_formatted"`
			Status         string `json:"status"`
		} `json:"attributes"`
	} `json:"data"`
}

// VerifyWebhook validates the webhook signature using PaymentVerifier
func VerifyWebhook(verifier PaymentVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("X-Signature")
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}

		if !verifier.Verify(signature, body) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid signature"})
			return
		}

		// Save payload for subsequent handlers
		c.Set("ls_payload", body)
		c.Next()
	}
}

// ParseEvent parses the raw payload extracted from the context
func ParseEvent(payload []byte) (*OrderEvent, error) {
	var event OrderEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
