.DEFAULT_GOAL := help

_YELLOW=\033[0;33m
_NC=\033[0m

PKG_DIRS = $(shell go list -f '{{.Dir}}' ./...)

.PHONY: help doc # generic commands
help: ## prints this help
	@grep -hE '^[\.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "${_YELLOW}%-16s${_NC} %s\n", $$1, $$2}'

setup: ## downloads dependencies
	go get -u github.com/robertkrimen/godocdown/godocdown
	go mod tidy

doc: ## generates markdown documentation
	@printf "${_YELLOW}generatic godoc documentation...${_NC}\n"
	@for d in ${PKG_DIRS}; do \
		godocdown -o $$d/README.md $$d; \
	done; \
