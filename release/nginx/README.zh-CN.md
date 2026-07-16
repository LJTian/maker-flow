[English](README.md) · **简体中文**

# Nginx（步骤 ⑥）

反向代理运维说明。SOP：`skills/deploy.md`。

## 宿主机准备（一次性）

```bash
sudo apt update && sudo apt install -y nginx
sudo mkdir -p /etc/nginx/snippets/maker-flow
```

在 `http` 块中引入：

```nginx
include /etc/nginx/snippets/maker-flow/*.conf;
```

## 每个 MVP

1. 复制 `snippets/mvp-server.conf.example` → `snippets/maker-flow/<name>.conf`
2. 设置 `server_name` 与 `proxy_pass http://127.0.0.1:<HOST_PORT>`
3. 生效：

```bash
sudo nginx -t && sudo systemctl reload nginx
```

## 静态通路测试

`snippets/static-test.conf.example` + `static/index.html` — 仅用于 Cloudflare→Nginx 冒烟。

## 调试

```bash
sudo tail -f /var/log/nginx/error.log
```
