# Nginx 网关（步骤 ⑥）

[English](README.md) · **简体中文**

共享网络 `maker-flow` 上的 Docker 反向代理。SOP：`skills/deploy.md` → `release/publish/vps-gateway.md`。

## 布局

- `docker-compose.yml` — `nginx:1.27-alpine`，宿主机 `80:80`
- `nginx.conf` — include `/etc/nginx/conf.d/*.conf`
- `conf.d/00-default.conf` — 默认 vhost，尚无 MVP 时也能启动网关
- `conf.d/` — 各 MVP 的 server block（由部署脚本写入）
- `snippets/mvp-server.conf.example` — 模板（`__DOMAIN__`、`__MVP_NAME__`、`__CONTAINER_PORT__`）

## 服务器一次性准备

只需 Docker（本流程不需要 apt / 系统 Nginx / sudo）：

```bash
docker network create maker-flow 2>/dev/null || true
# 网关文件由 push-and-route.sh 同步到 GATEWAY_PATH（默认 /opt/maker-flow/gateway）
```

## 每个 MVP

在产品仓根目录优先使用 `release/deploy/push-and-route.sh`（Agent 内部可用 `maker-flow deploy`）。它会：

1. 同步网关 + MVP
2. `docker compose up` MVP
3. `docker network connect --alias <MVP_NAME> maker-flow <container>`
4. 安装 `conf.d/<MVP_NAME>.conf`
5. `nginx -t` 通过后 reload

## 本地冒烟

```bash
docker network create maker-flow 2>/dev/null || true
cd release/nginx
docker compose up -d
docker compose logs -f nginx
```

## 调试

```bash
docker compose -f /opt/maker-flow/gateway/docker-compose.yml exec nginx nginx -t
docker compose -f /opt/maker-flow/gateway/docker-compose.yml logs -f nginx
```
