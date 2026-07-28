# Cloudflare (step 6)

**English** · [简体中文](README.zh-CN.md)

DNS / SSL / token automation for publish. Skill: `skills/cloudflare-dns.md` · scripts in this directory.

## Guides

| Doc / script | Purpose |
|--------------|---------|
| [`dns-api.md`](dns-api.md) | DNS CRUD reference (script flags) |
| [`dns.sh`](dns.sh) | Agent CLI — skill: [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md) |
| [`dns-upsert.sh`](dns-upsert.sh) | Back-compat wrapper → `dns.sh upsert` |
| [`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md) | Pages Direct Upload + optional custom domain DNS |
| [`subdomain-registry.example.md`](subdomain-registry.example.md) | Human registry of MVP subdomains |

## Human token checklist

**No account yet?** Open [https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)

| Variable | Needed for |
|----------|------------|
| `CLOUDFLARE_API_TOKEN` | Pages + DNS |
| `CLOUDFLARE_ACCOUNT_ID` | Pages (non-interactive) |
| `CLOUDFLARE_ZONE_ID` | DNS upsert |

Permissions: **Pages Edit** and/or **DNS Edit** on the zone. Prefer API over dashboard when set.

## Required state (custom domains)

- Zone active on Cloudflare (**NS** pointed at CF — one-time at registrar; not via API)
- SSL/TLS mode: `Full` or `Full (strict)` when proxying to an HTTPS origin (VPS)
- Record matches the publish target:

| Target | Record | Content | Proxy |
|--------|--------|---------|-------|
| VPS gateway | A / AAAA | server public IP(s) | Proxied |
| Cloudflare Pages | CNAME | `<project>.pages.dev` | Proxied |

## Per-MVP DNS

Register name + `MVP_NAME` in the subdomain registry before publish when the human keeps a list.

## Verify

```bash
curl -sfI "https://ideaN.your-domain.com/"
# VPS health example:
curl -sfI "https://ideaN.your-domain.com/health"
```
