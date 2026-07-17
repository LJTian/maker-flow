# ai-engine/

**English** · [简体中文](README.zh-CN.md)

Optional **LLM connection notes** for agents that call models via HTTP (OpenAI-compatible).  
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
```

Use your own HTTP client or IDE agent against that endpoint.  
Workflow authority remains `docs/workflow.md` + `skills/`, not this folder.
