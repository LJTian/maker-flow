# 下一步计划（Roadmap）

[English](roadmap.md) · **简体中文**

本文件是**工厂仓**（技能、模版、发布脚本、文档）的活清单，不是某个具体 MVP 的功能列表。

**最近审阅：** 2026-08-11

## 已经比较扎实的部分

六步流水线对 Go API + Vite SPA 类 MVP 已可端到端使用：

| 区域 | 状态 |
|------|------|
| 工作流与硬门禁 | `docs/workflow.md`、`AGENTS.md` |
| 技能（PRO → 匹配 → 组装 → 验收 → 发布） | `skills/CATALOG.md` |
| App 模版 | `go-api`、`go-cli`、`go-worker`、`web-vite` |
| Pattern | `templates/patterns/` 下 14 个（含 `persistence-sqlx`、OAuth 登录、Lemon Squeezy Webhook、cron 等） |
| 多 app 布局 | `templates/layouts/web-api/` + `skills/publish-split-web-api.md` |
| CI | `scripts/check.sh`、pattern `go test`、app compose / docker build |
| 线上示例 | [静态介绍站 → GitHub Pages](examples/static-intro-github-pages.zh-CN.md) |

---

## 未实现清单（诚实版）

目录/技能里曾经虚标，或至今仍缺的能力。Agent **禁止**假装工厂里已经有这些东西。

### Pattern / 支付 / 鉴权

| ID | 缺失项 | 说明 |
|----|--------|------|
| U-1 | **原生微信支付** pattern | `auth-oauth-jwt` 有微信**登录**；**收款**没有 |
| U-2 | **原生支付宝** pattern | 目录尚无 |
| U-3 | **原生 Stripe** pattern | 目录尚无（`payment-lemonsqueezy` 仅 MoR Webhook） |
| U-4 | 鉴权 → DB 拼装样例 | `auth-oauth-jwt` 仍有 `TODO: Save to DB` |
| U-5 | 支付 → DB 升级样例 | `payment-lemonsqueezy` 仍有 `TODO: Upgrade VIP` |
| U-6 | 生产级 **Sign in with Apple** | 当前 `apple.New(..., nil)` 仅为薄骨架 |
| U-7 | 轻量鉴权（API key / signed cookie） | 尚无；目前只有完整 OAuth |
| U-8 | 通用 **SMTP** 邮件 | `notify-email` **仅 Resend** |
| U-9 | 原生 **Anthropic** SDK | `ai-llm-client` 仅为 OpenAI 兼容 |

### App / 发布 / 示例

| ID | 缺失项 | 说明 |
|----|--------|------|
| U-10 | Python / Node API app | 目录仅 **Go + Vite** |
| U-11 | Vercel SSR / **Next.js** 模版 | `publish-vercel` 仅静态/SPA |
| U-12 | 更多线上示例 | 需 `go-api`→VPS、`web-api` 拆分 walkthrough（现仅静态 Pages） |
| U-13 | 静态发布**脚本** | Pages / Vercel / GH Pages 缺少与 VPS `push-and-route.sh` 对等的脚本 |

### 文档 / 工厂卫生

| ID | 缺失项 | 说明 |
|----|--------|------|
| U-14 | `CONTRIBUTING.zh-CN.md` | `CONTRIBUTING.md` 已链接但文件不存在 |
| U-15 | `skills/publish-*.zh-CN.md` | 多个发布技能缺中文姐妹 |
| U-16 | `prompts/pro-multi.example.zh-CN.md` | 英文有、中文缺 |
| U-17 | 首个 **git tag** / Changelog 切版 | CLI 写 `0.5.0`；Changelog 仍 `[Unreleased]` |
| U-18 | `web-vite` 自动化测试 | CONTRIBUTING 要求 app 有测；Vite 没有 |
| U-19 | 全部 Go pattern 的 `go.sum` | 多数只靠 CI `go mod tidy` |
| U-20 | Skill YAML frontmatter vs CONTRIBUTING | CONTRIBUTING 写了 frontmatter；现有 skill 都没有 — 需二选一 |

---

## 优先级 1 — 对真实孵化影响最大

| ID | 项 | 对应 |
|----|------|------|
| P1-1 | 更多端到端示例 | **U-12** |
| P1-2 | 静态发布脚本对齐 | **U-13** |
| P1-3 | 鉴权 + DB 拼装样例 | **U-4** |
| P1-4 | 支付 + DB 拼装样例 | **U-5** |
| P1-6 | 原生微信支付 pattern（可选下一步） | **U-1** |
| P1-7 | 原生支付宝 / Stripe pattern（可选） | **U-2** / **U-3** |

**建议顺序：** P1-1 → P1-2 → P1-3 → P1-4 → 若需要国内收款或 Stripe 再做 P1-6/P1-7。

---

## 优先级 2 — 工厂成熟度

| ID | 项 | 覆盖 |
|----|------|------|
| P2-1 | 打版本发布 | U-17 |
| P2-2 | 中文文档对齐 | U-14、U-15、U-16 |
| P2-3 | `web-vite` 测试 | U-18 |
| P2-4 | Pattern `go.sum` 覆盖 | U-19 |
| P2-5 | CONTRIBUTING 与现实一致（frontmatter） | U-20 |
| P2-6 | 栈边界写清 | 已在 `templates/CATALOG.md`（+ 链到本文）；保持更新 |

---

## 优先级 3 — 打磨

| ID | 项 | 覆盖 |
|----|------|------|
| P3-1 | 轻量鉴权 pattern | U-7 |
| P3-2 | 清理 `.gitignore` 重复项 | — |
| P3-3 | 加固 Apple 登录文档/接线 | U-6 |
| P3-4 | 队列后端（Redis/Asynq） | 默认非目标，除非产品真需要 |
| P3-5 | SMTP / 多厂商邮件 | U-8 |
| P3-6 | Anthropic 原生客户端 | U-9 |

---

## 现阶段明确不做

- 把组装后的 MVP **写进**本工厂仓（只写产品仓）
- 步骤 ② 的 PRO 要求 K8s / 微服务拆分
- 用强制自托管 LLM 替换宿主 Agent（Cursor / Claude）
- 假装 Lemon Squeezy 等于原生微信支付 / 支付宝 / Stripe

---

## 怎么用这份文档

- **人类：** 从优先级 1 选下一个 ID，开 Issue/PR；做完后把对应 **U-*** 行移入「已完成」。
- **Agent：** **不要**把本文件当成流水线步骤。执行仍以 `docs/workflow.md` + `skills/*` 为准。**禁止**把缺失的支付厂商当成 `payment-lemonsqueezy` 来选。
- 某项落地后：更新本表，并记入 `CHANGELOG.md`。

## 已完成（近期关掉的缺口）

| 项 | 位置 |
|------|------|
| 目录/技能诚实化（去掉假的 `stripe`/`alipay`/`wechat` 收款标签） | `templates/CATALOG*`、`patterns/index*`、`skills/template-matching*`、支付/鉴权 README |
| MVP 验收技能 + prompt | `skills/mvp-acceptance.md`、`prompts/05-accept-mvp.md` |
| `persistence-d1` pattern (线上 Cloudflare D1 + 本地 Docker) | `templates/patterns/persistence-d1/` |
| 前后端拆分发布 | `skills/publish-split-web-api.md`、`templates/layouts/web-api/` |
| 各发布目标技能 | `skills/publish-*.md` |
| Compose / pattern CI | `.github/workflows/ci.yml` |

---

## 相关链接

- [快速开始](getting-started.zh-CN.md)
- [示例](examples/)
- [模版目录](../templates/CATALOG.md)
- [技能目录](../skills/CATALOG.md)
- [Changelog](../CHANGELOG.md)
- [贡献指南](../CONTRIBUTING.md)
