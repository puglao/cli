BINARY_NAME=$(shell basename "$(PWD)")
GIT_TAG ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")
GIT_SHA ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

VERSION_PKG = github.com/puglao/cli/internal/version

LDFLAG += -X '$(VERSION_PKG)/version.tag=$(GIT_TAG)'
LDFLAG += -X '$(VERSION_PKG)/version.sha=$(GIT_SHA)'

.PHONY: pre-commit
pre-commit:
	@echo "Running pre-commit hooks"
	@pre-commit run --all-files

build:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAG)" -o bin/$(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: integration-test
integration-test:
	go test -v -tags=integration ./test/...
