# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
PROJECTNAME:=$(shell basename "$(PWD)")
COVERAGE=coverage.out

## vendor:update dependencies
.PHONY: vendor
vendor:
	go mod tidy && go mod vendor

## test:run unit tests in verbose mode
.PHONY: test
test:
	$(GOTEST) -v ./...

## clean:clean cached tests
.PHONY: test-clean
test-clean:
	$(GOCLEAN) -testcache
	
.PHONY: lint
lint:
	golangci-lint run -v --timeout=1m

## coverage:generate test coverage html
coverage:
	go test -coverprofile=$(COVERAGE) ./... && \
	go tool cover -html=$(COVERAGE)

## help:print help message
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo	