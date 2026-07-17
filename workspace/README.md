# workspace/

**English** · [简体中文](README.zh-CN.md)

Local **smoke / scratch** target when working inside the maker-flow factory repo.

## Recommended production layout

For real MVPs, use a **separate private product repository** and parallel directories:

→ [`docs/consumer-project.md`](../docs/consumer-project.md)  
→ Copy [`AGENTS.consumer.example.md`](../AGENTS.consumer.example.md) to your product repo as `AGENTS.md`

## Rules (factory repo only)

- MAY create `workspace/<kebab-name>/` for quick template smoke tests.
- MUST NOT commit generated apps (directory is gitignored except this README).
- MUST NOT treat `workspace/` as the long-term home for business MVPs you need to keep private.

Verify with `docker compose up --build` and PRO acceptance criteria (step 5).
