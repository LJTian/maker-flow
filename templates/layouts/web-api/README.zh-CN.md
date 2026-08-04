# 布局：web-api

[English](README.md) · **简体中文**

**布局 id：** `web-api`  
**Apps：** `go-api` + `web-vite`  
**何时用：** PRO 同时需要浏览器 UI 与 REST API。

> **Agent：** 以英文主版 [`README.md`](README.md) 为准。

本目录是**产品仓根骨架**，不是单独可部署服务。先把 apps 拷到 `api/`、`web/`，再把本布局的 `docker-compose.yml`、`.env.example` 拷到产品仓根。

## 本地

```bash
cp -n .env.example .env
docker compose up --build
```

关键：`VITE_API_BASE_URL`（构建时写入前端）、`CORS_ORIGINS`（API 允许的前端源）。

## 拆分上线

SPA → Pages / Vercel / GH Pages，API → VPS：见 [`skills/publish-split-web-api.md`](../../skills/publish-split-web-api.md)。
