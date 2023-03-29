# govc

`govc` is a vSphere CLI built on top of `govmomi`.

The CLI is designed to be a user friendly CLI alternative to the GUI and well suited for automation tasks.
It also acts as a [test harness](test) for the `govmomi` APIs and provides working examples of how to use the APIs.

## Installation

### Docker

The official `govc` [Docker images](https://hub.docker.com/r/vmware/govc) are built from this [Dockerfile](../Dockerfile.govc).

### Binaries
You can find prebuilt `govc` binaries on the [releases page](https://github.com/vmware/govmomi/releases).

You can download and install a binary locally like this:

```bash
# extract govc binary to /usr/local/bin
# note: the "tar" command must run with root permissions
curl -L -o - "https://github.com/vmware/govmomi/releases/latest/download/govc_$(uname -s)_$(uname -m).tar.gz" | tar -C /usr/local/bin -xvzf - govc
```

### Source

#### Install via `go install`

To build `govc` from source, first install the [Go
toolchain](https://golang.org/dl/). You can then install the latest `govc` from
Github using:

```bash
go install github.com/vmware/govmomi/govc@latest
```

**Note:** `govmomi` and its binaries use [Go
modules](https://golang.org/ref/mod), i.e. explicitly setting `GOPATH` is not
required anymore. To inject build variables (see details
[below](#install-via-goreleaser)) used by `govc version [-l]`, `GOFLAGS` can be
defined and are honored by `go get`.

⚠️ Make sure `$GOPATH/bin` is in your `PATH` to use the version installed from
source.

If you've made local modifications to the repository at
`$GOPATH/src/github.com/vmware/govmomi`, you can install using:

```bash
go install github.com/vmware/govmomi/govc
```

#### Install via `goreleaser`

You can also build `govc` following our release process using `goreleaser`
(requires [Go toolchain](https://golang.org/dl/)). This will ensure that build
time variables are correctly injected. Build (linker) flags and injection are
defined in [.goreleaser.yaml](./../.goreleaser.yml) and automatically set as
`GOFLAGS` when building with `goreleaser`.

Install `goreleaser` as per the installation
[instructions](https://goreleaser.com/install/), then:

```bash
git clone https://github.com/vmware/govmomi.git
cd govmomi

# pick a tag (>=v0.25.0)
RELEASE=v0.25.0

git checkout ${RELEASE}

# build for the host OS/ARCH, otherwise omit --single-target
# binaries are placed in respective subdirectories in ./dist/
goreleaser build --rm-dist --single-target
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
datacenter is specified, `govc` defaults to the only datacenter, if there is
only one. The datacenter itself can be specified as a pattern as well, enabling
the following arguments: `-dc='my*' -ds='*store'`. The datastore pattern is
looked up and matched relative to the datacenter which itself is specified as a
pattern.

Besides specifying managed entities as arguments, they can also be specified
using environment variables. The following environment variables are used by
`govc` to set defaults:

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

    ```bash
    export GOVC_TLS_CA_CERTS=~/.govc_ca.crt
    # Use path separator to specify multiple files:
    export GOVC_TLS_CA_CERTS=~/ca-certificates/bar.crt:~/ca-certificates/foo.crt
    ```

* `GOVC_TLS_KNOWN_HOSTS`: File(s) for thumbprint based certificate verification.

    Thumbprint based verification can be used in addition to or as an alternative to
    `GOVC_TLS_CA_CERTS` for self-signed certificates.  Example:

    ```bash
    export GOVC_TLS_KNOWN_HOSTS=~/.govc_known_hosts
    govc about.cert -u host -k -thumbprint | tee -a $GOVC_TLS_KNOWN_HOSTS
    govc about -u user:pass@host
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

* `GOVC_VI_JSON`: Uses JSON transport instead of SOAP (Experimental; Usable only
for vim25 APIs in vCenter 8.0u1)

## Troubleshooting

### Verbose Flag

The `-verbose` flag writes request and response data to stderr, in a format more compact than the `-trace` or `-debug` flags.
The output includes the request method name with abbreviated input parameters. The response data is more detailed and may include
structures formatted as Go code, such as Task property updates.  The value of some properties will the `govc` `object.collect`
command that can be used to view the actual value.

### Trace flag

The `-trace` flag writes HTTP request and response data to stderr.
XML bodies are formatted using `xmlstarlet` if installed and JSON bodies using `jq` if installed.
Formatting can be disabled via `export GOVC_DEBUG_FORMAT=false`.

If both `-trace` and `-verbose` flags are specified, request and response data is formatted as Go code.

### Debug Flag

The`-debug` flag traces vSphere API calls similar to the `-trace` flag, but saves to files rather than stderr.
When the `-debug` flag is specified, the default behavior is to put the output in `~/.govmomi/debug/<run timestamp>`.
In that directory will be four (4) files per API call.

```bash
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

The [debug-format](../scripts/debug-format.sh) script can be used to format the debug output similar to the `-trace` flag.

### Print Version Information

For troubleshooting and when filing issues, get build related details with:

```console
$ govc version -l
Build Version: v0.25.0-next
Build Commit: e86da96e
Build Date: 2021-04-19T10:29:57Z
```

#### stderr debug

If you prefer debug output to be sent to stderr and seen while the command is running you can override the file behavior
by setting the debug path to a dash: `export GOVC_DEBUG_PATH=-`

### Environment variables

If you're using environment variables to set `GOVC_URL`, verify the values are set as expected:

```bash
govc env
```

### Connection issues

Check your proxy settings:

```bash
env | grep -i https_proxy
```

Test connection using `curl`:

```bash
curl --verbose -k -X POST https://x.x.x.x/sdk
```

### MSYS2 (Windows)

Inventory path arguments with a leading '/' are subject
to [Posix path conversion](http://www.mingw.org/wiki/Posix_path_conversion).

### NotAuthenticated

When connecting to a non-TLS endpoint, Go's http.Client will not send Secure
cookies, resulting in a `NotAuthenticated` error. For example, running `govc`
directly against the vCenter vpxd endpoint at `http://127.0.0.1:8085`. Set the
environment variable `GOVMOMI_INSECURE_COOKIES=true` to workaround this:

```bash
GOVMOMI_INSECURE_COOKIES=true govc ls -u http://user:pass@127.0.0.1:8085
```

## Examples

Several examples are embedded in the govc command [help](USAGE.md)

* [Upload ssh public key to a VM](examples/lib/ssh.sh)

* [Create a CoreOS VM](https://github.com/vmware/govmomi/blob/main/toolbox/toolbox-test.sh)

* [Create a Debian VM](https://github.com/kubernetes/kubernetes/tree/master/cluster/vsphere)

* [Create a Windows VM](https://github.com/dougm/govc-windows-box/blob/master/provision-esx.sh)

* [Create an ESX VM](../scripts/vcsa/create-esxi-vm.sh)

* [Create a vCenter VM](../scripts/vcsa/create-vcsa-vm.sh)

* [Create a Cluster](../scripts/vcsa/create-cluster.sh)

## Status

Changes to the CLI are subject to [semantic versioning](http://semver.org).

Refer to the [CHANGELOG](../CHANGELOG.md) for version to version changes.

When new `govc` commands or flags are added, the PATCH version will be
incremented.  This enables you to require a minimum version from within a
script, for example:

```bash
govc version -require 0.24
```

## Projects using `govc`

* [Emacs govc package](./emacs)

* [Kubernetes vSphere Cloud Provider](https://github.com/kubernetes/cloud-provider-vsphere)

* [VMware VIC Engine](https://github.com/vmware/vic)

* [vSphere Docker Volume Service](https://github.com/vmware/docker-volume-vsphere)

* [golang/build](https://github.com/golang/build)

## Related projects

* [rvc](https://github.com/vmware/rvc)

## License

`govc` is available under the [Apache 2 license](../LICENSE).

## Name

Pronounced "go-v-c", short for "Go(lang) vCenter CLI".
