# 步骤 ⑥ — 发布（对话）

[English](06-publish.md) · **简体中文**

步骤 ⑤ 通过后加载。技能：`skills/deploy.md`。

## 对人类说（可改措辞）

MVP 已验收。上线前请确认：

1. **发什么？**（整站 / 仅前端 / 仅 API）
2. **发到哪？**
   - Cloudflare Pages / GitHub Pages / Vercel — 适合静态或 SPA
   - 自有 VPS（Docker 网关）— 适合 API、worker 或整包 compose
   - 可以拆分（例如前端 Pages + API VPS）
3. **域名？**（平台默认 URL，还是你已有的域名）
4. 本机平台登录 / Token 是否就绪？

你选定后由我执行发布步骤——**不需要**你运行 deploy 命令。

## 得到答复后

按 `release/publish/<target>.md` 执行，并回报公网 URL 与验证结果。
