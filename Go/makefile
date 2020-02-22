# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

# update dependencies
.PHONY: vendor
vendor:
	go mod tidy && go mod vendor

# run unit tests in verbose mode
.PHONY: test
test:
	$(GOTEST) -v ./...

# clean cached tests
.PHONY: test-clean
test-clean:
	$(GOCLEAN) -testcache
