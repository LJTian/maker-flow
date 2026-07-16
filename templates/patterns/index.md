# Pattern library catalog

**English** · [简体中文](index.zh-CN.md)

> **For AI:** After picking **app(s)** in step 4, select 0–N patterns by PRO tags; **copy/adapt** into `workspace/`. Do not deploy the pattern tree as a service.  
> **For humans:** Quick lookup for high-performance / concurrency snippets.

Each pattern has its own `go.mod` and can run `go test ./...` in-directory (or inside a container with `go-builder`).

## Overview

| id | Path | tags | When to use |
|----|------|------|-------------|
| `worker-pool` | [`worker-pool/`](worker-pool/) | `concurrency` `pool` `channel` | Fixed workers draining a job queue |
| `pipeline` | [`pipeline/`](pipeline/) | `concurrency` `fan-in` `fan-out` | Multi-stage pipeline processing |
| `singleflight-cache` | [`singleflight-cache/`](singleflight-cache/) | `cache` `singleflight` `ttl` | Collapse stampede + local TTL cache |
| `retry-backoff` | [`retry-backoff/`](retry-backoff/) | `retry` `backoff` `resilience` | Cancellable exponential backoff retry |
| `circuit-breaker` | [`circuit-breaker/`](circuit-breaker/) | `circuit-breaker` `resilience` | Simple circuit breaker |

## Agent usage

1. Match from this table by PRO keywords / tags.
2. Copy pattern package files into the **app that needs them** (`workspace/<project>/<app-id>/internal/...`).
3. **MUST NOT** expose a public deploy for a pattern alone.

## Registration rules

New pattern → add a row here + standalone directory (with `README.md` and testable code).
