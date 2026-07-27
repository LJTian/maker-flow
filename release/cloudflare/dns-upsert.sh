#!/usr/bin/env bash
# Upsert a Cloudflare DNS record (A / AAAA / CNAME / …).
# Agent-facing helper — see release/cloudflare/dns-api.md
#
# Required env:
#   CLOUDFLARE_API_TOKEN
#   CLOUDFLARE_ZONE_ID
# Record (flags or env):
#   --type / CF_DNS_TYPE
#   --name / CF_DNS_NAME      FQDN or relative name in the zone
#   --content / CF_DNS_CONTENT
#   --proxied / CF_DNS_PROXIED   true|false (default true)
set -euo pipefail

die() { echo "dns-upsert: $*" >&2; exit 1; }

TYPE="${CF_DNS_TYPE:-}"
NAME="${CF_DNS_NAME:-}"
CONTENT="${CF_DNS_CONTENT:-}"
PROXIED="${CF_DNS_PROXIED:-true}"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --type) TYPE="$2"; shift 2 ;;
    --name) NAME="$2"; shift 2 ;;
    --content) CONTENT="$2"; shift 2 ;;
    --proxied) PROXIED="$2"; shift 2 ;;
    -h|--help)
      sed -n '2,14p' "$0" | tr -d '#'
      exit 0
      ;;
    *) die "unknown option: $1" ;;
  esac
done

TOKEN="${CLOUDFLARE_API_TOKEN:-}"
ZONE="${CLOUDFLARE_ZONE_ID:-}"
[[ -n "$TOKEN" ]] || die "set CLOUDFLARE_API_TOKEN"
[[ -n "$ZONE" ]] || die "set CLOUDFLARE_ZONE_ID"
[[ -n "$TYPE" && -n "$NAME" && -n "$CONTENT" ]] || die "need --type --name --content"

case "$PROXIED" in
  true|false) ;;
  *) die "--proxied must be true or false" ;;
esac

AUTH=( -H "Authorization: Bearer ${TOKEN}" -H "Content-Type: application/json" )
API="https://api.cloudflare.com/client/v4/zones/${ZONE}/dns_records"

# URL-encode name lightly for query (space-safe; names should be hostnames)
ENC_NAME="$(python3 -c 'import urllib.parse,sys; print(urllib.parse.quote(sys.argv[1]))' "$NAME" 2>/dev/null \
  || printf '%s' "$NAME")"

LIST_JSON="$(curl -sS "${AUTH[@]}" "${API}?type=${TYPE}&name=${ENC_NAME}")"
SUCCESS="$(printf '%s' "$LIST_JSON" | python3 -c 'import json,sys; print(json.load(sys.stdin).get("success"))' 2>/dev/null || echo "")"
[[ "$SUCCESS" == "True" || "$SUCCESS" == "true" ]] || {
  echo "$LIST_JSON" >&2
  die "list dns_records failed"
}

RECORD_ID="$(printf '%s' "$LIST_JSON" | python3 -c '
import json,sys
data=json.load(sys.stdin)
res=data.get("result") or []
print(res[0]["id"] if res else "")
')"

BODY="$(python3 -c '
import json,sys
print(json.dumps({
  "type": sys.argv[1],
  "name": sys.argv[2],
  "content": sys.argv[3],
  "proxied": sys.argv[4] == "true",
  "ttl": 1,
}))
' "$TYPE" "$NAME" "$CONTENT" "$PROXIED")"

if [[ -n "$RECORD_ID" ]]; then
  echo "==> update ${TYPE} ${NAME} -> ${CONTENT} (proxied=${PROXIED}) id=${RECORD_ID}"
  RESP="$(curl -sS -X PATCH "${AUTH[@]}" "${API}/${RECORD_ID}" --data "$BODY")"
else
  echo "==> create ${TYPE} ${NAME} -> ${CONTENT} (proxied=${PROXIED})"
  RESP="$(curl -sS -X POST "${AUTH[@]}" "${API}" --data "$BODY")"
fi

printf '%s' "$RESP" | python3 -c '
import json,sys
data=json.load(sys.stdin)
if not data.get("success"):
  print(json.dumps(data, indent=2), file=sys.stderr)
  sys.exit(1)
r=data["result"]
print("ok id=%s %s %s -> %s proxied=%s" % (
  r.get("id"), r.get("type"), r.get("name"), r.get("content"), r.get("proxied")))
'
