# Template shared conventions

**English** · [简体中文](README.zh-CN.md)

For agents assembling in a **product repo** from any template.

## Images (Dockerfile fragments)

Do not invent upstream OS lines ad hoc. Compose from fragments in `templates/images/index.md`:

| Role | Upstream | Source dir |
|------|----------|------------|
| Go build | `golang:1.22-alpine` | `templates/images/go-builder` |
| Go runtime | `alpine:3.20` | `templates/images/go-runtime` |

Inline fragment lines into the app Dockerfile. Do **not** pre-build private `maker-flow/*` tags.

## Env names

`APP_NAME`, `APP_ENV`, `HTTP_ADDR`, `LOG_LEVEL`, `HOST_PORT` (optional **local** compose host map for `docker compose up` acceptance)

## Health

`GET /health` → `{"status":"ok"}`

## Ports

- **Local acceptance:** map with `HOST_PORT` (e.g. `8080:8080` or web `3000:80`) as convenient.
- **Production:** public entry is the Docker Nginx gateway on host **80**; MVP reachability is via network alias `MVP_NAME:CONTAINER_PORT` on `maker-flow` (see `release/`).
