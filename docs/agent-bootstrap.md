# Agent bootstrap

**English** · [简体中文](agent-bootstrap.zh-CN.md)

Load order when an agent session starts.  
Human quick start: [getting-started.md](getting-started.md) · Repo intro: [README.md](../README.md) · Consumer: [consumer-project.md](consumer-project.md) · i18n: [i18n.md](i18n.md)

## 1. Orient

1. `AGENTS.md` — audience, hard rules, layout, language rules  
   (In a **product repo**, the local `AGENTS.md` from `AGENTS.consumer.example.md` wins; factory is `$MAKER_FLOW_ROOT`.)
2. `docs/workflow.md` — state machine + gates
3. Determine **current step** from conversation or artifacts:
   - no PRO → step 2
   - PRO draft, not confirmed → wait at step 3
   - confirmed PRO, no assembled project yet → step 4
   - MVP exists, not approved → wait at step 5
   - approved MVP → step 6

## 2. Execute current step

| Step | Read first | Then |
|------|------------|------|
| 2 | `skills/pro-generation.md` | `prompts/02-pro-draft.md` + `prompts/pro.template.md` (sample: `pro.example.md`) |
| 4 | `template-matching.md` → `templates/CATALOG.md` → apps + patterns → `mvp-assembly.md` | **product repo root** (`maker-flow new <name>`) |
| 6 | `skills/deploy.md` | `maker-flow deploy` (wraps `release/deploy/push-and-route.sh`) |

Optional LLM config notes: `ai-engine/` (not required if the host agent is the LLM).

Read **English** primary files only for contracts.

## 3. Preflight checks (before coding)

- [ ] Confirmed PRO exists for step 4+
- [ ] Template(s) chosen via `templates/index.md`
- [ ] Target path is **product repo** (`maker-flow new <name>` or existing product checkout)
- [ ] Scope matches PRO out-of-scope list

## 4. Smoke template (optional)

Only to verify host Docker works — not a substitute for step 4:

```bash
mkdir -p /tmp/maker-flow-smoke && cp -r templates/apps/go-api /tmp/maker-flow-smoke/_smoke
cd /tmp/maker-flow-smoke/_smoke && cp .env.example .env && docker compose up --build
```
