# OpenAI 兼容网关

适用于 DeepSeek、Moonshot、硅基流动、OneAPI、LiteLLM 等转发服务。

## .env 模板

```bash
AI_BASE_URL=https://your-gateway.example.com/v1
AI_API_KEY=your-api-key
AI_MODEL=deepseek-chat
AI_TEMPERATURE=0.6
AI_MAX_TOKENS=8192
```

## DeepSeek 官方示例

```bash
AI_BASE_URL=https://api.deepseek.com/v1
AI_API_KEY=sk-...
AI_MODEL=deepseek-chat
```

## 选型要点

- 确认网关支持 `POST /v1/chat/completions` 与 `stream: true`
- 模型名以网关文档为准，与 Ollama 本地 tag 可能不同
