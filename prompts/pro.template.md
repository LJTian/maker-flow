# PRO blank skeleton

**English** · [简体中文](pro.template.zh-CN.md)

> **Purpose:** Fill this structure for step 2 output and step 3 confirmation paste.  
> **Authoritative sections:** [`skills/pro-generation.md`](../skills/pro-generation.md)  
> **Full sample:** [`pro.example.md`](pro.example.md)

After filling: draft → confirm at step 3 → write into `pro.md` in the **product repo** (or `03-pro-confirmed.example.md` for factory examples).

---

## 1. Summary

- **One-sentence goal:** (…)
- **MVP scope:** (boundary finishable in 1–2 days)
- **Out of scope:**
  - (…)
  - (…)

## 2. Business flow

1. (Main-path step)
2. (…)
3. (Boundaries: auth? data ownership? idempotency?)

## 3. Data model

| Field | Type | Description | Constraints |
|-------|------|-------------|-------------|
| id | … | … | … |

Optional (when persistence is required):

```sql
-- CREATE TABLE …
```

If no persistence, write “in-memory / no tables” and describe lifecycle.

## 4. API / interface contract

> Single app: the endpoints below are enough. Multi-app: split subsections by app (e.g. API / Worker / CLI).

### (app name or unified API)

#### `METHOD /path`

- **Request:**

```json
{}
```

- **Response:**

```json
{}
```

- **Main error codes:** (e.g. 400 / 404 / 500)

At least: health check (if applicable) + 2–4 core business interfaces/commands.

## 5. Acceptance criteria

Checkboxes for step 5:

- [ ] (One observable, reproducible item)
- [ ] (…)
- [ ] (…)

## 6. Template retrieval hints (optional — not final picks)

- **Preferred apps (1–N):** (e.g. `go-api`; or `go-api` + `go-worker`)
- **Preferred patterns (0–N):** (e.g. `retry-backoff`; or “none”)
- **Images / runtime:** (e.g. `go-builder` + `go-runtime` fragments, `docker compose`)
- **Complexity clues:** (expected QPS, DB needed?, etc.)
