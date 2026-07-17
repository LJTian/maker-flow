# Agent 启动

[English](agent-bootstrap.md) · **简体中文**

开启 Agent 会话时的加载顺序。  
人快速开始：[getting-started.zh-CN.md](getting-started.zh-CN.md) · 仓库介绍：[README.zh-CN.md](../README.zh-CN.md) · 消费侧：[consumer-project.zh-CN.md](consumer-project.zh-CN.md) · 国际化：[i18n.zh-CN.md](i18n.zh-CN.md)

## 1. 定向

1. `AGENTS.md` — 读者、硬规则、布局、语言规则  
   （在**产品仓**中，以本地 `AGENTS.md`（来自 `AGENTS.consumer.example.md`）为准；工厂在 `$MAKER_FLOW_ROOT`。）
2. `docs/workflow.md` — 状态机 + 门禁
3. 根据对话或产物判断 **当前步骤**：
   - 无 PRO → 步骤 ②
   - 有 PRO 草稿、未确认 → 停在步骤 ③
   - 已确认 PRO、尚无组装工程 → 步骤 ④
   - 有 MVP、未批准 → 停在步骤 ⑤
   - 已批准 MVP → 步骤 ⑥

## 2. 执行当前步骤

| 步 | 先读 | 然后 |
|----|------|------|
| 2 | `skills/pro-generation.md` | `prompts/02-pro-draft.md` + `prompts/pro.template.md`（样板：`pro.example.md`） |
| 4 | `template-matching.md` → `templates/CATALOG.md` → apps + patterns → `mvp-assembly.md` | **产品仓根目录**（`maker-flow new <名字>`） |
| 6 | `skills/deploy.md` | `maker-flow deploy`（封装 `release/deploy/push-and-route.sh`） |

可选 LLM 配置说明：`ai-engine/`（若宿主 Agent 本身就是 LLM 则不必）。

契约请只读 **英文** 主版。

## 3. 编码前预检

- [ ] 步骤 ④ 及之后已有确认 PRO
- [ ] 已通过 `templates/index.md` 选定模版
- [ ] 目标路径为**产品仓**（`maker-flow new <名字>` 或已有产品仓 checkout）
- [ ] 范围符合 PRO「不做」清单

## 4. 冒烟模版（可选）

仅用于验证本机 Docker，不能替代步骤 ④：

```bash
mkdir -p /tmp/maker-flow-smoke && cp -r templates/apps/go-api /tmp/maker-flow-smoke/_smoke
cd /tmp/maker-flow-smoke/_smoke && cp .env.example .env && docker compose up --build
```
