# persistence-d1

**English** · [简体中文](README.zh-CN.md)

Dual-mode persistence driver: **Local Docker / SQLite** during local development, and **Cloudflare D1 HTTP API** during online production deployment.

Tags: `db` `d1` `cloudflare` `sqlite` `persist` `docker`

## Strategy

- **Local (`DB_MODE=local`)**: Uses local SQLite file or Docker container volume (`SQLITE_PATH=/data/app.db`). Zero external dependencies or credentials required.
- **Online (`DB_MODE=d1`)**: Interacts directly with Cloudflare D1 REST API v4 using `CLOUDFLARE_ACCOUNT_ID`, `CLOUDFLARE_D1_DATABASE_ID`, and `CLOUDFLARE_API_TOKEN`.

## Env

| Variable | Required | Notes |
|----------|----------|-------|
| `DB_MODE` | No (default `local`) | `local` \| `d1` |
| `SQLITE_PATH` | Local mode | Used when `DB_MODE=local`; default `/data/app.db` |
| `CLOUDFLARE_ACCOUNT_ID` | Yes (d1 mode) | Cloudflare Account ID |
| `CLOUDFLARE_D1_DATABASE_ID` | Yes (d1 mode) | Cloudflare D1 Database ID |
| `CLOUDFLARE_API_TOKEN` | Yes (d1 mode) | Cloudflare API Token (D1 Edit permissions) |

## API Surface

- `ConfigFromEnv()` → Reads environment variables.
- `NewClient(cfg)` → Returns `*Client`.
- `client.ExecQuery(ctx, sql, params...)` → Unified interface executing locally or via Cloudflare D1 HTTP API.

## Free Tier & Limits

- **Cloudflare D1 Free Tier**: 5 Million reads / day, 100,000 writes / day, 5 GB storage.
- **Cost**: $0.00 / mo within free tier boundaries.
