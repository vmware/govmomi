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
8. Open a terminal window.
9. Run `make generate-types`.

    ---

    :warning: **Please note**

    This can take a while the first time because it has to build the image used to generate the types, and building the image includes building OpenSSL and Ruby2, since the GoVmomi generator requires Ruby2, and Ruby2 requires an older version of OpenSSL than is available on recent container images.

    ---
