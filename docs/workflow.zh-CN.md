# 流程（Agent 契约）

[English](workflow.md) · **简体中文**

Maker Flow Agent 的权威状态机。

```
① 人：需求
        │
        ▼
② Agent：起草 PRO ──► ③ 人：确认 PRO  [门禁]
        │
        ▼
④ Agent：匹配模版并组装 MVP
        │
        ▼
⑤ 人：确认 MVP  [门禁]
        │
        ▼
⑥ Agent：部署
```

## 步骤契约

### 1 — 需求（人）

输入：简短自然语言需求。  
Agent：若使用 `prompts/`，写入步骤 ② prompt 的用户需求区。

### 2 — 起草 PRO（Agent）

- **MUST 阅读：** `skills/pro-generation.md`
- **MUST 遵循结构：** `prompts/pro.template.md`（粒度见 `prompts/pro.example.md`）
- **MAY 使用：** `prompts/02-pro-draft.md` 作为 prompt 正文
- **MUST 输出：** 技能 / 模版规定的 PRO 章节
- **MUST NOT：** 写应用代码、最终选定模版、或组装进工厂仓

### 3 — 确认 PRO（人门禁）

- Agent 出示 PRO 并等待。
- 通过后将定稿 PRO 写入**产品仓**的 `pro.md`（或工厂示例 `prompts/03-pro-confirmed.example.md`；结构同 `pro.template.md`）。
- **未获明确人类批准 MUST NOT 进入步骤 ④。**

### 4 — 组装 MVP（Agent）

- **MUST 按序阅读：**
  1. `skills/template-matching.md`
  2. `templates/CATALOG.md` → `templates/index.md` → `templates/patterns/index.md` → `templates/images/index.md`
  3. `skills/mvp-assembly.md`
- **MAY 使用：** `prompts/04-assemble-mvp.md`
- **MUST：** 选定 **1～N 个** app ID，拷贝到**产品仓根**（多 app：`<产品根>/<app-id>/`），只实现 PRO 范围
- **MUST NOT：** 在模版外自创脚手架；部署

### 5 — 确认 MVP（人门禁）

Agent（或人）执行：

```bash
cd ~/projects/<name>   # 产品仓（maker-flow new <名字>）
cp -n .env.example .env
docker compose up --build
curl -sf http://localhost:8080/health
```

核验 PRO 验收标准。失败：迭代步骤 ④，或范围不对则回到步骤 ③。  
**未批准 MUST NOT 发布。**

### 6 — 发布（Agent）

- **MUST 阅读：** `skills/deploy.md`、`prompts/06-publish.md`、`release/publish/`
- **MUST** 询问人类选用哪些发布目标（Cloudflare Pages / GitHub Pages / Vercel / VPS 网关 / 拆分）
- **MUST NOT** 让人类执行 `maker-flow deploy`（仅 VPS 路径的 Agent 内部助手）
- 按对应的 `release/publish/<target>.md` 执行
- 前置：人类已批准 MVP；按所选目标具备凭证 / 主机访问

## 角色

| 角色 | 允许步骤 |
|------|----------|
| 人 | 1、3、5（可触发 6） |
| Agent | 2、4、6（6 仅在门禁 5 之后） |

## 相关

- `docs/architecture.md`
- `docs/agent-bootstrap.md`
- `docs/getting-started.md`（人）
- `docs/i18n.md`
- `AGENTS.md`
- `skills/README.md`
- `templates/index.md`
