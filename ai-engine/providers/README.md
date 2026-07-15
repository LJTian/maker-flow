# 后端配置示例

`ai-engine` 通过 **OpenAI 兼容 API** 连接任意后端。复制对应示例中的变量到 `ai-engine/.env` 即可。

| 示例 | 适用场景 |
|------|----------|
| [ollama.md](ollama.md) | 局域网 GPU 机自建推理 |
| [openai.md](openai.md) | OpenAI 官方 API |
| [openai-compatible.md](openai-compatible.md) | DeepSeek、硅基流动等兼容网关 |

若后端不支持某项生成参数（如 `top_p`），可忽略，脚本会照常发送。
