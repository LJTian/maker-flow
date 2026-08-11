# auth-oauth-jwt

Maker Flow MVPs 的统一鉴权模式。
结合了 `markbates/goth`（Google、GitHub、微信；Apple 为薄骨架）与 `golang-jwt/jwt/v5`（无状态 Token 签发）。

**范围（诚实）：** OAuth **登录** + JWT。**不是**微信支付 / 支付宝。用户落库仍是 README TODO — 见 [`docs/roadmap.zh-CN.md`](../../../docs/roadmap.zh-CN.md)。

## Agent 使用说明

1. **复制** `oauth.go` 和 `jwt.go` 到 `<产品根>/<app-id>/internal/auth/`。
2. **安装依赖**: `go get github.com/markbates/goth github.com/golang-jwt/jwt/v5`
3. **在 `main.go` 中初始化**:
   ```go
   import "your_app/internal/auth"

   auth.InitOAuth([]auth.Config{
       {Provider: "google", ClientID: os.Getenv("GOOGLE_KEY"), ClientSecret: os.Getenv("GOOGLE_SECRET"), CallbackURL: "http://localhost:8080/auth/google/callback"},
   })
   ```
4. **设置路由**:
   ```go
   authGroup := r.Group("/auth")
   auth.SetupRoutes(authGroup, func(c *gin.Context, user goth.User) {
       // user.Email, user.Name, user.AvatarURL
       // TODO: 保存到数据库或查找用户...
       
       token, _ := auth.GenerateToken(user.Email, time.Hour*24, os.Getenv("JWT_SECRET"))
       c.JSON(200, gin.H{"token": token})
   })
   ```
5. **保护私有路由**:
   ```go
   api := r.Group("/api")
   api.Use(auth.Middleware(os.Getenv("JWT_SECRET")))
   api.GET("/me", func(c *gin.Context) {
       userID := auth.GetUserID(c)
       c.JSON(200, gin.H{"me": userID})
   })
   ```
