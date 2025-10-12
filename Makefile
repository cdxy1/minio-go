.DEFAULT_GOAL := build

BIN_DIR = ./bin
APP_NAME = service

up:
	docker compose up -d
.PHONY: up

down:
	docker compose down
.PHONY: down

fmt:
	go fmt ./... >/dev/null
.PHONY: fmt

vet: fmt
	go vet ./... >/dev/null
.PHONY: vet

proto: vet
	protoc --go_out=./internal/ --go-grpc_out=./internal/ api/proto/*.proto
.PHONY: proto

build: proto
	@echo "Building service >>>>>>>>>"
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/app
.PHONY: build
