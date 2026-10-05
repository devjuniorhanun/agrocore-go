# HTTP Server

## Status

Initial implementation.

## Purpose

The HTTP server provides the network entry point for the AgroCore API.

The initial implementation uses Go's standard `net/http` package and does
not depend on an external HTTP framework or router.

## Responsibilities

The HTTP layer is responsible for:

- accepting HTTP requests;
- routing requests to handlers;
- writing HTTP responses;
- applying server-level timeouts;
- exposing application health information;
- supporting graceful shutdown.

## Initial Endpoints

| Method | Path | Description |
| --- | --- | --- |
| GET | `/health` | Reports whether the application process is running. |

## Health Response

A successful request returns:

```json
{
  "status": "ok"
}

## Server Lifecycle

The application handles `SIGINT` and `SIGTERM` to provide graceful HTTP
server shutdown.

```text
Application start
      |
      v
Load configuration
      |
      v
Create HTTP server
      |
      v
ListenAndServe
      |
      +--------------------+
      |                    |
      v                    v
SIGINT / SIGTERM      Server failure
      |
      v
Graceful shutdown
      |
      v
Application exit