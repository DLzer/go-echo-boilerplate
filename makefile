# .PHONY: migrate migrate_down migrate_up migrate_version docker prod docker_delve local swaggo test

# ==============================================================================
# Go goose postgresql

status:
	goose status

version:
	goose version

migrate_up:
	goose up

migrate_down:
	goose down


# ==============================================================================
# Docker compose commands

prod:
	echo "Starting docker prod environment"
	docker-compose -f docker-compose.prod.yml up --build

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build

dockerbuild:
	echo "Starting Docker Build"
	docker build . -t github.com/DLzer/go-echo-boilerplate && docker run -p 5000:5000 github.com/DLzer/go-echo-boilerplate:latest

# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go


# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...


# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)

# =============================================================================
# Profiling ( Requires pprof debug/middleware enabled and secondary debug server)

profile-web:
	go tool pprof -http=":8000" pprofbin http://localhost:5555

# ==============================================================================
# Versioning

release:
	goreleaser release --clean
