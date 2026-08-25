FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gochat .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY --from=builder /build/gochat .
COPY --from=builder /build/static ./static
COPY --from=builder /build/config ./config
COPY --from=builder /build/import.sql .

RUN mkdir -p /app/static/upload /app/logs

EXPOSE 8081

CMD ["./gochat", "server", "-p", "8081"]
