# release/publish/

**English** · [简体中文](README.zh-CN.md)

**Audience: agents** (step 6). Humans choose the target in chat; see `skills/deploy.md`.

## Principle

- **What** to run → PRO / templates  
- **Where** to publish → human confirmation in dialogue  
- **How** → this directory (+ `release/deploy/` for VPS)

Do **not** instruct humans to run `maker-flow deploy`. That CLI is agent-internal for the VPS path only.

## Targets

| id | File |
|----|------|
| `vps-gateway` | [vps-gateway.md](vps-gateway.md) |
| `cloudflare-pages` | [cloudflare-pages.md](cloudflare-pages.md) |
| `github-pages` | [github-pages.md](github-pages.md) |
| `vercel` | [vercel.md](vercel.md) |

## Related

- Skill: [`skills/deploy.md`](../../skills/deploy.md)
- Prompt: [`prompts/06-publish.md`](../../prompts/06-publish.md)
- VPS scripts: [`../deploy/`](../deploy/)
- DNS helpers: [`../cloudflare/`](../cloudflare/)
