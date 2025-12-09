# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api/main.go

CMD ["./main"]
