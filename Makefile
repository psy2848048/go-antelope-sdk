.PHONY: test-all
test-all: test cover upload-coverage

.PHONY: get-coverage
get-coverage: cover upload-coverage

# Run all unit tests
.PHONY: test
test:
	go test -short ./...

# Same as test but with coverage turned on
.PHONY: cover
cover:
	go test -short -cover -covermode=atomic -coverprofile coverage.txt ./...

# Apply https://github.com/golangci/golangci-lint to changes since forked from master branch
.PHONY: lint
lint:
	golangci-lint run --timeout=5m --enable=unparam --enable=misspell --enable=prealloc

# Remove all compiled binaries from the directory
.PHONY: clean
clean:
	go clean

# Analyze the code for any unused dependencies
.PHONY: prune-deps
prune-deps:
	go mod tidy
