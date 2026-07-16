# Image templates (base images)

Agent catalog for **inheritance-style** Docker bases. App templates (`go-api`, …) use `FROM maker-flow/<id>:<tag>` only — they MUST NOT redeclare OS packages already in the base.

## Build all Go bases

```bash
./scripts/build-images.sh
```

## Catalog

### go-builder

| Field | Value |
|-------|-------|
| **id** | `go-builder` |
| **path** | `templates/images/go-builder` |
| **tag** | `maker-flow/go-builder:1.22` |
| **tags** | `go`, `build`, `compile` |
| **when_to_use** | Go multi-stage build (compile stage) |
| **provides** | Go 1.22 toolchain, git, ca-certificates |

### go-runtime

| Field | Value |
|-------|-------|
| **id** | `go-runtime` |
| **path** | `templates/images/go-runtime` |
| **tag** | `maker-flow/go-runtime:1.22` |
| **tags** | `go`, `runtime`, `alpine` |
| **when_to_use** | Go API / binary runtime stage |
| **provides** | alpine 3.20, ca-certificates, tzdata, wget (healthcheck), non-root `nobody`, `WORKDIR /app` |

## Matching rule

If app template = `go-api` → depend on **both** `go-builder` + `go-runtime`.  
Build bases **before** `docker compose up --build` in `workspace/`.

## Extend

Add `templates/images/<id>/Dockerfile`, register here, update `scripts/build-images.sh`.
