#!/usr/bin/env bash
# Minimal factory checks for CI and local smoke.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

fail=0

search() {
  # usage: search <rg-pattern> [--glob GLOB] [PATH...]
  local pattern="$1"
  shift
  local glob=""
  local -a paths=()
  while (($#)); do
    case "$1" in
      --glob)
        glob="$2"
        shift 2
        ;;
      *)
        paths+=("$1")
        shift
        ;;
    esac
  done
  if ((${#paths[@]} == 0)); then
    paths=(.)
  fi
  if command -v rg >/dev/null 2>&1; then
    if [[ -n "$glob" ]]; then
      rg -n --glob "$glob" "$pattern" "${paths[@]}"
    else
      rg -n "$pattern" "${paths[@]}"
    fi
  else
    local -a grep_args=(-RIn --exclude-dir=.git)
    if [[ -n "$glob" ]]; then
      case "$glob" in
        \*\*/Dockerfile\* | Dockerfile\*) grep_args+=(--include='Dockerfile*') ;;
        *) grep_args+=(--include="$glob") ;;
      esac
    fi
    grep "${grep_args[@]}" -E "$pattern" "${paths[@]}"
  fi
}

echo "==> bash -n scripts + release shell"
bash -n scripts/maker-flow
bash -n scripts/install.sh
bash -n scripts/check.sh
bash -n release/deploy/push-and-route.sh
bash -n release/cloudflare/dns.sh
bash -n release/cloudflare/dns-upsert.sh

echo "==> forbid private maker-flow image FROM in Dockerfiles"
if search '^[[:space:]]*FROM[[:space:]]+maker-flow/' --glob '**/Dockerfile*' templates release; then
  echo "error: found FROM maker-flow/ in Dockerfile (use upstream images + fragments)" >&2
  fail=1
fi

echo "==> forbid workspace/ assemble target in contracts"
if command -v rg >/dev/null 2>&1; then
  if rg -n 'workspace/' --glob '!docs/superpowers/**' \
    AGENTS.md AGENTS.zh-CN.md skills docs templates prompts README.md README.zh-CN.md; then
    echo "error: found workspace/ reference (assemble only in product repos)" >&2
    fail=1
  fi
else
  if grep -RIn --exclude-dir=.git --exclude-dir=superpowers 'workspace/' \
    AGENTS.md AGENTS.zh-CN.md skills docs templates prompts README.md README.zh-CN.md; then
    echo "error: found workspace/ reference (assemble only in product repos)" >&2
    fail=1
  fi
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

echo "==> layouts catalog present"
test -f templates/layouts/index.md
test -f templates/layouts/web-api/docker-compose.yml

echo "==> maker-flow help"
./scripts/maker-flow help >/dev/null

if [[ "$fail" -ne 0 ]]; then
  echo "FAIL"
  exit 1
fi
echo "OK"
