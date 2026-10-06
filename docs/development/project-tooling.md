# Project Tooling

AgroCore provides a Makefile as the standard interface for common
development and quality-assurance tasks.

## Requirements

Local development currently requires:

- Go 1.27 or later;
- GNU Make;
- Git.

Docker is not required at the current stage of the project.

## Available Commands

Display the available development commands:

```bash
make help
```

### Formatting

Format all Go source files:

```bash
make fmt
```

Check formatting without modifying files:

```bash
make fmt-check
```

The `fmt-check` command is intended for automated validation and CI
environments.

### Static Analysis

Run Go static analysis:

```bash
make vet
```

### Tests

Run the automated test suite:

```bash
make test
```

Run the test suite with the Go race detector:

```bash
make test-race
```

### Build

Build the application:

```bash
make build
```

The generated executable is stored at:

```text
bin/agrocore
```

The `bin` directory contains generated artifacts and is not tracked by Git.

### Run

Run the application directly from the source code:

```bash
make run
```

The default HTTP configuration is used unless environment variables are
provided.

Example:

```bash
HTTP_PORT=9000 make run
```

### Clean

Remove generated build artifacts:

```bash
make clean
```

### Quality Gate

Run the complete local quality gate:

```bash
make check
```

The quality gate currently executes:

1. source code formatting validation;
2. `go vet`;
3. automated tests;
4. tests with the Go race detector;
5. application build.

Developers should run `make check` before pushing changes or opening a
pull request.

## Continuous Integration

The GitHub Actions CI workflow uses the same `make check` quality gate used
during local development.

This helps keep local validation and CI behavior consistent.