# 架构（Agent 映射）

[English](architecture.md) · **简体中文**

## 原则

- 模版 + 技能才是产品；Agent 只填业务逻辑。
- 两道人门禁：写代码前确认 PRO；部署前确认 MVP。
- Agent 按约定路径写产物，不要散落文件。
- 英文主版文档为权威；见 `docs/i18n.md`。
- **推荐布局：** 工厂（`~/.maker-flow` 或本仓）只读 skills/templates；每个 MVP 在**独立产品仓**（`maker-flow new <名字>`）。

## 组件关系

```mermaid
flowchart TB
    Human[人]
    Agent[Agent]
    Skills[skills/]
    Templates[templates/]
    Product[产品仓]
    Release[release/]

    Human -->|① 需求| Agent
    Agent -->|② PRO| Human
    Human -->|③ 确认 PRO| Agent
    Agent --> Skills
    Agent --> Templates
    Agent -->|④ 组装| Product
    Product -->|⑤ 确认 MVP| Human
    Product -->|⑥ 部署| Release
    Release --> MVP[公网 MVP]
```

## 目录

| 路径 | Agent 用途 |
|------|------------|
| `skills/` | 各步骤权威 HOW |
| `templates/` | 可检索脚手架；目录 = `index.md` |
| `prompts/` | 分阶段输入模版 |
| **产品仓** | 步骤 ④–⑥ **唯一**写出目标（见 `docs/consumer-project.md`；`maker-flow new <名字>`） |
| `release/` | 步骤 ⑥ 部署原语（网关 + `push-and-route`；优先 `maker-flow deploy`） |
| `scripts/` | `install.sh`、`maker-flow` CLI |
| `ai-engine/` | 可选远程 LLM 连接说明（宿主 Agent 场景通常不用） |
| `docs/` | 流程 / 架构契约 |

## 步骤 → 路径

| 步 | 角色 | 路径 |
|----|------|------|
| 1 | 人 | 需求文本 |
| 2 | Agent | `skills/pro-generation.md`、`prompts/02-pro-draft.md` |
| 3 | 人 | 定稿 PRO（产品仓 `pro.md`，或工厂示例） |
| 4 | Agent | `template-matching` + `mvp-assembly` + `templates/` → **产品仓根目录** |
| 5 | 人 | 组装工程 + PRO 验收 |
| 6 | Agent | `skills/deploy.md`、`maker-flow deploy` / `release/` |

## 相关

- `docs/workflow.md`
- `docs/consumer-project.md`
- `docs/agent-bootstrap.md`
- `docs/getting-started.md`（人）
- `docs/i18n.md`
- `AGENTS.md` · `AGENTS.consumer.example.md`
