# Changelog

All notable changes to AgroCore will be documented in this file.

The format is based on Keep a Changelog, and this project intends to follow
Semantic Versioning after the first release.

## [Unreleased]

### Added

- Initial Go module.
- Initial application entry point.
- Git Flow branching strategy.
- Initial project documentation.
- Architecture Decision Records for Go and Git Flow.
- HTTP server based on the Go standard library.
- Health check endpoint at `GET /health`.
- HTTP server timeout configuration.
- Graceful shutdown on `SIGINT` and `SIGTERM`.
- Environment-based HTTP server configuration.
- Validation for HTTP server port configuration.
- Automated tests for configuration, HTTP handlers, routing, and server setup.
- Architecture Decision Record for the initial HTTP implementation.