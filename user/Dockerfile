# Use official Golang image
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service .

# Use a minimal base image for production
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/user-service .

# Expose the service port (change if needed)
EXPOSE 8080

# Start the application
CMD ["./user-service"]
