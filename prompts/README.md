# 分阶段 Prompt

| 文件 | 步骤 | 说明 |
|------|------|------|
| [01-requirement.example.md](01-requirement.example.md) | ① | 用户需求 |
| [02-pro-draft.md](02-pro-draft.md) | ② | 生成 PRO（无代码） |
| [03-pro-confirmed.example.md](03-pro-confirmed.example.md) | ③ | 定稿 PRO（卡点） |
| [04-assemble-mvp.md](04-assemble-mvp.md) | ④ | 检索模版 + 组装 MVP |

```bash
./scripts/ai-run.sh prompts/02-pro-draft.md
# 确认 PRO 后
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

配套技能见 [skills/](../skills/)。
