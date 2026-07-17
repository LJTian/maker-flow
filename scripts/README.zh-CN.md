# scripts/

[English](README.md) · **简体中文**

| 脚本 / CLI | 用途 |
|------------|------|
| **`install.sh`** | 安装工厂到 `~/.maker-flow`，链接 `maker-flow` CLI，自动 PATH 提示 |
| **`maker-flow`** | CLI：`new`、`install`、`upgrade`、`init`、`root`、`version`（另有 Agent 内部用的 `deploy`） |
| `check.sh` | 工厂 CI / 本地自检 |

## 快速开始

```bash
# 远程（无需 git clone）
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash

# 新建 MVP → ~/projects/my-todo
maker-flow new my-todo
cd ~/projects/my-todo
```

## 给人用的命令

```bash
maker-flow new <name>              # 按需安装 + init（推荐）
maker-flow init <name>             # 仅产品仓
maker-flow init <name> --with-pro  # 同时创建草稿 pro.md
maker-flow upgrade                 # 更新 ~/.maker-flow
maker-flow root
maker-flow version
./scripts/check.sh                 # 本地 / CI 检查
```

**上线 / 发布：** 本地验收后与 Agent 对话，选择 Cloudflare Pages / GitHub Pages / Vercel / VPS。人类**不需要**跑 `maker-flow deploy`（见 `skills/deploy.md`）。

## Agent 内部

`maker-flow deploy …` 仅封装 VPS 网关路径的 `../release/deploy/push-and-route.sh`。人类确认该目标后，由 Agent 调用即可。

从 git clone：`./scripts/install.sh` 然后 `maker-flow new …`

指南：[`../docs/consumer-project.zh-CN.md`](../docs/consumer-project.zh-CN.md)  
发布目标：[`../release/publish/`](../release/publish/)  
Dockerfile 片段：`../templates/images/`（组装时内联；无需预构建）。
