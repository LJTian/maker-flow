# prompts/

[English](README.md) · **简体中文**

各阶段给 Agent 的输入契约。优先先读对应的 `skills/` 文件。

**Agent：** 默认只读英文主版（`*.md`），不要用 `.zh-CN.md` 作为 prompt/技能契约。见 [`docs/i18n.md`](../docs/i18n.md)。

| 文件 | 步骤 | Agent 职责 |
|------|------|------------|
| `01-requirement.example.md` | 1 | 采集 / 规范化人类需求 |
| `02-pro-draft.md` | 2 | 只起草 PRO（不写代码） |
| [`pro.template.md`](pro.template.md) | 2–3 | **PRO 空白骨架**（输出/定稿结构） |
| [`pro.example.md`](pro.example.md) | 2–3 | **PRO 完整样板**（待办 API） |
| `03-pro-confirmed.example.md` | 3 | 固化人工确认的 PRO（门禁产物） |
| `04-assemble-mvp.md` | 4 | 检索模版并组装（输出 = **产品仓**；见消费侧指南） |
| [`../AGENTS.consumer.example.md`](../AGENTS.consumer.example.zh-CN.md) | — | **产品仓** `AGENTS.md` 模版（复制到产品仓） |

## 规则

- 步骤 ② 正文：把需求注入 `02-pro-draft.md`（或等价消息）；结构 MUST 符合 `pro.template.md` / `skills/pro-generation.md`。
- 起草前优先阅读 `pro.example.md` 以对齐粒度。
- 步骤 ④ 正文：注入**已确认** PRO；若门禁 ③ 缺失则拒绝。

技能目录：`../skills/`。
