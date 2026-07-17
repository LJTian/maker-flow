# Parameters and output constraints

**English** · [简体中文](params.zh-CN.md)

Reference parameters for OpenAI-compatible calls (see `ai-engine/.env.example`).  
Host agents (e.g. Cursor) usually ignore this file and follow `skills/` instead.

## Connection

| Variable | Required | Description |
|----------|----------|-------------|
| `AI_BASE_URL` | yes | OpenAI-compatible API root |
| `AI_API_KEY` | no | Bearer token |
| `AI_MODEL` | yes | Model id |

## Generation

| Variable | Default | Step 2 PRO | Step 4 assemble |
|----------|---------|------------|-----------------|
| `AI_TEMPERATURE` | `0.6` | `0.5–0.6` | `0.3–0.5` |
| `AI_MAX_TOKENS` | `8192` | `4096–8192` | `8192+` |

## Per-step acceptance

### Step 2 (PRO)

Output MUST include every section from `skills/pro-generation.md`, with **no code**.

### Step 4 (assemble)

MUST include template selection rationale + runnable code under `workspace/<project-name>/`, per `skills/mvp-assembly.md`.
