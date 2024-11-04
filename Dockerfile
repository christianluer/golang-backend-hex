# Stage 1: Build the Go application
FROM golang:1.19-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Stage 2: Create a smaller image for running the application
FROM alpine:latest

# Set up a non-root user for security
RUN adduser -D -g '' appuser

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy .env file (optional, depending on security needs)
COPY .env .

# Set environment variables (optional, for readability)
ENV PORT=8080

# Expose the application port
EXPOSE 8080

# Run the application as a non-root user
USER appuser
CMD ["./main"]
