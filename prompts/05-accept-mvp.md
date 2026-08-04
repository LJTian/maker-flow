# Step 5 — accept MVP (local gate)

**English** · [简体中文](05-accept-mvp.zh-CN.md)

Agent loads this after step 4 assembly. Skill: `skills/mvp-acceptance.md`.

---

## Role

You run **local acceptance** against the confirmed PRO. The human only **approves or rejects** the gate. Do not publish.

## Inputs

- Product repo root (cwd for all commands)
- Confirmed PRO: `pro.md` (especially § Acceptance criteria and § API / interfaces)
- Running instructions left by assembly (compose, ports, `.env.example`)

## Required work

Follow `skills/mvp-acceptance.md` in order:

1. Bring up with `docker compose up --build`
2. Shape-aware baseline smoke (API / SPA / CLI / worker / multi-app)
3. Execute **every** acceptance checkbox with a concrete command
4. Return an **evidence package**, then ask: Approve step 5?

## Output format

```text
Step 5 evidence — <PRODUCT_NAME>
Baseline: …
Acceptance:
  [x] …
  [ ] … → FAIL …
Recommendation: pass gate | iterate step 4 | return to step 3
Approve step 5? (yes / no + notes)
```

## Constraints

- MUST NOT deploy or start step 6 until the human says the MVP is accepted
- MUST NOT invent criteria outside `pro.md`
- Prefer container commands; do not require host Go/Node installs

## Start

Read `pro.md`, detect compose services, then run the skill procedure.
