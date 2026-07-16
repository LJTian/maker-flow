# Deploy skill

**English** · [简体中文](deploy.zh-CN.md)

**Step:** 6 — production deploy  
**Prerequisite:** step 5 local MVP acceptance passed

## Goal

Expose the container(s) in `workspace/<project-name>/` to the public internet within about 10 minutes.

## Checklist

- [ ] `docker compose up` local health returns 200
- [ ] Subdomain and port registered in `release/cloudflare/subdomain-registry.example.md`
- [ ] Server Docker + Nginx ready
- [ ] Cloudflare DNS orange-cloud proxy enabled

## Deploy actions

From the MVP project directory:

```bash
export MVP_NAME=idea1
export MVP_PORT=8080
export DOMAIN=idea1.your-domain.com
export DEPLOY_HOST=deploy@your-server
export DEPLOY_PATH=/opt/mvps/idea1

/path/to/maker-flow/release/deploy/push-and-route.sh
```

## Nginx

Copy `release/nginx/snippets/mvp-server.conf.example`, then set `server_name` and `proxy_pass` port.

## Cloudflare

Add an A record to the server IP (Proxied). Details: `release/cloudflare/README.md`.

## Verify

```bash
curl -I https://idea1.your-domain.com/health
```

## Rollback

```bash
ssh $DEPLOY_HOST "cd $DEPLOY_PATH && docker compose down"
```

Also free the port and subdomain in the Nginx / DNS registry.

## Further reading

- [release/README.md](../release/README.md)
- [release/deploy/](../release/deploy/)
