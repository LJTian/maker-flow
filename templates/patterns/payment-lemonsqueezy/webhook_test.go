package ls

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestVerifyWebhook(t *testing.T) {
	secret := "test_lemon_secret"
	payload := []byte(`{"meta":{"event_name":"order_created","custom_data":{"user_id":"user_001"}}}`)

	// Generate correct signature
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	correctSignature := hex.EncodeToString(mac.Sum(nil))

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/webhook", VerifyWebhook(secret), func(c *gin.Context) {
		rawPayload := c.MustGet("ls_payload").([]byte)
		event, err := ParseEvent(rawPayload)
		if err != nil {
			c.String(http.StatusBadRequest, "parse error")
			return
		}
		c.String(http.StatusOK, event.Meta.CustomData.UserID)
	})

	t.Run("Valid Signature", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(payload))
		req.Header.Set("X-Signature", correctSignature)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
		if w.Body.String() != "user_001" {
			t.Errorf("expected user_001, got %s", w.Body.String())
		}
	})

	t.Run("Invalid Signature", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(payload))
		req.Header.Set("X-Signature", "invalid_signature_string")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", w.Code)
		}
	})

	t.Run("Missing Signature", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(payload))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("expected 401, got %d", w.Code)
		}
	})
}
