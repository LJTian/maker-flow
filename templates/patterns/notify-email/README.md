# notify-email

Lightweight email notification pattern designed for indie hackers using [Resend](https://resend.com/).
Provides a simple wrapper for sending transactional emails (welcome emails, password resets, etc.).

## 💰 Pricing / Free Tier
- **Free Tier**: 3,000 free emails per month (max 100/day).
- **Cost**: $20/mo for 50,000 emails once you exceed the free tier.
- Perfect for zero-cost MVP validation. Requires verifying a custom domain.

## Agent Usage Instructions

1. **Copy** `email.go` into `<product-root>/<app-id>/internal/notify/`.
2. **Initialize in `main.go`**:
   ```go
   import "your_app/internal/notify"

   emailClient := notify.NewClient(notify.ResendConfig{
       APIKey: os.Getenv("RESEND_API_KEY"),
       From:   "Acme <onboarding@resend.dev>",
   })
   ```
3. **Usage Example**:
   ```go
   err := emailClient.Send(
       []string{"user@example.com"},
       "Welcome to our platform!",
       "<h1>Welcome!</h1><p>Thanks for joining us.</p>",
   )
   ```
