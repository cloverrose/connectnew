export

# Setup Go Variables
GOPATH := $(shell go env GOPATH)
GOBIN := $(PWD)/bin

# Invoke shell with new path to enable access to bin
PATH := $(GOBIN):$(PATH)
PATH := $(shell aqua root-dir)/bin:$(PATH)
SHELL := env "PATH=$(PATH)" bash

# update updates dependencies.
.PHONY: update
update:
	make go/update
	make aqua/update
	make go/update/testdata

# aqua/install installs aqua dependencies.
.PHONY: aqua/install
aqua/install:
	aqua install

# aqua/update-checksum updates the checksums for the aqua dependencies.
.PHONY: aqua/update-checksum
aqua/update-checksum:
	aqua update-checksum --deep --prune

# aqua/update updates aqua dependencies.
.PHONY: aqua/update
aqua/update:
	aqua update-aqua
	aqua update
	aqua update-checksum --deep --prune
	aqua install

# aqua/reset removes all the aqua dependencies.
.PHONY: aqua/reset
aqua/reset:
	aqua rm -all

# lint runs the linters.
.PHONY: lint
lint:
	@golangci-lint version
	@golangci-lint run $(args) --config=./.golangci.yaml ./...

# fmt formats the files.
.PHONY: fmt
fmt:
	@find . -iname "*.go" -not -path "./vendor/**" | xargs gofmt -s -w
	gofumpt -w -extra .
	@# gci's option should match with .golangci.yaml
	gci write . --skip-generated --skip-vendor --custom-order -s standard -s default -s 'prefix(github.com/gostaticanalysis)' -s 'prefix(github.com/cloverrose)' -s 'prefix(github.com/cloverrose/connectnew)'

# tidy updates the go.mod and go.sum files.
.PHONY: tidy
tidy:
	go mod tidy -v

# go/update updates go dependencies.
.PHONY: go/update
go/update:
	go get -u all
	go mod tidy -v

# go/update/testdata updates go dependencies.
.PHONY: go/update/testdata
go/update/testdata:
	cd testdata/src/a
	go get -u all
	go mod tidy -v

# test runs the tests.
.PHONY: test
test:
	# @go test $(args) -race -cover ./...
	@go test $(args) ./...

# build creates the binaries.
.PHONY: build
build:
	make build/pkgdep
	make build/pkgdep-tidy

# build/pkgdep creates the pkgdep binary.
.PHONY: build/pkgdep
build/pkgdep:
	@CGO_ENABLED=0 go build -o bin/pkgdep -v ./cmd/pkgdep

# build/pkgdep-tidy creates the pkgdep-tidy binary.
.PHONY: build/pkgdep-tidy
build/pkgdep-tidy:
	@CGO_ENABLED=0 go build -o bin/pkgdep-tidy -v ./cmd/pkgdep-tidy

# goreleaser/local runs goreleaser locally.
# see https://goreleaser.com/quick-start/
.PHONY: goreleaser/local
goreleaser/local:
	goreleaser release --snapshot --clean
