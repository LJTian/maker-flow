[English](openai.md) · **简体中文**

# OpenAI

## .env 示例

```bash
AI_BASE_URL=https://api.openai.com/v1
AI_API_KEY=sk-your-key
AI_MODEL=gpt-4o-mini
AI_TEMPERATURE=0.6
AI_MAX_TOKENS=8192
```

## 注意

- API Key 仅写在 `ai-engine/.env`，勿提交仓库
- 按账号配额调整 `AI_MAX_TOKENS`
- 代码生成可选用 `gpt-4o` 等更强模型，成本更高
