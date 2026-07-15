# 快速开始

按六步流程跑通一次。详见 [workflow.md](workflow.md)。

## 0. 配置 AI

```bash
cp ai-engine/.env.example ai-engine/.env
# 参考 ai-engine/providers/ 填写
```

## 1–2. 需求 → PRO

编辑 `prompts/01-requirement.example.md`，同步到 `prompts/02-pro-draft.md` 的「用户需求」段。

```bash
chmod +x scripts/ai-run.sh
./scripts/ai-run.sh prompts/02-pro-draft.md
```

对照 `skills/pro-generation.md` 检查 PRO 章节是否齐全。

## 3. 确认 PRO

与 AI 迭代修改，定稿后粘贴到 `prompts/03-pro-confirmed.example.md`，勾选「已确认」。

## 4. 组装 MVP

```bash
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

AI 应选定模版并将代码落到 `workspace/<项目名>/`。若使用 Cursor Agent，也可直接要求它读技能后写入文件。

## 5. 本地验收

```bash
cd workspace/<项目名>
cp .env.example .env
docker compose up --build
curl http://localhost:8080/health
```

按 PRO 验收标准逐项勾选。

## 6. 部署

```bash
export MVP_NAME=idea1 MVP_PORT=8080 DOMAIN=idea1.your-domain.com
export DEPLOY_HOST=deploy@your-server DEPLOY_PATH=/opt/mvps/idea1
./release/deploy/push-and-route.sh
```

详见 `skills/deploy.md` 与 `release/`。

## 检查清单

- [ ] 步骤 ③ PRO 已确认后才执行步骤 ④
- [ ] `templates/index.md` 能解释模版选型
- [ ] `workspace/` 下 MVP 本地 health 200
- [ ] 步骤 ⑤ 通过后才部署
- [ ] 至少完成一次 ①→⑥ 闭环

## 模版冒烟（跳过 AI）

```bash
cp -r templates/go-api workspace/smoke-test
cd workspace/smoke-test && cp .env.example .env
docker compose up --build
```
