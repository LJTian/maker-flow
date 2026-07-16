# OpenAI

**English** · [简体中文](openai.zh-CN.md)

## .env example

```bash
AI_BASE_URL=https://api.openai.com/v1
AI_API_KEY=sk-your-key
AI_MODEL=gpt-4o-mini
AI_TEMPERATURE=0.6
AI_MAX_TOKENS=8192
```

## Notes

- Keep the API key only in `ai-engine/.env`; do not commit it
- Tune `AI_MAX_TOKENS` to account quotas
- For code generation you may prefer stronger models such as `gpt-4o` (higher cost)
