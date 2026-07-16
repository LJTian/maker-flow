# Build stage
FROM maker-flow/go-builder:1.22 AS builder
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /src/worker ./cmd/worker

FROM maker-flow/go-runtime:1.22
USER root
COPY --from=builder /src/worker /app/worker
RUN chown nobody:nobody /app/worker
USER nobody
ENV APP_ENV=production HTTP_ADDR=:8080 WORKERS=4 QUEUE_SIZE=64
EXPOSE 8080
ENTRYPOINT ["/app/worker"]
