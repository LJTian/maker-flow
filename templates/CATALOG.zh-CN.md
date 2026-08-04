[English](CATALOG.md) · **简体中文**

# 模版集检索目录

> **给 AI：** 步骤 ④ 选型前 **MUST** 先读本文件，再读明细。  
> **给人类：** 一眼看清 apps / images / patterns。

---

## 速览

| 类别 | 数量 | 检索明细 |
|------|:----:|----------|
| 应用模版 (apps) | 4 | [index.md](index.md) · [`apps/`](apps/) |
| 镜像基座 (images) | 2 | [images/index.md](images/index.md) |
| 模式库 (patterns) | 6 | [patterns/index.md](patterns/index.md) |
| 布局 (多 app 根) | 1 | [layouts/index.md](layouts/index.md) |

---

## 应用模版 (apps)

| id | 路径 | 标签 | 何时用 | 依赖镜像 |
|----|------|------|--------|----------|
| `go-api` | [`apps/go-api/`](apps/go-api/) | `go` `gin` `rest` `api` `docker` | Go + Gin REST API MVP | `go-builder` + `go-runtime` |
| `go-cli` | [`apps/go-cli/`](apps/go-cli/) | `go` `cli` `cobra` | 命令行工具 / 子命令骨架 | `go-builder`（+ runtime 可选） |
| `go-worker` | [`apps/go-worker/`](apps/go-worker/) | `go` `worker` `concurrency` `pool` | 多协程任务消费 + graceful shutdown | `go-builder` + `go-runtime` |
| `web-vite` | [`apps/web-vite/`](apps/web-vite/) | `web` `frontend` `vite` `react` `typescript` `tailwind` `spa` `docker` | 浏览器 UI / 落地页 / 简易面板 MVP | Node + Nginx（无 maker-flow 镜像基座） |

Agent：**1～N 个 app** 整目录复制到**产品仓**（多 app 时用子目录区分）。

---

## 镜像片段 (images)

| id | 上游 | 路径 |
|----|------|------|
| `go-builder` | `golang:1.22-alpine` | [`images/go-builder/`](images/go-builder/) |
| `go-runtime` | `alpine:3.20` | [`images/go-runtime/`](images/go-runtime/) |

拼装时内联进 app Dockerfile — 见 [`images/index.md`](images/index.md)。无需预构建。
---

## 模式库 (patterns)

| id | 路径 | tags |
|----|------|------|
| `worker-pool` | [`patterns/worker-pool/`](patterns/worker-pool/) | `concurrency` `pool` |
| `pipeline` | [`patterns/pipeline/`](patterns/pipeline/) | `fan-in` `fan-out` |
| `singleflight-cache` | [`patterns/singleflight-cache/`](patterns/singleflight-cache/) | `cache` `singleflight` |
| `retry-backoff` | [`patterns/retry-backoff/`](patterns/retry-backoff/) | `retry` `backoff` |
| `circuit-breaker` | [`patterns/circuit-breaker/`](patterns/circuit-breaker/) | `circuit-breaker` |
| `persistence-sqlx` | [`patterns/persistence-sqlx/`](patterns/persistence-sqlx/) | `db` `sqlx` `sqlite` `postgres` `mysql` |

Agent：先选 **1～N 个 app**，再选 **0～N 个 pattern**，**复制/改写**进对应 app 的产品仓目录，patterns 不单独部署。

明细 → [`patterns/index.md`](patterns/index.md)

---

## 布局（多 app 产品仓根）

| id | 路径 | Apps | 何时用 |
|----|------|------|--------|
| `web-api` | [`layouts/web-api/`](layouts/web-api/) | `go-api` + `web-vite` | 浏览器 UI + REST API 根 compose / env |

先拷 `api/` + `web/`，再拷布局根文件。明细 → [`layouts/index.md`](layouts/index.md)。拆分上线 → [`skills/publish-split-web-api.md`](../skills/publish-split-web-api.md)。

---

## 选型口令（Agent）

```
需要 REST API？     → go-api
需要 CLI？          → go-cli
需要后台 worker？   → go-worker
需要浏览器 UI？    → web-vite
需要 API + SPA？   → go-api + web-vite + layout web-api
需要数据库 / 表？  → go-api + persistence-sqlx
需要并发/韧性片段？ → 从 patterns/ 按 tags 追加
```

字段级契约 → [`index.md`](index.md)

---

## 登记规则

新增时同步更新：本文件 + `index.md` / `images/index.md` / `patterns/index.md` / `layouts/index.md`（若是布局）+ `skills/template-matching.md`
