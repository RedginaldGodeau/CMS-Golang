
.PHONY: up
up:
	@docker compose up -d

build: ./internal/app/main.go
	@docker compose exec command go build -o bin/build/app ./internal/app

.PHONY: run
run:
	@docker compose exec command go run ./internal/app

.PHONY: start_migration
start_migration:
	@docker compose exec command go run ./internal/cmd/start_migration

.PHONY: generate_account
generate_account:
	@docker compose exec command go run ./internal/cmd/new_api_account $(ARGS)

.PHONY: ci
ci:
	@docker compose exec command golangci-lint run