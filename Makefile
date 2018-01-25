VERSION=$$(git rev-parse --verify HEAD)
HERE=$(pwd)
THIS_FILE := $(lastword $(MAKEFILE_LIST))
NAME=asciifaceman/mailago


SHELL=/bin/bash -e -o pipefail
### PROG
BIN_GO := $(call pathsearch,go)
BIN_GOVENDOR := $(call pathsearch,govendor)
BIN_PERL := $(call pathsearch,perl)

### Targets

.DEFAULT_GOAL := all

.PHONY: all buildosx build

all: | clean buildosx docker

docker:
	docker build -t $(NAME):$(VERSION) .

run:
	go run main.go run

test:
	go test -v ./...

clean:
	rm mailago

buildosx:
	@GOOS=darwin GOARCH=amd64 go build

build:
	@GOOS=linux GOARCH=amd64 go build