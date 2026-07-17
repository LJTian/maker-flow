# Go runtime fragment — go-runtime

**English** · [简体中文](README.zh-CN.md)

Dockerfile **fragment** for the runtime stage. When assembling an app Dockerfile, inline `FROM` / `RUN apk` / `WORKDIR` / `USER` from this directory, then COPY the binary and set ENTRYPOINT/ENV/EXPOSE. Do **not** pre-build a local image tag.

Upstream: `alpine:3.20`. Includes `wget` for compose healthchecks. Default user is `nobody`.
