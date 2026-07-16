[English](README.md) · **简体中文**

# ai-engine/

可选的 **LLM 传输层**，供通过 HTTP（OpenAI 兼容）调用模型的 Agent 使用。  
若宿主产品（如 Cursor Agent）本身就是模型，则本目录不使用 — 仍须遵循 `skills/` 与 `docs/workflow.md`。

## 内容

| 路径 | 用途 |
|------|------|
| `.env.example` | `AI_BASE_URL`、`AI_MODEL`、生成参数 |
| `params.md` | 参数边界 + 分步骤验收 |
| `providers/` | 示例后端（Ollama、OpenAI、兼容网关） |

## 配置（使用时）

```bash
cp ai-engine/.env.example ai-engine/.env
# set AI_BASE_URL / AI_MODEL per providers/
./scripts/ai-run.sh prompts/02-pro-draft.md
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

流程权威仍是 `docs/workflow.md` + `skills/`，不是本目录。
