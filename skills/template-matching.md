# 模版检索技能

**适用步骤：** ④ AI 根据 PRO 检索模版  
**依赖：** `templates/index.md`（模版目录）

## 目标

根据**已确认 PRO**，从模版集中选出唯一最合适的模版，并说明理由。

## 依赖

- `prompts/03-pro-confirmed.example.md`（或项目内定稿 `pro.md`）
- [`templates/CATALOG.md`](../templates/CATALOG.md)（总览）
- `templates/index.md` + `templates/images/index.md`（明细）

## 输出

### 1. 检索结论

```markdown
## 选定模版
- **模版 ID**：go-api
- **路径**：templates/go-api
- **镜像依赖**：go-builder (`maker-flow/go-builder:1.22`) + go-runtime (`maker-flow/go-runtime:1.22`)
- **理由**：（对照 PRO 的技术需求逐条说明）
- **未选其他模版的原因**：（若仅一个模版可写「当前目录仅此一个」）
```

### 2. 匹配规则（按优先级）

| PRO 特征 | 推荐模版 | 依赖镜像（images/） |
|----------|----------|---------------------|
| REST API、Go、Gin、Docker | `go-api` | `go-builder` + `go-runtime` |
| 仅需静态通路测试 | 无代码模版，走 `release/nginx/static` | — |

阅读 `templates/CATALOG.md` → `templates/index.md` → `templates/images/index.md`。  
输出检索结论时 **MUST** 列出依赖的 `image_id` 与本地 tag。  
无完全匹配时，选最接近的并列出 PRO 中需手工调整的部分。

## 禁止

- 不得跳过目录自创脚手架
- 不得在 PRO 未确认时执行本技能
- 不得同时选多个模版（MVP 阶段单栈）
