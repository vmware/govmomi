// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"strings"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type EnvironmentBrowser struct {
	mo.EnvironmentBrowser

	QueryConfigTargetResponse           types.QueryConfigTargetResponse
	QueryConfigOptionResponse           types.QueryConfigOptionResponse
	QueryConfigOptionDescriptorResponse types.QueryConfigOptionDescriptorResponse
	QueryTargetCapabilitiesResponse     types.QueryTargetCapabilitiesResponse
}

func newEnvironmentBrowser(
	ctx *Context,
	hostRefs ...types.ManagedObjectReference) *types.ManagedObjectReference {

	env := new(EnvironmentBrowser)
	env.initDescriptorReturnVal(ctx, hostRefs...)
	ctx.Map.Put(env)
	return &env.Self
}

func (b *EnvironmentBrowser) addHost(
	ctx *Context, hostRef types.ManagedObjectReference) {

	// Get a set of unique hosts.
	hostSet := map[types.ManagedObjectReference]struct{}{
		hostRef: {},
	}
	for i := range b.QueryConfigOptionDescriptorResponse.Returnval {
		cod := b.QueryConfigOptionDescriptorResponse.Returnval[i]
		for j := range cod.Host {
			if _, ok := hostSet[cod.Host[j]]; !ok {
				hostSet[cod.Host[j]] = struct{}{}
			}
		}
	}

	// Get a list of unique hosts.
	var hostRefs []types.ManagedObjectReference
	for ref := range hostSet {
		hostRefs = append(hostRefs, ref)
	}

	// Clear the descriptor's return val.
	b.QueryConfigOptionDescriptorResponse.Returnval = nil

	b.initDescriptorReturnVal(ctx, hostRefs...)
}

func (b *EnvironmentBrowser) initDescriptorReturnVal(
	ctx *Context, hostRefs ...types.ManagedObjectReference) {

	// Get the max supported hardware version for this list of hosts.
	var maxHardwareVersion types.HardwareVersion
	maxHardwareVersionForHost := map[types.ManagedObjectReference]types.HardwareVersion{}
	for j := range hostRefs {
		ref := hostRefs[j]
		ctx.WithLock(ref, func() {
			host := ctx.Map.Get(ref).(*HostSystem)
			hostVersion := types.MustParseESXiVersion(host.Config.Product.Version)
			hostHardwareVersion := hostVersion.HardwareVersion()
			maxHardwareVersionForHost[ref] = hostHardwareVersion
			if !maxHardwareVersion.IsValid() {
				maxHardwareVersion = hostHardwareVersion
				return
			}
			if hostHardwareVersion > maxHardwareVersion {
				maxHardwareVersion = hostHardwareVersion
			}
		})
	}

	if !maxHardwareVersion.IsValid() {
		return
	}

	hardwareVersions := types.GetHardwareVersions()
	for i := range hardwareVersions {
		hv := hardwareVersions[i]
		dco := hv == maxHardwareVersion
		cod := types.VirtualMachineConfigOptionDescriptor{
			Key:                 hv.String(),
			Description:         hv.String(),
			DefaultConfigOption: types.NewBool(dco),
			CreateSupported:     types.NewBool(true),
			RunSupported:        types.NewBool(true),
			UpgradeSupported:    types.NewBool(true),
		}
		for hostRef, hostVer := range maxHardwareVersionForHost {
			if hostVer >= hv {
				cod.Host = append(cod.Host, hostRef)
			}
		}

		b.QueryConfigOptionDescriptorResponse.Returnval = append(
			b.QueryConfigOptionDescriptorResponse.Returnval, cod)

		if dco {
			break
		}
	}
}

func (b *EnvironmentBrowser) hosts(ctx *Context) []types.ManagedObjectReference {
	ctx.Map.m.Lock()
	defer ctx.Map.m.Unlock()
	for _, obj := range ctx.Map.objects {
		switch e := obj.(type) {
		case *mo.ComputeResource:
			if b.Self == *e.EnvironmentBrowser {
				return e.Host
			}
		case *ClusterComputeResource:
			if b.Self == *e.EnvironmentBrowser {
				return e.Host
			}
		}
	}
	return nil
}

func (b *EnvironmentBrowser) QueryConfigOption(req *types.QueryConfigOption) soap.HasFault {
	body := new(methods.QueryConfigOptionBody)

	opt := b.QueryConfigOptionResponse.Returnval
	if opt == nil {
		opt = &types.VirtualMachineConfigOption{
			Version:       esx.HardwareVersion,
			DefaultDevice: esx.VirtualDevice,
		}
	}

	body.Res = &types.QueryConfigOptionResponse{
		Returnval: opt,
	}

	return body
}

func guestFamily(id string) string {
	// TODO: We could capture the entire GuestOsDescriptor list from EnvironmentBrowser,
	// but it is a ton of data.. this should be good enough for now.
	switch {
	case strings.HasPrefix(id, "win"):
		return string(types.VirtualMachineGuestOsFamilyWindowsGuest)
	case strings.HasPrefix(id, "darwin"):
		return string(types.VirtualMachineGuestOsFamilyDarwinGuestFamily)
	default:
		return string(types.VirtualMachineGuestOsFamilyLinuxGuest)
	}
}

func (b *EnvironmentBrowser) QueryConfigOptionEx(req *types.QueryConfigOptionEx) soap.HasFault {
	body := new(methods.QueryConfigOptionExBody)

	opt := b.QueryConfigOptionResponse.Returnval
	if opt == nil {
		opt = &types.VirtualMachineConfigOption{
			Version:       esx.HardwareVersion,
			DefaultDevice: esx.VirtualDevice,
		}
	}

	if req.Spec != nil {
		// From the SDK QueryConfigOptionEx doc:
		// "If guestId is nonempty, the guestOSDescriptor array of the config option is filtered to match against the guest IDs in the spec.
		//  If there is no match, the whole list is returned."
		for _, id := range req.Spec.GuestId {
			for _, gid := range GuestID {
				if string(gid) == id {
					opt.GuestOSDescriptor = []types.GuestOsDescriptor{{
						Id:     id,
						Family: guestFamily(id),
					}}

					break
				}
			}
		}
	}

	if len(opt.GuestOSDescriptor) == 0 {
		for i := range GuestID {
			id := string(GuestID[i])
			opt.GuestOSDescriptor = append(opt.GuestOSDescriptor, types.GuestOsDescriptor{
				Id:     id,
				Family: guestFamily(id),
			})
		}
	}

	body.Res = &types.QueryConfigOptionExResponse{
		Returnval: opt,
	}

	return body
}

func (b *EnvironmentBrowser) QueryConfigOptionDescriptor(ctx *Context, req *types.QueryConfigOptionDescriptor) soap.HasFault {
	body := &methods.QueryConfigOptionDescriptorBody{
		Res: &types.QueryConfigOptionDescriptorResponse{
			Returnval: b.QueryConfigOptionDescriptorResponse.Returnval,
		},
	}

	return body
}

func (b *EnvironmentBrowser) QueryConfigTarget(ctx *Context, req *types.QueryConfigTarget) soap.HasFault {
	body := &methods.QueryConfigTargetBody{
		Res: &types.QueryConfigTargetResponse{
			Returnval: b.QueryConfigTargetResponse.Returnval,
		},
	}

	if body.Res.Returnval != nil {
		return body
	}

	target := &types.ConfigTarget{
		SmcPresent: types.NewBool(false),
	}
	body.Res.Returnval = target

	var hosts []types.ManagedObjectReference
	if req.Host == nil {
		hosts = b.hosts(ctx)
	} else {
		hosts = append(hosts, *req.Host)
	}

	seen := make(map[types.ManagedObjectReference]bool)

	for i := range hosts {
		host := ctx.Map.Get(hosts[i]).(*HostSystem)
		target.NumCpus += int32(host.Summary.Hardware.NumCpuPkgs)
		target.NumCpuCores += int32(host.Summary.Hardware.NumCpuCores)
		target.NumNumaNodes++

		for _, ref := range host.Datastore {
			if seen[ref] {
				continue
			}
			seen[ref] = true

			ds := ctx.Map.Get(ref).(*Datastore)
			target.Datastore = append(target.Datastore, types.VirtualMachineDatastoreInfo{
				VirtualMachineTargetInfo: types.VirtualMachineTargetInfo{
					Name: ds.Name,
				},
				Datastore:       ds.Summary,
				Capability:      ds.Capability,
				Mode:            string(types.HostMountModeReadWrite),
				VStorageSupport: string(types.FileSystemMountInfoVStorageSupportStatusVStorageUnsupported),
			})
		}

		for _, ref := range host.Network {
			if seen[ref] {
				continue
			}
			seen[ref] = true

			switch n := ctx.Map.Get(ref).(type) {
			case *mo.Network:
				target.Network = append(target.Network, types.VirtualMachineNetworkInfo{
					VirtualMachineTargetInfo: types.VirtualMachineTargetInfo{
						Name: n.Name,
					},
					Network: n.Summary.GetNetworkSummary(),
				})
			case *DistributedVirtualPortgroup:
				dvs := ctx.Map.Get(*n.Config.DistributedVirtualSwitch).(*DistributedVirtualSwitch)
				target.DistributedVirtualPortgroup = append(target.DistributedVirtualPortgroup, types.DistributedVirtualPortgroupInfo{
					SwitchName:                  dvs.Name,
					SwitchUuid:                  dvs.Uuid,
					PortgroupName:               n.Name,
					PortgroupKey:                n.Key,
					PortgroupType:               n.Config.Type,
					UplinkPortgroup:             false,
					Portgroup:                   n.Self,
					NetworkReservationSupported: types.NewBool(false),
				})
			case *DistributedVirtualSwitch:
				target.DistributedVirtualSwitch = append(target.DistributedVirtualSwitch, types.DistributedVirtualSwitchInfo{
					SwitchName:                  n.Name,
					SwitchUuid:                  n.Uuid,
					DistributedVirtualSwitch:    n.Self,
					NetworkReservationSupported: types.NewBool(false),
				})
			}
		}
	}

	return body
}

func (b *EnvironmentBrowser) QueryTargetCapabilities(ctx *Context, req *types.QueryTargetCapabilities) soap.HasFault {
	body := &methods.QueryTargetCapabilitiesBody{
		Res: &types.QueryTargetCapabilitiesResponse{
			Returnval: b.QueryTargetCapabilitiesResponse.Returnval,
		},
	}

	if body.Res.Returnval != nil {
		return body
	}

	body.Res.Returnval = &types.HostCapability{
		VmotionSupported:         true,
		MaintenanceModeSupported: true,
	}

	return body
}
