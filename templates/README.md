# 模版集

可独立复制、供 AI 在步骤 ④ **检索选用** 的预制工程。

## 目录

检索入口：[index.md](index.md)

| 模版 | 路径 | 默认端口 |
|------|------|----------|
| Go API | [go-api/](go-api/) | 8080 |

## 模版内容

每个模版包含：

- **镜像** — `Dockerfile` 锁定构建与运行环境
- **demo 源码** — 可运行的最小 API 骨架
- **文档** — `README.md`、`.env.example`

## 共享约定

见 [shared/](shared/)。

## 手动复制（不经过 AI）

```bash
cp -r templates/go-api workspace/my-idea-1
cd workspace/my-idea-1
cp .env.example .env
docker compose up --build
```

AI 组装时默认输出到 `workspace/<项目名>/`，见 `skills/mvp-assembly.md`。
