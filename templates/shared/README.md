# Template shared conventions

**English** · [简体中文](README.zh-CN.md)

For agents assembling under `workspace/` from any template.

## Images (Dockerfile fragments)

Do not invent upstream OS lines ad hoc. Compose from fragments in `templates/images/index.md`:

| Role | Upstream | Source dir |
|------|----------|------------|
| Go build | `golang:1.22-alpine` | `templates/images/go-builder` |
| Go runtime | `alpine:3.20` | `templates/images/go-runtime` |

Inline fragment lines into the app Dockerfile. Do **not** pre-build private `maker-flow/*` tags.

## Env names

`APP_NAME`, `APP_ENV`, `HTTP_ADDR`, `LOG_LEVEL`, `HOST_PORT` (compose host map → release port pool)

## Health

`GET /health` → `{"status":"ok"}`

## Ports

Server pool `8080–8090`; see `release/`.
