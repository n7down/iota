VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
MAKEFLAGS += --silent
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(GOPATH)/src/github.com/n7down/iota/cmd/kuiper/*.go

.PHONY: get 
get:
	echo "getting go dependencies..."
	@go get ./...
	echo "done"

.PHONY: generate
generate:
	echo "generating dependency files..."
	protoc --go_out=plugins=grpc:internal/pb/settings internal/pb/settings/settings.proto
	echo "done"

.PHONY: test-unit
test-unit:
	echo "running unit tests..."
	go test -tags=unit -v ./...
	echo "done"

.PHONY: test-integrations
test-integrations:
	echo "running integrations test"
	go test -tags=integrations -v ./...
	echo "done"

.PHONY: test
test: test-unit

.PHONY: lint
lint:
	golint ./...

.PHONY: stop
stop:
	echo "stopping docker containers..."
	docker-compose stop
	echo "done"

.PHONY: rm
rm:
	echo "removing docker containers..."
	docker-compose rm
	echo "done"

.PHONY: clean
clean: stop rm

.PHONY: build-apigeteway
build-apigateway:
	echo "building apigateway..."
	docker build -t "$(PROJECTNAME)"/apigateway:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/apigateway/Dockerfile .
	echo "done"

.PHONY: build-sensors
build-sensors:
	echo "building sensors..."
	docker build -t "$(PROJECTNAME)"/sensors:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/apigateway/Dockerfile .
	echo "done"

.PHONY: build-settings
build-settings:
	echo "building settings..."
	docker build -t "$(PROJECTNAME)"/settings:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/apigateway/Dockerfile .
	echo "done"

.PHONY: build-all
buildall: build-apigateway build-sensors build-settings


.PHONY: help
help:
	echo "Choose a command run in $(PROJECTNAME):"
