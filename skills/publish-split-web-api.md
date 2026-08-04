# Publish recipe: split web + API

**English** · [简体中文](publish-split-web-api.zh-CN.md)

**Skill id:** `publish-split-web-api`  
**Step:** 6 — after human confirms a **split** target via `skills/deploy.md`  
**Depends on:** [`publish-cloudflare-pages.md`](publish-cloudflare-pages.md) / [`publish-github-pages.md`](publish-github-pages.md) / [`publish-vercel.md`](publish-vercel.md) **and** [`publish-vps-gateway.md`](publish-vps-gateway.md)

## When

Product has both:

- a static/SPA frontend (`web-vite`, often under `web/`)
- a long-running API (`go-api`, often under `api/`)

Human wants **frontend on a static host** and **API on a VPS** (not “everything on Pages”, not “everything on one VPS container” unless they say so).

Local multi-app skeleton: [`templates/layouts/web-api/`](../templates/layouts/web-api/).

## Hard rules

- **MUST** publish API via `vps-gateway` (or refuse if the human has no VPS).
- **MUST** bake `VITE_API_BASE_URL` to the **public API origin** before building the SPA for production.
- **MUST** set API `CORS_ORIGINS` to the **public frontend origin(s)** (not `*` in production unless the human explicitly accepts it).
- **MUST** return **two URLs** (web + api) and verify both.
- **MUST NOT** put Postgres / Go API onto Cloudflare Pages / GitHub Pages / Vercel alone.

## Conversation inputs (confirm before shipping)

1. Frontend host: `cloudflare-pages` | `github-pages` | `vercel`
2. Frontend public origin: `https://<project>.pages.dev` and/or custom `https://app.example.com`
3. API public origin: `https://api.example.com` (recommended) — needs VPS + DNS
4. SSH `user@host` for VPS; Cloudflare token/zone if using custom DNS
5. Whether production CORS may stay `*` (default: **no** — set explicit origins)

## Order of operations

### 1. Freeze public origins

Example:

| Role | Origin |
|------|--------|
| Web | `https://todo.example.com` (or `https://todo.pages.dev`) |
| API | `https://api.todo.example.com` |

Write these into chat and the product root `.env` (do not commit secrets).

### 2. Ship API first (VPS)

From **product repo root**, with compose able to build service `api` (root `layouts/web-api` compose or `api/docker-compose.yml`):

1. On the server-facing config set:
   - `CORS_ORIGINS=https://todo.example.com` (comma-separate multiples)
   - `APP_ENV=production` when appropriate
2. Follow [`publish-vps-gateway.md`](publish-vps-gateway.md):
   - `--service api`
   - `--port 8080`
   - `--domain api.todo.example.com` (or the agreed API host)
3. DNS A/AAAA (proxied) for the API host via [`cloudflare-dns.md`](cloudflare-dns.md) when applicable.

Verify:

```bash
curl -sf "https://api.todo.example.com/health"
```

### 3. Rebuild SPA against public API

`VITE_*` is compile-time. From `web/` (or product root compose build for `web`):

```bash
export VITE_API_BASE_URL="https://api.todo.example.com"
# Prefer compose build so Dockerfile ARG is set:
docker compose build --build-arg VITE_API_BASE_URL="$VITE_API_BASE_URL" web
# Or from web/:
# docker compose build --build-arg VITE_API_BASE_URL="$VITE_API_BASE_URL"
```

Ensure the artifact in `web/dist` (or the image used only for extracting `dist`) was built with that URL. For Pages Direct Upload, produce `web/dist` with the public API base:

```bash
cd web
docker run --rm -v "$PWD:/app" -w /app \
  -e VITE_API_BASE_URL="https://api.todo.example.com" \
  node:22-alpine \
  sh -c "npm ci && npm run build"
```

### 4. Ship frontend (static host)

Follow the chosen skill with the freshly built `dist/`:

- [`publish-cloudflare-pages.md`](publish-cloudflare-pages.md)
- [`publish-github-pages.md`](publish-github-pages.md)
- [`publish-vercel.md`](publish-vercel.md)

Attach custom domain + DNS if agreed.

### 5. Dual-URL acceptance (required)

```bash
# Web
curl -sI "https://todo.example.com/" | head -n1
# API
curl -sf "https://api.todo.example.com/health"
# Browser CORS smoke (optional): open web origin, confirm UI reaches API /health
```

Present to the human:

```text
Split publish URLs
  Web: https://todo.example.com
  API: https://api.todo.example.com
Verify: web / → 200 · api /health → ok · CORS_ORIGINS includes web origin
```

## Rollback

- Frontend: redeploy previous static deployment (see the static skill).
- API: follow rollback in [`publish-vps-gateway.md`](publish-vps-gateway.md).

## Related

- Layout (compose / env): [`templates/layouts/web-api/`](../templates/layouts/web-api/)
- Parent gate: [`deploy.md`](deploy.md)
- Local gate: [`mvp-acceptance.md`](mvp-acceptance.md)
