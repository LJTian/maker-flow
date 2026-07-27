# Publish skill (step 6)

**English** · [简体中文](deploy.zh-CN.md)

**Step:** 6 — go live  
**Prerequisite:** step 5 local MVP acceptance passed  
**Skill id:** `deploy` (filename kept for catalog stability)

## Goal

Expose the approved MVP on the public internet. **Shape** (static vs runtime) comes from the PRO; **where** it ships is chosen by the **human in chat** — not by a human-facing CLI.

## Hard rules

- **MUST** stop and **ask the human** which publish target(s) to use before running any publish action.
- **MUST NOT** tell the human to run `maker-flow deploy` (that CLI is **agent-internal** only).
- **MUST NOT** publish until step 5 is approved.
- **MUST** refuse impossible pairs (e.g. Postgres-backed API → Cloudflare Pages alone). Propose a split (static frontend + VPS API) instead.
- Static vs non-static: decide from PRO / assembled apps; do not default to VPS or Pages.

## Conversation gate (required)

Before executing, confirm in chat:

1. **What** ships: whole product / frontend only / API only / worker (no public URL)?
2. **Where** (one or more targets):
   - `vps-gateway` — Docker on a VPS + shared Nginx gateway
   - `cloudflare-pages` — Cloudflare Pages (token + optional DNS API for custom hosts)
   - `github-pages` — GitHub Pages
   - `vercel` — Vercel
3. **Domain:** platform default URL vs custom hostname
4. **Credentials:** human confirms platform login / tokens are available (do not invent secrets). For Cloudflare prefer `CLOUDFLARE_API_TOKEN` (+ `CLOUDFLARE_ACCOUNT_ID`; + `CLOUDFLARE_ZONE_ID` if setting DNS).

Only after the human answers, follow the matching guide under `release/publish/`.

## Ports (VPS path only)

Three different numbers — do not conflate:

| Layer | Meaning | web-vite | go-api |
|-------|---------|----------|--------|
| Local `HOST_PORT` | Browser on the laptop | `3000` → container | `8080` → container |
| `CONTAINER_PORT` | Listen port **inside** the container | **80** | **8080** |
| Public entry | Cloudflare → gateway host `:80` | always gateway 80 | same |

Agent-internal VPS publish uses **`CONTAINER_PORT`**, not `HOST_PORT`.

| App template | Compose service | `CONTAINER_PORT` |
|--------------|-----------------|------------------|
| `go-api` | `api` | `8080` |
| `web-vite` | `web` | `80` |
| `go-worker` | `worker` | usually **no** public route |

## Target matrix

| Target | Good for | Not for | Agent guide |
|--------|----------|---------|-------------|
| `vps-gateway` | APIs, workers, full Docker compose, self-hosted static | Users without a VPS | [`release/publish/vps-gateway.md`](../release/publish/vps-gateway.md) |
| `cloudflare-pages` | Static / SPA (`web-vite` build) | DB, long-running Go API | [`release/publish/cloudflare-pages.md`](../release/publish/cloudflare-pages.md) (+ [`release/cloudflare/dns-api.md`](../release/cloudflare/dns-api.md) for custom host) |
| `github-pages` | Static / SPA | Same as above | [`release/publish/github-pages.md`](../release/publish/github-pages.md) |
| `vercel` | Static / SPA (and Vercel-native SSR later) | Self-hosted Postgres on Vercel free tier without redesign | [`release/publish/vercel.md`](../release/publish/vercel.md) |

Mixed products: publish frontend to Pages/Vercel **and** API to `vps-gateway` when the human wants that split.

## After publish

- Give the human the public URL(s).
- Verify (`curl` / open `/` or `/health` as appropriate).
- Record subdomain / project name if the human keeps a registry.

## Rollback

Follow the rollback section in the chosen `release/publish/<target>.md`.

## Further reading

- [`release/publish/README.md`](../release/publish/README.md)
- [`release/README.md`](../release/README.md)
- Cloudflare DNS API: [`release/cloudflare/dns-api.md`](../release/cloudflare/dns-api.md)
- Prompt shape: [`prompts/06-publish.md`](../prompts/06-publish.md)
