# Nginx 反向代理

## 安装（Ubuntu/Debian 示例）

```bash
sudo apt update && sudo apt install -y nginx
sudo mkdir -p /etc/nginx/snippets/maker-flow
```

## 主配置建议

在 `http` 块中：

```nginx
include /etc/nginx/snippets/maker-flow/*.conf;
```

## 新增 MVP

1. 复制 `snippets/mvp-server.conf.example` 为 `idea1.conf`
2. 修改 `server_name` 与 `proxy_pass` 端口
3. 测试并重载：

```bash
sudo nginx -t && sudo systemctl reload nginx
```

## 日志

默认 access/error log 按 server 块写入；排查时：

```bash
sudo tail -f /var/log/nginx/error.log
```

## 与 Cloudflare

- 源站可监听 80；Cloudflare 边缘终止 HTTPS
- 若需识别真实 IP，可配置 `real_ip` 与 Cloudflare IP 段（进阶）
