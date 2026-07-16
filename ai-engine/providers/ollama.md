# Ollama (local / LAN inference)

**English** · [简体中文](ollama.zh-CN.md)

Ollama exposes an OpenAI-compatible endpoint; suitable as a pure inference node on a ~16GB VRAM GPU host.

## .env example

```bash
AI_BASE_URL=http://192.168.1.100:11434/v1
AI_API_KEY=
AI_MODEL=deepseek-r1:14b
AI_TEMPERATURE=0.6
AI_MAX_TOKENS=8192
```

## Inference host prep

```bash
ollama pull deepseek-r1:14b
ollama pull qwen2.5-coder:14b
```

LAN bind (trusted networks only):

```bash
export OLLAMA_HOST=0.0.0.0:11434
```

## Verify connectivity

```bash
curl http://192.168.1.100:11434/v1/models
```

## Suggested model roles

| Model | Use |
|-------|-----|
| `deepseek-r1:14b` | Requirements breakdown, flow and schema |
| `qwen2.5-coder:14b` | Handlers / code snippets |
