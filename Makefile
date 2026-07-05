SHELL := /bin/sh

BINARY := bin/fotoforge
VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || printf unknown)
BUILD_DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w \
	-X main.buildVersion=$(VERSION) \
	-X main.buildCommit=$(COMMIT) \
	-X main.buildDate=$(BUILD_DATE)
GO_FILES := $(shell find . -type f -name '*.go' -not -path './vendor/*')

.PHONY: fmt vet test build smoke check clean

fmt:
	gofmt -w $(GO_FILES)

vet:
	go vet ./...

test:
	go test ./...

build:
	@mkdir -p bin
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd/fotoforge

smoke: build
	./$(BINARY) version
	./$(BINARY) --help >/dev/null

check:
	@test -z "$$(gofmt -l $(GO_FILES))" || { \
		printf '%s\n' 'Go files need formatting; run make fmt'; \
		gofmt -l $(GO_FILES); \
		exit 1; \
	}
	go vet ./...
	go test ./...
	$(MAKE) smoke

clean:
	rm -rf bin dist coverage.out
