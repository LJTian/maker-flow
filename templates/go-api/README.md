# go-api

Template id: `go-api`. REST API scaffold for step-4 assembly.

## Capabilities

- CORS, structured logs, panic recover
- `GET /health`, `GET /api/v1/ping`
- Docker + compose (host port via `HOST_PORT`, default 8080)

## Agent usage

1. Copy this directory to `workspace/<name>/` (do not edit the template in place).
2. Add handlers under `internal/handler/`; register routes in `internal/server/server.go`.
3. Follow `skills/mvp-assembly.md`.

## Layout

```
cmd/server/
internal/{config,handler,middleware,server}/
Dockerfile
docker-compose.yml
go.mod
```
