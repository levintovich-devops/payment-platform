# Architecture Overview

## Purpose

This reference architecture demonstrates a cloud-native payment platform built with Platform Engineering principles. It is designed for a single-platform-engineer mono-repository, local-first development, and future Kubernetes deployment.

## Key Principles

- Cloud-agnostic architecture
- Mono-repo for services, infrastructure, manifests, and documentation
- Go-based backend services
- API-first design with OpenAPI
- REST/JSON for internal service communication
- Apache Kafka for event streaming with JSON Schema contracts
- Kong as the external API gateway for north-south traffic only
- Separate service ownership of data with shared PostgreSQL instance and separate databases/schemas
- Jenkins-based CI/CD pipeline that validates local delivery into Docker Compose
- Lightweight optional monitoring via Docker Compose override
- Health endpoints with both `/healthz` and `/readyz`

## Components

- `api-gateway`: Kong gateway managing external client requests
- `payment-service`: handles payment intents, captures, history persistence, and event production
- `notification-service`: consumes Kafka events, persists notification status, and forwards notifications
- `postgres`: shared PostgreSQL instance with per-service database/schema separation
- `kafka`: Apache Kafka event bus for platform events
- `jenkins`: CI/CD orchestrator for build, test, image creation, deploy, and smoke tests

## Local Development Model

The first version uses Docker Compose as the canonical local platform. The repository supports a layered compose design so monitoring and future platform components can be enabled with optional override files.

## Observability and Operational Visibility

- Application logs are structured and developer-friendly
- Health endpoints support Kubernetes-style liveness and readiness checks
- Monitoring stack is optional and enabled separately via compose overrides

## Approved Architecture Review Scope

### Architecture Baseline v1
- Docker Compose-based local platform
- Kong as the external API gateway for north-south traffic
- REST/JSON internal service communication
- Kafka event streaming with JSON Schema contracts
- Separate service data ownership with a shared PostgreSQL instance
- Health endpoints with both `/healthz` and `/readyz`
- Jenkins-based CI/CD for the full local delivery flow
- Optional monitoring via Docker Compose override

### Roadmap v2
- Kubernetes deployment path with Helm charts and manifests
- Strimzi-managed Kafka for Kubernetes
- Prometheus and Grafana as first-class observability
- OpenTelemetry tracing support
- Expanded readiness checks and finer-grained metrics

### Future Enterprise Features
- GitOps workflow for Kubernetes deployments
- Enterprise schema registry and schema evolution strategy
- Retry, dead-letter queue, and notification resilience patterns
- Additional API management and gateway policy capabilities
- Developer self-service templates and operational runbooks

## Future Transition

The platform is architected for a later move to Kubernetes with Helm charts and a potential Strimzi-managed Kafka deployment. The initial implementation prioritizes a clear, complete developer workflow in a local environment.
