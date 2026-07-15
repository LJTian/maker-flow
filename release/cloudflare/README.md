# Cloudflare 配置指南

## 1. 添加站点

1. Cloudflare Dashboard → Add site → 输入 `your-domain.com`
2. 按提示将域名 NS 改为 Cloudflare 分配的记录
3. 等待 Active

## 2. SSL/TLS

推荐 MVP 阶段：

- **SSL/TLS encryption mode**: `Full` 或 `Full (strict)`（源站有自签证书时用 Full）
- 边缘证书由 Cloudflare 自动签发，无需手动申请 Let's Encrypt

## 3. DNS 记录

| 类型 | 名称 | 内容 | 代理 |
|------|------|------|------|
| A | `@` 或 `test` | 服务器公网 IP | 已代理（橙云） |
| A | `idea1` | 服务器公网 IP | 已代理 |
| A | `idea2` | 服务器公网 IP | 已代理 |

同一 IP 可挂多个子域名；Nginx 按 `server_name` 分流到不同端口。

## 4. 子域名分配表（手写维护）

复制 `subdomain-registry.example.md` 为 `subdomain-registry.md`（已 gitignore 可自建），避免端口冲突。

## 5. 可选：API Token

若后续用 API 自动创建 DNS 记录，创建 Token 权限：

- Zone → DNS → Edit
- 仅包含目标 Zone

环境变量：`CLOUDFLARE_API_TOKEN`、`CLOUDFLARE_ZONE_ID`

## 6. 验证

```bash
curl -I https://test.your-domain.com
```

应返回 `200` 或应用正常响应，且证书由 Cloudflare 签发。
