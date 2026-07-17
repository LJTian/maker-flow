# ai-engine/

**English** · [简体中文](README.zh-CN.md)

**Most users can ignore this directory.** Cursor / Claude (or any host agent that is already the LLM) should follow `skills/` and `docs/workflow.md` only — no `.env` setup required.

Optional notes for calling an OpenAI-compatible HTTP API yourself (local Ollama, gateways, etc.).

## Contents

| Path | Purpose |
|------|---------|
| `.env.example` | `AI_BASE_URL`, `AI_MODEL`, generation params |
| `params.md` | Parameter bounds + per-step acceptance |
| `providers/` | Example backends (Ollama, OpenAI, compatible gateways) |

## Setup (only if you call the API yourself)

```bash
cp ai-engine/.env.example ai-engine/.env
# set AI_BASE_URL / AI_MODEL per providers/
```

Use your own HTTP client against that endpoint.  
Workflow authority remains `docs/workflow.md` + `skills/`, not this folder.
