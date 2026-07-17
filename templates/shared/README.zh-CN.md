[English](README.md) · **简体中文**

# 模版共享约定

供 Agent 从任意模版组装到 `workspace/` 时遵循。

## 镜像（Dockerfile 片段）

不要临时发明上游 OS 行。从 `templates/images/index.md` 的片段拼装：

| 角色 | 上游 | 源目录 |
|------|------|--------|
| Go 构建 | `golang:1.22-alpine` | `templates/images/go-builder` |
| Go 运行 | `alpine:3.20` | `templates/images/go-runtime` |

把片段行内联进 app Dockerfile。**不要**预构建私有 `maker-flow/*` tag。

## 环境变量名

`APP_NAME`、`APP_ENV`、`HTTP_ADDR`、`LOG_LEVEL`、`HOST_PORT`（compose 宿主机映射 → release 端口池）

## 健康检查

`GET /health` → `{"status":"ok"}`

## 端口

服务端口池 `8080–8090`；见 `release/`。
