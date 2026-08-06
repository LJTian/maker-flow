# telemetry-posthog

Lightweight, dependency-free HTTP client wrapper for pushing events to PostHog (or compatible analytics platforms).
Helps indie developers track product usage and user flows easily without pulling in heavy SDKs.

## 💰 Pricing / Free Tier
- **Free Tier**: 1,000,000 events per month completely free (Cloud hosted).
- **Self-Hosting**: Free and open-source if you deploy it yourself.
- PostHog's free tier is extraordinarily generous and typically sufficient for MVP validation and early growth.

## Agent Usage Instructions

1. **Copy** `telemetry.go` into `<product-root>/<app-id>/internal/telemetry/`.
2. **Initialize in `main.go`**:
   ```go
   import "your_app/internal/telemetry"

   analytics := telemetry.NewClient(telemetry.Config{
       APIKey:  os.Getenv("POSTHOG_API_KEY"),
       BaseURL: "https://app.posthog.com", // Or self-hosted endpoint
   })
   ```
3. **Usage Example**:
   ```go
   // Capture a user action
   go func() {
       err := analytics.Capture(userID, "user_signed_up", map[string]interface{}{
           "plan": "premium",
           "source": "twitter",
       })
       if err != nil {
           log.Printf("Telemetry error: %v", err)
       }
   }()
   ```
