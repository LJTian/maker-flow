# Workflow (agent contract)

Authoritative state machine for Maker Flow agents.

```
① HUMAN: requirement
        │
        ▼
② AGENT: draft PRO ──► ③ HUMAN: confirm PRO  [GATE]
        │
        ▼
④ AGENT: match template + assemble MVP
        │
        ▼
⑤ HUMAN: confirm MVP  [GATE]
        │
        ▼
⑥ AGENT: deploy
```

## Step contracts

### 1 — Requirement (human)

Input: short natural-language requirement.  
Agent: store/copy into the step-2 prompt user-requirement section if using `prompts/`.

### 2 — Draft PRO (agent)

- **MUST read:** `skills/pro-generation.md`
- **MAY use:** `prompts/02-pro-draft.md` as the prompt body
- **MUST output:** PRO sections defined by the skill
- **MUST NOT:** write application code, pick a final template, or create `workspace/`

### 3 — Confirm PRO (human gate)

- Agent presents PRO and waits.
- On approval, persist confirmed PRO to `prompts/03-pro-confirmed.example.md` or `workspace/<name>/pro.md`.
- **MUST NOT** proceed to step 4 without explicit human approval.

### 4 — Assemble MVP (agent)

- **MUST read (in order):**
  1. `skills/template-matching.md`
  2. `templates/CATALOG.md` → `templates/index.md` → `templates/images/index.md`
  3. `skills/mvp-assembly.md`
- **MAY use:** `prompts/04-assemble-mvp.md`
- **MUST:** select one template ID, copy to `workspace/<name>/`, implement only PRO scope
- **MUST NOT:** invent scaffolding outside templates; deploy

### 5 — Confirm MVP (human gate)

Agent (or human) runs:

```bash
cd workspace/<name>
cp -n .env.example .env
docker compose up --build
curl -sf http://localhost:8080/health
```

Verify PRO acceptance criteria. On fail: iterate step 4, or return to step 3 if scope is wrong.  
**MUST NOT** deploy until approved.

### 6 — Deploy (agent)

- **MUST read:** `skills/deploy.md`
- **MUST use:** `release/deploy/push-and-route.sh` + nginx/cloudflare assets under `release/`
- Prerequisites: human-approved MVP; deploy host credentials available in env

## Roles

| Role | Allowed steps |
|------|----------------|
| Human | 1, 3, 5 (and may trigger 6) |
| Agent | 2, 4, 6 (6 only after gate 5) |

## Related

- `docs/architecture.md`
- `docs/agent-bootstrap.md`
- `docs/getting-started.md` (human)
- `AGENTS.md`
- `skills/README.md`
- `templates/index.md`
