# AgroCore

AgroCore is a backend application for agricultural management built with Go.

The project is based on real-world agricultural management requirements and
is being developed incrementally, with emphasis on maintainability,
testability, documentation, and explicit architectural decisions.

## Project Status

> Early development.

AgroCore currently provides its initial HTTP foundation, including
environment-based configuration, health checking, server timeouts,
graceful shutdown, and automated tests.

Features described in the roadmap should not be considered implemented
until they are marked as completed.

## Goals

AgroCore aims to provide backend capabilities for agricultural operations
while serving as a practical software engineering project.

The project will evolve from a small Go application into a production-oriented
backend, introducing architectural complexity only when required by actual
use cases.

## Current Technology Stack

- Go 1.27+
- Go Standard Library (`net/http`)
- Git
- Git Flow

No third-party runtime dependencies are currently required.

Additional technologies will be introduced only when they solve a concrete
technical or business requirement.

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

## HTTP API

Start the application:

```bash
go run .
```

By default, AgroCore listens on port `8080`.

### Health Check

```http
GET /health
```

Example:

```bash
curl http://localhost:8080/health
```

Response:

```json
{
  "status": "ok"
}
```

### Configuration

The HTTP server can be configured through environment variables.

| Variable | Default | Description |
| --- | --- | --- |
| `HTTP_HOST` | empty | HTTP interface/address to bind to |
| `HTTP_PORT` | `8080` | HTTP server port |

Example:

```bash
HTTP_PORT=9000 go run .
```

See [`docs/development/configuration.md`](docs/development/configuration.md)
for configuration details.

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
- [`docs/architecture/http-server.md`](docs/architecture/http-server.md) — HTTP server architecture.

## Architecture Decisions

Relevant technical decisions are documented using Architecture Decision
Records.

Current ADRs:

- [ADR 0001 — Use Go as the Backend Language](docs/decisions/0001-use-go.md)
- [ADR 0002 — Use Git Flow as the Branching Strategy](docs/decisions/0002-use-git-flow.md)
- [ADR 0003 — Use the Go Standard Library for the Initial HTTP Server](docs/decisions/0003-use-standard-library-http.md)

## Roadmap

### Foundation

- [x] Initialize Go module
- [x] Configure Git Flow
- [x] Establish project documentation
- [ ] Add automated quality checks
- [ ] Add continuous integration

### Backend

- [x] HTTP server
- [x] Health check endpoint
- [x] Environment-based HTTP configuration
- [x] Graceful shutdown
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