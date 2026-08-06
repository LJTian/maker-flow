# payment-lemonsqueezy

Lemon Squeezy payment integration pattern for individual developers.
Provides secure Webhook HMAC SHA256 signature verification to prevent fraudulent orders.

## 💰 Pricing / Free Tier
- **Setup Fee**: $0. No fixed monthly costs.
- **Transaction Fee**: Typically 5% + 50¢ per successful transaction.
- **Value**: Acts as a Merchant of Record (MoR). They handle global sales tax compliance and payout distributions for you, meaning you only pay when you make money.

## Agent Usage Instructions

1. **Copy** `webhook.go` into `<product-root>/<app-id>/internal/payment/`.
2. **Route setup** (Requires Gin):
   ```go
   import "your_app/internal/payment"

   r.POST("/api/webhooks/lemonsqueezy", payment.VerifyWebhook(os.Getenv("LEMON_SECRET")), func(c *gin.Context) {
       payload := c.MustGet("ls_payload").([]byte)
       event, err := payment.ParseEvent(payload)
       if err != nil {
           c.JSON(400, gin.H{"error": "bad format"})
           return
       }
       
       if event.Meta.EventName == "order_created" {
           userID := event.Meta.CustomData.UserID
           // TODO: Upgrade userID to VIP in database
       }
       c.JSON(200, gin.H{"status": "received"})
   })
   ```
3. **Frontend flow**: The frontend redirects the user directly to the Lemon Squeezy Checkout URL. **MUST** append `?checkout[custom][user_id]=<uid>` so the webhook can identify who paid.
