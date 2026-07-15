#!/usr/bin/env bash
# 在服务器上安装静态测试页与 Nginx 片段（需 sudo）
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
STATIC_SRC="${REPO_ROOT}/release/nginx/static"
SNIPPET_SRC="${REPO_ROOT}/release/nginx/snippets/static-test.conf.example"

sudo mkdir -p /var/www/maker-flow-test
sudo cp "${STATIC_SRC}/index.html" /var/www/maker-flow-test/
sudo mkdir -p /etc/nginx/snippets/maker-flow
sudo cp "${SNIPPET_SRC}" /etc/nginx/snippets/maker-flow/static-test.conf

echo "请确认 /etc/nginx/nginx.conf 的 http 块包含:"
echo '  include /etc/nginx/snippets/maker-flow/*.conf;'
echo "然后: sudo nginx -t && sudo systemctl reload nginx"
