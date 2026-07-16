# 模版集检索目录

> **给 AI：** 步骤 ④ 选型前 **MUST** 先读本文件，再读明细 [`index.md`](index.md) / [`images/index.md`](images/index.md)。  
> **给人类：** 一眼看清仓库里已有哪些模版；点链接看详情。

---

## 速览

| 类别 | 数量 | 检索明细 |
|------|:----:|----------|
| 应用模版 (apps) | 1 | [index.md](index.md) |
| 镜像基座 (images) | 2 | [images/index.md](images/index.md) |
| 模式库 (patterns) | 0 | 规划中（apps + patterns 分层） |

---

## 应用模版

| id | 路径 | 标签 | 何时用 | 依赖镜像 |
|----|------|------|--------|----------|
| `go-api` | [`go-api/`](go-api/) | `go` `gin` `rest` `api` `docker` | Go + Gin REST API MVP | `go-builder` + `go-runtime` |

**默认端口：** `8080` · **文档：** [go-api/README.md](go-api/README.md)

---

## 镜像基座（继承式）

应用 Dockerfile 只写业务层，通过 `FROM maker-flow/<id>:<tag>` 继承。

| id | 本地标签 | 路径 | 用途 |
|----|----------|------|------|
| `go-builder` | `maker-flow/go-builder:1.22` | [`images/go-builder/`](images/go-builder/) | Go 编译阶段 |
| `go-runtime` | `maker-flow/go-runtime:1.22` | [`images/go-runtime/`](images/go-runtime/) | 运行阶段 |

```bash
./scripts/build-images.sh   # 组装 / compose 前必做（本机首次或基座变更后）
```

---

## 选型口令（Agent）

```
需要 REST API（Go）？ → go-api（先 build images）
需要其它形态？       → 本目录尚无对应模版，先扩展 templates/ 并登记本文件 + index.md
```

字段级契约、决策树、新增规范 → **[`index.md`](index.md)**

---

## 登记规则

新增模版时 **同时** 更新：

1. 本文件（速览表）
2. [`index.md`](index.md) 或 [`images/index.md`](images/index.md)
3. [`skills/template-matching.md`](../skills/template-matching.md) 匹配表（若影响选型）
