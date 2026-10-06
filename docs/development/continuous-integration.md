# Continuous Integration

AgroCore uses GitHub Actions to validate changes automatically.

## Purpose

Continuous Integration provides an automated quality gate for changes
submitted to the repository.

The CI workflow complements local development checks and helps prevent
invalid changes from being integrated into the main development branches.

## Triggers

The workflow runs for:

- pushes to `develop`;
- pushes to `main`;
- pull requests targeting `develop`;
- pull requests targeting `main`.

## Quality Checks

The CI pipeline executes the project's standard quality gate:

```bash
make check
```

The quality gate validates:

1. source code formatting;
2. `go vet`;
3. automated tests;
4. the Go race detector;
5. application build.

Using the same quality gate locally and in CI reduces differences between
developer environments and automated validation.

## Local Validation

Before pushing changes, developers should run:

```bash
make check
```

Formatting can be fixed automatically with:

```bash
make fmt
```

CI validates the project but does not modify incorrectly formatted source
code.

## Dependencies

The workflow uses GitHub Actions for repository checkout and Go environment
setup.

These Actions are CI infrastructure dependencies and are not runtime
dependencies of the AgroCore application.