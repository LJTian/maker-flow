# Roadmap — next plans

**English** · [简体中文](roadmap.zh-CN.md)

Living backlog for the **factory repo** (skills, templates, release scripts, docs).  
Not a product-feature list for any one MVP.

**Last reviewed:** 2026-08-11

## What is already solid

The six-step pipeline is usable end-to-end for Go API + Vite SPA MVPs:

| Area | Status |
|------|--------|
| Workflow + hard gates | `docs/workflow.md`, `AGENTS.md` |
| Skills (PRO → match → assemble → accept → publish) | `skills/CATALOG.md` |
| Apps | `go-api`, `go-cli`, `go-worker`, `web-vite` |
| Patterns | 14 packages under `templates/patterns/` (incl. `persistence-sqlx`, auth, payments, cron, …) |
| Multi-app layout | `templates/layouts/web-api/` + `skills/publish-split-web-api.md` |
| CI | `scripts/check.sh`, pattern `go test`, app `docker compose` / `docker build` |
| Live example | [static intro → GitHub Pages](examples/static-intro-github-pages.md) |

Earlier gaps (acceptance skill, sqlx pattern, split publish, compose CI) are **done**. What follows is the remaining backlog.

---

## Priority 1 — high impact on real incubation

Ship these first. They unblock “idea → public validation” for more shapes than a single static landing page.

| ID | Item | Gap today | Plan |
|----|------|-----------|------|
| P1-1 | **More end-to-end examples** | Only one live walkthrough (`web-vite` → GitHub Pages) | Add walkthroughs (prefer live URLs): `go-api` → VPS gateway; `web-api` split (Pages/Vercel SPA + VPS API). Register under `docs/examples/` |
| P1-2 | **Static publish script parity** | VPS has `release/deploy/push-and-route.sh`; Pages / Vercel / GH Pages are skill + manual CLI only | Add thin agent-callable scripts under `release/publish/` (or sibling dirs) wrapping build + upload; keep SOPs in `skills/publish-*.md` |
| P1-3 | **Auth + DB assembly sample** | `auth-oauth-jwt` README still has `TODO: Save to DB` | Document + optional snippet wiring OAuth user → `persistence-sqlx` (users table, find-or-create, JWT subject) |
| P1-4 | **Payment + DB assembly sample** | `payment-lemonsqueezy` has `TODO: Upgrade VIP` | Same pattern: webhook → sqlx upgrade path; keep MoR flow in README |
| P1-5 | **Payment catalog honesty** | Tags include `stripe` / Alipay / WeChat but code is Lemon Squeezy only | Either implement extra providers **or** narrow tags/copy so agents do not mis-select |

**Suggested order:** P1-1 → P1-2 → P1-3 → P1-4 → P1-5.

---

## Priority 2 — factory maturity

| ID | Item | Gap today | Plan |
|----|------|-----------|------|
| P2-1 | **Versioned release** | CLI `VERSION=0.5.0`, `CHANGELOG` still all under `[Unreleased]`, no git tags | Cut first tagged release; move Unreleased notes into a version section; keep semver + Keep a Changelog |
| P2-2 | **Chinese doc parity** | Linked but missing: `CONTRIBUTING.zh-CN.md`; missing `skills/publish-*.zh-CN.md`, `prompts/pro-multi.example.zh-CN.md`, `CHANGELOG.zh-CN.md` (optional) | Ship ZH siblings for user-facing / human docs; English remains the agent contract ([i18n.md](i18n.md)) |
| P2-3 | **`web-vite` tests** | CONTRIBUTING requires app tests; Vite template has none | Add a minimal Vitest (or equivalent) smoke test; wire CI if cheap |
| P2-4 | **Pattern `go.sum` coverage** | Many patterns lack committed `go.sum`; CI relies on `go mod tidy` | Commit `go.sum` for every Go pattern (match `cron-scheduler` / `persistence-sqlx`) |
| P2-5 | **CONTRIBUTING ↔ reality** | CONTRIBUTING says skills need YAML frontmatter; current skills have none | Either add frontmatter to skills **or** relax CONTRIBUTING — pick one and stick to it |
| P2-6 | **Stack boundary clarity** | Catalog is Go + Vite only; easy to assume Python/Node APIs exist | Explicit “supported stacks / not in scope” note in `templates/CATALOG.md` (and README if needed). New stacks only when someone owns CI + acceptance |

---

## Priority 3 — polish (nice to have)

| ID | Item | Plan |
|----|------|------|
| P3-1 | Light auth pattern | Optional `api-key` or signed-cookie pattern for “personal tool login” without full OAuth |
| P3-2 | `.gitignore` cleanup | Deduplicate `.gomodcache/` / `.gocache/` entries |
| P3-3 | `ai-engine/` | Keep optional/docs-only unless Cursor-free LLM transport becomes a real need |
| P3-4 | Queue backends | Stay out of default PRO scope (no Redis/Asynq unless a real product needs it) |

---

## Explicit non-goals (for now)

- Assembling MVPs **into** this factory repo (product repos only)
- Kubernetes / multi-service mesh for step-2 PROs
- Replacing host agents (Cursor / Claude) with a mandatory self-hosted LLM
- Full Stripe / domestic payment MoR parity before Lemon Squeezy + DB wiring is solid

---

## How to use this doc

- **Humans:** pick the next ID from Priority 1, open an Issue/PR, shrink the row when done.
- **Agents:** do **not** treat this file as a workflow step. Active work still follows `docs/workflow.md` + `skills/*`.
- When an item ships: update this table (strike or move to “Done” below) and note it in `CHANGELOG.md`.

## Done (recently closed gaps)

| Item | Where |
|------|--------|
| MVP acceptance skill + prompt | `skills/mvp-acceptance.md`, `prompts/05-accept-mvp.md` |
| `persistence-sqlx` pattern | `templates/patterns/persistence-sqlx/` |
| Split web + API publish | `skills/publish-split-web-api.md`, `templates/layouts/web-api/` |
| Publish target skills | `skills/publish-*.md` |
| Compose / pattern CI | `.github/workflows/ci.yml` |
| OAuth / cron / payments / … patterns | `templates/patterns/` |

---

## Related

- [Getting started](getting-started.md)
- [Examples](examples/)
- [Template catalog](../templates/CATALOG.md)
- [Skills catalog](../skills/CATALOG.md)
- [Changelog](../CHANGELOG.md)
- [Contributing](../CONTRIBUTING.md)
