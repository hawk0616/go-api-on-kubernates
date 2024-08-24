FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o practice02 ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/practice02 .

EXPOSE 8080

CMD ["./practice02"]