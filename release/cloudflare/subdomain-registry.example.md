# Subdomain and port registry (example)

**English** · [简体中文](subdomain-registry.example.zh-CN.md)

Copy this file to `subdomain-registry.md` and maintain it.

| Subdomain | HOST_PORT | Project dir | Status | Notes |
|-----------|-----------|-------------|--------|-------|
| test.your-domain.com | 8080 | static-test | online | path smoke test |
| idea1.your-domain.com | 8080 | my-todo-api | online | Go API |
| idea2.your-domain.com | 8081 | — | free | — |

**Rule:** register before deploy, so two people/projects do not claim the same port.
