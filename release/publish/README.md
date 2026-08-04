# release/publish/

**English** · [简体中文](README.zh-CN.md)

Index for step-6 publish targets. **Agent SOPs live in `skills/`** — not here.

| Target id | Skill |
|-----------|-------|
| `vps-gateway` | [`skills/publish-vps-gateway.md`](../../skills/publish-vps-gateway.md) |
| `cloudflare-pages` | [`skills/publish-cloudflare-pages.md`](../../skills/publish-cloudflare-pages.md) |
| `github-pages` | [`skills/publish-github-pages.md`](../../skills/publish-github-pages.md) |
| `vercel` | [`skills/publish-vercel.md`](../../skills/publish-vercel.md) |
| `split-web-api` | [`skills/publish-split-web-api.md`](../../skills/publish-split-web-api.md) (SPA host + VPS API) |

## Flow

1. Human chooses target(s) in chat — [`skills/deploy.md`](../../skills/deploy.md) + [`prompts/06-publish.md`](../../prompts/06-publish.md)
2. Agent loads the matching `skills/publish-<target>.md`
3. Scripts and infra: [`../deploy/`](../deploy/) (VPS), [`../cloudflare/`](../cloudflare/) (DNS CLI), [`../nginx/`](../nginx/) (gateway)

Do **not** instruct humans to run `maker-flow deploy` (agent-internal for VPS only).

## Related

- Catalog: [`skills/CATALOG.md`](../../skills/CATALOG.md)
- Release overview: [`../README.md`](../README.md)
