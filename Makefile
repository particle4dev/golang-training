# Set an output prefix, which is the local directory if not specified
PREFIX?=$(shell pwd)

# Used to populate version variable in main package.
VERSION=$(shell git describe --match 'v[0-9]*' --dirty='.m' --always)

GO_LDFLAGS=-ldflags "-X `go list ./version`.Version=$(VERSION)"

#Source file target
SRCS  := $(shell find . -type f -name '*.go')

PKGS_AND_MOCKS := $(shell go list ./... | grep -v /vendor)
PKGS := $(shell echo $(PKGS_AND_MOCKS) | tr ' ' '\n' | grep -v /mock$)
PKGS_TEST := $(shell echo $(PKGS_AND_MOCKS) | tr ' ' '\n' | grep pkg$)

.PHONY: clean all fmt vet lint build test vendor-update containers check-docs e2e-test get-tools
.DEFAULT: all
all: clean fmt vet lint build test binaries

ci: fmt vet lint build test

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s', or use 'make fmt-save'" && false)

fmt-save:
	@echo "+ $@"
	@gofmt -s -l . 2>&1 | grep -v ^vendor/ | xargs gofmt -s -l -w

lint:
	@echo "+ $@"
	$(if $(shell which golint || echo ''), , \
		$(error Please install golint: `make get-tools`))
	@test -z "$$(golint ./... 2>&1 | grep -v ^vendor/ | grep -v mock/ | tee /dev/stderr)"

coverage:
	@echo "+ $@"
	@for pkg in $(PKGS_TEST); do \
		go test -test.short -race -coverprofile="./$$pkg/coverage.txt" $${pkg} || exit 1; \
	done

get-tools:
	@echo "+ $@"
	@go get -u \
		github.com/golang/lint/golint \
		github.com/wfarner/blockcheck \
		github.com/rancher/trash \
		github.com/stretchr/testify \
		github.com/ethereum/go-ethereum/log

logs:
	@echo "+ $@"
	@echo "PREFIX" ${PREFIX}
	@echo "VERSION" ${VERSION}
	@echo "GO_LDFLAGS" ${GO_LDFLAGS}
	@echo "SRCS" ${SRCS}
	@echo "PKGS_AND_MOCKS" ${PKGS_AND_MOCKS}
	@echo "PKGS" ${PKGS}
	@echo "PKGS_TEST" ${PKGS_TEST}

install:
	@echo "+ $@"
	@go install ${GO_LDFLAGS} $(PKGS)

test:
	@echo "+ $@"
	@go test -test.short -timeout 30s -race -v $(PKGS_TEST)

build:
	@echo "+ $@"
	@go build ${GO_LDFLAGS} $(PKGS)

vet:
	@echo "+ $@"
	@go vet $(PKGS)

vendor-update:
	@echo "+ $@"
	@trash -u