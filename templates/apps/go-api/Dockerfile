# Builder stage — composed from templates/images/go-builder
FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /src/server ./cmd/server

# Runtime stage — composed from templates/images/go-runtime
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata wget
WORKDIR /app
USER root
COPY --from=builder /src/server /app/server
RUN chown nobody:nobody /app/server
USER nobody
ENV APP_ENV=production HTTP_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]
