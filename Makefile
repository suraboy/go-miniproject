.PHONY: build run clean test fmt vet lint

# Build the application
build:
	go build -o bin/app app/main.go

# Run the application
run:
	go run app/main.go

# Clean build artifacts
clean:
	rm -rf bin/

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Run golangci-lint
lint:
	golangci-lint run

# Install dependencies
tidy:
	go mod tidy
	go mod download

# Run all checks
check: fmt vet lint test

# Default target
all: tidy fmt vet lint test build