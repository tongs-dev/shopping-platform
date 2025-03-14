# Shopping Platform

## Overview
The **Shopping Platform** is a microservices-based e-commerce system designed to handle different aspects of online shopping, including user management, orders, inventory, and payments. The system is built using **Go**, **gRPC**, and **MySQL**, with each microservice handling a specific domain of the platform.

## Features
- **User Service**: Handles user authentication, registration, and profile management.
- **Category Service**: Manages the categories for products, including adding, updating, deleting, and retrieving categories.
- **gRPC Communication**: Services interact via **gRPC** for efficient communication.
- **Consul Integration**: Uses Consul for service discovery and configuration management, ensuring that microservices can dynamically register and find each other.
- **Dockerized Deployment**: Uses **Docker & Docker Compose** for easy setup and scaling.
- **Monitoring & Logging**: Includes Prometheus, Grafana, and centralized logging for observability.

## Tech Stack
- **Go** (Golang) for backend services
- **gRPC** for inter-service communication
- **MySQL** for data persistence
- **GORM** as the ORM layer
- **Docker & Docker Compose** for containerization
- **Kubernetes (optional)** for scaling and orchestration
- **Consul** for service discovery and configuration management

---
## 🏗️ Microservices Overview

### **User Service**
📌 [User Service README](./user/README.md)
Handles authentication, user profile, and account management.

### **Category Service**
📌 [Category Service README](./category/README.md)
Handles category management, including creating, updating, deleting, and finding categories by different attributes like ID, name, and level. It integrates with Consul for dynamic configuration and service discovery.

---
## 📂 Project Structure
```
shopping-platform/
├── common/                # Shared libraries and utilities
├── user/                  # User Service (Auth, Registration)
│   ├── domain/            # Business logic and models
│   ├── handler/           # gRPC Handlers
│   ├── proto/             # gRPC Protobuf Definitions
│   ├── README.md          # User Service Docs
│   ├── main.go            # Service entry point
│   ├── Dockerfile         # Service containerization
│   ├── go.mod, go.sum     # Go dependencies
├── category/              # Category Service (Categories)
│   ├── README.md
├── docker-compose.yml     # Multi-container setup for all services
├── Makefile               # Build automation commands
├── README.md              # Shopping Platform Docs
```

--- 
## 🚀 Getting Started
