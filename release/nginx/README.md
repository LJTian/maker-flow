# Nginx (step 6)

**English** · [简体中文](README.zh-CN.md)

Agent ops for reverse proxy. SOP: `skills/deploy.md`.

## Host prep (once)

```bash
sudo apt update && sudo apt install -y nginx
sudo mkdir -p /etc/nginx/snippets/maker-flow
```

Include in `http` block:

```nginx
include /etc/nginx/snippets/maker-flow/*.conf;
```

## Per-MVP

1. Copy `snippets/mvp-server.conf.example` → `snippets/maker-flow/<name>.conf`
2. Set `server_name` and `proxy_pass http://127.0.0.1:<HOST_PORT>`
3. Apply:

```bash
sudo nginx -t && sudo systemctl reload nginx
```

## Static path test

`snippets/static-test.conf.example` + `static/index.html` — Cloudflare→Nginx smoke only.

## Debug

```bash
sudo tail -f /var/log/nginx/error.log
```
