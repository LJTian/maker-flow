# ai-engine/

Optional **LLM transport** for agents that invoke models via HTTP (OpenAI-compatible).  
If the host product (e.g. Cursor Agent) is already the model, this directory is unused — still follow `skills/` and `docs/workflow.md`.

## Contents

| Path | Purpose |
|------|---------|
| `.env.example` | `AI_BASE_URL`, `AI_MODEL`, generation params |
| `params.md` | Parameter bounds + per-step acceptance |
| `providers/` | Example backends (Ollama, OpenAI, compatible gateways) |

## Setup (when used)

```bash
cp ai-engine/.env.example ai-engine/.env
# set AI_BASE_URL / AI_MODEL per providers/
./scripts/ai-run.sh prompts/02-pro-draft.md
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

Workflow authority remains `docs/workflow.md` + `skills/`, not this folder.
