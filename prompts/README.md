# prompts/

Stage input contracts for agents. Prefer reading the matching `skills/` file first.

| File | Step | Agent duty |
|------|------|------------|
| `01-requirement.example.md` | 1 | Capture / normalize human requirement |
| `02-pro-draft.md` | 2 | Draft PRO only (no code) |
| [`pro.template.md`](pro.template.md) | 2–3 | **PRO 空白骨架**（输出/定稿结构） |
| [`pro.example.md`](pro.example.md) | 2–3 | **PRO 完整样板**（待办 API） |
| `03-pro-confirmed.example.md` | 3 | Persist human-approved PRO (gate artifact) |
| `04-assemble-mvp.md` | 4 | Match template + assemble into `workspace/` |

## Rules

- Step 2 body: inject requirement into `02-pro-draft.md` (or equivalent message); structure MUST match `pro.template.md` / `skills/pro-generation.md`.
- Prefer reading `pro.example.md` for granularity before drafting.
- Step 4 body: inject **confirmed** PRO; refuse if gate 3 missing.
- Optional CLI: `./scripts/ai-run.sh prompts/<file>.md`

Skills: `../skills/`.
