# Cloudflare (step 6)

**English** · [简体中文](README.zh-CN.md)

Agent checklist for DNS/SSL. Full deploy SOP: `skills/deploy.md`.

## Required state

- Zone active on Cloudflare (NS pointed)
- SSL/TLS mode: `Full` or `Full (strict)`
- Subdomain A/CNAME → server IP, **Proxied**

## Per-MVP DNS

| Type | Name | Content | Proxy |
|------|------|---------|-------|
| A | `ideaN` | server public IP | Proxied |

Register name + `HOST_PORT` in subdomain registry before assigning ports.

## Optional automation

Env: `CLOUDFLARE_API_TOKEN`, `CLOUDFLARE_ZONE_ID` (Zone DNS Edit). Prefer API over dashboard when available.

## Verify

```bash
curl -sfI "https://ideaN.your-domain.com/health"
```
