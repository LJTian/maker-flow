package ratelimit

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// IPRateLimiter stores rate limiters for each IP address.
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter creates a new rate limiter. 
// r is the rate (events per second) and b is the burst capacity.
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	// In a real production system, you would want a background routine 
	// to clean up old IP entries to prevent memory leaks over time.
	// For MVP simplicity, this is an unbounded map.
	go i.cleanupRoutine()

	return i
}

// AddCreates a new rate limiter for an IP if it doesn't exist.
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.RLock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.RUnlock()
		i.mu.Lock()
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
		i.mu.Unlock()
		return limiter
	}
	i.mu.RUnlock()
	return limiter
}

// cleanupRoutine periodically clears the map. A simple approach for MVPs.
func (i *IPRateLimiter) cleanupRoutine() {
	for {
		time.Sleep(time.Hour)
		i.mu.Lock()
		i.ips = make(map[string]*rate.Limiter)
		i.mu.Unlock()
	}
}

// Middleware is a Gin middleware that enforces the rate limit.
func Middleware(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		l := limiter.GetLimiter(ip)
		
		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests, please try again later",
			})
			return
		}
		
		c.Next()
	}
}
