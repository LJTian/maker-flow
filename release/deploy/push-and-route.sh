#!/usr/bin/env bash
# 将当前目录的 docker-compose 项目同步到服务器并重启
# 用法（在 MVP 项目根目录）:
#   MVP_NAME=idea1 MVP_PORT=8080 DOMAIN=idea1.your-domain.com \
#   DEPLOY_HOST=user@server DEPLOY_PATH=/opt/mvps/idea1 \
#   /path/to/maker-flow/release/deploy/push-and-route.sh
set -euo pipefail

MVP_NAME="${MVP_NAME:?请设置 MVP_NAME}"
MVP_PORT="${MVP_PORT:?请设置 MVP_PORT}"
DEPLOY_HOST="${DEPLOY_HOST:?请设置 DEPLOY_HOST，如 deploy@1.2.3.4}"
DEPLOY_PATH="${DEPLOY_PATH:-/opt/mvps/${MVP_NAME}}"
DOMAIN="${DOMAIN:-}"

echo "==> 同步 ${MVP_NAME} 到 ${DEPLOY_HOST}:${DEPLOY_PATH}"

ssh "$DEPLOY_HOST" "mkdir -p '${DEPLOY_PATH}'"
rsync -avz --delete \
  --exclude '.git' \
  --exclude '__pycache__' \
  --exclude 'bin' \
  ./ "${DEPLOY_HOST}:${DEPLOY_PATH}/"

ssh "$DEPLOY_HOST" bash -s <<EOF
set -euo pipefail
cd '${DEPLOY_PATH}'
export HOST_PORT=${MVP_PORT}
docker compose pull 2>/dev/null || true
docker compose up -d --build
docker compose ps
EOF

echo "==> 容器已在主机端口 ${MVP_PORT} 启动"
echo "==> 请在 Nginx 添加 server_name → 127.0.0.1:${MVP_PORT}"
if [[ -n "$DOMAIN" ]]; then
  echo "==> 请在 Cloudflare 添加 DNS: ${DOMAIN} → 服务器 IP（Proxied）"
fi
echo "完成。"
