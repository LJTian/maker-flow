# Scripts

| 脚本 | 说明 |
|------|------|
| [ai-run.sh](ai-run.sh) | 读取 `ai-engine/.env`，将 Prompt 发送至 AI 后端 |

```bash
cp ai-engine/.env.example ai-engine/.env

# 步骤 ② 生成 PRO
./scripts/ai-run.sh prompts/02-pro-draft.md

# 步骤 ④ 组装 MVP
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

部署脚本见 [release/deploy/](../release/deploy/)。
