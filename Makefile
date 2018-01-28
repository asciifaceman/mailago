VERSION=$$(git rev-parse --verify HEAD)
HERE=$(pwd)
THIS_FILE := $(lastword $(MAKEFILE_LIST))
NAME=asciifaceman/mailago


SHELL=/bin/bash -e -o pipefail
# PROG
BIN_GO := $(call pathsearch,go)
BIN_GOVENDOR := $(call pathsearch,govendor)
BIN_PERL := $(call pathsearch,perl)


COLOR = \
	use Term::ANSIColor; \
	printf("    %s %s\n", colored(["BOLD $$ARGV[0]"], "[$$ARGV[1]]"), join(" ", @ARGV[2..$$\#ARGV]));

COLOR_SECTION = \
	use Term::ANSIColor; \
	printf("\n  %s\n\n", colored(["BOLD GREEN"], join(" ", @ARGV)));

COLOR_INDENT = \
	use Term::ANSIColor; use Text::Wrap; \
	$$Text::Wrap::columns=128; $$Text::Wrap::separator="!!"; \
	$$INDENT = (length($$ARGV[1]) + 6 + 1); \
	@l = split(/!!/, wrap("", "", join(" ", @ARGV[2..$$\#ARGV]))); \
	printf("    %s %s\n", colored(["BOLD $$ARGV[0]"], "[$$ARGV[1]]"), shift(@l)); \
	for(@l) { printf("%s%s\n", " "x$$INDENT, $$_) };


# Targets

.DEFAULT_GOAL := all

.PHONY: help all frontend buildosx build

## MailaGo Make
## 
## ###### Commands #######

help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' #thanks githuh.com/prwhite

all: | test build ## Clean and build osx
	@echo "Not building frontend. (in case you don't have npm)"

docker: ## build docker image
	docker build -t $(NAME):$(VERSION) .

run: ## Run mailago
	@go run main.go run

frontend:
	@echo "Building frontend ..."
	npm run build --prefix frontend/
	-@rm -r static
	-@mkdir static
	@cp -R frontend/build/* static/

deploy: ## Deploy to local docker. Must have docker installed and docker-compose
	docker-compose up --build -d

destroy: ## Destroy cluster deployed from this docker compose file.
	docker-compose down

test: ## run tests
	go test github.com/asciifaceman/mailago/mailago --cover

clean: ## clean target directory
	-rm -r target

buildosx: ## build osx executable
	@echo "Building target/mailago ..."
	@GOOS=darwin GOARCH=amd64 go build -o target/mailago
	@echo "Done."

build: ## build linux executable
	@echo "Building target/mailago ..."
	@GOOS=linux GOARCH=amd64 go build -o target/mailago
	@echo "Done."

