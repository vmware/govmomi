# govc

govc is a vSphere CLI built on top of govmomi.

The CLI is designed to be a user friendly CLI alternative to the GUI and well suited for automation tasks.
It also acts as a [test harness](test) for the govmomi APIs and provides working examples of how to use the APIs.

## Installation

You can find prebuilt govc binaries on the [releases page](https://github.com/vmware/govmomi/releases).

Download and install a binary locally like this:

``` console
% curl -L $URL_TO_BINARY | gunzip > /usr/local/bin/govc
% chmod +x /usr/local/bin/govc
```

### Source

To build govc from source, first install the [Go toolchain](https://golang.org/dl/).

Make sure to set the environment variable [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH).

You can then install the latest govc from github using:

``` console
% go get -u github.com/vmware/govmomi/govc
```

Make sure `$GOPATH/bin` is in your `PATH` to use the version installed from source.

If you've made local modifications to the repository at `$GOPATH/src/github.com/vmware/govmomi`, you can install using:

``` console
% go install github.com/vmware/govmomi/govc
```

## Usage

For the complete list of commands and flags, refer to the [USAGE](USAGE.md) document.

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

    The URL scheme defaults to `https` and the URL path defaults to `/sdk`.
    This means that specifying `user:pass@host` is equivalent to
    `https://user:pass@host/sdk`.

    If username or password includes special characters like `\`, `#` or `:` you can use
    `GOVC_USERNAME` and `GOVC_PASSWORD` to have a simple `GOVC_URL`

    When using govc against VMware Workstation, GOVC_URL can be set to "localhost"
    without a user or pass, in which case local ticket based authentication is used.

* `GOVC_USERNAME`: USERNAME to use if not specified in GOVC_URL.

* `GOVC_PASSWORD`: PASSWORD to use if not specified in GOVC_URL.

* `GOVC_TLS_CA_CERTS`: Override system root certificate authorities.

    ``` console
    $ export GOVC_TLS_CA_CERTS=~/.govc_ca.crt
    # Use path separator to specify multiple files:
    $ export GOVC_TLS_CA_CERTS=~/ca-certificates/bar.crt:~/ca-certificates/foo.crt
    ```

* `GOVC_TLS_KNOWN_HOSTS`: File(s) for thumbprint based certificate verification.

    Thumbprint based verification can be used in addition to or as an alternative to
    `GOVC_TLS_CA_CERTS` for self-signed certificates.  Example:

    ``` console
    $ export GOVC_TLS_KNOWN_HOSTS=~/.govc_known_hosts
    $ govc about.cert -u host -k -thumbprint | tee -a $GOVC_TLS_KNOWN_HOSTS
    $ govc about -u user:pass@host
    ```

* `GOVC_TLS_HANDSHAKE_TIMEOUT`: Limits the time spent performing the TLS handshake.

* `GOVC_INSECURE`: Disable certificate verification.

    This option sets Go's `tls.Config.InsecureSkipVerify` flag and is false by default.
    Quoting https://golang.org/pkg/crypto/tls/#Config:
    > `InsecureSkipVerify` controls whether a client verifies the
    > server's certificate chain and host name.
    >
    > If `InsecureSkipVerify` is true, TLS accepts any certificate
    > presented by the server and any host name in that certificate.
    >
    > In this mode, TLS is susceptible to man-in-the-middle attacks.
    > This should be used only for testing.

* `GOVC_DATACENTER`

* `GOVC_DATASTORE`

* `GOVC_NETWORK`

* `GOVC_RESOURCE_POOL`

* `GOVC_HOST`

* `GOVC_GUEST_LOGIN`: Guest credentials for guest operations

* `GOVC_VIM_NAMESPACE`: Vim namespace defaults to `urn:vim25`

* `GOVC_VIM_VERSION`: Vim version defaults to `6.0`

## Troubleshooting

### Debug Flag
There is a `-debug` flag which you can use to debug the calls made to the vSphere API and there are some environment
variables you can use to configure how it works. If you turn on the `-debug` flag the default behavior is to put the
output in `~/.govmomi/debug/<run timestamp>`. In that directory will be four (4) files per API call.

```
1-0001.req.headers #headers from the request sent to the API
1-0001.req.xml #body content from request sent to the API
1-0001.res.headers #headers from the response from the API
1-0001.res.xml #body from the respnse from the API
```

In that filename the `0001` represents the an incremented call order and will increment for each time the SOAP client
makes an API call.

To configure the debug output you can use two environment variables.
* `GOVC_DEBUG_PATH`: defaults to ~/.govmomi/debug
* `GOVC_DEBUG_PATH_RUN`: defaults to timestamp of the run

#### stdout debug
If you prefer debug output to be sent to stdout and seen while the command is running you can override the file behavior
by setting the debug path to a dash: `export GOVC_DEBUG_PATH=-`

### Environment variables

If you're using environment variables to set `GOVC_URL`, verify the values are set as expected:

``` console
% govc env
```

### Connection issues

Check your proxy settings:

``` console
% env | grep -i https_proxy
```

Test connection using curl:
``` console
% curl --verbose -k -X POST https://x.x.x.x/sdk
```

### MSYS2 (Windows)

Inventory path arguments with a leading '/' are subject
to [Posix path conversion](http://www.mingw.org/wiki/Posix_path_conversion).

### NotAuthenticated

When connecting to a non-TLS endpoint, Go's http.Client will not send Secure cookies, resulting in a `NotAuthenticated` error.
For example, running govc directly against the vCenter vpxd endpoint at `http://127.0.0.1:8085`.
Set the environment variable `GOVMOMI_INSECURE_COOKIES=true` to workaround this:

``` console
% GOVMOMI_INSECURE_COOKIES=true govc ls -u http://user:pass@127.0.0.1:8085
```

## Examples

Several examples are embedded in the govc command [help](USAGE.md)

* [Upload ssh public key to a VM](examples/lib/ssh.sh)

* [Create a CoreOS VM](https://github.com/vmware/govmomi/blob/master/toolbox/toolbox-test.sh)

* [Create a Debian VM](https://github.com/kubernetes/kubernetes/tree/master/cluster/vsphere)

* [Create a Windows VM](https://github.com/dougm/govc-windows-box/blob/master/provision-esx.sh)

* [Create an ESX VM](../scripts/vcsa/create-esxi-vm.sh)

* [Create a vCenter VM](../scripts/vcsa/create-vcsa-vm.sh)

* [Create a Cluster](../scripts/vcsa/create-cluster.sh)

## Status

Changes to the cli are subject to [semantic versioning](http://semver.org).

Refer to the [CHANGELOG](CHANGELOG.md) for version to version changes.

When new govc commands or flags are added, the PATCH version will be incremented.  This enables you to require a minimum
version from within a script, for example:

``` console
% govc version -require 0.14
```

## Projects using govc

* [Emacs govc package](./emacs)

* [Kubernetes vSphere Provider](https://github.com/kubernetes/kubernetes/tree/master/cluster/vsphere)

* [VMware VIC Engine](https://github.com/vmware/vic)

* [vSphere Docker Volume Service](https://github.com/vmware/docker-volume-vsphere)

* [golang/build](https://github.com/golang/build)

## Related projects

* [rvc](https://github.com/vmware/rvc)

## License

govc is available under the [Apache 2 license](../LICENSE).

## Name

Pronounced "go-v-c", short for "Go(lang) vCenter CLI".
