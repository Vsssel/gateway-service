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
COPY --from=builder /app/swagger-ui ./swagger-ui
COPY --from=builder /app/protogen/golang/swagger/swagger.swagger.json ./swagger.json

EXPOSE 2222

CMD ["./server"]
