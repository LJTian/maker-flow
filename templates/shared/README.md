# 模版集共享约定

## Docker 基础镜像

| 语言 | 构建镜像 | 运行镜像 |
|------|----------|----------|
| Go | `golang:1.22-alpine` | `alpine:3.20` |

## 环境变量

- `APP_NAME` / `APP_ENV` / `HTTP_ADDR` / `LOG_LEVEL`
- `HOST_PORT` — 与 `release/` 端口池对齐

## 健康检查

`GET /health` → `{"status":"ok"}`

## 端口池

服务器 `8080–8090`，见 `release/nginx/`。
