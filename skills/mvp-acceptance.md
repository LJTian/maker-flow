# MVP acceptance skill

**English** · [简体中文](mvp-acceptance.zh-CN.md)

**Step:** 5 — local verification before the human gate  
**Prerequisite:** step 4 assembled a product repo; confirmed PRO at `pro.md` (same shape as `prompts/pro.template.md`)  
**Skill id:** `mvp-acceptance`

## Goal

Prove the assembled MVP meets **every** acceptance criterion in `pro.md`, with reproducible commands and evidence. The **human** only approves or rejects at the gate. Agent **MUST NOT** treat a green `/health` alone as “MVP passed.”

## Hard rules

- **MUST** run from the **product repo root** (not the factory).
- **MUST** walk every checkbox under PRO § Acceptance criteria (and any API/CLI contracts those checkboxes imply).
- **MUST** use containers for bring-up (`docker compose up --build`); do not require a host Go/Node toolchain.
- **MUST** present a pass/fail evidence summary, then **wait** for explicit human approval.
- **MUST NOT** start step 6 / any publish action until the human approves.
- **MUST NOT** invent acceptance items that are not in the PRO.
- On fail: fix in step 4, or return to step 3 if the PRO scope is wrong — do not paper over with deploy.

## Procedure

### 1. Orient

1. Confirm cwd is the product repo (`pro.md` + compose present).
2. Read `pro.md` — especially § Acceptance criteria and § API / interfaces.
3. Detect **shape** from layout / compose services:

| Shape | Signals | Default local check |
|-------|---------|---------------------|
| API (`go-api`) | service `api`, listen `:8080` | `curl` host `:8080` |
| SPA (`web-vite`) | service `web`, host port often `3000` | `curl` `/` or `/health` on host port |
| CLI (`go-cli`) | no long-running public port | `docker compose run --rm <svc> …` |
| Worker (`go-worker`) | service `worker`, no public URL | logs / process stays up; optional admin ping if PRO defines one |
| Multi-app | root compose with `api` + `web` (+ …) | verify **each** mapped app |

### 2. Bring up

```bash
cd ~/projects/<name>   # product repo
cp -n .env.example .env 2>/dev/null || true
docker compose up --build -d
docker compose ps
```

If multi-app uses nested compose files, follow the product README / root compose the assembly step left behind.

Wait until healthchecks are healthy (or a short retry loop on `curl`). On build/start failure → treat as **fail**, collect logs (`docker compose logs --tail=200`), iterate step 4.

### 3. Baseline smoke (shape-aware)

Run only what applies:

**API**

```bash
curl -sf "http://localhost:${HOST_PORT:-8080}/health"
# expect JSON status ok (or PRO-defined health body)
```

**SPA / static**

```bash
curl -sf "http://localhost:${HOST_PORT:-3000}/health" \
  || curl -sf -o /dev/null -w "%{http_code}" "http://localhost:${HOST_PORT:-3000}/"
# expect 200 (web-vite ships /health; fallback is document root)
```

**CLI**

```bash
docker compose run --rm <cli-service> --help
# plus one PRO happy-path command
```

**Worker**

```bash
docker compose logs --tail=50 <worker-service>
# process running; no crash loop; PRO-defined side effects if any
```

**Multi-app / split frontend+API (local)**

- Start the full root stack.
- Hit API `/health` **and** web `/` or `/health`.
- If `VITE_API_BASE_URL` is set, confirm the UI can reach the API (browser or a documented health ping). Do not require production CORS values at step 5 — only local compose values.

### 4. Walk PRO acceptance criteria

For each `- [ ]` item in PRO §5:

1. Turn it into **one concrete command** (curl / compose run / log assertion).
2. Execute it.
3. Record: criterion text · command · result (pass/fail) · short evidence (status code, snippet).

Cover implied contracts from § API when a checkbox references them (create → list → mutate, error codes, etc.). Prefer the PRO’s own examples over inventing payloads.

### 5. Evidence package (required before asking the human)

Present in chat:

1. Product path + compose services observed
2. Baseline smoke results
3. Table or checklist of every acceptance criterion with pass/fail
4. Notable logs on any failure
5. Explicit ask: **Approve step 5?** (pass → step 6; fail → agent iterates)

Example shape:

```text
Step 5 evidence — <PRODUCT_NAME>
Baseline: api /health 200 ok · web /health 200 ok
Acceptance:
  [x] create todo via POST /todos → 201
  [x] list includes created item → 200
  [ ] complete marks done → FAIL curl 500 (see logs)
Recommendation: iterate step 4 (handler panic) — not a PRO scope change.
Approve step 5? (yes / no + notes)
```

### 6. Human gate + feedback

| Human says | Agent does |
|------------|------------|
| Approved / pass | Stop verification; proceed only when human asks for publish → step 6 (`skills/deploy.md`) |
| Rejected — bug / missing behavior | Stay in / return to step 4; re-run this skill after fixes |
| Rejected — wrong scope / wrong PRO | Return to step 3; revise PRO before more code |
| Unclear | Ask which criteria failed; do not deploy |

## MUST NOT

- MUST NOT skip criteria that are “hard to automate” — document a manual browser step and ask the human to confirm that item only.
- MUST NOT deploy, open firewall ports, or change DNS at this step.
- MUST NOT mark the gate passed based on agent judgment alone.
- MUST NOT require the human to invent curl scripts — agent prepares them.

## Prompt

Stage prompt: [`prompts/05-accept-mvp.md`](../prompts/05-accept-mvp.md)

## Related

- Gate definition: [`docs/workflow.md`](../docs/workflow.md) step 5
- Assembly quality bar: [`mvp-assembly.md`](mvp-assembly.md)
- After approval: [`deploy.md`](deploy.md)
