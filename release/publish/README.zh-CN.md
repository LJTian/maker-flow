# release/publish/

[English](README.md) · **简体中文**

**读者：Agent**（步骤 ⑥）。人类在对话里选目标；见 `skills/deploy.md`。

## 原则

- **跑什么** → PRO / 模版  
- **发到哪** → 人类对话确认  
- **怎么发** → 本目录（VPS 另见 `release/deploy/`）

**不要**让人类执行 `maker-flow deploy`。该 CLI 仅供 Agent 在 VPS 路径内部调用。

## 目标

| id | 文件 |
|----|------|
| `vps-gateway` | [vps-gateway.md](vps-gateway.md) |
| `cloudflare-pages` | [cloudflare-pages.md](cloudflare-pages.md) |
| `github-pages` | [github-pages.md](github-pages.md) |
| `vercel` | [vercel.md](vercel.md) |

## 相关

- 技能：[`skills/deploy.md`](../../skills/deploy.md)
- 提示：[`prompts/06-publish.md`](../../prompts/06-publish.md)
- VPS 脚本：[`../deploy/`](../deploy/)
- DNS：[`../cloudflare/`](../cloudflare/)
