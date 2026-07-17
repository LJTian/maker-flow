# go-cli

**English** · [简体中文](README.zh-CN.md)

Template id: `go-cli`. Cobra CLI scaffold for step-4 assembly (commands / tools).

## Capabilities

- Cobra root + `version` / `run` subcommands
- Structured slog logging
- Context + signal graceful cancel
- Dockerfile composed from `go-builder` + `go-runtime` fragments (static binary)

## Agent usage

1. Copy to `workspace/<name>/`.
2. Add subcommands under `cmd/` / `internal/`.
3. Optional patterns: `retry-backoff`, `worker-pool`.
4. If customizing the Dockerfile, compose from `../../images/` fragments.

## Run

```bash
# host (optional)
go run ./cmd/cli --help

# container
docker build -t maker-flow/go-cli:local .
docker run --rm maker-flow/go-cli:local version
```
