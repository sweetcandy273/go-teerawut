hello:
	@echo "Hello, World"

tidy: ## Vendor and Tidy go modules
	@go mod tidy

run: ## Run application on default port 8000 and print console format
	@go run main.go

lint: ## Run golangci-lint
	@golangci-lint run ./...

# === CONFIGURATION ===
MIGRATE_DIR=./db/migrations

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir $(MIGRATE_DIR) $$name

migrate-up:
	@migrate -path $(MIGRATE_DIR) -database "$$DATABASE_URL" up

migrate-down:
	@migrate -path $(MIGRATE_DIR) -database "$$DATABASE_URL" down 1

migrate-status:
	migrate -path $(MIGRATE_DIR) -database "$$DATABASE_URL" version
