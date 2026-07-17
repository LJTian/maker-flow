[English](README.md) · **简体中文**

# Cloudflare（步骤 ⑥）

DNS/SSL 检查清单。完整部署 SOP：`skills/deploy.md`。

## 必需状态

- Zone 已在 Cloudflare 激活（NS 已指向）
- SSL/TLS 模式：`Full` 或 `Full (strict)`
- 子域名 A/CNAME → 服务器 IP，且为 **Proxied**

## 每个 MVP 的 DNS

| Type | Name | Content | Proxy |
|------|------|---------|-------|
| A | `ideaN` | server public IP | Proxied |

部署前在子域名登记表中登记名称 + `MVP_NAME`（以及可选的 `CONTAINER_PORT` / 服务名）。

## 可选自动化

环境变量：`CLOUDFLARE_API_TOKEN`、`CLOUDFLARE_ZONE_ID`（Zone DNS Edit）。可用时优先 API 而非控制台。

## 验证

```bash
curl -sfI "https://ideaN.your-domain.com/health"
```
