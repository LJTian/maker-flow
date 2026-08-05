# Template matching skill

**English** · [简体中文](template-matching.zh-CN.md)

**Step:** 4 — AI matches templates from the PRO  
**Depends on:** `templates/CATALOG.md` → `templates/index.md` → `templates/patterns/index.md`

## Goal

1. Select **1–N apps** (`templates/apps/`) — compose shapes the PRO needs, e.g. API + worker  
2. Select **0–N patterns** (`templates/patterns/`, by tags)  
3. List dependent **image** tags (union across selected apps)

## Input

- Confirmed PRO
- `templates/CATALOG.md`

## Output

```markdown
## Selected templates
- **Apps**:
  - go-api → templates/apps/go-api
  - go-worker → templates/apps/go-worker
- **Image deps**: go-builder + go-runtime
- **Patterns**: retry-backoff, worker-pool (may be empty)
- **Product layout**: `<product-root>/{api,worker}/` (or describe otherwise)
- **Rationale**: …
```

## Matching rules

| PRO signal | App | Common patterns |
|------------|-----|-----------------|
| REST API, Gin | `go-api` | `retry-backoff`, `circuit-breaker`, `singleflight-cache`, `persistence-sqlx` |
| Tables / DB / persistence | `go-api` | `persistence-sqlx` (sqlite \| postgres \| mysql via `DB_DRIVER`) |
| CLI / command-line tool | `go-cli` | `retry-backoff`, `worker-pool` |
| Background jobs / multi-goroutine consumers | `go-worker` | `worker-pool`, `pipeline`, `retry-backoff` |
| Browser UI / SPA / dashboard | `web-vite` | — (optional snippets in `src/lib/`) |
| API + browser UI | `go-api` + `web-vite` | layout [`web-api`](../templates/layouts/web-api/); patterns as needed |

- **DB / Durable Storage:** Add `persistence-sqlx` to your selection (only use this if state is required).
- **User Login / Auth:** Add `auth-oauth-jwt` to your selection if the user wants login, OAuth (Google, Apple, WeChat, GitHub), or JWT.
- **Payments / Subscriptions:** Add `payment-lemonsqueezy` to your selection if the user wants to collect payments, especially for individual developers (supports Alipay, WeChat Pay, global MoR).
- **Email/Notifications:** Add `notify-email` if the user wants to send transactional emails (welcome, reset password, etc.).
- **Storage/Uploads:** Add `storage-s3` if the user needs to upload files, avatars, or use S3/R2.
- **AI/LLMs:** Add `ai-llm-client` if the user wants to integrate ChatGPT, Claude, or any LLM streaming features.
- **Cron Jobs:** Add `cron-scheduler` if the user explicitly mentions scheduled tasks, recurring jobs, or daily/hourly cron.
- **Analytics:** Add `telemetry-posthog` if the user asks for user tracking, telemetry, or analytics (PostHog).
- **Resilience:** Add `circuit-breaker` or `retry-backoff` if high reliability against external systems is requested.
- **Concurrency:** Add `worker-pool` or `pipeline` if high-throughput background processing is needed.

Multi-app examples: `go-api` + `go-worker` (sync API + async consume); `go-api` + `go-cli` (service + ops commands); `go-api` + `web-vite` (API + browser UI — copy [`templates/layouts/web-api/`](../templates/layouts/web-api/) root files).

When matching `go-api` + `web-vite`, **MUST** mention layout `web-api` in the selection output (or explain why a custom root compose replaces it).

## MUST NOT

- MUST NOT skip the catalog and invent scaffolding
- MUST NOT run before the PRO is confirmed
- MUST NOT deploy a pattern as a standalone public service
- MUST NOT force-fit unrelated app shapes (each app MUST map to a PRO responsibility)
