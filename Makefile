.PHONY: help test build lint fmt clean all install-tools

help: ## Display this help message
	@echo "Learning Go The Hard Way - Available Commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

install-tools: ## Install development tools
	@echo "Installing Go tools..."
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed successfully!"

fmt: ## Format all Go code
	@echo "Formatting code..."
	go fmt ./...
	gofmt -s -w .

lint: ## Run linters
	@echo "Running linters..."
	golangci-lint run ./...

test: ## Run all tests
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

test-module: ## Run tests for specific module (usage: make test-module MODULE=01-basics)
	@echo "Running tests for $(MODULE)..."
	go test -v -race ./modules/$(MODULE)/...

coverage: test ## Generate coverage report
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

build: ## Build all executable examples
	@echo "Building examples..."
	@for dir in $$(find ./modules -name "cmd" -type d); do \
		(cd $$dir && go build -o ../../../bin/$$(basename $$(dirname $$dir)) .); \
	done

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/ coverage.out coverage.html

verify: fmt lint test ## Run all verification steps

all: install-tools verify build ## Run complete build pipeline
