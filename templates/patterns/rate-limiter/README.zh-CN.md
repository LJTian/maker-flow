# rate-limiter

基于 `golang.org/x/time/rate` 的内存级令牌桶 IP 限流器。
这是保护 MVP 免受“账单攻击 (Billing Attacks)”、暴力破解和 API 滥用（尤其是 LLM 调用或发送验证码邮件等昂贵接口）的最后一道防线。

## Agent 使用说明

1. **复制** `limiter.go` 到 `<产品根>/<app-id>/internal/ratelimit/`。
2. **安装依赖**: 
   `go get golang.org/x/time/rate github.com/gin-gonic/gin`
3. **在 `main.go` 中调用**:
   ```go
   import "your_app/internal/ratelimit"
   import "golang.org/x/time/rate"

   // 创建限流器：例如每秒补充 1 个令牌，最高并发突发 3 个
   limiter := ratelimit.NewIPRateLimiter(rate.Limit(1), 3)

   // 方案 A: 全局应用（保护所有接口）
   r.Use(ratelimit.Middleware(limiter))

   // 方案 B: 仅保护昂贵/危险的接口（推荐）
   aiGroup := r.Group("/api/ai")
   aiGroup.Use(ratelimit.Middleware(limiter))
   aiGroup.POST("/chat", handleChat)
   ```
