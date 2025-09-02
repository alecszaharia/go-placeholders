.PHONY: build run test clean deps dev

# Build the application
build:
	go build -o bin/antlr-test cmd/antlr/main.go

# Run the application
run:
	go run cmd/antlr/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Install development tools
install-tools:
	brew install antlr
	brew install golangci-lint
