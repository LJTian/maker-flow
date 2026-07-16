# go-worker

**English** · [简体中文](README.zh-CN.md)

Template id: `go-worker`. Concurrent worker-pool service with graceful shutdown.

## Capabilities

- Fixed-size worker pool + job channel
- Context cancel + WaitGroup drain
- HTTP `/health` for readiness (optional probe)
- Dockerfile: `go-builder` + `go-runtime`

## Agent usage

1. `./scripts/build-images.sh`
2. Copy to `workspace/<name>/`
3. Replace `Job` / `Process` with PRO business logic
4. Related patterns: `worker-pool`, `pipeline`, `retry-backoff`

## Run

```bash
./scripts/build-images.sh
docker compose up --build
curl http://localhost:8080/health
```
