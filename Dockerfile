# Build stage
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files and download deps first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the app
RUN go build -o web-analyzer ./cmd/webanalyzer/main.go

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app .

# Copy static files
COPY static ./static

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./web-analyzer"]
