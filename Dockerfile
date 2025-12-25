# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download all dependencies (if any)
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# CGO_ENABLED=0 ensures a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o aiproxy .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates for HTTPS requests to Google API
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/aiproxy .
COPY --from=builder /app/index.html .
COPY --from=builder /app/robots.txt .

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["./aiproxy"]
