# govc

govc is a CLI around govmomi. Its name is analogous to rvc for Ruby.

## Status

As of late August 2014, govc is considered to be an **alpha version**. The API
exposed by the CLI is in flux and may be changed without prior notice.

## Installation

If you don't have govmomi package yet, install govc directly with `go get github.com/vmware/govmomi/govc`.

If you already have govmomi package, install govc with `go install github.com/vmware/govmomi/govc`.

## Usage

govc exposes its functionality through subcommands. Option flags
to these subcommands are often shared.

Common flags include:

* `-u`: ESXi or vCenter URL (ex: `user:pass@host`)
* `-debug`: Trace requests and responses (to `~/.govmomi/debug`)

Managed entities can be referred to by their absolute path or by their relative
path. For example, when specifying a datastore to use for a subcommand, you can
either specify it as `/mydatacenter/datastore/mydatastore`, or as
`mydatastore`. If you're not sure about the name of the datastore, or even the
full path to the datastore, you can specify a pattern to match. Both
`/*center/*/my*` (absolute) and `my*store` (relative) will resolve to the same
datastore, given there are no other datastores that match those globs.

The relative path in this example can only be used if the command can
umambigously resolve a datacenter to use as origin for the query. If no
datacenter is specified, govc defaults to the only datacenter, if there is only
one. The datacenter itself can be specified as a pattern as well, enabling the
following arguments: `-dc='my*' -ds='*store'`. The datastore pattern is looked
up and matched relative to the datacenter which itself is specified as a
pattern.

Besides specifying managed entities as arguments, they can also be specified
using environment variables. The following environment variables are used by govc
to set defaults:

* `GOVC_URL`: URL of ESXi or vCenter instance to connect to.

  > The URL scheme defaults to `https` and the URL path defaults to `/sdk`.
  > This means that specifying `user:pass@host` is equivalent to
  > `https://user:pass@host/sdk`.

* `GOVC_INSECURE`: Allow establishing insecure connections.

  > Use this option when the host you're connecting is using self-signed
  > certificates, or is otherwise trusted. Set this option to `1` to enable.

* `GOVC_DATACENTER`

* `GOVC_DATASTORE`

* `GOVC_NETWORK`

* `GOVC_RESOURCE_POOL`

* `GOVC_HOST`

* `GOVC_GUEST_LOGIN`: Guest credentials for guest operations

## Examples

* [Upload ssh public key to a VM](examples/lib/ssh.sh)

* [Create and configure a vCenter VM](examples/vcsa.sh)

## Projects using govc

* [Kubernetes vSphere Provider](https://github.com/GoogleCloudPlatform/kubernetes/tree/master/cluster/vsphere)

## License

govc is available under the [Apache 2 license](../LICENSE).
