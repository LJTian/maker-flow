# skills/

Authoritative **HOW** contracts for agents. Prompt files say what to ask; skills say what “done” means.

| Skill | File | Step |
|-------|------|------|
| PRO generation | `pro-generation.md` | 2 |
| Template matching | `template-matching.md` | 4 |
| MVP assembly | `mvp-assembly.md` | 4 |
| Deploy | `deploy.md` | 6 |

## Agent rules

- MUST read the skill for the active step **before** acting.
- MUST treat skill sections labeled MUST / MUST NOT as hard constraints.
- If skill and prompt conflict, **skill wins**.

## Extend

Add a new `skills/<name>.md` and register it in this table + the matching step in `docs/workflow.md`.
