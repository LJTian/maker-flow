# Go 运行片段 — go-runtime

**English** · [简体中文](README.zh-CN.md)

runtime 阶段的 Dockerfile **片段**。拼装 app Dockerfile 时，内联本目录的 `FROM` / `RUN apk` / `WORKDIR` / `USER`，再 COPY 二进制并设置 ENTRYPOINT/ENV/EXPOSE。**不要**预构建本地镜像 tag。

上游：`alpine:3.20`。含 `wget` 供 compose 健康检查。默认用户为 `nobody`。
