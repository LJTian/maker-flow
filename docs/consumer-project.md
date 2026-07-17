# Consumer project (parallel directories)

**English** · [简体中文](consumer-project.zh-CN.md)

Use this when **maker-flow is the public tool repo** and each MVP lives in a **separate product repository** (often private).

Recommended layout: **parallel directories** — not submodule, not copying the whole factory into the product repo.

---

## Two repos, two roles

| Repo | Visibility | Contains |
|------|------------|----------|
| **maker-flow** | Public (tool) | `skills/`, `templates/`, `prompts/`, `release/`, docs |
| **Your MVP** (e.g. `my-todo`) | Private (product) | `pro.md`, assembled apps, `docker-compose.yml`, short `AGENTS.md` |

**Rule:** Never commit MVP business code into maker-flow. Others who clone maker-flow must not receive your ideas. Real MVPs go in product repos (`maker-flow new <name>`).

---

## Recommended directory layout

```text
~/.maker-flow/                # factory (maker-flow install)
~/projects/
  ├── my-todo/                # private git — one repo per idea
  ├── my-saas/
  └── ...
```

Git clone of maker-flow is for **contributors** only. End users: `curl … | bash` or `./scripts/install.sh`.

### Minimal product repo

```text
my-todo/
├── AGENTS.md                 # from maker-flow init / new
├── requirement.md            # step ①
├── pro.md                    # step ③ (optional; maker-flow init --with-pro)
├── docker-compose.yml        # after assembly (multi-app: root compose)
├── api/                      # copied from templates/apps/go-api + business logic
└── web/                      # optional: from templates/apps/web-vite
```

**Do not copy** into the product repo: entire `skills/`, `templates/`, `release/` trees (read them via `MAKER_FLOW_ROOT`).

---

## Environment variable

Usually **not required** — product `AGENTS.md` defaults to `MAKER_FLOW_ROOT=~/.maker-flow` (portable). Resolve with `maker-flow root`.

Optional override for a custom install directory only:

```bash
export MAKER_FLOW_ROOT=/custom/path/to/maker-flow
```

Agent contract:

- **Read (read-only):** `$MAKER_FLOW_ROOT/skills/`, `$MAKER_FLOW_ROOT/templates/`, `$MAKER_FLOW_ROOT/prompts/`, `$MAKER_FLOW_ROOT/release/`
- **Write:** current product repo only (`./`, `./api/`, `./web/`, …)
- **MUST NOT** modify files under `$MAKER_FLOW_ROOT`

---

## Bootstrap a new MVP

### One command (recommended)

```bash
# Install factory once (if needed)
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash

# New product repo → ~/projects/my-todo
maker-flow new my-todo
cd ~/projects/my-todo
```

### Or step by step

```bash
./scripts/install.sh              # from a git clone; installs to ~/.maker-flow
maker-flow init my-todo           # or: maker-flow init my-todo --with-pro
cd ~/projects/my-todo
```

### Manual (no CLI)

```bash
# 1. Factory
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash
# factory lives at ~/.maker-flow

# 2. Product repo
maker-flow init my-todo
cd ~/projects/my-todo
```

---

## Six-step mapping (consumer mode)

| Step | Where it happens |
|------|------------------|
| ① Requirement | Product repo (chat or local notes) |
| ② Draft PRO | Agent reads `$MAKER_FLOW_ROOT/skills/pro-generation.md` |
| ③ Confirm PRO | Write `pro.md` in **product repo** |
| ④ Assemble | Copy from `$MAKER_FLOW_ROOT/templates/...` into **product repo**; follow `mvp-assembly` (output = product root) |
| ⑤ Accept | `docker compose up --build` in **product repo** |
| ⑥ Deploy | Agent asks where to publish (Pages / Vercel / VPS); see `skills/deploy.md` — humans do not run a deploy CLI |

Go apps: compose Dockerfile fragments from `$MAKER_FLOW_ROOT/templates/images/` (already done in app templates). No pre-build step.

---

## Multi-app layout

```text
my-todo/
├── pro.md
├── docker-compose.yml        # orchestrates api + web services
├── api/                      # go-api
└── web/                      # web-vite
```

Set `VITE_API_BASE_URL` before building `web/` when pairing with `go-api`.

---

## Comparison (why parallel dirs)

| Approach | Product repo cleanliness | Version pin | Daily friction |
|----------|-------------------------|-------------|----------------|
| **Parallel dirs (recommended)** | Best | Manual (note factory commit) | Low |
| Submodule `vendor/maker-flow` | Good | Git submodule SHA | Medium |
| Copy whole factory into product | Poor | Frozen copy | Low |

Use **submodule** only when you need collaborators or CI to pin an exact factory commit on one product.

---

## Security checklist

- [ ] Product repos are **private** if MVPs must stay secret
- [ ] No `.env` with secrets in maker-flow
- [ ] No real subdomain/registry data committed to maker-flow `release/`
- [ ] `pro.md` with sensitive scope stays in product repo only

---

## Related

- Product `AGENTS.md` template: [`AGENTS.consumer.example.md`](../AGENTS.consumer.example.md)
- Factory agent entry: [`AGENTS.md`](../AGENTS.md)
- Workflow: [`workflow.md`](workflow.md)
- Getting started (factory-local): [`getting-started.md`](getting-started.md)
