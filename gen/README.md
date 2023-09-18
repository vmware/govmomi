# Generating the types

This document describes how a VMware engineer can generate the vim types from internal builds:

## Requirements

* Docker

## Steps

1. Find the desired build of `vcenter-all`.
2. Click on the `vsphere-h5client` dependency.
3. Click on the `vimbase` dependency.
4. Download the published deliverable `wsdl.zip` and inflate it into the directory `./gen/sdk`.
5. Navigate back to the `vsphere-h5client` dependency.
6. Click on the `eam-vcenter` dependency.
7. Download the published deliverable `eam-wsdl.zip` and copy its `eam-messagetypes.xsd` and `eam-types.xsd` files into the directory `./gen/sdk`.
8. Download the latest OpenAPI specification for the vSphere Web Services API from https://developer.vmware.com/apis by clicking on `Resources`, and then clicking on `Download OpenAPI`. Once downloaded, please copy it to the `./gen/sdk` directory, renaming it to `vim.yaml`.

    ---

    :warning: TODO(akutz)

    Add step for finding the vim.yaml file from internal builds

    ---

9.  Open a terminal window.
10. Run `make generate-types`.

    ---

    :warning: **Please note**

    This can take a while the first time because it has to build the image used to generate the types, and building the image includes building OpenSSL and Ruby2, since the GoVmomi generator requires Ruby2, and Ruby2 requires an older version of OpenSSL than is available on recent container images.

    ---
