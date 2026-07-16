<div align="center">

# Maker Flow

### A personal MVP incubation pipeline

**Heavy infrastructure, light logic.** Kill repeated setup so the path from idea → public validation has minimal friction.

You provide **requirements** and **two approvals**. AI agents execute against a skill library and template set.

<br/>

**English** · [简体中文](README.zh-CN.md)

<br/>

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Workflow](https://img.shields.io/badge/steps-6%20step-purple.svg)](#six-step-pipeline)
[![Stack](https://img.shields.io/badge/template-Go%20Gin-00ADD8.svg)](templates/apps/go-api/)
[![Agent](https://img.shields.io/badge/for-AI%20Agents-22c55e.svg)](AGENTS.md)

<br/>

[Getting started](docs/getting-started.md) · [Template catalog](templates/CATALOG.md) · [Skills catalog](skills/CATALOG.md) · [For agents](AGENTS.md) · [Docs index](docs/README.md)

</div>

---

## Why it exists

For indie builders, friction is often not business code — it is **rebuilding the environment every time**:

| Without Maker Flow | With Maker Flow |
|--------------------|-----------------|
| Pick a stack, write Docker, wire Nginx for every idea | Templates ready to run |
| AI codes first; wrong direction wastes the day | **Two gates**: PRO → MVP |
| Prompts and deploy steps live in your head | Skills encode SOPs; agents follow them |
| Many ideas, repeated plumbing | Focus on validation; public URL in ~10 minutes |

> This is **not one product**. It is a forkable, reusable **MVP factory**.

---

## Six-step pipeline

```mermaid
flowchart LR
    A["① You: one-line requirement"] --> B["② AI: draft PRO"]
    B --> C{"③ Approve PRO"}
    C -->|pass| D["④ AI: match templates · assemble MVP"]
    C -->|revise| B
    D --> E{"⑤ Approve MVP"}
    E -->|pass| F["⑥ Deploy"]
    E -->|iterate| D
    F --> G["🌐 Public traffic"]
```

| Step | You | AI agent |
|:----:|-----|----------|
| 1 | Provide requirement | — |
| 2 | — | Draft PRO (**no code**) |
| 3 | **Approve PRO** | Wait |
| 4 | — | Match templates → assemble under `workspace/` |
| 5 | **Local acceptance** | Fix against PRO |
| 6 | Trigger deploy | Run `release/` scripts |

Two gates are the core design: **align on what first, then how**.

---

## What’s in the repo

```
        ┌─────────────┐
        │  You + AI   │
        └──────┬──────┘
               │
    ┌──────────┼──────────┐
    ▼          ▼          ▼
 skills/   templates/   release/
  (HOW)      (WHAT)      (SHIP)
```

| Module | Path | One line |
|--------|------|----------|
| Skills | [`skills/`](skills/) · [**catalog**](skills/CATALOG.md) | How agents draft PRO, match templates, deploy |
| Templates | [`templates/`](templates/) · [**catalog**](templates/CATALOG.md) | apps + images + patterns |
| AI transport | [`ai-engine/`](ai-engine/) | Optional: Ollama / OpenAI-compatible APIs |
| Release | [`release/`](release/) | Nginx + Cloudflare + deploy scripts |
| Workspace | [`workspace/`](workspace/) | Where assembled MVPs land |

---

## 60-second start

```bash
git clone https://github.com/LJTian/maker-flow.git && cd maker-flow
```

**A — Cursor Agent (recommended)**

1. Open this repo in Cursor
2. Start a chat:

   > Follow @AGENTS.md and docs/workflow.md. I want to build [your idea], starting at step ①.

3. Approve PRO and MVP at steps ③ and ⑤

**B — CLI**

```bash
cp ai-engine/.env.example ai-engine/.env   # set AI_BASE_URL
chmod +x scripts/ai-run.sh
./scripts/ai-run.sh prompts/02-pro-draft.md
```

**C — Smoke a template (no AI)**

```bash
./scripts/build-images.sh   # build Go base images (first time)
cp -r templates/apps/go-api workspace/smoke-test
cd workspace/smoke-test && cp .env.example .env
docker compose up --build
curl http://localhost:8080/health
```

Full guide → **[docs/getting-started.md](docs/getting-started.md)** · [中文](docs/getting-started.zh-CN.md)

---

## Who it’s for

- People with **many ideas** who want fast public feedback
- Anyone using **Cursor / Claude** agents who is tired of prompting from scratch
- Builders who want a **forkable private pipeline**, not another todo demo
- Anyone who buys **heavy infra, light logic**: write plumbing once, run N ideas

---

## After you Star / Fork

| Action | Suggestion |
|--------|------------|
| Star | Track skill / template updates |
| Fork | Make it yours; replace domain and host under `release/` |
| Each new idea | Agent assembles into `workspace/<name>/` |
| Pin agent behavior | Add [AGENTS.md](AGENTS.md) to IDE rules, or `@AGENTS.md` in chat |

---

## Suggested machine roles

| Machine | Role |
|---------|------|
| GPU box (optional) | Inference only; main machine calls via `AI_BASE_URL` |
| Apple Silicon Mac | Dev, acceptance, `workspace/` |
| Cloud VPS | Nginx gateway; port pool `8080–8090` for multiple MVPs |

---

## Docs map

| Humans | AI agents |
|--------|-----------|
| [Getting started](docs/getting-started.md) | [AGENTS.md](AGENTS.md) |
| [Overview](docs/overview.md) | [workflow.md](docs/workflow.md) |
| [Template catalog](templates/CATALOG.md) · [Skills catalog](skills/CATALOG.md) | [agent-bootstrap.md](docs/agent-bootstrap.md) |
| [Docs index](docs/README.md) · [i18n](docs/i18n.md) | |

---

## Acknowledgments

Maker Flow stands on these projects — thanks to maintainers and communities:

| Use | Project | URL |
|-----|---------|-----|
| Go web (`go-api`) | **Gin** | https://github.com/gin-gonic/gin |
| Web UI (`web-vite`) | **Vite** · **React** · **Tailwind CSS** | https://vite.dev/ · https://react.dev/ · https://tailwindcss.com/ |
| CLI (`go-cli`) | **Cobra** | https://github.com/spf13/cobra |
| singleflight (pattern) | **golang.org/x/sync** | https://pkg.go.dev/golang.org/x/sync/singleflight |
| Build base image | **golang** (official) | https://hub.docker.com/_/golang |
| Runtime base image | **Alpine Linux** | https://alpinelinux.org/ · https://hub.docker.com/_/alpine |
| Containers | **Docker** | https://www.docker.com/ |
| Reverse proxy | **Nginx** | https://nginx.org/ |
| DNS / edge SSL | **Cloudflare** | https://www.cloudflare.com/ |
| Optional local LLM | **Ollama** | https://ollama.com/ · https://github.com/ollama/ollama |
| Agent protocol reference | **Model Context Protocol** | https://modelcontextprotocol.io/ |

Missing a credit? Open an Issue / PR.

---

<div align="center">

**If this pipeline helps you, a Star is the best nod to “heavy infrastructure.”**

<br/>

[Get started](docs/getting-started.md) · [Fork your factory](https://github.com/LJTian/maker-flow/fork)

</div>

## License

MIT — use freely and adapt as needed.
