# Docker 共享网络 Nginx 网关 实现计划

> **给代理执行者：** 必需子技能：使用 superpowers:subagent-driven-development（推荐）或 superpowers:executing-plans 逐任务实现此计划。步骤使用复选框（`- [ ]`）语法跟踪。

**目标：** 把步骤 ⑥ 的反向代理改为 Docker 共享网络上的 Nginx 网关，删除宿主机 Nginx / sudo / 静态测试通路。

**架构：** 常驻 `release/nginx` 网关容器接入外部网络 `maker-flow`；`push-and-route.sh` 部署 MVP 后把入口容器 connect 进该网络并设别名 `MVP_NAME`，渲染 `conf.d/<name>.conf`，经 `nginx -t` 成功后再 reload。

**技术栈：** Docker Compose、`nginx:1.27-alpine`、bash（rsync/ssh）、现有 Cloudflare 文档不变。

**规格：** [`docs/superpowers/specs/2026-07-17-docker-nginx-gateway-design.md`](../specs/2026-07-17-docker-nginx-gateway-design.md)

---

## 文件结构（创建 / 修改 / 删除）

| 路径 | 职责 |
|------|------|
| `release/nginx/docker-compose.yml` | 网关服务定义 |
| `release/nginx/nginx.conf` | 主配置，include conf.d |
| `release/nginx/conf.d/.gitkeep` | 运行时 conf 目录占位 |
| `release/nginx/snippets/mvp-server.conf.example` | 按别名反代的模板 |
| `release/nginx/README.md` / `.zh-CN.md` | 网关用法（无 sudo） |
| `release/deploy/push-and-route.sh` | 建网、同步网关、部署 MVP、连网、写 conf、reload |
| `release/README.md` / `.zh-CN.md` | 部署总览与前置条件 |
| `skills/deploy.md` / `.zh-CN.md` | 步骤 ⑥ SOP |
| `skills/template-matching.md` / `.zh-CN.md` | 去掉 static 通路行 |
| `docs/workflow.md` / `.zh-CN.md` | 若仍写宿主机 nginx，改为网关 |
| **删除** `release/deploy/install-static-test.sh` | |
| **删除** `release/nginx/static/`（含 `index.html`） | |
| **删除** `release/nginx/snippets/static-test.conf.example` | |

---

### Task 1: 网关 Compose + nginx 主配置

**Files:**
- Create: `release/nginx/docker-compose.yml`
- Create: `release/nginx/nginx.conf`
- Create: `release/nginx/conf.d/.gitkeep`

- [ ] **Step 1: 写入 `docker-compose.yml`**

```yaml
services:
  nginx:
    image: nginx:1.27-alpine
    ports:
      - "80:80"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf:ro
      - ./conf.d:/etc/nginx/conf.d:ro
    networks:
      - maker-flow
    restart: unless-stopped

networks:
  maker-flow:
    external: true
```

- [ ] **Step 2: 写入 `nginx.conf`**

```nginx
worker_processes auto;

events {
    worker_connections 1024;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;
    sendfile      on;
    keepalive_timeout 65;

    # Per-MVP server blocks written by push-and-route.sh
    include /etc/nginx/conf.d/*.conf;
}
```

- [ ] **Step 3: 创建空目录占位**

```bash
mkdir -p release/nginx/conf.d
touch release/nginx/conf.d/.gitkeep
```

- [ ] **Step 4: 本地冒烟（需 Docker）**

```bash
docker network create maker-flow 2>/dev/null || true
cd release/nginx
printf 'server { listen 80; return 200 "gateway-ok\\n"; add_header Content-Type text/plain; }\n' > conf.d/_default.conf
docker compose up -d
curl -sS http://127.0.0.1/ | grep gateway-ok
docker compose down
rm -f conf.d/_default.conf
```

Expected: curl 输出含 `gateway-ok`。

- [ ] **Step 5: Commit**

```bash
git add release/nginx/docker-compose.yml release/nginx/nginx.conf release/nginx/conf.d/.gitkeep
git commit -m "Add Docker Nginx gateway compose and main config."
```

---

### Task 2: 更新 MVP server 片段模板

**Files:**
- Modify: `release/nginx/snippets/mvp-server.conf.example`

- [ ] **Step 1: 将 `proxy_pass` 改为网络别名（不再用 127.0.0.1）**

整文件替换为：

```nginx
# MVP reverse-proxy snippet — copy to conf.d/<MVP_NAME>.conf
# Placeholders replaced by push-and-route.sh:
#   __DOMAIN__          e.g. idea1.your-domain.com
#   __MVP_NAME__        Docker network alias on maker-flow
#   __CONTAINER_PORT__  listen port inside the MVP container (8080 or 80)

server {
    listen 80;
    server_name __DOMAIN__;

    location / {
        proxy_pass http://__MVP_NAME__:__CONTAINER_PORT__;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

- [ ] **Step 2: Commit**

```bash
git add release/nginx/snippets/mvp-server.conf.example
git commit -m "Point MVP nginx snippet at Docker network aliases."
```

---

### Task 3: 删除宿主机静态测试与 sudo 安装脚本

**Files:**
- Delete: `release/deploy/install-static-test.sh`
- Delete: `release/nginx/static/index.html`（及空目录 `static/`）
- Delete: `release/nginx/snippets/static-test.conf.example`

- [ ] **Step 1: 删除文件**

```bash
rm -f release/deploy/install-static-test.sh
rm -f release/nginx/snippets/static-test.conf.example
rm -rf release/nginx/static
```

- [ ] **Step 2: 确认无残留引用**

```bash
rg -n 'install-static-test|/var/www/maker-flow-test|static-test\.conf' . || true
```

Expected: 操作文档与脚本无匹配（规格/计划中的删除清单文字可保留）。

- [ ] **Step 3: Commit**

```bash
git add -u release/deploy release/nginx
git commit -m "Remove host Nginx static-test install path."
```

---

### Task 4: 重写 `push-and-route.sh`

**Files:**
- Modify: `release/deploy/push-and-route.sh`

环境变量约定：

| 变量 | 必填 | 含义 |
|------|------|------|
| `MVP_NAME` | 是 | 网络别名 + conf 文件名 |
| `DOMAIN` | 是 | `server_name` |
| `DEPLOY_HOST` | 是 | `user@host` |
| `DEPLOY_PATH` | 否 | 默认 `/opt/mvps/${MVP_NAME}` |
| `GATEWAY_PATH` | 否 | 默认 `/opt/maker-flow/gateway` |
| `CONTAINER_PORT` | 否 | 容器内端口，默认 `8080`；兼容旧名 `MVP_PORT` |
| `MVP_SERVICE` | 否 | compose 服务名；未设时依次尝试 `api`、`web`、`worker` |
| `GATEWAY_SRC` | 否 | 本机网关源目录；默认相对本脚本的 `../nginx` |

- [ ] **Step 1: 用完整脚本替换 `push-and-route.sh`**

实现要点（写入文件时保持 `set -euo pipefail`）：

1. 校验 `GATEWAY_SRC/docker-compose.yml` 与 snippet 模板存在。
2. `ssh`：`docker network create maker-flow 2>/dev/null || true`
3. rsync 网关时 **排除整个 `conf.d`**，远端 `mkdir -p conf.d`，避免误删已有 MVP conf。
4. rsync 产品仓到 `DEPLOY_PATH`，远端 `docker compose up -d --build`。
5. 解析入口容器：若设 `MVP_SERVICE` 则用它；否则依次 `api` / `web` / `worker` 的 `docker compose ps -q`。
6. `docker network disconnect maker-flow <id> 2>/dev/null || true` 后 `docker network connect --alias "$MVP_NAME" maker-flow <id>`。
7. `sed` 替换 `__DOMAIN__` / `__MVP_NAME__` / `__CONTAINER_PORT__`，`scp` 到 `${GATEWAY_PATH}/conf.d/${MVP_NAME}.conf`。
8. 远端网关目录：`docker compose up -d`；`docker compose exec -T nginx nginx -t` 失败则 exit 1 且 **不** reload；成功则 `nginx -s reload`。
9. 打印 Cloudflare DNS 提示。

完整脚本正文见本任务下方「参考实现」。

- [ ] **Step 2: 语法检查**

```bash
bash -n release/deploy/push-and-route.sh
chmod +x release/deploy/push-and-route.sh
```

Expected: 无输出、exit 0。

- [ ] **Step 3: Commit**

```bash
git add release/deploy/push-and-route.sh
git commit -m "Deploy via shared maker-flow network and gateway reload."
```

#### 参考实现

```bash
#!/usr/bin/env bash
# Sync MVP compose project to server, attach to maker-flow network, reload gateway.
# Run from the MVP project root:
#   MVP_NAME=idea1 DOMAIN=idea1.your-domain.com DEPLOY_HOST=user@server \
#   CONTAINER_PORT=8080 MVP_SERVICE=api \
#   /path/to/maker-flow/release/deploy/push-and-route.sh
set -euo pipefail

MVP_NAME="${MVP_NAME:?set MVP_NAME}"
DOMAIN="${DOMAIN:?set DOMAIN}"
DEPLOY_HOST="${DEPLOY_HOST:?set DEPLOY_HOST, e.g. deploy@1.2.3.4}"
DEPLOY_PATH="${DEPLOY_PATH:-/opt/mvps/${MVP_NAME}}"
GATEWAY_PATH="${GATEWAY_PATH:-/opt/maker-flow/gateway}"
CONTAINER_PORT="${CONTAINER_PORT:-${MVP_PORT:-8080}}"
MVP_SERVICE="${MVP_SERVICE:-}"

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
GATEWAY_SRC="${GATEWAY_SRC:-${SCRIPT_DIR}/../nginx}"

if [[ ! -f "${GATEWAY_SRC}/docker-compose.yml" ]]; then
  echo "error: gateway source missing: ${GATEWAY_SRC}/docker-compose.yml" >&2
  exit 1
fi
if [[ ! -f "${GATEWAY_SRC}/snippets/mvp-server.conf.example" ]]; then
  echo "error: missing snippet template" >&2
  exit 1
fi

echo "==> ensure network maker-flow on ${DEPLOY_HOST}"
ssh "$DEPLOY_HOST" 'docker network create maker-flow 2>/dev/null || true'

echo "==> sync gateway → ${DEPLOY_HOST}:${GATEWAY_PATH}"
ssh "$DEPLOY_HOST" "mkdir -p '${GATEWAY_PATH}/conf.d'"
rsync -avz \
  --exclude 'conf.d' \
  "${GATEWAY_SRC}/" "${DEPLOY_HOST}:${GATEWAY_PATH}/"

echo "==> sync MVP → ${DEPLOY_HOST}:${DEPLOY_PATH}"
ssh "$DEPLOY_HOST" "mkdir -p '${DEPLOY_PATH}'"
rsync -avz --delete \
  --exclude '.git' \
  --exclude '__pycache__' \
  --exclude 'bin' \
  ./ "${DEPLOY_HOST}:${DEPLOY_PATH}/"

echo "==> compose up MVP"
ssh "$DEPLOY_HOST" bash -s <<EOF
set -euo pipefail
cd '${DEPLOY_PATH}'
docker compose up -d --build
docker compose ps
EOF

resolve_container() {
  local svc="$1"
  ssh "$DEPLOY_HOST" "cd '${DEPLOY_PATH}' && docker compose ps -q '${svc}'" 2>/dev/null || true
}

CONTAINER_ID=""
if [[ -n "${MVP_SERVICE}" ]]; then
  CONTAINER_ID="$(resolve_container "${MVP_SERVICE}")"
  [[ -n "${CONTAINER_ID}" ]] || { echo "error: no container for MVP_SERVICE=${MVP_SERVICE}" >&2; exit 1; }
else
  for svc in api web worker; do
    CONTAINER_ID="$(resolve_container "${svc}")"
    if [[ -n "${CONTAINER_ID}" ]]; then
      MVP_SERVICE="${svc}"
      break
    fi
  done
fi
[[ -n "${CONTAINER_ID}" ]] || {
  echo "error: could not find service api|web|worker; set MVP_SERVICE" >&2
  exit 1
}

echo "==> connect ${MVP_SERVICE} (${CONTAINER_ID:0:12}) as alias ${MVP_NAME}"
ssh "$DEPLOY_HOST" bash -s <<EOF
set -euo pipefail
docker network disconnect maker-flow '${CONTAINER_ID}' 2>/dev/null || true
docker network connect --alias '${MVP_NAME}' maker-flow '${CONTAINER_ID}'
EOF

TMP_CONF="$(mktemp)"
sed \
  -e "s/__DOMAIN__/${DOMAIN}/g" \
  -e "s/__MVP_NAME__/${MVP_NAME}/g" \
  -e "s/__CONTAINER_PORT__/${CONTAINER_PORT}/g" \
  "${GATEWAY_SRC}/snippets/mvp-server.conf.example" > "${TMP_CONF}"

echo "==> install conf.d/${MVP_NAME}.conf"
scp "${TMP_CONF}" "${DEPLOY_HOST}:${GATEWAY_PATH}/conf.d/${MVP_NAME}.conf"
rm -f "${TMP_CONF}"

echo "==> gateway up + nginx -t + reload"
ssh "$DEPLOY_HOST" bash -s <<EOF
set -euo pipefail
cd '${GATEWAY_PATH}'
docker compose up -d
if ! docker compose exec -T nginx nginx -t; then
  echo "error: nginx -t failed; conf left in place but NOT reloaded" >&2
  exit 1
fi
docker compose exec -T nginx nginx -s reload
docker compose ps
EOF

echo "==> done: ${DOMAIN} → ${MVP_NAME}:${CONTAINER_PORT} on maker-flow"
echo "==> Cloudflare: A record ${DOMAIN} → server IP (Proxied)"
```

---

### Task 5: 重写 nginx / release 文档（去 sudo）

**Files:**
- Modify: `release/nginx/README.md`
- Modify: `release/nginx/README.zh-CN.md`
- Modify: `release/README.md`
- Modify: `release/README.zh-CN.md`

- [ ] **Step 1: 重写英文 `release/nginx/README.md`**

内容须覆盖：布局、`docker network create maker-flow`、由 `push-and-route.sh` 完成 per-MVP、本地 smoke、`docker compose exec nginx` 调试。**禁止** `sudo` / `apt` / `systemctl` / `/etc/nginx` 宿主机路径。

- [ ] **Step 2: 同步中文 `README.zh-CN.md`**

- [ ] **Step 3: 更新 `release/README.md`**

- Prerequisites：仅 Docker + Compose（部署用户可用 docker）；不要求宿主机 Nginx。
- Agent sequence：由脚本写入网关 `conf.d`，不再「Install nginx server block on host」。
- Port pool：8080–8090 标为本地调试可选；生产走网关 80 + 别名。

- [ ] **Step 4: 同步 `release/README.zh-CN.md`**

- [ ] **Step 5: 确认 release 下无 sudo**

```bash
rg -n 'sudo|apt install.*nginx|systemctl|/etc/nginx|/var/www/maker-flow' release/
```

Expected: 无匹配。

- [ ] **Step 6: Commit**

```bash
git add release/nginx/README.md release/nginx/README.zh-CN.md release/README.md release/README.zh-CN.md
git commit -m "Document Docker gateway deploy; drop host Nginx sudo steps."
```

---

### Task 6: 更新 skills / workflow / template-matching

**Files:**
- Modify: `skills/deploy.md`
- Modify: `skills/deploy.zh-CN.md`
- Modify: `skills/template-matching.md`
- Modify: `skills/template-matching.zh-CN.md`
- Modify: `docs/workflow.md`（若仍写宿主机 nginx）
- Modify: `docs/workflow.zh-CN.md`

- [ ] **Step 1: 更新 `skills/deploy.md`**

- Checklist：`Server Docker ready; gateway on maker-flow`。
- Actions 示例改用 `DOMAIN` + `CONTAINER_PORT` + `MVP_SERVICE`，去掉「宿主机 MVP_PORT 映射」作为生产必填语义。
- Nginx 小节：脚本自动装 conf；手动时在网关容器内 `nginx -t` / reload。

- [ ] **Step 2: 同步 `skills/deploy.zh-CN.md`**

- [ ] **Step 3: 删 template-matching 静态通路行**

英文删除：`Static path smoke only | release/nginx/static`  
中文删除：`仅静态通路测试 | release/nginx/static`

- [ ] **Step 4: 核对 `docs/workflow.md`**

步骤 ⑥ 指向 `push-and-route.sh` + Docker 网关资产。

- [ ] **Step 5: 全仓扫操作文档**

```bash
rg -n 'sudo|install-static-test|/var/www/maker-flow-test|static-test\.conf' --glob '*.md' --glob '*.sh' .
```

Expected: `release/`、`skills/deploy*`、操作 README 无匹配。

- [ ] **Step 6: Commit**

```bash
git add skills/deploy.md skills/deploy.zh-CN.md skills/template-matching.md skills/template-matching.zh-CN.md docs/workflow.md docs/workflow.zh-CN.md
git commit -m "Align deploy skill and workflow with Docker gateway."
```

---

### Task 7: 规格状态 + 最终验证

**Files:**
- Modify: `docs/superpowers/specs/2026-07-17-docker-nginx-gateway-design.md`

- [ ] **Step 1: 规格状态改为 `已实现`**

- [ ] **Step 2: 最终检查**

```bash
test -f release/nginx/docker-compose.yml
test -f release/nginx/nginx.conf
test ! -e release/deploy/install-static-test.sh
test ! -e release/nginx/static
test ! -e release/nginx/snippets/static-test.conf.example
bash -n release/deploy/push-and-route.sh
rg -n 'sudo' release/ skills/deploy.md skills/deploy.zh-CN.md || true
```

Expected: 存在性正确；`bash -n` 通过；上述路径无 `sudo`。

- [ ] **Step 3: Commit**

```bash
git add docs/superpowers/specs/2026-07-17-docker-nginx-gateway-design.md
git commit -m "Mark Docker nginx gateway spec implemented."
```

---

## 规格覆盖对照

| 规格要求 | 任务 |
|----------|------|
| Docker Compose + nginx:alpine 网关 | Task 1 |
| 共享网络 `maker-flow` + 别名反代 | Task 2、4 |
| 无 sudo/apt/systemctl/宿主机 nginx | Task 3、5、6 |
| 删除 static / install-static-test | Task 3 |
| push-and-route：建网、同步、connect、conf、nginx -t、reload | Task 4 |
| 文档与 skills 更新 | Task 5、6 |
| CONTAINER_PORT / MVP_SERVICE | Task 4、6 |
| 本地 HOST_PORT 保留（不强制改 app compose） | 非目标；未改 templates compose |

## 自查

- 无 TBD 占位步骤；脚本有完整参考实现。
- rsync 排除整个 `conf.d`，避免误删远端 MVP conf。
- 占位符统一：`__DOMAIN__` / `__MVP_NAME__` / `__CONTAINER_PORT__`。
