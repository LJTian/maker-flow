# Deploy skill

**English** · [简体中文](deploy.zh-CN.md)

**Step:** 6 — production deploy  
**Prerequisite:** step 5 local MVP acceptance passed

## Goal

Expose the container(s) in the product repo (or `workspace/<project-name>/`) to the public internet within about 10 minutes.

## Checklist

- [ ] `docker compose up` local health returns 200
- [ ] Subdomain + `MVP_NAME` registered in `release/cloudflare/subdomain-registry.example.md`
- [ ] Server Docker ready; gateway will use shared network `maker-flow`
- [ ] Cloudflare DNS orange-cloud proxy enabled

## Deploy actions

From the MVP / product project directory:

```bash
maker-flow deploy \
  --domain idea1.your-domain.com \
  --host deploy@your-server \
  --service api \
  --port 8080
```

(`--name` defaults to `PRODUCT_NAME` in `AGENTS.md`. Equivalent env: `DOMAIN`, `DEPLOY_HOST`, `MVP_SERVICE`, `CONTAINER_PORT`.)

Low-level (same effect):

```bash
export MVP_NAME=idea1
export DOMAIN=idea1.your-domain.com
export DEPLOY_HOST=deploy@your-server
export DEPLOY_PATH=/opt/mvps/idea1
export CONTAINER_PORT=8080
export MVP_SERVICE=api

maker-flow root   # shows factory path
"$(maker-flow root)/release/deploy/push-and-route.sh"
```

## Nginx gateway

`push-and-route.sh` installs `conf.d/<MVP_NAME>.conf` on the Docker gateway and reloads after `nginx -t`.  
Manual path: render `release/nginx/snippets/mvp-server.conf.example` into the gateway `conf.d/`, then inside the gateway container run `nginx -t` and `nginx -s reload` (no sudo / no host Nginx).

## Cloudflare

Add an A record to the server IP (Proxied). Details: `release/cloudflare/README.md`.

## Verify

```bash
curl -I https://idea1.your-domain.com/health
```

## Rollback

```bash
ssh $DEPLOY_HOST "cd $DEPLOY_PATH && docker compose down"
ssh $DEPLOY_HOST "rm -f /opt/maker-flow/gateway/conf.d/${MVP_NAME}.conf && cd /opt/maker-flow/gateway && docker compose exec -T nginx nginx -t && docker compose exec -T nginx nginx -s reload"
```

Also free the subdomain in the DNS registry.

## Further reading

- [release/README.md](../release/README.md)
- [release/nginx/README.md](../release/nginx/README.md)
- [release/deploy/](../release/deploy/)
