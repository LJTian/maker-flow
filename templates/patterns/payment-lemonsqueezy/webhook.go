package ls

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
			StoreID       int    `json:"store_id"`
			Identifier    string `json:"identifier"`
			OrderNumber   int    `json:"order_number"`
			UserName      string `json:"user_name"`
			UserEmail     string `json:"user_email"`
			Currency      string `json:"currency"`
			TotalFormatted string `json:"total_formatted"`
			Status        string `json:"status"`
		} `json:"attributes"`
	} `json:"data"`
}

// VerifyWebhook validates the Lemon Squeezy X-Signature and injects the raw payload into context
func VerifyWebhook(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("X-Signature")
		if signature == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing signature"})
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
			return
		}

		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write(body)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature), []byte(expectedMAC)) {
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
