# ⚡ Build Stage
FROM --platform=linux/amd64 golang:1.22 AS builder

# Install required dependencies for CGO
RUN apt-get update && apt-get install -y \
    gcc \
    g++ \
    librdkafka-dev \
    && rm -rf /var/lib/apt/lists/*

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (leverage Docker caching)
COPY go.mod go.sum ./

# Ensure dependencies are installed
RUN go mod download

# Copy the entire project (including `cmd/`, `internal/`, `pkg/`)
COPY . .

# Change working directory to where main.go is located
WORKDIR /app/cmd

# Ensure proper compiler is used
ENV CC=gcc CXX=g++ CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# Build the Go binary
RUN go build -o /app/ingestion-go .

# 🎯 Minimal Runtime Stage
FROM --platform=linux/amd64 debian:bullseye-slim  

# Install required runtime libraries (use librdkafka1 instead of dev package)
RUN apt-get update && apt-get install -y --no-install-recommends librdkafka1 && rm -rf /var/lib/apt/lists/*

# Create a non-root user for security
RUN useradd -m appuser

# Set working directory inside container
WORKDIR /home/appuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ingestion-go .

# Change ownership and set permissions
RUN chmod +x ingestion-go && chown appuser:appuser ingestion-go

# Switch to non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./ingestion-go"]