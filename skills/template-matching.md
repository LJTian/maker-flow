# Template matching skill

**English** · [简体中文](template-matching.zh-CN.md)

**Step:** 4 — AI matches templates from the PRO  
**Depends on:** `templates/CATALOG.md` → `templates/index.md` → `templates/patterns/index.md`

## Goal

1. Select **1–N apps** (`templates/apps/`) — compose shapes the PRO needs, e.g. API + worker  
2. Select **0–N patterns** (`templates/patterns/`, by tags)  
3. List dependent **image** tags (union across selected apps)

## Input

- Confirmed PRO
- `templates/CATALOG.md`

## Output

```markdown
## Selected templates
- **Apps**:
  - go-api → templates/apps/go-api
  - go-worker → templates/apps/go-worker
- **Image deps**: go-builder + go-runtime
- **Patterns**: retry-backoff, worker-pool (may be empty)
- **Product layout**: `<product-root>/{api,worker}/` (or describe otherwise)
- **Rationale**: …
```

## Matching rules

| PRO signal | App | Common patterns |
|------------|-----|-----------------|
| REST API, Gin | `go-api` | `retry-backoff`, `circuit-breaker`, `singleflight-cache` |
| CLI / command-line tool | `go-cli` | `retry-backoff`, `worker-pool` |
| Background jobs / multi-goroutine consumers | `go-worker` | `worker-pool`, `pipeline`, `retry-backoff` |
| Browser UI / SPA / dashboard | `web-vite` | — (optional snippets in `src/lib/`) |

Multi-app examples: `go-api` + `go-worker` (sync API + async consume); `go-api` + `go-cli` (service + ops commands); `go-api` + `web-vite` (API + browser UI).

## MUST NOT

- MUST NOT skip the catalog and invent scaffolding
- MUST NOT run before the PRO is confirmed
- MUST NOT deploy a pattern as a standalone public service
- MUST NOT force-fit unrelated app shapes (each app MUST map to a PRO responsibility)
