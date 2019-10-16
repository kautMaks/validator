GO_BIN ?= go
CURL_BIN ?= curl
SHELL_BIN ?= sh

export PATH := $(PATH):/usr/local/go/bin

all: clean build

update:
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy
	cd buildin
	$(GO_BIN) get -u
	$(GO_BIN) mod tidy

linter-install: check-gopath
	cd ~
	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | $(SHELL_BIN) -s -- -b ${GOPATH}/bin v1.16.0
	$(GO_BIN) get -u github.com/Quasilyte/go-consistent
	$(GO_BIN) get -u github.com/mgechev/revive

test:
	$(GO_BIN) test -failfast ./...
	cd buildin
	$(GO_BIN) test -failfast ./...

lint:
	golangci-lint run ./...
	go-consistent -pedantic -v ./...
	revive -config revive.toml -formatter friendly ./...

check-gopath:
ifndef GOPATH
	$(error GOPATH is undefined)
endif
