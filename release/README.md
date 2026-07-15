# release/

Deploy primitives for **step 6**. Agents MUST follow `skills/deploy.md` and MUST wait for step-5 human approval.

## Layout

```
release/
├── nginx/          # reverse-proxy snippets
├── cloudflare/     # DNS / SSL / subdomain registry
└── deploy/         # push + route scripts
```

## Port pool

| Host port | Example |
|-----------|---------|
| 8080 | first MVP / test |
| 8081–8090 | subsequent MVPs |

Container listens on 8080; only host mapping increments.

## Agent deploy sequence

1. Confirm step-5 approval.
2. Register subdomain + port in cloudflare registry example / live registry.
3. Run `deploy/push-and-route.sh` with `MVP_NAME`, `MVP_PORT`, `DOMAIN`, `DEPLOY_HOST`, `DEPLOY_PATH`.
4. Install nginx server block from `nginx/snippets/mvp-server.conf.example`.
5. Ensure Cloudflare DNS Proxied; verify `GET /health` on public URL.

## Prerequisites

- Server: Docker, Compose, Nginx
- Domain NS → Cloudflare
- SSH access for `DEPLOY_HOST`
