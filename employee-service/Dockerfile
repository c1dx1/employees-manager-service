FROM golang:1.23 AS builder

WORKDIR /employee-service

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /employee-service/employee-service /employee-service

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y libc6

WORKDIR /

COPY --from=builder /employee-service/employee-service /employee-service
COPY --from=builder /employee-service/migrations /migrations
COPY --from=builder /employee-service/config /config

RUN chmod +x /employee-service

EXPOSE 50051

CMD ["/employee-service"]
