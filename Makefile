.PHONY: build clean test run install

BINARY_NAME=kawa
BUILD_DIR=bin
GO_FILES=$(shell find . -name '*.go' -type f)

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@go test -v ./...

run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

install: build
	@echo "Installing $(BINARY_NAME)..."
	@install -d ~/.local/bin
	@install -m 755 $(BUILD_DIR)/$(BINARY_NAME) ~/.local/bin/
	@echo "Installed to ~/.local/bin/$(BINARY_NAME)"
	@echo "Make sure ~/.local/bin is in your PATH"

dev:
	@echo "Running in development mode..."
	@go run main.go

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Linting code..."
	@golangci-lint run

install-global: build
	@echo "Installing $(BINARY_NAME) globally..."
	@sudo install -m 755 $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installed to /usr/local/bin/$(BINARY_NAME)"

help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  clean         - Remove build artifacts"
	@echo "  test          - Run tests"
	@echo "  run           - Build and run the binary"
	@echo "  install       - Install to ~/.local/bin (user install)"
	@echo "  install-global - Install to /usr/local/bin (requires sudo)"
	@echo "  dev           - Run without building"
	@echo "  fmt           - Format Go code"
	@echo "  lint          - Run linter"
	@echo "  help          - Show this help message"