# Cloudflare DNS API（Agent）

[English](dns-api.md) · **简体中文**

通过 Cloudflare API 做 DNS upsert，Zone 已在 CF 上后人类不必再点控制台。能力与 DDNS（A/AAAA）相同；Pages 自定义域常用 CNAME。

## 人类一次性准备

1. 域名 NS 已指向 Cloudflare（注册商侧 API 改不了）。
2. 创建 API Token：[API Tokens](https://dash.cloudflare.com/profile/api-tokens)。
3. 建议权限（可做在同一个 Token）：
   - **Account** → **Cloudflare Pages** → **Edit**（静态发布）
   - **Zone** → **DNS** → **Edit**（本指南）
   - Zone 资源：包含目标 Zone
4. 交给 Agent（优先环境变量）：
   - `CLOUDFLARE_API_TOKEN`
   - `CLOUDFLARE_ZONE_ID`
   - Pages 另需：`CLOUDFLARE_ACCOUNT_ID`

**禁止**编造或打印 Token。请人类自行 export。

## 何时用

| 记录 | 典型内容 | Proxied |
|------|----------|---------|
| `A` | VPS IPv4 | 通常 `true` |
| `AAAA` | VPS IPv6（同 DDNS） | 通常 `true` |
| `CNAME` | Pages 自定义域 → `<project>.pages.dev` | 通常 `true` |

有 Token + Zone ID 时优先 API（`skills/deploy.md`）。

## 助手脚本

```bash
export CLOUDFLARE_API_TOKEN=…
export CLOUDFLARE_ZONE_ID=…

"$(maker-flow root)/release/cloudflare/dns-upsert.sh" \
  --type CNAME \
  --name www.example.com \
  --content my-project.pages.dev \
  --proxied true
```

行为：按 `name` + `type` 查询 → 已有则 **PATCH**，否则 **POST**。可重复执行。

## Pages 自定义域流程

1. 按 [`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md) 发布，记下 `*.pages.dev` / 项目名。
2. 用 Pages Domains API 绑定主机名：
   ```bash
   curl -sS -X POST \
     -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" \
     -H "Content-Type: application/json" \
     "https://api.cloudflare.com/client/v4/accounts/${CLOUDFLARE_ACCOUNT_ID}/pages/projects/<PROJECT>/domains" \
     --data '{"name":"<HOST>"}'
   ```
3. 用本 API / `dns-upsert.sh` 写入 CNAME（或以 CF 返回的目标为准）。
4. 验证 `https://<自定义域>/`。

## VPS 网关流程

`vps-gateway` 发布后，对服务器公网地址 upsert **A** / **AAAA**，`proxied=true`。

## 验证

```bash
dig +short <name> A
dig +short <name> AAAA
dig +short <name> CNAME
curl -sI "https://<name>/"
```

## 相关

- Pages：[`../publish/cloudflare-pages.md`](../publish/cloudflare-pages.md)
- 登记表示例：[`subdomain-registry.example.zh-CN.md`](subdomain-registry.example.zh-CN.md)
