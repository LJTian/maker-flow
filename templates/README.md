# templates/

**English** · [简体中文](README.zh-CN.md)

**Catalog entry (humans + AI):** [`CATALOG.md`](CATALOG.md)

```
templates/
├── CATALOG.md      # overview
├── index.md        # apps field-level detail
├── apps/           # assemblable full projects → workspace/
├── images/         # inheritance-style Docker bases
├── patterns/       # compilable snippets → copy into workspace
└── shared/
```

## Catalog files

| File | Purpose |
|------|---------|
| [`CATALOG.md`](CATALOG.md) | Overview / quick lookup |
| [`index.md`](index.md) | Apps field-level detail |
| [`images/index.md`](images/index.md) | Image bases |
| [`patterns/index.md`](patterns/index.md) | Pattern library |

## Apps

| id | path | port | images |
|----|------|------|--------|
| `go-api` | `apps/go-api/` | 8080 | builder + runtime |
| `go-cli` | `apps/go-cli/` | — | builder + runtime |
| `go-worker` | `apps/go-worker/` | 8080 | builder + runtime |
| `web-vite` | `apps/web-vite/` | 3000 | node + nginx |

## Agent write rule

1. Open `CATALOG.md` → pick **1–N apps** → optional **patterns**.
2. `./scripts/build-images.sh` if bases missing.
3. Copy each app → `workspace/<name>/` or `workspace/<name>/<app-id>/`; merge pattern files into the app that needs them.
4. Prefer container build (`docker compose up --build`).

Do not leave assembled MVPs inside `templates/`. Do not deploy patterns alone.
