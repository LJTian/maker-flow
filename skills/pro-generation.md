# PRO generation skill

**English** · [简体中文](pro-generation.zh-CN.md)

**Step:** 2 — AI drafts PRO from the requirement  
**Prompt:** `prompts/02-pro-draft.md`  
**Blank skeleton:** [`prompts/pro.template.md`](../prompts/pro.template.md)  
**Full sample:** [`prompts/pro.example.md`](../prompts/pro.example.md)

## Goal

Turn a one-line human requirement into a confirmable **PRO** for step 3 human review.

## MUST NOT

- MUST NOT emit a full project codebase
- MUST NOT finalize template selection (leave for step 4; retrieval hints are clues only)
- MUST NOT assume the user has already confirmed scope

## Required PRO sections

Output structure MUST match `prompts/pro.template.md`. Use `prompts/pro.example.md` for granularity.

### 1. Summary

- One-sentence goal
- MVP scope (finishable in 1–2 days)
- Explicit **out of scope** list

### 2. Business flow

Numbered steps covering the main path and boundaries (auth needed? data ownership?).

### 3. Data model

Markdown table: field, type, description, constraints.  
Optional: `CREATE TABLE` statements when the PRO includes persistence. If no persistence, say so explicitly.

### 4. API / interface contract

Per endpoint or command: `METHOD /path` (or CLI subcommand), request/response examples, main error codes.  
At least: health check (if applicable) + 2–4 core business interfaces.  
For multi-app PROs, split subsections by app.

### 5. Acceptance criteria

Checkbox checklist for step 5. Examples:

- [ ] `GET /health` returns 200
- [ ] Can create and list todos
- [ ] …

### 6. Template retrieval hints (optional)

Clues for step 4 — not final picks:

- Preferred **apps (1–N)** (e.g. `go-api`, or `go-api` + `go-worker`)
- Preferred **patterns (0–N)** (or “none”)
- DB needed? expected QPS / complexity

## Quality constraints

- Keep scope at MVP; MUST NOT require K8s, message queues, or microservice splits
- Default to no auth unless the requirement explicitly asks
- Align terminology with tags in `templates/index.md` / `templates/patterns/index.md` for later retrieval

## Output format

Markdown with section headings matching `pro.template.md`, so humans can diff and annotate.
