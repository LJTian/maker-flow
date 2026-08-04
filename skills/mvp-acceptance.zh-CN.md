# MVP 验收技能

[English](mvp-acceptance.md) · **简体中文**

**步骤：** ⑤ — 人类门禁前的本地验证  
**前置：** 步骤 ④ 已在产品仓组装完成；确认版 PRO 在 `pro.md`  
**技能 id：** `mvp-acceptance`

> **Agent：** 契约以英文主版 [`mvp-acceptance.md`](mvp-acceptance.md) 为准。本文件仅供人类阅读。

## 目标

用可复现命令与证据，证明组装出的 MVP 满足 `pro.md` **每一条**验收标准。人类只在门禁上批准/拒绝。Agent **不得**把单独的 `/health` 通过当成「MVP 已过」。

## 硬规则

- **必须**在**产品仓根目录**执行（不是工厂仓）
- **必须**逐条走完 PRO「验收标准」勾选（以及勾选隐含的 API/CLI 契约）
- **必须**用容器拉起（`docker compose up --build`）
- **必须**先给出通过/失败证据摘要，再**等待**人类明确批准
- 人类批准前 **禁止**进入步骤 ⑥ / 任何发布动作
- 失败：回步骤 ④ 修代码，或 PRO 范围不对则回步骤 ③ — 禁止用上线糊弄

## 流程摘要

1. 读 `pro.md`，根据 compose 判定形态（API / SPA / CLI / Worker / 多 app）
2. `docker compose up --build -d` 并检查 `ps` / 日志
3. 按形态做基线探针（端口、`/health`、`compose run`、worker 日志等）
4. 把每条验收标准落到一条具体命令并执行
5. 输出证据包，询问：**是否批准步骤 ⑤？**
6. 通过 → 等人发起发布再进 ⑥；失败 → ④ 或 ③

多 app / 本地前后端联调：根 compose 全开；API 与 web 都要探；`VITE_API_BASE_URL` 用本地值即可，本步不要求生产 CORS。

## 配套

- Prompt：[`prompts/05-accept-mvp.md`](../prompts/05-accept-mvp.md)
- 英文全文（Agent 必读）：[`mvp-acceptance.md`](mvp-acceptance.md)
