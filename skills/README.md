# skills/

**English** · [简体中文](README.zh-CN.md)

Authoritative **HOW** contracts for agents. Prompt files say what to ask; skills say what “done” means.

**Catalog entry (humans + AI):** [`CATALOG.md`](CATALOG.md) ← start here

**Agents:** read English primary files only (`*.md`). Do not use `.zh-CN.md` as skill contracts. See [`docs/i18n.md`](../docs/i18n.md).

| Skill | File | Step |
|-------|------|------|
| PRO generation | `pro-generation.md` | 2 |
| Template matching | `template-matching.md` | 4 |
| MVP assembly | `mvp-assembly.md` | 4 |
| Deploy / publish | `deploy.md` | 6 |

## Agent rules

- MUST open [`CATALOG.md`](CATALOG.md) to locate the skill for the active step.
- MUST read the skill file **before** acting.
- MUST treat skill sections labeled MUST / MUST NOT as hard constraints.
- If skill and prompt conflict, **skill wins**.

## Extend

Add `skills/<name>.md`, then register in [`CATALOG.md`](CATALOG.md) + this table + `docs/workflow.md`.
