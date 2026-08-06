# rate-limiter

An in-memory token bucket API rate limiter based on `golang.org/x/time/rate`. 
Essential for protecting your MVP from billing attacks, brute-force logins, or API abuse (especially on costly routes like LLMs and Email sending).

## Agent Usage Instructions

1. **Copy** `limiter.go` into `<product-root>/<app-id>/internal/ratelimit/`.
2. **Install deps**: 
   `go get golang.org/x/time/rate github.com/gin-gonic/gin`
3. **Usage in `main.go`**:
   ```go
   import "your_app/internal/ratelimit"
   import "golang.org/x/time/rate"

   // Create a limiter: e.g. 1 request per second, maximum burst of 3
   limiter := ratelimit.NewIPRateLimiter(rate.Limit(1), 3)

   // Apply globally to all routes
   r.Use(ratelimit.Middleware(limiter))

   // OR apply specifically to expensive routes
   aiGroup := r.Group("/api/ai")
   aiGroup.Use(ratelimit.Middleware(limiter))
   aiGroup.POST("/chat", handleChat)
   ```
