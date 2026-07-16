[English](subdomain-registry.example.md) · **简体中文**

# 子域名与端口登记（示例）

复制本文件为 `subdomain-registry.md` 并维护。

| 子域名 | HOST_PORT | 项目目录 | 状态 | 备注 |
|--------|-----------|----------|------|------|
| test.your-domain.com | 8080 | static-test | 在线 | 通路验证 |
| idea1.your-domain.com | 8080 | my-todo-api | 在线 | Go API |
| idea2.your-domain.com | 8081 | — | 空闲 | — |

**规则**：先登记再部署，避免两人/两项目抢同一端口。
