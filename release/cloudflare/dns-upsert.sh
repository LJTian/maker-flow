#!/usr/bin/env bash
# Back-compat wrapper — prefer: dns.sh upsert
set -euo pipefail
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
exec "${SCRIPT_DIR}/dns.sh" upsert "$@"
