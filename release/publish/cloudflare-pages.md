# Publish target: cloudflare-pages

**English** · Agent-only. Human confirmation required first.

## When

Static or SPA output (typically assembled `web-vite`). No long-running API in this target alone.

## Human inputs (minimize clicks)

**Sign up:** [https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)

Prefer **token-only** after the Cloudflare account exists:

| Human provides (once) | Agent uses for |
|-----------------------|----------------|
| `CLOUDFLARE_API_TOKEN` | Pages deploy + optional DNS |
| `CLOUDFLARE_ACCOUNT_ID` | Pages project / deploy |
| `CLOUDFLARE_ZONE_ID` | Custom domain DNS upsert (optional) |
| Project name + whether to use `*.pages.dev` or a custom host | Naming |

Token permissions: **Account → Cloudflare Pages → Edit**; if custom DNS: also **Zone → DNS → Edit**. See [`../cloudflare/dns-api.md`](../cloudflare/dns-api.md).

Alternative: `npx wrangler login` on this machine (interactive; avoid when token is available).

## Prerequisites

- Human approved step 5 and chose Cloudflare Pages
- Auth confirmed in chat (token env **or** wrangler login)
- Project name / production branch agreed
- Custom hostname (optional) agreed; zone must already sit on Cloudflare NS

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

## Publish (Direct Upload API via Wrangler)

```bash
export CLOUDFLARE_API_TOKEN=…          # from human env
export CLOUDFLARE_ACCOUNT_ID=…         # required for non-interactive

npx wrangler pages project create <PROJECT> --production-branch main 2>/dev/null || true
npx wrangler pages deploy dist --project-name <PROJECT>
```

Wrangler calls Cloudflare’s Pages Direct Upload API under the hood — humans do not need the dashboard for upload.

Note the returned `https://<project>.pages.dev` URL and give it to the human.

## Custom domain (API)

When the human wants `https://www.example.com` (not only `pages.dev`):

1. **Attach hostname to the Pages project** (Pages Domains API — prefer this over the dashboard):
   ```bash
   export CLOUDFLARE_API_TOKEN=…
   export CLOUDFLARE_ACCOUNT_ID=…
   curl -sS -X POST \
     -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
     -H "Content-Type: application/json" \
     "https://api.cloudflare.com/client/v4/accounts/${CLOUDFLARE_ACCOUNT_ID}/pages/projects/<PROJECT>/domains" \
     --data '{"name":"<HOST>"}'
   ```
2. **Upsert DNS** so the name resolves to the Pages project — typically **CNAME → `<project>.pages.dev`**, proxied:
   ```bash
   export CLOUDFLARE_ZONE_ID=…
   "$(maker-flow root)/release/cloudflare/dns.sh" upsert \
     --type CNAME \
     --name <HOST> \
     --content <PROJECT>.pages.dev \
     --proxied true
   ```
   Order: attach domain on the project **then** (or around the same time) write DNS; if Cloudflare returns a different target than `<PROJECT>.pages.dev`, use the target from the API/dashboard response.
3. Full DNS SOP: [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md) and [`../cloudflare/dns-api.md`](../cloudflare/dns-api.md).

Do **not** use the VPS gateway `conf.d` path for Pages — that is only for `vps-gateway`.

## Verify

```bash
curl -sI "https://<project>.pages.dev/"
# if custom domain:
curl -sI "https://<HOST>/"
```

Open `/` in a browser. `/health` exists only if the static build includes it (Vite template health is Nginx-only — on Pages, rely on `/`).

## Rollback

Redeploy a previous Pages deployment in the Cloudflare dashboard, or `wrangler pages deployment list` + redeploy an older artifact if the human requests it.
