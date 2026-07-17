# Fragment: Go compile stage base.
# When assembling an app Dockerfile, inline these lines into the builder stage
# (do NOT pre-build a local maker-flow/* tag).
FROM golang:1.22-alpine
RUN apk add --no-cache git ca-certificates
WORKDIR /src
# --- end base fragment; app adds COPY / go build below ---
