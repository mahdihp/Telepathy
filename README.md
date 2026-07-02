# Telepathy
Message Platform - پیام رسان تله پاتی

# Project Vision

> A production-grade real-time messaging platform built with Go to learn distributed systems, ScyllaDB, and modern backend architecture.

---

# Vision

The goal of this project is **not** to build another Telegram or WhatsApp clone.

Instead, the objective is to design and implement a production-grade messaging platform that demonstrates modern backend engineering practices, distributed system design, and scalable architecture.

The project focuses on solving engineering problems rather than building UI features.

---

# Why This Project?

Real-time messaging systems combine many challenging backend concepts into a single product, including:

* High write throughput
* Low latency communication
* Real-time data synchronization
* Distributed architecture
* Horizontal scalability
* Event-driven communication
* Fault tolerance
* Efficient data modeling

This project serves as a hands-on environment to explore these concepts.

---

# Goals

The project aims to provide practical experience with:

* Go
* ScyllaDB
* WebSocket
* gRPC
* Event-Driven Architecture
* NATS JetStream
* Redis
* MinIO
* Docker
* Kubernetes
* OpenTelemetry
* Prometheus
* Grafana

---

# Design Principles

The system is designed around the following principles:

* Simplicity over unnecessary complexity
* Query-first data modeling
* Modular architecture
* Event-driven communication
* Horizontal scalability
* Production-ready code
* Observability by default
* Infrastructure as Code
* Clean Architecture

---

# Architecture Evolution

The project will evolve through multiple phases.

## Phase 1

Modular Monolith

Focus on:

* Domain Design
* ScyllaDB
* WebSocket
* Clean Architecture

---

## Phase 2

Distributed Services

Split modules into independent services.

Introduce:

* gRPC
* NATS JetStream
* Service-to-Service Communication

---

## Phase 3

Production Infrastructure

Deploy the platform using:

* Kubernetes
* Distributed Tracing
* Metrics
* Logging
* CI/CD

---

# Out of Scope

The following features are intentionally excluded from the early phases:

* Voice Calls
* Video Calls
* Stories
* Live Streaming
* AI Assistants
* Payment Systems
* End-to-End Encryption
* Multi-Region Deployment

These features may be explored in future versions.

---

# Success Criteria

The project will be considered successful when it demonstrates:

* Scalable backend architecture
* Efficient ScyllaDB data modeling
* Reliable real-time messaging
* Production-grade engineering practices
* Clear documentation
* Automated testing
* Easy deployment

---

# Target Audience

This repository is intended for:

* Backend Engineers
* Distributed Systems Enthusiasts
* Go Developers
* Recruiters
* Technical Interviewers
* Software Engineering Students

---

# Repository Philosophy

Every architectural decision in this repository should be documented.

Every database table should exist because of an access pattern.

Every module should have a single responsibility.

Every feature should prioritize correctness, simplicity, and maintainability over unnecessary complexity.

