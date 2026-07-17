# Fragment: Go runtime stage base.
# When assembling an app Dockerfile, inline these lines into the runtime stage
# (do NOT pre-build a local maker-flow/* tag).
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata wget
WORKDIR /app
USER nobody
# --- end base fragment; app adds COPY binary / ENTRYPOINT / ENV / EXPOSE below ---
