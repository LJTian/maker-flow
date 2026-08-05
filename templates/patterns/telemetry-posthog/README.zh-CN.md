# telemetry-posthog

无额外依赖的轻量级 HTTP 客户端封装，用于向 PostHog（或兼容的数据分析平台）推送事件。
帮助独立开发者轻松追踪产品使用情况和用户流失节点，而无需引入臃肿的官方全量 SDK。

## Agent 使用说明

1. **复制** `telemetry.go` 到 `<产品根>/<app-id>/internal/telemetry/`。
2. **在 `main.go` 中初始化**:
   ```go
   import "your_app/internal/telemetry"

   analytics := telemetry.NewClient(telemetry.Config{
       APIKey:  os.Getenv("POSTHOG_API_KEY"),
       BaseURL: "https://app.posthog.com", // 或填入您自建的实例地址
   })
   ```
3. **调用示例**:
   ```go
   // 捕获一个用户动作 (建议放入 Goroutine 异步执行以防阻塞主业务逻辑)
   go func() {
       err := analytics.Capture(userID, "user_signed_up", map[string]interface{}{
           "plan": "premium",
           "source": "twitter",
       })
       if err != nil {
           log.Printf("数据打点错误: %v", err)
       }
   }()
   ```
