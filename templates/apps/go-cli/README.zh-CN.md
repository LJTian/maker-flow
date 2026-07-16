[English](README.md) · **简体中文**

# go-cli

模版 id：`go-cli`。Cobra CLI 脚手架，用于步骤 ④ 组装（命令 / 工具）。

## 能力

- Cobra root + `version` / `run` 子命令
- 结构化 slog 日志
- Context + signal 优雅取消
- Dockerfile 继承 `maker-flow/go-builder:1.22`（静态二进制）

## Agent 用法

1. 若用 Docker 构建，先运行 `./scripts/build-images.sh`。
2. 复制到 `workspace/<name>/`。
3. 在 `cmd/` / `internal/` 下添加子命令。
4. 可选 patterns：`retry-backoff`、`worker-pool`。

## 运行

```bash
# host (optional)
go run ./cmd/cli --help

# container
./scripts/build-images.sh
docker build -t maker-flow/go-cli:local .
docker run --rm maker-flow/go-cli:local version
```
