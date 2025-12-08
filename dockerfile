# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

# Run stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .
COPY .env .

EXPOSE 8088

CMD ["./app"]
