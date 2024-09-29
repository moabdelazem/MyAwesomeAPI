include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

run:
	@air	

build:
	@go build -o bin/main cmd/main.go

test:
	@go test -v ./...

.PHONY: migrate-up
migration:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDRESS) up

.PHONY: migrate-down
migration-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDRESS) down $(filter-out $@,$(MAKECMDGOALS))