[English](README.md) · **简体中文**

# go-cli

模版 id：`go-cli`。Cobra CLI 脚手架，用于步骤 ④ 组装（命令 / 工具）。

## 能力

- Cobra root + `version` / `run` 子命令
- 结构化 slog 日志
- Context + signal 优雅取消
- Dockerfile 由 `go-builder` + `go-runtime` 片段拼装（静态二进制）

## Agent 用法

1. 复制到**产品仓根**（或 `<产品根>/<app-id>/`）。
2. 在 `cmd/` / `internal/` 下添加子命令。
3. 可选 patterns：`retry-backoff`、`worker-pool`。
4. 若自定义 Dockerfile，从 `../../images/` 片段拼装。

## 运行

```bash
# host (optional)
go run ./cmd/cli --help

# container
docker build -t maker-flow/go-cli:local .
docker run --rm maker-flow/go-cli:local version
```
