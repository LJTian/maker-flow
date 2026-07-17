# AGENTS.md — factory (maker-flow tool repo)

**English** · [简体中文](AGENTS.zh-CN.md)

**Audience: AI agents working on this factory repo only.**  
For MVP product work, open a product repo (`maker-flow new <name>`) and use its `AGENTS.md` (from [AGENTS.consumer.example.md](AGENTS.consumer.example.md)).

Human intro: [README.md](README.md) · Quick start: [docs/getting-started.md](docs/getting-started.md) · **Product repos:** [docs/consumer-project.md](docs/consumer-project.md) · i18n: [docs/i18n.md](docs/i18n.md)

This repository is an agent playbook for shipping personal MVPs. Humans only provide requirements and approve gates.

**Layout:** parallel **product repos** — read [consumer-project.md](docs/consumer-project.md); use [AGENTS.consumer.example.md](AGENTS.consumer.example.md) in each MVP repo (`maker-flow new <name>`). Assembled MVPs **MUST NOT** be written into this factory repo.

Principle: **heavy infrastructure, light logic.** Prefer templates and skills over inventing scaffolding.

## Language

- Canonical contracts are the **English** `.md` files.
- `.zh-CN.md` siblings are for humans only.
- **MUST NOT** treat Chinese siblings as authoritative for steps, skills, or gates unless the human explicitly asks for Chinese-facing output.

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
| 3 | Human | Approve PRO | — | confirmed PRO → `prompts/03-pro-confirmed.example.md` or project `pro.md` (same shape as `pro.template.md`) |
| 4 | Agent | Match template + assemble MVP | `skills/template-matching.md`, `skills/mvp-assembly.md`, `templates/index.md`, `prompts/04-assemble-mvp.md` | **Product repo root** (`maker-flow new <name>`) |
| 5 | Human | Approve MVP | PRO acceptance criteria | pass/fail |
| 6 | Agent (on approve) | Deploy | `skills/deploy.md`, `release/` | public URL |

Hard gates: **stop at 3 and 5 until human confirms.**

## Layout

```
skills/        # HOW — step SOPs (authoritative for agents)
templates/     # WHAT — searchable scaffolds; start at templates/index.md
prompts/       # inputs / stage contracts
release/       # deploy primitives (nginx, cloudflare, scripts)
scripts/       # helpers (install.sh, maker-flow CLI, check.sh)
docs/          # workflow + architecture contracts
```

## Hard rules

- MUST follow `skills/*` for the active step (English primary files).
- MUST NOT emit implementation code at step 2.
- MUST NOT assemble (step 4) without confirmed PRO (step 3).
- MUST NOT deploy (step 6) without MVP approval (step 5).
- MUST select **one or more** apps via `templates/CATALOG.md` / `templates/index.md` before coding (each app must map to a PRO responsibility).
- MUST resolve Dockerfile fragments via `templates/images/index.md` and **inline** them into app Dockerfiles (upstream `FROM` only — never `FROM maker-flow/*` private tags).
- MAY attach 0–N patterns from `templates/patterns/` (copy into the app that needs them; never deploy alone).
- MUST write assembled projects under the **product repo root** (`AGENTS.consumer.example.md`; create with `maker-flow new <name>`). Multi-app: `<product-root>/<app-id>/`.
- MUST NOT copy the `templates/images/` tree into a product repo; compose fragment lines into the product Dockerfile only.
- Factory smoke tests: `maker-flow new <name>` or copy a template to `/tmp/maker-flow-smoke/` — **MUST NOT** write assembled MVPs into this factory repo.
- Prefer **container builds** (`docker compose up --build`); do not require host Go toolchain for verification.

## Contracts

| Topic | Path |
|-------|------|
| Workflow | `docs/workflow.md` |
| Architecture | `docs/architecture.md` |
| Agent bootstrap | `docs/agent-bootstrap.md` |
| Skills index | `skills/CATALOG.md` |
| Template catalog | `templates/CATALOG.md` |
| Pattern catalog | `templates/patterns/index.md` |
| i18n | `docs/i18n.md` |
| Consumer / product repos | `docs/consumer-project.md` · `AGENTS.consumer.example.md` |
