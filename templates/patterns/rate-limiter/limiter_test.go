package ratelimit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func TestRateLimiterMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	// Allow 2 requests per second, burst of 2
	limiter := NewIPRateLimiter(rate.Limit(2), 2)
	
	r := gin.New()
	r.Use(Middleware(limiter))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Make 3 requests. First 2 should pass, 3rd should fail
	for i := 1; i <= 3; i++ {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "192.168.1.1:1234"
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if i <= 2 {
			if w.Code != http.StatusOK {
				t.Errorf("Request %d: Expected status 200, got %d", i, w.Code)
			}
		} else {
			if w.Code != http.StatusTooManyRequests {
				t.Errorf("Request %d: Expected status 429, got %d", i, w.Code)
			}
		}
	}
}
