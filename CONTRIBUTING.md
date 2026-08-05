# Contributing to Maker Flow

**English** · [简体中文](CONTRIBUTING.zh-CN.md)

Thanks for your interest in contributing! Maker Flow is a factory repository designed to orchestrate AI agents for building MVP products.

Because this repository focuses heavily on infrastructure, contracts, and templates rather than business logic, we ask that all contributions follow the guidelines below to maintain structural consistency.

## 1. Golden Rule: English is the Contract

To keep Agents strictly aligned:
- **Agents must only read the English (`.md`) files.**
- If you modify a skill, a prompt, or a template catalog, you **MUST** update the primary English file.
- `*.zh-CN.md` files are provided solely for human readability. They are not authoritative agent contracts. See [`docs/i18n.md`](docs/i18n.md) for details.

## 2. Adding a New Skill

If you introduce a new agent skill (e.g., `skills/publish-aws.md`):
1. **Register it:** Add it to `skills/CATALOG.md`.
2. **Follow format:** Include standard YAML frontmatter (`name`, `description`).
3. **Be deterministic:** Use `MUST` / `MUST NOT` keywords to set hard boundaries for the agent.
4. **Update docs:** Mention the new skill in `README.md` or `docs/workflow.md` if it changes the standard state machine.

## 3. Adding a New App Template

If you introduce a new app scaffold (e.g., `templates/apps/python-api`):
1. **Register it:** Add it to `templates/CATALOG.md` and `templates/index.md`.
2. **Tests are mandatory:** Include at least one test file (e.g., `test_main.py` or `handler_test.go`) to serve as a baseline for the agent in Step 4/5.
3. **Lock dependencies:** Always include the lockfile (`requirements.txt`, `poetry.lock`, `go.sum`, `package-lock.json`).
4. **Acceptance evidence:** Ensure `skills/mvp-acceptance.md` covers the verification commands for your new stack (e.g., `pytest`).

## 4. Submitting a Pull Request

1. **Local Checks:** Run `scripts/check.sh` locally to ensure no syntax errors and structural drift.
2. **Commit messages:** Use Conventional Commits (`feat: ...`, `fix: ...`, `docs: ...`).
3. **Tests:** All templates in `templates/patterns/` must pass `go test ./...`. Our GitHub Actions CI will enforce this.

## 5. Factory vs. Product

Never commit MVP/product business logic into this repository. This repo (`maker-flow`) is the *tool*. The products it builds live in separate repositories created by `maker-flow new <name>`. See [`docs/consumer-project.md`](docs/consumer-project.md).
