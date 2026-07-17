# 设计规格：Docker 共享网络 Nginx 网关（方案 B）

**日期：** 2026-07-17  
**状态：** 已批准（实现计划：`docs/superpowers/plans/2026-07-17-docker-nginx-gateway.md`）  
**范围：** 步骤 ⑥ 部署基础设施；去掉宿主机 Nginx / sudo

## 背景与目标

当前部署依赖宿主机安装 Nginx，并用 `sudo` 写入 `/etc/nginx`、`/var/www`。与「整条运行链路都在 Docker 内」冲突，且产品仓用户需要 root。

目标：

1. 反向代理以 **Docker Compose + `nginx:alpine`** 常驻网关运行。
2. MVP 与网关通过 **外部共享网络 `maker-flow`** 互通；Nginx 用 **网络别名** 反代（如 `http://my-todo:8080`）。
3. 文档与脚本中 **不得出现** `sudo` / `apt` / `systemctl` / 宿主机 `/etc/nginx`。
4. 删除旧的宿主机静态测试通路。

非目标（本次不做）：

- 不把 `maker-flow deploy` CLI 一并实现（可后续）。
- 不改 Cloudflare 流程（仍用 Proxied DNS → 服务器 80/443）。
- 不强制本地验收去掉 `HOST_PORT` 映射（本地仍可映射；生产流量走共享网络）。

## 架构

```text
Internet → Cloudflare (Proxied) → 宿主机 :80/:443
                                      ↓
                            [gateway] nginx 容器
                                      ↓  Docker network: maker-flow
                    ┌─────────────────┼─────────────────┐
                    ↓                 ↓                 ↓
              my-todo:8080      other-mvp:8080     …
              (compose 服务，network alias = MVP_NAME)
```

- 网关发布宿主机 `80`（可选 `443` 留给后续；MVP 阶段 Cloudflare 终止 TLS 时仅需 80）。
- 各 MVP **不必**为生产流量再占宿主机 8080–8090；端口池降级为「本地调试可选」。
- 服务发现：部署时把 MVP 容器接入 `maker-flow`，并设置别名 `MVP_NAME`（与 `server_name` / 注册表一致）。

## 组件与布局

```text
release/nginx/
├── docker-compose.yml      # 网关：nginx + 挂载 conf /（可选）static 已删
├── nginx.conf              # 主配置：include /etc/nginx/conf.d/*.conf;
├── conf.d/
│   └── .gitkeep            # 运行时写入 <mvp>.conf
├── snippets/
│   └── mvp-server.conf.example   # 模板：proxy_pass http://<alias>:<container_port>;
├── README.md / README.zh-CN.md
└── （删除）static/、snippets/static-test.conf.example

release/deploy/
├── push-and-route.sh       # 扩展：网络 + 别名 + 写入 conf + nginx -t + reload
└── （删除）install-static-test.sh
```

网关 compose 要点：

- `image: nginx:1.27-alpine`（与 web-vite 对齐或同系列）
- `ports: ["80:80"]`
- volumes：`./nginx.conf` → `/etc/nginx/nginx.conf`；`./conf.d` → `/etc/nginx/conf.d`
- `networks: maker-flow`（`external: true`）
- 首次部署前：`docker network create maker-flow`（幂等）

MVP server 片段模板（概念）：

```nginx
server {
    listen 80;
    server_name ${DOMAIN};

    location / {
        proxy_pass http://${MVP_NAME}:${CONTAINER_PORT};
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

`CONTAINER_PORT`：Go API/worker 默认 `8080`；`web-vite` 默认 `80`。由 env 传入，默认 `8080`。

## 部署数据流（`push-and-route.sh`）

前置：`MVP_NAME`、`DEPLOY_HOST`、`DOMAIN`；`MVP_PORT` 语义调整为 **容器内监听端口**（默认 8080），不再表示「宿主机映射端口」。若需兼容旧文档，可同时接受 `CONTAINER_PORT` 别名。

步骤：

1. SSH：`docker network create maker-flow 2>/dev/null || true`
2. 确保网关目录在服务器存在（rsync `release/nginx` 到固定路径，如 `/opt/maker-flow/gateway`，或从已安装的 `~/.maker-flow` / `MAKER_FLOW_ROOT` 同步）
3. rsync 产品仓 → `DEPLOY_PATH`
4. 远端：`docker compose up -d --build`（产品仓）
5. 将 MVP 主服务容器 `connect` 到 `maker-flow`，并 `--alias ${MVP_NAME}`  
   - 需能解析 compose 项目中的服务名（`api` / `worker` / `web`）；可用 env `MVP_SERVICE`（默认尝试常见名或由调用方指定）
6. 渲染 `mvp-server.conf.example` → 网关 `conf.d/${MVP_NAME}.conf`
7. `docker compose -f gateway/docker-compose.yml up -d`
8. `docker compose exec nginx nginx -t` → 成功则 `nginx -s reload`；失败则保留旧 conf、报错退出（不 reload）
9. 打印 Cloudflare DNS 提示

回滚：`docker compose down` 产品仓；删除 `conf.d/${MVP_NAME}.conf` 后 reload；可选 `docker network disconnect`。

## 产品仓 compose 约定

- **本地验收：** 可继续 `ports: HOST_PORT:CONTAINER_PORT`。
- **生产：** 不依赖宿主机端口；必须能加入外部网络 `maker-flow`。  
  推荐在组装时写入（或由 deploy 脚本 `docker network connect`，不强制改 compose 文件）：

```yaml
networks:
  default:
  maker-flow:
    external: true
```

本次优先用 **deploy 脚本 `docker network connect`**，减少对已有产品仓 compose 的硬依赖；文档说明「可选在 compose 声明 external network」。

服务容器名 / 别名：`MVP_NAME` 必须在共享网络上唯一。

## 删除清单

| 路径 | 动作 |
|------|------|
| `release/deploy/install-static-test.sh` | 删除 |
| `release/nginx/static/` | 删除 |
| `release/nginx/snippets/static-test.conf.example` | 删除 |
| 文档中的 `sudo` / 宿主机 nginx 安装步骤 | 改写为 Docker 网关流程 |
| `skills/template-matching` 中「仅静态通路测试 → release/nginx/static」 | 删除或改写 |

## 文档 / 技能更新

- `release/nginx/README*.md`、`release/README*.md`
- `skills/deploy*.md`、`docs/workflow*.md`（若提及宿主机 Nginx）
- `skills/template-matching*.md`（去掉 static 通路）
- 前置条件：服务器仅需 Docker（及部署用户可用 docker）；不再要求系统 Nginx

## 错误处理

- `nginx -t` 失败：不 reload，退出非零，保留上一份有效配置。
- 别名冲突：`docker network connect` 失败时明确报错。
- 网关 compose 未启动：步骤 7 先 `up -d`。

## 测试计划

1. 本机：`docker network create maker-flow`；启动网关；用假 upstream 或 smoke MVP 验证 `curl -H Host:… http://127.0.0.1/`。
2. 文档 grep：仓库内无 `sudo`（除历史无关处）、无 `install-static-test`、无 `/var/www/maker-flow-test`。
3. 对照 `skills/deploy.md` 走通一遍命令序列（可 dry-run 脚本路径）。

## 风险与缓解

| 风险 | 缓解 |
|------|------|
| 多服务 compose 不知连哪个容器 | 要求 `MVP_SERVICE` 或文档约定单入口服务名 |
| Cloudflare 仅指到 IP，网关未占 80 | README 明确网关必须 `80:80` |
| 本地仍映射端口与生产路径不一致 | 文档区分「验收用端口 / 生产用网络别名」 |
