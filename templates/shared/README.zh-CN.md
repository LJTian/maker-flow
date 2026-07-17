[English](README.md) · **简体中文**

# 模版共享约定

供 Agent 从任意模版组装到**产品仓**时遵循。

## 镜像（Dockerfile 片段）

不要临时发明上游 OS 行。从 `templates/images/index.md` 的片段拼装：

| 角色 | 上游 | 源目录 |
|------|------|--------|
| Go 构建 | `golang:1.22-alpine` | `templates/images/go-builder` |
| Go 运行 | `alpine:3.20` | `templates/images/go-runtime` |

把片段行内联进 app Dockerfile。**不要**预构建私有 `maker-flow/*` tag。

## 环境变量名

`APP_NAME`、`APP_ENV`、`HTTP_ADDR`、`LOG_LEVEL`、`HOST_PORT`（可选，**本地** compose 宿主机映射，用于 `docker compose up` 验收）

## 健康检查

`GET /health` → `{"status":"ok"}`

## 端口

- **本地验收：** 按需用 `HOST_PORT` 映射（如 `8080:8080`，或 web `3000:80`）。
- **生产：** 公网入口是 Docker Nginx 网关宿主机 **80**；MVP 经共享网络 `maker-flow` 上的别名 `MVP_NAME:CONTAINER_PORT` 可达（见 `release/`）。
