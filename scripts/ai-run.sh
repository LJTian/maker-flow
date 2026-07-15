#!/usr/bin/env bash
# 读取 ai-engine/.env，将 Prompt 文件发送至 OpenAI 兼容 API 并流式输出
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
CONFIG_FILE="${AI_ENGINE_CONFIG:-${REPO_ROOT}/ai-engine/.env}"

if [[ -f "$CONFIG_FILE" ]]; then
  set -a
  # shellcheck disable=SC1090
  source "$CONFIG_FILE"
  set +a
fi

PROMPT_FILE="${1:?用法: ai-run.sh <prompt.md> [model]}"
MODEL="${2:-${AI_MODEL:-}}"
BASE_URL="${AI_BASE_URL:-}"
API_KEY="${AI_API_KEY:-}"
TEMPERATURE="${AI_TEMPERATURE:-0.6}"
MAX_TOKENS="${AI_MAX_TOKENS:-8192}"
TOP_P="${AI_TOP_P:-0.9}"
SYSTEM_PROMPT_FILE="${AI_SYSTEM_PROMPT_FILE:-}"

[[ -n "$MODEL" ]] || { echo "错误: 请设置 AI_MODEL 或传入第二个参数" >&2; exit 1; }
[[ -n "$BASE_URL" ]] || { echo "错误: 请设置 AI_BASE_URL（见 ai-engine/.env.example）" >&2; exit 1; }
[[ -f "$PROMPT_FILE" ]] || { echo "错误: 找不到 Prompt 文件: $PROMPT_FILE" >&2; exit 1; }

USER_PROMPT=$(cat "$PROMPT_FILE")
ENDPOINT="${BASE_URL%/}/chat/completions"

SYSTEM_PROMPT=""
if [[ -n "$SYSTEM_PROMPT_FILE" ]]; then
  for candidate in "$SYSTEM_PROMPT_FILE" "${REPO_ROOT}/${SYSTEM_PROMPT_FILE}"; do
    if [[ -f "$candidate" ]]; then
      SYSTEM_PROMPT=$(cat "$candidate")
      break
    fi
  done
fi

if [[ -n "$SYSTEM_PROMPT" ]]; then
  PAYLOAD=$(jq -n \
    --arg model "$MODEL" \
    --arg system "$SYSTEM_PROMPT" \
    --arg user "$USER_PROMPT" \
    --argjson temperature "$TEMPERATURE" \
    --argjson max_tokens "$MAX_TOKENS" \
    --argjson top_p "$TOP_P" \
    '{model: $model, messages: [{role:"system",content:$system},{role:"user",content:$user}], temperature: $temperature, max_tokens: $max_tokens, top_p: $top_p, stream: true}')
else
  PAYLOAD=$(jq -n \
    --arg model "$MODEL" \
    --arg user "$USER_PROMPT" \
    --argjson temperature "$TEMPERATURE" \
    --argjson max_tokens "$MAX_TOKENS" \
    --argjson top_p "$TOP_P" \
    '{model: $model, messages: [{role:"user",content:$user}], temperature: $temperature, max_tokens: $max_tokens, top_p: $top_p, stream: true}')
fi

CURL_ARGS=(-sN "$ENDPOINT" -H "Content-Type: application/json")
if [[ -n "$API_KEY" ]]; then
  CURL_ARGS+=(-H "Authorization: Bearer $API_KEY")
fi

curl "${CURL_ARGS[@]}" -d "$PAYLOAD" | while IFS= read -r line; do
  [[ -z "$line" ]] && continue
  if [[ "$line" == data:* ]]; then
    payload="${line#data: }"
    [[ "$payload" == "[DONE]" ]] && break
    echo "$payload" | jq -r '.choices[0].delta.content // empty' 2>/dev/null || true
  else
    echo "$line" | jq -r '.choices[0].message.content // empty' 2>/dev/null || true
  fi
done

echo ""
