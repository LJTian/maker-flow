# Image templates (Dockerfile fragments)

**English** · [简体中文](index.zh-CN.md)

Agent catalog for **Dockerfile fragments** used when assembling app Dockerfiles.  
Inline the fragment lines into the builder / runtime stages — do **not** pre-build local `maker-flow/*` tags, and do **not** `FROM` a private base image.

## How to use (assembly)

1. Read the selected fragment Dockerfiles under `templates/images/<id>/`.
2. Copy their `FROM` / `RUN apk` / `WORKDIR` (and optional `USER`) lines into the app Dockerfile stages.
3. Append app-specific `COPY` / `go build` / `ENTRYPOINT` / `ENV` / `EXPOSE` below the fragment boundary.
4. Ship only the composed Dockerfile in the product repo — do **not** copy the whole `templates/images/` tree.

App templates (`go-api`, `go-cli`, `go-worker`) already ship composed Dockerfiles as the default.

## Catalog

### go-builder

| Field | Value |
|-------|-------|
| **id** | `go-builder` |
| **path** | `templates/images/go-builder` |
| **upstream** | `golang:1.22-alpine` |
| **stage** | builder |
| **tags** | `go`, `build`, `compile` |
| **when_to_use** | Go multi-stage build (compile stage) |
| **provides** | Go 1.22 toolchain, git, ca-certificates |

### go-runtime

| Field | Value |
|-------|-------|
| **id** | `go-runtime` |
| **path** | `templates/images/go-runtime` |
| **upstream** | `alpine:3.20` |
| **stage** | runtime |
| **tags** | `go`, `runtime`, `alpine` |
| **when_to_use** | Go API / binary runtime stage |
| **provides** | alpine 3.20, ca-certificates, tzdata, wget (healthcheck), non-root `nobody`, `WORKDIR /app` |

## Matching rule

If app template = `go-api` or `go-worker` → compose **both** `go-builder` + `go-runtime`.  
If app template = `go-cli` → same fragments for a containerized CLI.  
No pre-build step — `docker compose up --build` is enough.

## Extend

Add `templates/images/<id>/Dockerfile` (fragment with a clear `# --- end base fragment` marker), register here. Update app Dockerfiles that should use the new fragment.
