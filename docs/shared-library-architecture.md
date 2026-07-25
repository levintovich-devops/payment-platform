# Jenkins Shared Library Architecture

## Purpose

Design a reusable Jenkins Shared Library for the Payment Platform that is platform-oriented, configuration-driven, and extensible for future capabilities.

This document describes the architecture only; it does not implement pipelines or application changes.

---

## 1. Design Principles

- Reuse across services: one library, many pipelines.
- Externalize configuration: avoid hardcoded registry, repo, branch, credential, or parameter values.
- Separate concerns: configuration, versioning, registry abstraction, orchestration, and extension points.
- Keep pipeline entrypoints thin and declarative.

---

## 2. Proposed Directory Structure

```text
vars/
  platformPipeline.groovy

src/com/platform/jenkins/
  config/
    PlatformConfig.groovy
    RegistryConfig.groovy
  context/
    BuildContext.groovy
  version/
    VersionInfo.groovy
    VersionService.groovy
  registry/
    RegistryClient.groovy
    RegistryFactory.groovy
    DockerHubRegistry.groovy
  capabilities/
    Capability.groovy
    CapabilityRegistry.groovy
  stages/
    BuildStage.groovy
    TestStage.groovy
    DockerStage.groovy
    PublishStage.groovy

resources/
  platform/
    defaults.yml
```

---

## 3. Module Responsibilities

### vars/platformPipeline.groovy
Why it exists: entrypoint for Jenkins pipelines.

Responsibility:
- load shared configuration;
- calculate versions;
- create a build context;
- select the registry implementation;
- delegate execution to thin stage wrappers.

Interaction:
- calls the configuration, versioning, registry, capabilities, and stages modules.

### src/com/platform/jenkins/config/PlatformConfig.groovy
Why it exists: centralize platform-level configuration.

Responsibility:
- hold repository, branch, registry, credentials, and runtime settings;
- provide one object for all pipeline logic.

Interaction:
- consumed by versioning, image building, and registry selection.

### src/com/platform/jenkins/config/RegistryConfig.groovy
Why it exists: isolate registry-specific settings.

Responsibility:
- define registry type, URL, namespace, repository rules, and credentials reference.

Interaction:
- used by RegistryFactory and registry implementations.

### src/com/platform/jenkins/version/VersionService.groovy
Why it exists: centralize version calculation.

Responsibility:
- compute release version;
- compute build version;
- generate Docker tags for release, latest, Git SHA, and timestamp;
- incorporate Jenkins build number.

Interaction:
- returns VersionInfo to the build context and image naming layer.

### src/com/platform/jenkins/version/VersionInfo.groovy
Why it exists: make versions explicit and reusable.

Responsibility:
- hold structured version data as a typed value object.

Interaction:
- used by image naming and pipeline reporting.

### src/com/platform/jenkins/context/BuildContext.groovy
Why it exists: provide a single runtime context object.

Responsibility:
- aggregate configuration, versions, and derived metadata for the current build.

Interaction:
- passed to stages and capabilities.

### src/com/platform/jenkins/registry/RegistryClient.groovy
Why it exists: abstract registry operations behind a stable interface.

Responsibility:
- define common operations for authentication, tagging, image naming, and pushing images.
- keep image reference generation inside the registry abstraction because it depends on registry-specific rules.

Interaction:
- implemented by concrete registry providers.

### src/com/platform/jenkins/registry/RegistryFactory.groovy
Why it exists: select the correct registry implementation.

Responsibility:
- choose Docker Hub initially;
- support future migration to Artifactory, Harbor, AWS ECR, or similar registries.

Interaction:
- uses RegistryConfig to determine which registry implementation to instantiate.

### src/com/platform/jenkins/registry/DockerHubRegistry.groovy
Why it exists: support Docker Hub as the initial registry target.

Responsibility:
- implement the registry interface for Docker Hub.

Interaction:
- returned by RegistryFactory when Docker Hub is configured.

### src/com/platform/jenkins/capabilities/Capability.groovy
Why it exists: create a simple extension model.

Responsibility:
- define future capabilities such as Helm, Kubernetes, Argo CD, security scanning, artifact publishing, and notifications.

Interaction:
- implemented by optional modules and registered through CapabilityRegistry.

### src/com/platform/jenkins/capabilities/CapabilityRegistry.groovy
Why it exists: keep extension points modular.

Responsibility:
- register and resolve optional capabilities.

Interaction:
- used by the pipeline entrypoint when new capabilities are enabled.

### src/com/platform/jenkins/stages/BuildStage.groovy
Why it exists: isolate build-stage logic from the pipeline entrypoint.

Responsibility:
- implement the build step in a reusable stage component.

Interaction:
- invoked by platformPipeline as part of the build flow.

### src/com/platform/jenkins/stages/TestStage.groovy
Why it exists: isolate test-stage logic from the pipeline entrypoint.

Responsibility:
- implement the test step in a reusable stage component.

Interaction:
- invoked by platformPipeline as part of the validation flow.

### src/com/platform/jenkins/stages/DockerStage.groovy
Why it exists: isolate container build logic from the pipeline entrypoint.

Responsibility:
- implement Docker image build operations using the registry abstraction.

Interaction:
- invoked by platformPipeline after version and registry setup.

### src/com/platform/jenkins/stages/PublishStage.groovy
Why it exists: isolate publish logic from the pipeline entrypoint.

Responsibility:
- implement image publishing operations through the selected registry client.

Interaction:
- invoked by platformPipeline once the image is ready.

---

## 4. Configuration Model

### resources/platform/defaults.yml
Why it exists: provide baseline platform values without hardcoding them in Groovy.

Responsibility:
- define default registry, branch, naming conventions, and optional feature flags.

Interaction:
- loaded by PlatformConfig to form a usable runtime configuration.

---

## 5. Dependency Flow

```text
Jenkins job
  -> platformPipeline
  -> PlatformConfig
  -> VersionService
  -> VersionInfo
  -> BuildContext
  -> RegistryFactory
  -> RegistryClient / DockerHubRegistry
  -> BuildStage / TestStage / DockerStage / PublishStage
  -> capabilities
```

This keeps version logic centralized, registry behavior pluggable, and pipeline code minimal.

---

## 6. Recommended Initial Implementation Scope

The first implementation step should include only:
1. configuration layer;
2. versioning module;
3. shared-library entrypoint;
4. registry abstraction with Docker Hub support.

All other capabilities remain future extension points.
