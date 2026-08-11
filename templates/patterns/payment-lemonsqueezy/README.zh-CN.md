# payment-lemonsqueezy

针对个人开发者的 Lemon Squeezy 支付集成模式。
提供安全的 Webhook HMAC SHA256 签名验证，防止伪造的欺诈订单。

**范围（诚实）：** 仅 Lemon Squeezy MoR 结账 + Webhook。  
**不含：** 原生微信支付、支付宝、Stripe API — 见 [`docs/roadmap.zh-CN.md`](../../../docs/roadmap.zh-CN.md)。

## 💰 成本与免费额度 (Pricing)
- **基础月费**：0 元。没有固定的月租费，零门槛接入。
- **交易抽成**：每笔成功交易大约抽成 5% + $0.50（以实际条款为准）。
- **核心价值**：Lemon Squeezy 作为 MoR (Merchant of Record)，会替您合法处理全球各国的税务合规，并统一将税后利润打款给您。非常适合独立开发者“卖不出去不花钱，卖出去了只分利润”的模式。

## Agent 使用说明

1. **复制** `webhook.go` 到 `<产品根>/<app-id>/internal/payment/`。
2. **路由设置** (需要 Gin):
   ```go
   import "your_app/internal/payment"

   r.POST("/api/webhooks/lemonsqueezy", payment.VerifyWebhook(os.Getenv("LEMON_SECRET")), func(c *gin.Context) {
       payload := c.MustGet("ls_payload").([]byte)
       event, err := payment.ParseEvent(payload)
       if err != nil {
           c.JSON(400, gin.H{"error": "数据格式错误"})
           return
       }
       
       if event.Meta.EventName == "order_created" {
           userID := event.Meta.CustomData.UserID
           // TODO: 在数据库中将该 userID 升级为 VIP
       }
       c.JSON(200, gin.H{"status": "已接收"})
   })
   ```
3. **前端流程**: 前端将用户直接重定向至 Lemon Squeezy 结算页面。**必须**在链接后追加参数 `?checkout[custom][user_id]=<uid>`，以便 Webhook 回调时能识别是哪个用户支付的。
