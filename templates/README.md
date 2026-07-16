# templates/

**检索入口（人 + AI）：** [`CATALOG.md`](CATALOG.md)

```
templates/
├── CATALOG.md      # 总览
├── index.md        # apps 字段明细
├── apps/           # 可组装完整工程 → workspace/
├── images/         # 继承式 Docker 基座
├── patterns/       # 可编译片段 → 拷进 workspace
└── shared/
```

## Catalog files

| 文件 | 用途 |
|------|------|
| [`CATALOG.md`](CATALOG.md) | 总览 / 速查 |
| [`index.md`](index.md) | apps 字段级明细 |
| [`images/index.md`](images/index.md) | 镜像基座 |
| [`patterns/index.md`](patterns/index.md) | 模式库 |

## Apps

| id | path | port | images |
|----|------|------|--------|
| `go-api` | `apps/go-api/` | 8080 | builder + runtime |
| `go-cli` | `apps/go-cli/` | — | builder + runtime |
| `go-worker` | `apps/go-worker/` | 8080 | builder + runtime |

## Agent write rule

1. Open `CATALOG.md` → pick **1–N apps** → optional **patterns**.
2. `./scripts/build-images.sh` if bases missing.
3. Copy each app → `workspace/<name>/` or `workspace/<name>/<app-id>/`; merge pattern files into the app that needs them.
4. Prefer container build (`docker compose up --build`).

Do not leave assembled MVPs inside `templates/`. Do not deploy patterns alone.
