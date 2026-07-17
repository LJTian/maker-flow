# 发布技能（步骤 ⑥）

[English](deploy.md) · **简体中文**

**适用步骤：** ⑥ 上线  
**前置：** 步骤 ⑤ MVP 已本地验收通过  
**技能 id：** `deploy`（文件名保持不变，便于目录兼容）

## 目标

把已验收的 MVP 暴露到公网。**形态**（静态 / 常驻运行时）由 PRO 决定；**发到哪里**由**人类在对话中选定**——不是给人用的 CLI。

## 硬规则

- 执行任何发布动作前，**必须先询问人类**选用哪个（哪些）发布目标。
- **禁止**让人类去跑 `maker-flow deploy`（该 CLI 仅供 **Agent 内部**调用）。
- 未通过步骤 ⑤ **禁止**发布。
- **禁止**不可能的组合（例如带 Postgres 的 API 单独上 Cloudflare Pages）。应提议拆分（静态前端 + VPS API）。
- 静态 vs 非静态：跟 PRO / 已组装 app 走；不要默认 VPS 或 Pages。

## 对话门禁（必做）

执行前在对话中确认：

1. **发什么：** 整站 / 仅前端 / 仅 API / worker（可不对公网）？
2. **发到哪**（可多选）：
   - `vps-gateway` — VPS 上 Docker + 共享 Nginx 网关
   - `cloudflare-pages` — Cloudflare Pages
   - `github-pages` — GitHub Pages
   - `vercel` — Vercel
3. **域名：** 平台默认 URL，还是自定义域名
4. **凭证：** 人类确认已登录平台 / Token 可用（禁止编造密钥）

人类答完后，再按 `release/publish/` 下对应指南执行。

## 端口（仅 VPS 路径）

三层不要混：

| 层 | 含义 | web-vite | go-api |
|----|------|----------|--------|
| 本地 `HOST_PORT` | 本机浏览器 | `3000` → 容器 | `8080` → 容器 |
| `CONTAINER_PORT` | **容器内**监听端口 | **80** | **8080** |
| 公网入口 | Cloudflare → 网关宿主机 `:80` | 始终是网关 80 | 同左 |

Agent 内部 VPS 发布用的是 **`CONTAINER_PORT`**，不是 `HOST_PORT`。

| App 模版 | Compose 服务名 | `CONTAINER_PORT` |
|----------|----------------|------------------|
| `go-api` | `api` | `8080` |
| `web-vite` | `web` | `80` |
| `go-worker` | `worker` | 通常**不**对公网 |

## 目标矩阵

| 目标 | 适合 | 不适合 | Agent 指南 |
|------|------|--------|------------|
| `vps-gateway` | API、worker、整包 Compose、自托管静态 | 没有 VPS 的用户 | [`release/publish/vps-gateway.md`](../release/publish/vps-gateway.md) |
| `cloudflare-pages` | 静态 / SPA（`web-vite` 构建） | DB、长驻 Go API | [`release/publish/cloudflare-pages.md`](../release/publish/cloudflare-pages.md) |
| `github-pages` | 静态 / SPA | 同上 | [`release/publish/github-pages.md`](../release/publish/github-pages.md) |
| `vercel` | 静态 / SPA | 未改造就塞自建 Postgres | [`release/publish/vercel.md`](../release/publish/vercel.md) |

混合产品：人类若要求，前端上 Pages/Vercel，API 走 `vps-gateway`。

## 发布后

- 把公网 URL 交给人类。
- 按情况验证（`curl` / 打开 `/` 或 `/health`）。
- 若人类维护登记表，记下子域 / 项目名。

## 回滚

见所选 `release/publish/<target>.md` 的回滚小节。

## 详细文档

- [`release/publish/README.md`](../release/publish/README.md)
- [`release/README.md`](../release/README.md)
- 对话提示：[`prompts/06-publish.md`](../prompts/06-publish.md)
