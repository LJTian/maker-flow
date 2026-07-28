# skills/

[English](README.md) · **简体中文**

面向 Agent 的权威 **HOW** 契约。Prompt 说明「问什么」；技能说明「做到什么算完成」。

**检索入口（人 + AI）：** [`CATALOG.md`](CATALOG.md) ← 先读这个

**Agent：** 默认只读英文主版（`*.md`），不要用 `.zh-CN.md` 作为技能契约。见 [`docs/i18n.md`](../docs/i18n.md)。

| 技能 | 文件 | 步骤 |
|------|------|------|
| PRO 生成 | `pro-generation.md` | 2 |
| 模版检索 | `template-matching.md` | 4 |
| MVP 组装 | `mvp-assembly.md` | 4 |
| 发布 | `deploy.md` | 6 |
| 发布：VPS 网关 | `publish-vps-gateway.md` | 6 |
| 发布：Cloudflare Pages | `publish-cloudflare-pages.md` | 6 |
| 发布：GitHub Pages | `publish-github-pages.md` | 6 |
| 发布：Vercel | `publish-vercel.md` | 6 |
| Cloudflare DNS | `cloudflare-dns.md` | 6（Cloudflare DNS 时） |

## Agent 规则

- MUST 打开 [`CATALOG.md`](CATALOG.md) 定位当前步骤对应技能。
- MUST 在行动前阅读技能全文。
- MUST 将技能中标注 MUST / MUST NOT 的章节视为硬约束。
- 若技能与 prompt 冲突，**以技能为准**。

## 扩展

新增 `skills/<name>.md`，并登记到 [`CATALOG.md`](CATALOG.md)、本表与 `docs/workflow.md`。
