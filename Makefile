.PHONY: all build

version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR := $(ROOT_DIR)/bin
version := $(or $(version), $(shell cat /app/build-release | tr -d '\n'))

all: build

clean:
	rm -f $(BIN_DIR)/macross

build: codegen lint
	rm -f $(BIN_DIR)/kodiak
	go build -o $(BIN_DIR)/kodiak -v -ldflags \
		"-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/test/kodiak/main.go

lint:
	golangci-lint run

test: lint
	go test ./...

codegen:
	mkdir -p internal/sc
	abigen --abi assets/abis/kodiakvaultv1.abi --pkg sc --type KodiakV1 --out internal/sc/kodiak_v1.go
