# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE := on

# Versions.
K8S_VERSION=1.34.1

# Get the information about the platform on which the tools are built/run.
GOHOSTOS := $(shell go env GOHOSTOS)
GOHOSTARCH := $(shell go env GOHOSTARCH)
GOHOSTOSARCH := $(GOHOSTOS)_$(GOHOSTARCH)

# Default the GOOS and GOARCH values to be the same as the platform on which
# this Makefile is being executed.
export GOOS := $(GOHOSTOS)
export GOARCH := $(GOHOSTARCH)

# Directories.
BIN_DIR := bin

# Binaries.
CONTROLLER_GEN    := $(BIN_DIR)/$(GOHOSTOSARCH)/controller-gen
CONVERSION_GEN    := $(BIN_DIR)/$(GOHOSTOSARCH)/conversion-gen

## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Binaries
## --------------------------------------


controller-gen: $(CONTROLLER_GEN) ## Install controller-gen
$(CONTROLLER_GEN): go.mod
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) \
	go build -tags=govmomi_tools -o $(@) sigs.k8s.io/controller-tools/cmd/controller-gen

conversion-gen: $(CONVERSION_GEN) ## Install conversion-gen
$(CONVERSION_GEN): go.mod
	GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) \
	go build -tags=govmomi_tools -o $(@) k8s.io/code-generator/cmd/conversion-gen


## --------------------------------------
## Generate
## --------------------------------------

.PHONY: modules
modules: ## Runs go mod tidy to validate modules
	go mod tidy -v

.PHONY: modules-download
modules-download: ## Downloads and caches the modules
	go mod download

## --------------------------------------
## Cleanup / Verification
## --------------------------------------

.PHONY: clean
clean: ## Run all the clean targets
	rm -rf bin
