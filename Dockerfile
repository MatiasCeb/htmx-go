# Build stage
FROM golang:1.22.3-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy static files
COPY --from=builder /app/views ./views
COPY --from=builder /app/css ./css
COPY --from=builder /app/images ./images

# Expose port
EXPOSE 42069

# Command to run
CMD ["./main"]