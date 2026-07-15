# scripts/

| Script | Agent use |
|--------|-----------|
| `ai-run.sh` | POST prompt file to OpenAI-compatible API using `ai-engine/.env` |

```bash
./scripts/ai-run.sh prompts/02-pro-draft.md
./scripts/ai-run.sh prompts/04-assemble-mvp.md
```

Not required when the host agent is the LLM. Deploy: `../release/deploy/`.
