# Development Roadmap

## Architecture Baseline v1

- Docker Compose-based local platform
- Kong as the external API gateway for north-south traffic
- PostgreSQL and Kafka as the core local platform dependencies
- OpenAPI contracts for REST APIs
- JSON Schema contracts for Kafka events
- `/healthz` and `/readyz` endpoints for service health
- Jenkins pipeline for the full local delivery flow
- Optional Prometheus/Grafana monitoring via compose override

## Roadmap v2

- Helm charts and Kubernetes manifests
- Kubernetes-on-VM deployment support
- Strimzi-managed Kafka for Kubernetes
- Prometheus and Grafana as first-class observability stack
- OpenTelemetry tracing support for service requests and Kafka events
- Expanded readiness checks and finer-grained metrics

## Future Enterprise Features

- GitOps workflow for Kubernetes deployments
- Enterprise schema registry and schema evolution strategy
- Retry mechanisms, dead-letter queue, and notification delivery resilience
- API management and gateway policy capabilities in Kong
- Developer self-service templates and component catalog
- Operational runbooks, dashboards, and platform team onboarding
