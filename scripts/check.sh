#!/usr/bin/env bash
# Minimal factory checks for CI and local smoke.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

fail=0

echo "==> bash -n scripts"
bash -n scripts/maker-flow
bash -n scripts/install.sh
bash -n release/deploy/push-and-route.sh

echo "==> forbid private maker-flow image FROM in Dockerfiles"
if rg -n '^[[:space:]]*FROM[[:space:]]+maker-flow/' --glob '**/Dockerfile*' templates release 2>/dev/null; then
  echo "error: found FROM maker-flow/ in Dockerfile (use upstream images + fragments)" >&2
  fail=1
fi
echo "==> forbid workspace/ assemble target in contracts"
if rg -n 'workspace/' --glob '!docs/superpowers/**' AGENTS.md AGENTS.zh-CN.md skills docs templates prompts README.md README.zh-CN.md 2>/dev/null; then
  echo "error: found workspace/ reference (assemble only in product repos)" >&2
  fail=1
fi

echo "==> image fragment upstream lines present in Go app Dockerfiles"
for app in go-api go-cli go-worker; do
  df="templates/apps/${app}/Dockerfile"
  grep -q 'FROM golang:1.22-alpine' "$df" || { echo "error: $df missing go-builder upstream" >&2; fail=1; }
  grep -q 'FROM alpine:3.20' "$df" || { echo "error: $df missing go-runtime upstream" >&2; fail=1; }
  grep -q 'apk add --no-cache git ca-certificates' "$df" || { echo "error: $df missing builder apk line" >&2; fail=1; }
  grep -q 'apk add --no-cache ca-certificates tzdata wget' "$df" || { echo "error: $df missing runtime apk line" >&2; fail=1; }
done

echo "==> gateway default conf exists"
test -f release/nginx/conf.d/00-default.conf
test -f release/nginx/docker-compose.yml

echo "==> maker-flow help"
./scripts/maker-flow help >/dev/null

if [[ "$fail" -ne 0 ]]; then
  echo "FAIL"
  exit 1
fi
echo "OK"
