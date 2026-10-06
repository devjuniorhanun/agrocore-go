# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog, and this project intends to follow
Semantic Versioning when versioned releases begin.

## [Unreleased]

### Added

- Initial Go application bootstrap.
- Environment-based HTTP configuration.
- HTTP server foundation using the Go standard library.
- Health endpoint.
- Graceful HTTP server shutdown.
- HTTP server timeout configuration.
- Automated unit tests.
- Race detector validation.
- Continuous integration with GitHub Actions.
- Project development tooling with GNU Make.
- Initial architecture documentation.
- Architecture Decision Records.
- Initial Agricultural Year domain entity.
- Agricultural Year creation use case.
- Optional opening and closing business dates.
- Shared date value object for date-only business values.
- Agricultural Year repository contract.
- In-memory Agricultural Year repository.
- Agricultural Year unit tests.
- Agricultural Year repository tests.
- Agricultural Year internal integration test.
- Architecture documentation for the Agricultural Year module.
- ADR for modular application architecture.
- ADR for date-only business values.

### Changed

- Continuous integration now uses the project Makefile quality pipeline.
- Project formatting tooling targets Go source files explicitly.
- Architecture documentation now describes the modular business structure.
- Agricultural Year identifiers are assigned by persistence rather than create
  input.
- Agricultural Year creation and restoration are represented as separate
  domain operations.