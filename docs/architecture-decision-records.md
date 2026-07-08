# Architecture Decision Records (ADRs)

## ADR 001: Repository Strategy
- **Decision**: Use a mono-repository.
- **Rationale**: Simplifies architecture visibility, CI/CD, dependency management, and developer experience for a reference platform managed by a single platform engineer.

## ADR 002: Service Runtime
- **Decision**: Use Go for all backend services.
- **Rationale**: Aligns with cloud-native and Kubernetes ecosystems, simplifies deployment, and reflects modern platform engineering practices.

## ADR 003: Service Communication
- **Decision**: Use REST/JSON for internal service communication in v1.
- **Rationale**: Easier to demonstrate, debug, document, and integrate with API Gateway.

## ADR 004: Event Contracts
- **Decision**: Use JSON Schema for Kafka event contracts.
- **Rationale**: Prioritizes architectural clarity, readability, and ease of debugging; allows future upgrade to Avro/Protobuf.

## ADR 005: API Gateway Role
- **Decision**: Kong is used only as the external edge gateway for north-south traffic.
- **Rationale**: Simplifies architecture and reduces internal latency by allowing internal services to communicate directly.

## ADR 006: CI/CD Scope
- **Decision**: Jenkins pipeline includes full local delivery flow: build, unit tests, Docker image build, image validation, deployment to Docker Compose, and smoke tests.
- **Rationale**: Demonstrates complete platform engineering workflow and validates application delivery end-to-end.

## ADR 007: Notification Service Behavior
- **Decision**: `notification-service` persists notification history and delivery status.
- **Rationale**: Provides operational visibility and auditability while keeping v1 simple.

## ADR 008: Data Ownership
- **Decision**: Each service owns its own data.
- **Rationale**: Preserves separation of concerns; v1 uses one PostgreSQL instance with separate databases/schemas for each service.

## ADR 009: Docker Compose Design
- **Decision**: Use layered Docker Compose design.
- **Rationale**: Supports a simple single-command local start and optional overrides for monitoring and future components.

## ADR 010: Health Model
- **Decision**: Implement both `/healthz` and `/readyz` endpoints.
- **Rationale**: Aligns with Kubernetes best practices and enables meaningful smoke tests from v1.

## ADR 011: Observability
- **Decision**: Core platform runs with logs and health checks; Prometheus/Grafana are optional via compose override.
- **Rationale**: Keeps developer experience simple while demonstrating observability as a first-class capability.

## ADR 012: API Contracts
- **Decision**: Use an API-first approach with OpenAPI definitions from the start.
- **Rationale**: Ensures API contracts are central to architecture, improving consistency, documentation, testing, and future API Gateway integration.


