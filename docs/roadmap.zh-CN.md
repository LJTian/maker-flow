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
| Pattern | `templates/patterns/` 下 14 个（含 `persistence-sqlx`、鉴权、支付、cron 等） |
| 多 app 布局 | `templates/layouts/web-api/` + `skills/publish-split-web-api.md` |
| CI | `scripts/check.sh`、pattern `go test`、app compose / docker build |
| 线上示例 | [静态介绍站 → GitHub Pages](examples/static-intro-github-pages.zh-CN.md) |

更早缺口（验收技能、sqlx、拆分发布、compose CI）**已完成**。下文是剩余 backlog。

---

## 优先级 1 — 对真实孵化影响最大

先做这些，才能覆盖「不止一个静态落地页」的公网验证路径。

| ID | 项 | 当前缺口 | 规划 |
|----|------|----------|------|
| P1-1 | **更多端到端示例** | 只有一条线上 walkthrough（`web-vite` → GitHub Pages） | 补 walkthrough（最好带 live URL）：`go-api` → VPS 网关；`web-api` 拆分（Pages/Vercel SPA + VPS API）。登记到 `docs/examples/` |
| P1-2 | **静态发布脚本对齐** | VPS 有 `release/deploy/push-and-route.sh`；Pages / Vercel / GH Pages 主要靠 skill + 手工命令 | 在 `release/publish/`（或同级目录）增加薄封装脚本（构建 + 上传），SOP 仍放在 `skills/publish-*.md` |
| P1-3 | **鉴权 + DB 拼装样例** | `auth-oauth-jwt` README 仍有 `TODO: Save to DB` | 文档 + 可选片段：OAuth 用户 → `persistence-sqlx`（users 表、find-or-create、JWT subject） |
| P1-4 | **支付 + DB 拼装样例** | `payment-lemonsqueezy` 仍有 `TODO: Upgrade VIP` | 同样：webhook → sqlx 升级路径；README 保留 MoR 流程 |
| P1-5 | **支付目录诚实标注** | 标签含 `stripe` / 支付宝 / 微信，实现只有 Lemon Squeezy | 要么补实现，要么收窄标签/文案，避免 Agent 误选 |

**建议顺序：** P1-1 → P1-2 → P1-3 → P1-4 → P1-5。

---

## 优先级 2 — 工厂成熟度

| ID | 项 | 当前缺口 | 规划 |
|----|------|----------|------|
| P2-1 | **打版本发布** | CLI `VERSION=0.5.0`，`CHANGELOG` 仍全在 `[Unreleased]`，无 git tag | 打首个 tag；把 Unreleased 迁入版本节；继续 semver + Keep a Changelog |
| P2-2 | **中文文档对齐** | 已链接但缺失：`CONTRIBUTING.zh-CN.md`；缺 `skills/publish-*.zh-CN.md`、`prompts/pro-multi.example.zh-CN.md` 等 | 给人看的文档补 ZH；英文仍是 Agent 契约（见 [i18n.zh-CN.md](i18n.zh-CN.md)） |
| P2-3 | **`web-vite` 测试** | CONTRIBUTING 要求 app 有测试；Vite 模版没有 | 加最小 Vitest（或等价）冒烟；成本低则接入 CI |
| P2-4 | **Pattern `go.sum` 覆盖** | 多数 pattern 未提交 `go.sum`；CI 依赖 `go mod tidy` | 每个 Go pattern 提交 `go.sum`（对齐 `cron-scheduler` / `persistence-sqlx`） |
| P2-5 | **CONTRIBUTING 与现实一致** | 写明 skill 要 YAML frontmatter；现有 skill 都没有 | 要么给 skill 补 frontmatter，要么改 CONTRIBUTING — 二选一 |
| P2-6 | **栈边界写清** | 目录只有 Go + Vite，易被当成已有 Python/Node API | 在 `templates/CATALOG.md`（必要时 README）写明「支持 / 暂不做」；新栈必须有人维护 CI + 验收 |

---

## 优先级 3 — 打磨（有空再做）

| ID | 项 | 规划 |
|----|------|------|
| P3-1 | 轻量鉴权 pattern | 可选 API key / signed-cookie，覆盖「个人工具要登录」但不想上完整 OAuth |
| P3-2 | 清理 `.gitignore` | 去掉重复的 `.gomodcache/` / `.gocache/` |
| P3-3 | `ai-engine/` | 保持可选文档；除非真需要脱离 Cursor 的自托管 LLM 通道 |
| P3-4 | 队列后端 | 默认不进 PRO（无 Redis/Asynq，除非产品真需要） |

---

## 现阶段明确不做

- 把组装后的 MVP **写进**本工厂仓（只写产品仓）
- 步骤 ② 的 PRO 要求 K8s / 微服务拆分
- 用强制自托管 LLM 替换宿主 Agent（Cursor / Claude）
- 在 Lemon Squeezy + DB 接线扎实前，追求完整 Stripe / 国内支付 MoR 对等

---

## 怎么用这份文档

- **人类：** 从优先级 1 选下一个 ID，开 Issue/PR；做完后缩小或划掉对应行。
- **Agent：** **不要**把本文件当成流水线步骤。执行仍以 `docs/workflow.md` + `skills/*` 为准。
- 某项落地后：更新本表（划掉或移入下方「已完成」），并记入 `CHANGELOG.md`。

## 已完成（近期关掉的缺口）

| 项 | 位置 |
|------|------|
| MVP 验收技能 + prompt | `skills/mvp-acceptance.md`、`prompts/05-accept-mvp.md` |
| `persistence-sqlx` pattern | `templates/patterns/persistence-sqlx/` |
| 前后端拆分发布 | `skills/publish-split-web-api.md`、`templates/layouts/web-api/` |
| 各发布目标技能 | `skills/publish-*.md` |
| Compose / pattern CI | `.github/workflows/ci.yml` |
| OAuth / cron / 支付等 pattern | `templates/patterns/` |

---

## 相关链接

- [快速开始](getting-started.zh-CN.md)
- [示例](examples/)
- [模版目录](../templates/CATALOG.md)
- [技能目录](../skills/CATALOG.md)
- [Changelog](../CHANGELOG.md)
- [贡献指南](../CONTRIBUTING.md)
