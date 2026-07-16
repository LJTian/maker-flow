# 上线部署技能

[English](deploy.md) · **简体中文**

**适用步骤：** ⑥ 上线部署  
**前置：** 步骤 ⑤ MVP 已本地验收通过

## 目标

在 10 分钟内将 `workspace/<项目名>/` 中的容器暴露到公网。

## 检查清单

- [ ] `docker compose up` 本地 health 200
- [ ] 在 `release/cloudflare/subdomain-registry.example.md` 登记子域名与端口
- [ ] 服务器 Docker + Nginx 已就绪
- [ ] Cloudflare DNS 橙云代理开启

## 部署动作

在 MVP 项目目录：

```bash
export MVP_NAME=idea1
export MVP_PORT=8080
export DOMAIN=idea1.your-domain.com
export DEPLOY_HOST=deploy@your-server
export DEPLOY_PATH=/opt/mvps/idea1

/path/to/maker-flow/release/deploy/push-and-route.sh
```

## Nginx

复制 `release/nginx/snippets/mvp-server.conf.example`，修改 `server_name` 与 `proxy_pass` 端口。

## Cloudflare

添加 A 记录指向服务器 IP（Proxied）。详见 `release/cloudflare/README.md`。

## 验证

```bash
curl -I https://idea1.your-domain.com/health
```

## 回滚

```bash
ssh $DEPLOY_HOST "cd $DEPLOY_PATH && docker compose down"
```

并在 Nginx / DNS 登记表中释放端口与子域名。

## 详细文档

- [release/README.md](../release/README.md)
- [release/deploy/](../release/deploy/)
