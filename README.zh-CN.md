<div align="center">

# Maker Flow

### 个人专属 MVP 极速孵化流水线

**重基建，轻逻辑。** 把重复的配置消灭掉，让「想法 → 公网验证」的摩擦力降到最低。

你负责 **提需求** 和 **两次确认**，AI 智能体按预制技能库与模版集执行。

<br/>

[English](README.md) · **简体中文**

<br/>

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Workflow](https://img.shields.io/badge/steps-6%20step-purple.svg)](#六步流水线)
[![Stack](https://img.shields.io/badge/template-Go%20Gin-00ADD8.svg)](templates/apps/go-api/)
[![Agent](https://img.shields.io/badge/for-AI%20Agents-22c55e.svg)](AGENTS.md)

<br/>

[快速开始](docs/getting-started.zh-CN.md) · [模版检索](templates/CATALOG.md) · [技能检索](skills/CATALOG.md) · [给 Agent 看](AGENTS.zh-CN.md) · [文档索引](docs/README.zh-CN.md)

</div>

---

## 为什么需要它

独立开发者最大的摩擦，往往不在写业务代码，而在**每次从零配环境**：

| 没有 Maker Flow | 有 Maker Flow |
|-----------------|---------------|
| 每个点子重新选框架、写 Docker、配 Nginx | 模版集开箱即用 |
| AI 一口气生成，方向错了才发现 | **两次确认**：PRO → MVP |
| Prompt 和部署流程靠脑子记 | 技能库写死 SOP，Agent 照章办事 |
| 想法多，基建重复劳动 | 专注验证，10 分钟级公网上线 |

> **这不是某个具体产品**，而是一套可 Fork、可 Star、可反复复用的 **MVP 工厂**。

---

## 六步流水线

```mermaid
flowchart LR
    A["① 你：一句话需求"] --> B["② AI：出 PRO"]
    B --> C{"③ 确认 PRO"}
    C -->|通过| D["④ AI：选模版 · 组装 MVP"]
    C -->|修改| B
    D --> E{"⑤ 确认 MVP"}
    E -->|通过| F["⑥ 部署上线"]
    E -->|迭代| D
    F --> G["🌐 公网流量验证"]
```

| 步 | 你 | AI 智能体 |
|:--:|-----|-----------|
| 1 | 提供需求 | — |
| 2 | — | 输出 PRO（方案，**不写代码**） |
| 3 | **确认 PRO** | 等待 |
| 4 | — | 检索模版 → 组装到 `workspace/` |
| 5 | **本地验收** | 按 PRO 修改 |
| 6 | 触发部署 | 执行 `release/` 脚本 |

两次确认是核心设计：**先对齐「做什么」，再动手「怎么做」。**

---

## 仓库里有什么

```
        ┌─────────────┐
        │   你 + AI   │
        └──────┬──────┘
               │
    ┌──────────┼──────────┐
    ▼          ▼          ▼
 skills/   templates/   release/
 技能库      模版集       发布基建
  (HOW)      (WHAT)      (SHIP)
```

| 模块 | 目录 | 一句话 |
|------|------|--------|
| 技能库 | [`skills/`](skills/) · [**检索目录**](skills/CATALOG.md) | 约束 Agent：PRO 怎么写、模版怎么选、怎么部署 |
| 模版集 | [`templates/`](templates/) · [**检索目录**](templates/CATALOG.md) | apps + images + patterns |
| AI 连接 | [`ai-engine/`](ai-engine/) | 可选：Ollama / OpenAI 等兼容 API |
| 发布基建 | [`release/`](release/) | Nginx + Cloudflare + 一键部署脚本 |
| 工作区 | [`workspace/`](workspace/) | Agent 组装出的 MVP 落在这里 |

---

## 60 秒上手

```bash
git clone https://github.com/LJTian/maker-flow.git && cd maker-flow
```

**方式 A — Cursor Agent（推荐）**

1. 用 Cursor 打开本仓库
2. 新建对话，输入：

   > 按 @AGENTS.md 和 docs/workflow.md，我要做一个 [你的想法]，从步骤 ① 开始。

3. 在步骤 ③、⑤ 确认 PRO 和 MVP

**方式 B — 命令行**

```bash
cp ai-engine/.env.example ai-engine/.env   # 配置 AI_BASE_URL
chmod +x scripts/ai-run.sh
./scripts/ai-run.sh prompts/02-pro-draft.md
```

**方式 C — 先验模版（无需 AI）**

```bash
./scripts/build-images.sh   # 构建 Go 基座镜像（首次必做）
cp -r templates/apps/go-api workspace/smoke-test
cd workspace/smoke-test && cp .env.example .env
docker compose up --build
curl http://localhost:8080/health
```

完整教程 → **[docs/getting-started.zh-CN.md](docs/getting-started.zh-CN.md)**

---

## 适合谁

- 经常有**天马行空的想法**，想快速丢到公网看反馈
- 已经用 **Cursor / Claude** 等 Agent，但厌倦了每次从零 prompt
- 想要一套**可 Fork 的私人流水线**，而不是又一个 Todo Demo
- 认同 **重基建、轻逻辑**：基建写一次，点子跑 N 次

---

## Star / Fork 之后

| 动作 | 建议 |
|------|------|
| Star | 跟踪更新，有新的技能库 / 模版会推到这里 |
| Fork | 改成自己的流水线；替换 `release/` 里的域名与服务器 |
| 每个新点子 | Agent 组装到 `workspace/<名字>/`，互不干扰 |
| 固定 Agent 行为 | 把 [AGENTS.md](AGENTS.zh-CN.md) 加入 IDE 规则，或对话 `@AGENTS.md` |

---

## 推荐设备分工

| 设备 | 角色 |
|------|------|
| GPU 机（可选） | 纯推理节点，主力机通过 `AI_BASE_URL` 远程调用 |
| M 系 Mac | 开发、验收、`workspace/` |
| 云服务器 | Nginx 网关，端口池 `8080–8090` 挂多个 MVP |

---

## 文档导航

| 给人看 | 给 AI 智能体看 |
|--------|----------------|
| [快速开始](docs/getting-started.zh-CN.md) | [AGENTS.md](AGENTS.zh-CN.md) |
| [架构图解](docs/overview.zh-CN.md) | [workflow.md](docs/workflow.md) |
| [模版检索](templates/CATALOG.md) · [技能检索](skills/CATALOG.md) | [agent-bootstrap.md](docs/agent-bootstrap.md) |
| [文档索引](docs/README.zh-CN.md) · [国际化](docs/i18n.zh-CN.md) | |

---

## 致谢 · 开源依赖

Maker Flow 站在这些优秀项目之上，感谢维护者与社区：

| 用途 | 项目 | 地址 |
|------|------|------|
| Go Web 框架（`go-api`） | **Gin** | https://github.com/gin-gonic/gin |
| Web UI（`web-vite`） | **Vite** · **React** · **Tailwind CSS** | https://vite.dev/ · https://react.dev/ · https://tailwindcss.com/ |
| CLI 框架（`go-cli`） | **Cobra** | https://github.com/spf13/cobra |
| singleflight（pattern） | **golang.org/x/sync** | https://pkg.go.dev/golang.org/x/sync/singleflight |
| 编译基座镜像 | **golang** (official image) | https://hub.docker.com/_/golang |
| 运行基座镜像 | **Alpine Linux** | https://alpinelinux.org/ · https://hub.docker.com/_/alpine |
| 容器运行时 | **Docker** | https://www.docker.com/ |
| 反向代理 | **Nginx** | https://nginx.org/ |
| DNS / 边缘 SSL（发布层） | **Cloudflare** | https://www.cloudflare.com/ |
| 可选本地推理 | **Ollama** | https://ollama.com/ · https://github.com/ollama/ollama |
| Agent 协议参考 | **Model Context Protocol** | https://modelcontextprotocol.io/ |

若遗漏了你的项目，欢迎提 Issue / PR，我们会补上致谢。

---

<div align="center">

**如果这套流水线对你有用，欢迎 Star — 这是对「重基建」最好的认可。**

<br/>

[开始使用](docs/getting-started.zh-CN.md) · [Fork 成自己的工厂](https://github.com/LJTian/maker-flow/fork)

</div>

## License

MIT · 个人使用，按需调整。
