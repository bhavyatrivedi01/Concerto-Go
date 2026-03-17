# Concerto-Go: Automated Model-to-Go Pipeline

This repository hosts a specialized CI/CD environment designed to automate the lifecycle of Accord Project Concerto models. It focuses on the seamless transformation of Concerto Modeling Language (.cto) files into production-ready Go source code.

## Project Evolution

### Proof of Concept (Local Development)
The project began as an exploration of the Concerto ecosystem. By following the Concerto Go CodeGen specifications, I successfully established a local environment to manually generate Go components. This stage was critical for understanding the underlying mapping between Concerto namespaces and Go package structures.

### Infrastructure as Code (Dockerization)
To eliminate "it works on my machine" inconsistencies, I engineered a custom Docker environment based on `golang:1.25-alpine`. This containerized solution bundles the `concerto-cli` with a native Go build-stack, ensuring a deterministic and isolated environment for code generation.

### CI/CD Integration (GitHub Actions)
The current architecture features an automated GitHub Actions Pipeline. Upon every commit, the system:
* Initializes the specialized Docker build environment.
* Compiles the provided `model.cto` into Go source code.
* Verifies the syntactical integrity and compilability of the generated assets.

## Architecture and Tech Stack

* **Logic**: Concerto Modeling Language (.cto)
* **Engine**: @accordproject/concerto-cli
* **Environment**: Docker (Alpine Linux / Go 1.25 / Node.js)
* **Automation**: GitHub Actions

## Future Roadmap (GSoC Objectives)

Moving forward, I aim to expand this pipeline into a comprehensive verification framework:

* **Dynamic Test Injection**: Implementing automated stages to generate and execute Go test cases, specifically focusing on JSON serialization/deserialization integrity for generated structs.
* **Artifact Management**: Configuring the pipeline to package and export the generated .go files as versioned build artifacts for external consumption.
* **Parameterized Pipelines**: Enhancing the workflow to accept arbitrary .cto files as inputs, allowing the community to use this pipeline as a "CodeGen-as-a-Service" for their own models.
* **Static Analysis**: Integrating linting tools (e.g., staticcheck) to ensure the generated code adheres to idiomatic Go standards.
