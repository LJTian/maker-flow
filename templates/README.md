# templates/

Searchable scaffold set for step 4. Agents MUST NOT invent a new project layout when a catalog entry matches.

**检索入口（人 + AI）：** [`CATALOG.md`](CATALOG.md) ← 先读这个

## Catalog files

| 文件 | 用途 |
|------|------|
| [`CATALOG.md`](CATALOG.md) | 总览 / 速查（人类 + Agent） |
| [`index.md`](index.md) | 应用模版字段级明细（Agent 选型） |
| [`images/index.md`](images/index.md) | 镜像基座明细 |

| id | path | port | images |
|----|------|------|--------|
| `go-api` | `go-api/` | 8080 | `go-builder`, `go-runtime` |

## What each app template provides

- Thin `Dockerfile` that **inherits** image bases (no OS reinvent)
- `docker-compose.yml` + runnable demo source
- `.env.example` + README

Shared conventions: `shared/`.

## Agent write rule

1. Open [`CATALOG.md`](CATALOG.md); match app via `index.md`; resolve images via `images/index.md`.
2. Run `./scripts/build-images.sh` if base tags missing.
3. Copy selected **app** template → `workspace/<name>/` (not `images/`).
4. Implement PRO only (`skills/mvp-assembly.md`). Prefer container build.

Do not leave assembled MVPs inside `templates/`.
