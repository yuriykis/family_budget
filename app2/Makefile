.PHONY: build

BINARY_NAME=app

build:
	echo "Copying server.yml configuration file..."
	cp -f ./configs/server.yml ./bin/server.yml
	
	echo "Building..."
	go build -o bin/$(BINARY_NAME) -v ./cmd/app

.DEFAULT_GOAL := build