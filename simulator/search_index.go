// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"strings"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type SearchIndex struct {
	mo.SearchIndex
}

func (s *SearchIndex) FindByDatastorePath(ctx *Context, r *types.FindByDatastorePath) soap.HasFault {
	res := &methods.FindByDatastorePathBody{Res: new(types.FindByDatastorePathResponse)}

	ctx.Map.m.Lock()
	defer ctx.Map.m.Unlock()

	for ref, obj := range ctx.Map.objects {
		vm, ok := asVirtualMachineMO(obj)
		if !ok {
			continue
		}

		if vm.Config.Files.VmPathName == r.Path {
			res.Res.Returnval = &ref
			break
		}
	}

	return res
}

func (s *SearchIndex) FindByInventoryPath(ctx *Context, req *types.FindByInventoryPath) soap.HasFault {
	body := &methods.FindByInventoryPathBody{Res: new(types.FindByInventoryPathResponse)}

	root := ctx.Map.content().RootFolder
	o := &root

	if req.InventoryPath == "/" {
		body.Res.Returnval = o
		return body
	}

	split := func(c rune) bool {
		return c == '/'
	}
	path := strings.FieldsFunc(req.InventoryPath, split)
	if len(path) < 1 {
		return body
	}

	for _, name := range path {
		f := s.FindChild(ctx, &types.FindChild{Entity: *o, Name: name})

		o = f.(*methods.FindChildBody).Res.Returnval
		if o == nil {
			break
		}
	}

	body.Res.Returnval = o

	return body
}

func (s *SearchIndex) FindChild(ctx *Context, req *types.FindChild) soap.HasFault {
	body := &methods.FindChildBody{}

	obj := ctx.Map.Get(req.Entity)

	if obj == nil {
		body.Fault_ = Fault("", &types.ManagedObjectNotFound{Obj: req.Entity})
		return body
	}

	body.Res = new(types.FindChildResponse)

	var children []types.ManagedObjectReference

	switch e := obj.(type) {
	case *Datacenter:
		children = []types.ManagedObjectReference{e.VmFolder, e.HostFolder, e.DatastoreFolder, e.NetworkFolder}
	case *Folder:
		children = e.ChildEntity
	case *mo.ComputeResource:
		children = e.Host
		children = append(children, *e.ResourcePool)
	case *ClusterComputeResource:
		children = e.Host
		children = append(children, *e.ResourcePool)
	case *ResourcePool:
		children = e.ResourcePool.ResourcePool
		children = append(children, e.Vm...)
	case *VirtualApp:
		children = e.ResourcePool.ResourcePool
		children = append(children, e.Vm...)
	}

	match := ctx.Map.FindByName(req.Name, children)

	if match != nil {
		ref := match.Reference()
		body.Res.Returnval = &ref
	}

	return body
}

func (s *SearchIndex) FindByUuid(ctx *Context, req *types.FindByUuid) soap.HasFault {
	body := &methods.FindByUuidBody{Res: new(types.FindByUuidResponse)}

	ctx.Map.m.Lock()
	defer ctx.Map.m.Unlock()

	if req.VmSearch {
		// Find Virtual Machine using UUID
		for ref, obj := range ctx.Map.objects {
			vm, ok := asVirtualMachineMO(obj)
			if !ok {
				continue
			}
			if req.InstanceUuid != nil && *req.InstanceUuid {
				if vm.Config.InstanceUuid == req.Uuid {
					body.Res.Returnval = &ref
					break
				}
			} else {
				if vm.Config.Uuid == req.Uuid {
					body.Res.Returnval = &ref
					break
				}
			}
		}
	} else {
		// Find Host System using UUID
		for ref, obj := range ctx.Map.objects {
			host, ok := asHostSystemMO(obj)
			if !ok {
				continue
			}
			if host.Summary.Hardware.Uuid == req.Uuid {
				body.Res.Returnval = &ref
				break
			}
		}
	}

	return body
}

func (s *SearchIndex) FindAllByUuid(ctx *Context, req *types.FindAllByUuid) soap.HasFault {
	body := &methods.FindAllByUuidBody{Res: new(types.FindAllByUuidResponse)}

	ctx.Map.m.Lock()
	defer ctx.Map.m.Unlock()

	if req.VmSearch {
		// Find Virtual Machines using UUID
		for ref, obj := range ctx.Map.objects {
			vm, ok := asVirtualMachineMO(obj)
			if !ok {
				continue
			}
			if req.InstanceUuid != nil && *req.InstanceUuid {
				if vm.Config.InstanceUuid == req.Uuid {
					body.Res.Returnval = append(body.Res.Returnval, ref)
				}
			} else {
				if vm.Config.Uuid == req.Uuid {
					body.Res.Returnval = append(body.Res.Returnval, ref)
				}
			}
		}
	} else {
		// Find Host System using UUID
		for ref, obj := range ctx.Map.objects {
			host, ok := asHostSystemMO(obj)
			if !ok {
				continue
			}
			if host.Summary.Hardware.Uuid == req.Uuid {
				body.Res.Returnval = append(body.Res.Returnval, ref)
			}
		}
	}

	return body
}

func (s *SearchIndex) FindByDnsName(ctx *Context, req *types.FindByDnsName) soap.HasFault {
	body := &methods.FindByDnsNameBody{Res: new(types.FindByDnsNameResponse)}

	all := types.FindAllByDnsName(*req)

	switch r := s.FindAllByDnsName(ctx, &all).(type) {
	case *methods.FindAllByDnsNameBody:
		if len(r.Res.Returnval) > 0 {
			body.Res.Returnval = &r.Res.Returnval[0]
		}
	default:
		// no need until FindAllByDnsName below returns a Fault
	}

	return body
}

func (s *SearchIndex) FindAllByDnsName(ctx *Context, req *types.FindAllByDnsName) soap.HasFault {
	body := &methods.FindAllByDnsNameBody{Res: new(types.FindAllByDnsNameResponse)}

	ctx.Map.m.Lock()
	defer ctx.Map.m.Unlock()

	if req.VmSearch {
		// Find Virtual Machine using DNS name
		for ref, obj := range ctx.Map.objects {
			vm, ok := asVirtualMachineMO(obj)
			if !ok {
				continue
			}
			if vm.Guest.HostName == req.DnsName {
				body.Res.Returnval = append(body.Res.Returnval, ref)
			}
		}
	} else {
		// Find Host System using DNS name
		for ref, obj := range ctx.Map.objects {
			host, ok := asHostSystemMO(obj)
			if !ok {
				continue
			}
			for _, net := range host.Config.Network.NetStackInstance {
				if net.DnsConfig.GetHostDnsConfig().HostName == req.DnsName {
					body.Res.Returnval = append(body.Res.Returnval, ref)
				}
			}
		}
	}

	return body
}

func (s *SearchIndex) FindByIp(ctx *Context, req *types.FindByIp) soap.HasFault {
	body := &methods.FindByIpBody{Res: new(types.FindByIpResponse)}

	all := types.FindAllByIp(*req)

	switch r := s.FindAllByIp(ctx, &all).(type) {
	case *methods.FindAllByIpBody:
		if len(r.Res.Returnval) > 0 {
			body.Res.Returnval = &r.Res.Returnval[0]
		}
	default:
		// no need until FindAllByIp below returns a Fault
	}

	return body
}

func (s *SearchIndex) FindAllByIp(ctx *Context, req *types.FindAllByIp) soap.HasFault {
	body := &methods.FindAllByIpBody{Res: new(types.FindAllByIpResponse)}

	ctx.Map.m.Lock()
	defer ctx.Map.m.Unlock()

	if req.VmSearch {
		// Find Virtual Machine using IP
		for ref, obj := range ctx.Map.objects {
			vm, ok := asVirtualMachineMO(obj)
			if !ok {
				continue
			}
			if vm.Guest.IpAddress == req.Ip {
				body.Res.Returnval = append(body.Res.Returnval, ref)
			}
		}
	} else {
		// Find Host System using IP
		for ref, obj := range ctx.Map.objects {
			host, ok := asHostSystemMO(obj)
			if !ok {
				continue
			}
			for _, net := range host.Config.Network.Vnic {
				if net.Spec.Ip.IpAddress == req.Ip {
					body.Res.Returnval = append(body.Res.Returnval, ref)
				}
			}
		}
	}

	return body
}
