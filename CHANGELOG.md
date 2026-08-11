# Changelog

All notable changes to Maker Flow will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Roadmap unimplemented inventory**: Documented missing patterns (WeChat Pay, Alipay, Stripe, …) and factory gaps under `docs/roadmap.md`.
- **Roadmap**: Added `docs/roadmap.md` (+ zh-CN) and linked it from the README homepage and docs index.
- **Multi-app PRO example**: Added `prompts/pro-multi.example.md` (Web SPA + API).

### Changed
- **Catalog honesty**: Narrowed `payment-lemonsqueezy` tags/copy to Lemon Squeezy MoR only; clarified auth vs pay, OpenAI-compatible LLM, and Vercel static-only in matching/deploy skills.
- **Git fallback in install script**: Added `curl/wget` + `unzip` fallback when `git` is not installed during remote setup.
- **CI Workflows**: Added `.github/workflows/ci.yml` for automated testing and linting.
- **Failure Paths Contract**: Added rollback and error handling instructions to `docs/workflow.md`.
- **CONTRIBUTING.md**: Added guidelines for how to contribute templates, skills, and code to this factory repository.
- **CHANGELOG.md**: Added version history tracking.

### Fixed
- **Pattern Concurrency Flaws**: Fixed deep logical bugs in `singleflight-cache` (memory leak), `pipeline` (goroutine leak), `circuit-breaker` (stampede bypass), and `worker-pool` (panic bomb on close).
- **Test Integrity**: Fixed `persistence-sqlx` broken tests (package mismatch, function name mismatch).
- **Deployment Robustness**: Hardened `release/deploy/push-and-route.sh` by replacing `/` with `|` in `sed` delimiters.
- **Docker Best Practices**: Modernized all Go app Dockerfiles (`go-api`, `go-cli`, `go-worker`) to use `COPY --chown=nobody:nobody` instead of redundant `USER root` toggles.
- **Config Fixes**: Added missing `package-lock.json` to `web-vite` and `go.sum` to `go-cli` and `go-worker` to ensure reproducible builds.
- **Template Typos**: Fixed missing colon in `docker-compose.yml` restart policy for `go-api`.
