# Maker Flow

**个人专属 MVP 极速孵化流水线** — 重基建，轻逻辑。

AI 负责出方案与组装，模版与流程预制好；用户在两个卡点确认，避免方向跑偏。

## 六步流程

| 步 | 执行者 | 动作 | 文档 |
|----|--------|------|------|
| 1 | 用户 | 提供需求 | [prompts/01-requirement.example.md](prompts/01-requirement.example.md) |
| 2 | AI | 根据需求定出 **PRO** | [prompts/02-pro-draft.md](prompts/02-pro-draft.md) + [skills/pro-generation.md](skills/pro-generation.md) |
| 3 | 用户 ↔ AI | **确认 PRO** | [prompts/03-pro-confirmed.example.md](prompts/03-pro-confirmed.example.md) |
| 4 | AI | 检索模版，组建 **MVP** | [prompts/04-assemble-mvp.md](prompts/04-assemble-mvp.md) + [skills/](skills/) |
| 5 | 用户 ↔ AI | **确认 MVP** | 本地 `docker compose` 验收 |
| 6 | — | **上线部署** | [skills/deploy.md](skills/deploy.md) + [release/](release/) |

完整说明：[docs/workflow.md](docs/workflow.md)

## 仓库结构

```
maker-flow/
├── ai-engine/       # AI 连接配置（.env、providers、参数）
├── skills/          # 技能库：各阶段 SOP，约束 AI 怎么做
├── templates/       # 模版集：镜像、demo 源码、文档（供 AI 检索）
├── prompts/         # 分阶段 Prompt（PRO 与组装分开）
├── release/         # 部署基建（Nginx、Cloudflare、脚本）
├── scripts/         # ai-run.sh 等工具
└── docs/
```

## 快速开始

```bash
# 1. 配置 AI 连接
cp ai-engine/.env.example ai-engine/.env

# 2. 填写需求，生成 PRO（步骤 1-2）
# 编辑 prompts/01-requirement.example.md 与 prompts/02-pro-draft.md
./scripts/ai-run.sh prompts/02-pro-draft.md

# 3. 与 AI 迭代直到 PRO 确认，写入 prompts/03-pro-confirmed.example.md

# 4. 组装 MVP（步骤 4）
./scripts/ai-run.sh prompts/04-assemble-mvp.md

# 5. 本地验证
cd workspace/my-mvp   # 或 AI 指定的输出目录
docker compose up --build

# 6. 部署见 release/ 与 skills/deploy.md
```

## 三大块职责

| 块 | 目录 | 职责 |
|----|------|------|
| **AI 连接** | `ai-engine/` | 连什么模型、参数边界 |
| **技能库** | `skills/` | PRO 怎么写、模版怎么选、怎么部署 |
| **模版集** | `templates/` | 可检索的预制工程（含 `index.md` 目录） |

## License

Private / 个人使用。
