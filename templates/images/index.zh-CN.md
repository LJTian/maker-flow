# 镜像模版（Dockerfile 片段）

**English** · [简体中文](index.zh-CN.md)

供 Agent 在**拼装 app Dockerfile** 时使用的**片段目录**。  
把片段中的行内联进 builder / runtime 阶段——**不要**预构建本地 `maker-flow/*` tag，也**不要** `FROM` 私有基座镜像。

## 用法（拼装）

1. 读取 `templates/images/<id>/` 下选中的片段 Dockerfile。
2. 将其 `FROM` / `RUN apk` / `WORKDIR`（以及可选的 `USER`）复制进 app Dockerfile 对应阶段。
3. 在片段边界之后追加 app 专属的 `COPY` / `go build` / `ENTRYPOINT` / `ENV` / `EXPOSE`。
4. 产品仓只交付拼装后的 Dockerfile——**不要**把整个 `templates/images/` 树拷进去。

App 模版（`go-api`、`go-cli`、`go-worker`）已默认附带拼装好的 Dockerfile。

## 目录

### go-builder

| 字段 | 值 |
|------|-----|
| **id** | `go-builder` |
| **path** | `templates/images/go-builder` |
| **upstream** | `golang:1.22-alpine` |
| **stage** | builder |
| **tags** | `go`, `build`, `compile` |
| **when_to_use** | Go 多阶段构建（编译阶段） |
| **provides** | Go 1.22 工具链、git、ca-certificates |

### go-runtime

| 字段 | 值 |
|------|-----|
| **id** | `go-runtime` |
| **path** | `templates/images/go-runtime` |
| **upstream** | `alpine:3.20` |
| **stage** | runtime |
| **tags** | `go`, `runtime`, `alpine` |
| **when_to_use** | Go API / 二进制运行阶段 |
| **provides** | alpine 3.20、ca-certificates、tzdata、wget（健康检查）、非 root `nobody`、`WORKDIR /app` |

## 匹配规则

若 app 模版 = `go-api` 或 `go-worker` → 拼装 **两者** `go-builder` + `go-runtime`。  
若 app 模版 = `go-cli` → 容器化 CLI 使用相同片段。  
无需预构建步骤——直接 `docker compose up --build` 即可。

## 扩展

新增 `templates/images/<id>/Dockerfile`（带清晰的 `# --- end base fragment` 标记），在本文件登记。同步更新应使用该片段的 app Dockerfile。
