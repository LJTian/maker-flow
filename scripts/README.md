# scripts/

| Script | Agent use |
|--------|-----------|
| `ai-run.sh` | POST prompt file to OpenAI-compatible API using `ai-engine/.env` |
| `build-images.sh` | Build `maker-flow/go-*` base images before app `docker compose up --build` |

```bash
./scripts/build-images.sh
./scripts/ai-run.sh prompts/02-pro-draft.md
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

`ai-run.sh` not required when the host agent is the LLM. Deploy: `../release/deploy/`.