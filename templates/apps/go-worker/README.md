# go-worker

**English** · [简体中文](README.zh-CN.md)

Template id: `go-worker`. Concurrent worker-pool service with graceful shutdown.

## Capabilities

- Fixed-size worker pool + job channel
- Context cancel + WaitGroup drain
- HTTP `/health` for readiness (optional probe)
- Dockerfile composed from `go-builder` + `go-runtime` fragments

## Agent usage

1. Copy to `workspace/<name>/`
2. Replace `Job` / `Process` with PRO business logic
3. Related patterns: `worker-pool`, `pipeline`, `retry-backoff`
4. If customizing the Dockerfile, compose from `../../images/` fragments

## Run

```bash
docker compose up --build
curl http://localhost:8080/health
```
