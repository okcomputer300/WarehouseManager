FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./server

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m -u 10001 appuser

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/web/templates ./templates
COPY --from=builder /app/web/static ./static
COPY --from=builder /app/migrations ./migrations

RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 8080

CMD ["./server"]
