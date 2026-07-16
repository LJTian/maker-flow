[English](README.md) · **简体中文**

# 模版共享约定

供 Agent 从任意模版组装到 `workspace/` 时遵循。

## 镜像（继承）

不要在应用 Dockerfile 中写死上游 OS。使用 `templates/images/index.md` 中的标签：

| 角色 | 本地标签 | 源目录 |
|------|----------|--------|
| Go 构建 | `maker-flow/go-builder:1.22` | `templates/images/go-builder` |
| Go 运行 | `maker-flow/go-runtime:1.22` | `templates/images/go-runtime` |

构建：`./scripts/build-images.sh`

## 环境变量名

`APP_NAME`、`APP_ENV`、`HTTP_ADDR`、`LOG_LEVEL`、`HOST_PORT`（compose 宿主机映射 → release 端口池）

## 健康检查

`GET /health` → `{"status":"ok"}`

## 端口

服务端口池 `8080–8090`；见 `release/`。
