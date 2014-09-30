/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govmomi

import "github.com/vmware/govmomi/vim25/types"

type Reference interface {
	Reference() types.ManagedObjectReference
}

func NewReference(e types.ManagedObjectReference) Reference {
	switch e.Type {
	case "Folder":
		return &Folder{ManagedObjectReference: e}
	case "StoragePod":
		return &StoragePod{
			Folder{ManagedObjectReference: e},
		}
	case "Datacenter":
		return &Datacenter{ManagedObjectReference: e}
	case "VirtualMachine":
		return &VirtualMachine{ManagedObjectReference: e}
	case "VirtualApp":
		return &VirtualApp{
			ResourcePool{ManagedObjectReference: e},
		}
	case "ComputeResource":
		return &ComputeResource{ManagedObjectReference: e}
	case "ClusterComputeResource":
		return &ClusterComputeResource{
			ComputeResource{ManagedObjectReference: e},
		}
	case "HostSystem":
		return &HostSystem{ManagedObjectReference: e}
	case "Network":
		return &Network{ManagedObjectReference: e}
	case "ResourcePool":
		return &ResourcePool{ManagedObjectReference: e}
	case "DistributedVirtualSwitch":
		return &DistributedVirtualSwitch{ManagedObjectReference: e}
	case "VmwareDistributedVirtualSwitch":
		return &VmwareDistributedVirtualSwitch{
			DistributedVirtualSwitch{ManagedObjectReference: e},
		}
	case "DistributedVirtualPortgroup":
		return &DistributedVirtualPortgroup{ManagedObjectReference: e}
	case "Datastore":
		return &Datastore{ManagedObjectReference: e}
	default:
		panic("Unknown managed entity: " + e.Type)
	}
}
