// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vmdk"
)

type VirtualMachine struct {
	mo.VirtualMachine
	DataSets map[string]*DataSet

	log string
	sid int32
	svm *simVM
	uid uuid.UUID
	imc *types.CustomizationSpec
}

func asVirtualMachineMO(obj mo.Reference) (*mo.VirtualMachine, bool) {
	vm, ok := getManagedObject(obj).Addr().Interface().(*mo.VirtualMachine)
	return vm, ok
}

func NewVirtualMachine(ctx *Context, parent types.ManagedObjectReference, spec *types.VirtualMachineConfigSpec) (*VirtualMachine, types.BaseMethodFault) {
	vm := &VirtualMachine{}
	vm.Parent = &parent
	ctx.Map.reference(vm)

	folder := ctx.Map.Get(parent)

	if spec.Name == "" {
		return vm, &types.InvalidVmConfig{Property: "configSpec.name"}
	}

	if spec.Files == nil || spec.Files.VmPathName == "" {
		return vm, &types.InvalidVmConfig{Property: "configSpec.files.vmPathName"}
	}

	rspec := types.DefaultResourceConfigSpec()
	vm.Guest = &types.GuestInfo{}
	vm.Config = &types.VirtualMachineConfigInfo{
		ExtraConfig:        []types.BaseOptionValue{&types.OptionValue{Key: "govcsim", Value: "TRUE"}},
		Tools:              &types.ToolsConfigInfo{},
		MemoryAllocation:   &rspec.MemoryAllocation,
		CpuAllocation:      &rspec.CpuAllocation,
		LatencySensitivity: &types.LatencySensitivity{Level: types.LatencySensitivitySensitivityLevelNormal},
		BootOptions:        &types.VirtualMachineBootOptions{},
		CreateDate:         types.NewTime(time.Now()),
	}
	vm.Layout = &types.VirtualMachineFileLayout{}
	vm.LayoutEx = &types.VirtualMachineFileLayoutEx{
		Timestamp: time.Now(),
	}
	vm.Snapshot = nil // intentionally set to nil until a snapshot is created
	vm.Storage = &types.VirtualMachineStorageInfo{
		Timestamp: time.Now(),
	}
	vm.Summary.Guest = &types.VirtualMachineGuestSummary{}
	vm.Summary.Vm = &vm.Self
	vm.Summary.Storage = &types.VirtualMachineStorageSummary{
		Timestamp: time.Now(),
	}

	vmx := vm.vmx(spec)
	if vmx.Path == "" {
		// Append VM Name as the directory name if not specified
		vmx.Path = spec.Name
	}

	dc := ctx.Map.getEntityDatacenter(folder.(mo.Entity))
	ds := ctx.Map.FindByName(vmx.Datastore, dc.Datastore).(*Datastore)
	dir := ds.resolve(ctx, vmx.Path)

	if path.Ext(vmx.Path) == ".vmx" {
		dir = path.Dir(dir)
		// Ignore error here, deferring to createFile
		_ = os.Mkdir(dir, 0700)
	} else {
		// Create VM directory, renaming if already exists
		name := dir

		for i := 0; i < 1024; /* just in case */ i++ {
			err := os.Mkdir(name, 0700)
			if err != nil {
				if os.IsExist(err) {
					name = fmt.Sprintf("%s (%d)", dir, i)
					continue
				}
				return nil, &types.FileFault{File: name}
			}
			break
		}
		vmx.Path = path.Join(path.Base(name), spec.Name+".vmx")
	}

	spec.Files.VmPathName = vmx.String()

	dsPath := path.Dir(spec.Files.VmPathName)
	vm.uid = internal.OID(spec.Files.VmPathName)

	defaults := types.VirtualMachineConfigSpec{
		NumCPUs:           1,
		NumCoresPerSocket: types.NewInt32(1),
		MemoryMB:          32,
		Uuid:              vm.uid.String(),
		InstanceUuid:      newUUID(strings.ToUpper(spec.Files.VmPathName)),
		Version:           esx.HardwareVersion,
		Firmware:          string(types.GuestOsDescriptorFirmwareTypeBios),
		VAppConfig:        spec.VAppConfig,
		Files: &types.VirtualMachineFileInfo{
			SnapshotDirectory: dsPath,
			SuspendDirectory:  dsPath,
			LogDirectory:      dsPath,
		},
	}

	// Add the default devices
	defaults.DeviceChange, _ = object.VirtualDeviceList(esx.VirtualDevice).ConfigSpec(types.VirtualDeviceConfigSpecOperationAdd)

	err := vm.configure(ctx, &defaults)
	if err != nil {
		return vm, err
	}

	vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOff
	vm.Runtime.ConnectionState = types.VirtualMachineConnectionStateConnected
	vm.Summary.Runtime = vm.Runtime

	vm.Capability.ChangeTrackingSupported = changeTrackingSupported(spec)

	vm.Summary.QuickStats.GuestHeartbeatStatus = types.ManagedEntityStatusGray
	vm.Summary.OverallStatus = types.ManagedEntityStatusGreen
	vm.ConfigStatus = types.ManagedEntityStatusGreen
	vm.DataSets = make(map[string]*DataSet)

	// put vm in the folder only if no errors occurred
	f, _ := asFolderMO(folder)
	folderPutChild(ctx, f, vm)

	return vm, nil
}

func (o *VirtualMachine) RenameTask(ctx *Context, r *types.Rename_Task) soap.HasFault {
	return RenameTask(ctx, o, r)
}

func (*VirtualMachine) Reload(*types.Reload) soap.HasFault {
	return &methods.ReloadBody{Res: new(types.ReloadResponse)}
}

func (vm *VirtualMachine) event(ctx *Context) types.VmEvent {
	host := ctx.Map.Get(*vm.Runtime.Host).(*HostSystem)

	return types.VmEvent{
		Event: types.Event{
			Datacenter:      datacenterEventArgument(ctx, host),
			ComputeResource: host.eventArgumentParent(ctx),
			Host:            host.eventArgument(),
			Ds:              ctx.Map.Get(vm.Datastore[0]).(*Datastore).eventArgument(),
			Vm: &types.VmEventArgument{
				EntityEventArgument: types.EntityEventArgument{Name: vm.Name},
				Vm:                  vm.Self,
			},
		},
	}
}

func (vm *VirtualMachine) hostInMM(ctx *Context) bool {
	return ctx.Map.Get(*vm.Runtime.Host).(*HostSystem).Runtime.InMaintenanceMode
}

func (vm *VirtualMachine) apply(spec *types.VirtualMachineConfigSpec) {
	if spec.Files == nil {
		spec.Files = new(types.VirtualMachineFileInfo)
	}

	apply := []struct {
		src string
		dst *string
	}{
		{spec.AlternateGuestName, &vm.Config.AlternateGuestName},
		{spec.Annotation, &vm.Config.Annotation},
		{spec.Firmware, &vm.Config.Firmware},
		{spec.InstanceUuid, &vm.Config.InstanceUuid},
		{spec.LocationId, &vm.Config.LocationId},
		{spec.NpivWorldWideNameType, &vm.Config.NpivWorldWideNameType},
		{spec.Name, &vm.Name},
		{spec.Name, &vm.Config.Name},
		{spec.Name, &vm.Summary.Config.Name},
		{spec.GuestId, &vm.Config.GuestId},
		{spec.GuestId, &vm.Config.GuestFullName},
		{spec.GuestId, &vm.Summary.Guest.GuestId},
		{spec.GuestId, &vm.Summary.Config.GuestId},
		{spec.GuestId, &vm.Summary.Config.GuestFullName},
		{spec.Uuid, &vm.Config.Uuid},
		{spec.Uuid, &vm.Summary.Config.Uuid},
		{spec.InstanceUuid, &vm.Config.InstanceUuid},
		{spec.InstanceUuid, &vm.Summary.Config.InstanceUuid},
		{spec.Version, &vm.Config.Version},
		{spec.Version, &vm.Summary.Config.HwVersion},
		{spec.Files.VmPathName, &vm.Config.Files.VmPathName},
		{spec.Files.VmPathName, &vm.Summary.Config.VmPathName},
		{spec.Files.SnapshotDirectory, &vm.Config.Files.SnapshotDirectory},
		{spec.Files.SuspendDirectory, &vm.Config.Files.SuspendDirectory},
		{spec.Files.LogDirectory, &vm.Config.Files.LogDirectory},
		{spec.FtEncryptionMode, &vm.Config.FtEncryptionMode},
		{spec.MigrateEncryption, &vm.Config.MigrateEncryption},
	}

	for _, f := range apply {
		if f.src != "" {
			*f.dst = f.src
		}
	}

	applyb := []struct {
		src *bool
		dst **bool
	}{
		{spec.NestedHVEnabled, &vm.Config.NestedHVEnabled},
		{spec.CpuHotAddEnabled, &vm.Config.CpuHotAddEnabled},
		{spec.CpuHotRemoveEnabled, &vm.Config.CpuHotRemoveEnabled},
		{spec.GuestAutoLockEnabled, &vm.Config.GuestAutoLockEnabled},
		{spec.MemoryHotAddEnabled, &vm.Config.MemoryHotAddEnabled},
		{spec.MemoryReservationLockedToMax, &vm.Config.MemoryReservationLockedToMax},
		{spec.MessageBusTunnelEnabled, &vm.Config.MessageBusTunnelEnabled},
		{spec.NpivTemporaryDisabled, &vm.Config.NpivTemporaryDisabled},
		{spec.NpivOnNonRdmDisks, &vm.Config.NpivOnNonRdmDisks},
		{spec.ChangeTrackingEnabled, &vm.Config.ChangeTrackingEnabled},
	}

	for _, f := range applyb {
		if f.src != nil {
			*f.dst = f.src
		}
	}

	if spec.Flags != nil {
		vm.Config.Flags = *spec.Flags
	}

	if spec.LatencySensitivity != nil {
		vm.Config.LatencySensitivity = spec.LatencySensitivity
	}

	if spec.ManagedBy != nil {
		if spec.ManagedBy.ExtensionKey == "" {
			spec.ManagedBy = nil
		}
		vm.Config.ManagedBy = spec.ManagedBy
		vm.Summary.Config.ManagedBy = spec.ManagedBy
	}

	if spec.BootOptions != nil {
		vm.Config.BootOptions = spec.BootOptions
	}

	if spec.RepConfig != nil {
		vm.Config.RepConfig = spec.RepConfig
	}

	if spec.Tools != nil {
		vm.Config.Tools = spec.Tools
	}

	if spec.ConsolePreferences != nil {
		vm.Config.ConsolePreferences = spec.ConsolePreferences
	}

	if spec.CpuAffinity != nil {
		vm.Config.CpuAffinity = spec.CpuAffinity
	}

	if spec.CpuAllocation != nil {
		vm.Config.CpuAllocation = spec.CpuAllocation
	}

	if spec.MemoryAffinity != nil {
		vm.Config.MemoryAffinity = spec.MemoryAffinity
	}

	if spec.MemoryAllocation != nil {
		vm.Config.MemoryAllocation = spec.MemoryAllocation
	}

	if spec.LatencySensitivity != nil {
		vm.Config.LatencySensitivity = spec.LatencySensitivity
	}

	if spec.MemoryMB != 0 {
		vm.Config.Hardware.MemoryMB = int32(spec.MemoryMB)
		vm.Summary.Config.MemorySizeMB = vm.Config.Hardware.MemoryMB
	}

	if spec.NumCPUs != 0 {
		vm.Config.Hardware.NumCPU = spec.NumCPUs
		vm.Summary.Config.NumCpu = vm.Config.Hardware.NumCPU
	}

	if spec.NumCoresPerSocket != nil {
		vm.Config.Hardware.NumCoresPerSocket = spec.NumCoresPerSocket
	}

	if spec.GuestId != "" {
		vm.Guest.GuestFamily = guestFamily(spec.GuestId)
	}

	vm.Config.Modified = time.Now()
}

// updateVAppProperty updates the simulator VM with the specified VApp properties.
func (vm *VirtualMachine) updateVAppProperty(spec *types.VmConfigSpec) types.BaseMethodFault {
	if vm.Config.VAppConfig == nil {
		vm.Config.VAppConfig = &types.VmConfigInfo{}
	}

	info := vm.Config.VAppConfig.GetVmConfigInfo()
	propertyInfo := info.Property
	productInfo := info.Product

	for _, prop := range spec.Property {
		var foundIndex int
		exists := false
		// Check if the specified property exists or not. This helps rejecting invalid
		// operations (e.g., Adding a VApp property that already exists)
		for i, p := range propertyInfo {
			if p.Key == prop.Info.Key {
				exists = true
				foundIndex = i
				break
			}
		}

		switch prop.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			propertyInfo = append(propertyInfo, *prop.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			propertyInfo[foundIndex] = *prop.Info
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			propertyInfo = append(propertyInfo[:foundIndex], propertyInfo[foundIndex+1:]...)
		}
	}

	for _, prod := range spec.Product {
		var foundIndex int
		exists := false
		// Check if the specified product exists or not. This helps rejecting invalid
		// operations (e.g., Adding a VApp product that already exists)
		for i, p := range productInfo {
			if p.Key == prod.Info.Key {
				exists = true
				foundIndex = i
				break
			}
		}

		switch prod.Operation {
		case types.ArrayUpdateOperationAdd:
			if exists {
				return new(types.InvalidArgument)
			}
			productInfo = append(productInfo, *prod.Info)
		case types.ArrayUpdateOperationEdit:
			if !exists {
				return new(types.InvalidArgument)
			}
			productInfo[foundIndex] = *prod.Info
		case types.ArrayUpdateOperationRemove:
			if !exists {
				return new(types.InvalidArgument)
			}
			productInfo = append(productInfo[:foundIndex], productInfo[foundIndex+1:]...)
		}
	}

	info.Product = productInfo
	info.Property = propertyInfo

	return nil
}

var extraConfigAlias = map[string]string{
	"ip0": "SET.guest.ipAddress",
}

func extraConfigKey(key string) string {
	if k, ok := extraConfigAlias[key]; ok {
		return k
	}
	return key
}

func (vm *VirtualMachine) applyExtraConfig(ctx *Context, spec *types.VirtualMachineConfigSpec) types.BaseMethodFault {
	if len(spec.ExtraConfig) == 0 {
		return nil
	}
	var removedContainerBacking bool
	var changes []types.PropertyChange
	field := mo.Field{Path: "config.extraConfig"}

	for _, c := range spec.ExtraConfig {
		val := c.GetOptionValue()
		key := strings.TrimPrefix(extraConfigKey(val.Key), "SET.")
		if key == val.Key {
			field.Key = key
			op := types.PropertyChangeOpAssign
			keyIndex := -1
			for i := range vm.Config.ExtraConfig {
				bov := vm.Config.ExtraConfig[i]
				if bov == nil {
					continue
				}
				ov := bov.GetOptionValue()
				if ov == nil {
					continue
				}
				if ov.Key == key {
					keyIndex = i
					break
				}
			}
			if keyIndex < 0 {
				op = types.PropertyChangeOpAdd
				vm.Config.ExtraConfig = append(vm.Config.ExtraConfig, c)
			} else {
				if s, ok := val.Value.(string); ok && s == "" {
					op = types.PropertyChangeOpRemove
					if key == ContainerBackingOptionKey {
						removedContainerBacking = true
					}
					// Remove existing element
					vm.Config.ExtraConfig = append(
						vm.Config.ExtraConfig[:keyIndex],
						vm.Config.ExtraConfig[keyIndex+1:]...)
					val = nil
				} else {
					// Update existing element
					vm.Config.ExtraConfig[keyIndex] = val
				}
			}

			changes = append(changes, types.PropertyChange{Name: field.String(), Val: val, Op: op})
			continue
		}
		changes = append(changes, types.PropertyChange{Name: key, Val: val.Value})

		switch key {
		case "guest.ipAddress":
			if len(vm.Guest.Net) > 0 {
				ip := val.Value.(string)
				vm.Guest.Net[0].IpAddress = []string{ip}
				changes = append(changes,
					types.PropertyChange{Name: "summary." + key, Val: ip},
					types.PropertyChange{Name: "guest.net", Val: vm.Guest.Net},
				)
			}
		case "guest.hostName":
			changes = append(changes,
				types.PropertyChange{Name: "summary." + key, Val: val.Value},
			)
		}
	}

	// create the container backing before we publish the updates so the simVM is available before handlers
	// get triggered
	var fault types.BaseMethodFault
	if vm.svm == nil {
		vm.svm = createSimulationVM(vm)

		// check to see if the VM is already powered on - if so we need to retroactively hit that path here
		if vm.Runtime.PowerState == types.VirtualMachinePowerStatePoweredOn {
			err := vm.svm.start(ctx)
			if err != nil {
				// don't attempt to undo the changes already made - just return an error
				// we'll retry the svm.start operation on pause/restart calls
				fault = &types.VAppConfigFault{
					VimFault: types.VimFault{
						MethodFault: types.MethodFault{
							FaultCause: &types.LocalizedMethodFault{
								Fault:            &types.SystemErrorFault{Reason: err.Error()},
								LocalizedMessage: err.Error()}}}}
			}
		}
	} else if removedContainerBacking {
		err := vm.svm.remove(ctx)
		if err == nil {
			// remove link from container to VM so callbacks no longer reflect state
			vm.svm.vm = nil
			// nil container backing reference to return this to a pure in-mem simulated VM
			vm.svm = nil

		} else {
			// don't attempt to undo the changes already made - just return an error
			// we'll retry the svm.start operation on pause/restart calls
			fault = &types.VAppConfigFault{
				VimFault: types.VimFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault:            &types.SystemErrorFault{Reason: err.Error()},
							LocalizedMessage: err.Error()}}}}
		}
	}

	change := types.PropertyChange{Name: field.Path, Val: vm.Config.ExtraConfig}
	ctx.Update(vm, append(changes, change))

	return fault
}

func validateGuestID(id string) types.BaseMethodFault {
	for _, x := range GuestID {
		if id == string(x) {
			return nil
		}
	}

	return &types.InvalidArgument{InvalidProperty: "configSpec.guestId"}
}

func (vm *VirtualMachine) configure(ctx *Context, spec *types.VirtualMachineConfigSpec) (result types.BaseMethodFault) {
	defer func() {
		if result == nil {
			vm.updateLastModifiedAndChangeVersion(ctx)
		}
	}()

	vm.apply(spec)

	if spec.MemoryAllocation != nil {
		if err := updateResourceAllocation("memory", spec.MemoryAllocation, vm.Config.MemoryAllocation); err != nil {
			return err
		}
	}

	if spec.CpuAllocation != nil {
		if err := updateResourceAllocation("cpu", spec.CpuAllocation, vm.Config.CpuAllocation); err != nil {
			return err
		}
	}

	if spec.GuestId != "" {
		if err := validateGuestID(spec.GuestId); err != nil {
			return err
		}
	}

	if o := spec.BootOptions; o != nil {
		if isTrue(o.EfiSecureBootEnabled) && vm.Config.Firmware != string(types.GuestOsDescriptorFirmwareTypeEfi) {
			return &types.InvalidVmConfig{Property: "msg.hostd.configSpec.efi"}
		}
	}

	if spec.VAppConfig != nil {
		if err := vm.updateVAppProperty(spec.VAppConfig.GetVmConfigSpec()); err != nil {
			return err
		}
	}

	if spec.Crypto != nil {
		if err := vm.updateCrypto(ctx, spec.Crypto); err != nil {
			return err
		}
	}

	if err := vm.updateTagSpec(ctx, spec.TagSpecs); err != nil {
		return err
	}

	return vm.configureDevices(ctx, spec)
}

func getVMFileType(fileName string) types.VirtualMachineFileLayoutExFileType {
	var fileType types.VirtualMachineFileLayoutExFileType

	fileExt := path.Ext(fileName)
	fileNameNoExt := strings.TrimSuffix(fileName, fileExt)

	switch fileExt {
	case ".vmx":
		fileType = types.VirtualMachineFileLayoutExFileTypeConfig
	case ".core":
		fileType = types.VirtualMachineFileLayoutExFileTypeCore
	case ".vmdk":
		fileType = types.VirtualMachineFileLayoutExFileTypeDiskDescriptor
		if strings.HasSuffix(fileNameNoExt, "-digest") {
			fileType = types.VirtualMachineFileLayoutExFileTypeDigestDescriptor
		}

		extentSuffixes := []string{"-flat", "-delta", "-s", "-rdm", "-rdmp"}
		for _, suffix := range extentSuffixes {
			if strings.HasSuffix(fileNameNoExt, suffix) {
				fileType = types.VirtualMachineFileLayoutExFileTypeDiskExtent
			} else if strings.HasSuffix(fileNameNoExt, "-digest"+suffix) {
				fileType = types.VirtualMachineFileLayoutExFileTypeDigestExtent
			}
		}
	case ".psf":
		fileType = types.VirtualMachineFileLayoutExFileTypeDiskReplicationState
	case ".vmxf":
		fileType = types.VirtualMachineFileLayoutExFileTypeExtendedConfig
	case ".vmft":
		fileType = types.VirtualMachineFileLayoutExFileTypeFtMetadata
	case ".log":
		fileType = types.VirtualMachineFileLayoutExFileTypeLog
	case ".nvram":
		fileType = types.VirtualMachineFileLayoutExFileTypeNvram
	case ".png", ".bmp":
		fileType = types.VirtualMachineFileLayoutExFileTypeScreenshot
	case ".vmsn":
		fileType = types.VirtualMachineFileLayoutExFileTypeSnapshotData
	case ".vmsd":
		fileType = types.VirtualMachineFileLayoutExFileTypeSnapshotList
	case ".xml":
		if strings.HasSuffix(fileNameNoExt, "-aux") {
			fileType = types.VirtualMachineFileLayoutExFileTypeSnapshotManifestList
		}
	case ".stat":
		fileType = types.VirtualMachineFileLayoutExFileTypeStat
	case ".vmss":
		fileType = types.VirtualMachineFileLayoutExFileTypeSuspend
	case ".vmem":
		if strings.Contains(fileNameNoExt, "Snapshot") {
			fileType = types.VirtualMachineFileLayoutExFileTypeSnapshotMemory
		} else {
			fileType = types.VirtualMachineFileLayoutExFileTypeSuspendMemory
		}
	case ".vswp":
		if strings.HasPrefix(fileNameNoExt, "vmx-") {
			fileType = types.VirtualMachineFileLayoutExFileTypeUwswap
		} else {
			fileType = types.VirtualMachineFileLayoutExFileTypeSwap
		}
	case "":
		if strings.HasPrefix(fileNameNoExt, "imcf-") {
			fileType = types.VirtualMachineFileLayoutExFileTypeGuestCustomization
		}
	}

	return fileType
}

func (vm *VirtualMachine) addFileLayoutEx(ctx *Context, datastorePath object.DatastorePath, fileSize int64) int32 {
	var newKey int32
	for _, layoutFile := range vm.LayoutEx.File {
		if layoutFile.Name == datastorePath.String() {
			return layoutFile.Key
		}

		if layoutFile.Key >= newKey {
			newKey = layoutFile.Key + 1
		}
	}

	fileType := getVMFileType(filepath.Base(datastorePath.Path))

	switch fileType {
	case types.VirtualMachineFileLayoutExFileTypeNvram, types.VirtualMachineFileLayoutExFileTypeSnapshotList:
		if !slices.Contains(vm.Layout.ConfigFile, datastorePath.Path) {
			vm.Layout.ConfigFile = append(vm.Layout.ConfigFile, datastorePath.Path)
		}
	case types.VirtualMachineFileLayoutExFileTypeLog:
		if !slices.Contains(vm.Layout.LogFile, datastorePath.Path) {
			vm.Layout.LogFile = append(vm.Layout.LogFile, datastorePath.Path)
		}
	case types.VirtualMachineFileLayoutExFileTypeSwap:
		vm.Layout.SwapFile = datastorePath.String()
	}

	vm.LayoutEx.File = append(vm.LayoutEx.File, types.VirtualMachineFileLayoutExFileInfo{
		Accessible:      types.NewBool(true),
		BackingObjectId: "",
		Key:             newKey,
		Name:            datastorePath.String(),
		Size:            fileSize,
		Type:            string(fileType),
		UniqueSize:      fileSize,
	})

	vm.LayoutEx.Timestamp = time.Now()

	vm.updateStorage(ctx)

	return newKey
}

func (vm *VirtualMachine) addSnapshotLayout(ctx *Context, snapshot types.ManagedObjectReference, dataKey int32) {
	for _, snapshotLayout := range vm.Layout.Snapshot {
		if snapshotLayout.Key == snapshot {
			return
		}
	}

	var snapshotFiles []string
	for _, file := range vm.LayoutEx.File {
		if file.Key == dataKey || file.Type == "diskDescriptor" {
			snapshotFiles = append(snapshotFiles, file.Name)
		}
	}

	vm.Layout.Snapshot = append(vm.Layout.Snapshot, types.VirtualMachineFileLayoutSnapshotLayout{
		Key:          snapshot,
		SnapshotFile: snapshotFiles,
	})
}

func (vm *VirtualMachine) addSnapshotLayoutEx(ctx *Context, snapshot types.ManagedObjectReference, dataKey int32, memoryKey int32) {
	for _, snapshotLayoutEx := range vm.LayoutEx.Snapshot {
		if snapshotLayoutEx.Key == snapshot {
			return
		}
	}

	vm.LayoutEx.Snapshot = append(vm.LayoutEx.Snapshot, types.VirtualMachineFileLayoutExSnapshotLayout{
		DataKey:   dataKey,
		Disk:      vm.LayoutEx.Disk,
		Key:       snapshot,
		MemoryKey: memoryKey,
	})

	vm.LayoutEx.Timestamp = time.Now()

	vm.updateStorage(ctx)
}

// Updates both vm.Layout.Disk and vm.LayoutEx.Disk
func (vm *VirtualMachine) updateDiskLayouts(ctx *Context) types.BaseMethodFault {
	var disksLayout []types.VirtualMachineFileLayoutDiskLayout
	var disksLayoutEx []types.VirtualMachineFileLayoutExDiskLayout

	disks := object.VirtualDeviceList(vm.Config.Hardware.Device).SelectByType((*types.VirtualDisk)(nil))
	for _, disk := range disks {
		disk := disk.(*types.VirtualDisk)
		diskBacking, ok := disk.Backing.(types.BaseVirtualDeviceFileBackingInfo)
		if !ok {
			continue
		}

		diskLayout := &types.VirtualMachineFileLayoutDiskLayout{Key: disk.Key}
		diskLayoutEx := &types.VirtualMachineFileLayoutExDiskLayout{Key: disk.Key}

		// Iterate through disk and its parents
		for diskBacking != nil {

			dFileName := diskBacking.GetVirtualDeviceFileBackingInfo().FileName

			var fileKeys []int32

			// Add disk descriptor and extent files
			for _, diskName := range vdmNames(dFileName) {
				// get full path including datastore location
				p, fault := parseDatastorePath(diskName)
				if fault != nil {
					return fault
				}

				datastore := vm.useDatastore(ctx, p.Datastore)
				dFilePath := datastore.resolve(ctx, p.Path)

				var fileSize int64
				// If file can not be opened - fileSize will be 0
				if dFileInfo, err := os.Stat(dFilePath); err == nil {
					fileSize = dFileInfo.Size()
				}

				diskKey := vm.addFileLayoutEx(ctx, *p, fileSize)
				fileKeys = append(fileKeys, diskKey)
			}

			diskLayout.DiskFile = append(diskLayout.DiskFile, dFileName)
			diskLayoutEx.Chain = append(diskLayoutEx.Chain, types.VirtualMachineFileLayoutExDiskUnit{
				FileKey: fileKeys,
			})

			switch tBack := diskBacking.(type) {
			case *types.VirtualDiskFlatVer1BackingInfo:
				if tBack.Parent == nil {
					diskBacking = nil
				} else {
					diskBacking = tBack.Parent
				}
			case *types.VirtualDiskFlatVer2BackingInfo:
				if tBack.Parent == nil {
					diskBacking = nil
				} else {
					diskBacking = tBack.Parent
				}
			case *types.VirtualDiskSeSparseBackingInfo:
				if tBack.Parent == nil {
					diskBacking = nil
				} else {
					diskBacking = tBack.Parent
				}
			case *types.VirtualDiskSparseVer1BackingInfo:
				if tBack.Parent == nil {
					diskBacking = nil
				} else {
					diskBacking = tBack.Parent
				}
			case *types.VirtualDiskSparseVer2BackingInfo:
				if tBack.Parent == nil {
					diskBacking = nil
				} else {
					diskBacking = tBack.Parent
				}
			default:
				diskBacking = nil
			}
		}

		disksLayout = append(disksLayout, *diskLayout)
		disksLayoutEx = append(disksLayoutEx, *diskLayoutEx)
	}

	vm.Layout.Disk = disksLayout

	vm.LayoutEx.Disk = disksLayoutEx
	vm.LayoutEx.Timestamp = time.Now()

	vm.updateStorage(ctx)

	return nil
}

func (vm *VirtualMachine) updateStorage(ctx *Context) types.BaseMethodFault {
	// Committed - sum of Size for each file in vm.LayoutEx.File
	// Unshared  - sum of Size for each disk (.vmdk) in vm.LayoutEx.File
	// Uncommitted - disk capacity minus disk usage (only currently used disk)
	var datastoresUsage []types.VirtualMachineUsageOnDatastore

	disks := object.VirtualDeviceList(vm.Config.Hardware.Device).SelectByType((*types.VirtualDisk)(nil))

	for _, file := range vm.LayoutEx.File {
		p, fault := parseDatastorePath(file.Name)
		if fault != nil {
			return fault
		}

		datastore := vm.useDatastore(ctx, p.Datastore)
		dsUsage := &types.VirtualMachineUsageOnDatastore{
			Datastore: datastore.Self,
		}

		for idx, usage := range datastoresUsage {
			if usage.Datastore == datastore.Self {
				datastoresUsage = append(datastoresUsage[:idx], datastoresUsage[idx+1:]...)
				dsUsage = &usage
				break
			}
		}

		dsUsage.Committed += file.Size

		if path.Ext(file.Name) == ".vmdk" {
			dsUsage.Unshared += file.Size
		}

		for _, disk := range disks {
			disk := disk.(*types.VirtualDisk)
			backing := disk.Backing.(types.BaseVirtualDeviceFileBackingInfo).GetVirtualDeviceFileBackingInfo()

			if backing.FileName == file.Name {
				dsUsage.Uncommitted += disk.CapacityInBytes
			}
		}

		datastoresUsage = append(datastoresUsage, *dsUsage)
	}

	vm.Storage.PerDatastoreUsage = datastoresUsage
	vm.Storage.Timestamp = time.Now()

	storageSummary := &types.VirtualMachineStorageSummary{
		Timestamp: time.Now(),
	}

	for i, usage := range datastoresUsage {
		datastoresUsage[i].Uncommitted -= usage.Committed
		storageSummary.Committed += usage.Committed
		storageSummary.Uncommitted += usage.Uncommitted
		storageSummary.Unshared += usage.Unshared
	}

	vm.Summary.Storage = storageSummary

	return nil
}

func (vm *VirtualMachine) RefreshStorageInfo(ctx *Context, req *types.RefreshStorageInfo) soap.HasFault {
	body := new(methods.RefreshStorageInfoBody)

	if vm.Runtime.Host == nil {
		// VM not fully created
		return body
	}

	// Validate that all files in vm.LayoutEx.File can still be found
	for idx := len(vm.LayoutEx.File) - 1; idx >= 0; idx-- {
		file := vm.LayoutEx.File[idx]

		p, fault := parseDatastorePath(file.Name)
		if fault != nil {
			body.Fault_ = Fault("", fault)
			return body
		}

		if _, err := os.Stat(p.String()); err != nil {
			vm.LayoutEx.File = append(vm.LayoutEx.File[:idx], vm.LayoutEx.File[idx+1:]...)
		}
	}

	vmPathName := vm.Config.Files.VmPathName
	// vm.Config.Files.VmPathName can be a directory or full path to .vmx
	if path.Ext(vmPathName) == ".vmx" {
		vmPathName = path.Dir(vmPathName)
	}

	// Directories will be used to locate VM files.
	// Does not include information about virtual disk file locations.
	locations := []string{
		vmPathName,
		vm.Config.Files.SnapshotDirectory,
		vm.Config.Files.LogDirectory,
		vm.Config.Files.SuspendDirectory,
		vm.Config.Files.FtMetadataDirectory,
	}

	for _, directory := range slices.Compact(locations) {
		if directory == "" {
			continue
		}

		p, fault := parseDatastorePath(directory)
		if fault != nil {
			body.Fault_ = Fault("", fault)
			return body
		}

		datastore := vm.useDatastore(ctx, p.Datastore)
		directory := datastore.resolve(ctx, p.Path)

		if _, err := os.Stat(directory); err != nil {
			// Can not access the directory
			continue
		}

		files, err := os.ReadDir(directory)
		if err != nil {
			body.Fault_ = Fault("", ctx.Map.FileManager().fault(directory, err, new(types.CannotAccessFile)))
			return body
		}

		for _, file := range files {
			datastorePath := object.DatastorePath{
				Datastore: p.Datastore,
				Path:      path.Join(p.Path, file.Name()),
			}
			info, _ := file.Info()
			vm.addFileLayoutEx(ctx, datastorePath, info.Size())
		}
	}

	fault := vm.updateDiskLayouts(ctx)
	if fault != nil {
		body.Fault_ = Fault("", fault)
		return body
	}

	vm.LayoutEx.Timestamp = time.Now()

	body.Res = new(types.RefreshStorageInfoResponse)

	return body
}

func (vm *VirtualMachine) findDatastore(ctx *Context, name string) *Datastore {
	host := ctx.Map.Get(*vm.Runtime.Host).(*HostSystem)

	return ctx.Map.FindByName(name, host.Datastore).(*Datastore)
}

func (vm *VirtualMachine) useDatastore(ctx *Context, name string) *Datastore {
	ds := vm.findDatastore(ctx, name)
	if FindReference(vm.Datastore, ds.Self) == nil {
		vm.Datastore = append(vm.Datastore, ds.Self)
	}

	return ds
}

func (vm *VirtualMachine) vmx(spec *types.VirtualMachineConfigSpec) object.DatastorePath {
	var p object.DatastorePath
	vmx := vm.Config.Files.VmPathName
	if spec != nil {
		vmx = spec.Files.VmPathName
	}
	p.FromString(vmx)
	return p
}

func (vm *VirtualMachine) createFile(ctx *Context, spec string, name string, register bool) (*os.File, types.BaseMethodFault) {
	p, fault := parseDatastorePath(spec)
	if fault != nil {
		return nil, fault
	}

	ds := vm.useDatastore(ctx, p.Datastore)

	nhost := len(ds.Host)
	if internal.IsDatastoreVSAN(ds.Datastore) && nhost < 3 {
		fault := new(types.CannotCreateFile)
		fault.FaultMessage = []types.LocalizableMessage{
			{
				Key:     "vob.vsanprovider.object.creation.failed",
				Message: "Failed to create object.",
			},
			{
				Key:     "vob.vsan.clomd.needMoreFaultDomains2",
				Message: fmt.Sprintf("There are currently %d usable fault domains. The operation requires %d more usable fault domains.", nhost, 3-nhost),
			},
		}
		fault.File = p.Path
		return nil, fault
	}

	file := ds.resolve(ctx, p.Path)

	if name != "" {
		if path.Ext(p.Path) == ".vmx" {
			file = path.Dir(file) // vm.Config.Files.VmPathName can be a directory or full path to .vmx
		}

		file = path.Join(file, name)
	}

	if register {
		f, err := os.Open(filepath.Clean(file))
		if err != nil {
			log.Printf("register %s: %s", vm.Reference(), err)
			if os.IsNotExist(err) {
				return nil, &types.NotFound{}
			}

			return nil, &types.InvalidArgument{}
		}

		return f, nil
	}

	_, err := os.Stat(file)
	if err == nil {
		switch path.Ext(file) {
		case ".nvram":
			f, err := os.Open(file)
			if err != nil {
				return nil, &types.FileFault{
					File: file,
				}
			}
			return f, nil
		default:
			fault := &types.FileAlreadyExists{FileFault: types.FileFault{File: file}}
			log.Printf("%T: %s", fault, file)
			return nil, fault
		}
	}

	// Create parent directory if needed
	dir := path.Dir(file)
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			_ = os.Mkdir(dir, 0700)
		}
	}

	f, err := os.Create(file)
	if err != nil {
		log.Printf("create(%s): %s", file, err)
		return nil, &types.FileFault{
			File: file,
		}
	}

	return f, nil
}

// Rather than keep an fd open for each VM, open/close the log for each messages.
// This is ok for now as we do not do any heavy VM logging.
func (vm *VirtualMachine) logPrintf(format string, v ...any) {
	f, err := os.OpenFile(vm.log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.New(f, "vmx ", log.Flags()).Printf(format, v...)
	_ = f.Close()
}

func (vm *VirtualMachine) create(ctx *Context, spec *types.VirtualMachineConfigSpec, register bool) types.BaseMethodFault {
	vm.apply(spec)

	if spec.Version != "" {
		v := strings.TrimPrefix(spec.Version, "vmx-")
		_, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("unsupported hardware version: %s", spec.Version)
			return new(types.NotSupported)
		}
	}

	files := []struct {
		spec string
		name string
		use  *string
	}{
		{vm.Config.Files.VmPathName, "", nil},
		{vm.Config.Files.VmPathName, fmt.Sprintf("%s.nvram", vm.Name), nil},
		{vm.Config.Files.LogDirectory, "vmware.log", &vm.log},
	}

	for _, file := range files {
		f, err := vm.createFile(ctx, file.spec, file.name, register)
		if err != nil {
			return err
		}
		if file.use != nil {
			*file.use = f.Name()
		}
		_ = f.Close()
	}

	vm.logPrintf("created")

	return vm.configureDevices(ctx, spec)
}

var vmwOUI = net.HardwareAddr([]byte{0x0, 0xc, 0x29})

// From https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-networking-8-0/mac-addresses/mac-address-generation-on-esxi-hosts.html
// > The host generates generateMAC addresses that consists of the VMware OUI 00:0C:29 and the last three octets in hexadecimal
// > format of the virtual machine UUID.  The virtual machine UUID is based on a hash calculated by using the UUID of the
// > ESXi physical machine and the path to the configuration file (.vmx) of the virtual machine.
func (vm *VirtualMachine) generateMAC(unit int32) string {
	id := []byte(vm.Config.Uuid)

	offset := len(id) - len(vmwOUI)
	key := id[offset] + byte(unit) // add device unit number, giving each VM NIC a unique MAC
	id = append([]byte{key}, id[offset+1:]...)

	mac := append(vmwOUI, id...)

	return mac.String()
}

func numberToString(n int64, sep rune) string {
	buf := &bytes.Buffer{}
	if n < 0 {
		n = -n
		buf.WriteRune('-')
	}
	s := strconv.FormatInt(n, 10)
	pos := 3 - (len(s) % 3)
	for i := 0; i < len(s); i++ {
		if pos == 3 {
			if i != 0 {
				buf.WriteRune(sep)
			}
			pos = 0
		}
		pos++
		buf.WriteByte(s[i])
	}

	return buf.String()
}

func getDiskSize(disk *types.VirtualDisk) int64 {
	if disk.CapacityInBytes == 0 {
		return disk.CapacityInKB * 1024
	}
	return disk.CapacityInBytes
}

func changedDiskSize(oldDisk *types.VirtualDisk, newDiskSpec *types.VirtualDisk) (int64, bool) {
	// capacity cannot be decreased
	if newDiskSpec.CapacityInBytes > 0 && newDiskSpec.CapacityInBytes < oldDisk.CapacityInBytes {
		return 0, false
	}
	if newDiskSpec.CapacityInKB > 0 && newDiskSpec.CapacityInKB < oldDisk.CapacityInKB {
		return 0, false
	}

	// NOTE: capacity is ignored if specified value is same as before
	if newDiskSpec.CapacityInBytes == oldDisk.CapacityInBytes {
		return newDiskSpec.CapacityInKB * 1024, true
	}
	if newDiskSpec.CapacityInKB == oldDisk.CapacityInKB {
		return newDiskSpec.CapacityInBytes, true
	}

	// if both set, CapacityInBytes and CapacityInKB must be the same
	if newDiskSpec.CapacityInBytes > 0 && newDiskSpec.CapacityInKB > 0 {
		if newDiskSpec.CapacityInBytes != newDiskSpec.CapacityInKB*1024 {
			return 0, false
		}
	}

	return newDiskSpec.CapacityInBytes, true
}

func (vm *VirtualMachine) validateSwitchMembers(ctx *Context, id string) types.BaseMethodFault {
	var dswitch *DistributedVirtualSwitch

	var find func(types.ManagedObjectReference)
	find = func(child types.ManagedObjectReference) {
		s, ok := ctx.Map.Get(child).(*DistributedVirtualSwitch)
		if ok && s.Uuid == id {
			dswitch = s
			return
		}
		walk(ctx.Map.Get(child), find)
	}
	f := ctx.Map.getEntityDatacenter(vm).NetworkFolder
	walk(ctx.Map.Get(f), find) // search in NetworkFolder and any sub folders

	if dswitch == nil {
		log.Printf("DVS %s cannot be found", id)
		return new(types.NotFound)
	}

	h := ctx.Map.Get(*vm.Runtime.Host).(*HostSystem)
	c := hostParent(ctx, &h.HostSystem)
	isMember := func(val types.ManagedObjectReference) bool {
		for _, mem := range dswitch.Summary.HostMember {
			if mem == val {
				return true
			}
		}
		log.Printf("%s is not a member of VDS %s", h.Name, dswitch.Name)
		return false
	}

	for _, ref := range c.Host {
		if !isMember(ref) {
			return &types.InvalidArgument{InvalidProperty: "spec.deviceChange.device.port.switchUuid"}
		}
	}

	return nil
}

func (vm *VirtualMachine) configureDevice(
	ctx *Context,
	devices object.VirtualDeviceList,
	spec *types.VirtualDeviceConfigSpec,
	oldDevice types.BaseVirtualDevice) types.BaseMethodFault {

	device := spec.Device
	d := device.GetVirtualDevice()
	var controller types.BaseVirtualController

	key := d.Key
	if d.Key <= 0 {
		// Keys can't be negative; Key 0 is reserved
		d.Key = devices.NewKey()
		d.Key *= -1
	}

	// Update device controller's key reference
	if key != d.Key {
		if device := devices.FindByKey(d.ControllerKey); device != nil {
			if c, ok := device.(types.BaseVirtualController); ok {
				c := c.GetVirtualController()
				for i := range c.Device {
					if c.Device[i] == key {
						c.Device[i] = d.Key
						break
					}
				}
			}
		}
	}

	// Choose a unique key
	for {
		if devices.FindByKey(d.Key) == nil {
			break
		}
		d.Key++
	}

	label := devices.Name(device)
	summary := label
	dc := ctx.Map.getEntityDatacenter(ctx.Map.Get(*vm.Parent).(mo.Entity))

	switch x := device.(type) {
	case types.BaseVirtualEthernetCard:
		controller = devices.PickController((*types.VirtualPCIController)(nil))
		var net types.ManagedObjectReference
		var name string

		if b, ok := d.Backing.(*types.VirtualEthernetCardOpaqueNetworkBackingInfo); ok &&
			b.OpaqueNetworkType == "nsx.LogicalSwitch" {

			// For NSX opaque networks, replace the backing with the actual DVPG.
			var dvpg *DistributedVirtualPortgroup

			var find func(types.ManagedObjectReference)
			find = func(child types.ManagedObjectReference) {
				d, ok := ctx.Map.Get(child).(*DistributedVirtualPortgroup)
				if ok && d.Config.LogicalSwitchUuid == b.OpaqueNetworkId {
					dvpg = d
					return
				}
				walk(ctx.Map.Get(child), find)
			}
			f := ctx.Map.getEntityDatacenter(vm).NetworkFolder
			walk(ctx.Map.Get(f), find) // search in NetworkFolder and any sub folders

			if dvpg == nil {
				log.Printf("DPVG for NSX LogicalSwitch %s cannot be found", b.OpaqueNetworkId)
				fault := new(types.NotFound)
				fault.FaultMessage = []types.LocalizableMessage{
					{
						Key: "com.vmware.nsx.attachFailed",
						Message: fmt.Sprintf("The operation failed due to An error occurred during host configuration: "+
							"Failed to attach VIF: The requested object : LogicalSwitch/%s could not be found. Object identifiers are case sensitive.", b.OpaqueNetworkId),
					},
				}
				return fault
			}

			dvs := ctx.Map.Get(*dvpg.Config.DistributedVirtualSwitch).(*DistributedVirtualSwitch)

			d.Backing = &types.VirtualEthernetCardDistributedVirtualPortBackingInfo{
				Port: types.DistributedVirtualSwitchPortConnection{
					PortgroupKey: dvpg.Key,
					SwitchUuid:   dvs.Uuid,
				},
			}
		}

		switch b := d.Backing.(type) {
		case *types.VirtualEthernetCardNetworkBackingInfo:
			name = b.DeviceName
			summary = name
			net = ctx.Map.FindByName(b.DeviceName, dc.Network).Reference()
			b.Network = &net
		case *types.VirtualEthernetCardDistributedVirtualPortBackingInfo:
			summary = fmt.Sprintf("DVSwitch: %s", b.Port.SwitchUuid)
			net.Type = "DistributedVirtualPortgroup"
			net.Value = b.Port.PortgroupKey
			if err := vm.validateSwitchMembers(ctx, b.Port.SwitchUuid); err != nil {
				return err
			}
		}

		ctx.Update(vm, []types.PropertyChange{
			{Name: "summary.config.numEthernetCards", Val: vm.Summary.Config.NumEthernetCards + 1},
			{Name: "network", Val: append(vm.Network, net)},
		})

		c := x.GetVirtualEthernetCard()
		if c.MacAddress == "" {
			if c.UnitNumber == nil {
				devices.AssignController(device, controller)
			}
			c.MacAddress = vm.generateMAC(*c.UnitNumber - 7) // Note 7 == PCI offset
		}

		vm.Guest.Net = append(vm.Guest.Net, types.GuestNicInfo{
			Network:        name,
			IpAddress:      nil,
			MacAddress:     c.MacAddress,
			Connected:      true,
			DeviceConfigId: c.Key,
		})

		if spec.Operation == types.VirtualDeviceConfigSpecOperationAdd {
			if c.ResourceAllocation == nil {
				c.ResourceAllocation = &types.VirtualEthernetCardResourceAllocation{
					Reservation: types.NewInt64(0),
					Share: types.SharesInfo{
						Shares: 50,
						Level:  "normal",
					},
					Limit: types.NewInt64(-1),
				}
			}
		}
	case *types.VirtualDisk:
		if oldDevice == nil {
			// NOTE: either of capacityInBytes and capacityInKB may not be specified
			x.CapacityInBytes = getDiskSize(x)
			x.CapacityInKB = getDiskSize(x) / 1024
		} else {
			if oldDisk, ok := oldDevice.(*types.VirtualDisk); ok {
				diskSize, ok := changedDiskSize(oldDisk, x)
				if !ok {
					return &types.InvalidDeviceOperation{}
				}
				x.CapacityInBytes = diskSize
				x.CapacityInKB = diskSize / 1024
			}
		}

		getCryptoKeyID := func(desc *vmdk.Descriptor) *types.CryptoKeyId {
			if desc == nil {
				return nil
			}
			if ek := desc.EncryptionKeys; ek != nil {
				if l := ek.List; len(l) > 0 {
					if p := l[0].Pair; p != nil {
						if l := p.Locker; l != nil {
							if i := l.Indirect; i != nil {
								var (
									keyID      = i.FQID.KeyID
									providerID = i.FQID.KeyServerID
								)

								if keyID == "" && providerID == "" {
									return nil
								}

								cki := &types.CryptoKeyId{
									KeyId: i.FQID.KeyID,
								}

								if providerID != "" {
									cki.ProviderId = &types.KeyProviderId{
										Id: providerID,
									}
								}

								return cki
							}
						}
					}
				}
			}
			return nil
		}

		var setCryptoKeyID func(desc *vmdk.Descriptor)

		switch b := d.Backing.(type) {
		case *types.VirtualDiskSparseVer2BackingInfo:
			// Sparse disk creation not supported in ESX
			return &types.DeviceUnsupportedForVmPlatform{
				InvalidDeviceSpec: types.InvalidDeviceSpec{
					InvalidVmConfig: types.InvalidVmConfig{Property: "VirtualDeviceSpec.device.backing"},
				},
			}
		case types.BaseVirtualDeviceFileBackingInfo:
			var (
				parent string
				crypto types.BaseCryptoSpec
			)

			switch backing := d.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				if backing.Parent != nil {
					parent = backing.Parent.FileName
				}
				if spec.Backing != nil {
					crypto = spec.Backing.Crypto
					switch tCrypto := crypto.(type) {
					case *types.CryptoSpecEncrypt:
						backing.KeyId = &tCrypto.CryptoKeyId
					case *types.CryptoSpecShallowRecrypt:
						backing.KeyId = &tCrypto.NewKeyId
					case *types.CryptoSpecDeepRecrypt:
						backing.KeyId = &tCrypto.NewKeyId
					case *types.CryptoSpecDecrypt:
						backing.KeyId = nil
					}
				}
				setCryptoKeyID = func(desc *vmdk.Descriptor) {
					backing.KeyId = getCryptoKeyID(desc)
				}
			case *types.VirtualDiskSeSparseBackingInfo:
				if backing.Parent != nil {
					parent = backing.Parent.FileName
				}
				if spec.Backing != nil {
					crypto = spec.Backing.Crypto
					switch tCrypto := crypto.(type) {
					case *types.CryptoSpecEncrypt:
						backing.KeyId = &tCrypto.CryptoKeyId
					case *types.CryptoSpecShallowRecrypt:
						backing.KeyId = &tCrypto.NewKeyId
					case *types.CryptoSpecDeepRecrypt:
						backing.KeyId = &tCrypto.NewKeyId
					case *types.CryptoSpecDecrypt:
						backing.KeyId = nil
					}
				}
				setCryptoKeyID = func(desc *vmdk.Descriptor) {
					backing.KeyId = getCryptoKeyID(desc)
				}
			case *types.VirtualDiskSparseVer2BackingInfo:
				if backing.Parent != nil {
					parent = backing.Parent.FileName
				}
				if spec.Backing != nil {
					crypto = spec.Backing.Crypto
					switch tCrypto := crypto.(type) {
					case *types.CryptoSpecEncrypt:
						backing.KeyId = &tCrypto.CryptoKeyId
					case *types.CryptoSpecShallowRecrypt:
						backing.KeyId = &tCrypto.NewKeyId
					case *types.CryptoSpecDeepRecrypt:
						backing.KeyId = &tCrypto.NewKeyId
					case *types.CryptoSpecDecrypt:
						backing.KeyId = nil
					}
				}
				setCryptoKeyID = func(desc *vmdk.Descriptor) {
					backing.KeyId = getCryptoKeyID(desc)
				}
			}

			if parent != "" {
				desc, _, err := ctx.Map.FileManager().DiskDescriptor(ctx, &dc.Self, parent)
				if err != nil {
					return &types.InvalidDeviceSpec{
						InvalidVmConfig: types.InvalidVmConfig{
							Property: "virtualDeviceSpec.device.backing.parent.fileName",
						},
					}
				}

				// Disk Capacity is always same as the parent's
				x.CapacityInBytes = int64(desc.Capacity())
				x.CapacityInKB = x.CapacityInBytes / 1024
			}

			summary = fmt.Sprintf("%s KB", numberToString(x.CapacityInKB, ','))

			info := b.GetVirtualDeviceFileBackingInfo()
			var path object.DatastorePath
			path.FromString(info.FileName)

			if path.Path == "" {
				filename, err := vm.genVmdkPath(ctx, path)
				if err != nil {
					return err
				}

				info.FileName = filename
			}

			desc, err := vdmCreateVirtualDisk(ctx, spec.FileOperation, &types.CreateVirtualDisk_Task{
				Datacenter: &dc.Self,
				Name:       info.FileName,
				Spec: &types.FileBackedVirtualDiskSpec{
					CapacityKb: x.CapacityInKB,
					Crypto:     crypto,
				},
			})
			if err != nil {
				return err
			}

			if desc != nil && setCryptoKeyID != nil {
				setCryptoKeyID(desc)
			}

			ctx.Update(vm, []types.PropertyChange{
				{Name: "summary.config.numVirtualDisks", Val: vm.Summary.Config.NumVirtualDisks + 1},
			})

			p, _ := parseDatastorePath(info.FileName)
			ds := vm.findDatastore(ctx, p.Datastore)
			info.Datastore = &ds.Self

			if oldDevice != nil {
				if oldDisk, ok := oldDevice.(*types.VirtualDisk); ok {
					// add previous capacity to datastore freespace
					ctx.WithLock(ds, func() {
						ds.Summary.FreeSpace += getDiskSize(oldDisk)
						ds.Info.GetDatastoreInfo().FreeSpace = ds.Summary.FreeSpace
					})
				}
			}

			// then subtract new capacity from datastore freespace
			// XXX: compare disk size and free space until windows stat is supported
			ctx.WithLock(ds, func() {
				ds.Summary.FreeSpace -= getDiskSize(x)
				ds.Info.GetDatastoreInfo().FreeSpace = ds.Summary.FreeSpace
			})

			vm.updateDiskLayouts(ctx)

			if disk, ok := b.(*types.VirtualDiskFlatVer2BackingInfo); ok {
				// These properties default to false
				props := []**bool{
					&disk.EagerlyScrub,
					&disk.ThinProvisioned,
					&disk.WriteThrough,
					&disk.Split,
					&disk.DigestEnabled,
				}
				for _, prop := range props {
					if *prop == nil {
						*prop = types.NewBool(false)
					}
				}

				if disk.Uuid == "" {
					disk.Uuid = virtualDiskUUID(&dc.Self, info.FileName)
				}
			}
		}
	case *types.VirtualCdrom:
		if b, ok := d.Backing.(types.BaseVirtualDeviceFileBackingInfo); ok {
			summary = "ISO " + b.GetVirtualDeviceFileBackingInfo().FileName
		}
	case *types.VirtualFloppy:
		if b, ok := d.Backing.(types.BaseVirtualDeviceFileBackingInfo); ok {
			summary = "Image " + b.GetVirtualDeviceFileBackingInfo().FileName
		}
	case *types.VirtualSerialPort:
		switch b := d.Backing.(type) {
		case types.BaseVirtualDeviceFileBackingInfo:
			summary = "File " + b.GetVirtualDeviceFileBackingInfo().FileName
		case *types.VirtualSerialPortURIBackingInfo:
			summary = "Remote " + b.ServiceURI
		}
	}

	if d.UnitNumber == nil && controller != nil {
		devices.AssignController(device, controller)
	}

	if d.DeviceInfo == nil {
		d.DeviceInfo = &types.Description{
			Label:   label,
			Summary: summary,
		}
	} else {
		info := d.DeviceInfo.GetDescription()
		if info.Label == "" {
			info.Label = label
		}
		if info.Summary == "" {
			info.Summary = summary
		}
	}

	switch device.(type) {
	case types.BaseVirtualEthernetCard, *types.VirtualCdrom, *types.VirtualFloppy, *types.VirtualUSB, *types.VirtualSerialPort:
		if d.Connectable == nil {
			d.Connectable = &types.VirtualDeviceConnectInfo{StartConnected: true, Connected: true}
		}
	}

	// device can be connected only if vm is powered on
	if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
		if d.Connectable != nil {
			d.Connectable.Connected = false
		}
	}

	return nil
}

func (vm *VirtualMachine) removeDevice(ctx *Context, devices object.VirtualDeviceList, spec *types.VirtualDeviceConfigSpec) object.VirtualDeviceList {
	key := spec.Device.GetVirtualDevice().Key

	for i, d := range devices {
		if d.GetVirtualDevice().Key != key {
			continue
		}

		devices = append(devices[:i], devices[i+1:]...)

		switch device := spec.Device.(type) {
		case *types.VirtualDisk:
			if spec.FileOperation == types.VirtualDeviceConfigSpecFileOperationDestroy {
				var file string

				switch b := device.Backing.(type) {
				case types.BaseVirtualDeviceFileBackingInfo:
					file = b.GetVirtualDeviceFileBackingInfo().FileName

					p, _ := parseDatastorePath(file)
					ds := vm.findDatastore(ctx, p.Datastore)

					ctx.WithLock(ds, func() {
						ds.Summary.FreeSpace += getDiskSize(device)
						ds.Info.GetDatastoreInfo().FreeSpace = ds.Summary.FreeSpace
					})
				}

				if file != "" {
					dc := ctx.Map.getEntityDatacenter(vm)
					dm := ctx.Map.VirtualDiskManager()
					if dc == nil {
						continue // parent was destroyed
					}
					res := dm.DeleteVirtualDiskTask(ctx, &types.DeleteVirtualDisk_Task{
						Name:       file,
						Datacenter: &dc.Self,
					})
					ctask := ctx.Map.Get(res.(*methods.DeleteVirtualDisk_TaskBody).Res.Returnval).(*Task)
					ctask.Wait()
				}
			}
			ctx.Update(vm, []types.PropertyChange{
				{Name: "summary.config.numVirtualDisks", Val: vm.Summary.Config.NumVirtualDisks - 1},
			})

			vm.RefreshStorageInfo(ctx, nil)
		case types.BaseVirtualEthernetCard:
			var net types.ManagedObjectReference

			switch b := device.GetVirtualEthernetCard().Backing.(type) {
			case *types.VirtualEthernetCardNetworkBackingInfo:
				net = *b.Network
			case *types.VirtualEthernetCardDistributedVirtualPortBackingInfo:
				net.Type = "DistributedVirtualPortgroup"
				net.Value = b.Port.PortgroupKey
			}

			for j, nicInfo := range vm.Guest.Net {
				if nicInfo.DeviceConfigId == key {
					vm.Guest.Net = append(vm.Guest.Net[:j], vm.Guest.Net[j+1:]...)
					break
				}
			}

			networks := vm.Network
			RemoveReference(&networks, net)
			ctx.Update(vm, []types.PropertyChange{
				{Name: "summary.config.numEthernetCards", Val: vm.Summary.Config.NumEthernetCards - 1},
				{Name: "network", Val: networks},
			})
		}

		break
	}

	return devices
}

func (vm *VirtualMachine) genVmdkPath(ctx *Context, p object.DatastorePath) (string, types.BaseMethodFault) {
	if p.Datastore == "" {
		p.FromString(vm.Config.Files.VmPathName)
	}
	if p.Path == "" {
		p.Path = vm.Config.Name
	} else {
		p.Path = path.Dir(p.Path)
	}
	vmdir := p.String()
	index := 0
	for {
		var filename string
		if index == 0 {
			filename = fmt.Sprintf("%s.vmdk", vm.Config.Name)
		} else {
			filename = fmt.Sprintf("%s_%d.vmdk", vm.Config.Name, index)
		}

		f, err := vm.createFile(ctx, vmdir, filename, false)
		if err != nil {
			switch err.(type) {
			case *types.FileAlreadyExists:
				index++
				continue
			default:
				return "", err
			}
		}

		_ = f.Close()
		_ = os.Remove(f.Name())

		return path.Join(vmdir, filename), nil
	}
}

// Encrypt requires powered off VM with no snapshots.
// Decrypt requires powered off VM.
// Deep recrypt requires powered off VM with no snapshots.
// Shallow recrypt works with VMs in any power state and even if snapshots are
// present as long as it is a single chain and not a tree.
func (vm *VirtualMachine) updateCrypto(
	ctx *Context,
	spec types.BaseCryptoSpec) types.BaseMethodFault {

	const configKeyId = "config.keyId"

	assertEncrypted := func() types.BaseMethodFault {
		if vm.Config.KeyId == nil {
			return newInvalidStateFault("vm is not encrypted")
		}
		return nil
	}

	assertPoweredOff := func() types.BaseMethodFault {
		if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOff {
			return &types.InvalidPowerState{
				ExistingState:  vm.Runtime.PowerState,
				RequestedState: types.VirtualMachinePowerStatePoweredOff,
			}
		}
		return nil
	}

	assertNoSnapshots := func(allowSingleChain bool) types.BaseMethodFault {
		hasSnapshots := vm.Snapshot != nil && vm.Snapshot.CurrentSnapshot != nil
		if !hasSnapshots {
			return nil
		}
		if !allowSingleChain {
			return newInvalidStateFault("vm has snapshots")
		}
		type node = types.VirtualMachineSnapshotTree
		var isTreeFn func(nodes []node) types.BaseMethodFault
		isTreeFn = func(nodes []node) types.BaseMethodFault {
			switch len(nodes) {
			case 0:
				return nil
			case 1:
				return isTreeFn(nodes[0].ChildSnapshotList)
			default:
				return newInvalidStateFault("vm has snapshot tree")
			}
		}
		return isTreeFn(vm.Snapshot.RootSnapshotList)
	}

	doRecrypt := func(newKeyID types.CryptoKeyId) types.BaseMethodFault {
		if err := assertEncrypted(); err != nil {
			return err
		}

		var providerID *types.KeyProviderId
		if pid := newKeyID.ProviderId; pid != nil {
			providerID = &types.KeyProviderId{
				Id: pid.Id,
			}
		}

		keyID := newKeyID.KeyId
		if providerID == nil {
			if p, k := getDefaultProvider(ctx, vm, true); p != "" && k != "" {
				providerID = &types.KeyProviderId{
					Id: p,
				}
				keyID = k
			}
		} else if keyID == "" {
			keyID = generateKeyForProvider(ctx, providerID.Id)
		}

		ctx.Update(vm, []types.PropertyChange{
			{
				Name: configKeyId,
				Op:   types.PropertyChangeOpAssign,
				Val: &types.CryptoKeyId{
					KeyId:      keyID,
					ProviderId: providerID,
				},
			},
		})
		return nil
	}

	switch tspec := spec.(type) {
	case *types.CryptoSpecDecrypt:
		if err := assertPoweredOff(); err != nil {
			return err
		}
		if err := assertNoSnapshots(false); err != nil {
			return err
		}
		if err := assertEncrypted(); err != nil {
			return err
		}
		ctx.Update(vm, []types.PropertyChange{
			{
				Name: configKeyId,
				Op:   types.PropertyChangeOpRemove,
				Val:  nil,
			},
		})

	case *types.CryptoSpecDeepRecrypt:
		if err := assertPoweredOff(); err != nil {
			return err
		}
		if err := assertNoSnapshots(false); err != nil {
			return err
		}
		return doRecrypt(tspec.NewKeyId)

	case *types.CryptoSpecShallowRecrypt:
		if err := assertNoSnapshots(true); err != nil {
			return err
		}
		return doRecrypt(tspec.NewKeyId)

	case *types.CryptoSpecEncrypt:
		if err := assertPoweredOff(); err != nil {
			return err
		}
		if err := assertNoSnapshots(false); err != nil {
			return err
		}
		if vm.Config.KeyId != nil {
			return newInvalidStateFault("vm is already encrypted")
		}

		var providerID *types.KeyProviderId
		if pid := tspec.CryptoKeyId.ProviderId; pid != nil {
			providerID = &types.KeyProviderId{
				Id: pid.Id,
			}
		}

		keyID := tspec.CryptoKeyId.KeyId
		if providerID == nil {
			if p, k := getDefaultProvider(ctx, vm, true); p != "" && k != "" {
				providerID = &types.KeyProviderId{
					Id: p,
				}
				keyID = k
			}
		} else if keyID == "" {
			keyID = generateKeyForProvider(ctx, providerID.Id)
		}

		ctx.Update(vm, []types.PropertyChange{
			{
				Name: configKeyId,
				Op:   types.PropertyChangeOpAssign,
				Val: &types.CryptoKeyId{
					KeyId:      keyID,
					ProviderId: providerID,
				},
			},
		})

	case *types.CryptoSpecNoOp,
		*types.CryptoSpecRegister:

		// No-op
	}

	return nil
}

func (vm *VirtualMachine) configureDevices(ctx *Context, spec *types.VirtualMachineConfigSpec) types.BaseMethodFault {
	var changes []types.PropertyChange
	field := mo.Field{Path: "config.hardware.device"}
	devices := object.VirtualDeviceList(vm.Config.Hardware.Device)

	var err types.BaseMethodFault
	for i, change := range spec.DeviceChange {
		dspec := change.GetVirtualDeviceConfigSpec()
		device := dspec.Device.GetVirtualDevice()
		invalid := &types.InvalidDeviceSpec{DeviceIndex: int32(i)}
		change := types.PropertyChange{}

		switch dspec.FileOperation {
		case types.VirtualDeviceConfigSpecFileOperationCreate:
			switch dspec.Device.(type) {
			case *types.VirtualDisk:
				if device.UnitNumber == nil {
					return invalid
				}
			}
		}

		switch dspec.Operation {
		case types.VirtualDeviceConfigSpecOperationAdd:
			change.Op = types.PropertyChangeOpAdd

			if devices.FindByKey(device.Key) != nil && device.ControllerKey == 0 {
				// Note: real ESX does not allow adding base controllers (ControllerKey = 0)
				// after VM is created (returns success but device is not added).
				continue
			} else if device.UnitNumber != nil && devices.SelectByType(dspec.Device).Select(func(d types.BaseVirtualDevice) bool {
				base := d.GetVirtualDevice()
				if base.UnitNumber != nil {
					if base.ControllerKey != device.ControllerKey {
						return false
					}
					return *base.UnitNumber == *device.UnitNumber
				}
				return false
			}) != nil {
				// UnitNumber for this device type is taken
				return invalid
			}

			key := device.Key
			err = vm.configureDevice(ctx, devices, dspec, nil)
			if err != nil {
				return err
			}

			devices = append(devices, dspec.Device)
			change.Val = dspec.Device
			if key != device.Key {
				// Update ControllerKey refs
				for i := range spec.DeviceChange {
					ckey := &spec.DeviceChange[i].GetVirtualDeviceConfigSpec().Device.GetVirtualDevice().ControllerKey
					if *ckey == key {
						*ckey = device.Key
					}
				}
			}
		case types.VirtualDeviceConfigSpecOperationEdit:
			rspec := *dspec
			oldDevice := devices.FindByKey(device.Key)
			if oldDevice == nil {
				return invalid
			}
			rspec.Device = oldDevice
			devices = vm.removeDevice(ctx, devices, &rspec)
			if device.DeviceInfo != nil {
				device.DeviceInfo.GetDescription().Summary = "" // regenerate summary
			}

			err = vm.configureDevice(ctx, devices, dspec, oldDevice)
			if err != nil {
				return err
			}

			devices = append(devices, dspec.Device)
			change.Val = dspec.Device
		case types.VirtualDeviceConfigSpecOperationRemove:
			change.Op = types.PropertyChangeOpRemove

			devices = vm.removeDevice(ctx, devices, dspec)
		}

		field.Key = device.Key
		change.Name = field.String()
		changes = append(changes, change)
	}

	if len(changes) != 0 {
		change := types.PropertyChange{Name: field.Path, Val: []types.BaseVirtualDevice(devices)}
		ctx.Update(vm, append(changes, change))
	}

	err = vm.updateDiskLayouts(ctx)
	if err != nil {
		return err
	}

	// Do this after device config, as some may apply to the devices themselves (e.g. ethernet -> guest.net)
	err = vm.applyExtraConfig(ctx, spec)
	if err != nil {
		return err
	}

	return nil
}

type powerVMTask struct {
	*VirtualMachine

	state types.VirtualMachinePowerState
	ctx   *Context
}

func (c *powerVMTask) Run(task *Task) (types.AnyType, types.BaseMethodFault) {
	c.logPrintf("running power task: requesting %s, existing %s",
		c.state, c.VirtualMachine.Runtime.PowerState)

	if c.VirtualMachine.Runtime.PowerState == c.state {
		return nil, &types.InvalidPowerState{
			RequestedState: c.state,
			ExistingState:  c.VirtualMachine.Runtime.PowerState,
		}
	}

	var boot types.AnyType
	if c.state == types.VirtualMachinePowerStatePoweredOn {
		boot = time.Now()
	}

	event := c.event(c.ctx)
	switch c.state {
	case types.VirtualMachinePowerStatePoweredOn:
		if c.VirtualMachine.hostInMM(c.ctx) {
			return nil, new(types.InvalidState)
		}

		err := c.svm.start(c.ctx)
		if err != nil {
			return nil, &types.MissingPowerOnConfiguration{
				VAppConfigFault: types.VAppConfigFault{
					VimFault: types.VimFault{
						MethodFault: types.MethodFault{
							FaultCause: &types.LocalizedMethodFault{
								Fault:            &types.SystemErrorFault{Reason: err.Error()},
								LocalizedMessage: err.Error()}}}}}
		}
		c.ctx.postEvent(
			&types.VmStartingEvent{VmEvent: event},
			&types.VmPoweredOnEvent{VmEvent: event},
		)
		c.customize(c.ctx)
	case types.VirtualMachinePowerStatePoweredOff:
		c.svm.stop(c.ctx)
		c.ctx.postEvent(
			&types.VmStoppingEvent{VmEvent: event},
			&types.VmPoweredOffEvent{VmEvent: event},
		)
	case types.VirtualMachinePowerStateSuspended:
		if c.VirtualMachine.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
			return nil, &types.InvalidPowerState{
				RequestedState: types.VirtualMachinePowerStatePoweredOn,
				ExistingState:  c.VirtualMachine.Runtime.PowerState,
			}
		}

		c.svm.pause(c.ctx)
		c.ctx.postEvent(
			&types.VmSuspendingEvent{VmEvent: event},
			&types.VmSuspendedEvent{VmEvent: event},
		)
	}

	// copy devices to prevent data race
	devices := c.VirtualMachine.cloneDevice()
	for _, d := range devices {
		conn := d.GetVirtualDevice().Connectable
		if conn == nil {
			continue
		}

		if c.state == types.VirtualMachinePowerStatePoweredOn {
			// apply startConnected to current connection
			conn.Connected = conn.StartConnected
		} else {
			conn.Connected = false
		}
	}

	c.ctx.Update(c.VirtualMachine, []types.PropertyChange{
		{Name: "runtime.powerState", Val: c.state},
		{Name: "summary.runtime.powerState", Val: c.state},
		{Name: "summary.runtime.bootTime", Val: boot},
		{Name: "config.hardware.device", Val: devices},
	})

	return nil, nil
}

func (vm *VirtualMachine) PowerOnVMTask(ctx *Context, c *types.PowerOnVM_Task) soap.HasFault {
	if vm.Config.Template {
		return &methods.PowerOnVM_TaskBody{
			Fault_: Fault("cannot powerOn a template", &types.InvalidState{}),
		}
	}

	runner := &powerVMTask{vm, types.VirtualMachinePowerStatePoweredOn, ctx}
	task := CreateTask(runner.Reference(), "powerOn", runner.Run)

	return &methods.PowerOnVM_TaskBody{
		Res: &types.PowerOnVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) PowerOffVMTask(ctx *Context, c *types.PowerOffVM_Task) soap.HasFault {
	runner := &powerVMTask{vm, types.VirtualMachinePowerStatePoweredOff, ctx}
	task := CreateTask(runner.Reference(), "powerOff", runner.Run)

	return &methods.PowerOffVM_TaskBody{
		Res: &types.PowerOffVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) SuspendVMTask(ctx *Context, req *types.SuspendVM_Task) soap.HasFault {
	runner := &powerVMTask{vm, types.VirtualMachinePowerStateSuspended, ctx}
	task := CreateTask(runner.Reference(), "suspend", runner.Run)

	return &methods.SuspendVM_TaskBody{
		Res: &types.SuspendVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) ResetVMTask(ctx *Context, req *types.ResetVM_Task) soap.HasFault {
	task := CreateTask(vm, "reset", func(task *Task) (types.AnyType, types.BaseMethodFault) {
		res := vm.PowerOffVMTask(ctx, &types.PowerOffVM_Task{This: vm.Self})
		ctask := ctx.Map.Get(res.(*methods.PowerOffVM_TaskBody).Res.Returnval).(*Task)
		ctask.Wait()
		if ctask.Info.Error != nil {
			return nil, ctask.Info.Error.Fault
		}

		res = vm.PowerOnVMTask(ctx, &types.PowerOnVM_Task{This: vm.Self})
		ctask = ctx.Map.Get(res.(*methods.PowerOnVM_TaskBody).Res.Returnval).(*Task)
		ctask.Wait()

		return nil, nil
	})

	return &methods.ResetVM_TaskBody{
		Res: &types.ResetVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) RebootGuest(ctx *Context, req *types.RebootGuest) soap.HasFault {
	body := new(methods.RebootGuestBody)

	if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
		body.Fault_ = Fault("", &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOn,
			ExistingState:  vm.Runtime.PowerState,
		})
		return body
	}

	if vm.Guest.ToolsRunningStatus == string(types.VirtualMachineToolsRunningStatusGuestToolsRunning) {
		vm.svm.restart(ctx)
		body.Res = new(types.RebootGuestResponse)
	} else {
		body.Fault_ = Fault("", new(types.ToolsUnavailable))
	}

	return body
}

func (vm *VirtualMachine) ReconfigVMTask(ctx *Context, req *types.ReconfigVM_Task) soap.HasFault {
	task := CreateTask(vm, "reconfigVm", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		ctx.postEvent(&types.VmReconfiguredEvent{
			VmEvent:    vm.event(ctx),
			ConfigSpec: req.Spec,
		})

		if vm.Config.Template {
			expect := types.VirtualMachineConfigSpec{
				Name:       req.Spec.Name,
				Annotation: req.Spec.Annotation,
			}
			if !reflect.DeepEqual(&req.Spec, &expect) {
				log.Printf("template reconfigure only allows name and annotation change")
				return nil, new(types.NotSupported)
			}
		}

		err := vm.configure(ctx, &req.Spec)

		return nil, err
	})

	return &methods.ReconfigVM_TaskBody{
		Res: &types.ReconfigVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) UpgradeVMTask(ctx *Context, req *types.UpgradeVM_Task) soap.HasFault {
	body := &methods.UpgradeVM_TaskBody{}

	task := CreateTask(vm, "upgradeVm", func(t *Task) (types.AnyType, types.BaseMethodFault) {

		// InvalidPowerState
		//
		// 1. Is VM's power state anything other than powered off?
		if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOff {
			return nil, &types.InvalidPowerStateFault{
				ExistingState:  vm.Runtime.PowerState,
				RequestedState: types.VirtualMachinePowerStatePoweredOff,
			}
		}

		// InvalidState
		//
		// 1. Is host on which VM is scheduled in maintenance mode?
		// 2. Is VM a template?
		// 3. Is VM already the latest hardware version?
		var (
			ebRef                     *types.ManagedObjectReference
			latestHardwareVersion     string
			hostRef                   = vm.Runtime.Host
			supportedHardwareVersions = map[string]struct{}{}
			vmHardwareVersionString   = vm.Config.Version
		)
		if hostRef != nil {
			var hostInMaintenanceMode bool
			ctx.WithLock(*hostRef, func() {
				host := ctx.Map.Get(*hostRef).(*HostSystem)
				hostInMaintenanceMode = host.Runtime.InMaintenanceMode
				switch host.Parent.Type {
				case "ClusterComputeResource":
					obj := ctx.Map.Get(*host.Parent).(*ClusterComputeResource)
					ebRef = obj.EnvironmentBrowser
				case "ComputeResource":
					obj := ctx.Map.Get(*host.Parent).(*mo.ComputeResource)
					ebRef = obj.EnvironmentBrowser
				}
			})
			if hostInMaintenanceMode {
				return nil, newInvalidStateFault("%s in maintenance mode", hostRef.Value)
			}
		}
		if vm.Config.Template {
			return nil, newInvalidStateFault("%s is template", vm.Reference().Value)
		}
		if ebRef != nil {
			ctx.WithLock(*ebRef, func() {
				eb := ctx.Map.Get(*ebRef).(*EnvironmentBrowser)
				for i := range eb.QueryConfigOptionDescriptorResponse.Returnval {
					cod := eb.QueryConfigOptionDescriptorResponse.Returnval[i]
					for j := range cod.Host {
						if cod.Host[j].Value == hostRef.Value {
							supportedHardwareVersions[cod.Key] = struct{}{}
						}
						if latestHardwareVersion == "" {
							if def := cod.DefaultConfigOption; def {
								latestHardwareVersion = cod.Key
							}
						}
					}
				}
			})
		}

		if latestHardwareVersion == "" {
			latestHardwareVersion = esx.HardwareVersion
		}
		if vmHardwareVersionString == latestHardwareVersion {
			return nil, newInvalidStateFault("%s is latest version", vm.Reference().Value)
		}
		if req.Version == "" {
			req.Version = latestHardwareVersion
		}

		// NotSupported
		targetVersion, _ := types.ParseHardwareVersion(req.Version)
		if targetVersion.IsValid() {
			req.Version = targetVersion.String()
		}
		if _, ok := supportedHardwareVersions[req.Version]; !ok {
			msg := fmt.Sprintf("%s not supported", req.Version)
			return nil, &types.NotSupported{
				RuntimeFault: types.RuntimeFault{
					MethodFault: types.MethodFault{
						FaultCause: &types.LocalizedMethodFault{
							Fault: &types.SystemErrorFault{
								Reason: msg,
							},
							LocalizedMessage: msg,
						},
					},
				},
			}
		}

		// AlreadyUpgraded
		vmHardwareVersion, _ := types.ParseHardwareVersion(vmHardwareVersionString)
		if targetVersion.IsValid() && vmHardwareVersion.IsValid() &&
			targetVersion <= vmHardwareVersion {

			return nil, &types.AlreadyUpgradedFault{}
		}

		// InvalidArgument
		if targetVersion < types.VMX3 {
			return nil, &types.InvalidArgument{}
		}

		ctx.Update(vm, []types.PropertyChange{
			{
				Name: "config.version", Val: targetVersion.String(),
			},
			{
				Name: "summary.config.hwVersion", Val: targetVersion.String(),
			},
		})

		return nil, nil
	})

	body.Res = &types.UpgradeVM_TaskResponse{
		Returnval: task.Run(ctx),
	}

	return body
}

func (vm *VirtualMachine) DestroyTask(ctx *Context, req *types.Destroy_Task) soap.HasFault {
	dc := ctx.Map.getEntityDatacenter(vm)

	task := CreateTask(vm, "destroy", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		if dc == nil {
			return nil, &types.ManagedObjectNotFound{Obj: vm.Self} // If our Parent was destroyed, so were we.
			// TODO: should this also trigger container removal?
		}

		r := vm.UnregisterVM(ctx, &types.UnregisterVM{
			This: req.This,
		})

		if r.Fault() != nil {
			return nil, r.Fault().VimFault().(types.BaseMethodFault)
		}

		// Remove all devices
		devices := object.VirtualDeviceList(vm.Config.Hardware.Device)
		spec, _ := devices.ConfigSpec(types.VirtualDeviceConfigSpecOperationRemove)
		vm.configureDevices(ctx, &types.VirtualMachineConfigSpec{DeviceChange: spec})

		// Delete VM files from the datastore (ignoring result for now)
		m := ctx.Map.FileManager()

		_ = m.DeleteDatastoreFileTask(ctx, &types.DeleteDatastoreFile_Task{
			This:       m.Reference(),
			Name:       vm.Config.Files.LogDirectory,
			Datacenter: &dc.Self,
		})

		err := vm.svm.remove(ctx)
		if err != nil {
			return nil, &types.RuntimeFault{
				MethodFault: types.MethodFault{
					FaultCause: &types.LocalizedMethodFault{
						Fault:            &types.SystemErrorFault{Reason: err.Error()},
						LocalizedMessage: err.Error()}}}
		}

		return nil, nil
	})

	return &methods.Destroy_TaskBody{
		Res: &types.Destroy_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) SetCustomValue(ctx *Context, req *types.SetCustomValue) soap.HasFault {
	return SetCustomValue(ctx, req)
}

func (vm *VirtualMachine) UnregisterVM(ctx *Context, c *types.UnregisterVM) soap.HasFault {
	r := &methods.UnregisterVMBody{}

	if vm.Runtime.PowerState == types.VirtualMachinePowerStatePoweredOn {
		r.Fault_ = Fault("", &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOff,
			ExistingState:  vm.Runtime.PowerState,
		})

		return r
	}

	host := ctx.Map.Get(*vm.Runtime.Host).(*HostSystem)
	ctx.Map.RemoveReference(ctx, host, &host.Vm, vm.Self)

	if vm.ResourcePool != nil {
		switch pool := ctx.Map.Get(*vm.ResourcePool).(type) {
		case *ResourcePool:
			ctx.Map.RemoveReference(ctx, pool, &pool.Vm, vm.Self)
		case *VirtualApp:
			ctx.Map.RemoveReference(ctx, pool, &pool.Vm, vm.Self)
		}
	}

	for i := range vm.Datastore {
		ds := ctx.Map.Get(vm.Datastore[i]).(*Datastore)
		ctx.Map.RemoveReference(ctx, ds, &ds.Vm, vm.Self)
	}

	ctx.postEvent(&types.VmRemovedEvent{VmEvent: vm.event(ctx)})
	if f, ok := asFolderMO(ctx.Map.getEntityParent(vm, "Folder")); ok {
		folderRemoveChild(ctx, f, c.This)
	}

	r.Res = new(types.UnregisterVMResponse)

	return r
}

type vmFolder interface {
	CreateVMTask(ctx *Context, c *types.CreateVM_Task) soap.HasFault
}

func (vm *VirtualMachine) cloneDevice() []types.BaseVirtualDevice {
	src := types.ArrayOfVirtualDevice{
		VirtualDevice: vm.Config.Hardware.Device,
	}
	dst := types.ArrayOfVirtualDevice{}
	deepCopy(src, &dst)
	return dst.VirtualDevice
}

func (vm *VirtualMachine) worldID() int {
	return int(binary.BigEndian.Uint32(vm.uid[0:4]))
}

func (vm *VirtualMachine) CloneVMTask(ctx *Context, req *types.CloneVM_Task) soap.HasFault {
	pool := req.Spec.Location.Pool
	if pool == nil {
		if !vm.Config.Template {
			pool = vm.ResourcePool
		}
	}

	destHost := vm.Runtime.Host

	if req.Spec.Location.Host != nil {
		destHost = req.Spec.Location.Host
	}

	folder, _ := asFolderMO(ctx.Map.Get(req.Folder))
	host := ctx.Map.Get(*destHost).(*HostSystem)
	event := vm.event(ctx)

	ctx.postEvent(&types.VmBeingClonedEvent{
		VmCloneEvent: types.VmCloneEvent{
			VmEvent: event,
		},
		DestFolder: folderEventArgument(folder),
		DestName:   req.Name,
		DestHost:   *host.eventArgument(),
	})

	vmx := vm.vmx(nil)
	vmx.Path = req.Name
	if ref := req.Spec.Location.Datastore; ref != nil {
		ds := ctx.Map.Get(*ref).(*Datastore).Name
		vmx.Datastore = ds
	}

	task := CreateTask(vm, "cloneVm", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		if pool == nil {
			return nil, &types.InvalidArgument{InvalidProperty: "spec.location.pool"}
		}
		if obj := ctx.Map.FindByName(req.Name, folder.ChildEntity); obj != nil {
			return nil, &types.DuplicateName{
				Name:   req.Name,
				Object: obj.Reference(),
			}
		}
		config := types.VirtualMachineConfigSpec{
			Name:    req.Name,
			Version: vm.Config.Version,
			GuestId: vm.Config.GuestId,
			Files: &types.VirtualMachineFileInfo{
				VmPathName: vmx.String(),
			},
		}

		// Copying hardware properties
		config.NumCPUs = vm.Config.Hardware.NumCPU
		config.MemoryMB = int64(vm.Config.Hardware.MemoryMB)
		config.NumCoresPerSocket = vm.Config.Hardware.NumCoresPerSocket
		config.VirtualICH7MPresent = vm.Config.Hardware.VirtualICH7MPresent
		config.VirtualSMCPresent = vm.Config.Hardware.VirtualSMCPresent

		defaultDevices := object.VirtualDeviceList(esx.VirtualDevice)
		devices := vm.cloneDevice()

		for _, device := range devices {
			var fop types.VirtualDeviceConfigSpecFileOperation

			if defaultDevices.Find(object.VirtualDeviceList(devices).Name(device)) != nil {
				// Default devices are added during CreateVMTask
				continue
			}

			switch disk := device.(type) {
			case *types.VirtualDisk:
				// TODO: consider VirtualMachineCloneSpec.DiskMoveType
				fop = types.VirtualDeviceConfigSpecFileOperationCreate

				// Leave FileName empty so CreateVM will just create a new one under VmPathName
				disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo).FileName = ""
				disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo).Parent = nil
				// Clear UUID so a new unique UUID is generated for the cloned disk
				disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo).Uuid = ""
			}

			config.DeviceChange = append(config.DeviceChange, &types.VirtualDeviceConfigSpec{
				Operation:     types.VirtualDeviceConfigSpecOperationAdd,
				Device:        device,
				FileOperation: fop,
			})
		}

		if dst, src := config, req.Spec.Config; src != nil {
			dst.ExtraConfig = src.ExtraConfig
			copyNonEmptyValue(&dst.Uuid, &src.Uuid)
			copyNonEmptyValue(&dst.InstanceUuid, &src.InstanceUuid)
			copyNonEmptyValue(&dst.NumCPUs, &src.NumCPUs)
			copyNonEmptyValue(&dst.MemoryMB, &src.MemoryMB)
		}

		res := ctx.Map.Get(req.Folder).(vmFolder).CreateVMTask(ctx, &types.CreateVM_Task{
			This:   folder.Self,
			Config: config,
			Pool:   *pool,
			Host:   destHost,
		})

		ctask := ctx.Map.Get(res.(*methods.CreateVM_TaskBody).Res.Returnval).(*Task)
		ctask.Wait()
		if ctask.Info.Error != nil {
			return nil, ctask.Info.Error.Fault
		}

		ref := ctask.Info.Result.(types.ManagedObjectReference)
		clone := ctx.Map.Get(ref).(*VirtualMachine)
		clone.configureDevices(ctx, &types.VirtualMachineConfigSpec{DeviceChange: req.Spec.Location.DeviceChange})
		if req.Spec.Config != nil && req.Spec.Config.DeviceChange != nil {
			clone.configureDevices(ctx, &types.VirtualMachineConfigSpec{DeviceChange: req.Spec.Config.DeviceChange})
		}
		clone.DataSets = copyDataSetsForVmClone(vm.DataSets)

		if req.Spec.Template {
			_ = clone.MarkAsTemplate(&types.MarkAsTemplate{This: clone.Self})
		}

		ctx.postEvent(&types.VmClonedEvent{
			VmCloneEvent: types.VmCloneEvent{VmEvent: clone.event(ctx)},
			SourceVm:     *event.Vm,
		})

		return ref, nil
	})

	return &methods.CloneVM_TaskBody{
		Res: &types.CloneVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func copyNonEmptyValue[T comparable](dst, src *T) {
	if dst == nil || src == nil {
		return
	}
	var t T
	if *src == t {
		return
	}
	*dst = *src
}

func (vm *VirtualMachine) RelocateVMTask(ctx *Context, req *types.RelocateVM_Task) soap.HasFault {
	task := CreateTask(vm, "relocateVm", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		var changes []types.PropertyChange

		if ref := req.Spec.Datastore; ref != nil {
			ds := ctx.Map.Get(*ref).(*Datastore)
			ctx.Map.RemoveReference(ctx, ds, &ds.Vm, *ref)

			// TODO: migrate vm.Config.Files, vm.Summary.Config.VmPathName, vm.Layout and vm.LayoutEx

			changes = append(changes, types.PropertyChange{Name: "datastore", Val: []types.ManagedObjectReference{*ref}})
		}

		if ref := req.Spec.Pool; ref != nil {
			pool := ctx.Map.Get(*ref).(*ResourcePool)
			ctx.Map.RemoveReference(ctx, pool, &pool.Vm, *ref)

			changes = append(changes, types.PropertyChange{Name: "resourcePool", Val: ref})
		}

		if ref := req.Spec.Host; ref != nil {
			host := ctx.Map.Get(*ref).(*HostSystem)
			ctx.Map.RemoveReference(ctx, host, &host.Vm, *ref)

			changes = append(changes,
				types.PropertyChange{Name: "runtime.host", Val: ref},
				types.PropertyChange{Name: "summary.runtime.host", Val: ref},
			)
		}

		if ref := req.Spec.Folder; ref != nil {
			folder := ctx.Map.Get(*ref).(*Folder)
			ctx.WithLock(folder, func() {
				res := folder.MoveIntoFolderTask(ctx, &types.MoveIntoFolder_Task{
					List: []types.ManagedObjectReference{vm.Self},
				}).(*methods.MoveIntoFolder_TaskBody).Res
				// Wait for task to complete while we hold the Folder lock
				ctx.Map.Get(res.Returnval).(*Task).Wait()
			})
		}

		cspec := &types.VirtualMachineConfigSpec{DeviceChange: req.Spec.DeviceChange}
		if err := vm.configureDevices(ctx, cspec); err != nil {
			return nil, err
		}

		ctx.postEvent(&types.VmMigratedEvent{
			VmEvent:          vm.event(ctx),
			SourceHost:       *ctx.Map.Get(*vm.Runtime.Host).(*HostSystem).eventArgument(),
			SourceDatacenter: datacenterEventArgument(ctx, vm),
			SourceDatastore:  ctx.Map.Get(vm.Datastore[0]).(*Datastore).eventArgument(),
		})

		ctx.Update(vm, changes)

		return nil, nil
	})

	return &methods.RelocateVM_TaskBody{
		Res: &types.RelocateVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) customize(ctx *Context) {
	if vm.imc == nil {
		return
	}

	event := types.CustomizationEvent{VmEvent: vm.event(ctx)}
	ctx.postEvent(&types.CustomizationStartedEvent{CustomizationEvent: event})

	changes := []types.PropertyChange{
		{Name: "config.tools.pendingCustomization", Val: ""},
	}

	if len(vm.Guest.Net) != len(vm.imc.NicSettingMap) {
		ctx.postEvent(&types.CustomizationNetworkSetupFailed{
			CustomizationFailed: types.CustomizationFailed{
				CustomizationEvent: event,
				Reason:             "NicSettingMismatch",
			},
		})

		vm.imc = nil
		ctx.Update(vm, changes)
		return
	}

	hostname := ""
	address := ""

	switch c := vm.imc.Identity.(type) {
	case *types.CustomizationLinuxPrep:
		hostname = customizeName(vm, c.HostName)
	case *types.CustomizationSysprep:
		hostname = customizeName(vm, c.UserData.ComputerName)
	}

	cards := object.VirtualDeviceList(vm.Config.Hardware.Device).SelectByType((*types.VirtualEthernetCard)(nil))

	for i, s := range vm.imc.NicSettingMap {
		nic := &vm.Guest.Net[i]
		if s.MacAddress != "" {
			nic.MacAddress = strings.ToLower(s.MacAddress) // MacAddress in guest will always be lowercase
			card := cards[i].(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
			card.MacAddress = s.MacAddress // MacAddress in Virtual NIC can be any case
			card.AddressType = string(types.VirtualEthernetCardMacTypeManual)
		}
		if nic.DnsConfig == nil {
			nic.DnsConfig = new(types.NetDnsConfigInfo)
		}
		if s.Adapter.DnsDomain != "" {
			nic.DnsConfig.DomainName = s.Adapter.DnsDomain
		}
		if len(s.Adapter.DnsServerList) != 0 {
			nic.DnsConfig.IpAddress = s.Adapter.DnsServerList
		}
		if hostname != "" {
			nic.DnsConfig.HostName = hostname
		}
		if len(vm.imc.GlobalIPSettings.DnsSuffixList) != 0 {
			nic.DnsConfig.SearchDomain = vm.imc.GlobalIPSettings.DnsSuffixList
		}
		if nic.IpConfig == nil {
			nic.IpConfig = new(types.NetIpConfigInfo)
		}

		switch ip := s.Adapter.Ip.(type) {
		case *types.CustomizationCustomIpGenerator:
		case *types.CustomizationDhcpIpGenerator:
		case *types.CustomizationFixedIp:
			if address == "" {
				address = ip.IpAddress
			}
			nic.IpAddress = []string{ip.IpAddress}
			nic.IpConfig.IpAddress = []types.NetIpConfigInfoIpAddress{{
				IpAddress: ip.IpAddress,
			}}
		case *types.CustomizationUnknownIpGenerator:
		}
	}

	if len(vm.imc.NicSettingMap) != 0 {
		changes = append(changes, types.PropertyChange{Name: "guest.net", Val: vm.Guest.Net})
	}
	if hostname != "" {
		changes = append(changes, types.PropertyChange{Name: "guest.hostName", Val: hostname})
		changes = append(changes, types.PropertyChange{Name: "summary.guest.hostName", Val: hostname})
	}
	if address != "" {
		changes = append(changes, types.PropertyChange{Name: "guest.ipAddress", Val: address})
		changes = append(changes, types.PropertyChange{Name: "summary.guest.ipAddress", Val: address})
	}

	vm.imc = nil
	ctx.Update(vm, changes)
	ctx.postEvent(&types.CustomizationSucceeded{CustomizationEvent: event})
}

func (vm *VirtualMachine) CustomizeVMTask(ctx *Context, req *types.CustomizeVM_Task) soap.HasFault {
	task := CreateTask(vm, "customizeVm", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		if vm.hostInMM(ctx) {
			return nil, new(types.InvalidState)
		}

		if vm.Runtime.PowerState == types.VirtualMachinePowerStatePoweredOn {
			return nil, &types.InvalidPowerState{
				RequestedState: types.VirtualMachinePowerStatePoweredOff,
				ExistingState:  vm.Runtime.PowerState,
			}
		}
		if vm.Config.Tools.PendingCustomization != "" {
			return nil, new(types.CustomizationPending)
		}
		if len(vm.Guest.Net) != len(req.Spec.NicSettingMap) {
			return nil, &types.NicSettingMismatch{
				NumberOfNicsInSpec: int32(len(req.Spec.NicSettingMap)),
				NumberOfNicsInVM:   int32(len(vm.Guest.Net)),
			}
		}

		vm.imc = &req.Spec
		vm.Config.Tools.PendingCustomization = uuid.New().String()

		return nil, nil
	})

	return &methods.CustomizeVM_TaskBody{
		Res: &types.CustomizeVM_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) CreateSnapshotTask(ctx *Context, req *types.CreateSnapshot_Task) soap.HasFault {
	body := &methods.CreateSnapshot_TaskBody{}

	r := &types.CreateSnapshotEx_Task{
		Name:        req.Name,
		Description: req.Description,
		Memory:      req.Memory,
	}

	if req.Quiesce {
		r.QuiesceSpec = &types.VirtualMachineGuestQuiesceSpec{}
	}

	res := vm.CreateSnapshotExTask(ctx, r)

	if res.Fault() != nil {
		body.Fault_ = res.Fault()
	} else {
		body.Res = &types.CreateSnapshot_TaskResponse{
			Returnval: res.(*methods.CreateSnapshotEx_TaskBody).Res.Returnval,
		}
	}

	return body
}

func (vm *VirtualMachine) CreateSnapshotExTask(ctx *Context, req *types.CreateSnapshotEx_Task) soap.HasFault {
	task := CreateTask(vm, "createSnapshotEx", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		var changes []types.PropertyChange

		if vm.Snapshot == nil {
			vm.Snapshot = &types.VirtualMachineSnapshotInfo{}
		}

		snapshot := &VirtualMachineSnapshot{}
		snapshot.Vm = vm.Reference()
		snapshot.Config = copyConfigFromVmConfig(vm.Config)
		snapshot.DataSets = copyDataSetsForVmClone(vm.DataSets)

		ctx.Map.Put(snapshot)

		quiesced := false
		if req.QuiesceSpec != nil {
			quiesced = true
		}

		snapPowerState := vm.Runtime.PowerState
		if !req.Memory {
			snapPowerState = types.VirtualMachinePowerStatePoweredOff
		}

		treeItem := types.VirtualMachineSnapshotTree{
			Snapshot:        snapshot.Self,
			Vm:              snapshot.Vm,
			Name:            req.Name,
			Description:     req.Description,
			Id:              atomic.AddInt32(&vm.sid, 1),
			CreateTime:      time.Now(),
			State:           snapPowerState,
			Quiesced:        quiesced,
			BackupManifest:  "",
			ReplaySupported: types.NewBool(false),
		}

		cur := vm.Snapshot.CurrentSnapshot
		if cur != nil {
			parent := ctx.Map.Get(*cur).(*VirtualMachineSnapshot)
			parent.ChildSnapshot = append(parent.ChildSnapshot, snapshot.Self)

			ss := findSnapshotInTree(vm.Snapshot.RootSnapshotList, *cur)
			ss.ChildSnapshotList = append(ss.ChildSnapshotList, treeItem)
		} else {
			changes = append(changes, types.PropertyChange{
				Name: "snapshot.rootSnapshotList",
				Val:  append(vm.Snapshot.RootSnapshotList, treeItem),
			})
			changes = append(changes, types.PropertyChange{
				Name: "rootSnapshot",
				Val:  append(vm.RootSnapshot, treeItem.Snapshot),
			})
		}

		snapshot.createSnapshotFiles(ctx)

		changes = append(changes, types.PropertyChange{Name: "snapshot.currentSnapshot", Val: snapshot.Self})
		ctx.Update(vm, changes)

		return snapshot.Self, nil
	})

	return &methods.CreateSnapshotEx_TaskBody{
		Res: &types.CreateSnapshotEx_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) RevertToCurrentSnapshotTask(ctx *Context, req *types.RevertToCurrentSnapshot_Task) soap.HasFault {
	body := &methods.RevertToCurrentSnapshot_TaskBody{}

	if vm.Snapshot == nil || vm.Snapshot.CurrentSnapshot == nil {
		body.Fault_ = Fault("snapshot not found", &types.NotFound{})

		return body
	}
	snapshot := ctx.Map.Get(*vm.Snapshot.CurrentSnapshot).(*VirtualMachineSnapshot)

	task := CreateTask(vm, "revertSnapshot", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		vm.DataSets = copyDataSetsForVmClone(snapshot.DataSets)
		return nil, nil
	})

	body.Res = &types.RevertToCurrentSnapshot_TaskResponse{
		Returnval: task.Run(ctx),
	}

	return body
}

func (vm *VirtualMachine) RemoveAllSnapshotsTask(ctx *Context, req *types.RemoveAllSnapshots_Task) soap.HasFault {
	task := CreateTask(vm, "RemoveAllSnapshots", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		if vm.Snapshot == nil {
			return nil, nil
		}

		refs := allSnapshotsInTree(vm.Snapshot.RootSnapshotList)

		ctx.Update(vm, []types.PropertyChange{
			{Name: "snapshot", Val: nil},
			{Name: "rootSnapshot", Val: nil},
		})

		for _, ref := range refs {
			ctx.Map.Get(ref).(*VirtualMachineSnapshot).removeSnapshotFiles(ctx)
			ctx.Map.Remove(ctx, ref)
		}

		return nil, nil
	})

	return &methods.RemoveAllSnapshots_TaskBody{
		Res: &types.RemoveAllSnapshots_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) fcd(ctx *Context, ds types.ManagedObjectReference, id types.ID) *VStorageObject {
	m := ctx.Map.VStorageObjectManager()
	if ds.Value != "" {
		return m.objects[ds][id]
	}
	for _, set := range m.objects {
		for key, val := range set {
			if key == id {
				return val
			}
		}
	}
	return nil
}

func (vm *VirtualMachine) AttachDiskTask(ctx *Context, req *types.AttachDisk_Task) soap.HasFault {
	task := CreateTask(vm, "attachDisk", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		fcd := vm.fcd(ctx, req.Datastore, req.DiskId)
		if fcd == nil {
			return nil, new(types.InvalidArgument)
		}

		fcd.Config.ConsumerId = []types.ID{{Id: vm.Config.Uuid}}

		// TODO: add device

		return nil, nil
	})

	return &methods.AttachDisk_TaskBody{
		Res: &types.AttachDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) DetachDiskTask(ctx *Context, req *types.DetachDisk_Task) soap.HasFault {
	task := CreateTask(vm, "detachDisk", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		fcd := vm.fcd(ctx, types.ManagedObjectReference{}, req.DiskId)
		if fcd == nil {
			return nil, new(types.InvalidArgument)
		}

		fcd.Config.ConsumerId = nil

		// TODO: remove device

		return nil, nil
	})

	return &methods.DetachDisk_TaskBody{
		Res: &types.DetachDisk_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) PromoteDisksTask(ctx *Context, req *types.PromoteDisks_Task) soap.HasFault {
	task := CreateTask(vm, "promoteDisks", func(t *Task) (types.AnyType, types.BaseMethodFault) {
		devices := object.VirtualDeviceList(vm.Config.Hardware.Device)
		devices = devices.SelectByType((*types.VirtualDisk)(nil))
		var cap int64

		for i := range req.Disks {
			d := devices.FindByKey(req.Disks[i].Key)
			if d == nil {
				return nil, &types.InvalidArgument{InvalidProperty: "disks"}
			}

			disk := d.(*types.VirtualDisk)

			switch backing := disk.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				if backing.Parent != nil {
					cap += disk.CapacityInBytes
					if req.Unlink {
						backing.Parent = nil
					}
				}
			case *types.VirtualDiskSeSparseBackingInfo:
				if backing.Parent != nil {
					cap += disk.CapacityInBytes
					if req.Unlink {
						backing.Parent = nil
					}
				}
			case *types.VirtualDiskSparseVer2BackingInfo:
				if backing.Parent != nil {
					cap += disk.CapacityInBytes
					if req.Unlink {
						backing.Parent = nil
					}
				}
			}
		}

		// Built-in default delay. `simulator.TaskDelay` can be used to add additional time
		// Translates to roughly 1s per 1GB
		sleep := time.Duration(cap/units.MB) * time.Millisecond
		if sleep > 0 {
			log.Printf("%s: sleep %s for %s", t.Info.DescriptionId, sleep, units.ByteSize(cap))
			time.Sleep(sleep)
		}

		return nil, nil
	})

	return &methods.PromoteDisks_TaskBody{
		Res: &types.PromoteDisks_TaskResponse{
			Returnval: task.Run(ctx),
		},
	}
}

func (vm *VirtualMachine) ShutdownGuest(ctx *Context, c *types.ShutdownGuest) soap.HasFault {
	r := &methods.ShutdownGuestBody{}

	if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
		r.Fault_ = Fault("", &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOn,
			ExistingState:  vm.Runtime.PowerState,
		})

		return r
	}

	event := vm.event(ctx)
	ctx.postEvent(&types.VmGuestShutdownEvent{VmEvent: event})

	_ = CreateTask(vm, "shutdownGuest", func(*Task) (types.AnyType, types.BaseMethodFault) {
		vm.svm.stop(ctx)

		ctx.Update(vm, []types.PropertyChange{
			{Name: "runtime.powerState", Val: types.VirtualMachinePowerStatePoweredOff},
			{Name: "summary.runtime.powerState", Val: types.VirtualMachinePowerStatePoweredOff},
		})

		ctx.postEvent(&types.VmPoweredOffEvent{VmEvent: event})

		return nil, nil
	}).Run(ctx)

	r.Res = new(types.ShutdownGuestResponse)

	return r
}

func (vm *VirtualMachine) StandbyGuest(ctx *Context, c *types.StandbyGuest) soap.HasFault {
	r := &methods.StandbyGuestBody{}

	if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
		r.Fault_ = Fault("", &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOn,
			ExistingState:  vm.Runtime.PowerState,
		})

		return r
	}

	event := vm.event(ctx)
	ctx.postEvent(&types.VmGuestStandbyEvent{VmEvent: event})

	_ = CreateTask(vm, "standbyGuest", func(*Task) (types.AnyType, types.BaseMethodFault) {
		vm.svm.pause(ctx)

		ctx.Update(vm, []types.PropertyChange{
			{Name: "runtime.powerState", Val: types.VirtualMachinePowerStateSuspended},
			{Name: "summary.runtime.powerState", Val: types.VirtualMachinePowerStateSuspended},
		})

		ctx.postEvent(&types.VmSuspendedEvent{VmEvent: event})

		return nil, nil
	}).Run(ctx)

	r.Res = new(types.StandbyGuestResponse)

	return r
}

func (vm *VirtualMachine) MarkAsTemplate(req *types.MarkAsTemplate) soap.HasFault {
	r := &methods.MarkAsTemplateBody{}

	if vm.Config.Template {
		r.Fault_ = Fault("", new(types.NotSupported))
		return r
	}

	if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOff {
		r.Fault_ = Fault("", &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOff,
			ExistingState:  vm.Runtime.PowerState,
		})
		return r
	}

	vm.Config.Template = true
	vm.Summary.Config.Template = true
	vm.ResourcePool = nil

	r.Res = new(types.MarkAsTemplateResponse)

	return r
}

func (vm *VirtualMachine) MarkAsVirtualMachine(req *types.MarkAsVirtualMachine) soap.HasFault {
	r := &methods.MarkAsVirtualMachineBody{}

	if !vm.Config.Template {
		r.Fault_ = Fault("", new(types.NotSupported))
		return r
	}

	if vm.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOff {
		r.Fault_ = Fault("", &types.InvalidPowerState{
			RequestedState: types.VirtualMachinePowerStatePoweredOff,
			ExistingState:  vm.Runtime.PowerState,
		})
		return r
	}

	vm.Config.Template = false
	vm.Summary.Config.Template = false
	vm.ResourcePool = &req.Pool
	if req.Host != nil {
		vm.Runtime.Host = req.Host
	}

	r.Res = new(types.MarkAsVirtualMachineResponse)

	return r
}

func findSnapshotInTree(tree []types.VirtualMachineSnapshotTree, ref types.ManagedObjectReference) *types.VirtualMachineSnapshotTree {
	if tree == nil {
		return nil
	}

	for i, ss := range tree {
		if ss.Snapshot == ref {
			return &tree[i]
		}

		target := findSnapshotInTree(ss.ChildSnapshotList, ref)
		if target != nil {
			return target
		}
	}

	return nil
}

func findParentSnapshot(tree types.VirtualMachineSnapshotTree, ref types.ManagedObjectReference) *types.ManagedObjectReference {
	for _, ss := range tree.ChildSnapshotList {
		if ss.Snapshot == ref {
			return &tree.Snapshot
		}

		res := findParentSnapshot(ss, ref)
		if res != nil {
			return res
		}
	}

	return nil
}

func findParentSnapshotInTree(tree []types.VirtualMachineSnapshotTree, ref types.ManagedObjectReference) *types.ManagedObjectReference {
	if tree == nil {
		return nil
	}

	for _, ss := range tree {
		res := findParentSnapshot(ss, ref)
		if res != nil {
			return res
		}
	}

	return nil
}

func removeSnapshotInTree(tree []types.VirtualMachineSnapshotTree, ref types.ManagedObjectReference, removeChildren bool) []types.VirtualMachineSnapshotTree {
	if tree == nil {
		return tree
	}

	var result []types.VirtualMachineSnapshotTree

	for _, ss := range tree {
		if ss.Snapshot == ref {
			if !removeChildren {
				result = append(result, ss.ChildSnapshotList...)
			}
		} else {
			ss.ChildSnapshotList = removeSnapshotInTree(ss.ChildSnapshotList, ref, removeChildren)
			result = append(result, ss)
		}
	}

	return result
}

func allSnapshotsInTree(tree []types.VirtualMachineSnapshotTree) []types.ManagedObjectReference {
	var result []types.ManagedObjectReference

	if tree == nil {
		return result
	}

	for _, ss := range tree {
		result = append(result, ss.Snapshot)
		result = append(result, allSnapshotsInTree(ss.ChildSnapshotList)...)
	}

	return result
}

func changeTrackingSupported(spec *types.VirtualMachineConfigSpec) bool {
	for _, device := range spec.DeviceChange {
		if dev, ok := device.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
			switch dev.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				return true
			case *types.VirtualDiskSparseVer2BackingInfo:
				return true
			case *types.VirtualDiskRawDiskMappingVer1BackingInfo:
				return true
			case *types.VirtualDiskRawDiskVer2BackingInfo:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func (vm *VirtualMachine) updateLastModifiedAndChangeVersion(ctx *Context) {
	modified := time.Now()
	ctx.Update(vm, []types.PropertyChange{
		{
			Name: "config.changeVersion",
			Val:  fmt.Sprintf("%d", modified.UnixNano()),
			Op:   types.PropertyChangeOpAssign,
		},
		{
			Name: "config.modified",
			Val:  modified,
			Op:   types.PropertyChangeOpAssign,
		},
	})
}

func (vm *VirtualMachine) updateTagSpec(
	ctx *Context,
	specs []types.TagSpec) types.BaseMethodFault {

	if len(specs) == 0 {
		return nil
	}

	// If the VAPI simulator is not loaded, the tagManager will be nil.
	if ctx.Map.tagManager == nil {
		return nil
	}

	vmRef := vm.Reference()

	for _, spec := range specs {
		tagID := spec.Id.Uuid
		if spec.Id.NameId != nil {
			t, err := ctx.Map.tagManager.GetTagByCategoryAndName(
				spec.Id.NameId.Category,
				spec.Id.NameId.Tag)
			if err != nil {
				return err
			}
			tagID = t
		}

		switch spec.Operation {
		case types.ArrayUpdateOperationAdd:
			if err := ctx.Map.tagManager.AttachTag(vmRef, tagID); err != nil {
				return err
			}
		case types.ArrayUpdateOperationRemove:
			if err := ctx.Map.tagManager.DetachTag(vmRef, tagID); err != nil {
				return err
			}
		default:
			return &types.InvalidArgument{
				InvalidProperty: "tagSpecs.operation",
			}
		}
	}

	return nil
}
