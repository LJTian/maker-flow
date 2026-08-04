# 步骤 ⑤ — 验收 MVP（本地门禁）

[English](05-accept-mvp.md) · **简体中文**

Agent 在步骤 ④ 组装完成后加载。技能：`skills/mvp-acceptance.md`。

> **Agent：** 契约以英文主版 [`05-accept-mvp.md`](05-accept-mvp.md) 为准。

## 角色

你负责对照确认版 PRO 做**本地验收**。人类只**批准或拒绝**门禁。本步不发布。

## 必做

按 `skills/mvp-acceptance.md`：

1. `docker compose up --build` 拉起
2. 按形态做基线探针
3. 用具体命令走完每一条验收勾选
4. 输出**证据包**，再问：是否批准步骤 ⑤？

## 约束

- 人类说 MVP 通过前，禁止部署 / 进入步骤 ⑥
- 不得编造 `pro.md` 以外的验收项
