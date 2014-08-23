# govmomi

govmomi is a Go library for interacting with VMware ESXi and/or vCenter.

This repository is also host to a CLI on top of govmomi, named [govc](./govc).

## Status

As of late August 2014, govmomi is considered to be an **alpha version**. The
API exposed by the library is in flux and may be changed without prior notice.

After the library reaches v1.0, changes to the API will be subject to [semantic
versioning](http://semver.org).

## Compatibility

govmomi is built for and tested against ESXi and vCenter 5.5.

If you're able to use it against older versions of ESXi and/or vCenter, please
leave a note and we'll include it in this compatibility list.

## Documentation

The APIs exposed by this library very closely follow the API described in the
[VMware vSphere API Reference Documentation][apiref].
Refer to this document to become familiar with the upstream API.

[apiref]:http://pubs.vmware.com/vsphere-55/index.jsp#com.vmware.wssdk.apiref.doc/right-pane.html

**TODO(PN)**: Link to godoc.

## License

govmomi is published under the [Apache 2 license](LICENSE).
