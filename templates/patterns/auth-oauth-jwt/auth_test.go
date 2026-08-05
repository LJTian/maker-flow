package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestJWTGenerateAndMiddleware(t *testing.T) {
	secret := "test-secret"
	userID := "user_123"

	// 1. Generate Token
	token, err := GenerateToken(userID, time.Minute, secret)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}
	if token == "" {
		t.Fatal("token is empty")
	}

	// 2. Setup Gin with Middleware
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Middleware(secret))
	r.GET("/protected", func(c *gin.Context) {
		uid := GetUserID(c)
		c.String(http.StatusOK, uid)
	})

	// 3. Test Without Token (Should fail)
	req, _ := http.NewRequest(http.MethodGet, "/protected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}

	// 4. Test With Token (Should succeed and extract userID)
	req, _ = http.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != userID {
		t.Errorf("expected body %s, got %s", userID, w.Body.String())
	}
}
