# templates/

Searchable scaffold set for step 4. Agents MUST NOT invent a new project layout when a catalog entry matches.

## Catalog

**Entry point:** `index.md` (tags, `when_to_use`, path).

| id | path | port |
|----|------|------|
| `go-api` | `go-api/` | 8080 |

## What each template provides

- Locked `Dockerfile` + `docker-compose.yml`
- Runnable demo source
- `.env.example` + README

Shared conventions: `shared/`.

## Agent write rule

1. Match via `index.md` (`skills/template-matching.md`).
2. Copy selected template → `workspace/<name>/`.
3. Implement PRO only (`skills/mvp-assembly.md`).

Do not leave assembled MVPs inside `templates/`.
