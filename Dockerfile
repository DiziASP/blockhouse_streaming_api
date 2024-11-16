## BUILD STAGE ##
FROM golang:1.22.8-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o blockhouse_streaming_api ./cmd/main.go

## DEV STAGE ##
FROM alpine:3.18

ENV APP_PORT=8081
ENV KAFKA_BROKERS="redpanda-0:19092"

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/blockhouse_streaming_api .
COPY ./config ./config


RUN chown -R appuser:appgroup /app && chmod +x blockhouse_streaming_api

USER appuser

EXPOSE 8081

CMD ["./blockhouse_streaming_api"]
