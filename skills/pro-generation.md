# PRO 生成技能

**适用步骤：** ② AI 根据需求定 PRO  
**配套 Prompt：** `prompts/02-pro-draft.md`  
**空白骨架：** [`prompts/pro.template.md`](../prompts/pro.template.md)  
**完整样板：** [`prompts/pro.example.md`](../prompts/pro.example.md)

## 目标

把用户一句话需求转化为可确认的 **PRO**，供步骤 ③ 人工审核。

## 禁止

- 不输出完整项目代码
- 不选定最终模版（留到步骤 ④；检索提示仅作线索）
- 不假设用户已确认范围

## PRO 必须包含的章节

输出结构 MUST 与 `prompts/pro.template.md` 一致。粒度参考 `prompts/pro.example.md`。

### 1. 摘要

- 一句话目标
- MVP 范围（1–2 天可完成）
- 明确 **不做** 的功能

### 2. 业务流程

编号步骤，写清主流程与边界（是否需要登录、数据归属等）。

### 3. 数据模型

Markdown 表格：字段、类型、说明、约束。  
可选：`CREATE TABLE` 语句（若 PRO 含持久化）。无持久化须显式写明。

### 4. 接口契约

每个端点/命令：`METHOD /path`（或 CLI 子命令）、请求/响应示例、主要错误码。  
至少：健康检查（若适用）+ 核心业务 2–4 个接口。  
多 app 时按 app 分小节。

### 5. 验收标准

可勾选的 checklist，供步骤 ⑤ 使用。例如：

- [ ] `GET /health` 返回 200
- [ ] 可创建并列表查询待办
- [ ] …

### 6. 模版检索提示（可选）

给步骤 ④ 的线索，不最终拍板：

- 倾向 **apps（1～N）**（如 `go-api`，或 `go-api` + `go-worker`）
- 倾向 **patterns（0～N）**（可写「无」）
- 是否需要 DB / 预期 QPS / 复杂度

## 质量约束

- 范围控制在 MVP，禁止 K8s、消息队列、微服务拆分
- 默认无鉴权，除非需求明确要求
- 术语与 `templates/index.md` / `templates/patterns/index.md` 标签可对齐，便于后续检索

## 输出格式

使用 Markdown，章节标题与 `pro.template.md` 一致，便于用户 diff 和注释修改。
