# ADR 0004: Use a Modular Application Architecture

## Status

Accepted

## Context

AgroCore will grow to support multiple agricultural business capabilities.

Organizing the entire application only by technical layers can cause unrelated
business concepts to become coupled through large global packages.

Starting with microservices would introduce operational and architectural
complexity before the application has demonstrated a need for independent
services.

## Decision

AgroCore will start as a modular monolith organized primarily around business
capabilities.

Domain, application, and infrastructure responsibilities will be separated
when concrete requirements justify those boundaries.

Dependencies must point toward business and application logic.

Domain code must not depend directly on HTTP, databases, messaging systems, or
other infrastructure concerns.

Interfaces will be introduced at useful architectural boundaries and should be
defined close to their consumers when appropriate.

The project will evolve incrementally. Empty modules, speculative abstractions,
and infrastructure that is not yet required should not be introduced merely to
match an architectural diagram.

## Consequences

### Positive

- Business capabilities remain easier to identify.
- Domain logic remains independent from infrastructure.
- Modules can evolve independently inside the monolith.
- Testing business behavior does not require external infrastructure.
- Future architectural boundaries can emerge from real requirements.

### Negative

- Developers must maintain clear module boundaries.
- Some architectural decisions are intentionally deferred.
- Additional discipline is required to avoid coupling between modules.

## Alternatives Considered

### Global Layer-First Architecture

Rejected as the primary organization because large global controller, service,
repository, and model packages can mix unrelated business capabilities.

### Microservices From the Beginning

Rejected because the current system does not require the operational complexity
of independently deployed services.