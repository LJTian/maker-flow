package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/apple"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/wechat"
)

// Config represents standard OAuth setup
type Config struct {
	Provider     string
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

// InitOAuth providers (call this in main.go)
func InitOAuth(configs []Config) {
	var providers []goth.Provider
	for _, c := range configs {
		switch c.Provider {
		case "google":
			providers = append(providers, google.New(c.ClientID, c.ClientSecret, c.CallbackURL))
		case "github":
			providers = append(providers, github.New(c.ClientID, c.ClientSecret, c.CallbackURL))
		case "apple":
			providers = append(providers, apple.New(c.ClientID, c.ClientSecret, c.CallbackURL, nil))
		case "wechat":
			providers = append(providers, wechat.New(c.ClientID, c.ClientSecret, c.CallbackURL))
		}
	}
	goth.UseProviders(providers...)
}

// SetupRoutes registers the /auth/:provider and /auth/:provider/callback endpoints.
// handleSuccess is your business logic that receives the goth.User, generates a JWT, and returns it.
func SetupRoutes(r *gin.RouterGroup, handleSuccess func(*gin.Context, goth.User)) {
	r.GET("/:provider", func(c *gin.Context) {
		// Set provider into request context for gothic
		c.Request = c.Request.WithContext(c.Request.Context())
		q := c.Request.URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request.URL.RawQuery = q.Encode()

		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	r.GET("/:provider/callback", func(c *gin.Context) {
		q := c.Request.URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request.URL.RawQuery = q.Encode()

		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("oauth failed: %v", err)})
			return
		}
		handleSuccess(c, user)
	})
}
