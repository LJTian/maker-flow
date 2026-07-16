# Go runtime base — maker-flow/go-runtime:1.22

Infrastructure only. App Dockerfiles MUST `FROM maker-flow/go-runtime:1.22` and only COPY the binary + set ENTRYPOINT/ENV/EXPOSE.

Runs as `nobody`. Includes `wget` for compose healthchecks.

```bash
docker build -t maker-flow/go-runtime:1.22 templates/images/go-runtime
# or: ./scripts/build-images.sh
```
