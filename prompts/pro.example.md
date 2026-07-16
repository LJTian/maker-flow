# PRO full sample — Todo API

**English** · [简体中文](pro.example.zh-CN.md)

> **Purpose:** Reference sample showing “enough” granularity.  
> **Blank skeleton:** [`pro.template.md`](pro.template.md)  
> **Generation skill:** [`skills/pro-generation.md`](../skills/pro-generation.md)

This sample matches the requirement: “Build a mini todo API: create, complete, and list — no user system.”

---

## 1. Summary

- **One-sentence goal:** Provide an unauthenticated todo REST API supporting create, mark complete, and list.
- **MVP scope:** Single-process in-memory store; Gin HTTP; health check + 3 business endpoints; acceptable within 1 day.
- **Out of scope:**
  - User system / login / JWT
  - Pagination, sorting, filtering, tags
  - Persistent database, backups
  - Delete todo, bulk ops, attachments

## 2. Business flow

1. Client `POST /todos` creates a todo (title required); service assigns `id`, default `done=false`.
2. Client `GET /todos` returns all current todos (no pagination).
3. Client `POST /todos/{id}/complete` marks the todo complete; repeating on an already-complete todo still succeeds (idempotent).
4. Unknown `id` → `404`; empty title → `400`.
5. Process restart loses in-memory data (acceptable for MVP).

## 3. Data model

In-memory structure (no tables):

| Field | Type | Description | Constraints |
|-------|------|-------------|-------------|
| id | string | Unique id | Non-empty, server-generated |
| title | string | Title | Non-empty, recommend ≤ 200 chars |
| done | bool | Completed? | Default `false` |
| created_at | string | Created time RFC3339 | Server-written |

No `CREATE TABLE` (this MVP does not persist).

## 4. API / interface contract

### `GET /health`

- **Request:** no body
- **Response:**

```json
{ "status": "ok" }
```

- **Error codes:** none (process alive → 200)

### `POST /todos`

- **Request:**

```json
{ "title": "Buy milk" }
```

- **Response `201`:**

```json
{
  "id": "todo_01HXYZ",
  "title": "Buy milk",
  "done": false,
  "created_at": "2026-07-16T05:00:00Z"
}
```

- **Error codes:** `400` empty title

### `GET /todos`

- **Request:** no body
- **Response `200`:**

```json
{
  "items": [
    {
      "id": "todo_01HXYZ",
      "title": "Buy milk",
      "done": false,
      "created_at": "2026-07-16T05:00:00Z"
    }
  ]
}
```

### `POST /todos/{id}/complete`

- **Request:** no body
- **Response `200`:**

```json
{
  "id": "todo_01HXYZ",
  "title": "Buy milk",
  "done": true,
  "created_at": "2026-07-16T05:00:00Z"
}
```

- **Error codes:** `404` unknown id

## 5. Acceptance criteria

- [ ] After `docker compose up --build`, `GET /health` returns 200 with `status=ok`
- [ ] `POST /todos` creates an item; response includes `id` / `title` / `done=false`
- [ ] `GET /todos` shows the newly created item
- [ ] After `POST /todos/{id}/complete`, that item has `done=true`; repeat call still returns 200
- [ ] Empty title create returns 400; complete on unknown id returns 404
- [ ] No user/auth-related code or dependencies introduced

## 6. Template retrieval hints (optional — not final picks)

- **Preferred apps (1–N):** `go-api`
- **Preferred patterns (0–N):** none
- **Images / runtime:** `maker-flow/go-builder` + `go-runtime`, `docker compose`
- **Complexity clues:** low QPS, no DB, single service
