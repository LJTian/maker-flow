# persistence-d1

[English](README.md) · **简体中文**

解耦的数据库持久化模式组件：提供统一的 **`DB` 接口抽象层**，并由 **`D1Driver`**（Cloudflare D1 REST API）与 **`LocalSQLDriver`**（Docker / 本地 SQL 数据库）提供具体实现。

标签: `db` `d1` `cloudflare` `sqlite` `persist` `docker`

## 解耦架构设计

```
                 ┌──────────────────┐
                 │   DB 接口抽象层   │
                 └────────┬─────────┘
                          │
          ┌───────────────┴───────────────┐
          ▼                               ▼
  ┌───────────────┐               ┌───────────────┐
  │   D1Driver    │               │LocalSQLDriver │
  │(Cloudflare D1)│               │(Docker / SQL) │
  └───────────────┘               └───────────────┘
```

## 使用示例

```go
import "your_app/internal/persistd1"

// 从环境变量读取配置 (DB_MODE=local 或 d1)
cfg := persistd1.ConfigFromEnv()

// 工厂构造函数返回 DB 接口对象
db, err := persistd1.NewDB(cfg)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// 通过统一的 DB 接口执行 SQL 读写
res, err := db.ExecQuery(ctx, "SELECT * FROM users WHERE id = ?", userID)
```
