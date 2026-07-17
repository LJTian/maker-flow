# Go compile fragment — go-builder

**English** · [简体中文](README.zh-CN.md)

Dockerfile **fragment** for the builder stage. When assembling an app Dockerfile, inline `FROM` / `RUN apk` / `WORKDIR` from this directory, then append project COPY/build steps. Do **not** pre-build a local image tag.

Upstream: `golang:1.22-alpine`
