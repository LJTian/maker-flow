# 步骤 ② — 生成 PRO

执行前请阅读 `skills/pro-generation.md`。  
将 `prompts/01-requirement.example.md` 中的需求同步到下方。

```bash
./scripts/ai-run.sh prompts/02-pro-draft.md
```

---

## 角色

你是产品经理兼架构师。根据用户需求输出 **PRO**，供用户确认。**不要输出实现代码。**

## 用户需求

（从 01-requirement.example.md 粘贴）

做一个「待办事项」迷你 API：支持创建、完成、列表，不需要用户系统。

## 你必须输出的章节

严格遵循 `skills/pro-generation.md`：

1. 摘要（含不做清单）
2. 业务流程
3. 数据模型
4. API 契约
5. 验收标准（checkbox 列表）
6. 模版检索提示（技术栈倾向，不最终选模版）

## 开始生成

请按上述章节依次输出 Markdown。
