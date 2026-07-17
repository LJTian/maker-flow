[English](README.md) · **简体中文**

# ai-engine/

**多数用户可忽略本目录。** Cursor / Claude（或任何本身就是 LLM 的宿主 Agent）只需遵循 `skills/` 与 `docs/workflow.md`，无需配置 `.env`。

若你自行调用 OpenAI 兼容 HTTP API（本机 Ollama、网关等），这里提供可选连接说明。

## 内容

| 路径 | 用途 |
|------|------|
| `.env.example` | `AI_BASE_URL`、`AI_MODEL`、生成参数 |
| `params.md` | 参数边界 + 分步骤验收 |
| `providers/` | 示例后端（Ollama、OpenAI、兼容网关） |

## 配置（仅在你自己调 API 时）

```bash
cp ai-engine/.env.example ai-engine/.env
# set AI_BASE_URL / AI_MODEL per providers/
```

用你自己的 HTTP 客户端对接该端点。  
流程权威仍是 `docs/workflow.md` + `skills/`，不是本目录。
