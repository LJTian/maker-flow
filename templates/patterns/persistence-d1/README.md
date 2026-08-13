# persistence-d1

**English** · [简体中文](README.zh-CN.md)

Decoupled database persistence pattern: **`DB` interface abstraction layer** with concrete implementations for **`D1Driver`** (Cloudflare D1 REST API) and **`LocalSQLDriver`** (Docker / local SQL database).

Tags: `db` `d1` `cloudflare` `sqlite` `persist` `docker`

## Decoupled Architecture

```
                 ┌──────────────────┐
                 │   DB Interface   │
                 └────────┬─────────┘
                          │
          ┌───────────────┴───────────────┐
          ▼                               ▼
  ┌───────────────┐               ┌───────────────┐
  │   D1Driver    │               │LocalSQLDriver │
  │(Cloudflare D1)│               │(Docker / SQL) │
  └───────────────┘               └───────────────┘
```

## Usage

```go
import "your_app/internal/persistd1"

// Read config from env (DB_MODE=local or d1)
cfg := persistd1.ConfigFromEnv()

// Factory constructor returns DB interface
db, err := persistd1.NewDB(cfg)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Execute queries cleanly against DB interface
res, err := db.ExecQuery(ctx, "SELECT * FROM users WHERE id = ?", userID)
```
