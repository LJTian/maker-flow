# notify-email

专为独立开发者设计的轻量级邮件通知模式，默认采用 [Resend](https://resend.com/)。
提供了极其简便的事务性邮件（欢迎信、密码重置等）发送接口。

## 💰 成本与免费额度 (Pricing)
- **免费额度**：每月免费发送 3,000 封（每日上限 100 封）。
- **超额成本**：超出后基础套餐为 $20/月（含 5 万封）。
- 极其适合 MVP 前期零成本启动，前提是您需要拥有一个自定义域名并完成 DNS 验证。

## Agent 使用说明

1. **复制** `email.go` 到 `<产品根>/<app-id>/internal/notify/` 目录下。
2. **在 `main.go` 中初始化**:
   ```go
   import "your_app/internal/notify"

   emailClient := notify.NewClient(notify.ResendConfig{
       APIKey: os.Getenv("RESEND_API_KEY"),
       From:   "Acme <onboarding@resend.dev>",
   })
   ```
3. **调用示例**:
   ```go
   err := emailClient.Send(
       []string{"user@example.com"},
       "欢迎加入我们！",
       "<h1>欢迎！</h1><p>感谢您的注册。</p>",
   )
   ```
