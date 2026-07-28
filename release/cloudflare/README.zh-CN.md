# Cloudflare（步骤 ⑥）

[English](README.md) · **简体中文**

发布用的 DNS / SSL / Token 自动化。总 SOP：`skills/deploy.md` → `release/publish/`。

## 指南

| 文档 / 脚本 | 用途 |
|-------------|------|
| [`dns-api.zh-CN.md`](dns-api.zh-CN.md) | DNS CRUD 参考（脚本参数） |
| [`dns.sh`](dns.sh) | Agent CLI — 技能：[`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md) |
| [`dns-upsert.sh`](dns-upsert.sh) | 兼容包装 → `dns.sh upsert` |
| [`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md) | Pages Direct Upload + 可选自定义域 DNS |
| [`subdomain-registry.example.zh-CN.md`](subdomain-registry.example.zh-CN.md) | MVP 子域登记表示例 |

## 人类 Token 清单

**还没有账号？** 打开 [https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)

| 变量 | 用于 |
|------|------|
| `CLOUDFLARE_API_TOKEN` | Pages + DNS |
| `CLOUDFLARE_ACCOUNT_ID` | Pages（非交互） |
| `CLOUDFLARE_ZONE_ID` | DNS upsert |

权限：**Pages Edit** 和/或 **DNS Edit**。已配置时优先 API。

## 自定义域必需状态

- Zone 已在 Cloudflare（**NS** 已指向 CF — 注册商侧一次性；API 做不了）
- 反代到 HTTPS 源站（VPS）时 SSL/TLS：`Full` 或 `Full (strict)`
- 记录与发布目标匹配：

| 目标 | 记录 | 内容 | Proxy |
|------|------|------|-------|
| VPS 网关 | A / AAAA | 服务器公网 IP | Proxied |
| Cloudflare Pages | CNAME | `<project>.pages.dev` | Proxied |

## 每个 MVP 的 DNS

若人类维护清单，发布前在登记表写入名称 + `MVP_NAME`。

## 验证

```bash
curl -sfI "https://ideaN.your-domain.com/"
curl -sfI "https://ideaN.your-domain.com/health"
```
