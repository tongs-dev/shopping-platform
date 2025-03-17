# Product Service

## Overview
The Product Service is a microservice that handles product management functionality for the e-commerce platform. It provides capabilities such as adding, updating, deleting, and fetching product details. This service is built using Go and gRPC for efficient communication between services.

It uses MySQL for data persistence and is integrated with Consul for service discovery and Jaeger for distributed tracing.

## Project Structure
```
product/
├── domain/
│   ├── model/                  # Data Models
│   ├── repository/             # Database Operations
│   ├── service/                # Business Logic
│
├── handler/                    # gRPC Handlers
├── proto/                      # gRPC Protobuf Definitions
│   ├── product/
│   │   ├── product.proto       # gRPC API Specification
│   │   ├── product.pb.go       # Generated Proto Go Code
│   │   ├── product.pb.micro.go
│
├── Dockerfile                  # Docker Build Configuration
├── docker-compose.yml          # Multi-Container Setup (MySQL & Service)
├── main.go                     # Service Entry Point
├── product_client.go           # Client for interacting with the product service
├── Makefile                    # Build Automation
├── go.mod                      # Dependencies
├── go.sum                      # Package Checksum
├── README.md                   # Documentation
```

## Features

- Add Product: Add a new product to the catalog with details such as name, SKU, price, and description.
- Update Product: Update details of an existing product.
- Delete Product: Delete a product from the catalog by its ID.
- Find Product: Retrieve product details by ID, name, or other criteria.
- Product Observability: Integrated with Jaeger for distributed tracing and monitoring of product service interactions.

## Technologies Used

- Go (Golang): Backend service implementation.
- gRPC: Communication protocol for inter-service communication.
- MySQL: Relational database for data storage.
- GORM: ORM library for interacting with the MySQL database.
- Consul: Service discovery and configuration management.
- Jaeger: Distributed tracing for monitoring and debugging.


## Setup & Installation

1. Clone the Repository
```shell
git clone https://github.com/your-org/shopping-platform.git
cd shopping-platform/product
```

2. Install Dependencies
```shell
go mod tidy
```

3. Start MySQL, Consul, and Jaeger (optional)
```shell
make docker-start
```

4. Setup Mysql Config in Consul
In Consul, create new Key/Value pair in `/micro/config` folder
```json
{
  "host": "127.0.0.1",
  "user":"root",
  "pwd": "123456",
  "database":"productdb",
  "port": 3306
}
```

5. Stop docker containers
```shell
make docker-stop
```

## Running the Service

Locally (without Docker)
```shell
make docker-start
go run main.go
```

## Running Tests

1. Unit Tests
```shell
make test
```

2. Integration Tests (with MySQL in Docker)
```shell
make docker-start
make test
```

## Development Guidelines

**Generating gRPC Code** <br>
If you update the user.proto file, regenerate the gRPC files:

```shell
# Install go micro and required plugins 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/micro/micro/v2/cmd/protoc-gen-micro@latest

# Generate go code from protobuf
make proto
```


## Database Migrations

To initialize the database schema, uncomment the following in main.go and run:
```go
// productRepo := repository.NewUserRepository(db)
// err = productRepo.InitTable()
```
Run the service, then comment it back once tables are created.
