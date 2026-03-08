# ======================
# Stage 1: Build
# ======================
FROM golang:1.23-alpine AS builder

# Install gcc for cgo dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for better layer caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o pos-app ./cmd/main.go

# ======================
# Stage 2: Run
# ======================
FROM alpine:latest

# Install ca-certificates for HTTPS calls if needed
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy only the built binary from builder stage
COPY --from=builder /app/pos-app .

# Copy .env file
COPY .env .

# Expose app port
EXPOSE 3000

# Run the binary
CMD ["./pos-app"]