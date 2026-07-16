# OpenAI-compatible gateways

**English** · [简体中文](openai-compatible.zh-CN.md)

For DeepSeek, Moonshot, SiliconFlow, OneAPI, LiteLLM, and similar forwarding services.

## .env template

```bash
AI_BASE_URL=https://your-gateway.example.com/v1
AI_API_KEY=your-api-key
AI_MODEL=deepseek-chat
AI_TEMPERATURE=0.6
AI_MAX_TOKENS=8192
```

## DeepSeek official example

```bash
AI_BASE_URL=https://api.deepseek.com/v1
AI_API_KEY=sk-...
AI_MODEL=deepseek-chat
```

## Selection notes

- Confirm the gateway supports `POST /v1/chat/completions` and `stream: true`
- Model names follow gateway docs; they may differ from Ollama local tags
