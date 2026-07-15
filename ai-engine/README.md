# The AI Engine

可配置的 **AI 连接层** — 不限定厂商，支持 OpenAI 兼容 API。

流程与步骤见 [docs/workflow.md](../docs/workflow.md)；各阶段 SOP 见 [skills/](../skills/)；分阶段 Prompt 见 [prompts/](../prompts/)。

## 目录

```
ai-engine/
├── .env.example        # 连接与生成参数
├── params.md           # 参数说明与输出约束
└── providers/          # 各后端配置示例
```

## 快速配置

```bash
cp ai-engine/.env.example ai-engine/.env
# 参考 providers/ 填写 AI_BASE_URL、AI_MODEL
```

## 执行 Prompt

```bash
./scripts/ai-run.sh prompts/02-pro-draft.md      # 步骤 ②
./scripts/ai-run.sh prompts/04-assemble-mvp.md # 步骤 ④
```

## 相关文档

- [参数与约束](params.md)
- [后端示例](providers/)
- [六步工作流](../docs/workflow.md)
