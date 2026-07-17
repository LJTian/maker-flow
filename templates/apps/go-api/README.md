# go-api

**English** · [简体中文](README.zh-CN.md)

Template id: `go-api`. Gin REST API scaffold for step-4 assembly.

## Capabilities

- Gin router + binding-friendly handlers
- CORS, structured logs, panic recover
- `GET /health`, `GET /api/v1/ping`
- Docker + compose (host port via `HOST_PORT`, default 8080)
- Dockerfile composed from `go-builder` + `go-runtime` fragments (`golang:1.22-alpine` / `alpine:3.20`)

## Agent usage

1. Copy this directory to the **product repo root** (or `<product-root>/<app-id>/` for multi-app). Do not edit the template in place.
2. Add handlers under `internal/handler/` (`func(c *gin.Context)`); register routes in `internal/server/server.go`.
3. Follow `skills/mvp-assembly.md`. Optional patterns: see `templates/patterns/index.md`.
4. If customizing the Dockerfile, compose from `../../images/` fragments (see `images/index.md`).

## Layout

```
cmd/server/
internal/{config,handler,middleware,server}/
Dockerfile
docker-compose.yml
go.mod
```

## Container-first

Do **not** require host `go build` / `go mod tidy`.  
Resolve modules and compile via Docker:

```bash
docker compose up --build
```

`go.sum` is generated inside the build stage (`go mod tidy`); commit optional later if you want lockfile reproducibility.
