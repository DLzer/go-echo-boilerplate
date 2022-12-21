# Go Echo Boilerplate

## Overview

This project is meant to serve as a boilerplate for building a simple API in the Echo v4 framework. 

## Index

* [Development](#development)
* [Architectual Outline](#architectual-outline)
* [Migrations](#migrations)
* [API Documentation](#api-documentation)
* [Semantic Versioning](#semantic-versioning)

## Development and Deployment


### Development
Included in this project is a `Makefile` that can be used to perform various tasks. Previously we used custom bash scripts but multiple scripts can become unruly so the decision was made to consolidate it into a single Makefile.

For local deployments and to run the development environment it's best to use the following:

```makefile
make local // runs all containers
make run // runs the app, and allows easy attachment of debugger or re-running the app 
```

In the directory are 2 docker files that can be run directly via Docker CLI

```bash
docker-compose.local.yml - Will run Prometheus/Grafana/Jaeger
docker-compose.dev.yml - Will run docker development environment
```

## Architectual Outline

### Entrypoint
The entrypoint of the entire application is the `cmd/api/main.go` file which assembles the various utilities including the Config via `Viper`, the Database via `sqlx` and Jaeger distributed tracing via `Jaeger`. The end of the file runs the server package and starts the application.

### Server
The next tier down in the project is the server which is located at `/internal/server.go`. This file along with the associated `handler.go` starts various services like the Prometheus endpoint for metrics as well as all the app handlers and repositories to actually run the API.

### Implementation
The ideal flow of the application is:
```go
/orders
    /controller
        /http
            handlers.go <-- Defines the handler for accepting/returning http data
            routes.go <-- Defines the routes to handle
    /repository
        pq_repository.go <-- Defines the repository/data layer for Postgres
        mongo_repository.go <-- Defines the repository/data layer for Mongo
        sql_queries.go <-- Defines a safe place for shared service queries
    /service
        service.go <-- Defines the business logic/use case for the handler
    delivery.go <-- The interface for exposing handler functions
    service.go <-- The interface for exposing service/logic functions
    repository.go <-- The interface for exposing repository/storage functions
```

## Migrations

Database migrations are managed by a tool called [Goose](https://github.com/pressly/goose) by Pressly. Migrations can be found under the directory `/migrations`

### Simple usage of goose within a service

**Create a migration file:**
```bash
$ goose create update_a_column sql
2022110725125_update_a_column.sql generated
```

Within the file define the UP and DOWN migration statements.

```sql
-- +goose Up
-- +goose StatementBegin
ALTER TABLE .....
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE .....
-- +goose StatementEnd
```

**Run the migrations**
```bash
$ goose up
2022-11-07 OK 2022110725125_update_a_column
``` 

### Configuration notes

After installing goose it's necessary to set two environment variables for goose to know where and how to connect to the database in question. The variables to set are:

- `GOOSE_DBSTRING` Which is the DSN connection string for the DB
- `GOOSE_DRIVER` Which is a string that specifies the DB driver to use e.a `postgres`, `sql`, `sqlite`

## API Documentation

API Documentation is automatically generated via the Swagger. The godocs written on the route handlers are scanned and turned into OpenAPI documentation inside the `/docs` directory.

To run the API generate use the makefile command:
```bash
make swaggo
```

This should scan and generate accompanying documentation.

## Semantic Versioning

The versioning and release system used by the application is as follows:
```
v.{currentYear:YY}.{currentMonth:MM}.{currentDay:DD}.{latestCommitHash:-7}-{alpha-beta}
```