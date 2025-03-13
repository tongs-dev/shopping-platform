# Use Go as the base image
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy the entire monorepo to the container
COPY . /app/user

# Set Go module path for the user service
WORKDIR /app/user

# Ensure modules are linked properly
RUN go mod tidy

# Build the user service binary
RUN go build -o user-service .

# Use a lightweight image for runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user/user-service .

# Expose port and run the application
EXPOSE 8080
CMD ["./user-service"]
