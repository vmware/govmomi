# © Broadcom. All Rights Reserved.
# The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
# SPDX-License-Identifier: Apache-2.0

# If you update this file, please follow
# https://www.thapaliya.com/en/writings/well-documented-makefiles/

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL := /usr/bin/env bash

# Print the help/usage when make is executed without any other arguments
.DEFAULT_GOAL := help

# CRI_BIN is the path to the container runtime binary.
ifeq (,$(strip $(GITHUB_RUN_ID)))
# Prefer podman locally.
CRI_BIN := $(shell command -v podman 2>/dev/null || command -v docker 2>/dev/null)
else
# Prefer docker in GitHub actions.
CRI_BIN := $(shell command -v docker 2>/dev/null || command -v podman 2>/dev/null)
endif
export CRI_BIN


## --------------------------------------
## Help
## --------------------------------------

.PHONY: help
help: ## Display usage
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target] \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)


## --------------------------------------
## Image
## --------------------------------------

IMAGE_NAME ?= govmomi-gen-types
IMAGE_TAG ?= latest
IMAGE ?= $(IMAGE_NAME):$(IMAGE_TAG)

.PHONY: image-build
image-build: ## Build the image for generating types
	$(CRI_BIN) build -t $(IMAGE) .


## --------------------------------------
## Generate
## --------------------------------------

ABS_PATH_PARENT_DIR := $(abspath $(dir $(shell pwd)))

#
# Please note the use of the .Gemfile.lock.tmp file below. This is to prevent
# the container from modifying the local copy of the Gemfile.lock that would
# otherwise be bind mounted into the container courtesy of the first bind mount.
#

.PHONY: generate-types
generate-types: image-build
generate-types: ## Generate the types
	@cp -f Gemfile.lock .Gemfile.lock.tmp
	$(CRI_BIN) run -it --rm \
	  -v $(ABS_PATH_PARENT_DIR):/govmomi \
	  -v $(ABS_PATH_PARENT_DIR)/gen/.Gemfile.lock.tmp:/govmomi/gen/Gemfile.lock \
	  $(IMAGE) \
	  /bin/bash -c 'bundle update --bundler && bundle install && ./gen.sh'


