#!make
include .env.example

BIN_APP="./bin/anti-bruteforce"

config:
	@if [ ! -f ".env" ] ; then echo "creating .env" ; cp .env.example .env; fi

build: config
	go build -v -o $(BIN_APP) ./cmd/anti-bruteforce/

run: build
	$(BIN_APP)

test:
	go test -cover ./...

integra-test:
	go test -tags integration ./tests/integration/...

generate:
	 go generate ./...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps
	golangci-lint run

up: config
	docker compose -f docker-compose.yml up -d --build

down:
	docker compose -f docker-compose.yml down

.PHONY: build run test lint up down config
