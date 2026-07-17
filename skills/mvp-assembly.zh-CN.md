# MVP 组装技能

[English](mvp-assembly.md) · **简体中文**

**适用步骤：** ④ AI 根据 PRO 与选定模版组建 MVP  
**前置：** `skills/template-matching.md` 已产出选定模版

## 目标

产出**可本地运行**的最小工程，不含 PRO 中未要求的功能。

## 输出目录约定

**推荐（消费侧 / 产品仓）：** 写到**当前产品仓根目录**（由 `maker-flow new` / `init` 创建）。见 `docs/consumer-project.md` 与 `AGENTS.consumer.example.md`。

**仅工厂冒烟：**

```
workspace/<项目名>/
```

`<项目名>` 由 PRO 摘要推导（kebab-case，如 `todo-api`）。

若工作树中的 `AGENTS.md` 配置了 `PRODUCT_NAME` / `MAKER_FLOW_ROOT`（消费侧模式），**MUST** 在该产品仓组装——**禁止**把私有 MVP 代码写进工厂 `workspace/`。

## 组装步骤

1. **拼装 Dockerfile** — 若 app 需要 Go 镜像片段，读 `templates/images/index.md`，内联 `go-builder` / `go-runtime`（或直接沿用 app 模版里已拼装好的 Dockerfile）。  
   仅使用上游镜像（`golang:…`、`alpine:…`）。**禁止** `FROM maker-flow/*` 或预构建本地 tag。  
   **禁止**把整个 `templates/images/` 树拷进产品——只把片段行写进产品 Dockerfile。
2. **复制模版** — 将每个选定的 `templates/apps/<id>/` 复制到输出根：
   - 单 app：产品仓根（或冒烟用 `workspace/<项目名>/`）
   - 多 app：`<输出>/<id>/`（如 `api/`、`worker/`、`cli/`），或 PRO 约定的布局
3. **合并 patterns（可选）** — 按检索结果，将 pattern 包拷入**需要它的那个 app** 目录下的 `internal/...` 并接线
4. **修改配置** — 各 app 的 `.env.example` → `.env`；多 app 时端口 / 名称勿冲突
5. **实现业务** — 按 PRO 与各 app 技术栈（Gin / Cobra / worker）分别实现
6. **更新 compose** — 多服务可在项目根用 compose 编排多个 build context，或各 app 独立 compose
7. **自检清单** — 对照 PRO 验收标准（覆盖所有选定 app）

## 代码原则

- 沿用各选定 app 模版的中间件 / 日志 / 入口风格
- 多 app 时明确进程边界与通信方式（HTTP / 队列占位等），避免糊成单进程
- 合并 pattern 时保持包名清晰、可测
- 不引入 PRO 未列出的依赖
- 改动尽量小，能跑优先
- **依赖与编译在容器内完成**（各 app `docker compose up --build`）；不要求本机 Go 工具链

## 输出格式

1. 目录树（仅新增/修改的文件）
2. 每个关键文件的**完整内容**或清晰 diff 块
3. 本地运行命令：

```bash
# 产品仓（推荐）或 workspace/<项目名>（冒烟）
cp .env.example .env
docker compose up --build
```

## 禁止

- 不把 `templates/patterns/` 单独部署为公网服务
- 不重写模版基础设施代码；不修改 `templates/images/` 片段（除非任务明确要求）
- 不在应用 Dockerfile 中重复安装片段已提供的包（ca-certificates 等）
- 不输出 PRO 范围外的功能
- 不在此步骤执行部署（留给步骤 ⑥）
