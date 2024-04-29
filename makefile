hello:
	@echo "Hello, World"

tidy: ## Vendor and Tidy go modules
	@go mod tidy

run: ## Run application on default port 8000 and print console format
	@go run main.go

lint: ## Run golangci-lint
	@golangci-lint run ./...