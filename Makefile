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

proto: vet
	mkdir -p ./internal/grpc
	protoc --go_out=./internal/grpc --go-grpc_out=./internal/grpc api/proto/*.proto
.PHONY: proto

build: vet
	@echo "Building service >>>>>>>>>"
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/app
.PHONY: build
