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
| Pattern library (patterns) | 5 | [patterns/index.md](patterns/index.md) |

---

## App templates (apps)

| id | Path | Tags | When to use | Image deps |
|----|------|------|-------------|------------|
| `go-api` | [`apps/go-api/`](apps/go-api/) | `go` `gin` `rest` `api` `docker` | Go + Gin REST API MVP | `go-builder` + `go-runtime` |
| `go-cli` | [`apps/go-cli/`](apps/go-cli/) | `go` `cli` `cobra` | CLI tools / subcommand scaffold | `go-builder` (+ runtime optional) |
| `go-worker` | [`apps/go-worker/`](apps/go-worker/) | `go` `worker` `concurrency` `pool` | Multi-goroutine job consumption + graceful shutdown | `go-builder` + `go-runtime` |
| `web-vite` | [`apps/web-vite/`](apps/web-vite/) | `web` `frontend` `vite` `react` `typescript` `tailwind` `spa` `docker` | Browser UI / landing / dashboard MVP | Node + Nginx (no maker-flow image bases) |

Agent: copy **1–N apps** as whole directories into `workspace/` (use subdirs when multi-app).

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

Agent: pick **1–N apps** first, then **0–N patterns**; **copy/adapt** into the matching app under workspace. Patterns are never deployed alone.

Detail → [`patterns/index.md`](patterns/index.md)

---

## Selection cues (Agent)

```
Need REST API?              → go-api
Need CLI?                   → go-cli
Need background worker?     → go-worker
Need browser UI / SPA?      → web-vite
Need concurrency/resilience → append from patterns/ by tags
```

Field-level contract → [`index.md`](index.md)

---

## Registration rules

When adding: update this file + `index.md` / `images/index.md` / `patterns/index.md` + `skills/template-matching.md`
