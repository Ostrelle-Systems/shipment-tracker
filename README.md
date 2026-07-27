# shipment-tracker

Core shipment events service for the Ostrelle Systems logistics platform. Ingests carrier scan events, tracks shipment state transitions, and publishes status updates consumed by `dispatch` and `carrier-portal`.

**Stack:** Go 1.26, SQLite/PostgreSQL, REST + WebSocket streaming

## Prerequisites

- Go 1.26 or later
- A local Postgres instance (or SQLite for lightweight dev)

## Local setup

```bash
go mod download
cp gotify-server.env.example .env   # fill in DB connection + auth secrets
go run app.go serve
```

The service listens on `:80` by default; override with `GOTIFY_SERVER_PORT` (see `.env`).

## Testing

```bash
make test          # go tests + coverage + JS tests
make check-go      # lint/vet
```

## Project layout

- `api/` — HTTP handlers
- `auth/` — session and token auth
- `database/` — persistence layer
- `ui/` — internal ops UI (embedded at build time)

## Related services

- `dispatch` — consumes shipment events for SLA monitoring
- `carrier-portal` — surfaces shipment status to customers
- `route-engine` — triggers workflow steps on status changes
