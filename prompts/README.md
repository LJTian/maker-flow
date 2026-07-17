# prompts/

**English** · [简体中文](README.zh-CN.md)

Stage input contracts for agents. Prefer reading the matching `skills/` file first.

**Agents:** read English primary files only (`*.md`). Do not use `.zh-CN.md` as prompt/skill contracts. See [`docs/i18n.md`](../docs/i18n.md).

| File | Step | Agent duty |
|------|------|------------|
| `01-requirement.example.md` | 1 | Capture / normalize human requirement |
| `02-pro-draft.md` | 2 | Draft PRO only (no code) |
| [`pro.template.md`](pro.template.md) | 2–3 | **PRO blank skeleton** (output / confirm structure) |
| [`pro.example.md`](pro.example.md) | 2–3 | **PRO full sample** (todo API) |
| `03-pro-confirmed.example.md` | 3 | Persist human-approved PRO (gate artifact) |
| `04-assemble-mvp.md` | 4 | Match template + assemble (output = **product repo**; see consumer guide) |
| [`../AGENTS.consumer.example.md`](../AGENTS.consumer.example.md) | — | **Product repo** `AGENTS.md` template (copy out) |

## Rules

- Step 2 body: inject requirement into `02-pro-draft.md` (or equivalent message); structure MUST match `pro.template.md` / `skills/pro-generation.md`.
- Prefer reading `pro.example.md` for granularity before drafting.
- Step 4 body: inject **confirmed** PRO; refuse if gate 3 is missing.

Skills: `../skills/`.
