# Start from the official Golang image for building the app
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for Go module fetching
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code and .env file
COPY . .

# Build the Go app
RUN go build -o server .

# Use a minimal image for running the app
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the built binary from builder
COPY --from=builder /app/server .
# ✅ Copy the .env file as well
COPY --from=builder /app/.env .env

# Expose the port Gin runs on (default 9000)
EXPOSE 9000

# Run the server
CMD ["./server"]
