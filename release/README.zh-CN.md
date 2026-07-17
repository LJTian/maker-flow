[English](README.md) · **简体中文**

# release/

**步骤 ⑥** 部署基建。Agent **MUST** 遵循 `skills/deploy.md`，且 **MUST** 等待步骤 ⑤ 人工批准。

## 布局

```
release/
├── nginx/          # Docker Nginx 网关（共享网络 maker-flow）
├── cloudflare/     # DNS / SSL / 子域名登记
└── deploy/         # 推送 + 路由脚本
```

## 端口

| 端口 | 角色 |
|------|------|
| **80**（宿主机） | 仅网关 — Cloudflare 公网入口 |
| `8080` / `80`（容器内） | MVP 容器监听端口（`CONTAINER_PORT`） |
| `8080–8090`（宿主机） | **可选**本地调试映射；生产不需要 |

生产流量：Cloudflare → 网关 `:80` → Docker 网络别名 `MVP_NAME:CONTAINER_PORT`。

## Agent 部署顺序

1. 确认步骤 ⑤ 已批准。
2. 在 cloudflare 登记示例 / 线上登记表中登记子域名。
3. 在产品仓根目录运行 `maker-flow deploy`（或 `deploy/push-and-route.sh`），传入 `DOMAIN`、`DEPLOY_HOST`，可选 `CONTAINER_PORT` / `MVP_SERVICE`。
4. 脚本同步网关、将 MVP 接入 `maker-flow`、写入 `conf.d/<MVP_NAME>.conf`、执行 `nginx -t` 后 reload。
5. 确保 Cloudflare DNS 已 Proxied；在公网 URL 验证 `GET /health`。

## 前置条件

- 服务器：Docker + Compose（部署用户可运行 docker；**不需要宿主机 Nginx**）
- 域名 NS → Cloudflare
- 对 `DEPLOY_HOST` 有 SSH 访问
