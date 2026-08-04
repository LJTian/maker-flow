# 发布配方：拆分 Web + API

[English](publish-split-web-api.md) · **简体中文**

**技能 id：** `publish-split-web-api`  
**步骤：** ⑥ — 人类确认**拆分**目标后  
**依赖：** 静态发布 skill + [`publish-vps-gateway.md`](publish-vps-gateway.md)

> **Agent：** 契约以英文主版为准。

## 何时用

前端 `web-vite` 上 Pages/Vercel/GH Pages，API `go-api` 上 VPS。本地骨架见 [`templates/layouts/web-api/`](../templates/layouts/web-api/)。

## 顺序（摘要）

1. 冻结两个公网 origin（web + api）
2. 先发 API（VPS），生产设置 `CORS_ORIGINS` 为前端 origin
3. 用公网 API 地址重建 SPA（`VITE_API_BASE_URL`）
4. 再发前端到静态主机
5. **双 URL** 验收后交给人类

禁止只把带 DB 的 Go API 丢到 Pages。
