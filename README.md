# AgroCore

AgroCore is a backend application for agricultural management, developed in Go
with a focus on clean architecture, modularity, maintainability, testing, and
professional software engineering practices.

The project is also an evolving engineering case study that demonstrates the
construction of a production-oriented backend from its foundations to more
advanced topics such as persistence, authentication, multi-tenancy, messaging,
observability, CI/CD, and cloud infrastructure.

> AgroCore is currently under active development.

## Project Goals

AgroCore aims to provide a modular backend for agricultural management while
serving as a practical demonstration of backend engineering with Go.

The project evolves incrementally from simple foundations to production-oriented
architecture.

The main goals are:

- build a modular agricultural management backend;
- keep business rules independent from infrastructure;
- design explicit application boundaries;
- provide a REST API;
- use professional automated testing practices;
- document important architectural decisions;
- automate quality checks with continuous integration;
- introduce infrastructure only when concrete requirements justify it;
- evolve toward a production-ready application.

## Current Status

The project currently provides:

- Go application bootstrap;
- environment-based configuration;
- HTTP server foundation;
- graceful shutdown;
- health endpoint;
- HTTP server timeouts;
- automated unit tests;
- race detector validation;
- continuous integration;
- development tooling through Make;
- modular application architecture;
- shared business date value object;
- Agricultural Year domain entity;
- Agricultural Year creation use case;
- repository abstraction;
- in-memory Agricultural Year repository;
- internal integration testing;
- Architecture Decision Records.

The first agricultural business capability currently being implemented is
Agricultural Year.

## Technology Stack

Current technologies:

- Go;
- Go standard library;
- `net/http`;
- `context`;
- `testing`;
- `httptest`;
- Git;
- Git Flow;
- GNU Make;
- GitHub Actions.

Planned technologies will be introduced only when required by concrete
features.

Examples include:

- PostgreSQL;
- Redis;
- Docker;
- OpenAPI;
- authentication;
- multi-tenancy;
- messaging;
- observability;
- cloud infrastructure;
- Kubernetes.

These technologies are roadmap items, not current runtime dependencies.

## Architecture

AgroCore starts as a modular monolith.

Code is organized primarily around business capabilities rather than around
large global technical layers.

The architecture evolves incrementally.

Current high-level structure:

```text
.
├── .github/
│   └── workflows/
│       └── ci.yml
│
├── docs/
│   ├── architecture/
│   │   ├── agricultural/
│   │   │   └── agricultural-year.md
│   │   ├── http-server.md
│   │   └── README.md
│   │
│   ├── decisions/
│   │   ├── 0001-use-go.md
│   │   ├── 0002-use-git-flow.md
│   │   ├── 0003-use-standard-library-http.md
│   │   ├── 0004-use-modular-application-architecture.md
│   │   └── 0005-use-date-value-object.md
│   │
│   └── development/
│       ├── configuration.md
│       ├── continuous-integration.md
│       ├── git-workflow.md
│       └── project-tooling.md
│
├── internal/
│   ├── agricultural/
│   │   └── agriculturalyear/
│   │       ├── memory/
│   │       │   ├── repository.go
│   │       │   └── repository_test.go
│   │       ├── agricultural_year.go
│   │       ├── agricultural_year_test.go
│   │       ├── create.go
│   │       ├── create_test.go
│   │       └── create_integration_test.go
│   │
│   ├── config/
│   ├── httpserver/
│   │
│   └── shared/
│       └── date/
│           ├── date.go
│           └── date_test.go
│
├── .editorconfig
├── .env.example
├── .gitignore
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── Makefile
├── README.md
├── go.mod
└── main.go
```

More details are available in:

```text
docs/architecture/README.md
```

## Agricultural Year

Agricultural Year is the first business module implemented in AgroCore.

The current entity contains:

```text
AgriculturalYear
├── ID
├── Name
├── OpeningDate
├── ClosingDate
└── Status
```

New entities are created without a persistence identifier.

The persistence adapter is responsible for assigning the ID.

Example:

```text
New("2026/2027")

ID = 0
        |
        v
Repository.Save()
        |
        v
ID = 1
```

Opening and closing dates are optional.

Business dates use the shared `Date` value object and intentionally represent
only:

- year;
- month;
- day.

They do not expose time-of-day or timezone semantics.

See:

```text
docs/architecture/agricultural/agricultural-year.md
```

## Application Flow

The current Agricultural Year creation flow is:

```text
CreateInput
    |
    v
Creator.Execute
    |
    v
AgriculturalYear.New
    |
    v
Repository.ExistsByName
    |
    v
Repository.Save
    |
    v
Persisted AgriculturalYear
```

The application depends on the repository interface rather than on a concrete
persistence technology.

The current implementation uses an in-memory adapter.

PostgreSQL persistence will be introduced later without requiring the
application use case to depend directly on PostgreSQL.

## HTTP Server

AgroCore currently uses Go's standard `net/http` package.

The server includes:

- configurable host;
- configurable port;
- HTTP timeouts;
- health endpoint;
- graceful shutdown.

Current endpoint:

```text
GET /health
```

Expected response:

```json
{
  "status": "ok"
}
```

The health endpoint is infrastructure-oriented and does not depend on
agricultural business modules.

## Configuration

Configuration is loaded from environment variables.

Current variables:

```text
HTTP_HOST=
HTTP_PORT=8080
```

An example is available in:

```text
.env.example
```

The application intentionally does not automatically load `.env`.

Environment injection is considered an infrastructure responsibility.

## Getting Started

### Requirements

Install:

- Go;
- Git;
- Git Flow;
- GNU Make.

Docker will become a requirement when persistent infrastructure is introduced.

Check the Go installation:

```bash
go version
```

Check Git:

```bash
git --version
```

Check Git Flow:

```bash
git flow version
```

Check Make:

```bash
make --version
```

## Running the Application

Run directly:

```bash
go run .
```

Or:

```bash
make run
```

By default, the server listens on:

```text
:8080
```

Test the health endpoint:

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{"status":"ok"}
```

## Development Commands

Format Go source files:

```bash
make fmt
```

Check formatting:

```bash
make fmt-check
```

Run static analysis:

```bash
make vet
```

Run tests:

```bash
make test
```

Run tests with the race detector:

```bash
make test-race
```

Build the application:

```bash
make build
```

Run all project quality checks:

```bash
make check
```

The quality pipeline currently validates:

```text
formatting
    |
    v
go vet
    |
    v
tests
    |
    v
race detector
    |
    v
build
```

## Testing

AgroCore uses Go's standard testing tools.

Current testing includes:

- domain unit tests;
- application unit tests;
- configuration tests;
- HTTP handler tests;
- HTTP server tests;
- repository tests;
- internal integration tests;
- race detector validation.

Run all tests:

```bash
go test ./...
```

Run tests with detailed output:

```bash
go test -v ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run the race detector:

```bash
go test -race ./...
```

## Continuous Integration

GitHub Actions executes the project quality pipeline for changes targeting the
main development branches.

The workflow runs:

```bash
make check
```

This keeps local and CI validation aligned.

See:

```text
docs/development/continuous-integration.md
```

## Git Workflow

AgroCore uses Git Flow.

Main branches:

```text
main
develop
```

Feature branches use:

```text
feature/<feature-name>
```

Example:

```bash
git flow feature start agricultural-year
```

The normal development flow is:

```text
feature branch
      |
      v
Pull Request
      |
      v
CI
      |
      v
Squash and Merge
      |
      v
develop
```

See:

```text
docs/development/git-workflow.md
```

## Commit Convention

The project uses Conventional Commits.

Examples:

```text
feat: add agricultural year creation
fix: validate HTTP port
docs: document agricultural year architecture
test: add agricultural year repository tests
refactor: simplify server initialization
chore: update development tooling
ci: add continuous integration workflow
```

Common types:

```text
feat
fix
docs
test
refactor
perf
build
ci
chore
```

## Architecture Decision Records

Important technical decisions are documented as ADRs.

Current ADRs:

```text
ADR 0001 - Use Go as Backend Language
ADR 0002 - Use Git Flow as Branching Strategy
ADR 0003 - Use Standard Library HTTP
ADR 0004 - Use a Modular Application Architecture
ADR 0005 - Represent Business Dates Without Time-of-Day Semantics
```

ADRs are stored under:

```text
docs/decisions
```

## Development Principles

AgroCore follows these principles:

- business capabilities before technical layers;
- explicit dependencies;
- small interfaces;
- interfaces close to their consumers;
- infrastructure isolated from business logic;
- tests as part of implementation;
- automated quality checks;
- documented architectural decisions;
- incremental architecture;
- no speculative abstractions;
- no premature microservices.

## Roadmap

The project will evolve incrementally.

### Foundation

- [x] Go project bootstrap
- [x] Git Flow
- [x] project documentation
- [x] HTTP server
- [x] environment configuration
- [x] health endpoint
- [x] graceful shutdown
- [x] automated tests
- [x] CI
- [x] development tooling
- [x] modular architecture
- [x] date-only business value object

### Agricultural Domain

- [x] Agricultural Year foundation
- [ ] Culture
- [ ] Crop
- [ ] Crop and Culture relationship
- [ ] Variety Culture
- [ ] Owner
- [ ] Producer
- [ ] Farm
- [ ] Field
- [ ] Supplier types
- [ ] Suppliers
- [ ] Supplier bank information
- [ ] Warehouses
- [ ] Lanyards
- [ ] Drivers

Additional agricultural modules will be introduced according to the project's
business model and migration sequence.

### Persistence

- [ ] PostgreSQL
- [ ] database migrations
- [ ] PostgreSQL repository adapters
- [ ] integration tests with PostgreSQL
- [ ] database constraints and indexes

### API

- [ ] business HTTP handlers
- [ ] request validation
- [ ] consistent error responses
- [ ] JSON date serialization
- [ ] OpenAPI documentation
- [ ] API versioning strategy

### Security

- [ ] authentication
- [ ] authorization
- [ ] multi-tenancy
- [ ] tenant isolation
- [ ] security hardening

### Infrastructure

- [ ] Docker
- [ ] Docker Compose
- [ ] Redis
- [ ] background processing
- [ ] messaging
- [ ] observability
- [ ] structured application metrics

### Delivery

- [ ] production CI/CD
- [ ] cloud deployment
- [ ] container orchestration
- [ ] Kubernetes when justified by operational requirements

## Documentation

Project documentation is organized under:

```text
docs/
├── architecture/
├── decisions/
└── development/
```

Architecture documentation explains how the system is structured.

ADRs explain why important technical decisions were made.

Development documentation explains how contributors build, test, configure,
and evolve the project.

## Contributing

Contribution guidelines are available in:

```text
CONTRIBUTING.md
```

Before submitting changes, run:

```bash
make check
git diff --check
```

## License

See:

```text
LICENSE
```