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
| REST API、Gin | `go-api` | `retry-backoff`, `circuit-breaker`, `singleflight-cache` |
| CLI / 命令行工具 | `go-cli` | `retry-backoff`, `worker-pool` |
| 后台任务 / 多协程消费 | `go-worker` | `worker-pool`, `pipeline`, `retry-backoff` |
| 浏览器 UI / SPA / 面板 | `web-vite` | —（可选片段放 `src/lib/`） |

多 app 示例：`go-api` + `go-worker`（同步 API + 异步消费）；`go-api` + `go-cli`（服务 + 运维命令）；`go-api` + `web-vite`（API + 浏览器 UI）。

## 禁止

- 不得跳过目录自创脚手架
- 不得在 PRO 未确认时执行
- 不得把 pattern 当作独立公网服务部署
- 不得为无关形态硬凑 app（每个 app 须能对应 PRO 中的职责）
