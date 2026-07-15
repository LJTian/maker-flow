# Architecture (agent map)

## Principles

- Templates + skills are the product; agent fills business logic only.
- Two human gates: PRO before code; MVP before deploy.
- Agent writes artifacts to defined paths; do not scatter files.

## Component map

```mermaid
flowchart TB
    Human[Human]
    Agent[Agent]
    Skills[skills/]
    Templates[templates/]
    Workspace[workspace/]
    Release[release/]

    Human -->|① requirement| Agent
    Agent -->|② PRO| Human
    Human -->|③ approve PRO| Agent
    Agent --> Skills
    Agent --> Templates
    Agent -->|④ assemble| Workspace
    Workspace -->|⑤ approve MVP| Human
    Workspace -->|⑥ deploy| Release
    Release --> MVP[public MVP]
```

## Directories

| Path | Agent use |
|------|-----------|
| `skills/` | Authoritative HOW for each step |
| `templates/` | Searchable scaffolds; catalog = `index.md` |
| `prompts/` | Stage input templates |
| `workspace/` | ONLY place for assembled MVP code |
| `release/` | Deploy primitives for step 6 |
| `ai-engine/` | Optional remote LLM config for `scripts/ai-run.sh` |
| `docs/` | Workflow / architecture contracts |

## Step → path map

| Step | Actor | Paths |
|------|-------|-------|
| 1 | Human | requirement text |
| 2 | Agent | `skills/pro-generation.md`, `prompts/02-pro-draft.md` |
| 3 | Human | confirmed PRO file |
| 4 | Agent | `skills/template-matching.md`, `skills/mvp-assembly.md`, `templates/index.md`, `workspace/<name>/` |
| 5 | Human | `workspace/<name>/` + PRO acceptance |
| 6 | Agent | `skills/deploy.md`, `release/` |

## Related

- `docs/workflow.md`
- `docs/agent-bootstrap.md`
- `docs/getting-started.md` (human)
- `AGENTS.md`
