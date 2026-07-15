# prompts/

Stage input contracts for agents. Prefer reading the matching `skills/` file first.

| File | Step | Agent duty |
|------|------|------------|
| `01-requirement.example.md` | 1 | Capture / normalize human requirement |
| `02-pro-draft.md` | 2 | Draft PRO only (no code) |
| `03-pro-confirmed.example.md` | 3 | Persist human-approved PRO (gate artifact) |
| `04-assemble-mvp.md` | 4 | Match template + assemble into `workspace/` |

## Rules

- Step 2 body: inject requirement into `02-pro-draft.md` (or equivalent message).
- Step 4 body: inject **confirmed** PRO; refuse if gate 3 missing.
- Optional CLI: `./scripts/ai-run.sh prompts/<file>.md`

Skills: `../skills/`.
