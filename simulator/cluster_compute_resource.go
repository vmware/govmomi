// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"log"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type ClusterComputeResource struct {
	mo.ClusterComputeResource

	ruleKey int32
}

func (c *ClusterComputeResource) RenameTask(ctx *Context, req *types.Rename_Task) soap.HasFault {
	return RenameTask(ctx, c, req)
}

type addHost struct {
	*ClusterComputeResource

	req *types.AddHost_Task
}

func (add *addHost) Run(task *Task) (types.AnyType, types.BaseMethodFault) {
	spec := add.req.Spec

	if spec.HostName == "" {
		return nil, &types.NoHost{}
	}

	cr := add.ClusterComputeResource
	template := esx.HostSystem

	if h := task.ctx.Map.FindByName(spec.UserName, cr.Host); h != nil {
		// "clone" an existing host from the inventory
		template = h.(*HostSystem).HostSystem
		template.Vm = nil
	} else {
		template.Network = cr.Network[:1] // VM Network
	}

	host := NewHostSystem(task.ctx, template)
	host.configure(task.ctx, spec, add.req.AsConnected)

	task.ctx.Map.PutEntity(cr, task.ctx.Map.NewEntity(host))
	host.Summary.Host = &host.Self
	host.Config.Host = host.Self

	task.ctx.Map.WithLock(task.ctx, *cr.EnvironmentBrowser, func() {
		eb := task.ctx.Map.Get(*cr.EnvironmentBrowser).(*EnvironmentBrowser)
		eb.addHost(task.ctx, host.Self)
	})

	cr.Host = append(cr.Host, host.Reference())
	addComputeResource(cr.Summary.GetComputeResourceSummary(), host)

	if cr.vsanIsEnabled() {
		cr.addStorageHost(task.ctx, host.Self)
	}

	return host.Reference(), nil
}

func (c *ClusterComputeResource) AddHostTask(ctx *Context, add *types.AddHost_Task) soap.HasFault {
	return &methods.AddHost_TaskBody{
		Res: &types.AddHost_TaskResponse{
			Returnval: NewTask(&addHost{c, add}).Run(ctx),
		},
	}
}

func (c *ClusterComputeResource) vsanIsEnabled() bool {
	if cfg := c.ConfigurationEx.(*types.ClusterConfigInfoEx).VsanConfigInfo; cfg != nil {
		return isTrue(cfg.Enabled)
	}
	return false
}

func (c *ClusterComputeResource) update(_ *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	if cspec.DasConfig != nil {
		if val := cspec.DasConfig.Enabled; val != nil {
			cfg.DasConfig.Enabled = val
		}
		if val := cspec.DasConfig.AdmissionControlEnabled; val != nil {
			cfg.DasConfig.AdmissionControlEnabled = val
		}
	}
	if cspec.DrsConfig != nil {
		if val := cspec.DrsConfig.Enabled; val != nil {
			cfg.DrsConfig.Enabled = val
		}
		if val := cspec.DrsConfig.DefaultVmBehavior; val != "" {
			cfg.DrsConfig.DefaultVmBehavior = val
		}
	}

	return nil
}

func (c *ClusterComputeResource) updateRules(_ *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	for _, spec := range cspec.RulesSpec {
		var i int
		exists := false

		match := func(info types.BaseClusterRuleInfo) bool {
			return info.GetClusterRuleInfo().Name == spec.Info.GetClusterRuleInfo().Name
		}

		if spec.Operation == types.ArrayUpdateOperationRemove {
			match = func(rule types.BaseClusterRuleInfo) bool {
				return rule.GetClusterRuleInfo().Key == spec.ArrayUpdateSpec.RemoveKey.(int32)
			}
		}

		for i = range cfg.Rule {
			if match(cfg.Rule[i].GetClusterRuleInfo()) {
				exists = true
				break
			}
		}

		switch spec.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			info := spec.Info.GetClusterRuleInfo()
			info.Key = atomic.AddInt32(&c.ruleKey, 1)
			info.RuleUuid = uuid.New().String()
			cfg.Rule = append(cfg.Rule, spec.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.Rule[i] = spec.Info
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.Rule = append(cfg.Rule[:i], cfg.Rule[i+1:]...)
		}
	}

	return nil
}

func (c *ClusterComputeResource) updateGroups(_ *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	for _, spec := range cspec.GroupSpec {
		var i int
		exists := false

		match := func(info types.BaseClusterGroupInfo) bool {
			return info.GetClusterGroupInfo().Name == spec.Info.GetClusterGroupInfo().Name
		}

		if spec.Operation == types.ArrayUpdateOperationRemove {
			match = func(info types.BaseClusterGroupInfo) bool {
				return info.GetClusterGroupInfo().Name == spec.ArrayUpdateSpec.RemoveKey.(string)
			}
		}

		for i = range cfg.Group {
			if match(cfg.Group[i].GetClusterGroupInfo()) {
				exists = true
				break
			}
		}

		switch spec.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			cfg.Group = append(cfg.Group, spec.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.Group[i] = spec.Info
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.Group = append(cfg.Group[:i], cfg.Group[i+1:]...)
		}
	}

	return nil
}

func (c *ClusterComputeResource) updateOverridesDAS(_ *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	for _, spec := range cspec.DasVmConfigSpec {
		var i int
		var key types.ManagedObjectReference
		exists := false

		if spec.Operation == types.ArrayUpdateOperationRemove {
			key = spec.RemoveKey.(types.ManagedObjectReference)
		} else {
			key = spec.Info.Key
		}

		for i = range cfg.DasVmConfig {
			if cfg.DasVmConfig[i].Key == key {
				exists = true
				break
			}
		}

		switch spec.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			cfg.DasVmConfig = append(cfg.DasVmConfig, *spec.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			src := spec.Info.DasSettings
			if src == nil {
				return new(types.InvalidArgument)
			}
			dst := cfg.DasVmConfig[i].DasSettings
			if src.RestartPriority != "" {
				dst.RestartPriority = src.RestartPriority
			}
			if src.RestartPriorityTimeout != 0 {
				dst.RestartPriorityTimeout = src.RestartPriorityTimeout
			}
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.DasVmConfig = append(cfg.DasVmConfig[:i], cfg.DasVmConfig[i+1:]...)
		}
	}

	return nil
}

func (c *ClusterComputeResource) updateOverridesDRS(_ *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	for _, spec := range cspec.DrsVmConfigSpec {
		var i int
		var key types.ManagedObjectReference
		exists := false

		if spec.Operation == types.ArrayUpdateOperationRemove {
			key = spec.RemoveKey.(types.ManagedObjectReference)
		} else {
			key = spec.Info.Key
		}

		for i = range cfg.DrsVmConfig {
			if cfg.DrsVmConfig[i].Key == key {
				exists = true
				break
			}
		}

		switch spec.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			cfg.DrsVmConfig = append(cfg.DrsVmConfig, *spec.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			if spec.Info.Enabled != nil {
				cfg.DrsVmConfig[i].Enabled = spec.Info.Enabled
			}
			if spec.Info.Behavior != "" {
				cfg.DrsVmConfig[i].Behavior = spec.Info.Behavior
			}
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.DrsVmConfig = append(cfg.DrsVmConfig[:i], cfg.DrsVmConfig[i+1:]...)
		}
	}

	return nil
}

func (c *ClusterComputeResource) updateOverridesVmOrchestration(_ *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	for _, spec := range cspec.VmOrchestrationSpec {
		var i int
		var key types.ManagedObjectReference
		exists := false

		if spec.Operation == types.ArrayUpdateOperationRemove {
			key = spec.RemoveKey.(types.ManagedObjectReference)
		} else {
			key = spec.Info.Vm
		}

		for i = range cfg.VmOrchestration {
			if cfg.VmOrchestration[i].Vm == key {
				exists = true
				break
			}
		}

		switch spec.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			cfg.VmOrchestration = append(cfg.VmOrchestration, *spec.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			if spec.Info.VmReadiness.ReadyCondition != "" {
				cfg.VmOrchestration[i].VmReadiness.ReadyCondition = spec.Info.VmReadiness.ReadyCondition
			}
			if spec.Info.VmReadiness.PostReadyDelay != 0 {
				cfg.VmOrchestration[i].VmReadiness.PostReadyDelay = spec.Info.VmReadiness.PostReadyDelay
			}
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			cfg.VmOrchestration = append(cfg.VmOrchestration[:i], cfg.VmOrchestration[i+1:]...)
		}
	}

	return nil
}

func (c *ClusterComputeResource) addStorageHost(ctx *Context, ref types.ManagedObjectReference) types.BaseMethodFault {
	ds := ctx.Map.Get(ref).(*HostSystem).ConfigManager.DatastoreSystem
	hds := ctx.Map.Get(*ds).(*HostDatastoreSystem)
	return hds.createVsanDatastore(ctx)
}

func (c *ClusterComputeResource) updateVSAN(ctx *Context, cfg *types.ClusterConfigInfoEx, cspec *types.ClusterConfigSpecEx) types.BaseMethodFault {
	if cspec.VsanConfig == nil {
		return nil
	}

	if cfg.VsanConfigInfo == nil {
		cfg.VsanConfigInfo = cspec.VsanConfig
		if cfg.VsanConfigInfo.DefaultConfig == nil {
			cfg.VsanConfigInfo.DefaultConfig = new(types.VsanClusterConfigInfoHostDefaultInfo)
		}
	} else {
		if cspec.VsanConfig.Enabled != nil {
			cfg.VsanConfigInfo.Enabled = cspec.VsanConfig.Enabled
		}
	}

	if cfg.VsanConfigInfo.DefaultConfig.Uuid == "" {
		cfg.VsanConfigInfo.DefaultConfig.Uuid = uuid.NewString()
	}

	if isTrue(cfg.VsanConfigInfo.Enabled) {
		for _, ref := range c.Host {
			if err := c.addStorageHost(ctx, ref); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *ClusterComputeResource) ReconfigureComputeResourceTask(ctx *Context, req *types.ReconfigureComputeResource_Task) soap.HasFault {
	task := CreateTask(c, "reconfigureCluster", func(*Task) (types.AnyType, types.BaseMethodFault) {
		spec, ok := req.Spec.(*types.ClusterConfigSpecEx)
		if !ok {
			return nil, new(types.InvalidArgument)
		}

		updates := []func(*Context, *types.ClusterConfigInfoEx, *types.ClusterConfigSpecEx) types.BaseMethodFault{
			c.update,
			c.updateRules,
			c.updateGroups,
			c.updateOverridesDAS,
			c.updateOverridesDRS,
			c.updateOverridesVmOrchestration,
			c.updateVSAN,
		}

		for _, update := range updates {
			if err := update(ctx, c.ConfigurationEx.(*types.ClusterConfigInfoEx), spec); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return &methods.ReconfigureComputeResource_TaskBody{
		Res: &types.ReconfigureComputeResource_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (c *ClusterComputeResource) MoveIntoTask(ctx *Context, req *types.MoveInto_Task) soap.HasFault {
	task := CreateTask(c, "moveInto", func(*Task) (types.AnyType, types.BaseMethodFault) {
		for _, ref := range req.Host {
			host := ctx.Map.Get(ref).(*HostSystem)

			if *host.Parent == c.Self {
				return nil, new(types.DuplicateName) // host already in this cluster
			}

			switch parent := ctx.Map.Get(*host.Parent).(type) {
			case *ClusterComputeResource:
				if !host.Runtime.InMaintenanceMode {
					return nil, new(types.InvalidState)
				}

				RemoveReference(&parent.Host, ref)
			case *mo.ComputeResource:
				ctx.Map.Remove(ctx, parent.Self)
			}

			c.Host = append(c.Host, ref)
			host.Parent = &c.Self
		}

		return nil, nil
	})

	return &methods.MoveInto_TaskBody{
		Res: &types.MoveInto_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (c *ClusterComputeResource) PlaceVm(ctx *Context, req *types.PlaceVm) soap.HasFault {
	body := new(methods.PlaceVmBody)

	if len(c.Host) == 0 {
		body.Fault_ = Fault("", new(types.InvalidState))
		return body
	}

	res := types.ClusterRecommendation{
		Key:        "1",
		Type:       "V1",
		Time:       time.Now(),
		Rating:     1,
		Reason:     string(types.RecommendationReasonCodeXvmotionPlacement),
		ReasonText: string(types.RecommendationReasonCodeXvmotionPlacement),
		Target:     &c.Self,
	}

	hosts := req.PlacementSpec.Hosts
	if len(hosts) == 0 {
		hosts = c.Host
	}

	datastores := req.PlacementSpec.Datastores
	if len(datastores) == 0 {
		datastores = c.Datastore
	}

	switch types.PlacementSpecPlacementType(req.PlacementSpec.PlacementType) {
	case types.PlacementSpecPlacementTypeClone, types.PlacementSpecPlacementTypeCreate:
		spec := &types.VirtualMachineRelocateSpec{
			Datastore: &datastores[rand.Intn(len(c.Datastore))],
			Host:      &hosts[rand.Intn(len(c.Host))],
			Pool:      c.ResourcePool,
		}
		res.Action = append(res.Action, &types.PlacementAction{
			Vm:           req.PlacementSpec.Vm,
			TargetHost:   spec.Host,
			RelocateSpec: spec,
		})
	case types.PlacementSpecPlacementTypeReconfigure:
		// Validate input.
		if req.PlacementSpec.ConfigSpec == nil {
			body.Fault_ = Fault("", &types.InvalidArgument{
				InvalidProperty: "PlacementSpec.configSpec",
			})
			return body
		}

		// Update PlacementResult.
		vmObj := ctx.Map.Get(*req.PlacementSpec.Vm).(*VirtualMachine)
		spec := &types.VirtualMachineRelocateSpec{
			Datastore:    &vmObj.Datastore[0],
			Host:         vmObj.Runtime.Host,
			Pool:         vmObj.ResourcePool,
			DiskMoveType: string(types.VirtualMachineRelocateDiskMoveOptionsMoveAllDiskBackingsAndAllowSharing),
		}
		res.Action = append(res.Action, &types.PlacementAction{
			Vm:           req.PlacementSpec.Vm,
			TargetHost:   spec.Host,
			RelocateSpec: spec,
		})
	case types.PlacementSpecPlacementTypeRelocate:
		// Validate fields of req.PlacementSpec, if explicitly provided.
		if !validatePlacementSpecForPlaceVmRelocate(ctx, req, body) {
			return body
		}

		// After validating req.PlacementSpec, we must have a valid req.PlacementSpec.Vm.
		vmObj := ctx.Map.Get(*req.PlacementSpec.Vm).(*VirtualMachine)

		// Populate RelocateSpec's common fields, if not explicitly provided.
		populateRelocateSpecForPlaceVmRelocate(&req.PlacementSpec.RelocateSpec, vmObj)

		// Update PlacementResult.
		res.Action = append(res.Action, &types.PlacementAction{
			Vm:           req.PlacementSpec.Vm,
			TargetHost:   req.PlacementSpec.RelocateSpec.Host,
			RelocateSpec: req.PlacementSpec.RelocateSpec,
		})
	default:
		log.Printf("unsupported placement type: %s", req.PlacementSpec.PlacementType)
		body.Fault_ = Fault("", new(types.NotSupported))
		return body
	}

	body.Res = &types.PlaceVmResponse{
		Returnval: types.PlacementResult{
			Recommendations: []types.ClusterRecommendation{res},
		},
	}

	return body
}

// validatePlacementSpecForPlaceVmRelocate validates the fields of req.PlacementSpec for a relocate placement type.
// Returns true if the fields are valid, false otherwise.
func validatePlacementSpecForPlaceVmRelocate(ctx *Context, req *types.PlaceVm, body *methods.PlaceVmBody) bool {
	if req.PlacementSpec.Vm == nil {
		body.Fault_ = Fault("", &types.InvalidArgument{
			InvalidProperty: "PlacementSpec",
		})
		return false
	}

	// Oddly when the VM is not found, PlaceVm complains about configSpec being invalid, despite this being
	// a relocate placement type. Possibly due to treating the missing VM as a create placement type
	// internally, which requires the configSpec to be present.
	vmObj, exist := ctx.Map.Get(*req.PlacementSpec.Vm).(*VirtualMachine)
	if !exist {
		body.Fault_ = Fault("", &types.InvalidArgument{
			InvalidProperty: "PlacementSpec.configSpec",
		})
		return false
	}

	return validateRelocateSpecForPlaceVmRelocate(ctx, req.PlacementSpec.RelocateSpec, body, vmObj)
}

// validateRelocateSpecForPlaceVmRelocate validates the fields of req.PlacementSpec.RelocateSpec for a relocate
// placement type. Returns true if the fields are valid, false otherwise.
func validateRelocateSpecForPlaceVmRelocate(ctx *Context, spec *types.VirtualMachineRelocateSpec, body *methods.PlaceVmBody, vmObj *VirtualMachine) bool {
	if spec == nil {
		// An empty relocate spec is valid, as it will be populated with default values.
		return true
	}

	if spec.Host != nil {
		if _, exist := ctx.Map.Get(*spec.Host).(*HostSystem); !exist {
			body.Fault_ = Fault("", &types.ManagedObjectNotFound{
				Obj: *spec.Host,
			})
			return false
		}
	}

	if spec.Datastore != nil {
		if _, exist := ctx.Map.Get(*spec.Datastore).(*Datastore); !exist {
			body.Fault_ = Fault("", &types.ManagedObjectNotFound{
				Obj: *spec.Datastore,
			})
			return false
		}
	}

	if spec.Pool != nil {
		if _, exist := ctx.Map.Get(*spec.Pool).(*ResourcePool); !exist {
			body.Fault_ = Fault("", &types.ManagedObjectNotFound{
				Obj: *spec.Pool,
			})
			return false
		}
	}

	if spec.Disk != nil {
		deviceList := object.VirtualDeviceList(vmObj.Config.Hardware.Device)
		vdiskList := deviceList.SelectByType(&types.VirtualDisk{})
		for _, disk := range spec.Disk {
			var diskFound bool
			for _, vdisk := range vdiskList {
				if disk.DiskId == vdisk.GetVirtualDevice().Key {
					diskFound = true
					break
				}
			}
			if !diskFound {
				body.Fault_ = Fault("", &types.InvalidArgument{
					InvalidProperty: "PlacementSpec.vm",
				})
				return false
			}

			// Unlike a non-existing spec.Datastore that throws ManagedObjectNotFound, a non-existing disk.Datastore
			// throws InvalidArgument.
			if _, exist := ctx.Map.Get(disk.Datastore).(*Datastore); !exist {
				body.Fault_ = Fault("", &types.InvalidArgument{
					InvalidProperty: "RelocateSpec",
				})
				return false
			}
		}
	}

	return true
}

// populateRelocateSpecForPlaceVmRelocate populates the fields of req.PlacementSpec.RelocateSpec for a relocate
// placement type, if not explicitly provided.
func populateRelocateSpecForPlaceVmRelocate(specPtr **types.VirtualMachineRelocateSpec, vmObj *VirtualMachine) {
	if *specPtr == nil {
		*specPtr = &types.VirtualMachineRelocateSpec{}
	}

	spec := *specPtr

	if spec.DiskMoveType == "" {
		spec.DiskMoveType = string(types.VirtualMachineRelocateDiskMoveOptionsMoveAllDiskBackingsAndDisallowSharing)
	}

	if spec.Datastore == nil {
		spec.Datastore = &vmObj.Datastore[0]
	}

	if spec.Host == nil {
		spec.Host = vmObj.Runtime.Host
	}

	if spec.Pool == nil {
		spec.Pool = vmObj.ResourcePool
	}

	if spec.Disk == nil {
		deviceList := object.VirtualDeviceList(vmObj.Config.Hardware.Device)
		for _, vdisk := range deviceList.SelectByType(&types.VirtualDisk{}) {
			spec.Disk = append(spec.Disk, types.VirtualMachineRelocateSpecDiskLocator{
				DiskId:       vdisk.GetVirtualDevice().Key,
				Datastore:    *spec.Datastore,
				DiskMoveType: spec.DiskMoveType,
			})
		}
	}
}

func CreateClusterComputeResource(ctx *Context, f *Folder, name string, spec types.ClusterConfigSpecEx) (*ClusterComputeResource, types.BaseMethodFault) {
	if e := ctx.Map.FindByName(name, f.ChildEntity); e != nil {
		return nil, &types.DuplicateName{
			Name:   e.Entity().Name,
			Object: e.Reference(),
		}
	}

	cluster := &ClusterComputeResource{}
	cluster.EnvironmentBrowser = newEnvironmentBrowser(ctx)
	cluster.Name = name
	cluster.Network = ctx.Map.getEntityDatacenter(f).defaultNetwork()
	cluster.Summary = &types.ClusterComputeResourceSummary{
		UsageSummary: new(types.ClusterUsageSummary),
	}

	config := &types.ClusterConfigInfoEx{}
	cluster.ConfigurationEx = config

	config.VmSwapPlacement = string(types.VirtualMachineConfigInfoSwapPlacementTypeVmDirectory)
	config.DrsConfig.Enabled = types.NewBool(true)

	pool := NewResourcePool(ctx)
	ctx.Map.PutEntity(cluster, ctx.Map.NewEntity(pool))
	cluster.ResourcePool = &pool.Self

	folderPutChild(ctx, &f.Folder, cluster)
	pool.Owner = cluster.Self

	return cluster, nil
}
