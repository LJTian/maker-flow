# AGENTS.md — 消费侧产品仓

[English](AGENTS.consumer.example.md) · **简体中文**

**请把本文件复制到 MVP 产品仓，并命名为 `AGENTS.md`。**  
不要把它当作 maker-flow 工厂仓内的产品契约提交。

工厂（只读）：用 `maker-flow root` 解析（默认 `~/.maker-flow`）。  
指南：[`docs/consumer-project.zh-CN.md`](docs/consumer-project.zh-CN.md)

---

## 上下文

- **本仓** = 一个 MVP 产品（只放业务代码）。
- **maker-flow** = `$MAKER_FLOW_ROOT` 处的共享工厂（skills、templates、release）。
- **禁止**把整棵 `skills/` / `templates/` / `release/` 拷进本仓。
- **禁止**修改 `$MAKER_FLOW_ROOT` 下的文件。

## 配置（按项目修改）

```bash
# 可移植默认（展开 ~）。仅自定义安装目录时覆盖。
# 运行时解析：maker-flow root
MAKER_FLOW_ROOT=~/.maker-flow

# 本产品（kebab-case）
PRODUCT_NAME=my-todo
```

## 输出路径

| 产物 | 本仓路径 |
|------|----------|
| 已确认 PRO | `pro.md` |
| 单 app | `./` 或 `./<app-id>/` |
| 多 app | `./api/`、`./web/` … 按 PRO |
| 发布工作目录 | 本仓根 |

## Agent 入口

1. 读 `$MAKER_FLOW_ROOT/docs/workflow.md`（门禁不变）。必要时展开 `MAKER_FLOW_ROOT` 中的 `~`。
2. 从 `$MAKER_FLOW_ROOT/skills/CATALOG.md` 加载步骤技能。
3. 经 `$MAKER_FLOW_ROOT/templates/CATALOG.md` 匹配模版。
4. **只在本产品仓写入** — 不要写 `$MAKER_FLOW_ROOT`。

## 六步状态机

| 步骤 | 动作 | 阅读 | 写入 |
|------|------|------|------|
| 1 | 人：需求 | — | 对话 / 笔记 |
| 2 | Agent：起草 PRO | `$MAKER_FLOW_ROOT/skills/pro-generation.md`、`prompts/pro.template.md` | 对话（无代码） |
| 3 | 人：确认 PRO | — | **`pro.md`** |
| 4 | Agent：组装 | `template-matching`、`mvp-assembly`、`templates/` | **本仓**（`./`、`./api/` …） |
| 5 | 人：验收 MVP | `pro.md` 标准 | — |
| 6 | Agent：发布 | `$MAKER_FLOW_ROOT/skills/deploy.md`、`prompts/06-publish.md`、`release/publish/` | **人类在对话中选定目标后** 得到公网 URL |

硬门禁：**在步骤 ③、⑤ 人类确认前必须停下。**  
步骤 ⑥：**先问发到哪里**（Pages / Vercel / VPS / 拆分）。**不要**让人类跑 `maker-flow deploy`（仅 Agent 内部可用）。

## 组装规则

- 从 `$MAKER_FLOW_ROOT/templates/apps/<id>/` **拷贝**到本仓。
- Patterns：拷进需要它的 app；不得单独部署 pattern。
- Go Dockerfile：从 `$MAKER_FLOW_ROOT/templates/images/` 拼装片段（内联进 app Dockerfile；无需预构建）。
- 拷贝后，把各 Go app 的 `go.mod` `module` 改成产品路径（如 `github.com/<you>/${PRODUCT_NAME}` 或 `.../${PRODUCT_NAME}/api`）——不要保留 `.../maker-flow/templates/...`。
- 验收优先在本仓 `docker compose up --build`。

## 发布

步骤 ⑤ 后，加载 `$MAKER_FLOW_ROOT/prompts/06-publish.md`，与人类确认目标，再执行 `$MAKER_FLOW_ROOT/release/publish/<target>.md`。

## 首条消息模版

```
请阅读本 AGENTS.md 与 $MAKER_FLOW_ROOT/docs/workflow.md。
MAKER_FLOW_ROOT=~/.maker-flow（或：maker-flow root）。当前为产品仓「${PRODUCT_NAME}」。
从步骤 ① 开始，我的需求是：…
```
