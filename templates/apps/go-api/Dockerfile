# Build stage — inherits maker-flow/go-builder (run ./scripts/build-images.sh first)
# Dependencies resolve inside the container (no host go mod tidy required).
FROM maker-flow/go-builder:1.22 AS builder
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /src/server ./cmd/server

# Runtime stage — inherits maker-flow/go-runtime
FROM maker-flow/go-runtime:1.22
USER root
COPY --from=builder /src/server /app/server
RUN chown nobody:nobody /app/server
USER nobody
ENV APP_ENV=production HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]
