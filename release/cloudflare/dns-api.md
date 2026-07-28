# Cloudflare DNS API (script reference)

**English** · [简体中文](dns-api.zh-CN.md)

CLI and raw API reference for `release/cloudflare/dns.sh`. **Agent SOP:** [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md).

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

- **Pages custom domain:** [`skills/publish-cloudflare-pages.md`](../../skills/publish-cloudflare-pages.md) — attach hostname on project, then `dns.sh upsert` CNAME.
- **VPS gateway:** [`skills/publish-vps-gateway.md`](../../skills/publish-vps-gateway.md) — `dns.sh upsert` A / AAAA.

## Verify

```bash
"$DNS" list --name api.example.com
dig +short api.example.com A
curl -sI "https://api.example.com/"
```

## Related

- [`README.md`](README.md)
- [`subdomain-registry.example.md`](subdomain-registry.example.md)
