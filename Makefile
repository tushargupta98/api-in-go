#------ This is to faciliate local setup ------#

### Environment Variables
# Environment Identifier
export RDM_ENVIRONMENT_TYPE ?= local

# DB Environment Varliables
export PG_CERTIFICATE ?= 
export PG_DATABASE ?= api_in_go
export PG_PASSWORD ?= api_in_go_password
export PG_SCHEMA ?= api_in_go
export PG_URL ?= localhost
export PG_PORT ?= 5434
export PG_USERNAME ?= api_in_go

# Redis Environment Valriables
export REDIS_DATABASE ?= 7
export REDIS_URL ?= localhost

# Go variables
GO_FILES := $(wildcard *.go)
GO_TEST_FILES := $(wildcard ./test/unit/*.go) 

# Local variables
LOCAL_WAIT_DB_TIMEOUT := 5
LOCAL_WAIT_APP_TIMEOUT := 5
DB_CONNECTION_STR := "user=$(PG_USERNAME) password=$(PG_PASSWORD) dbname=$(PG_DATABASE) host=$(PG_URL) port=$(PG_PORT) sslmode=disable search_path=$(PG_SCHEMA)"


.PHONY: help
help:
	@echo "Makefile Usage:"
	@echo "  make local      - Run local development environment"
	@echo "  make goose_up   - Migrate the database schema up"
	@echo "  make goose_down - Migrate the database schema down"
	@echo "  make goose_create - Create a new migration. Usage: make goose_create name=<migration_name>"
	@echo "  make unit_tests_run - Run unit tests"
	@echo "  make int_tests_run  - Run integration tests"
	@echo "  make clean      - Clean project"
	@echo "  make help       - Display this help message"

.PHONY: local
local: docker_up local_setup goose_up app_start

## Docker Commands
.PHONY: docker_up
docker_up:
	@docker-compose -f docker-compose.local.yml up -d
	@sleep $(LOCAL_WAIT_DB_TIMEOUT)

.PHONY: docker_down
docker_down:
	@docker-compose -f docker-compose.local.yml down

## Local Setup
.PHONY: local_setup
local_setup:
	chmod +x cmd/bash/*
	@./cmd/bash/local/local-setup.sh
	@sleep $(LOCAL_WAIT_DB_TIMEOUT)

.PHONY: local_setup_down
local_setup_down:
	@make docker_down

## Goose Commands
.PHONY: goose_up
goose_up:
	@goose -dir ./db/migrations postgres $(DB_CONNECTION_STR) up

.PHONY: goose_down
goose_down:
	@goose -dir ./db/migrations postgres $(DB_CONNECTION_STR) down

.PHONY: goose_create
goose_create:
    ifeq ($(name),)
	    @echo "Please provide the migration name. Usage: make goose_create name=<migration_name>"
	    @exit 2
    endif
	goose -dir ./db/migrations create $(name) sql

## Go APP commands
.PHONY: app_start 
app_start: app_build app_run

.PHONY: app_build
# Build the api-in-go api
app_build:
	go clean
	go build ./...
	go mod tidy

.PHONY: app_run
# Run the api-in-go api in developer mode
app_run:
	air -c ./air.toml
#go run -race -gcflags "all=-N -l" ./cmd/api/main.go

.PHONY: unit_tests_run
unit_tests_run:
	go test -v -short ./test/unit/...

.PHONY: int_tests_run
int_tests_run:
	go test -v -short ./test/integration/...

# Swagger 
.PHONY: swagger_docs
swagger_docs:
	swag init --generalInfo ./cmd/api/main.go --output ./pkg/swagger/docs

.PHONY: clean
clean:
	rm -rf $(GO_FILES) $(GO_TEST_FILES) $(BINARY)

# Code Analysis
.PHONY: golangci-lint
golangci-lint:
	golangci-lint run ./...