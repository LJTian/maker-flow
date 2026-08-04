# Layout catalog

**English** · [简体中文](index.zh-CN.md)

> **For AI:** Layouts are optional **product-root skeletons** used when assembling **2+ apps**. Copy apps first, then copy the layout’s root files. Layouts are never deployed alone.

| id | Path | Apps | When to use |
|----|------|------|-------------|
| `web-api` | [`web-api/`](web-api/) | `go-api` + `web-vite` | Browser UI + REST API (local compose; optional split publish) |

## Agent usage

1. Select apps via `templates/CATALOG.md`.
2. If the pair matches a layout row, copy that layout’s root files into the **product repo root**.
3. Continue with `skills/mvp-assembly.md`.
4. Split publish (Pages/Vercel frontend + VPS API): `skills/publish-split-web-api.md`.
