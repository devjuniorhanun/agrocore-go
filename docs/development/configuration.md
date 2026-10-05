# Configuration

AgroCore follows environment-based application configuration.

The application currently uses only the Go standard library to read and
validate configuration.

## Environment Variables

| Variable | Required | Default | Description |
| --- | --- | --- | --- |
| `HTTP_HOST` | No | empty | HTTP interface/address to bind to |
| `HTTP_PORT` | No | `8080` | HTTP server port |

`HTTP_PORT` must contain an integer between `1` and `65535`.

## Examples

Run with the default configuration:

```bash
go run .
```

Run on port `9000`:

```bash
HTTP_PORT=9000 go run .
```

Bind explicitly to localhost:

```bash
HTTP_HOST=127.0.0.1 HTTP_PORT=8080 go run .
```

## `.env.example`

The `.env.example` file documents the available environment variables.

AgroCore does not currently load `.env` files automatically. Environment
variables must be provided by the shell, container runtime, deployment
platform, or another process-level configuration mechanism.

## Invalid Configuration

AgroCore validates configuration during application startup.

Invalid values cause the application to terminate before starting the HTTP
server.