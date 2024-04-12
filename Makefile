TAG=0.0.1

.PHONY: build
build: # Build the project
	@go build -o bin/ ./...

.PHONY: test
test: # Test the project
	@go test -v ./...

.PHONY: format
format: # Format all go code in project
	@gofmt -s -w ./

.PHONY: help
help: # Print valid Make targets
	@echo "Valid targets:"
	@grep --extended-regexp --no-filename '^[a-zA-Z/_-]+:' Makefile | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
