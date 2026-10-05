APP_NAME := agrocore
BUILD_DIR := bin
BINARY := $(BUILD_DIR)/$(APP_NAME)

.PHONY: help fmt fmt-check vet test test-race build run clean check

help:
	@echo "AgroCore development commands:"
	@echo ""
	@echo "  make help        Show available commands"
	@echo "  make fmt         Format Go source code"
	@echo "  make fmt-check   Check Go source code formatting"
	@echo "  make vet         Run Go static analysis"
	@echo "  make test        Run automated tests"
	@echo "  make test-race   Run tests with the race detector"
	@echo "  make build       Build the application"
	@echo "  make run         Run the application"
	@echo "  make clean       Remove generated build artifacts"
	@echo "  make check       Run all quality checks"

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || \
		(echo "The following files are not formatted:"; \
		gofmt -l .; \
		exit 1)

vet:
	go vet ./...

test:
	go test -count=1 ./...

test-race:
	go test -race -count=1 ./...

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BINARY) .

run:
	go run .

clean:
	rm -rf $(BUILD_DIR)

check: fmt-check vet test test-race build