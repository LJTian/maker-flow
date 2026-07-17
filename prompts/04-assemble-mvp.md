# Step 4 — Match templates and assemble MVP

**English** · [简体中文](04-assemble-mvp.zh-CN.md)

Read in order before running:

1. `skills/template-matching.md`
2. `skills/mvp-assembly.md`

Confirm the PRO in `prompts/03-pro-confirmed.example.md` is marked **Confirmed**. Template overview: `templates/CATALOG.md`.

---

## Role

You are a full-stack engineer. From the **confirmed PRO**, match templates and assemble a runnable MVP.

## Confirmed PRO

(Paste the finalized PRO from 03-pro-confirmed.example.md)

## Template catalog

Search using `templates/index.md` in this repo.

## Required output

### Part A — Template matching (follow template-matching skill)

- Selected template IDs, paths, rationale

### Part B — MVP assembly (follow mvp-assembly skill)

- Output directory: `workspace/<project-name>/`
- Directory tree (changed files)
- Full contents of key files
- Local run commands

## Constraints

- MUST NOT change the confirmed PRO scope
- MUST NOT deploy to the public internet (that is step 6)
- Reuse template middleware and project structure

## Start

Output Part A first, then Part B.
