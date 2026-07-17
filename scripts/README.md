# scripts/

**English** · [简体中文](README.zh-CN.md)

| Script / CLI | Purpose |
|--------------|---------|
| **`install.sh`** | Install factory to `~/.maker-flow`, link `maker-flow` CLI, auto PATH hint |
| **`maker-flow`** | CLI: `new`, `install`, `upgrade`, `init`, `root`, `version` (+ agent-internal `deploy`) |
| `check.sh` | Factory CI / local sanity checks |

## Quick start

```bash
# Remote (no git clone)
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash

# New MVP → ~/projects/my-todo
maker-flow new my-todo
cd ~/projects/my-todo
```

## Commands (human)

```bash
maker-flow new <name>              # install if needed + init (recommended)
maker-flow init <name>             # product repo only
maker-flow init <name> --with-pro  # also create draft pro.md
maker-flow upgrade                 # update ~/.maker-flow
maker-flow root
maker-flow version
./scripts/check.sh                 # local / CI checks
```

**Publish / go-live:** talk to the agent after local acceptance — choose Cloudflare Pages / GitHub Pages / Vercel / VPS. Humans do **not** need `maker-flow deploy` (`skills/deploy.md`).

## Agent-internal

`maker-flow deploy …` wraps `../release/deploy/push-and-route.sh` for the **VPS gateway** target only. Prefer invoking it from the agent after the human confirms that target.

From a git clone: `./scripts/install.sh` then `maker-flow new …`

Guide: [`../docs/consumer-project.md`](../docs/consumer-project.md)  
Publish targets: [`../release/publish/`](../release/publish/)  
Dockerfile fragments: `../templates/images/` (inline when assembling; no pre-build step).
