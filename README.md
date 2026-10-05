# AgroCore

AgroCore is a backend application for agricultural management built with Go.

The project is based on real-world agricultural management requirements and
is being developed incrementally, with emphasis on maintainability,
testability, documentation, and explicit architectural decisions.

## Project Status

> Early development.

The project is currently in its bootstrap phase. Features described in the
roadmap are planned capabilities and should not be considered implemented
until they are marked as completed.

## Goals

AgroCore aims to provide backend capabilities for agricultural operations
while serving as a practical software engineering project.

The project will evolve from a small Go application into a production-oriented
backend, introducing architectural complexity only when required by actual
use cases.

## Current Technology Stack

- Go 1.27+
- Git
- Git Flow

Additional technologies will be introduced only when they solve a concrete
technical or business requirement.

## Getting Started

### Requirements

- Go 1.27 or newer
- Git

Check the installed Go version:

```bash
go version
```

### Clone

```bash
git clone https://github.com/devjuniorhanun/agrocore-go.git
cd agrocore-go
```

> The remote repository may not be available until the initial project
> bootstrap is published.

### Run

```bash
go run .
```

Expected output:

```text
AgroCore API
```

### Build

```bash
go build -o agrocore .
```

Run the compiled binary:

```bash
./agrocore
```

## Development

AgroCore uses Git Flow.

Permanent branches:

- `main` — production-ready code.
- `develop` — integration branch for ongoing development.

Supporting branches:

- `feature/*`
- `bugfix/*`
- `release/*`
- `hotfix/*`

Development workflow details are documented in
[`docs/development/git-workflow.md`](docs/development/git-workflow.md).

## Documentation

Project documentation is maintained alongside the source code.

- [`docs/architecture/`](docs/architecture/) — architecture documentation.
- [`docs/decisions/`](docs/decisions/) — Architecture Decision Records (ADRs).
- [`docs/development/`](docs/development/) — development workflows and conventions.

## Architecture Decisions

Relevant technical decisions are documented using Architecture Decision
Records.

Current ADRs:

- [ADR 0001 — Use Go as the Backend Language](docs/decisions/0001-use-go.md)
- [ADR 0002 — Use Git Flow as the Branching Strategy](docs/decisions/0002-use-git-flow.md)

## Roadmap

### Foundation

- [x] Initialize Go module
- [x] Configure Git Flow
- [x] Establish project documentation
- [ ] Add automated quality checks
- [ ] Add continuous integration

### Backend

- [ ] HTTP server
- [ ] REST API
- [ ] Domain modeling
- [ ] PostgreSQL persistence
- [ ] Database migrations
- [ ] Authentication and authorization
- [ ] Multi-tenancy
- [ ] Redis integration
- [ ] Automated tests
- [ ] OpenAPI documentation

### Platform

- [ ] Docker
- [ ] Structured logging
- [ ] Metrics
- [ ] Distributed tracing
- [ ] Asynchronous messaging
- [ ] gRPC
- [ ] CI/CD
- [ ] Kubernetes
- [ ] Cloud deployment

The roadmap may change as the project requirements evolve.

## Contributing

Contribution guidelines are available in
[`CONTRIBUTING.md`](CONTRIBUTING.md).

## License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE)
for details.