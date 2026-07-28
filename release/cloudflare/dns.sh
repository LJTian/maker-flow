#!/usr/bin/env bash
# Cloudflare DNS CRUD — see release/cloudflare/dns-api.md
#
# Env: CLOUDFLARE_API_TOKEN, CLOUDFLARE_ZONE_ID
set -euo pipefail

die() { echo "dns: $*" >&2; exit 1; }

require_env() {
  [[ -n "${CLOUDFLARE_API_TOKEN:-}" ]] || die "set CLOUDFLARE_API_TOKEN"
  [[ -n "${CLOUDFLARE_ZONE_ID:-}" ]] || die "set CLOUDFLARE_ZONE_ID"
}

api_base() {
  printf '%s' "https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_ID}/dns_records"
}

curl_api() {
  curl -sS -H "Authorization: Bearer ${CLOUDFLARE_API_TOKEN}" -H "Content-Type: application/json" "$@"
}

check_success() {
  local json="$1"
  local label="${2:-api}"
  printf '%s' "$json" | python3 -c '
import json,sys
data=json.load(sys.stdin)
if not data.get("success"):
  print(json.dumps(data, indent=2), file=sys.stderr)
  sys.exit(1)
' || die "${label} failed"
}

urlencode() {
  python3 -c 'import urllib.parse,sys; print(urllib.parse.quote(sys.argv[1]))' "$1" 2>/dev/null || printf '%s' "$1"
}

find_record_id() {
  local type="$1" name="$2"
  local enc q json
  enc="$(urlencode "$name")"
  q="type=${type}&name=${enc}"
  json="$(curl_api "$(api_base)?${q}")"
  check_success "$json" "list"
  printf '%s' "$json" | python3 -c '
import json,sys
data=json.load(sys.stdin)
res=data.get("result") or []
print(res[0]["id"] if res else "")
'
}

usage() {
  cat <<'EOF'
Usage: dns.sh <command> [options]

Commands:
  list    List DNS records in the zone
  get     Get one record by id
  create  Create a record
  update  Update a record (by --id or --type + --name)
  delete  Delete a record (by --id or --type + --name)
  upsert  Create or update (same type + name)

Env:
  CLOUDFLARE_API_TOKEN   API token (Zone DNS Edit)
  CLOUDFLARE_ZONE_ID     Target zone id

Examples:
  dns.sh list
  dns.sh list --type A --name api.example.com
  dns.sh get --id RECORD_ID
  dns.sh create --type A --name api.example.com --content 203.0.113.10
  dns.sh update --type A --name api.example.com --content 203.0.113.11
  dns.sh delete --type A --name api.example.com
  dns.sh upsert --type AAAA --name home.example.com --content 2001:db8::1

Legacy: dns-upsert.sh calls upsert (same flags as create).
EOF
}

cmd_list() {
  local type="" name="" json_flag=0 per_page=100
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --type) type="$2"; shift 2 ;;
      --name) name="$2"; shift 2 ;;
      --per-page) per_page="$2"; shift 2 ;;
      --json) json_flag=1; shift ;;
      *) die "list: unknown option $1" ;;
    esac
  done

  local page=1 query="per_page=${per_page}"
  [[ -n "$type" ]] && query="${query}&type=${type}"
  [[ -n "$name" ]] && query="${query}&name=$(urlencode "$name")"

  local all='[]'
  while true; do
    local json
    json="$(curl_api "$(api_base)?${query}&page=${page}")"
    check_success "$json" "list"
    local chunk done
    read -r chunk done < <(printf '%s' "$json" | python3 -c '
import json,sys
data=json.load(sys.stdin)
print(json.dumps(data.get("result") or []))
info=data.get("result_info") or {}
page=info.get("page") or 1
total=info.get("total_pages") or 1
print("done" if page >= total else "more")
')
    all="$(python3 -c 'import json,sys; a=json.loads(sys.argv[1]); b=json.loads(sys.argv[2]); print(json.dumps(a+b))' "$all" "$chunk")"
    [[ "$done" == "done" ]] && break
    page=$((page + 1))
  done

  if [[ "$json_flag" -eq 1 ]]; then
    printf '%s\n' "$all" | python3 -m json.tool
    return
  fi

  printf '%s' "$all" | python3 -c '
import json,sys
rows=json.load(sys.stdin)
if not rows:
  print("(no records)")
  sys.exit(0)
for r in rows:
  print("%s\t%s\t%s\t%s\tproxied=%s" % (
    r.get("id",""), r.get("type",""), r.get("name",""), r.get("content",""), r.get("proxied")))
print("---")
print("%d record(s)" % len(rows))
'
}

cmd_get() {
  local id="" json_flag=0
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --id) id="$2"; shift 2 ;;
      --json) json_flag=1; shift ;;
      *) die "get: unknown option $1" ;;
    esac
  done
  [[ -n "$id" ]] || die "get: need --id"

  local json
  json="$(curl_api "$(api_base)/${id}")"
  check_success "$json" "get"
  if [[ "$json_flag" -eq 1 ]]; then
    printf '%s\n' "$json" | python3 -m json.tool
  else
    printf '%s' "$json" | python3 -c '
import json,sys
r=json.load(sys.stdin)["result"]
print("id=%s" % r.get("id"))
print("type=%s" % r.get("type"))
print("name=%s" % r.get("name"))
print("content=%s" % r.get("content"))
print("proxied=%s" % r.get("proxied"))
print("ttl=%s" % r.get("ttl"))
'
  fi
}

build_body() {
  local type="$1" name="$2" content="$3" proxied="${4:-true}"
  python3 -c '
import json,sys
print(json.dumps({
  "type": sys.argv[1],
  "name": sys.argv[2],
  "content": sys.argv[3],
  "proxied": sys.argv[4] == "true",
  "ttl": 1,
}))
' "$type" "$name" "$content" "$proxied"
}

parse_record_flags() {
  TYPE="" NAME="" CONTENT="" PROXIED="true"
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --type) TYPE="$2"; shift 2 ;;
      --name) NAME="$2"; shift 2 ;;
      --content) CONTENT="$2"; shift 2 ;;
      --proxied) PROXIED="$2"; shift 2 ;;
      *) die "unknown option: $1" ;;
    esac
  done
  case "$PROXIED" in true|false) ;; *) die "--proxied must be true or false" ;; esac
}

print_record_ok() {
  local json="$1" verb="$2"
  check_success "$json" "$verb"
  printf '%s' "$json" | python3 -c '
import json,sys
r=json.load(sys.stdin)["result"]
print("ok %s id=%s %s %s -> %s proxied=%s" % (
  sys.argv[1], r.get("id"), r.get("type"), r.get("name"), r.get("content"), r.get("proxied")))
' "$verb"
}

cmd_create() {
  parse_record_flags "$@"
  [[ -n "$TYPE" && -n "$NAME" && -n "$CONTENT" ]] || die "create: need --type --name --content"
  local body json
  body="$(build_body "$TYPE" "$NAME" "$CONTENT" "$PROXIED")"
  echo "==> create ${TYPE} ${NAME} -> ${CONTENT} (proxied=${PROXIED})"
  json="$(curl_api -X POST "$(api_base)" --data "$body")"
  print_record_ok "$json" "create"
}

cmd_update() {
  local id="" 
  TYPE="" NAME="" CONTENT="" PROXIED=""
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --id) id="$2"; shift 2 ;;
      --type) TYPE="$2"; shift 2 ;;
      --name) NAME="$2"; shift 2 ;;
      --content) CONTENT="$2"; shift 2 ;;
      --proxied) PROXIED="$2"; shift 2 ;;
      *) die "update: unknown option $1" ;;
    esac
  done

  if [[ -z "$id" ]]; then
    [[ -n "$TYPE" && -n "$NAME" ]] || die "update: need --id or --type + --name"
    id="$(find_record_id "$TYPE" "$NAME")"
    [[ -n "$id" ]] || die "update: no record for type=${TYPE} name=${NAME}"
  fi

  if [[ -z "$TYPE" || -z "$NAME" || -z "$CONTENT" || -z "$PROXIED" ]]; then
    local cur
    cur="$(curl_api "$(api_base)/${id}")"
    check_success "$cur" "get"
    read -r TYPE NAME CONTENT PROXIED < <(printf '%s' "$cur" | python3 -c '
import json,sys
r=json.load(sys.stdin)["result"]
print(r.get("type",""), r.get("name",""), r.get("content",""), str(r.get("proxied")).lower())
')
  fi

  [[ -n "$CONTENT" ]] || die "update: need --content (or existing record must have content)"
  PROXIED="${PROXIED:-true}"

  local body json
  body="$(build_body "$TYPE" "$NAME" "$CONTENT" "$PROXIED")"
  echo "==> update ${TYPE} ${NAME} -> ${CONTENT} (proxied=${PROXIED}) id=${id}"
  json="$(curl_api -X PATCH "$(api_base)/${id}" --data "$body")"
  print_record_ok "$json" "update"
}

cmd_delete() {
  local id="" type="" name=""
  while [[ $# -gt 0 ]]; do
    case "$1" in
      --id) id="$2"; shift 2 ;;
      --type) type="$2"; shift 2 ;;
      --name) name="$2"; shift 2 ;;
      *) die "delete: unknown option $1" ;;
    esac
  done

  if [[ -z "$id" ]]; then
    [[ -n "$type" && -n "$name" ]] || die "delete: need --id or --type + --name"
    id="$(find_record_id "$type" "$name")"
    [[ -n "$id" ]] || die "delete: no record for type=${type} name=${name}"
  fi

  echo "==> delete id=${id}"
  local json
  json="$(curl_api -X DELETE "$(api_base)/${id}")"
  check_success "$json" "delete"
  printf '%s' "$json" | python3 -c '
import json,sys
r=json.load(sys.stdin)["result"]
print("ok deleted id=%s" % r.get("id"))
'
}

cmd_upsert() {
  parse_record_flags "$@"
  [[ -n "$TYPE" && -n "$NAME" && -n "$CONTENT" ]] || die "upsert: need --type --name --content"
  local id
  id="$(find_record_id "$TYPE" "$NAME")"
  if [[ -n "$id" ]]; then
    cmd_update --id "$id" --type "$TYPE" --name "$NAME" --content "$CONTENT" --proxied "$PROXIED"
  else
    cmd_create --type "$TYPE" --name "$NAME" --content "$CONTENT" --proxied "$PROXIED"
  fi
}

main() {
  [[ $# -gt 0 ]] || { usage; exit 0; }
  case "$1" in
    -h|--help|help) usage; exit 0 ;;
    list|get|create|update|delete|upsert) local cmd="$1"; shift; require_env; "cmd_${cmd}" "$@" ;;
    *) die "unknown command: $1 (try: list|get|create|update|delete|upsert)" ;;
  esac
}

main "$@"
