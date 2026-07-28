# Cloudflare DNS API (agent)

**English** · [简体中文](dns-api.zh-CN.md)

CRUD DNS records via the Cloudflare API. Same idea as DDNS (A/AAAA); also used for Pages custom domains (CNAME).

**Agent skill:** [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md) — load when listing or changing DNS on Cloudflare.

## Human one-time setup

1. Sign up: [https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)
2. Domain NS already pointed at Cloudflare (API cannot change the registrar).
3. Create an API token: [API Tokens](https://dash.cloudflare.com/profile/api-tokens) → Create Token.
4. Permissions: **Zone → DNS → Edit** (include target zone). For Pages publish add **Account → Cloudflare Pages → Edit**.
5. Export for the agent:
   - `CLOUDFLARE_API_TOKEN` (**required**)
   - `CLOUDFLARE_ZONE_ID` (**optional**; script can discover/select zones from token)
   - `CLOUDFLARE_ACCOUNT_ID` (**optional**; helps scope discovery, needed by some Pages API calls)

**MUST NOT** invent or log tokens.

## CLI: `dns.sh`

Path: `$(maker-flow root)/release/cloudflare/dns.sh`

```bash
export CLOUDFLARE_API_TOKEN=…
DNS="$(maker-flow root)/release/cloudflare/dns.sh"

# Optional: preselect account/zone
# export CLOUDFLARE_ACCOUNT_ID=…
# export CLOUDFLARE_ZONE_ID=…

# List (read)
"$DNS" list
"$DNS" list --type A --name api.example.com
"$DNS" list --json

# Discover from token
"$DNS" accounts
"$DNS" zones
"$DNS" --zone-name example.com list

# Get one record by id
"$DNS" get --id RECORD_ID

# Create
"$DNS" create --type A --name api.example.com --content 203.0.113.10 --proxied true

# Update (by id, or by type + name)
"$DNS" update --type A --name api.example.com --content 203.0.113.11
"$DNS" update --id RECORD_ID --content 203.0.113.12

# Delete
"$DNS" delete --type A --name api.example.com
"$DNS" delete --id RECORD_ID

# Upsert (create or update same type + name) — used by publish flows
"$DNS" upsert --type CNAME --name www.example.com --content my-project.pages.dev
```

`dns-upsert.sh` is a thin wrapper: `dns.sh upsert …` (backward compatible).

Requires: `curl`, `python3`.

## Raw API (reference)

Base: `https://api.cloudflare.com/client/v4/zones/{zone_id}/dns_records`

| Action | HTTP |
|--------|------|
| List | `GET …/dns_records` |
| Get | `GET …/dns_records/{id}` |
| Create | `POST …/dns_records` |
| Update | `PATCH …/dns_records/{id}` |
| Delete | `DELETE …/dns_records/{id}` |

Docs: [DNS records API](https://developers.cloudflare.com/api/resources/dns/subresources/records/).

When token can access multiple accounts/zones, `dns.sh` supports:
- interactive selection (TTY)
- or explicit `--account-id` / `--zone-id` / `--zone-name`
- or `--non-interactive` to fail fast in automation

## Publish flows

- **Pages custom domain:** [`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md) — attach hostname on project, then `dns.sh upsert` CNAME.
- **VPS gateway:** [`../publish/vps-gateway.md`](../publish/vps-gateway.md) — `dns.sh upsert` A / AAAA.

## Verify

```bash
"$DNS" list --name api.example.com
dig +short api.example.com A
curl -sI "https://api.example.com/"
```

## Related

- [`README.md`](README.md)
- [`subdomain-registry.example.md`](subdomain-registry.example.md)
