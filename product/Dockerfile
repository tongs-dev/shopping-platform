# Use Go as the base image
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy the entire monorepo to the container
COPY . /app/product

# Set Go module path for the user service
WORKDIR /app/product

# Ensure modules are linked properly
RUN go mod tidy

# Build the user service binary
RUN go build -o product-service .

# Use a lightweight image for runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/product/product-service .

# Expose port and run the application
EXPOSE 8082
CMD ["./product-service"]
