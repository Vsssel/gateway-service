# syntax=docker/dockerfile:1.4
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/server ./cmd

# Final image
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /bin/server .
COPY swagger-ui /swagger-ui

EXPOSE 7088 7089 7090

CMD ["./server"]
