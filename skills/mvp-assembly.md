# MVP 组装技能

**适用步骤：** ④ AI 根据 PRO 与选定模版组建 MVP  
**前置：** `skills/template-matching.md` 已产出选定模版

## 目标

产出**可本地运行**的最小工程，不含 PRO 中未要求的功能。

## 输出目录约定

```
workspace/<项目名>/
```

`<项目名>` 由 PRO 摘要推导（kebab-case，如 `todo-api`）。

## 组装步骤

1. **构建镜像基座** — 若应用依赖 `templates/images/`（如 `go-api`），先执行：
   ```bash
   ./scripts/build-images.sh
   ```
   确认 `docker images` 可见 `maker-flow/go-builder:1.22` 与 `maker-flow/go-runtime:1.22`。  
   **MUST NOT** 把基座 Dockerfile 复制进 `workspace/`；应用只通过 `FROM` 继承。
2. **复制模版** — 将 `templates/<模版ID>/` 复制到 `workspace/<项目名>/`（列出需复制的路径）
3. **修改配置** — `.env.example` → `.env`，`APP_NAME`、端口等
4. **实现业务** — 按 PRO 的 API 契约与数据模型（Gin）：
   - 路由注册（`internal/server/server.go`，`r.GET` / `Group`）
   - handler（`internal/handler/`，`func(c *gin.Context)`）
   - request/response struct + `ShouldBindJSON`（按需）
   - 可选 SQL migration 或 init 脚本
5. **更新 compose** — 若 PRO 需要 DB，取消 `docker-compose.yml` 中 postgres 注释并接线
6. **自检清单** — 对照 PRO 验收标准，标明哪些已覆盖

## 代码原则

- 沿用模版已有中间件（CORS、日志、异常恢复）与 **Gin** handler 风格
- 不引入 PRO 未列出的依赖
- 改动尽量小，能跑优先
- **依赖与编译在容器内完成**（`docker compose up --build`）；不要求本机 `go mod tidy` / `go build`

## 输出格式

1. 目录树（仅新增/修改的文件）
2. 每个关键文件的**完整内容**或清晰 diff 块
3. 本地运行命令：

```bash
./scripts/build-images.sh   # once per machine / after image template change
cd workspace/<项目名>
cp .env.example .env
docker compose up --build
```

## 禁止

- 不重写模版基础设施代码；不修改 `templates/images/` 基座（除非任务明确要求）
- 不在应用 Dockerfile 中重复安装基座已提供的包（ca-certificates 等）
- 不输出 PRO 范围外的功能
- 不在此步骤执行部署（留给步骤 ⑥）
