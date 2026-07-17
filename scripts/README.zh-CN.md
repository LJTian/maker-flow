# scripts/

[English](README.md) · **简体中文**

| 脚本 / CLI | 用途 |
|------------|------|
| **`install.sh`** | 安装工厂到 `~/.maker-flow`，链接 CLI，提示/写入 PATH |
| **`maker-flow`** | CLI：`new`、`install`、`upgrade`、`init`、`deploy`、`root`、`version` |

## 快速开始

```bash
# 远程安装（无需 git clone）
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash

# 新建 MVP → ~/projects/my-todo
maker-flow new my-todo
cd ~/projects/my-todo
```

## 命令

```bash
maker-flow new <名字>              # 缺工厂则先安装 + init（推荐）
maker-flow init <名字>             # 仅建产品仓
maker-flow init <名字> --with-pro  # 同时生成 pro.md 草稿
maker-flow deploy --domain … --host …   # 在产品仓根目录执行
maker-flow upgrade                 # 更新 ~/.maker-flow
maker-flow root
maker-flow version
```

已 git clone 时：`./scripts/install.sh` 然后 `maker-flow new …`

指南：[`../docs/consumer-project.zh-CN.md`](../docs/consumer-project.zh-CN.md)

部署封装 `../release/deploy/push-and-route.sh`（共享网络 `maker-flow` 上的 Docker 网关）。  
Dockerfile 片段：`../templates/images/`（拼装时内联；无需预构建）。
