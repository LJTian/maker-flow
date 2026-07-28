# Cloudflare DNS API（脚本参考）

[English](dns-api.md) · **简体中文**

`release/cloudflare/dns.sh` 的 CLI 与原始 API 参考。**Agent SOP：** [`skills/cloudflare-dns.md`](../../skills/cloudflare-dns.md)。

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

- Pages：[`skills/publish-cloudflare-pages.md`](../../skills/publish-cloudflare-pages.md)
- VPS：[`skills/publish-vps-gateway.md`](../../skills/publish-vps-gateway.md)

## 相关

- [`README.zh-CN.md`](README.zh-CN.md)
