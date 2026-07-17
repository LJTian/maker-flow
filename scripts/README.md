# scripts/

**English** · [简体中文](README.zh-CN.md)

| Script / CLI | Purpose |
|--------------|---------|
| **`install.sh`** | Install factory to `~/.maker-flow`, link `maker-flow` CLI, auto PATH hint |
| **`maker-flow`** | CLI: `new`, `install`, `upgrade`, `init`, `root`, `version` |

## Quick start

```bash
# Remote (no git clone)
curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash

# New MVP → ~/projects/my-todo
maker-flow new my-todo
cd ~/projects/my-todo
```

## Commands

```bash
maker-flow new <name>              # install if needed + init (recommended)
maker-flow init <name>             # product repo only
maker-flow init <name> --with-pro  # also create draft pro.md
maker-flow upgrade                 # update ~/.maker-flow
maker-flow root
maker-flow version
```

From a git clone: `./scripts/install.sh` then `maker-flow new …`

Guide: [`../docs/consumer-project.md`](../docs/consumer-project.md)

Deploy: `../release/deploy/`.  
Dockerfile fragments: `../templates/images/` (inline when assembling; no pre-build step).
