# Subdomain registry (example)

**English** · [简体中文](subdomain-registry.example.zh-CN.md)

Copy this file to `subdomain-registry.md` and maintain it.

| Subdomain | MVP_NAME | CONTAINER_PORT | Service | Status | Notes |
|-----------|----------|----------------|---------|--------|-------|
| idea1.your-domain.com | idea1 | 8080 | api | online | Go API |
| idea2.your-domain.com | idea2 | 80 | web | online | web-vite |
| free.your-domain.com | — | — | — | free | — |

**Rule:** register before deploy so two projects do not claim the same `MVP_NAME` (Docker network alias) or subdomain.

Production traffic uses the gateway on host port **80** and Docker network aliases — host `8080–8090` mappings are optional for local debug only.
