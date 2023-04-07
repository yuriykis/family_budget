# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTIDY=$(GOCMD) mod tidy
BINARY_NAME=user_service
BINARY_UNIX=$(BINARY_NAME)_unix

# Directories
CMD_DIR=./cmd
INTERNAL_DIR=./internal
PKG_DIR=./pkg
API_DIR=./api
WEB_DIR=./web
CONFIG_DIR=./config
TEST_DIR=./test
BIN_DIR=./bin

# Docker parameters
DOCKER_IMAGE_NAME=user_service
DOCKER_IMAGE_TAG=latest

# Targets
.PHONY: all build clean test run deps build-linux docker-build docker-clean

all: test build

build: go-tidy
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v $(CMD_DIR)/app

clean:
	$(GOCLEAN)
	rm -f $(BIN_DIR)/$(BINARY_NAME)
	rm -f $(BIN_DIR)/$(BINARY_UNIX)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) -v $(CMD_DIR)/app
	./$(BIN_DIR)/$(BINARY_NAME)

deps:
	$(GOGET) ./...

go-tidy:
	$(GOTIDY)
	 
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_DIR)/$(BINARY_UNIX) -v $(CMD_DIR)/app

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) .

docker-clean:
	docker rmi $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
