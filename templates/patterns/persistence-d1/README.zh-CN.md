# persistence-d1

[English](README.md) · **简体中文**

双模式持久化驱动：**本地开发**使用 Docker 容器 / 本地 SQLite 文件；**线上部署**使用 **Cloudflare D1** HTTP API。

标签: `db` `d1` `cloudflare` `sqlite` `persist` `docker`

## 开发与部署策略

- **本地开发 (`DB_MODE=local`)**：使用本地 SQLite 文件或 Docker 容器挂载卷 (`SQLITE_PATH=/data/app.db`)，无需配置任何云端凭证。
- **线上部署 (`DB_MODE=d1`)**：通过 `CLOUDFLARE_ACCOUNT_ID`、`CLOUDFLARE_D1_DATABASE_ID` 和 `CLOUDFLARE_API_TOKEN` 直接调用 Cloudflare D1 REST API v4 执行数据库读写。

## 环境变量

| 变量名 | 是否必填 | 说明 |
|--------|----------|------|
| `DB_MODE` | 否（默认 `local`） | `local` \| `d1` |
| `SQLITE_PATH` | 本地模式 | 本地文件路径，默认 `/data/app.db` |
| `CLOUDFLARE_ACCOUNT_ID` | 线上模式必填 | Cloudflare 账户 ID |
| `CLOUDFLARE_D1_DATABASE_ID` | 线上模式必填 | Cloudflare D1 数据库 ID |
| `CLOUDFLARE_API_TOKEN` | 线上模式必填 | 包含 D1 编辑权限的 Cloudflare API Token |

## 核心接口

- `ConfigFromEnv()` → 从环境变量读取配置。
- `NewClient(cfg)` → 初始化并返回 `*Client`。
- `client.ExecQuery(ctx, sql, params...)` → 统一查询接口，支持本地与 D1 REST 调用的无缝切换。

## 免费额度说明

- **Cloudflare D1 免费额度**：每天 500 万次读取，每天 10 万次写入，5 GB 存储空间。
- **费用**：免费额度内 0 元/月。
