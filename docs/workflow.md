# Workflow (agent contract)

**English** · [简体中文](workflow.zh-CN.md)

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
- **MUST follow structure:** `prompts/pro.template.md` (see `prompts/pro.example.md` for granularity)
- **MAY use:** `prompts/02-pro-draft.md` as the prompt body
- **MUST output:** PRO sections defined by the skill / template
- **MUST NOT:** write application code, pick a final template, or assemble into the factory repo

### 3 — Confirm PRO (human gate)

- Agent presents PRO and waits.
- On approval, persist confirmed PRO to `pro.md` in the **product repo** (or `prompts/03-pro-confirmed.example.md` for factory-local examples; same section shape as `pro.template.md`).
- **MUST NOT** proceed to step 4 without explicit human approval.

### 4 — Assemble MVP (agent)

- **MUST read (in order):**
  1. `skills/template-matching.md`
  2. `templates/CATALOG.md` → `templates/index.md` → `templates/patterns/index.md` → `templates/images/index.md`
  3. `skills/mvp-assembly.md`
- **MAY use:** `prompts/04-assemble-mvp.md`
- **MUST:** select **one or more** app IDs, copy to the **product repo root** (multi-app: `<product-root>/<app-id>/`), implement only PRO scope
- **MUST NOT:** invent scaffolding outside templates; deploy

### 5 — Confirm MVP (human gate)

- **MUST read:** `skills/mvp-acceptance.md`
- **MAY use:** `prompts/05-accept-mvp.md`
- Agent runs verification in the **product repo** (shape-aware: API / SPA / CLI / worker / multi-app), walks every PRO acceptance criterion, and presents an evidence package.
- Human only approves/rejects at this gate.
- On fail: iterate step 4, or return to step 3 if scope is wrong.
- **MUST NOT** deploy until approved.

Typical bring-up (agent expands per skill):

```bash
cd ~/projects/<name>   # product repo (maker-flow new <name>)
cp -n .env.example .env
docker compose up --build
# then criterion-by-criterion checks — see skills/mvp-acceptance.md
```

### 6 — Publish (agent)

- **MUST read:** `skills/deploy.md`, `prompts/06-publish.md`, and the chosen `skills/publish-<target>.md` (see [`skills/CATALOG.md`](skills/CATALOG.md))
- **MUST also read** `skills/cloudflare-dns.md` when changing Cloudflare DNS (custom domain, VPS A/AAAA, DDNS, or human asks to manage records)
- **MUST** ask the human which publish target(s) to use (Cloudflare Pages / GitHub Pages / Vercel / VPS gateway / split)
- **MUST NOT** instruct the human to run `maker-flow deploy` (agent-internal helper for VPS only)
- Execute the matching `skills/publish-<target>.md`
- Prerequisites: human-approved MVP; credentials / host access as required by the chosen target(s)

## Failure paths

- **Step 3 (PRO gate) rejected by human:** Agent must revise the draft in-place or generate a new draft (`v2`) based on feedback, then wait for approval again. DO NOT proceed to Step 4 until the human confirms the revised PRO.
- **Step 4 (Assemble) encounters ambiguous PRO:** If the Agent realizes during assembly that the PRO is contradictory or underspecified, the Agent MUST pause and return to Step 3, requesting human clarification before continuing to write code.
- **Step 5 (MVP gate) rejected by human:**
  - If the code has a bug or deviates from the PRO: iterate in Step 4 and fix the code.
  - If the human realizes the PRO itself was wrong (scope change): return to Step 3, update the `pro.md` in the product repo, then return to Step 4 to adjust the code.
- **Step 6 (Publish) failure:** Agent MUST follow the rollback instructions specified in the respective `skills/publish-<target>.md`. If none exist, ensure the system is left in a stable state and ask the human for help.

## Roles

| Role | Allowed steps |
|------|----------------|
| Human | 1, 3, 5 (approve/reject; and may trigger 6) |
| Agent | 2, 4, 5 (run acceptance evidence), 6 (only after gate 5) |

## Related

- `docs/architecture.md`
- `docs/agent-bootstrap.md`
- `docs/getting-started.md` (human)
- `docs/i18n.md`
- `AGENTS.md`
- `skills/README.md`
- `templates/index.md`
