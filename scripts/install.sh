#!/usr/bin/env bash
# Install Maker Flow factory to ~/.maker-flow and link `maker-flow` into ~/.local/bin
#
# Local:  ./scripts/install.sh
# Remote: curl -fsSL https://raw.githubusercontent.com/LJTian/maker-flow/main/scripts/install.sh | bash
set -euo pipefail

REPO_URL="${MAKER_FLOW_REPO_URL:-https://github.com/LJTian/maker-flow.git}"
SOURCE=""

if [[ -n "${BASH_SOURCE[0]:-}" && "${BASH_SOURCE[0]}" != bash && "${BASH_SOURCE[0]}" != sh ]]; then
  _dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  _root="$(cd "${_dir}/.." && pwd)"
  if [[ -d "${_root}/templates" && -d "${_root}/skills" ]]; then
    SOURCE="${_root}"
  fi
fi

if [[ -z "${SOURCE}" ]]; then
  if ! command -v git >/dev/null 2>&1; then
    echo "install.sh: git is required for remote install" >&2
    exit 1
  fi
  TMP="$(mktemp -d)"
  trap 'rm -rf "${TMP}"' EXIT
  echo "==> Cloning ${REPO_URL} ..."
  git clone --depth 1 "${REPO_URL}" "${TMP}/maker-flow"
  SOURCE="${TMP}/maker-flow"
fi

export MAKER_FLOW_INSTALL_DIR="${MAKER_FLOW_INSTALL_DIR:-${HOME}/.maker-flow}"
exec "${SOURCE}/scripts/maker-flow" install --from "${SOURCE}" "$@"
