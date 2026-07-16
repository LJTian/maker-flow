[English](README.md) · **简体中文**

# Go 编译基座 — maker-flow/go-builder:1.22

仅基建。应用 Dockerfile **MUST** `FROM maker-flow/go-builder:1.22 AS builder`，只追加项目 COPY/构建步骤。

```bash
docker build -t maker-flow/go-builder:1.22 templates/images/go-builder
# or: ./scripts/build-images.sh
```
