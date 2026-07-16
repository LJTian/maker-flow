[English](index.md) · **简体中文**

# 镜像模版（基座镜像）

Agent 用的**继承式** Docker 基座目录。应用模版（`go-api` 等）只使用 `FROM maker-flow/<id>:<tag>`，**MUST NOT** 再声明基座里已有的系统包。

## 构建全部 Go 基座

```bash
./scripts/build-images.sh
```

## 目录

### go-builder

| 字段 | 值 |
|------|-----|
| **id** | `go-builder` |
| **path** | `templates/images/go-builder` |
| **tag** | `maker-flow/go-builder:1.22` |
| **tags** | `go`, `build`, `compile` |
| **when_to_use** | Go 多阶段构建（编译阶段） |
| **provides** | Go 1.22 工具链、git、ca-certificates |

### go-runtime

| 字段 | 值 |
|------|-----|
| **id** | `go-runtime` |
| **path** | `templates/images/go-runtime` |
| **tag** | `maker-flow/go-runtime:1.22` |
| **tags** | `go`, `runtime`, `alpine` |
| **when_to_use** | Go API / 二进制运行阶段 |
| **provides** | alpine 3.20、ca-certificates、tzdata、wget（healthcheck）、非 root `nobody`、`WORKDIR /app` |

## 匹配规则

若应用模版为 `go-api` 或 `go-worker` → **同时**需要 `go-builder` + `go-runtime`。  
若应用模版为 `go-cli` → 容器化 CLI 使用相同标签。  
在 `workspace/` 执行 `docker compose up --build` / `docker build` **之前**先构建基座。

## 扩展

新增 `templates/images/<id>/Dockerfile`，在本文件登记，并更新 `scripts/build-images.sh`。
