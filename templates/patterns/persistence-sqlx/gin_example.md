# Gin wiring example (agent)

Copy ideas into the product `go-api` app. Adjust package paths after module rewrite.

```go
// cmd/server/main.go
package main

import (
	_ "modernc.org/sqlite" // or pq / mysql driver

	"product/internal/persist"
	// ...
)

func main() {
	cfg, err := persist.FromEnv()
	// handle err
	db, err := persist.Open(cfg)
	// handle err
	defer db.Close()
	if err := persist.Migrate(db, cfg.Driver); err != nil {
		// fatal
	}
	store := persist.NewStore(db)
	// pass store into server / handlers
}
```

```go
// health with DB ping
func Health(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(503, gin.H{"status": "degraded", "db": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "ok"})
	}
}
```

Suggested notes routes (sketch):

| Method | Path | Action |
|--------|------|--------|
| POST | `/api/v1/notes` | Create |
| GET | `/api/v1/notes` | List |
| GET | `/api/v1/notes/:id` | Get |
| PUT | `/api/v1/notes/:id` | Update |
| DELETE | `/api/v1/notes/:id` | Delete |

Replace `notes` with PRO resources when assembling.
