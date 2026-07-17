# release/

**English** · [简体中文](README.zh-CN.md)

Deploy primitives for **step 6**. Agents MUST follow `skills/deploy.md` and MUST wait for step-5 human approval.

## Layout

```
release/
├── nginx/          # Docker Nginx gateway (shared network maker-flow)
├── cloudflare/     # DNS / SSL / subdomain registry
└── deploy/         # push + route scripts
```

## Ports

| Port | Role |
|------|------|
| **80** (host) | Gateway only — public entry via Cloudflare |
| `8080` / `80` (container) | MVP listen port inside its container (`CONTAINER_PORT`) |
| `8080–8090` (host) | **Optional** local debug mapping only; not required for production |

Production traffic: Cloudflare → gateway `:80` → Docker network alias `MVP_NAME:CONTAINER_PORT`.

## Agent deploy sequence

1. Confirm step-5 approval.
2. Register subdomain in cloudflare registry example / live registry.
3. From the product repo root, run `maker-flow deploy` (or `deploy/push-and-route.sh`) with `DOMAIN`, `DEPLOY_HOST`, optional `CONTAINER_PORT` / `MVP_SERVICE`.
4. Script syncs the gateway, attaches the MVP to `maker-flow`, writes `conf.d/<MVP_NAME>.conf`, runs `nginx -t`, then reloads.
5. Ensure Cloudflare DNS Proxied; verify `GET /health` on public URL.

## Prerequisites

- Server: Docker + Compose (deploy user can run docker; **no host Nginx required**)
- Domain NS → Cloudflare
- SSH access for `DEPLOY_HOST`
