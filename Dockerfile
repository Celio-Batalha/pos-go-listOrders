# Build stage
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem

# Production stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copiar binário do builder
COPY --from=builder /app/main .

# Copiar .env se existir, ou criar vazio
COPY --from=builder /app/cmd/ordersystem/.env .

EXPOSE 8000 8080 50051
CMD ["./main"]
