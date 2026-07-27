# Cloudflare DNS API (agent)

**English** · [简体中文](dns-api.zh-CN.md)

Upsert DNS records via the Cloudflare API so humans need not click the dashboard after the zone is on Cloudflare. Same capability as DDNS (A/AAAA); also used for Pages custom domains (CNAME).

## Human one-time setup

1. Domain NS already pointed at Cloudflare (API cannot change the registrar).
2. Create an API token: [API Tokens](https://dash.cloudflare.com/profile/api-tokens) → Create Token.
3. Recommended permissions (can be one token):
   - **Account** → **Cloudflare Pages** → **Edit** (static publish)
   - **Zone** → **DNS** → **Edit** (this guide)
   - Zone Resources: include the target zone (or all zones in the account)
4. Give the agent (env or chat, prefer env):
   - `CLOUDFLARE_API_TOKEN`
   - `CLOUDFLARE_ZONE_ID` (zone Overview → Zone ID)
   - For Pages also: `CLOUDFLARE_ACCOUNT_ID`

**MUST NOT** invent or log tokens. Ask the human to export them.

## When to use

| Record | Typical content | Proxied |
|--------|-----------------|--------|
| `A` | VPS IPv4 | usually `true` |
| `AAAA` | VPS IPv6 (same idea as DDNS) | usually `true` |
| `CNAME` | `<project>.pages.dev` for Pages custom host | usually `true` |

Prefer API over dashboard whenever token + zone id are available (`skills/deploy.md`).

## Helper script

From any cwd (factory must be installed / checkout available):

```bash
export CLOUDFLARE_API_TOKEN=…
export CLOUDFLARE_ZONE_ID=…

"$(maker-flow root)/release/cloudflare/dns-upsert.sh" \
  --type CNAME \
  --name www.example.com \
  --content my-project.pages.dev \
  --proxied true
```

Equivalent env-style:

```bash
CF_DNS_TYPE=A CF_DNS_NAME=api.example.com CF_DNS_CONTENT=203.0.113.10 CF_DNS_PROXIED=true \
  "$(maker-flow root)/release/cloudflare/dns-upsert.sh"
```

Behavior: list records for `name` + `type` → **PATCH** if one exists, else **POST**. Idempotent enough for re-runs and DDNS-like updates.

## Raw API (reference)

Base: `https://api.cloudflare.com/client/v4`

```bash
AUTH="Authorization: Bearer ${CLOUDFLARE_API_TOKEN}"
ZONE="${CLOUDFLARE_ZONE_ID}"

# List
curl -sS -H "$AUTH" \
  "https://api.cloudflare.com/client/v4/zones/${ZONE}/dns_records?type=CNAME&name=www.example.com"

# Create
curl -sS -X POST -H "$AUTH" -H "Content-Type: application/json" \
  "https://api.cloudflare.com/client/v4/zones/${ZONE}/dns_records" \
  --data '{"type":"CNAME","name":"www.example.com","content":"my-project.pages.dev","proxied":true,"ttl":1}'

# Update (replace RECORD_ID)
curl -sS -X PATCH -H "$AUTH" -H "Content-Type: application/json" \
  "https://api.cloudflare.com/client/v4/zones/${ZONE}/dns_records/RECORD_ID" \
  --data '{"type":"CNAME","name":"www.example.com","content":"my-project.pages.dev","proxied":true,"ttl":1}'
```

`ttl: 1` means “automatic” when proxied.

Docs: [DNS records API](https://developers.cloudflare.com/api/resources/dns/subresources/records/).

## Pages custom domain flow

1. Publish with [`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md) → note `*.pages.dev` URL / project name.
2. Attach custom hostname via Pages Domains API:
   ```bash
   curl -sS -X POST \
     -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
     -H "Content-Type: application/json" \
     "https://api.cloudflare.com/client/v4/accounts/${CLOUDFLARE_ACCOUNT_ID}/pages/projects/<PROJECT>/domains" \
     --data '{"name":"<HOST>"}'
   ```
3. Upsert **CNAME** (or the target Cloudflare returns) with `dns-upsert.sh`.
4. Verify `https://<custom-host>/`.

SSL for proxied records is handled by Cloudflare; keep zone SSL/TLS at **Full** or **Full (strict)** when origin is HTTPS (VPS). For Pages, orange-cloud CNAME is enough for HTTPS on the custom host.

## VPS gateway flow

After `release/publish/vps-gateway.md`, upsert **A** and/or **AAAA** to the server public address(es), `proxied=true`. Same script.

## Verify

```bash
dig +short <name> A
dig +short <name> AAAA
dig +short <name> CNAME
curl -sI "https://<name>/"
```

## Related

- Pages publish: [`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md)
- Registry notes: [`subdomain-registry.example.md`](subdomain-registry.example.md)
