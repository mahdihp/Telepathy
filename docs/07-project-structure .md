# Project Structure

## Overview

Telepathy follows a Modular Monolith architecture.

Each module has its own application, domain, repository, and transport layer.

The project is organized to make future migration to Microservices simple.

---

# Directory Structure

```text
telepathy/

├── api/
│
├── cmd/
│   └── telepathy/
│       └── main.go
│
├── configs/
│
├── deployments/
│
├── docker/
│
├── docs/
│
├── internal/
│   │
│   ├── auth/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── transport/
│   │   └── service/
│   │
│   ├── user/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── transport/
│   │   └── service/
│   │
│   ├── conversation/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── transport/
│   │   └── service/
│   │
│   ├── message/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── transport/
│   │   └── service/
│   │
│   ├── presence/
│   │   ├── application/
│   │   ├── domain/
│   │   ├── repository/
│   │   ├── transport/
│   │   └── service/
│   │
│   └── shared/
│       ├── config/
│       ├── logger/
│       ├── middleware/
│       ├── errors/
│       ├── websocket/
│       └── utils/
│
├── pkg/
│
├── scripts/
│
├── test/
│
├── Makefile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

# Layers

## Domain

Contains business rules and entities.

No dependency on infrastructure.

---

## Application

Contains use cases.

Coordinates domain objects.

---

## Repository

Handles data persistence.

ScyllaDB implementation lives here.

---

## Transport

HTTP and WebSocket handlers.

Responsible for request/response mapping.

---

## Service

Coordinates module-specific operations.

---

# Shared Package

Contains reusable components.

* Logger
* Configuration
* Middleware
* Error Handling
* WebSocket Manager
* Utilities

---

# Entry Point

```text
cmd/telepathy/main.go
```

Application startup begins here.

---

# Configuration

All configuration files are stored in:

```text
configs/
```

---

# API

REST API definitions:

```text
api/
```

---

# Documentation

Project documentation:

```text
docs/
```

---

# Deployment

Docker and Kubernetes manifests:

```text
deployments/
```

---

# Scripts

Development scripts:

```text
scripts/
```

---

# Future Structure

When migrating to Microservices, each module can become an independent service without major structural changes.

Example:

```text
Auth Service

User Service

Conversation Service

Message Service

Presence Service
```
