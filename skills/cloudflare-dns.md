# Cloudflare DNS skill

**English** · [简体中文](cloudflare-dns.zh-CN.md)

**Skill id:** `cloudflare-dns`  
**Typical step:** 6 (publish) — whenever DNS records must be listed, created, updated, or deleted on Cloudflare  
**Also:** standalone DNS / DDNS tasks the human requests outside a full publish

## Goal

Configure and manage DNS records on a Cloudflare zone via API + `release/cloudflare/dns.sh` — **without** asking the human to use the dashboard for routine CRUD.

## When to load

Load this skill **in addition to** `skills/deploy.md` when any of:

- Publishing to **Cloudflare Pages** with a **custom hostname**
- Publishing to **VPS gateway** and pointing a subdomain at the server (A / AAAA)
- Human asks to **list / add / change / remove** DNS records on Cloudflare
- DDNS-style updates (IPv4 / IPv6) on Cloudflare

**MUST NOT** load for DNS-only tasks on non-Cloudflare providers (use provider docs).

## Human configuration (ask in chat)

Before running `dns.sh`, confirm:

| Item | Required | Notes |
|------|----------|-------|
| Account | First time | [https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up) |
| `CLOUDFLARE_API_TOKEN` | Yes | **Zone → DNS → Edit** on target zone(s) |
| `CLOUDFLARE_ZONE_ID` | Optional | If omitted, `dns.sh` can discover zones from token and let human choose |
| `CLOUDFLARE_ACCOUNT_ID` | Pages only | Not needed for pure DNS CRUD |
| Zone on Cloudflare | Yes | NS already at Cloudflare (registrar change is human one-time) |

**MUST** prefer env vars the human exports locally. **MUST NOT** invent tokens or paste them into chat logs.

Token creation: [API Tokens](https://dash.cloudflare.com/profile/api-tokens) → Custom token → Zone DNS Edit.

## Agent tool

```bash
DNS="$(maker-flow root)/release/cloudflare/dns.sh"
export CLOUDFLARE_API_TOKEN=…   # from human env
# Optional: export CLOUDFLARE_ZONE_ID / CLOUDFLARE_ACCOUNT_ID
# If omitted, dns.sh can list/discover and prompt selection.
```

| Action | Command |
|--------|---------|
| **List** | `"$DNS" list` · `"$DNS" list --type A --name host.example.com` · `"$DNS" list --json` |
| **List accounts/zones** | `"$DNS" accounts` · `"$DNS" zones` |
| **Get** | `"$DNS" get --id RECORD_ID` |
| **Create** | `"$DNS" create --type TYPE --name NAME --content CONTENT [--proxied true\|false]` |
| **Update** | `"$DNS" update --type TYPE --name NAME --content CONTENT` or `--id RECORD_ID --content …` |
| **Delete** | `"$DNS" delete --type TYPE --name NAME` or `--id RECORD_ID` |
| **Upsert** | `"$DNS" upsert --type TYPE --name NAME --content CONTENT` (publish default) |

Requires `curl` + `python3`. Back-compat: `dns-upsert.sh` → `dns.sh upsert`.

Full reference: [`release/cloudflare/dns-api.md`](../release/cloudflare/dns-api.md).

## Record types (common)

| Type | Use | `proxied` |
|------|-----|-----------|
| `A` | IPv4 → VPS | usually `true` |
| `AAAA` | IPv6 → VPS (DDNS) | usually `true` |
| `CNAME` | Pages custom host → `<project>.pages.dev` | usually `true` |

`ttl: 1` (auto) is set by the script when proxied.

## Publish integration

| Publish target | DNS skill action |
|----------------|------------------|
| `cloudflare-pages` + custom domain | Pages Domains API (see [`release/publish/cloudflare-pages.md`](../release/publish/cloudflare-pages.md)) **then** `dns.sh upsert` CNAME |
| `vps-gateway` | `dns.sh upsert` A and/or AAAA to server IP(s) |
| `*.pages.dev` only | **Skip** DNS skill (no zone record needed) |

Order for custom Pages host: attach domain on project → upsert DNS → verify HTTPS.

## Workflow

1. Ask human what DNS change is needed (or infer from publish target + domain).
2. Confirm token is available. Zone/account IDs are optional.
3. **List** first when unsure (`dns.sh list --name …`) to avoid duplicates.
4. Run **create** / **update** / **upsert** / **delete** as appropriate.
5. **Verify:** `dns.sh list --name …`, `dig +short`, `curl -sI https://…`.

## Hard rules

- **MUST** use `dns.sh` for Cloudflare CRUD when token exists; let script discover zone/account if IDs are missing.
- **MUST** list before delete when multiple records might match (delete by `--id` if ambiguous).
- **MUST NOT** tell humans to run `dns.sh` unless they explicitly want CLI themselves — agent runs it.
- **MUST NOT** change DNS before step-5 MVP approval when DNS is part of publish step 6.
- **MUST** register subdomain in [`release/cloudflare/subdomain-registry.example.md`](../release/cloudflare/subdomain-registry.example.md) when the human keeps a registry.

## Errors

- `list dns_records failed` → check token permissions and selected/derived zone.
- `no record for type=… name=…` on update/delete → list zone or use create/upsert.
- Wrong zone → human must supply the zone id for the correct domain.

## Related

- Publish: [`deploy.md`](deploy.md) · [`prompts/06-publish.md`](../prompts/06-publish.md)
- Cloudflare overview: [`release/cloudflare/README.md`](../release/cloudflare/README.md)
