-include .env

run:
	@go run ./cmd/api --db-dsn=$(POSTGRES_DSN)