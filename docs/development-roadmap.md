# Development Roadmap

## Phase 1: Reference Platform Delivery

- Implement Docker Compose-based local platform
- Build Go services: `payment-service` and `notification-service`
- Configure Kong as external API Gateway for north-south traffic
- Provision PostgreSQL and Kafka for local development
- Define OpenAPI contracts for REST APIs
- Define JSON Schema contracts for Kafka events
- Implement `/healthz` and `/readyz` for all services
- Build Jenkins pipeline for full local delivery flow
- Support optional Prometheus/Grafana monitoring via compose override

## Phase 2: Kubernetes and Observability

- Add Helm charts and Kubernetes manifests
- Transition local deployment support toward Kubernetes-on-VM
- Introduce Strimzi-managed Kafka for Kubernetes
- Add Prometheus and Grafana as first-class observability stack
- Add OpenTelemetry tracing support for service requests and Kafka events
- Expand readiness checks and finer-grained metrics

## Phase 3: Platform Engineering Maturity

- Add GitOps workflow for Kubernetes deployments
- Introduce enterprise schema registry and schema evolution strategy
- Add retry mechanisms, dead-letter queue, and notification delivery resilience
- Add API management / gateway policy examples in Kong
- Add developer self-service templates and component catalog
- Enhance operational runbooks, dashboards, and platform team onboarding
