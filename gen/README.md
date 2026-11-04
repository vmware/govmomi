# Generating the types

This document describes how a VMware engineer can generate the vim types from internal builds:

## Requirements

* Docker or Podman

## Inputs

Stage the vSphere SDK files in `./gen/sdk`:

1. The public files of `wsdl.zip`.
2. `eam-types.xsd` and `eam-messagetypes.xsd`.
3. `vi-json.yaml`, once per namespace, as `vim.yaml`, `pbm.yaml`, `sms.yaml`, `vslm.yaml` and `eam.yaml`.

Set `VIM_VERSION` in `gen.sh`.

## Steps

1. Open a terminal window.
2. Run `make generate-types`.

    ---

    :warning: **Please note**

    This can take a while the first time because it has to build the image used to generate the types.

    ---
