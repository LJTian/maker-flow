# Publish target: vps-gateway

**Skill id:** `publish-vps-gateway`  
**Step:** 6 — after human confirms target via `skills/deploy.md`

Human confirmation required first ([`deploy.md`](deploy.md)).

## When

Dockerized API / worker / self-hosted static (`web-vite` Nginx image) on a server the human controls.

## Prerequisites

- Human approved step 5
- Human provided: SSH `user@host`, public hostname, compose **service** name
- Server: Docker + Compose; deploy user can run `docker`
- Optional: Cloudflare zone for DNS (Proxied)

## Ports

Use **container listen port**, not laptop `HOST_PORT`:

| Service | `CONTAINER_PORT` |
|---------|------------------|
| `api` (go-api) | `8080` |
| `web` (web-vite) | `80` |

## Agent-internal execution

From the **product repo root** (do not ask the human to type this):

```bash
# Prefer wrapping script (agent-internal CLI):
maker-flow deploy \
  --domain <DOMAIN> \
  --host <DEPLOY_HOST> \
  --service <MVP_SERVICE> \
  --port <CONTAINER_PORT>

# Equivalent:
export MVP_NAME=<PRODUCT_NAME>
export DOMAIN=<DOMAIN>
export DEPLOY_HOST=<DEPLOY_HOST>
export DEPLOY_PATH=/opt/mvps/<PRODUCT_NAME>
export CONTAINER_PORT=<CONTAINER_PORT>
export MVP_SERVICE=<MVP_SERVICE>
"$(maker-flow root)/release/deploy/push-and-route.sh"
```

`--service` is required. Script syncs gateway + MVP, attaches container to network `maker-flow`, writes `conf.d/<MVP_NAME>.conf`, reloads Nginx.

Then ensure Cloudflare A/CNAME → server IP (Proxied) if using a custom domain. Prefer API upsert:

```bash
export CLOUDFLARE_API_TOKEN=…
export CLOUDFLARE_ZONE_ID=…
"$(maker-flow root)/release/cloudflare/dns.sh" upsert \
  --type A --name <DOMAIN> --content <SERVER_IPV4> --proxied true
# IPv6 (optional, same as DDNS):
# "$(maker-flow root)/release/cloudflare/dns.sh" upsert \
#   --type AAAA --name <DOMAIN> --content <SERVER_IPV6> --proxied true
```

See [`cloudflare-dns.md`](cloudflare-dns.md) and [`release/cloudflare/dns-api.md`](../release/cloudflare/dns-api.md). Dashboard: [`release/cloudflare/README.md`](../release/cloudflare/README.md).

## Verify

```bash
curl -sI "https://<DOMAIN>/" 
curl -sf "https://<DOMAIN>/health" || true
```

## Rollback

```bash
ssh "$DEPLOY_HOST" "cd '$DEPLOY_PATH' && docker compose down"
ssh "$DEPLOY_HOST" "rm -f /opt/maker-flow/gateway/conf.d/${MVP_NAME}.conf && cd /opt/maker-flow/gateway && docker compose exec -T nginx nginx -t && docker compose exec -T nginx nginx -s reload"
```
