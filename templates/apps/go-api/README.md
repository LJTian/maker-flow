# go-api

Template id: `go-api`. Gin REST API scaffold for step-4 assembly.

## Capabilities

- Gin router + binding-friendly handlers
- CORS, structured logs, panic recover
- `GET /health`, `GET /api/v1/ping`
- Docker + compose (host port via `HOST_PORT`, default 8080)
- Dockerfile inherits `maker-flow/go-builder:1.22` + `maker-flow/go-runtime:1.22`

## Agent usage

1. Ensure bases exist: `./scripts/build-images.sh` (see `../../images/index.md`).
2. Copy this directory to `workspace/<name>/` (do not edit the template in place).
3. Add handlers under `internal/handler/` (`func(c *gin.Context)`); register routes in `internal/server/server.go`.
4. Follow `skills/mvp-assembly.md`. Optional patterns: see `templates/patterns/index.md`.

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
./scripts/build-images.sh
docker compose up --build
```

`go.sum` is generated inside the build stage (`go mod tidy`); commit optional later if you want lockfile reproducibility.