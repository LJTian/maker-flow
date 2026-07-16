[English](README.md) · **简体中文**

# Go 运行基座 — maker-flow/go-runtime:1.22

仅基建。应用 Dockerfile **MUST** `FROM maker-flow/go-runtime:1.22`，只 COPY 二进制并设置 ENTRYPOINT/ENV/EXPOSE。

以 `nobody` 运行。含 `wget` 供 compose healthcheck。

```bash
docker build -t maker-flow/go-runtime:1.22 templates/images/go-runtime
# or: ./scripts/build-images.sh
```
