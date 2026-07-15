# 六步工作流

Maker Flow 的核心是 **两次人工确认**，AI 在步骤 2 和 4 出力，用户在步骤 3 和 5 把关。

```
① 用户提供需求
        │
        ▼
② AI 定 PRO ──────────► ③ 用户确认 PRO  ✓ GATE
        │
        ▼
④ AI 检索模版 · 组建 MVP
        │
        ▼
⑤ 用户确认 MVP  ✓ GATE（docker compose 验收）
        │
        ▼
⑥ 上线部署
```

## 步骤详解

### 1. 用户提供需求

在 `prompts/01-requirement.example.md` 用一句话或短段落描述想法。  
不要求技术细节，边界越清晰，PRO 越准。

### 2. AI 定 PRO

执行：

```bash
./scripts/ai-run.sh prompts/02-pro-draft.md
```

AI 遵循 `skills/pro-generation.md`，输出 **PRO（Product Requirements Outline）**，包含：

- 业务目标与范围（含不做什么）
- 业务流程
- 数据模型 / 表结构
- API 契约
- 验收标准

**此步骤不输出实现代码。**

### 3. 与用户确认 PRO

将 PRO 与用户（或你自己）对齐：改范围、砍功能、补约束。  
定稿后写入 `prompts/03-pro-confirmed.example.md`（或项目内 `pro.md`）。

未通过此卡点，**不得进入步骤 4**。

### 4. AI 检索模版 · 组建 MVP

执行：

```bash
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

AI 依次遵循：

1. `skills/template-matching.md` — 根据 PRO 从 `templates/index.md` 选模版
2. `skills/mvp-assembly.md` — 复制模版、填入业务代码、给出目录树

产出：可 `docker compose up` 的最小工程（默认建议放在 `workspace/<项目名>/`）。

### 5. 与用户确认 MVP

```bash
cd workspace/<项目名>
docker compose up --build
curl http://localhost:8080/health
```

按 PRO 中的验收标准逐项检查。不通过则回到步骤 4 迭代，或回到步骤 3 修订 PRO。

### 6. 上线部署

遵循 `skills/deploy.md`，使用 `release/` 下脚本与配置：

- `release/deploy/push-and-route.sh`
- Nginx 片段 + Cloudflare DNS

## 角色分工

| 角色 | 步骤 |
|------|------|
| 用户 | ① 需求、③ 确认 PRO、⑤ 确认 MVP、⑥ 触发部署 |
| AI | ② PRO、④ 模版检索与组装 |
| 技能库 | 约束 AI 各阶段输出格式与动作 |
| 模版集 | 步骤 4 的预制零件 |

## 相关文档

- [架构说明](architecture.md)
- [快速开始](getting-started.md)
- [技能库](../skills/README.md)
- [模版目录](../templates/index.md)
