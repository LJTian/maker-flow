# Build Go base images for inheritance-style app Dockerfiles.
# Tags must match templates/images/index.md
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
IMAGES="${REPO_ROOT}/templates/images"

echo "==> building maker-flow/go-builder:1.22"
docker build -t maker-flow/go-builder:1.22 "${IMAGES}/go-builder"

echo "==> building maker-flow/go-runtime:1.22"
docker build -t maker-flow/go-runtime:1.22 "${IMAGES}/go-runtime"

echo "==> done"
docker images 'maker-flow/go-*'
