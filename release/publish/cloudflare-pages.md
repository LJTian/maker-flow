# Publish target: cloudflare-pages

**English** · Agent-only. Human confirmation required first.

## When

Static or SPA output (typically assembled `web-vite`). No long-running API in this target alone.

## Prerequisites

- Human approved step 5 and chose Cloudflare Pages
- Human has a Cloudflare account; confirm auth method:
  - `npx wrangler login` already done on this machine, **or**
  - `CLOUDFLARE_API_TOKEN` (+ account id if required) in the environment
- Project name / production branch agreed in chat

## Build

Prefer container build so the host need not own a global Node toolchain. From the web app directory (product root or `web/`):

```bash
docker compose run --rm --no-deps \
  --entrypoint sh web \
  -c "npm ci && npm run build" 2>/dev/null \
|| docker run --rm -v "$PWD:/app" -w /app node:22-alpine \
  sh -c "npm ci && npm run build"
```

Expect `dist/` (Vite default). Adjust if the PRO changed `outDir`.

## Publish

```bash
npx wrangler pages project create <PROJECT> --production-branch main 2>/dev/null || true
npx wrangler pages deploy dist --project-name <PROJECT>
```

Custom domain: follow Cloudflare dashboard / `wrangler pages` domain docs after human confirms the hostname. DNS helpers for **VPS** A records are under `release/cloudflare/` — Pages custom domains are configured on the Pages project, not the Docker gateway.

## Verify

Open the `*.pages.dev` URL (or custom domain) returned by wrangler. Check `/` loads; `/health` exists only if the static build includes it (Vite template health is Nginx-only — on Pages, rely on `/`).

## Rollback

Redeploy a previous Pages deployment in the Cloudflare dashboard, or `wrangler pages deployment list` + redeploy an older artifact if the human requests it.
