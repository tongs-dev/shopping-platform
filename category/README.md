# Category Service

## Overview

The Category Service is part of the shopping platform and is responsible for handling category-related operations like creating, updating, retrieving, and deleting categories. It interacts with a MySQL database and uses Consul for service discovery and configuration management.

## Project Structure
```
category/
│
├── common/                     # Shared utilities and configurations
│   ├── config.go               # Configuration management
│   ├── mysql.go                # MySQL connection utility
│   ├── swap.go                 # Data mapping utility
│
├── domain/
│   ├── model/                  # Data Models
│   ├── repository/             # Database Operations
│   ├── service/                # Business Logic
│
├── handler/                    # gRPC Handlers
├── proto/                      # GRPC proto files
│   ├── category/
│   │   ├── category.proto      # gRPC API Specification
│   │   ├── category.pb.go      # Generated Proto Go Code
│   │   ├── category.pb.micro.go
│
├── Dockerfile                  # Docker Build Configuration
├── docker-compose.yml          # Multi-Container Setup (MySQL & Service)
├── main.go                     # Service Entry Point
├── Makefile                    # Build Automation
├── go.mod                      # Dependencies
├── go.sum                      # Package Checksum
├── README.md                   # Documentation
```

## Features

- Category CRUD operations
- MySQL Database Integration
- gRPC for Inter-Service Communication
- Docker and Docker Compose for containerized deployment
- Consul for service discovery and configuration management

## Technologies Used

- Go (Golang)
- gRPC (Protocol Buffers)
- MySQL (Database)
- GORM (ORM for Go)
- Micro (Go Micro v2 framework for microservices)
- Docker & Docker Compose (Containerization & Deployment)
- Unit Testing (with mock repository & MySQL integration tests)
- Consul

## Setup & Installation

1. Clone the Repository
```shell
git clone https://github.com/your-org/shopping-platform.git
cd shopping-platform/category
```

2. Install Dependencies
```shell
go mod tidy
```

3. Start MySQL & Consul using Docker
```shell
make docker-start
```
This will start MySQL and the Category Service.

4. Stop docker containers
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
If you update the category.proto file, regenerate the gRPC files:

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
// categoryRepo := repository.NewCategoryRepository(db)
// err = categoryRepo.InitTable()
```
Run the service, then comment it back once tables are created.
