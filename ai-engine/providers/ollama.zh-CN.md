[English](ollama.md) · **简体中文**

# Ollama（本地 / 局域网推理）

Ollama 提供 OpenAI 兼容端点，适合 16G 显存 GPU 机作纯推理节点。

## .env 示例

```bash
AI_BASE_URL=http://192.168.1.100:11434/v1
AI_API_KEY=
AI_MODEL=deepseek-r1:14b
AI_TEMPERATURE=0.6
AI_MAX_TOKENS=8192
```

## 推理机准备

```bash
ollama pull deepseek-r1:14b
ollama pull qwen2.5-coder:14b
```

局域网暴露（仅可信网络）：

```bash
export OLLAMA_HOST=0.0.0.0:11434
```

## 验证连接

```bash
curl http://192.168.1.100:11434/v1/models
```

## 模型分工建议

| 模型 | 用途 |
|------|------|
| `deepseek-r1:14b` | 需求拆解、流程与表结构 |
| `qwen2.5-coder:14b` | Handler / 代码片段 |
