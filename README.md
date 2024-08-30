# Go API Project using Echo v4

Boilerplate for spinning up a Go API quickly using the Echo framework.
Features:
- Ready to go with local docker development and auto-migrations
- Prepped for Authentication ( JWT or Cookies )
- Robust logging
- Prometheus Metrics endpoint + Pprof server
- Example CICD and Deployment using
    - Docker Compose
    - Watchtower
    - Traefik

[![Go Report Card](https://goreportcard.com/badge/github.com/DLzer/go-echo-boilerplate)](https://goreportcard.com/report/github.com/DLzer/go-echo-boilerplate) ![GitHub Release](https://img.shields.io/github/v/release/DLzer/go-echo-boilerplate)

## Prerequesites
- Go >= v1.22.4
- Docker

## Index

* [Getting Started](#getting-started)
* [Extending](#extending)
* [Migrations](#migrations)
* [API Documentation](#api-documentation)
* [Semantic Versioning](#semantic-versioning)


## Getting Started
Clone the project and run it locally with Docker
```bash
$ git clone https://github.com/DLzer/go-echo-boilerplate.git
$ cd go-echo-boilerplate
$ make local
```

Should resolve with
```bash
 ✔ Container go-echo-boilerplate-db-1  Healthy
 ✔ Container echo_api                  Started
 ✔ Container pgadmin4_container        Running 
```

Test that all services are up and running by vistiting the following in your browser

Visit `localhost:8008/v1/health`
```json
{"status": "ok"}
```

Visit `localhost:8008/v1/users?page=1&size=10`
```json
{
  "total_count": 1,
  "total_pages": 1,
  "page": 1,
  "size": 10,
  "has_more": false,
  "values": [
    {
      "uuid": "9edac806-3fe8-49ef-9522-e137d9a5bb4b",
      "email": "testuser@goboilerplate.io",
      "first_name": "John",
      "last_name": "Smith",
      "created_at": "2024-08-30T13:16:49.190144Z",
      "updated_at": "2024-08-30T13:16:49.190144Z",
      "roles": [
        "USER",
        "ADMIN"
      ]
    }
  ]
}
```


## Extending

A brief overview of the project structure
```yaml
project
│cmd/api/main.go # The entrypoint
│
└───internal # Where all API login lives
│   │
│   └───server
│   |   │   handler.go # Handler for all Middleware and API handlers
│   |   │   server.go # Inception point for all API servers
│   │
│   └───users # A domain directory
│   │   └───http # Routes and Route Handlers
│   │   └───repository # Queries and Repository layers
│   │   └───service # Service layer for handler->repository transporting and logic
│   |   │   service.go # Service interface
│   |   │   repository.go # Repository interface
│   |   │   handler.go # Handler interface
│
└───pkg # Public packages
    └───utils # Utilities for configs, JWT claims, image parsing, etc..
    └───logging # Uber zap logger library
    └───middleware # API shared middleware
```

To extend the API by adding another domain following the created pattern you will create a domain folder under `internal/YourDomain`. Lets say you want to create a *Businesses* domain. Create the businesses directory, and inside create the shared interface files (Service, Repository, Handler). Following that create the directories that hold the actual logic. 

The pattern above does not need to be followed exactly. For something like a simple static route you could get away with just a Handler that returns x.txt or something. However, the handler does need to be implemented in `/internal/server/handler.go` to define a route and the handling function.

## Migrations

Database migrations are managed by a tool called [Goose](https://github.com/pressly/goose) by Pressly. Migrations can be found under the directory `cmd/api/migrations`.

**Create a migration file**
```bash
$ goose create create_business_table sql
2022110725125_create_business_table.sql generated
```

**Write the migration**

```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE.....
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE .....
-- +goose StatementEnd
```

**Running migrations**

Migrations are automatically applied on a rebuild/restart of the API. Migraiton files are embedded into the `main.go` file as well as migration scripts. 

## API Documentation

API Documentation can be automatically generated via [goswag](https://github.com/swaggo/swag). The godocs written on the handlers functions and models are scanned and turned into OpenAPI documentation inside the `/docs` directory.

To run the API generate use the makefile command:
```bash
make swaggo
```

*Note:* At the moment this only supports **OpenAPIv2** or **Swagger**. Still looking for a good solution for auto generation of OpenAPIv3.

## Semantic Versioning

The versioning and release system used by the application follows [Semver2.0](https://semver.org/):
```
v.{Major}.{Minor}.{Patch}-{alpha-beta-release}
```

## TODO
- [ ] Writing Tests + Coverage
- [ ] Toggleable Auth+Middleware
- [ ] Microservice example with included OpenTelemetry