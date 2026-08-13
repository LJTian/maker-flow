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
| REST API, Gin | `go-api` | `retry-backoff`, `circuit-breaker`, `singleflight-cache`, `persistence-d1` |
| Tables / DB / persistence | `go-api` | `persistence-d1` (Cloudflare D1 online + Docker local) |
| CLI / command-line tool | `go-cli` | `retry-backoff`, `worker-pool` |
| Background jobs / multi-goroutine consumers | `go-worker` | `worker-pool`, `pipeline`, `retry-backoff` |
| Browser UI / SPA / dashboard | `web-vite` | — (optional snippets in `src/lib/`) |
| API + browser UI | `go-api` + `web-vite` | layout [`web-api`](../templates/layouts/web-api/); patterns as needed |

- **DB / Durable Storage:** Add `persistence-d1` (Cloudflare D1 online + Docker local) to your selection (only use this if state is required).
- **User Login / Auth:** Add `auth-oauth-jwt` for login / OAuth (Google, GitHub, WeChat; Apple is a thin scaffold) / JWT. This is **login only** — not WeChat Pay.
- **Payments / Subscriptions:** Add `payment-lemonsqueezy` **only** for Lemon Squeezy MoR checkout + webhook verification. **MUST NOT** pick it for native WeChat Pay, Alipay, or Stripe — those patterns are not in the catalog yet (see `docs/roadmap.md`).
- **Email/Notifications:** Add `notify-email` for Resend transactional email (Resend only — no generic SMTP pattern).
- **Storage/Uploads:** Add `storage-s3` for S3-compatible object storage (AWS S3 / R2 / MinIO via `aws-sdk-go-v2` custom endpoint).
- **AI/LLMs:** Add `ai-llm-client` for **OpenAI-compatible** streaming (ChatGPT or gateways that speak `/v1/chat/completions`; not a native Anthropic SDK).
- **Cron Jobs:** Add `cron-scheduler` if the user explicitly mentions scheduled tasks, recurring jobs, or daily/hourly cron.
- **Analytics:** Add `telemetry-posthog` if the user asks for user tracking, telemetry, or analytics (PostHog).
- **Security/Rate Limiting:** Add `rate-limiter` if the user wants to prevent abuse, protect APIs, or limit request frequency.
- **Resilience:** Add `circuit-breaker` or `retry-backoff` if high reliability against external systems is requested.
- **Concurrency:** Add `worker-pool` or `pipeline` if high-throughput background processing is needed.

Multi-app examples: `go-api` + `go-worker` (sync API + async consume); `go-api` + `go-cli` (service + ops commands); `go-api` + `web-vite` (API + browser UI — copy [`templates/layouts/web-api/`](../templates/layouts/web-api/) root files).

When matching `go-api` + `web-vite`, **MUST** mention layout `web-api` in the selection output (or explain why a custom root compose replaces it).

## MUST NOT

- MUST NOT skip the catalog and invent scaffolding
- MUST NOT run before the PRO is confirmed
- MUST NOT deploy a pattern as a standalone public service
- MUST NOT force-fit unrelated app shapes (each app MUST map to a PRO responsibility)
