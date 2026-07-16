# Agent bootstrap

Load order when an agent session starts on this repo.  
Human quick start: [getting-started.md](getting-started.md) · Repo intro: [README.md](../README.md)

## 1. Orient

1. `AGENTS.md` — audience, hard rules, layout
2. `docs/workflow.md` — state machine + gates
3. Determine **current step** from conversation or artifacts:
   - no PRO → step 2
   - PRO draft, not confirmed → wait at step 3
   - confirmed PRO, no `workspace/<name>/` → step 4
   - MVP exists, not approved → wait at step 5
   - approved MVP → step 6

## 2. Execute current step

| Step | Read first | Then |
|------|------------|------|
| 2 | `skills/pro-generation.md` | `prompts/02-pro-draft.md` |
| ④ | `template-matching.md` → `templates/CATALOG.md` → `mvp-assembly.md` | write under `workspace/` |
| 6 | `skills/deploy.md` | `release/` |

Optional LLM transport: `ai-engine/.env` + `scripts/ai-run.sh` (not required if the host agent is the LLM).

## 3. Preflight checks (before coding)

- [ ] Confirmed PRO exists for step 4+
- [ ] Template chosen via `templates/index.md`
- [ ] Target path is `workspace/<kebab-name>/`
- [ ] Scope matches PRO “不做” list

## 4. Smoke template (optional)

Only to verify host Docker works — not a substitute for step 4:

```bash
cp -r templates/go-api workspace/_smoke
cd workspace/_smoke && cp .env.example .env && docker compose up --build
```
