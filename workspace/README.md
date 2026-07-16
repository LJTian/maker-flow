# workspace/

**English** · [简体中文](README.zh-CN.md)

Agent **write target** for step-4 assembled MVPs.

## Rules

- MUST create `workspace/<kebab-name>/` for each MVP.
- MUST NOT commit generated apps to git by default (directory is gitignored except README / `.gitkeep`).
- Verify with `docker compose up --build` and PRO acceptance criteria (step 5).

Each subdirectory is an independent runnable project.
