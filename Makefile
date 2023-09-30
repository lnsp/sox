GOCMD=go
YARNCMD=yarnpkg
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
YARNBUILD=$(YARNCMD) generate

PREFIX=/usr/local/bin

BIN_FOLDER=bin

SERVER_BINARY_NAME=sox
CLIENT_BINARY_NAME=sox-cli
UI_BINARY_NAME=sox-ui

VERSION := $(shell date -u +"%Y.%m.%d-%s")
export VERSION

all: build-ui build

proto:
	protoc -I api/ --go_out=api/ --go_opt=paths=source_relative --go-grpc_out=api/ --go-grpc_opt=paths=source_relative data.proto service.proto

$(BIN_FOLDER):
	mkdir -p $(BIN_FOLDER)

build: $(BIN_FOLDER)
	$(GOBUILD) -ldflags '-X github.com/lnsp/sox/meta.Version=$(VERSION)' -o $(BIN_FOLDER)/$(SERVER_BINARY_NAME) -v ./cmd/server
	$(GOBUILD) -ldflags '-X github.com/lnsp/sox/meta.Version=$(VERSION)' -o $(BIN_FOLDER)/$(CLIENT_BINARY_NAME) -v ./cmd/client
	$(GOBUILD) -ldflags '-X github.com/lnsp/sox/meta.Version=$(VERSION)' -o $(BIN_FOLDER)/$(UI_BINARY_NAME) -v ./cmd/ui

install:
	unlink $(PREFIX)/$(SERVER_BINARY_NAME); cp $(BIN_FOLDER)/$(SERVER_BINARY_NAME) $(PREFIX)/$(SERVER_BINARY_NAME)
	unlink $(PREFIX)/$(CLIENT_BINARY_NAME); cp $(BIN_FOLDER)/$(CLIENT_BINARY_NAME) $(PREFIX)/$(CLIENT_BINARY_NAME)
	unlink $(PREFIX)/$(UI_BINARY_NAME); cp $(BIN_FOLDER)/$(UI_BINARY_NAME) $(PREFIX)/$(UI_BINARY_NAME)

build-ui:
	cd ui && $(YARNBUILD)

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BIN_FOLDER)
