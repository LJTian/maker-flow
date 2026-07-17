# Go 编译片段 — go-builder

**English** · [简体中文](README.zh-CN.md)

builder 阶段的 Dockerfile **片段**。拼装 app Dockerfile 时，内联本目录的 `FROM` / `RUN apk` / `WORKDIR`，再追加项目 COPY/构建步骤。**不要**预构建本地镜像 tag。

上游：`golang:1.22-alpine`
