# Cloudflare DNS 技能

[English](cloudflare-dns.md) · **简体中文**

**技能 id：** `cloudflare-dns`  
**常见步骤：** ⑥ 发布 — 需要在 Cloudflare 上查、增、改、删 DNS 时  
**也可：** 人类单独要求 DNS / DDNS，不一定完整发布

## 目标

用 API + `release/cloudflare/dns.sh` 管理 Cloudflare Zone 上的 DNS，**不必**让人类为日常 CRUD 去点控制台。

## 何时加载

在 **`skills/deploy.md` 之外** 额外加载本技能，当：

- **Cloudflare Pages** + **自定义域名**
- **VPS 网关** + 子域指向服务器（A / AAAA）
- 人类要求 **列出 / 添加 / 修改 / 删除** Cloudflare DNS
- Cloudflare 上的 DDNS（IPv4 / IPv6）

非 Cloudflare 服务商 **不要** 用本技能。

## 人类配置（对话确认）

执行 `dns.sh` 前确认：

| 项 | 必需 | 说明 |
|----|------|------|
| 账号 | 首次 | [https://dash.cloudflare.com/sign-up](https://dash.cloudflare.com/sign-up) |
| `CLOUDFLARE_API_TOKEN` | 是 | 目标 Zone：**Zone → DNS → Edit** |
| `CLOUDFLARE_ZONE_ID` | 是 | Zone 概览 → Zone ID |
| `CLOUDFLARE_ACCOUNT_ID` | 仅 Pages | 纯 DNS CRUD 不需要 |
| Zone 在 CF 上 | 是 | NS 已指向 CF（注册商侧一次性） |

**必须**用人类本机环境变量。**禁止**编造 Token 或写入对话记录。

创建 Token：[API Tokens](https://dash.cloudflare.com/profile/api-tokens) → 自定义 → Zone DNS Edit。

## Agent 工具

```bash
DNS="$(maker-flow root)/release/cloudflare/dns.sh"
export CLOUDFLARE_API_TOKEN=…
export CLOUDFLARE_ZONE_ID=…
```

| 操作 | 命令 |
|------|------|
| **查列表** | `"$DNS" list` · `"$DNS" list --type A --name host.example.com` · `"$DNS" list --json` |
| **查单条** | `"$DNS" get --id RECORD_ID` |
| **增** | `"$DNS" create --type TYPE --name NAME --content CONTENT [--proxied true\|false]` |
| **改** | `"$DNS" update --type TYPE --name NAME --content CONTENT` 或 `--id RECORD_ID --content …` |
| **删** | `"$DNS" delete --type TYPE --name NAME` 或 `--id RECORD_ID` |
| **有则改无则增** | `"$DNS" upsert --type TYPE --name NAME --content CONTENT`（发布默认） |

依赖：`curl`、`python3`。兼容：`dns-upsert.sh` = `dns.sh upsert`。

详见：[`release/cloudflare/dns-api.zh-CN.md`](../release/cloudflare/dns-api.zh-CN.md)。

## 常见记录类型

| 类型 | 用途 | `proxied` |
|------|------|-----------|
| `A` | IPv4 → VPS | 通常 `true` |
| `AAAA` | IPv6 → VPS（DDNS） | 通常 `true` |
| `CNAME` | Pages 自定义域 → `<project>.pages.dev` | 通常 `true` |

## 与发布集成

| 发布目标 | DNS 技能 |
|----------|----------|
| Pages + 自定义域 | 先 Pages Domains API，再 `dns.sh upsert` CNAME |
| VPS 网关 | `dns.sh upsert` A / AAAA |
| 仅 `*.pages.dev` | **跳过** DNS 技能 |

## 流程

1. 确认要改什么 DNS（或从发布目标推断）。
2. 确认 Token + Zone ID。
3. 不确定时先 **list**。
4. 执行 create / update / upsert / delete。
5. **验证：** list、`dig`、`curl`。

## 硬规则

- 有 Token + Zone ID 时 **必须** 用 `dns.sh`。
- 可能多条匹配时删之前先 **list**，优先 `--id` 删除。
- **不要**让人类自己跑 `dns.sh`（除非人类明确要求）— 由 Agent 执行。
- 发布相关 DNS **不得**在步骤 ⑤ 之前改。
- 人类维护登记表时写入 [`subdomain-registry.example.zh-CN.md`](../release/cloudflare/subdomain-registry.example.zh-CN.md)。

## 相关

- 发布：[`deploy.zh-CN.md`](deploy.zh-CN.md) · [`prompts/06-publish.zh-CN.md`](../prompts/06-publish.zh-CN.md)
- 总览：[`release/cloudflare/README.zh-CN.md`](../release/cloudflare/README.zh-CN.md)
