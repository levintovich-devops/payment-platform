# Repository Structure

The repository is organized to support a mono-repo Platform Engineering reference architecture.

## Top-level Layout

- `docs/` — architecture documentation and decision records
- `services/` — application service source trees
  - `payment-service/`
  - `notification-service/`
- `infrastructure/` — infrastructure definitions and environment bootstrap files
- `k8s/` — Kubernetes manifests and resources
- `helm/` — Helm charts for future Kubernetes deployment
- `ci/` — Jenkins pipeline definitions and CI utilities
- `deploy/` — Docker Compose files and environment overrides
- `schemas/` — JSON Schema contracts for Kafka events
- `openapi/` — OpenAPI API contract definitions
- `README.MD` — repository overview and developer instructions

## Notes

- `services/` contains service-oriented code and service-specific configs.
- `deploy/` contains the canonical local Docker Compose stack plus optional override files for monitoring and additional platform components.
- `ci/` holds Jenkins pipeline files and scripts to execute the full local delivery flow.
- `k8s/` and `helm/` are ready for the Kubernetes migration path; v1 keeps them as architecture artifacts.
- `schemas/` stores event contract artifacts for Kafka topics.
- `openapi/` stores API-first REST contract definitions that drive service design.
