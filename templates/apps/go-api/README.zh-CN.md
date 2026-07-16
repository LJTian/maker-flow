[English](README.md) · **简体中文**

# go-api

模版 id：`go-api`。Gin REST API 脚手架，用于步骤 ④ 组装。

## 能力

- Gin 路由 + 便于 binding 的 handlers
- CORS、结构化日志、panic recover
- `GET /health`、`GET /api/v1/ping`
- Docker + compose（宿主机端口由 `HOST_PORT` 控制，默认 8080）
- Dockerfile 继承 `maker-flow/go-builder:1.22` + `maker-flow/go-runtime:1.22`

## Agent 用法

1. 确保基座存在：`./scripts/build-images.sh`（见 `../../images/index.md`）。
2. 将本目录复制到 `workspace/<name>/`（不要原地改模版）。
3. 在 `internal/handler/` 添加 handlers（`func(c *gin.Context)`）；在 `internal/server/server.go` 注册路由。
4. 遵循 `skills/mvp-assembly.md`。可选 patterns：见 `templates/patterns/index.md`。

## 布局

```
cmd/server/
internal/{config,handler,middleware,server}/
Dockerfile
docker-compose.yml
go.mod
```

## 容器优先

**不要**要求宿主机执行 `go build` / `go mod tidy`。  
通过 Docker 解析模块并编译：

```bash
./scripts/build-images.sh
docker compose up --build
```

`go.sum` 在构建阶段生成（`go mod tidy`）；若需要锁文件可复现，可稍后提交。
