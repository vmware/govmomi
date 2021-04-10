# vcsim - A vCenter and ESXi API based simulator

This package implements a vSphere Web Services (SOAP) SDK endpoint intended for
testing consumers of the API.  While the mock framework is written in the Go
language, it can be used by any language that can talk to the vSphere API.

## Installation

```sh
% export GOPATH=$HOME/gopath
% go get -u github.com/vmware/govmomi/vcsim
% $GOPATH/bin/vcsim -h
```

## Usage

The **vcsim** program by default creates a *vCenter* model with a datacenter,
hosts, cluster, resource pools, networks and a datastore.  The naming is similar
to that of the original *vcsim* mode that was included with vCenter.  The number
of resources can be increased or decreased using the various resource type
flags.  Resources can also be created and removed using the API. In fact, vcsim
itself uses the vSphere API generate its inventory.

``` console
% vcsim -h # pruned to model type flags used in the Examples section
Usage of vcsim:
  -app int
        Number of virtual apps per compute resource
  -cluster int
        Number of clusters (default 1)
  -dc int
        Number of datacenters (default 1)
  -ds int
        Number of local datastores (default 1)
  -folder int
        Number of folders
  -host int
        Number of hosts per cluster (default 3)
  -nsx int
        Number of NSX backed opaque networks
  -pg int
        Number of port groups (default 1)
  -pod int
        Number of storage pods per datacenter
  -pool int
        Number of resource pools per compute resource
  -standalone-host int
        Number of standalone hosts (default 1)
  -vm int
        Number of virtual machines per resource pool (default 2)
```

[model]:https://godoc.org/github.com/vmware/govmomi/simulator#Model

### Version Information

To print detailed (build) information for vcsim run: `vcsim version`. 

## Examples

The following examples illustrate how **vcsim** flags can be used to change the
generated inventory.  Each example assumes **GOVC_URL** is set to vcsim's default
[listen address](#listen-address):

```console
% export GOVC_URL=https://user:pass@127.0.0.1:8989
```

### Default vCenter inventory

```console
% $GOPATH/bin/vcsim

% govc find -l
Folder                       /
Datacenter                   /DC0
Folder                       /DC0/vm
VirtualMachine               /DC0/vm/DC0_H0_VM0
VirtualMachine               /DC0/vm/DC0_H0_VM1
VirtualMachine               /DC0/vm/DC0_C0_RP0_VM0
VirtualMachine               /DC0/vm/DC0_C0_RP0_VM1
Folder                       /DC0/host
ComputeResource              /DC0/host/DC0_H0
HostSystem                   /DC0/host/DC0_H0/DC0_H0
ResourcePool                 /DC0/host/DC0_H0/Resources
ClusterComputeResource       /DC0/host/DC0_C0
HostSystem                   /DC0/host/DC0_C0/DC0_C0_H0
HostSystem                   /DC0/host/DC0_C0/DC0_C0_H1
HostSystem                   /DC0/host/DC0_C0/DC0_C0_H2
ResourcePool                 /DC0/host/DC0_C0/Resources
Folder                       /DC0/datastore
Datastore                    /DC0/datastore/LocalDS_0
Folder                       /DC0/network
Network                      /DC0/network/VM Network
DistributedVirtualSwitch     /DC0/network/DVS0
DistributedVirtualPortgroup  /DC0/network/DVS0-DVUplinks-9
DistributedVirtualPortgroup  /DC0/network/DC0_DVPG0
```

### Default standalone ESX inventory

With the `-esx` flag, vcsim behaves as a standalone ESX host without any vCenter
specific features.

```console
% $GOPATH/vcsim -esx

% govc find
/
/ha-datacenter
/ha-datacenter/vm
/ha-datacenter/vm/ha-host_VM0
/ha-datacenter/vm/ha-host_VM1
/ha-datacenter/host
/ha-datacenter/host/localhost.localdomain
/ha-datacenter/host/localhost.localdomain/localhost.localdomain
/ha-datacenter/host/localhost.localdomain/Resources
/ha-datacenter/datastore
/ha-datacenter/datastore/LocalDS_0
/ha-datacenter/network
/ha-datacenter/network/VM Network
```

### Customizing inventory

Model flags can be specified to increase or decrease the generated inventory.

``` console
% vcsim -dc 2 -folder 1 -ds 4 -pod 1 -nsx 2 -pool 2 -app 1

% govc find -l
Folder                       /
Datacenter                   /DC0
Folder                       /DC0/vm
VirtualMachine               /DC0/vm/DC0_H0_VM0
VirtualMachine               /DC0/vm/DC0_H0_VM1
VirtualMachine               /DC0/vm/DC0_C0_RP0_VM0
VirtualMachine               /DC0/vm/DC0_C0_RP0_VM1
VirtualMachine               /DC0/vm/DC0_C0_APP0_VM0
VirtualMachine               /DC0/vm/DC0_C0_APP0_VM1
Folder                       /DC0/host
ComputeResource              /DC0/host/DC0_H0
HostSystem                   /DC0/host/DC0_H0/DC0_H0
ResourcePool                 /DC0/host/DC0_H0/Resources
ClusterComputeResource       /DC0/host/DC0_C0
HostSystem                   /DC0/host/DC0_C0/DC0_C0_H0
HostSystem                   /DC0/host/DC0_C0/DC0_C0_H1
HostSystem                   /DC0/host/DC0_C0/DC0_C0_H2
ResourcePool                 /DC0/host/DC0_C0/Resources
ResourcePool                 /DC0/host/DC0_C0/Resources/DC0_C0_RP1
ResourcePool                 /DC0/host/DC0_C0/Resources/DC0_C0_RP2
VirtualApp                   /DC0/host/DC0_C0/Resources/DC0_C0_APP0
Folder                       /DC0/datastore
StoragePod                   /DC0/datastore/DC0_POD0
Datastore                    /DC0/datastore/LocalDS_0
Datastore                    /DC0/datastore/LocalDS_1
Datastore                    /DC0/datastore/LocalDS_2
Datastore                    /DC0/datastore/LocalDS_3
Folder                       /DC0/network
Network                      /DC0/network/VM Network
DistributedVirtualSwitch     /DC0/network/DVS0
DistributedVirtualPortgroup  /DC0/network/DVS0-DVUplinks-10
DistributedVirtualPortgroup  /DC0/network/DC0_DVPG0
OpaqueNetwork                /DC0/network/DC0_NSX0
OpaqueNetwork                /DC0/network/DC0_NSX1
Folder                       /F0
Datacenter                   /F0/DC1
Folder                       /F0/DC1/vm
Folder                       /F0/DC1/vm/F0
VirtualMachine               /F0/DC1/vm/F0/DC1_H0_VM0
VirtualMachine               /F0/DC1/vm/F0/DC1_H0_VM1
VirtualMachine               /F0/DC1/vm/F0/DC1_C0_RP0_VM0
VirtualMachine               /F0/DC1/vm/F0/DC1_C0_RP0_VM1
VirtualMachine               /F0/DC1/vm/F0/DC1_C0_APP0_VM0
VirtualMachine               /F0/DC1/vm/F0/DC1_C0_APP0_VM1
Folder                       /F0/DC1/host
Folder                       /F0/DC1/host/F0
ComputeResource              /F0/DC1/host/F0/DC1_H0
HostSystem                   /F0/DC1/host/F0/DC1_H0/DC1_H0
ResourcePool                 /F0/DC1/host/F0/DC1_H0/Resources
ClusterComputeResource       /F0/DC1/host/F0/DC1_C0
HostSystem                   /F0/DC1/host/F0/DC1_C0/DC1_C0_H0
HostSystem                   /F0/DC1/host/F0/DC1_C0/DC1_C0_H1
HostSystem                   /F0/DC1/host/F0/DC1_C0/DC1_C0_H2
ResourcePool                 /F0/DC1/host/F0/DC1_C0/Resources
ResourcePool                 /F0/DC1/host/F0/DC1_C0/Resources/DC1_C0_RP1
ResourcePool                 /F0/DC1/host/F0/DC1_C0/Resources/DC1_C0_RP2
VirtualApp                   /F0/DC1/host/F0/DC1_C0/Resources/DC1_C0_APP0
Folder                       /F0/DC1/datastore
StoragePod                   /F0/DC1/datastore/DC1_POD0
Folder                       /F0/DC1/datastore/F0
Datastore                    /F0/DC1/datastore/F0/LocalDS_0
Datastore                    /F0/DC1/datastore/F0/LocalDS_1
Datastore                    /F0/DC1/datastore/F0/LocalDS_2
Datastore                    /F0/DC1/datastore/F0/LocalDS_3
Folder                       /F0/DC1/network
Network                      /F0/DC1/network/VM Network
Folder                       /F0/DC1/network/F0
DistributedVirtualSwitch     /F0/DC1/network/F0/DVS0
DistributedVirtualPortgroup  /F0/DC1/network/F0/DVS0-DVUplinks-69
DistributedVirtualPortgroup  /F0/DC1/network/F0/DC1_DVPG0
OpaqueNetwork                /F0/DC1/network/F0/DC1_NSX0
OpaqueNetwork                /F0/DC1/network/F0/DC1_NSX1
```

Create yourself a Datastore cluster:

``` console
% govc object.mv /F0/DC0/datastore/F0/LocalDS_[123] /DC1/datastore/DC0_POD0

% govc find -l /DC0/datastore
Folder      /DC0/datastore
StoragePod  /DC0/datastore/DC0_POD0
Datastore   /DC0/datastore/LocalDS_0
Datastore   /DC0/datastore/DC0_POD0/LocalDS_1
Datastore   /DC0/datastore/DC0_POD0/LocalDS_2
Datastore   /DC0/datastore/DC0_POD0/LocalDS_3
```

### Starting with empty inventory

The model flags when set to 0 can be used to turn off generation of any type.
With **Datacenter** generation turned off, the inventory will be empty:

```console
% $GOPATH/vcsim -dc 0
% govc find
/
```

You can create your own inventory using the API or govc:

``` console
% govc datacenter.create godc

% govc cluster.create gocluster

% govc cluster.add -hostname gohost1 -username user -password pass -noverify

% govc datastore.create -type local -name gostore -path /tmp gocluster/*

% govc vm.create -ds gostore -cluster gocluster govm1

% govc find -l
Folder                  /
Datacenter              /godc
Folder                  /godc/vm
VirtualMachine          /godc/vm/govm1
Folder                  /godc/host
ClusterComputeResource  /godc/host/gocluster
HostSystem              /godc/host/gocluster/gohost1
ResourcePool            /godc/host/gocluster/Resources
Folder                  /godc/datastore
Datastore               /godc/datastore/gostore
Folder                  /godc/network
Network                 /godc/network/VM Network
```

## Generated inventory names

The generated names include a prefix per-type and integer suffix per-instance.
See the [simulator.Model][model] documentation for a complete list of type prefixes.
For example, the name **DC0_C1_RP0_VM6** is composed of:

| Prefix | Instance | Type                   | Flag     |
|--------|----------|------------------------|----------|
| DC     | 0        | Datacenter             | -dc      |
| C      | 1        | ClusterComputeResource | -cluster |
| RP     | 0        | ResourcePool           | -pool    |
| VM     | 6        | VirtualMachine         | -vm      |

VMs with Clusters include the ResourcePool in their name, while VMs in standalone
hosts include the host name instead.
For example, the name **DC0_H0_VM1** is composed of:

| Prefix | Instance | Type           | Flag             |
|--------|----------|----------------|------------------|
| DC     | 0        | Datacenter     | -dc              |
| H      | 0        | HostSystem     | -standalone-host |
| VM     | 1        | VirtualMachine | -vm              |

## Supported methods

The simulator supports a subset of API methods.  However, the generated govmomi
code includes all types and methods defined in the vmodl, which can be used to
implement any method documented in the [VMware vSphere API Reference][apiref].

To see the list of supported methods:

```console
% curl -sk https://user:pass@127.0.0.1:8989/about
```

[apiref]:https://code.vmware.com/apis/196/vsphere

## Listen address

The default vcsim listen address is `127.0.0.1:8989`.  Use the `-l` flag to
listen on another address:

```console
% vcsim -l 10.118.69.224:8989 # specific address

% vcsim -l :8989 # any address
```

When given a port value of `0`, an unused port will be chosen.  You can then
source the GOVC_URL from another process, for example:

```console
% govc_sim_env=$TMPDIR/vcsim-$(uuidgen)

% mkfifo $govc_sim_env

% vcsim -l 127.0.0.1:0 -E $govc_sim_env &

% eval "$(cat $govc_sim_env)"

# ... run tests ...

% kill $GOVC_SIM_PID
% rm -f $govc_sim_env
```

Tests written in Go can also use the [simulator package](https://godoc.org/github.com/vmware/govmomi/simulator)
directly, rather than the vcsim binary.

## Feature Details

For more details on vcsim features, see the project [wiki](https://github.com/vmware/govmomi/wiki/vcsim-features).

## Projects using vcsim

* [VMware VIC Engine](https://github.com/vmware/vic)

* [Kubernetes](https://github.com/kubernetes/kubernetes/tree/master/pkg/cloudprovider/providers/vsphere)

* [Ansible](https://github.com/ansible/vcenter-test-container)

* [Telegraf](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/vsphere)

## Blog posts

* [Beginning vCenter Server simulation with vcsim]( by Abhijeet Kasurde

* [vCenter & ESXi API based simulator] by William Lam

* [vCenter Simulator Docker Container] by Brian Bunke

[Beginning vCenter Server simulation with vcsim]: https://opensourceforu.com/2017/10/vcenter-server-simulation-govcsim/

[vCenter & ESXi API based simulator]: https://www.virtuallyghetto.com/2017/04/govcsim-neat-incubation-project-vcenter-server-esxi-api-based-simulator.html

[vCenter Simulator Docker Container]: https://www.brianbunke.com/blog/2018/12/31/vcenter-simulator-ci/

## Related projects

* [LocalStack](https://github.com/localstack/localstack/blob/master/README.md#why-localstack)

## License

vcsim is available under the [Apache 2 license](../LICENSE).

## Name

Pronounced "v-c-sim", short for "vCenter Simulator"
