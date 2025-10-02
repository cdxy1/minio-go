.DEFAULT_GOAL := build

BIN_DIR = ./bin
APP_NAME = service

fmt:
	go fmt ./... >/dev/null
.PHONY: fmt

lint: fmt
	go lint ./... >/dev/null
.PHONY: lint

vet: fmt
	go vet ./... >/dev/null
.PHONY: vet

build: vet
	@echo "Building service >>>>>>>>>"
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/app
.PHONY: build
