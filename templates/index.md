# Template catalog (for AI search)

**English** · [简体中文](index.zh-CN.md)

Step 4: read [`CATALOG.md`](CATALOG.md) first, then this file and [`patterns/index.md`](patterns/index.md).

## Apps

### go-api

| Field | Value |
|-------|-------|
| **id** | `go-api` |
| **path** | `templates/apps/go-api` |
| **tags** | `go`, `rest`, `api`, `docker`, `gin`, `high-concurrency` |
| **default_port** | `8080` |
| **when_to_use** | Go + Gin REST API MVP |
| **includes** | Dockerfile (FROM bases), docker-compose, CORS, recover, `/health` |
| **images** | `go-builder`, `go-runtime` |
| **suggested_patterns** | `retry-backoff`, `circuit-breaker`, `singleflight-cache` |
| **docs** | [apps/go-api/README.md](apps/go-api/README.md) |

### go-cli

| Field | Value |
|-------|-------|
| **id** | `go-cli` |
| **path** | `templates/apps/go-cli` |
| **tags** | `go`, `cli`, `cobra`, `command` |
| **when_to_use** | CLI tools, ops scripts, local batch entrypoints |
| **includes** | Cobra root + version/run, signal cancel, Dockerfile |
| **images** | `go-builder`, `go-runtime` |
| **suggested_patterns** | `retry-backoff`, `worker-pool` |
| **docs** | [apps/go-cli/README.md](apps/go-cli/README.md) |

### go-worker

| Field | Value |
|-------|-------|
| **id** | `go-worker` |
| **path** | `templates/apps/go-worker` |
| **tags** | `go`, `worker`, `concurrency`, `pool`, `docker` |
| **default_port** | `8080` (health only) |
| **when_to_use** | Background job consumption, multi-goroutine processing, graceful shutdown |
| **includes** | worker pool, job channel, `/health`, compose |
| **images** | `go-builder`, `go-runtime` |
| **suggested_patterns** | `worker-pool`, `pipeline`, `retry-backoff` |
| **docs** | [apps/go-worker/README.md](apps/go-worker/README.md) |

### web-vite

| Field | Value |
|-------|-------|
| **id** | `web-vite` |
| **path** | `templates/apps/web-vite` |
| **tags** | `web`, `frontend`, `vite`, `react`, `typescript`, `tailwind`, `spa`, `docker` |
| **default_port** | `3000` |
| **when_to_use** | Browser UI, landing page, simple dashboard; pairs with `go-api` |
| **includes** | Vite + React + TS + Tailwind, Nginx static serve, `/health`, compose |
| **images** | `node:22-alpine` (build), `nginx:1.27-alpine` (runtime) |
| **suggested_patterns** | — (optional: copy fetch/retry helpers into `src/lib/`) |
| **docs** | [apps/web-vite/README.md](apps/web-vite/README.md) |

## Image bases

See [images/index.md](images/index.md). Before assembly: `./scripts/build-images.sh`

## Pattern library

See [patterns/index.md](patterns/index.md). Selection: 1–N apps first, then 0–N patterns.

## Decision tree

```
Shape? (multi-select OK)
  ├─ REST API     → + go-api
  ├─ CLI          → + go-cli
  ├─ Worker/queue → + go-worker
  ├─ Web UI / SPA → + web-vite
  └─ Other        → extend apps/ and register

Need concurrency/resilience snippets? → append from patterns/ by tags into the matching app (do not deploy alone)
Multi-app layout → workspace/<project>/<app-id>/
```

## Adding templates

- App: place under `apps/<id>/`, register in this file + CATALOG
- Pattern: place under `patterns/<id>/`, register in `patterns/index.md` + CATALOG
