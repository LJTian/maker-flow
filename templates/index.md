# 模版目录（供 AI 检索）

步骤 ④ 执行 `skills/template-matching.md` 时：先读 [`CATALOG.md`](CATALOG.md)，再读本文件。

## go-api

| 字段 | 值 |
|------|-----|
| **id** | `go-api` |
| **path** | `templates/go-api` |
| **tags** | `go`, `rest`, `api`, `docker`, `gin`, `high-concurrency` |
| **default_port** | `8080` |
| **when_to_use** | REST API MVP；需要 CORS、结构化日志、健康检查；Go + Gin |
| **includes** | Dockerfile (FROM 基座), docker-compose.yml, CORS, 全局异常, GET /health |
| **images** | `go-builder`, `go-runtime` — 见 [images/index.md](images/index.md) |
| **optional** | docker-compose 内注释掉的 PostgreSQL，PRO 需要 DB 时启用 |
| **docs** | [go-api/README.md](go-api/README.md) |

## 镜像基座（继承式）

应用模版不内置 OS 层；组装前构建：

```bash
./scripts/build-images.sh
```

目录与标签：[images/index.md](images/index.md)

## 选型决策树

```
需要 REST API？
  ├─ 是 → go-api
  └─ 否 → 在 templates/ 下新增模版后再检索
```

## 新增模版

复制 `go-api` 结构，在本文件追加条目，填写 `id`、`tags`、`when_to_use`。
