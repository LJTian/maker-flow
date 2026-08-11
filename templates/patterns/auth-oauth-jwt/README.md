# auth-oauth-jwt

Unified authentication pattern for Maker Flow MVPs.
Combines `markbates/goth` (Google, GitHub, WeChat; Apple is a thin scaffold) and `golang-jwt/jwt/v5` (stateless token issuance).

**Scope (honest):** OAuth **login** + JWT. **Not** WeChat Pay / Alipay. Persisting users to DB is still a README TODO — see [`docs/roadmap.md`](../../../docs/roadmap.md).

## Agent Usage Instructions

1. **Copy** `oauth.go` and `jwt.go` into `<product-root>/<app-id>/internal/auth/`.
2. **Install deps**: `go get github.com/markbates/goth github.com/golang-jwt/jwt/v5`
3. **Initialize in `main.go`**:
   ```go
   import "your_app/internal/auth"

   auth.InitOAuth([]auth.Config{
       {Provider: "google", ClientID: os.Getenv("GOOGLE_KEY"), ClientSecret: os.Getenv("GOOGLE_SECRET"), CallbackURL: "http://localhost:8080/auth/google/callback"},
   })
   ```
4. **Setup Routes**:
   ```go
   authGroup := r.Group("/auth")
   auth.SetupRoutes(authGroup, func(c *gin.Context, user goth.User) {
       // user.Email, user.Name, user.AvatarURL
       // TODO: Save to DB or find user...
       
       token, _ := auth.GenerateToken(user.Email, time.Hour*24, os.Getenv("JWT_SECRET"))
       c.JSON(200, gin.H{"token": token})
   })
   ```
5. **Protect Private Routes**:
   ```go
   api := r.Group("/api")
   api.Use(auth.Middleware(os.Getenv("JWT_SECRET")))
   api.GET("/me", func(c *gin.Context) {
       userID := auth.GetUserID(c)
       c.JSON(200, gin.H{"me": userID})
   })
   ```
