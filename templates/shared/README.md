# Template shared conventions

**English** · [简体中文](README.zh-CN.md)

For agents assembling under `workspace/` from any template.

## Images (inheritance)

Do not hardcode upstream OS in app Dockerfiles. Use tags from `templates/images/index.md`:

| Role | Local tag | Source dir |
|------|-----------|------------|
| Go build | `maker-flow/go-builder:1.22` | `templates/images/go-builder` |
| Go runtime | `maker-flow/go-runtime:1.22` | `templates/images/go-runtime` |

Build: `./scripts/build-images.sh`

## Env names

`APP_NAME`, `APP_ENV`, `HTTP_ADDR`, `LOG_LEVEL`, `HOST_PORT` (compose host map → release port pool)

## Health

`GET /health` → `{"status":"ok"}`

## Ports

Server pool `8080–8090`; see `release/`.
