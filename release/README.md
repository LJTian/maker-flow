# release/

**English** · [简体中文](README.zh-CN.md)

Step **6** publish primitives. Agents MUST follow `skills/deploy.md`, confirm the target with the **human in chat**, and MUST wait for step-5 approval.

Humans are **not** expected to run `maker-flow deploy`. That CLI is agent-internal for the VPS path.

## Layout

```
release/
├── publish/        # Target index → skills/publish-*.md
├── nginx/          # Docker Nginx gateway (shared network maker-flow)
├── cloudflare/     # DNS CLI + reference (skill: skills/cloudflare-dns.md)
└── deploy/         # VPS push + route scripts (agent-internal)
```

## Choose a target

See [`publish/README.md`](publish/README.md) for skill file names. Shape from PRO; location from human dialogue.

## Ports (VPS only)

| Port | Role |
|------|------|
| **80** (host) | Gateway only — public entry via Cloudflare |
| `8080` / `80` (container) | MVP listen port inside its container (`CONTAINER_PORT`) |
| `3000` / `8080` (host) | **Local** `HOST_PORT` mapping for laptop acceptance only |

Production (VPS): Cloudflare → gateway `:80` → Docker alias `MVP_NAME:CONTAINER_PORT`.  
Static hosts (Pages / Vercel): no `CONTAINER_PORT` — build `dist/` and upload.

## Prerequisites (depend on target)

- **VPS:** Docker on server, SSH, optional Cloudflare DNS
- **Cloudflare Pages / Vercel / GitHub Pages:** platform auth on the agent machine; static or SPA build
