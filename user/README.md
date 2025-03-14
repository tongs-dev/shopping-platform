# User Service

## Overview

The User Service is a microservice in the Shopping Platform project responsible for user authentication, registration, and account management. It is built using Go, gRPC, and MySQL, following a microservices architecture.

## Project Structure
```
user/
├── domain/
│   ├── model/              # Data Models
│   ├── repository/         # Database Operations
│   ├── service/            # Business Logic
│
├── handler/                # gRPC Handlers
├── proto/                  # gRPC Protobuf Definitions
│   ├── user/
│   │   ├── user.proto      # gRPC API Specification
│   │   ├── user.pb.go      # Generated Proto Go Code
│   │   ├── user.pb.micro.go
│
├── Dockerfile              # Docker Build Configuration
├── docker-compose.yml      # Multi-Container Setup (MySQL & Service)
├── main.go                 # Service Entry Point
├── Makefile                # Build Automation
├── go.mod                  # Dependencies
├── go.sum                  # Package Checksum
├── README.md               # Documentation
```

## Features

- User Registration
- User Authentication (Login)
- Password Hashing using bcrypt
- MySQL Database Integration
- gRPC for Inter-Service Communication
- Docker and Docker Compose for containerized deployment

## Technologies Used

- Go (Golang)
- gRPC (Protocol Buffers)
- MySQL (Database)
- GORM (ORM for Go)
- Micro (Go Micro v2 framework for microservices)
- Docker & Docker Compose (Containerization & Deployment)
- Unit Testing (with mock repository & MySQL integration tests)


## Setup & Installation

1. Clone the Repository
```shell
git clone https://github.com/your-org/shopping-platform.git
cd shopping-platform/user
```

2. Install Dependencies
```shell
go mod tidy
```

3. Start MySQL & User Service using Docker
```shell
make docker-start
```
This will start MySQL and the User Service.

4. Stop docker containers
```shell
make docker-stop
```

## Running the Service

1. Locally (without Docker)
```shell
docker-compose up -d mysql
go run main.go
```

2. Inside Docker
```shell
make docker-build
make docker-start
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
// userRepo := repository.NewUserRepository(db)
// err = userRepo.InitTable()
```
Run the service, then comment it back once tables are created.
