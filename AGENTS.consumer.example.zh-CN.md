# AGENTS.md — 消费侧产品仓

[English](AGENTS.consumer.example.md) · **简体中文**

**将本文件复制到你的 MVP 产品仓库，重命名为 `AGENTS.md`。**  
不要把它当作 maker-flow 工厂仓内的产品契约提交。

工厂（只读）：把 `MAKER_FLOW_ROOT` 设为本机 maker-flow 克隆路径。  
指南：[`docs/consumer-project.zh-CN.md`](docs/consumer-project.zh-CN.md)

---

## 上下文

- **本仓库** = 单个 MVP 产品（仅业务代码）。
- **maker-flow** = 共享工厂，路径 `$MAKER_FLOW_ROOT`（skills、templates、release）。
- **禁止** 将整棵 `skills/` / `templates/` / `release/` 拷入本仓。
- **禁止** 修改 `$MAKER_FLOW_ROOT` 下任何文件。

## 配置（按项目修改）

```bash
# 必填 — maker-flow 克隆的绝对路径
MAKER_FLOW_ROOT=/path/to/maker-flow

# 本产品名（kebab-case）
PRODUCT_NAME=my-todo
```

## 输出路径（覆盖工厂默认 `workspace/`）

| 产物 | 本仓路径 |
|------|----------|
| 定稿 PRO | `pro.md` |
| 单 app | `./` 或 `./<app-id>/` |
| 多 app | `./api/`、`./web/` … 按 PRO |
| 部署工作目录 | 本仓根目录 |

## Agent 入口

1. 读 `$MAKER_FLOW_ROOT/docs/workflow.md`（门禁不变）。
2. 从 `$MAKER_FLOW_ROOT/skills/CATALOG.md` 加载当前步骤技能。
3. 经 `$MAKER_FLOW_ROOT/templates/CATALOG.md` 匹配模版。
4. **只在本产品仓写入** — 不要写 `$MAKER_FLOW_ROOT/workspace/`。

## 六步状态机

| 步 | 动作 | 读取 | 写入 |
|----|------|------|------|
| 1 | 人：需求 | — | 对话 / 笔记 |
| 2 | Agent：起草 PRO | `$MAKER_FLOW_ROOT/skills/pro-generation.md`、`prompts/pro.template.md` | 对话（无代码） |
| 3 | 人：确认 PRO | — | **`pro.md`** |
| 4 | Agent：组装 | `template-matching`、`mvp-assembly`、`templates/` | **本仓**（`./`、`./api/` …） |
| 5 | 人：验收 MVP | `pro.md` 验收标准 | — |
| 6 | Agent：部署 | `$MAKER_FLOW_ROOT/skills/deploy.md`、`release/` | 在本仓执行 `maker-flow deploy` |

硬门禁：**步骤 ③、⑤ 人类确认前必须停下。**

## 组装规则

- 从 `$MAKER_FLOW_ROOT/templates/apps/<id>/` **拷贝**到本仓。
- Patterns：拷进需要它的 app；禁止单独部署 pattern。
- Go Dockerfile：从 `$MAKER_FLOW_ROOT/templates/images/` 拼装片段（内联进 app Dockerfile；无需预构建）。
- 验收优先在本仓 `docker compose up --build`。

## 部署

在本仓根目录：

```bash
maker-flow deploy \
  --domain my-todo.your-domain.com \
  --host deploy@your-server \
  --service api \
  --port 8080
```

## 首条对话模版

```
请阅读本 AGENTS.md 与 $MAKER_FLOW_ROOT/docs/workflow.md。
MAKER_FLOW_ROOT 已配置。当前为产品仓「${PRODUCT_NAME}」。
从步骤 ① 开始，我的需求是：…
```
