-include .env

# ============================================================================== #
# HELPERS
# ============================================================================== #

## help: prints this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo 'Are you sure? [y/N] ' && read ans  && [ $${ans:-N} = y ] 

# ============================================================================== #
# DEVELOPMENT
# ============================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run ./cmd/api -db-dsn=${POSTGRES_DSN}

## db/migrations/new name=$1: creates a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration file for ${name}...'
	@goose create -s ${name} sql

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations'
	@goose up

# ============================================================================== #
# QUALITT CONTROL
# ============================================================================== #

## tidy: format all .go files and tidy module dependencies
.PHONY: tidy
tidy:
	@echo 'Formatting go files'
	@go fmt ./...
	@echo 'Tidying module dependencies'
	@go mod tidy
	@echo 'Verifyig module dependencies...'
	@go mod verify

## audit: run quality control checks
.PHONY: audit
audit:
	@echo 'Checking module dependencies'
	@go mod tidy -diff
	@go mod verify
	@echo 'Vetting code'
	@go vet ./...
	@echo 'Running tests'
	@go test -race -vet=off ./...