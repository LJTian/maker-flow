# 技能库

约束 AI 在各阶段**怎么做**。Prompt 负责**说什么**，技能库负责**做到什么标准**。

| 技能 | 文件 | 用于步骤 |
|------|------|----------|
| PRO 生成 | [pro-generation.md](pro-generation.md) | ② |
| 模版检索 | [template-matching.md](template-matching.md) | ④ |
| MVP 组装 | [mvp-assembly.md](mvp-assembly.md) | ④ |
| 上线部署 | [deploy.md](deploy.md) | ⑥ |

## 使用方式

在对应 Prompt 中引用技能路径，或在对话中要求 AI 先读技能再执行：

```
请先阅读 skills/pro-generation.md，再按 prompts/02-pro-draft.md 生成 PRO。
```

## 扩展

新增场景时增加技能文件即可，例如：

- `skills/add-postgres.md` — PRO 含数据库时的组装约束
- `skills/api-only.md` — 无 DB 的纯 API MVP
