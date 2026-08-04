# persistence-sqlx

[English](README.md) · **简体中文**

基于 [sqlx](https://github.com/jmoiron/sqlx) 的多库持久化。支持 **SQLite / PostgreSQL / MySQL**（`DB_DRIVER` 切换）。

Tags: `db` `sqlx` `sqlite` `postgres` `mysql` `persist`

## 何时用

PRO 有表 / 需要持久化。配合 `go-api`，拷贝到 `internal/persist/`。

**禁止**单独部署本 pattern。

## 环境变量

| 变量 | 必需 | 说明 |
|------|------|------|
| `DB_DRIVER` | 否（默认 `sqlite`） | `sqlite` \| `postgres` \| `mysql` |
| `DATABASE_URL` | 是* | 完整 DSN / URL |
| `SQLITE_PATH` | sqlite 回退 | `DATABASE_URL` 为空时；默认 `/data/app.db` |

## 驱动（产品仓 blank-import）

| 驱动 | import | 名称 |
|------|--------|------|
| SQLite（纯 Go，推荐） | `_ "modernc.org/sqlite"` | `sqlite` |
| PostgreSQL | `_ "github.com/lib/pq"` 等 | `postgres` |
| MySQL | `_ "github.com/go-sql-driver/mysql"` | `mysql` |

本 pattern 只硬依赖 **sqlx**；驱动写在产品仓。

## Agent 步骤

1. 选 `go-api` + `persistence-sqlx`
2. 拷到 `internal/persist/`，改 module path
3. 在 `main` blank-import 驱动
4. 接线 Open → Migrate → Store（见 [`gin_example.md`](gin_example.md)）
5. 合并 [`compose.snippet.yml`](compose.snippet.yml)
6. 用 PRO 实体替换/扩展 notes

## 校验

```bash
cd templates/patterns/persistence-sqlx
go test ./...
```
