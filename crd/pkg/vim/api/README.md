# vim.vmware.com APIs

This module contains the actual VIM Kubernetes APIs.

## Environment Browser

* [`ClusterConfigTarget`](./v1alpha1/cluster_config_target_types.go) is a cluster-scoped resource surfaced for each vSphere cluster visible to Supervisor. The resource's status provides information for all of the cluster's physical attributes and hardware available to provision workloads.
* [`ConfigTarget`](./v1alpha1/config_target_types.go) is a namespace-scoped resource surfaced once per namespace. It allows administrators to control the availability of physical resources in a given namespace, even across zones.
* [`VirtualMachineConfigOptions`](./v1alpha1/virtualmachine_config_option_types.go) is a cluster-scoped resource surfaced for each distinct hardware version. Each resource then contains all of the virtual hardware and configuration options for the object's respective hardware version, for each guest ID supported.

## Virtual Machines

* [`VirtualMachine`](./v1alpha1/virtualmachine_types.go) is a namespace-scoped resource that may be used to provision and manage VMs. Unlike the VM Operator API's `VirtualMachine` CRD, the VIM CRD exposes everything that VIM would, enabling projects like VM Operator or the spherelet to build on top of this API instead of directly talking to vSphere.
  