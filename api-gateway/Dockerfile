FROM golang:1.23 AS builder

WORKDIR /api-gateway

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /api-gateway/api-gateway /api-gateway

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y libc6

WORKDIR /

COPY --from=builder /api-gateway/api-gateway /api-gateway
COPY --from=builder /api-gateway/config /config

RUN chmod +x /api-gateway

EXPOSE 8080

CMD ["/api-gateway"]
