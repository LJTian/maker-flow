# Go API 模版

高并发 REST API 最小骨架，供步骤 ④ AI 组装时复制使用。

## 预设能力

- CORS、结构化日志、全局异常恢复
- `GET /health`、`GET /api/v1/ping`
- Dockerfile + docker-compose.yml

## 本地运行

```bash
cp .env.example .env
docker compose up --build
curl http://localhost:8080/health
```

## 目录

```
go-api/
├── cmd/server/
├── internal/{config,handler,middleware,server}/
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## AI 组装时

按 PRO 在 `internal/handler/` 增加业务 handler，在 `internal/server/server.go` 注册路由。  
输出目录约定见 `skills/mvp-assembly.md`。
