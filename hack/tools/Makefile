# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# Directories
BIN_DIR := bin

# Get information about this OS.
HOST_OS :=   $(shell go env GOHOSTOS)
HOST_ARCH := $(shell go env GOHOSTARCH)

# Binaries.
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint


## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


## --------------------------------------
## Binaries
## --------------------------------------

.PHONY: $(GOLANGCI_LINT)
golangci-lint: $(GOLANGCI_LINT) ## Install golangci-lint
$(GOLANGCI_LINT): go.mod
	go build -tags=govmomi_tools -o $@ github.com/golangci/golangci-lint/cmd/golangci-lint


## --------------------------------------
## Generate
## --------------------------------------

.PHONY: mod
mod: ## Runs go mod tidy to validate modules
	go mod tidy -v

.PHONY: mod-get
mod-get: ## Downloads and caches the modules
	go mod download


## --------------------------------------
## Cleanup / Verification
## --------------------------------------

.PHONY: clean
clean: ## Removes the generated binaries
	go clean -i -v .
	rm -rf $(BIN_DIR)
