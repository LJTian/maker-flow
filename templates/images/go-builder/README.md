# Go compile base — maker-flow/go-builder:1.22

**English** · [简体中文](README.zh-CN.md)

Infrastructure only. App Dockerfiles MUST `FROM maker-flow/go-builder:1.22 AS builder` and only add project COPY/build steps.

```bash
docker build -t maker-flow/go-builder:1.22 templates/images/go-builder
# or: ./scripts/build-images.sh
```
