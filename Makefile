# Get the latest git tag
TAG=$(shell git describe --tags --abbrev=0)

.PHONY: build
build: # Build the project
	@go build -o bin/ ./...

.PHONY: test
test: # Test the project
	@go test -v ./...

.PHONY: release
release: # Release the project
	@go mod vendor
	@RELEASE_VERSION=${TAG} ./hack/build.sh

.PHONY: format
format: # Format all go code in project
	@gofmt -s -w ./

## Dev Helpers ##

.PHONY: dev/start-nomad
dev/start-nomad: # Start a local Nomad server for testing
	@nomad agent -dev -bind=0.0.0.0 > /tmp/nomad.log 2>&1 &
	@echo "Nomad started on http://localhost:4646"
	@echo "Nomad logs are in /tmp/nomad.log"

.PHONY: dev/stop-nomad
dev/stop-nomad: # Stop a local Nomad server
	@pkill nomad
	@echo "Nomad stopped"

## END Dev Helpers END ##

.PHONY: help
help: # Print valid Make targets
	@echo "Valid targets:"
	@grep --extended-regexp --no-filename '^[a-zA-Z/_-]+:' Makefile | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
