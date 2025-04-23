## Example Makefile that apps would use

# Application configuration
APP_NAME := checker

# Go configuration
GOPATH := $(shell go env GOPATH)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GO_FILES := $(shell find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./.git/*")
GO_PACKAGES := $(shell go list ./... | grep -v /vendor/)

# Directories
BIN_DIR := ./bin
REPORTS_DIR := ./reports

# Tool binaries
GOLANGCI_LINT := $(GOPATH)/bin/golangci-lint
GOTESTSUM := $(GOPATH)/bin/gotestsum

# Main targets
.PHONY: build clean lint setup unittest

# Default target runs setup, lint, test and build
all: setup lint unittest build

# Setup development environment and fetch dependencies
setup:
	@echo "Setting up development environment..."
	go mod download
	
	@echo "Installing development tools..."
	echo $(GOLANGCI_LINT)
	@if [ ! -f $(GOLANGCI_LINT) ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@if [ ! -f $(GOTESTSUM) ]; then \
		go install gotest.tools/gotestsum@latest; \
	fi
	
	@echo "Setup complete"

# Build the application
build:
	@echo "Building application..."
	mkdir -p $(BIN_DIR)
	go build -v -ldflags="-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)" -o $(BIN_DIR)/$(APP_NAME) main.go
	@echo "Build complete: $(BIN_DIR)/$(APP_NAME)"

# Run linters
lint:
	@echo "Running linters..."
	$(GOLANGCI_LINT) run --timeout 5m
	@echo "Linting complete"

# Run unit tests
unittest:
	@echo "Running unit tests..."
	mkdir -p $(REPORTS_DIR)
	$(GOTESTSUM) --junitfile $(REPORTS_DIR)/unit-tests.xml -- -short -race -coverprofile=$(REPORTS_DIR)/coverage.out ./...
	go tool cover -html=$(REPORTS_DIR)/coverage.out -o $(REPORTS_DIR)/coverage.html
	@echo "Unit tests complete"

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)
	rm -rf $(DIST_DIR)
	rm -rf $(REPORTS_DIR)
	go clean
	@echo "Clean complete"

help:
	@echo "Available targets:"
	@echo "  all              : Run setup, lint, build, and unittest (default)"
	@echo "  setup            : Set up development environment and fetch dependencies"
	@echo "  build            : Build the application"
	@echo "  lint             : Run linters"
	@echo "  unittest         : Run unit tests"
	@echo "  clean            : Clean build artifacts"
	@echo "  help             : Show this help message"