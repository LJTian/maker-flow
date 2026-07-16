# AGENTS.md

[English](AGENTS.md) · **简体中文**

**读者：AI 智能体。** 给人看的介绍：[README.zh-CN.md](README.zh-CN.md) · 快速开始：[docs/getting-started.zh-CN.md](docs/getting-started.zh-CN.md) · 国际化：[docs/i18n.zh-CN.md](docs/i18n.zh-CN.md)

本仓库是个人 MVP 交付的 Agent 剧本。人类只提供需求并做门禁确认。

原则：**重基建，轻逻辑。** 优先用模版与技能，不要自造脚手架。

## 语言

- 权威契约是 **英文** `.md` 主版。
- `.zh-CN.md` 仅供人类阅读。
- **除非**用户明确要求中文面向输出，否则 **不得** 把中文副版当作步骤 / 技能 / 门禁的权威来源。

## Agent 入口

1. 读 `docs/workflow.md`（状态机 + 硬门禁）。
2. 从 `skills/` 加载当前步骤技能（先看 `skills/CATALOG.md`）。
3. 使用 `prompts/` 输入（或等价用户消息）。
4. 步骤 ④ 先打开 `templates/CATALOG.md` 再匹配。
5. 按技能规定的路径写出产物。

**不要**跳过门禁。`templates/index.md` 有匹配时 **不要** 自创技术栈。

## 六步状态机

| 步 | 角色 | 动作 | 必读 | 产出 |
|----|------|------|------|------|
| 1 | 人 | 提供需求 | — | 需求文本 |
| 2 | Agent | 起草 PRO | `skills/pro-generation.md`、`prompts/02-pro-draft.md`、`prompts/pro.template.md` | PRO Markdown（无代码） |
| 3 | 人 | 确认 PRO | — | 定稿 PRO → `prompts/03-pro-confirmed.example.md` 或项目 `pro.md`（结构同 `pro.template.md`） |
| 4 | Agent | 匹配模版并组装 MVP | `skills/template-matching.md`、`skills/mvp-assembly.md`、`templates/index.md`、`prompts/04-assemble-mvp.md` | `workspace/<name>/` |
| 5 | 人 | 确认 MVP | PRO 验收标准 | 通过 / 不通过 |
| 6 | Agent（通过后） | 部署 | `skills/deploy.md`、`release/` | 公网 URL |

硬门禁：**在步骤 ③、⑤ 人类确认前必须停下。**

## 目录布局

```
ai-engine/     # 可选 LLM 传输配置（OpenAI 兼容）
skills/        # HOW — 步骤 SOP（对 Agent 权威）
templates/     # WHAT — 可检索脚手架；从 templates/index.md 开始
prompts/       # 分阶段输入契约
workspace/     # Agent 组装 MVP 的写出目录
release/       # 部署原语（nginx、cloudflare、脚本）
scripts/       # 助手脚本（如 ai-run.sh）
docs/          # 流程与架构契约
```

## 硬规则

- 当前步骤 MUST 遵循 `skills/*`（英文主版）。
- 步骤 ② MUST NOT 输出实现代码。
- 未确认 PRO（步骤 ③）MUST NOT 组装（步骤 ④）。
- 未确认 MVP（步骤 ⑤）MUST NOT 部署（步骤 ⑥）。
- 编码前 MUST 经 `templates/CATALOG.md` / `templates/index.md` 选定 **1～N 个** app（每个 app 须对应 PRO 职责）。
- MUST 经 `templates/images/index.md` 解析镜像基座；缺失时运行 `./scripts/build-images.sh`。
- MAY 从 `templates/patterns/` 附加 0～N 个 pattern（拷进需要它的 app；不得单独部署）。
- MUST 把组装项目写到 `workspace/<kebab-name>/`（多 app：`workspace/<name>/<app-id>/`）。
- MUST NOT 把 `templates/images/` 拷进 `workspace/`；仅通过 `FROM` 继承。
- 优先 **容器构建**（`./scripts/build-images.sh` 然后 `docker compose up --build`）；验收不要求本机 Go 工具链。

## 契约索引

| 主题 | 路径 |
|------|------|
| 流程 | `docs/workflow.md` |
| 架构 | `docs/architecture.md` |
| Agent 启动 | `docs/agent-bootstrap.md` |
| 技能索引 | `skills/CATALOG.md` |
| 模版目录 | `templates/CATALOG.md` |
| Pattern 目录 | `templates/patterns/index.md` |
| 国际化 | `docs/i18n.md` |
