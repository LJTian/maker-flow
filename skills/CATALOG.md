# Skills catalog

**English** · [简体中文](CATALOG.zh-CN.md)

> **For AI:** before the matching step, **MUST** read this file to locate the skill, then open the full skill text.  
> **For humans:** see which skill constrains each pipeline step at a glance.  
> **Agents:** read English primary files only. Do not use `.zh-CN.md` as skill contracts. See [`docs/i18n.md`](../docs/i18n.md).

---

## Overview

| Step | Skill id | File | One-liner |
|:----:|----------|------|-----------|
| 2 | `pro-generation` | [`pro-generation.md`](pro-generation.md) | PRO only — no code |
| 4 | `template-matching` | [`template-matching.md`](template-matching.md) | Pick 1–N apps + 0–N patterns + images |
| 4 | `mvp-assembly` | [`mvp-assembly.md`](mvp-assembly.md) | Copy apps, merge patterns, run in containers |
| 6 | `deploy` | [`deploy.md`](deploy.md) | Ship after acceptance |

**Prompts:** [`../prompts/`](../prompts/) · **Template catalog:** [`../templates/CATALOG.md`](../templates/CATALOG.md) · **Patterns:** [`../templates/patterns/index.md`](../templates/patterns/index.md)

---

## Load by step

| Current step | Required reads |
|--------------|----------------|
| 2 — draft PRO | `pro-generation.md` + `prompts/02-pro-draft.md`; structure `prompts/pro.template.md`, sample `prompts/pro.example.md` |
| 4 — assemble | `template-matching.md` → `templates/CATALOG.md` → apps + patterns → `mvp-assembly.md` |
| 6 — deploy | `deploy.md` + `release/` |

Hard gates: no confirmed PRO at step 3 → MUST NOT run step 4; no MVP approval at step 5 → MUST NOT run step 6. See [`docs/workflow.md`](../docs/workflow.md).

---

## Skill contract summary

| Skill | MUST | MUST NOT |
|-------|------|----------|
| PRO generation | Include summary / flow / model / API / acceptance | Emit implementation code or final template picks |
| Template matching | Read CATALOG + index; list image deps | Invent scaffolding; select before PRO confirmed |
| MVP assembly | Write under `workspace/<name>/`; build images first | Edit image base directory; deploy in this step |
| Deploy | Follow `release/` scripts and port pool | Skip local / acceptance gates |

---

## Registration

When adding a skill, update **all** of:

1. This file (overview table)
2. [`README.md`](README.md) (agent rules table)
3. Matching step in [`docs/workflow.md`](../docs/workflow.md)
4. [`AGENTS.md`](../AGENTS.md) state-machine table (if a new step touchpoint)
