# The Release

将本地验证通过的 Docker 容器在 **10 分钟内** 暴露到公网。

## 目录

```
release/
├── nginx/          # 反向代理配置与片段
├── cloudflare/     # DNS、SSL、子域名分配
└── deploy/         # 推送与路由脚本
```

## 端口池

| 子域名示例 | 主机端口 | 用途 |
|------------|----------|------|
| `test.your-domain.com` | 8080 | 通路测试 |
| `idea1.your-domain.com` | 8080 | MVP #1 |
| `idea2.your-domain.com` | 8081 | MVP #2 |
| … | 8082–8090 | 预留 |

容器内端口固定为 8080，仅 **主机映射端口** 按 MVP 递增。

## 标准上线流程

1. 本地 `docker compose up` 通过
2. `release/deploy/push-and-route.sh` 同步镜像/compose 到服务器并重启
3. 复制 `nginx/snippets/mvp-server.conf.example` 为新 server 块
4. Cloudflare 添加子域名（Proxied 橙云）
5. 访问 `https://ideaN.your-domain.com`

## 前置条件

- 云服务器：Docker、Docker Compose、Nginx
- 域名 NS 已指向 Cloudflare
- SSH 免密或 `DEPLOY_SSH` 环境变量
