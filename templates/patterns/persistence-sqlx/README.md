# persistence-sqlx

**English** · [简体中文](README.zh-CN.md)

Multi-database persistence with [sqlx](https://github.com/jmoiron/sqlx). Supports **SQLite**, **PostgreSQL**, and **MySQL** via `DB_DRIVER`.

Tags: `db` `sqlx` `sqlite` `postgres` `mysql` `persist`

## When to use

PRO has tables / needs durable storage. Pair with `go-api` (copy into `internal/persist/`).

**MUST NOT** deploy this pattern alone.

## Env

| Variable | Required | Notes |
|----------|----------|-------|
| `DB_DRIVER` | No (default `sqlite`) | `sqlite` \| `postgres` \| `mysql` |
| `DATABASE_URL` | Yes* | Full DSN / URL |
| `SQLITE_PATH` | sqlite fallback | Used when `DATABASE_URL` empty; default `/data/app.db` |

## Drivers (product blank-import)

| Driver | Blank import | `sql.Open` name |
|--------|--------------|-----------------|
| SQLite (pure Go, recommended) | `_ "modernc.org/sqlite"` | `sqlite` |
| SQLite (CGO) | `_ "github.com/mattn/go-sqlite3"` | change `Open` to `sqlite3` or wrap |
| PostgreSQL | `_ "github.com/lib/pq"` or `_ "github.com/jackc/pgx/v5/stdlib"` | `postgres` (pgx registers `pgx` — set accordingly) |
| MySQL | `_ "github.com/go-sql-driver/mysql"` | `mysql` |

This pattern depends on **sqlx** only. Drivers live in the product `go.mod` / `main`.

## Agent steps

1. Select `go-api` + `persistence-sqlx`.
2. Copy package files into product app `internal/persist/` (rewrite module path).
3. Blank-import chosen driver in `cmd/server/main.go`.
4. Wire Open → Migrate → `NewStore` (see [`gin_example.md`](gin_example.md)).
5. Merge [`compose.snippet.yml`](compose.snippet.yml); set env.
6. Replace / extend `notes` with PRO entities.

## API surface

- `FromEnv()` / `Open(Config)` → `*sqlx.DB`
- `Migrate(db, driver)` — embed dialect SQL
- `Store` — example notes CRUD

## Verify

Pattern tests use CGO `mattn/go-sqlite3` (`Driver: sqlite3`). Product apps should prefer pure-Go `modernc.org/sqlite` (`DB_DRIVER=sqlite`).

```bash
cd templates/patterns/persistence-sqlx
CGO_ENABLED=1 go test ./...
```
