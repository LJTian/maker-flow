# Step 2 — Draft PRO

**English** · [简体中文](02-pro-draft.zh-CN.md)

Read `skills/pro-generation.md` before running.  
Match output structure to [`pro.template.md`](pro.template.md); use [`pro.example.md`](pro.example.md) for granularity.  
Sync the requirement from `prompts/01-requirement.example.md` into the section below.

---

## Role

You are a product manager and architect. From the user requirement, output a **PRO** for human confirmation. **MUST NOT emit implementation code.**

## User requirement

(Paste from 01-requirement.example.md)

Build a mini “todo” API: create, complete, and list — no user system.

## Required sections

Follow `skills/pro-generation.md` and `pro.template.md` strictly:

1. Summary (including out-of-scope list)
2. Business flow
3. Data model
4. API / interface contract (split by app when multi-app)
5. Acceptance criteria (checkbox list)
6. Template retrieval hints (preferred apps 1–N / patterns 0–N — not final picks)

## Start

Output Markdown section by section in the order above.
