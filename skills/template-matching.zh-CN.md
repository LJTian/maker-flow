# 模版检索技能

[English](template-matching.md) · **简体中文**

**适用步骤：** ④ AI 根据 PRO 检索模版  
**依赖：** `templates/CATALOG.md` → `templates/index.md` → `templates/patterns/index.md`

## 目标

1. 选定 **1～N 个 app**（`templates/apps/`）— 按 PRO 需要的形态组合，例如 API + worker  
2. 选定 **0～N 个 pattern**（`templates/patterns/`，按 tags）  
3. 列出依赖的 **image** tags（对所选 apps 取并集）

## 输入

- 已确认 PRO
- `templates/CATALOG.md`

## 输出

```markdown
## 选定模版
- **Apps**：
  - go-api → templates/apps/go-api
  - go-worker → templates/apps/go-worker
- **镜像依赖**：go-builder + go-runtime
- **Patterns**：retry-backoff, worker-pool（可为空）
- **产品布局**：`<产品根>/{api,worker}/`（或分别说明）
- **理由**：…
```

## 匹配规则

| PRO 特征 | App | 常用 Patterns |
|----------|-----|---------------|
| REST API、Gin | `go-api` | `retry-backoff`, `circuit-breaker`, `singleflight-cache`, `persistence-sqlx` |
| 表结构 / 数据库 / 持久化 | `go-api` | `persistence-sqlx`（`DB_DRIVER`：sqlite \| postgres \| mysql） |
| CLI / 命令行工具 | `go-cli` | `retry-backoff`, `worker-pool` |
| 后台任务 / 多协程消费 | `go-worker` | `worker-pool`, `pipeline`, `retry-backoff` |
| 浏览器 UI / SPA / 面板 | `web-vite` | —（可选片段放 `src/lib/`） |
| API + 浏览器 UI | `go-api` + `web-vite` | 布局 [`web-api`](../templates/layouts/web-api/)；patterns 按需 |

- **数据库 / 持久化存储**：选中 `persistence-sqlx`（用于 SQL/SQLite/Postgres）或 `persistence-d1`（用于线上 Cloudflare D1 + 本地 Docker）。
- **用户登录 / Auth:** 若需要登录 / OAuth（Google、GitHub、微信；Apple 仅为薄骨架）/ JWT，加入 `auth-oauth-jwt`。这是**登录**，不是微信支付。
- **支付 / 订阅:** **仅**在使用 Lemon Squeezy MoR 结账 + Webhook 验签时加入 `payment-lemonsqueezy`。**MUST NOT** 用它冒充原生微信支付、支付宝或 Stripe — 这些 pattern 尚未入库（见 `docs/roadmap.md`）。
- **通知 / 邮件:** 事务邮件用 `notify-email`（仅 Resend，无通用 SMTP pattern）。
- **存储 / 上传:** S3 兼容对象存储用 `storage-s3`（AWS S3 / R2 / MinIO，经 `aws-sdk-go-v2` 自定义 endpoint）。
- **AI / 大模型:** `ai-llm-client` 为 **OpenAI 兼容**流式客户端（ChatGPT 或兼容 `/v1/chat/completions` 的网关；非 Anthropic 原生 SDK）。
- **定时任务:** 若明确要求定时调度、周期性后台任务或 Cron，请加入 `cron-scheduler`。
- **数据埋点:** 若需要用户追踪、使用量分析或对接 PostHog，请加入 `telemetry-posthog`。
- **安全 / 限流:** 若需要防止接口被刷、账单攻击或限制 IP 请求频率，请加入 `rate-limiter`。
- **系统韧性:** 若要求对外部系统进行高可靠性调用，请加入 `circuit-breaker` 或 `retry-backoff`。
- **并发处理:** 若需要高吞吐量的后台处理，请加入 `worker-pool` 或 `pipeline`。

多 app 示例：`go-api` + `go-worker`（同步 API + 异步消费）；`go-api` + `go-cli`（服务 + 运维命令）；`go-api` + `web-vite`（API + 浏览器 UI — 拷贝 [`templates/layouts/web-api/`](../templates/layouts/web-api/) 根文件）。

匹配 `go-api` + `web-vite` 时，选型输出中 **必须** 写明布局 `web-api`（或说明为何用自定义根 compose）。

## 禁止

- 不得跳过目录自创脚手架
- 不得在 PRO 未确认时执行
- 不得把 pattern 当作独立公网服务部署
- 不得为无关形态硬凑 app（每个 app 须能对应 PRO 中的职责）
