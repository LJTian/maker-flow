# 模版集检索目录

> **给 AI：** 步骤 ④ 选型前 **MUST** 先读本文件，再读明细。  
> **给人类：** 一眼看清 apps / images / patterns。

---

## 速览

| 类别 | 数量 | 检索明细 |
|------|:----:|----------|
| 应用模版 (apps) | 3 | [index.md](index.md) · [`apps/`](apps/) |
| 镜像基座 (images) | 2 | [images/index.md](images/index.md) |
| 模式库 (patterns) | 5 | [patterns/index.md](patterns/index.md) |

---

## 应用模版 (apps)

| id | 路径 | 标签 | 何时用 | 依赖镜像 |
|----|------|------|--------|----------|
| `go-api` | [`apps/go-api/`](apps/go-api/) | `go` `gin` `rest` `api` `docker` | Go + Gin REST API MVP | `go-builder` + `go-runtime` |
| `go-cli` | [`apps/go-cli/`](apps/go-cli/) | `go` `cli` `cobra` | 命令行工具 / 子命令骨架 | `go-builder`（+ runtime 可选） |
| `go-worker` | [`apps/go-worker/`](apps/go-worker/) | `go` `worker` `concurrency` `pool` | 多协程任务消费 + graceful shutdown | `go-builder` + `go-runtime` |

Agent：**1～N 个 app** 整目录复制到 `workspace/`（多 app 时用子目录区分）。

---

## 镜像基座 (images)

| id | 本地标签 | 路径 |
|----|----------|------|
| `go-builder` | `maker-flow/go-builder:1.22` | [`images/go-builder/`](images/go-builder/) |
| `go-runtime` | `maker-flow/go-runtime:1.22` | [`images/go-runtime/`](images/go-runtime/) |

```bash
./scripts/build-images.sh
```

---

## 模式库 (patterns)

| id | 路径 | tags |
|----|------|------|
| `worker-pool` | [`patterns/worker-pool/`](patterns/worker-pool/) | `concurrency` `pool` |
| `pipeline` | [`patterns/pipeline/`](patterns/pipeline/) | `fan-in` `fan-out` |
| `singleflight-cache` | [`patterns/singleflight-cache/`](patterns/singleflight-cache/) | `cache` `singleflight` |
| `retry-backoff` | [`patterns/retry-backoff/`](patterns/retry-backoff/) | `retry` `backoff` |
| `circuit-breaker` | [`patterns/circuit-breaker/`](patterns/circuit-breaker/) | `circuit-breaker` |

Agent：先选 **1～N 个 app**，再选 **0～N 个 pattern**，**复制/改写**进对应 app 的 workspace 目录，patterns 不单独部署。

明细 → [`patterns/index.md`](patterns/index.md)

---

## 选型口令（Agent）

```
需要 REST API？     → go-api
需要 CLI？          → go-cli
需要后台 worker？   → go-worker
需要并发/韧性片段？ → 从 patterns/ 按 tags 追加
```

字段级契约 → [`index.md`](index.md)

---

## 登记规则

新增时同步更新：本文件 + `index.md` / `images/index.md` / `patterns/index.md` + `skills/template-matching.md`
