# 消费侧项目（并列目录）

[English](consumer-project.md) · **简体中文**

当 **maker-flow 是公开工具仓**，每个 MVP 在 **独立产品仓库**（常为私有）中开发时，使用本文。

推荐布局：**并列目录** — 不用 submodule，也不把整棵工厂拷进产品仓。

---

## 两个仓库，两种角色

| 仓库 | 可见性 | 内容 |
|------|--------|------|
| **maker-flow** | 公开（工具） | `skills/`、`templates/`、`prompts/`、`release/`、文档 |
| **你的 MVP**（如 `my-todo`） | 私有（产品） | `pro.md`、组装后的 app、`docker-compose.yml`、短 `AGENTS.md` |

**规则：** 不要把 MVP 业务代码提交进 maker-flow。别人 clone 工厂时不应拿到你的想法。真实 MVP 放在产品仓（`maker-flow new <名字>`）。

---

## 推荐目录布局

```text
~/.maker-flow/                # 工厂（maker-flow install）
~/projects/
  ├── my-todo/                # 私有 git — 一个点子一个仓
  ├── my-saas/
  └── ...
```

克隆 maker-flow 仓库仅供**贡献者**；使用者请 `curl … | bash` 或 `./scripts/install.sh`。

### 产品仓最小结构

```text
my-todo/
├── AGENTS.md                 # maker-flow init / new 生成
├── requirement.md            # 步骤 ①
├── pro.md                    # 步骤 ③（可选；maker-flow init --with-pro）
├── docker-compose.yml        # 组装后（多 app 用根 compose）
├── api/
└── web/                      # 可选
```

**不要**整棵拷贝进产品仓：`skills/`、`templates/`、`release/`（通过 `MAKER_FLOW_ROOT` 只读引用）。

---

## 环境变量

通常 **不必配置** — 产品仓 `AGENTS.md` 默认 `MAKER_FLOW_ROOT=~/.maker-flow`（可移植）。用 `maker-flow root` 解析。

仅自定义安装目录时覆盖：

```bash
export MAKER_FLOW_ROOT=/custom/path/to/maker-flow
```

Agent 约定：

- **只读：** `$MAKER_FLOW_ROOT/skills/`、`templates/`、`prompts/`、`release/`
- **写入：** 仅当前产品仓（`./`、`./api/`、`./web/` …）
- **禁止** 修改 `$MAKER_FLOW_ROOT` 下任何文件

---

## 新建 MVP 流程

### 一条命令（推荐）

```bash
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash
maker-flow new my-todo
cd ~/projects/my-todo
```

### 分步

```bash
./scripts/install.sh
maker-flow init my-todo              # 或 maker-flow init my-todo --with-pro
cd ~/projects/my-todo
```

### 手动（无 CLI）

```bash
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash
maker-flow init my-todo
cd ~/projects/my-todo
```

---

## 六步映射（消费侧）

| 步 | 发生位置 |
|----|----------|
| ① 需求 | 产品仓（对话或本地笔记） |
| ② 起草 PRO | Agent 读 `$MAKER_FLOW_ROOT/skills/pro-generation.md` |
| ③ 确认 PRO | 写入产品仓 **`pro.md`** |
| ④ 组装 | 从 `$MAKER_FLOW_ROOT/templates/...` **拷贝**到产品仓；输出目录 = 产品根 |
| ⑤ 验收 | 在产品仓 `docker compose up --build` |
| ⑥ 部署 | 产品仓根目录执行 `maker-flow deploy --domain … --host …` |

Go 模版：从 `$MAKER_FLOW_ROOT/templates/images/` 拼装 Dockerfile 片段（app 模版已内联）。无需预构建步骤。

---

## 多 app 布局

```text
my-todo/
├── pro.md
├── docker-compose.yml
├── api/
└── web/
```

与 `go-api` 联调时，构建 `web/` 前设置 `VITE_API_BASE_URL`。

---

## 方式对比（为何用并列目录）

| 方式 | 产品仓干净度 | 版本钉死 | 日常摩擦 |
|------|--------------|----------|----------|
| **并列目录（推荐）** | 最好 | 手动记工厂 commit | 低 |
| Submodule `vendor/maker-flow` | 较好 | submodule SHA | 中 |
| 整棵工厂拷进产品 | 差 | 冻结副本 | 低 |

仅当某个产品需要协作或 CI 钉死工厂版本时，再对该产品仓使用 **submodule**。

---

## 安全清单

- [ ] 需保密的产品仓设为 **private**
- [ ] maker-flow 不提交含密钥的 `.env`
- [ ] 真实子域/登记数据不进 maker-flow 的 `release/`
- [ ] 含敏感范围的 `pro.md` 只留在产品仓

---

## 相关

- 产品仓 `AGENTS.md` 模版：[`AGENTS.consumer.example.md`](../AGENTS.consumer.example.md)
- 工厂 Agent 入口：[`AGENTS.md`](../AGENTS.md)
- 流程：[`workflow.md`](workflow.md)
- 快速开始（工厂本仓）：[`getting-started.md`](getting-started.zh-CN.md)
