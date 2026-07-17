# 技能库检索目录

[English](CATALOG.md) · **简体中文**

> **给 AI：** 执行对应步骤前 **MUST** 先读本文件定位技能，再打开技能全文。  
> **给人类：** 一眼看清流水线各步由哪个技能约束。  
> **Agent：** 契约以英文主版为准；默认只读 `CATALOG.md`，不要用本中文副版作为技能契约。

---

## 速览

| 步骤 | 技能 id | 文件 | 一句话 |
|:----:|---------|------|--------|
| ② | `pro-generation` | [`pro-generation.md`](pro-generation.md) | 只出 PRO，不写代码 |
| ④ | `template-matching` | [`template-matching.md`](template-matching.md) | 选 1～N 个 apps + 0～N patterns + images |
| ④ | `mvp-assembly` | [`mvp-assembly.md`](mvp-assembly.md) | 复制 app、合并 patterns、容器可跑 |
| ⑥ | `deploy` | [`deploy.md`](deploy.md) | 验收通过后上线 |

**配套 Prompt：** [`../prompts/`](../prompts/) · **模版检索：** [`../templates/CATALOG.md`](../templates/CATALOG.md) · **Patterns：** [`../templates/patterns/index.md`](../templates/patterns/index.md)

---

## 按步骤加载

| 当前步骤 | 必读 |
|----------|------|
| ② 出 PRO | `pro-generation.md` + `prompts/02-pro-draft.md`；结构 `prompts/pro.template.md`，样板 `prompts/pro.example.md` |
| ④ 组装 | `template-matching.md` → `templates/CATALOG.md` → apps + patterns → `mvp-assembly.md` |
| ⑥ 部署 | `deploy.md` + `release/` |

硬门禁：③ 未确认 PRO → 禁止 ④；⑤ 未确认 MVP → 禁止 ⑥。见 [`docs/workflow.md`](../docs/workflow.md)。

---

## 技能契约摘要

| 技能 | MUST | MUST NOT |
|------|------|----------|
| PRO 生成 | 含摘要/流程/模型/API/验收 | 输出实现代码、最终选定模版 |
| 模版检索 | 读 CATALOG + index；写出 image 依赖 | 自创脚手架、未确认 PRO 就选 |
| MVP 组装 | 输出到**产品仓根**；从镜像片段拼装 Dockerfile | 拷贝整个 `templates/images/` 树；本步部署；写入工厂仓 |
| 部署 | 跟 `release/` 脚本与端口池 | 跳过本地/验收门禁 |

---

## 登记规则

新增技能时 **同时** 更新：

1. 本文件（速览表）
2. [`README.md`](README.md)（Agent 规则表）
3. [`docs/workflow.md`](../docs/workflow.md) 对应步骤
4. [`AGENTS.md`](../AGENTS.md) 状态机表（若新增步骤触点）
