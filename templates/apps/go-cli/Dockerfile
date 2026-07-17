# Builder stage — composed from templates/images/go-builder
FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /src/cli ./cmd/cli

# Runtime stage — composed from templates/images/go-runtime
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata wget
WORKDIR /app
USER root
COPY --from=builder /src/cli /app/cli
RUN chown nobody:nobody /app/cli
USER nobody
ENTRYPOINT ["/app/cli"]
CMD ["version"]
