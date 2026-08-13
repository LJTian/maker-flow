# Template set catalog

**English** · [简体中文](CATALOG.zh-CN.md)

> **For AI:** Before step-4 selection, **MUST** read this file first, then the detail indexes.  
> **For humans:** One-glance view of apps / images / patterns.

---

## Overview

| Category | Count | Detail index |
|----------|:-----:|--------------|
| App templates (apps) | 4 | [index.md](index.md) · [`apps/`](apps/) |
| Image bases (images) | 2 | [images/index.md](images/index.md) |
| Pattern library (patterns) | 14 | [patterns/index.md](patterns/index.md) |
| Layouts (multi-app roots) | 1 | [layouts/index.md](layouts/index.md) |

---

## App templates (apps)

| id | Path | Tags | When to use | Image deps |
|----|------|------|-------------|------------|
| `go-api` | [`apps/go-api/`](apps/go-api/) | `go` `gin` `rest` `api` `docker` | Go + Gin REST API MVP | `go-builder` + `go-runtime` |
| `go-cli` | [`apps/go-cli/`](apps/go-cli/) | `go` `cli` `cobra` | CLI tools / subcommand scaffold | `go-builder` (+ runtime optional) |
| `go-worker` | [`apps/go-worker/`](apps/go-worker/) | `go` `worker` `concurrency` `pool` | Multi-goroutine job consumption + graceful shutdown | `go-builder` + `go-runtime` |
| `web-vite` | [`apps/web-vite/`](apps/web-vite/) | `web` `frontend` `vite` `react` `typescript` `tailwind` `spa` `docker` | Browser UI / landing / dashboard MVP | Node + Nginx (no maker-flow image bases) |

Agent: copy **1–N apps** as whole directories into the **product repo** (use subdirs when multi-app).

---

## Image fragments (images)

| id | Upstream | Path |
|----|----------|------|
| `go-builder` | `golang:1.22-alpine` | [`images/go-builder/`](images/go-builder/) |
| `go-runtime` | `alpine:3.20` | [`images/go-runtime/`](images/go-runtime/) |

Inline into app Dockerfiles when assembling — see [`images/index.md`](images/index.md). No pre-build step.
---

## Pattern library (patterns)

| id | Path | tags |
|----|------|------|
| `worker-pool` | [`patterns/worker-pool/`](patterns/worker-pool/) | `concurrency` `pool` |
| `pipeline` | [`patterns/pipeline/`](patterns/pipeline/) | `fan-in` `fan-out` |
| `singleflight-cache` | [`patterns/singleflight-cache/`](patterns/singleflight-cache/) | `cache` `singleflight` |
| `retry-backoff` | [`patterns/retry-backoff/`](patterns/retry-backoff/) | `retry` `backoff` |
| `circuit-breaker` | [`patterns/circuit-breaker/`](patterns/circuit-breaker/) | `circuit-breaker` |
| `persistence-d1` | [`patterns/persistence-d1/`](patterns/persistence-d1/) | `db` `d1` `cloudflare` `sqlite` `persist` `docker` |
| `auth-oauth-jwt` | [`patterns/auth-oauth-jwt/`](patterns/auth-oauth-jwt/) | `auth` `oauth` `jwt` `login` |
| `payment-lemonsqueezy` | [`patterns/payment-lemonsqueezy/`](patterns/payment-lemonsqueezy/) | `payment` `lemonsqueezy` `mor` |
| `notify-email` | [`patterns/notify-email/`](patterns/notify-email/) | `email` `resend` `notification` |
| `storage-s3` | [`patterns/storage-s3/`](patterns/storage-s3/) | `storage` `s3` `r2` `upload` |
| `ai-llm-client` | [`patterns/ai-llm-client/`](patterns/ai-llm-client/) | `ai` `llm` `openai` `streaming` |
| `cron-scheduler` | [`patterns/cron-scheduler/`](patterns/cron-scheduler/) | `cron` `schedule` `job` |
| `telemetry-posthog` | [`patterns/telemetry-posthog/`](patterns/telemetry-posthog/) | `telemetry` `analytics` `posthog` |
| `rate-limiter` | [`patterns/rate-limiter/`](patterns/rate-limiter/) | `rate-limit` `security` `api` |

Agent: pick **1–N apps** first, then **0–N patterns**; **copy/adapt** into the matching app in the **product repo**. Patterns are never deployed alone.

Detail → [`patterns/index.md`](patterns/index.md)

---

## Layouts (multi-app product roots)

| id | Path | Apps | When to use |
|----|------|------|-------------|
| `web-api` | [`layouts/web-api/`](layouts/web-api/) | `go-api` + `web-vite` | Browser UI + REST API root compose / env |

Copy apps into `api/` + `web/`, then copy the layout’s root files. Detail → [`layouts/index.md`](layouts/index.md). Split publish → [`skills/publish-split-web-api.md`](../skills/publish-split-web-api.md).

---

## Selection cues (Agent)

```
Need REST API?              → go-api
Need CLI?                   → go-cli
Need Background Worker?     → go-worker
Need Browser UI?            → web-vite
Need API + SPA?             → go-api + web-vite + layout web-api
Need DB / Tables?           → go-api + persistence-d1 (Cloudflare D1 online + Docker local)
Need User Login/Auth?       → go-api + auth-oauth-jwt
Need Lemon Squeezy (MoR) pay? → go-api + payment-lemonsqueezy
Need WeChat/Alipay/Stripe pay? → **not in catalog yet** — see [`docs/roadmap.md`](../docs/roadmap.md)
Need to Send Emails?        → go-api + notify-email (Resend only)
Need File Uploads?          → go-api + storage-s3 (Cloudflare R2 online + MinIO Docker local)
Need AI/LLM Features?       → go-api + ai-llm-client (OpenAI-compatible streaming)
Need Scheduled Jobs?        → go-api + cron-scheduler / go-worker + cron-scheduler
Need Product Analytics?     → go-api + telemetry-posthog
Need API Rate Limiting?     → go-api + rate-limiter
Need Concurrency/Resilience?→ Add from patterns/ by tags
```

**Stack boundary (honest):** apps are Go + Vite only. No Python/Node API, no native WeChat Pay / Alipay / Stripe, no Vercel SSR / Next.js templates. Unimplemented list → [`docs/roadmap.md`](../docs/roadmap.md).

Field-level contract → [`index.md`](index.md)

---

## Registration rules

When adding: update this file + `index.md` / `images/index.md` / `patterns/index.md` / `layouts/index.md` (if layout) + `skills/template-matching.md`
