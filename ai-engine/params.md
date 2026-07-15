# 参数与输出约束

`ai-run.sh` 读取 `ai-engine/.env` 中的参数。

## 连接参数

| 变量 | 必填 | 说明 |
|------|------|------|
| `AI_BASE_URL` | 是 | OpenAI 兼容 API 根路径 |
| `AI_API_KEY` | 否 | Bearer Token |
| `AI_MODEL` | 是 | 模型标识 |

## 生成参数

| 变量 | 默认 | 步骤 ② PRO | 步骤 ④ 组装 |
|------|------|------------|-------------|
| `AI_TEMPERATURE` | `0.6` | `0.5–0.6` | `0.3–0.5` |
| `AI_MAX_TOKENS` | `8192` | `4096–8192` | `8192+` |

## 分步骤验收

### 步骤 ②（PRO）

输出须含 `skills/pro-generation.md` 全部章节，**无代码**。

### 步骤 ④（组装）

须含模版检索结论 + `workspace/<项目名>/` 可运行代码，符合 `skills/mvp-assembly.md`。

## 临时覆盖模型

```bash
./scripts/ai-run.sh prompts/04-assemble-mvp.md qwen2.5-coder:14b
```
