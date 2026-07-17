[English](README.md) · **简体中文**

# go-worker

模版 id：`go-worker`。带 graceful shutdown 的并发 worker-pool 服务。

## 能力

- 固定大小 worker pool + job channel
- Context 取消 + WaitGroup 排空
- HTTP `/health` 就绪探针（可选）
- Dockerfile 由 `go-builder` + `go-runtime` 片段拼装

## Agent 用法

1. 复制到 `workspace/<name>/`
2. 用 PRO 业务逻辑替换 `Job` / `Process`
3. 相关 patterns：`worker-pool`、`pipeline`、`retry-backoff`
4. 若自定义 Dockerfile，从 `../../images/` 片段拼装

## 运行

```bash
docker compose up --build
curl http://localhost:8080/health
```
