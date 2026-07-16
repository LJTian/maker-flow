# 模版目录（供 AI 检索）

步骤 ④：先读 [`CATALOG.md`](CATALOG.md)，再读本文件与 [`patterns/index.md`](patterns/index.md)。

## Apps

### go-api

| 字段 | 值 |
|------|-----|
| **id** | `go-api` |
| **path** | `templates/apps/go-api` |
| **tags** | `go`, `rest`, `api`, `docker`, `gin`, `high-concurrency` |
| **default_port** | `8080` |
| **when_to_use** | Go + Gin REST API MVP |
| **includes** | Dockerfile (FROM 基座), docker-compose, CORS, recover, `/health` |
| **images** | `go-builder`, `go-runtime` |
| **suggested_patterns** | `retry-backoff`, `circuit-breaker`, `singleflight-cache` |
| **docs** | [apps/go-api/README.md](apps/go-api/README.md) |

### go-cli

| 字段 | 值 |
|------|-----|
| **id** | `go-cli` |
| **path** | `templates/apps/go-cli` |
| **tags** | `go`, `cli`, `cobra`, `command` |
| **when_to_use** | 命令行工具、运维脚本、本地批处理入口 |
| **includes** | Cobra root + version/run、signal cancel、Dockerfile |
| **images** | `go-builder`, `go-runtime` |
| **suggested_patterns** | `retry-backoff`, `worker-pool` |
| **docs** | [apps/go-cli/README.md](apps/go-cli/README.md) |

### go-worker

| 字段 | 值 |
|------|-----|
| **id** | `go-worker` |
| **path** | `templates/apps/go-worker` |
| **tags** | `go`, `worker`, `concurrency`, `pool`, `docker` |
| **default_port** | `8080`（health only） |
| **when_to_use** | 后台任务消费、多协程处理、需 graceful shutdown |
| **includes** | worker pool、job channel、`/health`、compose |
| **images** | `go-builder`, `go-runtime` |
| **suggested_patterns** | `worker-pool`, `pipeline`, `retry-backoff` |
| **docs** | [apps/go-worker/README.md](apps/go-worker/README.md) |

## 镜像基座

见 [images/index.md](images/index.md)。组装前：`./scripts/build-images.sh`

## 模式库

见 [patterns/index.md](patterns/index.md)。选型：先 1～N 个 app，再 0～N 个 pattern。

## 选型决策树

```
形态？（可多选组合）
  ├─ REST API     → + go-api
  ├─ CLI          → + go-cli
  ├─ Worker/队列  → + go-worker
  └─ 其它         → 扩展 apps/ 并登记

需要并发/韧性片段？ → patterns/ 按 tags 追加到对应 app（不单独部署）
多 app 布局        → workspace/<project>/<app-id>/
```

## 新增模版

- App：放入 `apps/<id>/`，登记本文件 + CATALOG  
- Pattern：放入 `patterns/<id>/`，登记 `patterns/index.md` + CATALOG
