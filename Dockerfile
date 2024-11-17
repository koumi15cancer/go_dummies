# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy and install dependencies
COPY go.mod ./
# COPY go.sum ./
RUN go mod download

# Copy application code and build
COPY . .
RUN go build -o /app/api cmd/api/main.go  # Ensure binary is placed at /app/api

# Run Stage
FROM alpine:latest
WORKDIR /root/

# Copy the built binary from the builder image
COPY --from=builder /app/api .

# Ensure the binary has execute permissions
RUN chmod +x /root/api

# Expose port and run
EXPOSE 8080
CMD ["/root/api"]  # Correct the path to /root/api
