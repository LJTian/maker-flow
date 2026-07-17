[English](index.md) · **简体中文**

# 模式库目录

> **给 AI：** 步骤 ④ 在选定 **app** 后，按 PRO 标签选 0～N 个 pattern；**复制/改写**进产品 app，不要整仓当服务部署。  
> **给人类：** 高性能 / 并发片段速查。

每个 pattern 独立 `go.mod`，可在目录内 `go test ./...`（或容器内用 `go-builder` 跑测试）。

## 速览

| id | 路径 | tags | 何时用 |
|----|------|------|--------|
| `worker-pool` | [`worker-pool/`](worker-pool/) | `concurrency` `pool` `channel` | 固定 worker 消化任务队列 |
| `pipeline` | [`pipeline/`](pipeline/) | `concurrency` `fan-in` `fan-out` | 多阶段流水线处理 |
| `singleflight-cache` | [`singleflight-cache/`](singleflight-cache/) | `cache` `singleflight` `ttl` | 防击穿 + 本地 TTL 缓存 |
| `retry-backoff` | [`retry-backoff/`](retry-backoff/) | `retry` `backoff` `resilience` | 可取消的指数退避重试 |
| `circuit-breaker` | [`circuit-breaker/`](circuit-breaker/) | `circuit-breaker` `resilience` | 简易熔断器 |

## Agent 用法

1. 从本表按 PRO 关键词 / tags 匹配。
2. 将 pattern 包文件拷入**需要它的 app** 目录（`<产品根>/<app-id>/internal/...`）。
3. **MUST NOT** 为 pattern 单独开公网部署。

## 登记规则

新增 pattern → 本文件追加一行 + 独立目录（含 `README.md`、可测代码）。
