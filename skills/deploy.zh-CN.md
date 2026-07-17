# 上线部署技能

[English](deploy.md) · **简体中文**

**适用步骤：** ⑥ 上线部署  
**前置：** 步骤 ⑤ MVP 已本地验收通过

## 目标

在约 10 分钟内将产品仓（或 `workspace/<项目名>/`）中的容器暴露到公网。

## 检查清单

- [ ] `docker compose up` 本地 health 200
- [ ] 在 `release/cloudflare/subdomain-registry.example.md` 登记子域名
- [ ] 服务器 Docker 已就绪；网关使用共享网络 `maker-flow`
- [ ] Cloudflare DNS 橙云代理开启

## 部署动作

在 MVP / 产品项目目录：

```bash
export MVP_NAME=idea1
export DOMAIN=idea1.your-domain.com
export DEPLOY_HOST=deploy@your-server
export DEPLOY_PATH=/opt/mvps/idea1
export CONTAINER_PORT=8080   # web-vite: 80
export MVP_SERVICE=api       # 或 web / worker

/path/to/maker-flow/release/deploy/push-and-route.sh
```

## Nginx 网关

`push-and-route.sh` 会在 Docker 网关写入 `conf.d/<MVP_NAME>.conf`，`nginx -t` 通过后 reload。  
手动路径：把 `release/nginx/snippets/mvp-server.conf.example` 渲染进网关 `conf.d/`，在网关容器内执行 `nginx -t` 与 `nginx -s reload`（无需 sudo / 宿主机 Nginx）。

## Cloudflare

添加 A 记录指向服务器 IP（Proxied）。详见 `release/cloudflare/README.md`。

## 验证

```bash
curl -I https://idea1.your-domain.com/health
```

## 回滚

```bash
ssh $DEPLOY_HOST "cd $DEPLOY_PATH && docker compose down"
ssh $DEPLOY_HOST "rm -f /opt/maker-flow/gateway/conf.d/${MVP_NAME}.conf && cd /opt/maker-flow/gateway && docker compose exec -T nginx nginx -t && docker compose exec -T nginx nginx -s reload"
```

并在 DNS 登记表中释放子域名。

## 详细文档

- [release/README.md](../release/README.md)
- [release/nginx/README.md](../release/nginx/README.md)
- [release/deploy/](../release/deploy/)
