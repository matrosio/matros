.DEFAULT_GOAL := help

.SILENT : help
.PHONY  : build clean help install

company  := Matros
project  := Matros Workload Orchestrator
version  := 0.1.0

DEVEL = single

## build matros binary
build:
	@go build -o bin/matros

## clean up project
clean:
	@rm -rf ./bin

## start local development environment
devel:
	@vagrant up ${DEVEL}

## install dependencies
install:
	@go get

## run tests
tests:
	@go test -v ./...

## show help screen
help:
	printf "$(company) - $(project) v$(version)\n"
	printf "\nUsage:\n  make <target>\n\nTargets:\n"
	awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  \033[32m%-15s\033[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
