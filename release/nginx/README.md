# Nginx gateway (step 6)

**English** · [简体中文](README.zh-CN.md)

Docker reverse proxy on shared network `maker-flow`. SOP: `skills/deploy.md`.

## Layout

- `docker-compose.yml` — `nginx:1.27-alpine`, host `80:80`
- `nginx.conf` — includes `/etc/nginx/conf.d/*.conf`
- `conf.d/` — per-MVP server blocks (written by deploy script)
- `snippets/mvp-server.conf.example` — template (`__DOMAIN__`, `__MVP_NAME__`, `__CONTAINER_PORT__`)

## One-time on server

Needs Docker only (no apt/system Nginx, no sudo for this flow):

```bash
docker network create maker-flow 2>/dev/null || true
# Gateway files are synced by push-and-route.sh to GATEWAY_PATH (default /opt/maker-flow/gateway)
```

## Per-MVP

Prefer `release/deploy/push-and-route.sh` from the product repo root. It will:

1. Sync gateway + MVP
2. `docker compose up` the MVP
3. `docker network connect --alias <MVP_NAME> maker-flow <container>`
4. Install `conf.d/<MVP_NAME>.conf`
5. `nginx -t` then reload

## Local smoke

```bash
docker network create maker-flow 2>/dev/null || true
cd release/nginx
docker compose up -d
docker compose logs -f nginx
```

## Debug

```bash
docker compose -f /opt/maker-flow/gateway/docker-compose.yml exec nginx nginx -t
docker compose -f /opt/maker-flow/gateway/docker-compose.yml logs -f nginx
```
