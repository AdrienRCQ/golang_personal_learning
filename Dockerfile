# Stage 1: Build the Go binary
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

COPY . .

# Build the Go binary
RUN go build -o hello .

# Stage 2: Create a minimal runtime environment
FROM debian:bullseye-slim

# Set up a non-root user
RUN useradd -m appuser

# Copy the binary from the builder stage
COPY --from=builder /app/hello /usr/local/bin/hello

# Switch to the non-root user
USER appuser

# Command to run the binary
ENTRYPOINT ["hello"]
