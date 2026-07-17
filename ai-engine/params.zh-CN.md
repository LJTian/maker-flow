[English](params.md) · **简体中文**

# 参数与输出约束

OpenAI 兼容调用的参考参数（见 `ai-engine/.env.example`）。  
宿主 Agent（如 Cursor）通常忽略本文件，直接遵循 `skills/`。

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

须含模版检索结论 + **产品仓根**可运行代码，符合 `skills/mvp-assembly.md`。
