# AGENTS.md

**Audience: AI agents.** Human-oriented intro: [README.md](README.md) · Quick start: [docs/getting-started.md](docs/getting-started.md)

This repository is an agent playbook for shipping personal MVPs. Humans only provide requirements and approve gates.

Principle: **heavy infrastructure, light logic.** Prefer templates and skills over inventing scaffolding.

## Agent entry

1. Read `docs/workflow.md` (state machine + hard gates).
2. Load the skill for the current step from `skills/` (start at `skills/CATALOG.md`).
3. Use inputs from `prompts/` (or user message equivalent).
4. For step 4, open `templates/CATALOG.md` before matching.
5. Write outputs to the paths defined by the skill.

Do **not** skip gates. Do **not** invent a new stack when `templates/index.md` has a match.

## Six-step state machine

| Step | Actor | Action | Required reads | Output |
|------|-------|--------|----------------|--------|
| 1 | Human | Provide requirement | — | requirement text |
| 2 | Agent | Draft PRO | `skills/pro-generation.md`, `prompts/02-pro-draft.md`, `prompts/pro.template.md` | PRO markdown (no code) |
| 3 | Human | Approve PRO | — | confirmed PRO → `prompts/03-pro-confirmed.example.md` or project `pro.md`（结构同 `pro.template.md`） |
| 4 | Agent | Match template + assemble MVP | `skills/template-matching.md`, `skills/mvp-assembly.md`, `templates/index.md`, `prompts/04-assemble-mvp.md` | `workspace/<name>/` |
| 5 | Human | Approve MVP | PRO acceptance criteria | pass/fail |
| 6 | Agent (on approve) | Deploy | `skills/deploy.md`, `release/` | public URL |

Hard gates: **stop at 3 and 5 until human confirms.**

## Layout

```
ai-engine/     # optional LLM transport config (OpenAI-compatible)
skills/        # HOW — step SOPs (authoritative for agents)
templates/     # WHAT — searchable scaffolds; start at templates/index.md
prompts/       # inputs / stage contracts
workspace/     # agent write target for assembled MVPs
release/       # deploy primitives (nginx, cloudflare, scripts)
scripts/       # helpers (e.g. ai-run.sh)
docs/          # workflow + architecture contracts
```

## Hard rules

- MUST follow `skills/*` for the active step.
- MUST NOT emit implementation code at step 2.
- MUST NOT assemble (step 4) without confirmed PRO (step 3).
- MUST NOT deploy (step 6) without MVP approval (step 5).
- MUST select **one or more** apps via `templates/CATALOG.md` / `templates/index.md` before coding (each app must map to a PRO responsibility).
- MUST resolve image bases via `templates/images/index.md` and run `./scripts/build-images.sh` when bases are missing.
- MAY attach 0–N patterns from `templates/patterns/` (copy into the app that needs them; never deploy alone).
- MUST write assembled projects under `workspace/<kebab-name>/` (multi-app: `workspace/<name>/<app-id>/`).
- MUST NOT copy `templates/images/` into `workspace/`; inherit via `FROM` only.
- Prefer **container builds** (`./scripts/build-images.sh` then `docker compose up --build`); do not require host Go toolchain for verification.

## Contracts

| Topic | Path |
|-------|------|
| Workflow | `docs/workflow.md` |
| Architecture | `docs/architecture.md` |
| Agent bootstrap | `docs/agent-bootstrap.md` |
| Skills index | `skills/CATALOG.md` |
| Template catalog | `templates/CATALOG.md` |
| Pattern catalog | `templates/patterns/index.md` |
