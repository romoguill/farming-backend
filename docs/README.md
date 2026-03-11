# Farming App — Backend

## 1. Overview

The backend is a Go HTTP API that powers the farming app. It manages farms, crops, and related domain data, exposing a REST interface consumed by the frontend. It is designed for clarity and maintainability: idiomatic Go, layered architecture, explicit dependencies, and plain SQL for database access.

---

## 2. Architecture

The backend follows a layered architecture with a strict dependency direction:

```
handler → service → repository → database
```

- **handler**: Parses HTTP requests, validates input, calls the service layer, and writes responses. No business logic here.
- **service**: Contains business logic and orchestrates calls to one or more repositories.
- **repository**: Data access only. Wraps sqlc-generated code behind interfaces.
- **database**: PostgreSQL, accessed exclusively through the repository layer.

### Key design principles

- **Interfaces at the consumer**: Interfaces are defined where they are used (e.g., in `service/`), not where they are implemented. This follows the Go idiom and keeps packages decoupled.
- **Dependency injection via constructors**: Dependencies are passed explicitly through `NewXxx(dep Dep) *Xxx` constructor functions. No globals, no init tricks.
- **Explicit error handling**: Errors are returned and handled at each layer. No panics for control flow.
- **Context propagation**: `context.Context` is threaded through all function calls that touch I/O (DB, external services).

---

## 3. Project Structure

```
farming-backend/
├── cmd/
│   └── api/
│       └── main.go           # Entry point: wires dependencies, starts server
├── internal/
│   ├── config/               # Loads env vars into a typed Config struct
│   ├── domain/               # Core business types (pure structs, no DB/HTTP tags)
│   ├── handler/              # HTTP layer: parse → validate → call service → respond
│   │   ├── farm/
│   │   ├── crop/
│   │   └── middleware/
│   ├── repository/           # Data access layer (wraps sqlc-generated code)
│   │   └── postgres/
│   └── service/              # Business logic, orchestration
├── db/
│   ├── migrations/           # golang-migrate SQL files (up/down pairs)
│   └── queries/              # sqlc .sql query files
├── docs/
│   └── README.md
├── docker/
│   └── Dockerfile            # Multi-stage build
├── docker-compose.yml        # Local dev: app + postgres
├── .github/
│   └── workflows/
│       └── ci.yml
├── sqlc.yaml
├── Makefile
├── go.mod
└── go.sum
```

All application code lives under `internal/` (not importable by external packages). The `cmd/api/main.go` entry point is the only place where all layers are wired together.

---

## 4. Tech Stack

| Concern | Tool | Notes |
|---|---|---|
| HTTP router | Gin | Request routing + middleware |
| Database | PostgreSQL 17 | Primary datastore |
| SQL queries | sqlc | Type-safe query generation from `.sql` files |
| Migrations | golang-migrate | Plain SQL up/down files |
| Driver | pgx/v5 | PostgreSQL driver |
| Config | godotenv + os.Getenv | Simple env-var loading |
| Testing | testify + testcontainers-go | Assertions + real DB in tests |

---

## 5. Database: sqlc Workflow

[sqlc](https://sqlc.dev) generates type-safe Go code from plain SQL queries. The workflow is:

1. Write SQL queries in `db/queries/*.sql`, annotated with sqlc comments:
   ```sql
   -- name: GetFarm :one
   SELECT * FROM farms WHERE id = $1;
   ```
2. Run `sqlc generate` — produces Go structs and query functions in `internal/repository/postgres/`.
3. Wrap the generated code in repository structs that implement the interfaces defined in the service layer.
4. `sqlc.yaml` at the repo root configures the database engine, output paths, and Go package names.

Never edit the generated files directly. All changes go through the `.sql` query files, followed by regeneration.

---

## 6. Migrations

Database schema is managed with [golang-migrate](https://github.com/golang-migrate/migrate).

- Migration files live in `db/migrations/` and follow the naming convention:
  ```
  000001_create_farms.up.sql
  000001_create_farms.down.sql
  ```
- Each migration has an `up` (apply) and `down` (rollback) file.
- Common Makefile targets:

  ```
  make migrate-up       # Apply all pending migrations
  make migrate-down     # Roll back the last migration
  make migrate-create name=<description>  # Scaffold a new migration pair
  ```

- Migrations are also applied automatically in the local docker-compose setup on startup.

---

## 7. Containerization

### Dockerfile (`docker/Dockerfile`)

Multi-stage build to keep the final image small:

1. **Build stage** (`golang:1.25-alpine`): compiles the binary with CGO disabled.
2. **Final stage** (`alpine`): copies only the compiled binary and any required static files.

### docker-compose.yml

Runs the full local development environment:

- `postgres`: PostgreSQL 17 with a health check.
- `app`: the API server, depends on postgres being healthy, environment loaded from `.env`.

```
docker-compose up       # Start all services
docker-compose down     # Stop and remove containers
```

### Environment variables

Copy `.env.example` to `.env` and fill in the values before running locally. The `.env` file is gitignored; `.env.example` is committed and kept up to date.

---

## 8. CI/CD (GitHub Actions + Coolify)

### Pipeline (`.github/workflows/ci.yml`)

Triggered on every push to `main`:

1. Run unit and integration tests.
2. Build the Docker image using the multi-stage Dockerfile.
3. Push the image to GitHub Container Registry (GHCR) tagged with the commit SHA and `latest`.

### Deployment (Coolify)

A self-hosted [Coolify](https://coolify.io) instance on the VPS is configured to watch GHCR for new image tags. On push, Coolify automatically pulls the latest image and redeploys the container. No manual SSH or deployment scripts needed.

---

## 9. Testing Strategy

Tests are organized into three levels:

### Unit tests
- `_test.go` files live alongside the source files they test.
- Service and handler logic is tested with mocked dependencies.
- Mocks are generated or written manually implementing the consumer-defined interfaces.
- Run with: `make test`

### Integration tests
- Use [testcontainers-go](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/) to spin up a real PostgreSQL container for the duration of the test.
- Repository layer is tested against an actual database with real migrations applied.
- Run with: `make test-integration`

### E2E / HTTP tests
- Use Go's `net/http/httptest` package with a fully wired Gin router.
- Cover the full request/response cycle including middleware.
- Run with: `make test-integration` (included in the integration suite)

---

## 10. Development Setup

### Prerequisites

- Go 1.25+
- Docker and docker-compose
- `sqlc` CLI: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- `golang-migrate` CLI: see [install docs](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- `make`

### Steps

```bash
# 1. Clone the repo
git clone <repo-url>
cd farming/backend

# 2. Set up environment variables
cp .env.example .env
# Edit .env with your local values

# 3. Start the database (and app) with docker-compose
docker-compose up -d

# 4. Apply migrations
make migrate-up

# 5. (Optional) Regenerate sqlc code after query changes
sqlc generate

# 6. Run the server locally (outside Docker)
make run

# 7. Run tests
make test                  # Unit tests only
make test-integration      # Unit + integration (requires Docker)
```

The API will be available at `http://localhost:8080` by default (configurable via `PORT` in `.env`).
