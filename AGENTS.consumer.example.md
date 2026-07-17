# AGENTS.md — consumer product repo

**English** · [简体中文](AGENTS.consumer.example.zh-CN.md)

**Copy this file to your MVP product repository as `AGENTS.md`.**  
Do not commit it inside maker-flow as your product contract.

Factory (read-only): resolve with `maker-flow root` (default `~/.maker-flow`).  
Guide: [`docs/consumer-project.md`](docs/consumer-project.md)

---

## Context

- **This repo** = one MVP product (business code only).
- **maker-flow** = shared factory at `$MAKER_FLOW_ROOT` (skills, templates, release).
- **MUST NOT** copy entire `skills/` / `templates/` / `release/` into this repo.
- **MUST NOT** modify files under `$MAKER_FLOW_ROOT`.

## Configuration (edit per project)

```bash
# Portable default (expand ~). Override only for a custom install dir.
# Resolve at runtime: maker-flow root
MAKER_FLOW_ROOT=~/.maker-flow

# This product (kebab-case)
PRODUCT_NAME=my-todo
```

## Output paths

| Artifact | Path in this repo |
|----------|-------------------|
| Confirmed PRO | `pro.md` |
| Single app | `./` or `./<app-id>/` |
| Multi-app | `./api/`, `./web/`, … per PRO |
| Publish cwd | this repo root |

## Agent entry

1. Read `$MAKER_FLOW_ROOT/docs/workflow.md` (gates unchanged). Expand `~` in `MAKER_FLOW_ROOT` if needed.
2. Load step skill from `$MAKER_FLOW_ROOT/skills/CATALOG.md`.
3. Match templates via `$MAKER_FLOW_ROOT/templates/CATALOG.md`.
4. **Write only under this product repo** — never under `$MAKER_FLOW_ROOT`.

## Six-step state machine

| Step | Action | Reads | Writes |
|------|--------|-------|--------|
| 1 | Human: requirement | — | chat / notes |
| 2 | Agent: draft PRO | `$MAKER_FLOW_ROOT/skills/pro-generation.md`, `prompts/pro.template.md` | chat (no code) |
| 3 | Human: approve PRO | — | **`pro.md`** |
| 4 | Agent: assemble | `template-matching`, `mvp-assembly`, `templates/` | **this repo** (`./`, `./api/`, …) |
| 5 | Human: accept MVP | `pro.md` criteria | — |
| 6 | Agent: publish | `$MAKER_FLOW_ROOT/skills/deploy.md`, `prompts/06-publish.md`, `release/publish/` | public URL(s) after **human chooses target(s) in chat** |

Hard gates: **stop at 3 and 5 until human confirms.**  
Step 6: **ask where to publish** (Pages / Vercel / VPS / split). Do **not** tell the human to run `maker-flow deploy` (agent-internal only).

## Assembly rules

- Copy apps from `$MAKER_FLOW_ROOT/templates/apps/<id>/` into this repo.
- Patterns: copy into the app that needs them; never deploy patterns alone.
- Go Dockerfiles: compose fragments from `$MAKER_FLOW_ROOT/templates/images/` (inline into the app Dockerfile; no pre-build step).
- After copy, rewrite each Go app `go.mod` `module` line to a product path (e.g. `github.com/<you>/${PRODUCT_NAME}` or `github.com/<you>/${PRODUCT_NAME}/api`) — do not keep `.../maker-flow/templates/...`.
- Prefer `docker compose up --build` in **this repo** for acceptance.

## Publish

After step 5, load `$MAKER_FLOW_ROOT/prompts/06-publish.md` and confirm targets with the human, then execute `$MAKER_FLOW_ROOT/release/publish/<target>.md`.

## First message template

```
Read this AGENTS.md and $MAKER_FLOW_ROOT/docs/workflow.md.
MAKER_FLOW_ROOT=~/.maker-flow (or: maker-flow root). We are in product repo "${PRODUCT_NAME}".
Start at step ① with my requirement: …
```
