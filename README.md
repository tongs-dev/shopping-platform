# Shopping Platform

## Overview
The **Shopping Platform** is a microservices-based e-commerce system designed to handle different aspects of online shopping, including user management, orders, inventory, and payments. The system is built using **Go**, **gRPC**, and **MySQL**, with each microservice handling a specific domain of the platform.

## Features
- **User Service**: Handles user authentication, registration, and profile management.
- **Category Service**: Manages the categories for products, including adding, updating, deleting, and retrieving categories.
- **gRPC Communication**: Services interact via **gRPC** for efficient communication.
- **Consul Integration**: Uses Consul for service discovery and configuration management, ensuring that microservices can dynamically register and find each other.
- **Tracing**: Uses **Jaeger** for distributed tracing and monitoring of service interactions.
- **Monitoring & Logging**: Includes Prometheus, Grafana, and centralized logging for observability.
- **Dockerized Deployment**: Uses **Docker & Docker Compose** for easy setup and scaling.

## Tech Stack
- **Go** (Golang) for backend services
- **gRPC** for inter-service communication
- **MySQL** for data persistence
- **GORM** as the ORM layer
- **Docker & Docker Compose** for containerization
- **Kubernetes (optional)** for scaling and orchestration
- **Consul** for service discovery and configuration management
- **Jaeger** for distributed tracing

---
## 🏗️ Microservices Overview

### **User Service**
📌 [User Service README](./user/README.md)
Handles authentication, user profile, and account management.

### **Category Service**
📌 [Category Service README](./category/README.md)
Handles category management, including creating, updating, deleting, and finding categories by different attributes like ID, name, and level. It integrates with Consul for dynamic configuration and service discovery.

### **Product Service**
📌 [Product Service README](./product/README.md)
Handles product management, including creating, updating, deleting, and finding products. It integrates with Consul for dynamic configuration and service discovery and Jaeger for distributed tracing and monitoring of product service interactions.

---
## 📂 Project Structure
```
shopping-platform/
├── common/                # Shared libraries and utilities
├── user/                  # User Service (Auth, Registration)
│   ├── domain/
│   ├── ...
├── category/              # Category Service (Categories)
│   ├── domain/
│   ├── ...
├── product/               # Product Service (Products)
│   ├── domain/
│   ├── ...
├── docker-compose.yml     # Multi-container setup for all services
├── Makefile               # Build automation commands
├── README.md              # Shopping Platform Docs
```

--- 
## 🚀 Getting Started
