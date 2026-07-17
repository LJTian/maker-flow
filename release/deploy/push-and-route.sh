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
