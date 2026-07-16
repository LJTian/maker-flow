[English](README.md) · **简体中文**

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

## 目录文件

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

## Agent 写入规则

1. 打开 `CATALOG.md` → 选 **1～N 个 app** → 可选 **patterns**。
2. 基座缺失时运行 `./scripts/build-images.sh`。
3. 将每个 app 复制到 `workspace/<name>/` 或 `workspace/<name>/<app-id>/`；把 pattern 文件合并进需要它的 app。
4. 优先容器构建（`docker compose up --build`）。

不要把已组装的 MVP 留在 `templates/` 内。不要单独部署 patterns。
