[English](README.md) · **简体中文**

# release/

**步骤 ⑥** 部署基建。Agent **MUST** 遵循 `skills/deploy.md`，且 **MUST** 等待步骤 ⑤ 人工批准。

## 布局

```
release/
├── nginx/          # 反向代理片段
├── cloudflare/     # DNS / SSL / 子域名登记
└── deploy/         # 推送 + 路由脚本
```

## 端口池

| 宿主机端口 | 示例 |
|------------|------|
| 8080 | 首个 MVP / 测试 |
| 8081–8090 | 后续 MVP |

容器监听 8080；仅宿主机映射递增。

## Agent 部署顺序

1. 确认步骤 ⑤ 已批准。
2. 在 cloudflare 登记示例 / 线上登记表中登记子域名 + 端口。
3. 运行 `deploy/push-and-route.sh`，传入 `MVP_NAME`、`MVP_PORT`、`DOMAIN`、`DEPLOY_HOST`、`DEPLOY_PATH`。
4. 从 `nginx/snippets/mvp-server.conf.example` 安装 nginx server block。
5. 确保 Cloudflare DNS 已 Proxied；在公网 URL 验证 `GET /health`。

## 前置条件

- 服务器：Docker、Compose、Nginx
- 域名 NS → Cloudflare
- 对 `DEPLOY_HOST` 有 SSH 访问
