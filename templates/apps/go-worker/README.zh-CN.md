[English](README.md) · **简体中文**

# go-worker

模版 id：`go-worker`。带 graceful shutdown 的并发 worker-pool 服务。

## 能力

- 固定大小 worker pool + job channel
- Context 取消 + WaitGroup 排空
- HTTP `/health` 就绪探针（可选）
- Dockerfile：`go-builder` + `go-runtime`

## Agent 用法

1. `./scripts/build-images.sh`
2. 复制到 `workspace/<name>/`
3. 用 PRO 业务逻辑替换 `Job` / `Process`
4. 相关 patterns：`worker-pool`、`pipeline`、`retry-backoff`

## 运行

```bash
./scripts/build-images.sh
docker compose up --build
curl http://localhost:8080/health
```
