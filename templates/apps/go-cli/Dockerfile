# Build static CLI binary — inherits maker-flow/go-builder
FROM maker-flow/go-builder:1.22 AS builder
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /src/cli ./cmd/cli

FROM maker-flow/go-runtime:1.22
USER root
COPY --from=builder /src/cli /app/cli
RUN chown nobody:nobody /app/cli
USER nobody
ENTRYPOINT ["/app/cli"]
CMD ["version"]
