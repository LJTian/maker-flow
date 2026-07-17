# web-vite

**English** · [简体中文](README.zh-CN.md)

Template id: `web-vite`. Vite + React + TypeScript + Tailwind static web UI for step-4 assembly.

## Capabilities

- Vite 6 + React 19 + TypeScript
- Tailwind CSS 3
- SPA shell with optional backend health ping via `VITE_API_BASE_URL`
- Multi-stage Docker: Node build → Nginx static serve
- `GET /health` on Nginx (JSON `{"status":"ok"}`)
- Compose host port via `HOST_PORT` (default **3000**)

## Agent usage

1. Copy this directory to the **product repo root** or `<product-root>/web/` (multi-app).
2. Replace `src/App.tsx` and add components under `src/` per the PRO.
3. Set `VITE_API_BASE_URL` when pairing with `go-api` or another backend.
4. Follow `skills/mvp-assembly.md`. Patterns are optional; copy into `src/lib/` if needed.

## Layout

```
src/
  App.tsx
  main.tsx
  index.css
index.html
vite.config.ts
tailwind.config.js
Dockerfile
docker-compose.yml
nginx.conf
```

## Container-first

Do **not** require host `npm install` for acceptance. Build and serve via Docker:

```bash
cp .env.example .env
docker compose up --build
curl http://localhost:3000/health
```

Local dev (optional): `npm install && npm run dev` → http://localhost:5173

## Pairing with go-api

Multi-app example: `my-app/api/` (`go-api`) + `my-app/web/` (`web-vite`).  
Set `VITE_API_BASE_URL=http://localhost:8080` (or your API host) before `docker compose build`.
