GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BIN_FOLDER=bin

SERVER_BINARY_NAME=$(BIN_FOLDER)/virtm
CLIENT_BINARY_NAME=$(BIN_FOLDER)/virtm-cli
UI_BINARY_NAME=$(BIN_FOLDER)/virtm-ui

VERSION := $(shell date -u +"%Y.%m.%d-%s")
export VERSION

all: build

proto:
	protoc -I api/ --go_out=api/ --go_opt=paths=source_relative --go-grpc_out=api/ --go-grpc_opt=paths=source_relative data.proto service.proto

$(BIN_FOLDER):
	mkdir -p $(BIN_FOLDER)

build: bin
	$(GOBUILD) -ldflags '-X github.com/valar/virtm/meta.Version=$(VERSION)' -o $(SERVER_BINARY_NAME) -v ./cmd/server
	$(GOBUILD) -ldflags '-X github.com/valar/virtm/meta.Version=$(VERSION)' -o $(CLIENT_BINARY_NAME) -v ./cmd/client
	$(GOBUILD) -ldflags '-X github.com/valar/virtm/meta.Version=$(VERSION)' -o $(UI_BINARY_NAME) -v ./cmd/ui

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BIN_FOLDER)
