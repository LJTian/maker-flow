[English](subdomain-registry.example.md) · **简体中文**

# 子域名登记（示例）

复制本文件为 `subdomain-registry.md` 并维护。

| 子域名 | MVP_NAME | CONTAINER_PORT | Service | 状态 | 备注 |
|--------|----------|----------------|---------|------|------|
| idea1.your-domain.com | idea1 | 8080 | api | 在线 | Go API |
| idea2.your-domain.com | idea2 | 80 | web | 在线 | web-vite |
| free.your-domain.com | — | — | — | 空闲 | — |

**规则**：先登记再部署，避免两项目抢同一 `MVP_NAME`（Docker 网络别名）或同一子域名。

生产流量走网关宿主机 **80** + Docker 网络别名；宿主机 `8080–8090` 映射仅可选本地调试用。
