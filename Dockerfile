FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./.env ./.env
COPY ./migration ./migration
COPY ./data ./data

RUN go build -o warehouse ./cmd/warehouse_api/main.go
CMD ["/app/warehouse"]
