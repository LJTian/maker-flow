# Template shared conventions

For agents assembling under `workspace/` from any template.

## Images

| Lang | Build | Runtime |
|------|-------|---------|
| Go | `golang:1.22-alpine` | `alpine:3.20` |

## Env names

`APP_NAME`, `APP_ENV`, `HTTP_ADDR`, `LOG_LEVEL`, `HOST_PORT` (compose host map → release port pool)

## Health

`GET /health` → `{"status":"ok"}`

## Ports

Server pool `8080–8090`; see `release/`.
