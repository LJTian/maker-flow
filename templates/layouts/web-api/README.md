# Layout: web-api

**English** · [简体中文](README.zh-CN.md)

**Layout id:** `web-api`  
**Apps:** `go-api` + `web-vite`  
**When:** PRO needs a browser UI **and** a REST API (local compose + optional split publish).

This directory is a **product-root skeleton**, not a deployable app by itself. Copy its files into the product repo **after** copying `apps/go-api` → `api/` and `apps/web-vite` → `web/`.

## Product layout after assembly

```text
<product-root>/
├── pro.md
├── docker-compose.yml      # from this layout
├── .env.example            # from this layout
├── .env                    # local secrets (not committed)
├── api/                    # templates/apps/go-api
└── web/                    # templates/apps/web-vite
```

## Agent steps (step 4)

1. Copy `templates/apps/go-api/` → `<product>/api/`
2. Copy `templates/apps/web-vite/` → `<product>/web/`
3. Copy this layout’s `docker-compose.yml` + `.env.example` → `<product>/`
4. Rewrite Go module path under `api/`
5. Implement PRO logic; keep process boundaries (API ≠ SPA)
6. Local: `cp .env.example .env && docker compose up --build`

## Env contract

| Variable | Used by | Notes |
|----------|---------|-------|
| `VITE_API_BASE_URL` | `web` **build arg** | Browser-reachable API origin. Local: `http://localhost:8080`. Split publish: `https://api.example.com` |
| `CORS_ORIGINS` | `api` | Comma-separated. Local default `*`. Split: set to Pages origin(s), e.g. `https://app.example.com` |
| `API_HOST_PORT` | laptop → api | Default `8080` |
| `WEB_HOST_PORT` | laptop → web | Default `3000` |

`VITE_*` is baked at **image build** time. Changing the public API URL requires rebuild of `web`.

## Local acceptance (step 5)

```bash
cp -n .env.example .env
docker compose up --build -d
curl -sf http://localhost:8080/health
curl -sf http://localhost:3000/health
# then walk pro.md criteria — see skills/mvp-acceptance.md
```

## Split publish (step 6)

When the human wants **SPA on Pages/Vercel/GitHub Pages** and **API on VPS**, follow:

[`skills/publish-split-web-api.md`](../../skills/publish-split-web-api.md)

Do **not** ship the API to Cloudflare Pages alone.
