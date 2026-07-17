# web-vite

[English](README.md) · **简体中文**

模版 id：`web-vite`。步骤 ④ 组装用的 Vite + React + TypeScript + Tailwind 静态 Web UI。

## 能力

- Vite 6 + React 19 + TypeScript
- Tailwind CSS 3
- SPA 壳页；可通过 `VITE_API_BASE_URL` 可选探测后端 `/health`
- 多阶段 Docker：Node 构建 → Nginx 静态托管
- Nginx 提供 `GET /health`（JSON `{"status":"ok"}`）
- Compose 宿主机端口 `HOST_PORT`（默认 **3000**）

## Agent 用法

1. 拷贝本目录到**产品仓根**或 `<产品根>/web/`（多 app）。
2. 按 PRO 改写 `src/App.tsx`，在 `src/` 下新增组件。
3. 与 `go-api` 等后端联调时设置 `VITE_API_BASE_URL`。
4. 遵循 `skills/mvp-assembly.md`；patterns 可选，需要时拷到 `src/lib/`。

## 目录

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

## 容器优先

验收 **不要求** 本机 `npm install`。用 Docker 构建并运行：

```bash
cp .env.example .env
docker compose up --build
curl http://localhost:3000/health
```

本地开发（可选）：`npm install && npm run dev` → http://localhost:5173

## 与 go-api 组合

多 app 示例：`my-app/api/`（`go-api`）+ `my-app/web/`（`web-vite`）。  
`docker compose build` 前设置 `VITE_API_BASE_URL=http://localhost:8080`（或你的 API 地址）。
