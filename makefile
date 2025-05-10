# Packages and tools
PKG=./...

LINT=golangci-lint run
GEN=go generate ./...

.PHONY: all test lint generate start

# Run all essential checks
all: test lint

# Run unit tests with coverage
test:
	go test -v -cover $(PKG)

# Lint code using golangci-lint
lint:
	@command -v golangci-lint > /dev/null || (echo "golangci-lint is not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	$(LINT)

# Generate all necessary code (e.g., mocks, wire)
generate:
	$(GEN)

# Start application stack via Docker Compose
start:
	docker-compose up --build --force-recreate
