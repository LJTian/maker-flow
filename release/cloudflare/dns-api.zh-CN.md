# Cloudflare DNS API（Agent）

[English](dns-api.md) · **简体中文**

通过 Cloudflare API 对 DNS 做 **增删改查**。

**Agent 技能：** [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md) — 在 Cloudflare 上查改 DNS 时加载。

## 人类一次性准备

1. 注册：[https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up)
2. 域名 NS 已指向 Cloudflare。
3. 创建 Token：[API Tokens](https://dash.cloudflare.com/profile/api-tokens)。
   - 最小权限建议：
     - `Account → Cloudflare Pages → Edit`（发布静态页面）
     - `Zone → DNS → Edit`（增删改查 DNS，含目标 Zone）
4. 交给 Agent：
   - `CLOUDFLARE_API_TOKEN`（**必需**）
   - `CLOUDFLARE_ZONE_ID`（**可选**，脚本可基于 token 自动发现/选择）
   - `CLOUDFLARE_ACCOUNT_ID`（**可选**，用于限定发现范围；部分 Pages API 会用到）

**禁止**编造或打印 Token。

## CLI：`dns.sh`

路径：`$(maker-flow root)/release/cloudflare/dns.sh`

```bash
export CLOUDFLARE_API_TOKEN=…
DNS="$(maker-flow root)/release/cloudflare/dns.sh"

# 可选：预先指定
# export CLOUDFLARE_ACCOUNT_ID=…
# export CLOUDFLARE_ZONE_ID=…

# 查
"$DNS" list
"$DNS" list --type A --name api.example.com
"$DNS" list --json
"$DNS" accounts
"$DNS" zones
"$DNS" --zone-name example.com list
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

当 token 下有多个账号/Zone 时，`dns.sh` 支持：
- 交互式选择（TTY）
- 或显式 `--account-id` / `--zone-id` / `--zone-name`
- 或 `--non-interactive` 在自动化里快速失败

## 发布流程

- Pages：[`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md)
- VPS：[`../publish/vps-gateway.md`](../publish/vps-gateway.md)

## 相关

- [`README.zh-CN.md`](README.zh-CN.md)
