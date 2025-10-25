.DEFAULT_GOAL := build

BIN_DIR = ./bin
FILE_SERVICE = file
METADATA_SERVICE = metadata
GATEWAY_SERVICE = gateway

up:
	docker compose up -d && docker exec kafka sh /kafka-init.sh
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
	go build -o $(BIN_DIR)/$(FILE_SERVICE) ./cmd/file_service/main.go
	go build -o $(BIN_DIR)/$(METADATA_SERVICE) ./cmd/metadata_service/main.go
	go build -o $(BIN_DIR)/$(GATEWAY_SERVICE) ./cmd/gateway/main.go
.PHONY: build
