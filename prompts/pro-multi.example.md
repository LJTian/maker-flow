# PRO full sample — Multi-app Todo (Web SPA + API)

**English** · [简体中文](pro-multi.example.zh-CN.md)

> **Purpose:** Reference sample showing granularity for a multi-app (frontend + backend) MVP.  
> **Blank skeleton:** [`pro.template.md`](pro.template.md)  
> **Generation skill:** [`skills/pro-generation.md`](../skills/pro-generation.md)

This sample matches the requirement: “Build a todo app with a React frontend and Go backend, supporting create, mark complete, and list — no user system.”

---

## 1. Summary

- **One-sentence goal:** Provide an unauthenticated full-stack todo application with a React SPA frontend and a Go REST API backend.
- **MVP scope:** Single-process in-memory store for backend; React + Tailwind for frontend; health check + 3 business endpoints; acceptable within 1 day.
- **Out of scope:**
  - User system / login / JWT
  - Pagination, sorting, filtering, tags
  - Persistent database, backups
  - Delete todo, bulk ops, attachments

## 2. Business flow

1. **Frontend:** User opens the web app and sees a list of todos fetched from `GET /api/v1/todos`.
2. **Frontend:** User types a title and clicks "Add". The frontend calls `POST /api/v1/todos` and then refreshes the list.
3. **Frontend:** User clicks a checkbox on a todo. The frontend calls `POST /api/v1/todos/{id}/complete` and updates the UI state to checked.
4. **Backend:** Process restart loses in-memory data (acceptable for MVP).

## 3. Data model

In-memory structure (no tables):

| Field | Type | Description | Constraints |
|-------|------|-------------|-------------|
| id | string | Unique id | Non-empty, server-generated |
| title | string | Title | Non-empty, recommend ≤ 200 chars |
| done | bool | Completed? | Default `false` |
| created_at | string | Created time RFC3339 | Server-written |

## 4. API / interface contract

### `GET /health` (API)
- **Response:** `{ "status": "ok" }`

### `POST /api/v1/todos` (API)
- **Request:** `{ "title": "Buy milk" }`
- **Response `201`:** `{ "id": "todo_01", "title": "Buy milk", "done": false, "created_at": "..." }`

### `GET /api/v1/todos` (API)
- **Response `200`:** `{ "items": [ { "id": "todo_01", "title": "Buy milk", "done": false, "created_at": "..." } ] }`

### `POST /api/v1/todos/{id}/complete` (API)
- **Response `200`:** `{ "id": "todo_01", "title": "Buy milk", "done": true, "created_at": "..." }`

## 5. Acceptance criteria

- [ ] `docker compose up --build` successfully starts both `api` and `web` services.
- [ ] Backend `GET /health` returns 200 `status=ok`.
- [ ] Frontend is accessible via `http://localhost:3000` (or similar) and loads without JS errors.
- [ ] Creating a todo on the frontend adds it to the list visually and persists it to the backend memory.
- [ ] Clicking complete on a frontend item updates its visual state to done and sends the correct API call.
- [ ] No user/auth-related code introduced.
- [ ] Cross-Origin Resource Sharing (CORS) is properly configured on the backend so the frontend can call it.

## 6. Template retrieval hints (optional — not final picks)

- **Preferred apps (1–N):** `go-api` (backend), `web-vite` (frontend)
- **Preferred patterns (0–N):** none
- **Images / runtime:** `go-builder`, `go-runtime`, `node-builder`, `node-runtime` (Nginx for static assets).
- **Complexity clues:** Requires `docker-compose.yml` to orchestrate both services and handle frontend-to-backend proxying or CORS.
