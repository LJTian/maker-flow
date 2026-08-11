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
| Patterns | 14 packages under `templates/patterns/` (incl. `persistence-sqlx`, OAuth login, Lemon Squeezy webhook, cron, …) |
| Multi-app layout | `templates/layouts/web-api/` + `skills/publish-split-web-api.md` |
| CI | `scripts/check.sh`, pattern `go test`, app `docker compose` / `docker build` |
| Live example | [static intro → GitHub Pages](examples/static-intro-github-pages.md) |

---

## Unimplemented inventory (honest)

Things that were over-claimed in catalogs/skills, or are still missing. Agents **MUST NOT** invent these as if they already exist in the factory.

### Patterns / payments / auth

| ID | Missing | Notes |
|----|---------|--------|
| U-1 | **Native WeChat Pay** pattern | WeChat **OAuth login** exists in `auth-oauth-jwt`; **payment** does not |
| U-2 | **Native Alipay** pattern | Not in catalog |
| U-3 | **Native Stripe** pattern | Not in catalog (`payment-lemonsqueezy` is MoR webhook only) |
| U-4 | Auth → DB assembly sample | `auth-oauth-jwt` still has `TODO: Save to DB` |
| U-5 | Payment → DB upgrade sample | `payment-lemonsqueezy` still has `TODO: Upgrade VIP` |
| U-6 | Production-ready **Sign in with Apple** | Current `apple.New(..., nil)` is a thin scaffold (no key/team wiring docs) |
| U-7 | Light auth (API key / signed cookie) | No pattern yet; full OAuth only |
| U-8 | Generic **SMTP** email | `notify-email` is **Resend-only** |
| U-9 | Native **Anthropic** SDK | `ai-llm-client` is OpenAI-compatible only |

### Apps / publish / examples

| ID | Missing | Notes |
|----|---------|--------|
| U-10 | Python / Node API apps | Catalog is **Go + Vite** only |
| U-11 | Vercel SSR / **Next.js** template | `publish-vercel` is static/SPA only |
| U-12 | Extra live examples | Need `go-api`→VPS and `web-api` split walkthroughs (only static Pages demo today) |
| U-13 | Static publish **scripts** | Pages / Vercel / GH Pages lack script parity with VPS `push-and-route.sh` |

### Docs / factory hygiene

| ID | Missing | Notes |
|----|---------|--------|
| U-14 | `CONTRIBUTING.zh-CN.md` | Linked from `CONTRIBUTING.md` but file missing |
| U-15 | `skills/publish-*.zh-CN.md` | Several publish skills lack ZH siblings |
| U-16 | `prompts/pro-multi.example.zh-CN.md` | EN example exists; ZH missing |
| U-17 | First **git tag** / changelog cut | CLI says `0.5.0`; changelog still `[Unreleased]` |
| U-18 | `web-vite` automated tests | CONTRIBUTING requires app tests; Vite has none |
| U-19 | Pattern `go.sum` for all Go patterns | Many rely on CI `go mod tidy` only |
| U-20 | Skill YAML frontmatter vs CONTRIBUTING | CONTRIBUTING claims frontmatter; skills have none — pick one |

---

## Priority 1 — high impact on real incubation

| ID | Item | Plan |
|----|------|------|
| P1-1 | More end-to-end examples | Cover **U-12** |
| P1-2 | Static publish script parity | Cover **U-13** |
| P1-3 | Auth + DB assembly sample | Cover **U-4** |
| P1-4 | Payment + DB assembly sample | Cover **U-5** |
| P1-6 | Native WeChat Pay pattern (optional next) | Cover **U-1** when prioritized |
| P1-7 | Native Alipay / Stripe patterns (optional) | Cover **U-2** / **U-3** when prioritized |

**Suggested order:** P1-1 → P1-2 → P1-3 → P1-4 → then pick P1-6/P1-7 if domestic or Stripe pay is needed.

---

## Priority 2 — factory maturity

| ID | Item | Covers |
|----|------|--------|
| P2-1 | Versioned release | U-17 |
| P2-2 | Chinese doc parity | U-14, U-15, U-16 |
| P2-3 | `web-vite` tests | U-18 |
| P2-4 | Pattern `go.sum` coverage | U-19 |
| P2-5 | CONTRIBUTING ↔ reality (frontmatter) | U-20 |
| P2-6 | Stack boundary clarity | Done in `templates/CATALOG.md` (+ link here); keep updated |

---

## Priority 3 — polish

| ID | Item | Covers |
|----|------|--------|
| P3-1 | Light auth pattern | U-7 |
| P3-2 | `.gitignore` dedupe | — |
| P3-3 | Harden Apple Sign In docs/wiring | U-6 |
| P3-4 | Queue backends (Redis/Asynq) | Explicit non-goal unless a product needs it |
| P3-5 | SMTP / multi-provider email | U-8 |
| P3-6 | Native Anthropic client | U-9 |

---

## Explicit non-goals (for now)

- Assembling MVPs **into** this factory repo (product repos only)
- Kubernetes / multi-service mesh for step-2 PROs
- Replacing host agents (Cursor / Claude) with a mandatory self-hosted LLM
- Pretending Lemon Squeezy covers WeChat Pay / Alipay / Stripe native APIs

---

## How to use this doc

- **Humans:** pick the next ID from Priority 1, open an Issue/PR; when done, move the matching **U-*** row to Done.
- **Agents:** do **not** treat this file as a workflow step. Active work still follows `docs/workflow.md` + `skills/*`. **MUST NOT** select missing payment providers as if they were `payment-lemonsqueezy`.
- When an item ships: update this table and note it in `CHANGELOG.md`.

## Done (recently closed gaps)

| Item | Where |
|------|--------|
| Catalog / skill honesty pass (remove fake `stripe`/`alipay`/`wechat` pay tags) | `templates/CATALOG*`, `patterns/index*`, `skills/template-matching*`, payment/auth READMEs |
| MVP acceptance skill + prompt | `skills/mvp-acceptance.md`, `prompts/05-accept-mvp.md` |
| `persistence-sqlx` pattern | `templates/patterns/persistence-sqlx/` |
| Split web + API publish | `skills/publish-split-web-api.md`, `templates/layouts/web-api/` |
| Publish target skills | `skills/publish-*.md` |
| Compose / pattern CI | `.github/workflows/ci.yml` |

---

## Related

- [Getting started](getting-started.md)
- [Examples](examples/)
- [Template catalog](../templates/CATALOG.md)
- [Skills catalog](../skills/CATALOG.md)
- [Changelog](../CHANGELOG.md)
- [Contributing](../CONTRIBUTING.md)
