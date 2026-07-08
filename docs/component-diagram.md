# Component Diagram

## Components

- **API Gateway (Kong)**
  - External edge gateway for north-south traffic
  - Handles routing, authentication, rate limiting, and observability hooks

- **Payment Service**
  - REST API for payment intent creation and capture
  - Persists payment history in PostgreSQL
  - Publishes domain events to Kafka
  - Exposes `/healthz` and `/readyz`

- **Notification Service**
  - Consumes Kafka events related to payment status
  - Persists notification records and delivery status
  - Emits downstream notifications (simulated for demo)
  - Exposes `/healthz` and `/readyz`

- **Kafka**
  - Event bus for domain events
  - Topics include `payment.intent.created`, `payment.intent.captured`, `payment.payment.updated`
  - Uses JSON Schema contracts for event payload validation

- **PostgreSQL**
  - Single shared instance for v1
  - Separate databases or schemas for `payment-service` and `notification-service`

- **Jenkins**
  - CI/CD orchestrator executing build, tests, image creation, deploy, and smoke tests
  - Validates the complete local delivery pipeline

- **Optional Monitoring Stack**
  - Prometheus + Grafana enabled through Docker Compose override
  - Provides observability without being required for core platform startup

## High-Level Flow

1. External clients call Kong on `/payments/intent` or `/payments/capture`.
2. Kong routes requests to `payment-service`.
3. `payment-service` persists history, updates state, and publishes Kafka events.
4. `notification-service` consumes Kafka events, writes notification records, and sends notifications.
5. Optional monitoring stack collects metrics and displays dashboards when enabled.
