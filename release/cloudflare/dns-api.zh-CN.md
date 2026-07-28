# Cloudflare DNS API（Agent）

[English](dns-api.md) · **简体中文**

通过 Cloudflare API 对 DNS 做 **增删改查**。

**Agent 技能：** [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md) — 在 Cloudflare 上查改 DNS 时加载。

## 人类一次性准备

1. 注册：[https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)
2. 域名 NS 已指向 Cloudflare。
3. 创建 Token：[API Tokens](https://dash.cloudflare.com/profile/api-tokens)，权限 **Zone → DNS → Edit**（含目标 Zone）。
4. 交给 Agent：
   - `CLOUDFLARE_API_TOKEN`
   - `CLOUDFLARE_ZONE_ID`
   - Pages 另需：`CLOUDFLARE_ACCOUNT_ID`

**禁止**编造或打印 Token。

## CLI：`dns.sh`

路径：`$(maker-flow root)/release/cloudflare/dns.sh`

```bash
export CLOUDFLARE_API_TOKEN=…
export CLOUDFLARE_ZONE_ID=…
DNS="$(maker-flow root)/release/cloudflare/dns.sh"

# 查
"$DNS" list
"$DNS" list --type A --name api.example.com
"$DNS" list --json
"$DNS" get --id RECORD_ID

# 增
"$DNS" create --type A --name api.example.com --content 203.0.113.10

# 改
"$DNS" update --type A --name api.example.com --content 203.0.113.11
"$DNS" update --id RECORD_ID --content 203.0.113.12

# 删
"$DNS" delete --type A --name api.example.com
"$DNS" delete --id RECORD_ID

# 有则改、无则增（发布流程常用）
"$DNS" upsert --type CNAME --name www.example.com --content my-project.pages.dev
```

`dns-upsert.sh` = `dns.sh upsert` 的兼容包装。

依赖：`curl`、`python3`。

## 发布流程

- Pages：[`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md)
- VPS：[`../publish/vps-gateway.md`](../publish/vps-gateway.md)

## 相关

- [`README.zh-CN.md`](README.zh-CN.md)
