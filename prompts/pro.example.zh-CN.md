# PRO 完整示例 — 待办事项 API

[English](pro.example.md) · **简体中文**

> **用途：** 对照样板，展示「写到什么粒度算够」。  
> **空白骨架：** [`pro.template.md`](pro.template.md)  
> **生成技能：** [`skills/pro-generation.md`](../skills/pro-generation.md)

本示例对应需求：「做一个待办事项迷你 API：支持创建、完成、列表，不需要用户系统。」

---

## 1. 摘要

- **一句话目标：** 提供无鉴权的待办 REST API，支持创建、标记完成、列表查询。
- **MVP 范围：** 单进程内存存储；Gin HTTP；健康检查 + 3 个业务端点；1 天内可验收。
- **不做：**
  - 用户系统 / 登录 / JWT
  - 分页、排序、筛选、标签
  - 持久化数据库、备份
  - 删除待办、批量操作、附件

## 2. 业务流程

1. 调用方 `POST /todos` 创建一条待办（标题必填），服务分配 `id`，默认 `done=false`。
2. 调用方 `GET /todos` 获取当前全部待办（无分页）。
3. 调用方 `POST /todos/{id}/complete` 将指定待办标为完成；已完成再调用仍返回成功（幂等）。
4. 不存在的 `id` → `404`；标题为空 → `400`。
5. 进程重启后内存数据丢失（MVP 可接受）。

## 3. 数据模型

内存结构（无表）：

| 字段 | 类型 | 说明 | 约束 |
|------|------|------|------|
| id | string | 唯一标识 | 非空，服务端生成 |
| title | string | 标题 | 非空，建议 ≤ 200 字符 |
| done | bool | 是否完成 | 默认 `false` |
| created_at | string | 创建时间 RFC3339 | 服务端写入 |

无 `CREATE TABLE`（本 MVP 不持久化）。

## 4. 接口契约

### `GET /health`

- **请求：** 无 body
- **响应：**

```json
{ "status": "ok" }
```

- **错误码：** 无（进程存活即 200）

### `POST /todos`

- **请求：**

```json
{ "title": "买牛奶" }
```

- **响应 `201`：**

```json
{
  "id": "todo_01HXYZ",
  "title": "买牛奶",
  "done": false,
  "created_at": "2026-07-16T05:00:00Z"
}
```

- **错误码：** `400` 标题为空

### `GET /todos`

- **请求：** 无 body
- **响应 `200`：**

```json
{
  "items": [
    {
      "id": "todo_01HXYZ",
      "title": "买牛奶",
      "done": false,
      "created_at": "2026-07-16T05:00:00Z"
    }
  ]
}
```

### `POST /todos/{id}/complete`

- **请求：** 无 body
- **响应 `200`：**

```json
{
  "id": "todo_01HXYZ",
  "title": "买牛奶",
  "done": true,
  "created_at": "2026-07-16T05:00:00Z"
}
```

- **错误码：** `404` 未知 id

## 5. 验收标准

- [ ] `docker compose up --build` 后 `GET /health` 返回 200 且 `status=ok`
- [ ] `POST /todos` 可创建，响应含 `id` / `title` / `done=false`
- [ ] `GET /todos` 能看到刚创建的条目
- [ ] `POST /todos/{id}/complete` 后该条目 `done=true`；重复调用仍 200
- [ ] 空标题创建返回 400；不存在的 id complete 返回 404
- [ ] 未引入用户/鉴权相关代码或依赖

## 6. 模版检索提示（可选，不最终拍板）

- **倾向 apps（1～N）：** `go-api`
- **倾向 patterns（0～N）：** 无
- **镜像 / 运行：** `go-builder` + `go-runtime` Dockerfile 片段（上游 `golang` / `alpine`），`docker compose`
- **复杂度线索：** 低 QPS、无 DB、单服务
