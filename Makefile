# Go parameters
BINARY_NAME=app
GO=go

# Build the application
build:
	$(GO) build -o bin/$(BINARY_NAME) cmd/main.go

# Run the application
run:
	swag init -g cmd/main.go --parseDependency --parseInternal
	$(GO) run cmd/main.go

# Clean build files
clean:
	rm -f bin/$(BINARY_NAME)
	rm -f *.log
	rm -f dump.rdb

# Run tests
test:
	$(GO) test -v ./... -cover

test-bench:
	$(GO) test -v -bench=. ./...	

# Run tests with coverage
test-coverage:
	$(GO) test -v -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# Download dependencies
deps:
	$(GO) mod download

# Tidy go.mod
tidy:
	$(GO) mod tidy

# Run linter
lint:
	golangci-lint run

# Build and run docker
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8000:8000 $(BINARY_NAME)

# Default target
.PHONY: build run clean test test-coverage deps tidy lint docker-build docker-run

default: build
