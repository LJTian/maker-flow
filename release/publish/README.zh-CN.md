# release/publish/

[English](README.md) · **简体中文**

步骤 ⑥ 发布目标索引。**Agent 执行 SOP 在 `skills/`**，本目录只作跳转。

| 目标 id | 技能 |
|---------|------|
| `vps-gateway` | [`skills/publish-vps-gateway.md`](../../skills/publish-vps-gateway.md) |
| `cloudflare-pages` | [`skills/publish-cloudflare-pages.md`](../../skills/publish-cloudflare-pages.md) |
| `github-pages` | [`skills/publish-github-pages.md`](../../skills/publish-github-pages.md) |
| `vercel` | [`skills/publish-vercel.md`](../../skills/publish-vercel.md) |
| `split-web-api` | [`skills/publish-split-web-api.md`](../../skills/publish-split-web-api.md)（SPA 主机 + VPS API） |

## 流程

1. 人类在对话里选目标 — [`skills/deploy.md`](../../skills/deploy.md) + [`prompts/06-publish.md`](../../prompts/06-publish.md)
2. Agent 加载对应的 `skills/publish-<target>.md`
3. 脚本与基建：[`../deploy/`](../deploy/)（VPS）、[`../cloudflare/`](../cloudflare/)（DNS CLI）、[`../nginx/`](../nginx/)（网关）

**不要**让人类运行 `maker-flow deploy`（VPS 路径仅 Agent 内部使用）。

## 相关

- 技能目录：[`skills/CATALOG.md`](../../skills/CATALOG.md)
- 发布总览：[`../README.md`](../README.md)
