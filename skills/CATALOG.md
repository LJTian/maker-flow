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
| 5 | `mvp-acceptance` | [`mvp-acceptance.md`](mvp-acceptance.md) | Local evidence vs PRO criteria; human gate |
| 6 | `deploy` | [`deploy.md`](deploy.md) | Dialogue: choose publish target(s), then ship |
| 6 | `publish-vps-gateway` | [`publish-vps-gateway.md`](publish-vps-gateway.md) | VPS + shared Nginx gateway |
| 6 | `publish-cloudflare-pages` | [`publish-cloudflare-pages.md`](publish-cloudflare-pages.md) | Cloudflare Pages static / SPA |
| 6 | `publish-github-pages` | [`publish-github-pages.md`](publish-github-pages.md) | GitHub Pages static / SPA |
| 6 | `publish-vercel` | [`publish-vercel.md`](publish-vercel.md) | Vercel static / SPA |
| 6 | `publish-split-web-api` | [`publish-split-web-api.md`](publish-split-web-api.md) | SPA on Pages/Vercel + API on VPS |
| 6 | `cloudflare-dns` | [`cloudflare-dns.md`](cloudflare-dns.md) | Cloudflare DNS CRUD via `dns.sh` (with deploy when DNS needed) |

**Prompts:** [`../prompts/`](../prompts/) · **Template catalog:** [`../templates/CATALOG.md`](../templates/CATALOG.md) · **Patterns:** [`../templates/patterns/index.md`](../templates/patterns/index.md)

---

## Load by step

| Current step | Required reads |
|--------------|----------------|
| 2 — draft PRO | `pro-generation.md` + `prompts/02-pro-draft.md`; structure `prompts/pro.template.md`, sample `prompts/pro.example.md` |
| 4 — assemble | `template-matching.md` → `templates/CATALOG.md` → apps + patterns → `mvp-assembly.md` |
| 5 — accept MVP | `mvp-acceptance.md` + `prompts/05-accept-mvp.md` + product `pro.md` |
| 6 — publish | `deploy.md` + `prompts/06-publish.md` + chosen `publish-<target>.md` |
| 6 — Cloudflare DNS | `cloudflare-dns.md` + `release/cloudflare/dns-api.md` (when listing or changing DNS) |

Hard gates: no confirmed PRO at step 3 → MUST NOT run step 4; no MVP approval at step 5 → MUST NOT run step 6. See [`docs/workflow.md`](../docs/workflow.md).

---

## Skill contract summary

| Skill | MUST | MUST NOT |
|-------|------|----------|
| PRO generation | Include summary / flow / model / API / acceptance | Emit implementation code or final template picks |
| Template matching | Read CATALOG + index; list image deps | Invent scaffolding; select before PRO confirmed |
| MVP assembly | Write under **product repo root**; compose Dockerfiles from image fragments | Copy `templates/images/` tree; deploy in this step; write into factory repo |
| MVP acceptance | Walk every PRO acceptance item with evidence; wait for human approve | Pass gate on `/health` alone; deploy before human yes; invent criteria |
| Deploy (publish) | Ask human for target(s); follow `publish-<target>.md`; VPS CLI is agent-internal | Skip gates; tell humans to run `maker-flow deploy`; ship static apps to Pages with a required DB API |
| Split web+API | API→VPS then SPA with public `VITE_API_BASE_URL`; dual-URL verify | Pages-only for Go API; rebuild SPA without public API origin |
| Cloudflare DNS | Use `dns.sh` CRUD; confirm token + zone id; list before ambiguous delete | Invent tokens; use dashboard when API creds exist; DNS before step-5 gate |

---

## Registration

When adding a skill, update **all** of:

1. This file (overview table)
2. [`README.md`](README.md) (agent rules table)
3. Matching step in [`docs/workflow.md`](../docs/workflow.md)
4. [`AGENTS.md`](../AGENTS.md) state-machine table (if a new step touchpoint)
