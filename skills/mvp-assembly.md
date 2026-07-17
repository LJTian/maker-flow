# MVP assembly skill

**English** · [简体中文](mvp-assembly.zh-CN.md)

**Step:** 4 — AI assembles an MVP from the PRO and selected templates  
**Prerequisite:** `skills/template-matching.md` has produced selected templates

## Goal

Produce a **locally runnable** minimal project with no features outside the PRO.

## Output directory

```
workspace/<project-name>/
```

`<project-name>` is derived from the PRO summary (kebab-case, e.g. `todo-api`).

## Assembly steps

1. **Compose Dockerfiles** — if apps need Go image fragments, read `templates/images/index.md` and inline `go-builder` / `go-runtime` (or keep the already-composed Dockerfile from the app template).  
   Use upstream images only (`golang:…`, `alpine:…`). **MUST NOT** `FROM maker-flow/*` or pre-build local tags.  
   **MUST NOT** copy the `templates/images/` tree into `workspace/` — only fragment lines in the product Dockerfile.
2. **Copy templates** — copy each selected `templates/apps/<id>/` into the workspace:
   - Single app: `workspace/<project-name>/`
   - Multi-app: `workspace/<project-name>/<id>/` (e.g. `api/`, `worker/`, `cli/`), or the layout agreed in the PRO
3. **Merge patterns (optional)** — copy pattern packages into `internal/...` of the **app that needs them** and wire them up
4. **Config** — per app `.env.example` → `.env`; avoid port / name clashes across apps
5. **Implement business logic** — per PRO and each app stack (Gin / Cobra / worker)
6. **Update compose** — multi-service: root compose with multiple build contexts, or per-app compose
7. **Self-check** — walk PRO acceptance criteria (cover every selected app)

## Code principles

- Keep middleware / logging / entry style from each selected app template
- For multi-app, keep process boundaries and communication clear (HTTP / queue placeholders); MUST NOT collapse into one process
- Keep pattern package names clear and testable when merging
- MUST NOT add dependencies not listed in the PRO
- Prefer small diffs; runnable first
- **Deps and builds happen inside containers** (`docker compose up --build` per app); host Go toolchain is not required

## Output format

1. Directory tree (new/changed files only)
2. Full contents or clear diffs for each key file
3. Local run commands:

```bash
cd workspace/<project-name>
cp .env.example .env
docker compose up --build
```

## MUST NOT

- MUST NOT deploy `templates/patterns/` alone as a public service
- MUST NOT rewrite template infrastructure; MUST NOT edit `templates/images/` fragments unless the task explicitly requires it
- MUST NOT duplicate `apk` packages already provided by the inlined image fragment (e.g. ca-certificates)
- MUST NOT ship features outside the PRO
- MUST NOT deploy in this step (that is step 6)
