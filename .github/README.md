# GitHub Actions

Factory CI lives in [`workflows/ci.yml`](workflows/ci.yml).

| Job | What it does |
|-----|----------------|
| `factory-check` | `scripts/check.sh` (shell syntax, Dockerfile contracts, layouts, CLI help) |
| `pattern-tests` | For each `templates/patterns/*` with Go sources: `go mod tidy && go test ./...` |
| `compose-build` | Matrix: `docker compose build` for go-api / go-worker / web-vite; `docker build` for go-cli |

This is **factory CI only**. It does **not** publish product MVPs (step 6 stays chat + `skills/publish-*.md`).
