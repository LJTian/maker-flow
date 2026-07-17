# release/

[English](README.md) · **简体中文**

**步骤 ⑥** 发布基建。Agent **MUST** 遵循 `skills/deploy.md`，在**对话中与人类确认目标**，且 **MUST** 等待步骤 ⑤ 批准。

**不要**指望人类执行 `maker-flow deploy`。该 CLI 仅供 VPS 路径的 Agent 内部使用。

## 布局

```
release/
├── publish/        # 分目标指南（Pages / Vercel / VPS）— 从这里开始
├── nginx/          # Docker Nginx 网关（共享网络 maker-flow）
├── cloudflare/     # DNS / SSL 助手（主要服务 VPS 自定义域）
└── deploy/         # VPS 推送与路由脚本（Agent 内部）
```

## 选择目标

见 [`publish/README.zh-CN.md`](publish/README.zh-CN.md)。形态跟 PRO；落点跟人类对话。

## 端口（仅 VPS）

| 端口 | 角色 |
|------|------|
| **80**（宿主机） | 仅网关 — 经 Cloudflare 的公网入口 |
| `8080` / `80`（容器内） | MVP 容器内监听（`CONTAINER_PORT`） |
| `3000` / `8080`（宿主机） | **本地** `HOST_PORT` 映射，仅本机验收 |

生产（VPS）：Cloudflare → 网关 `:80` → Docker 别名 `MVP_NAME:CONTAINER_PORT`。  
静态托管（Pages / Vercel）：无 `CONTAINER_PORT` — 构建 `dist/` 并上传。

## 前置（视目标而定）

- **VPS：** 服务器 Docker、SSH、可选 Cloudflare DNS
- **Cloudflare Pages / Vercel / GitHub Pages：** Agent 机器上的平台登录；静态或 SPA 构建
