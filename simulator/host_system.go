// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	hostPortUnique = os.Getenv("VCSIM_HOST_PORT_UNIQUE") == "true"

	globalLock sync.Mutex
	// globalHostCount is used to construct unique hostnames. Should be consumed under globalLock.
	globalHostCount = 0
)

type HostSystem struct {
	mo.HostSystem

	sh  *simHost
	mme *ManagedMethodExecuter
	dtm *DynamicTypeManager

	types.QueryTpmAttestationReportResponse
}

func asHostSystemMO(obj mo.Reference) (*mo.HostSystem, bool) {
	h, ok := getManagedObject(obj).Addr().Interface().(*mo.HostSystem)
	return h, ok
}

func NewHostSystem(ctx *Context, host mo.HostSystem) *HostSystem {
	if hostPortUnique { // configure unique port for each host
		port := &esx.HostSystem.Summary.Config.Port
		*port++
		host.Summary.Config.Port = *port
	}

	now := time.Now()

	hs := &HostSystem{
		HostSystem: host,
	}

	hs.Name = hs.Summary.Config.Name
	hs.Summary.Runtime = &hs.Runtime
	hs.Summary.Runtime.BootTime = &now

	// shallow copy Summary.Hardware, as each host will be assigned its own .Uuid
	hardware := *host.Summary.Hardware
	hs.Summary.Hardware = &hardware

	if hs.Hardware == nil {
		// shallow copy Hardware, as each host will be assigned its own .Uuid
		info := *esx.HostHardwareInfo
		hs.Hardware = &info
	}
	if hs.Capability == nil {
		capability := *esx.HostCapability
		hs.Capability = &capability
	}

	cfg := new(types.HostConfigInfo)
	deepCopy(hs.Config, cfg)
	hs.Config = cfg

	// copy over the reference advanced options so each host can have it's own, allowing hosts to be configured for
	// container backing individually
	deepCopy(esx.AdvancedOptions, &cfg.Option)

	// add a supported option to the AdvancedOption manager
	simOption := types.OptionDef{ElementDescription: types.ElementDescription{Key: advOptContainerBackingImage}}
	// TODO: how do we enter patterns here? Or should we stick to a list in the value?
	// patterns become necessary if we want to enforce correctness on options for RUN.underlay.<pnic> or allow RUN.port.xxx
	hs.Config.OptionDef = append(hs.Config.OptionDef, simOption)

	config := []struct {
		ref **types.ManagedObjectReference
		obj mo.Reference
	}{
		{&hs.ConfigManager.DatastoreSystem, &HostDatastoreSystem{Host: &hs.HostSystem}},
		{&hs.ConfigManager.NetworkSystem, NewHostNetworkSystem(&hs.HostSystem)},
		{&hs.ConfigManager.VirtualNicManager, NewHostVirtualNicManager(&hs.HostSystem)},
		{&hs.ConfigManager.AdvancedOption, NewOptionManager(nil, nil, &hs.Config.Option)},
		{&hs.ConfigManager.FirewallSystem, NewHostFirewallSystem(&hs.HostSystem)},
		{&hs.ConfigManager.StorageSystem, NewHostStorageSystem(&hs.HostSystem)},
		{&hs.ConfigManager.CertificateManager, NewHostCertificateManager(ctx, &hs.HostSystem)},
	}

	for _, c := range config {
		ref := ctx.Map.Put(c.obj).Reference()

		*c.ref = &ref
	}

	return hs
}

func (h *HostSystem) configure(ctx *Context, spec types.HostConnectSpec, connected bool) {
	h.Runtime.ConnectionState = types.HostSystemConnectionStateDisconnected
	if connected {
		h.Runtime.ConnectionState = types.HostSystemConnectionStateConnected
	}

	// lets us construct non-conflicting hostname automatically if omitted
	// does not use the unique port instead to avoid constraints on port, such as >1024

	globalLock.Lock()
	instanceID := globalHostCount
	globalHostCount++
	globalLock.Unlock()

	if spec.HostName == "" {
		spec.HostName = fmt.Sprintf("esx-%d", instanceID)
	} else if net.ParseIP(spec.HostName) != nil {
		h.Config.Network.Vnic[0].Spec.Ip.IpAddress = spec.HostName
	}

	h.Summary.Config.Name = spec.HostName
	h.Name = h.Summary.Config.Name
	id := newUUID(h.Name)
	h.Summary.Hardware.Uuid = id
	h.Hardware.SystemInfo.Uuid = id

	var err error
	h.sh, err = createSimulationHost(ctx, h)
	if err != nil {
		panic("failed to create simulation host and no path to return error: " + err.Error())
	}
}

// configureContainerBacking sets up _this_ host for simulation using a container backing.
// Args:
//
//		image - the container image with which to simulate the host
//		mounts - array of mount info that should be translated into /vmfs/volumes/... mounts backed by container volumes
//	 	networks - names of bridges to use for underlays. Will create a pNIC for each. The first will be treated as the management network.
//
// Restrictions adopted from createSimulationHost:
// * no mock of VLAN connectivity
// * only a single vmknic, used for "the management IP"
// * pNIC connectivity does not directly impact VMs/vmks using it as uplink
//
// The pnics will be named using standard pattern, ie. vmnic0, vmnic1, ...
// This will sanity check the NetConfig for "management" nicType to ensure that it maps through PortGroup->vSwitch->pNIC to vmnic0.
func (h *HostSystem) configureContainerBacking(ctx *Context, image string, mounts []types.HostFileSystemMountInfo, networks ...string) error {
	option := &types.OptionValue{
		Key:   advOptContainerBackingImage,
		Value: image,
	}

	advOpts := ctx.Map.Get(h.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	fault := advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option}}).Fault()
	if fault != nil {
		panic(fault)
	}

	h.Config.FileSystemVolume = nil
	if mounts != nil {
		h.Config.FileSystemVolume = &types.HostFileSystemVolumeInfo{
			VolumeTypeList: []string{"VMFS", "OTHER"},
			MountInfo:      mounts,
		}
	}

	// force at least a management network
	if len(networks) == 0 {
		networks = []string{defaultUnderlayBridgeName}
	}

	// purge pNICs from the template - it makes no sense to keep them for a sim host
	h.Config.Network.Pnic = make([]types.PhysicalNic, len(networks))

	// purge any IPs and MACs associated with existing NetConfigs for the host
	for cfgIdx := range h.Config.VirtualNicManagerInfo.NetConfig {
		config := &h.Config.VirtualNicManagerInfo.NetConfig[cfgIdx]
		for candidateIdx := range config.CandidateVnic {
			candidate := &config.CandidateVnic[candidateIdx]
			candidate.Spec.Ip.IpAddress = "0.0.0.0"
			candidate.Spec.Ip.SubnetMask = "0.0.0.0"
			candidate.Spec.Mac = "00:00:00:00:00:00"
		}
	}

	// The presence of a pNIC is used to indicate connectivity to a specific underlay. We construct an empty pNIC entry and specify the underly via
	// host.ConfigManager.AdvancedOptions. The pNIC will be populated with the MAC (accurate) and IP (divergence - we need to stash it somewhere) for the veth.
	// We create a NetConfig "management" entry for the first pNIC - this will be populated with the IP of the "host" container.

	// create a pNIC for each underlay
	for i, net := range networks {
		name := fmt.Sprintf("vmnic%d", i)

		// we don't have a natural field for annotating which pNIC is connected to which network, so stash it in an adv option.
		option := &types.OptionValue{
			Key:   advOptPrefixPnicToUnderlayPrefix + name,
			Value: net,
		}
		fault = advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option}}).Fault()
		if fault != nil {
			panic(fault)
		}

		h.Config.Network.Pnic[i] = types.PhysicalNic{
			Key:             "key-vim.host.PhysicalNic-" + name,
			Device:          name,
			Pci:             fmt.Sprintf("0000:%2d:00.0", i+1),
			Driver:          "vcsim-bridge",
			DriverVersion:   "1.2.10.0",
			FirmwareVersion: "1.57, 0x80000185",
			LinkSpeed: &types.PhysicalNicLinkInfo{
				SpeedMb: 10000,
				Duplex:  true,
			},
			ValidLinkSpecification: []types.PhysicalNicLinkInfo{
				{
					SpeedMb: 10000,
					Duplex:  true,
				},
			},
			Spec: types.PhysicalNicSpec{
				Ip:                            &types.HostIpConfig{},
				LinkSpeed:                     (*types.PhysicalNicLinkInfo)(nil),
				EnableEnhancedNetworkingStack: types.NewBool(false),
				EnsInterruptEnabled:           types.NewBool(false),
			},
			WakeOnLanSupported: false,
			Mac:                "00:00:00:00:00:00",
			FcoeConfiguration: &types.FcoeConfig{
				PriorityClass: 3,
				SourceMac:     "00:00:00:00:00:00",
				VlanRange: []types.FcoeConfigVlanRange{
					{},
				},
				Capabilities: types.FcoeConfigFcoeCapabilities{},
				FcoeActive:   false,
			},
			VmDirectPathGen2Supported:             types.NewBool(false),
			VmDirectPathGen2SupportedMode:         "",
			ResourcePoolSchedulerAllowed:          types.NewBool(false),
			ResourcePoolSchedulerDisallowedReason: nil,
			AutoNegotiateSupported:                types.NewBool(true),
			EnhancedNetworkingStackSupported:      types.NewBool(false),
			EnsInterruptSupported:                 types.NewBool(false),
			RdmaDevice:                            "",
			DpuId:                                 "",
		}
	}

	// sanity check that everything's hung together sufficiently well
	details, err := h.getNetConfigInterface(ctx, "management")
	if err != nil {
		return err
	}

	if details.uplink == nil || details.uplink.Device != "vmnic0" {
		return fmt.Errorf("Config provided for host %s does not result in a consistent 'management' NetConfig that's bound to 'vmnic0'", h.Name)
	}

	return nil
}

// netConfigDetails is used to packaged up all the related network entities associated with a NetConfig binding
type netConfigDetails struct {
	nicType   string
	netconfig *types.VirtualNicManagerNetConfig
	vmk       *types.HostVirtualNic
	netstack  *types.HostNetStackInstance
	portgroup *types.HostPortGroup
	vswitch   *types.HostVirtualSwitch
	uplink    *types.PhysicalNic
}

// getNetConfigInterface returns the set of constructs active for a given nicType (eg. "management", "vmotion")
// This method is provided because the Config structure held by HostSystem is heavily interconnected but serialized and not cross-linked with pointers.
// As such there's a _lot_ of cross-referencing that needs to be done to navigate.
// The pNIC returned is the uplink associated with the vSwitch for the netconfig
func (h *HostSystem) getNetConfigInterface(ctx *Context, nicType string) (*netConfigDetails, error) {
	details := &netConfigDetails{
		nicType: nicType,
	}

	for i := range h.Config.VirtualNicManagerInfo.NetConfig {
		if h.Config.VirtualNicManagerInfo.NetConfig[i].NicType == nicType {
			details.netconfig = &h.Config.VirtualNicManagerInfo.NetConfig[i]
			break
		}
	}
	if details.netconfig == nil {
		return nil, fmt.Errorf("no matching NetConfig for NicType=%s", nicType)
	}

	if details.netconfig.SelectedVnic == nil {
		return details, nil
	}

	vnicKey := details.netconfig.SelectedVnic[0]
	for i := range details.netconfig.CandidateVnic {
		if details.netconfig.CandidateVnic[i].Key == vnicKey {
			details.vmk = &details.netconfig.CandidateVnic[i]
			break
		}
	}
	if details.vmk == nil {
		panic(fmt.Sprintf("NetConfig for host %s references non-existant vNIC key %s for %s nicType", h.Name, vnicKey, nicType))
	}

	portgroupName := details.vmk.Portgroup
	netstackKey := details.vmk.Spec.NetStackInstanceKey

	for i := range h.Config.Network.NetStackInstance {
		if h.Config.Network.NetStackInstance[i].Key == netstackKey {
			details.netstack = &h.Config.Network.NetStackInstance[i]
			break
		}
	}
	if details.netstack == nil {
		panic(fmt.Sprintf("NetConfig for host %s references non-existant NetStack key %s for %s nicType", h.Name, netstackKey, nicType))
	}

	for i := range h.Config.Network.Portgroup {
		// TODO: confirm correctness of this - seems weird it references the Spec.Name instead of the key like everything else.
		if h.Config.Network.Portgroup[i].Spec.Name == portgroupName {
			details.portgroup = &h.Config.Network.Portgroup[i]
			break
		}
	}
	if details.portgroup == nil {
		panic(fmt.Sprintf("NetConfig for host %s references non-existant PortGroup name %s for %s nicType", h.Name, portgroupName, nicType))
	}

	vswitchKey := details.portgroup.Vswitch
	for i := range h.Config.Network.Vswitch {
		if h.Config.Network.Vswitch[i].Key == vswitchKey {
			details.vswitch = &h.Config.Network.Vswitch[i]
			break
		}
	}
	if details.vswitch == nil {
		panic(fmt.Sprintf("NetConfig for host %s references non-existant vSwitch key %s for %s nicType", h.Name, vswitchKey, nicType))
	}

	if len(details.vswitch.Pnic) != 1 {
		// to change this, look at the Active NIC in the NicTeamingPolicy, but for now not worth it
		panic(fmt.Sprintf("vSwitch %s for host %s has multiple pNICs associated which is not supported.", vswitchKey, h.Name))
	}

	pnicKey := details.vswitch.Pnic[0]
	for i := range h.Config.Network.Pnic {
		if h.Config.Network.Pnic[i].Key == pnicKey {
			details.uplink = &h.Config.Network.Pnic[i]
			break
		}
	}
	if details.uplink == nil {
		panic(fmt.Sprintf("NetConfig for host %s references non-existant pNIC key %s for %s nicType", h.Name, pnicKey, nicType))
	}

	return details, nil
}

func (h *HostSystem) event(ctx *Context) types.HostEvent {
	return types.HostEvent{
		Event: types.Event{
			Datacenter:      datacenterEventArgument(ctx, h),
			ComputeResource: h.eventArgumentParent(ctx),
			Host:            h.eventArgument(),
		},
	}
}

func (h *HostSystem) eventArgument() *types.HostEventArgument {
	return &types.HostEventArgument{
		Host:                h.Self,
		EntityEventArgument: types.EntityEventArgument{Name: h.Name},
	}
}

func (h *HostSystem) eventArgumentParent(ctx *Context) *types.ComputeResourceEventArgument {
	parent := hostParent(ctx, &h.HostSystem)

	return &types.ComputeResourceEventArgument{
		ComputeResource:     parent.Self,
		EntityEventArgument: types.EntityEventArgument{Name: parent.Name},
	}
}

func hostParent(ctx *Context, host *mo.HostSystem) *mo.ComputeResource {
	switch parent := ctx.Map.Get(*host.Parent).(type) {
	case *mo.ComputeResource:
		return parent
	case *ClusterComputeResource:
		return &parent.ComputeResource
	default:
		return nil
	}
}

func addComputeResource(s *types.ComputeResourceSummary, h *HostSystem) {
	s.TotalCpu += h.Summary.Hardware.CpuMhz
	s.TotalMemory += h.Summary.Hardware.MemorySize
	s.NumCpuCores += h.Summary.Hardware.NumCpuCores
	s.NumCpuThreads += h.Summary.Hardware.NumCpuThreads
	s.EffectiveCpu += h.Summary.Hardware.CpuMhz
	s.EffectiveMemory += h.Summary.Hardware.MemorySize
	s.NumHosts++
	s.NumEffectiveHosts++
	s.OverallStatus = types.ManagedEntityStatusGreen
}

// CreateDefaultESX creates a standalone ESX
// Adds objects of type: Datacenter, Network, ComputeResource, ResourcePool and HostSystem
func CreateDefaultESX(ctx *Context, f *Folder) {
	dc := NewDatacenter(ctx, &f.Folder)

	host := NewHostSystem(ctx, esx.HostSystem)

	summary := new(types.ComputeResourceSummary)
	addComputeResource(summary, host)

	cr := &mo.ComputeResource{
		Summary: summary,
		Network: esx.Datacenter.Network,
	}
	cr.Self = *host.Parent
	cr.Name = host.Name
	cr.Host = append(cr.Host, host.Reference())
	host.Network = cr.Network
	ctx.Map.PutEntity(cr, host)
	cr.EnvironmentBrowser = newEnvironmentBrowser(ctx, host.Reference())

	pool := NewResourcePool(ctx)
	cr.ResourcePool = &pool.Self
	ctx.Map.PutEntity(cr, pool)
	pool.Owner = cr.Self

	folderPutChild(ctx, &ctx.Map.Get(dc.HostFolder).(*Folder).Folder, cr)
}

// CreateStandaloneHost uses esx.HostSystem as a template, applying the given spec
// and creating the ComputeResource parent and ResourcePool sibling.
func CreateStandaloneHost(ctx *Context, f *Folder, spec types.HostConnectSpec) (*HostSystem, types.BaseMethodFault) {
	if spec.HostName == "" {
		return nil, &types.NoHost{}
	}

	template := esx.HostSystem
	network := ctx.Map.getEntityDatacenter(f).defaultNetwork()

	if p := ctx.Map.FindByName(spec.UserName, f.ChildEntity); p != nil {
		cr := p.(*mo.ComputeResource)
		h := ctx.Map.Get(cr.Host[0])
		// "clone" an existing host from the inventory
		template = h.(*HostSystem).HostSystem
		template.Vm = nil
		network = cr.Network
	}

	pool := NewResourcePool(ctx)
	host := NewHostSystem(ctx, template)
	host.configure(ctx, spec, false)

	summary := new(types.ComputeResourceSummary)
	addComputeResource(summary, host)

	cr := &mo.ComputeResource{
		ConfigurationEx: &types.ComputeResourceConfigInfo{
			VmSwapPlacement: string(types.VirtualMachineConfigInfoSwapPlacementTypeVmDirectory),
		},
		Summary: summary,
	}

	ctx.Map.PutEntity(cr, ctx.Map.NewEntity(host))
	cr.EnvironmentBrowser = newEnvironmentBrowser(ctx, host.Reference())

	host.Summary.Host = &host.Self
	host.Config.Host = host.Self

	ctx.Map.PutEntity(cr, ctx.Map.NewEntity(pool))

	cr.Name = host.Name
	cr.Network = network
	cr.Host = append(cr.Host, host.Reference())
	cr.ResourcePool = &pool.Self

	folderPutChild(ctx, &f.Folder, cr)
	pool.Owner = cr.Self
	host.Network = cr.Network

	return host, nil
}

func (h *HostSystem) DestroyTask(ctx *Context, req *types.Destroy_Task) soap.HasFault {
	task := CreateTask(h, "destroy", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		if len(h.Vm) > 0 {
			return nil, &types.ResourceInUse{}
		}

		ctx.postEvent(&types.HostRemovedEvent{HostEvent: h.event(ctx)})

		f := ctx.Map.getEntityParent(h, "Folder").(*Folder)
		folderRemoveChild(ctx, &f.Folder, h.Reference())
		err := h.sh.remove(ctx)

		if err != nil {
			return nil, &types.RuntimeFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						Fault:            &types.SystemErrorFault{Reason: err.Error()},
						LocalizedMessage: err.Error()}}}
		}

		// TODO: should there be events on lifecycle operations as with VMs?

		return nil, nil
	})

	return &methods.Destroy_TaskBody{
		Res: &types.Destroy_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (h *HostSystem) EnterMaintenanceModeTask(ctx *Context, spec *types.EnterMaintenanceMode_Task) soap.HasFault {
	task := CreateTask(h, "enterMaintenanceMode", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		h.Runtime.InMaintenanceMode = true
		return nil, nil
	})

	return &methods.EnterMaintenanceMode_TaskBody{
		Res: &types.EnterMaintenanceMode_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (h *HostSystem) ExitMaintenanceModeTask(ctx *Context, spec *types.ExitMaintenanceMode_Task) soap.HasFault {
	task := CreateTask(h, "exitMaintenanceMode", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		h.Runtime.InMaintenanceMode = false
		return nil, nil
	})

	return &methods.ExitMaintenanceMode_TaskBody{
		Res: &types.ExitMaintenanceMode_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (h *HostSystem) DisconnectHostTask(ctx *Context, spec *types.DisconnectHost_Task) soap.HasFault {
	task := CreateTask(h, "disconnectHost", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		h.Runtime.ConnectionState = types.HostSystemConnectionStateDisconnected
		return nil, nil
	})

	return &methods.DisconnectHost_TaskBody{
		Res: &types.DisconnectHost_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (h *HostSystem) ReconnectHostTask(ctx *Context, spec *types.ReconnectHost_Task) soap.HasFault {
	task := CreateTask(h, "reconnectHost", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		h.Runtime.ConnectionState = types.HostSystemConnectionStateConnected
		return nil, nil
	})

	return &methods.ReconnectHost_TaskBody{
		Res: &types.ReconnectHost_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (s *HostSystem) QueryTpmAttestationReport(ctx *Context, req *types.QueryTpmAttestationReport) soap.HasFault {
	body := new(methods.QueryTpmAttestationReportBody)

	if ctx.Map.IsVPX() {
		body.Res = &s.QueryTpmAttestationReportResponse
	} else {
		body.Fault_ = Fault("", new(types.NotSupported))
	}

	return body
}
