[English](README.md) · **简体中文**

# scripts/

| 脚本 | Agent 用途 |
|------|------------|
| `ai-run.sh` | 使用 `ai-engine/.env`，将 prompt 文件 POST 到 OpenAI 兼容 API |
| `build-images.sh` | 在应用 `docker compose up --build` 前构建 `maker-flow/go-*` 基座镜像 |

```bash
./scripts/build-images.sh
./scripts/ai-run.sh prompts/02-pro-draft.md
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

宿主 Agent 本身即 LLM 时不需要 `ai-run.sh`。部署见 `../release/deploy/`。
