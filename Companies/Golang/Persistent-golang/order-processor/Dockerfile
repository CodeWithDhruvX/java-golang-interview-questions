# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
# COPY go.sum ./ 

RUN go mod download

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o order-processor main.go

# Production Stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/order-processor .

EXPOSE 8080

CMD ["./order-processor"]
