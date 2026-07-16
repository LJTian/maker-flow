# skills/

Authoritative **HOW** contracts for agents. Prompt files say what to ask; skills say what “done” means.

**检索入口（人 + AI）：** [`CATALOG.md`](CATALOG.md) ← 先读这个

| Skill | File | Step |
|-------|------|------|
| PRO generation | `pro-generation.md` | 2 |
| Template matching | `template-matching.md` | 4 |
| MVP assembly | `mvp-assembly.md` | 4 |
| Deploy | `deploy.md` | 6 |

## Agent rules

- MUST open [`CATALOG.md`](CATALOG.md) to locate the skill for the active step.
- MUST read the skill file **before** acting.
- MUST treat skill sections labeled MUST / MUST NOT as hard constraints.
- If skill and prompt conflict, **skill wins**.

## Extend

Add `skills/<name>.md`, then register in [`CATALOG.md`](CATALOG.md) + this table + `docs/workflow.md`.
