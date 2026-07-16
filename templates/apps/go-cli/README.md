# go-cli

Template id: `go-cli`. Cobra CLI scaffold for step-4 assembly (commands / tools).

## Capabilities

- Cobra root + `version` / `run` subcommands
- Structured slog logging
- Context + signal graceful cancel
- Dockerfile inherits `maker-flow/go-builder:1.22` (static binary)

## Agent usage

1. `./scripts/build-images.sh` if building with Docker.
2. Copy to `workspace/<name>/`.
3. Add subcommands under `cmd/` / `internal/`.
4. Optional patterns: `retry-backoff`, `worker-pool`.

## Run

```bash
# host (optional)
go run ./cmd/cli --help

# container
./scripts/build-images.sh
docker build -t maker-flow/go-cli:local .
docker run --rm maker-flow/go-cli:local version
```
