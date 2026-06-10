package convert

import (
	"fmt"

	mo "github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"k8s.io/apimachinery/pkg/api/resource"

	vimv1 "github.com/vmware/govmomi/crd/pkg/vim/api/v1alpha1"
)

type ConvertFunc func(dstObj, srcObj any) error

var (
	vim2CRDFuncs = map[string]ConvertFunc{}
	crd2VIMFuncs = map[string]ConvertFunc{}
)

func ConvertVIMToCRD(dst, src any) error {
	switch tDst := dst.(type) {

	case *vimv1.HostCPUIDInfo:
		var tSrc types.HostCpuIdInfo
		switch s := src.(type) {
		case *types.HostCpuIdInfo:
			tSrc = *s
		case types.HostCpuIdInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.HostCpuIdInfo")
		}
		if err := convertVIMToCRD_HostCPUIDInfo(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.HostCpuIdInfo: %w", err)
		}

	case *vimv1.LatencySensitivity:
		var tSrc types.LatencySensitivity
		switch s := src.(type) {
		case *types.LatencySensitivity:
			tSrc = *s
		case types.LatencySensitivity:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.LatencySensitivity")
		}
		if err := convertVIMToCRD_LatencySensitivity(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.LatencySensitivity: %w", err)
		}

	case *vimv1.ManagedByInfo:
		var tSrc types.ManagedByInfo
		switch s := src.(type) {
		case *types.ManagedByInfo:
			tSrc = *s
		case types.ManagedByInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.ManagedByInfo")
		}
		if err := convertVIMToCRD_ManagedByInfo(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.ManagedByInfo: %w", err)
		}

	case *vimv1.ReplicationConfigSpec:
		var tSrc types.ReplicationConfigSpec
		switch s := src.(type) {
		case *types.ReplicationConfigSpec:
			tSrc = *s
		case types.ReplicationConfigSpec:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.ReplicationConfigSpec")
		}
		if err := convertVIMToCRD_ReplicationConfigSpec(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.ReplicationConfigSpec: %w", err)
		}

	case *vimv1.ResourceAllocationInfo:
		var tSrc types.ResourceAllocationInfo
		switch s := src.(type) {
		case *types.ResourceAllocationInfo:
			tSrc = *s
		case types.ResourceAllocationInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.ResourceAllocationInfo")
		}
		if err := convertVIMToCRD_ResourceAllocationInfo(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.ResourceAllocationInfo: %w", err)
		}

	case *vimv1.ScheduledHardwareUpgradeInfo:
		var tSrc types.ScheduledHardwareUpgradeInfo
		switch s := src.(type) {
		case *types.ScheduledHardwareUpgradeInfo:
			tSrc = *s
		case types.ScheduledHardwareUpgradeInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.ScheduledHardwareUpgradeInfo")
		}
		if err := convertVIMToCRD_ScheduledHardwareUpgradeInfo(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.ScheduledHardwareUpgradeInfo: %w",
				err)
		}

	case *vimv1.ToolsConfigInfo:
		var tSrc types.ToolsConfigInfo
		switch s := src.(type) {
		case *types.ToolsConfigInfo:
			tSrc = *s
		case types.ToolsConfigInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.ToolsConfigInfo")
		}
		if err := convertVIMToCRD_ToolsConfigInfo(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.ToolsConfigInfo: %w", err)
		}

	case *vimv1.VirtualHardwareSpec:
		var tSrc types.VirtualHardware
		switch s := src.(type) {
		case *types.VirtualHardware:
			tSrc = *s
		case types.VirtualHardware:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualHardware")
		}
		if err := convertVIMToCRD_VirtualHardwareSpec(tDst, tSrc); err != nil {
			return fmt.Errorf("failed to convert types.VirtualHardware: %w", err)
		}

	case *vimv1.VirtualMachine:
		var tSrc mo.VirtualMachine
		switch s := src.(type) {
		case *mo.VirtualMachine:
			tSrc = *s
		case mo.VirtualMachine:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s", "mo.VirtualMachine")
		}
		if tSrc.Config != nil {
			tDst.Spec.Config = &vimv1.VirtualMachineConfigInfoSpec{}
			if err := ConvertVIMToCRD(tDst.Spec.Config, tSrc.Config); err != nil {
				return fmt.Errorf("failed to convert moVM.Config: %w", err)
			}
		}

	case *vimv1.VirtualMachineConfigInfoSpec:
		var tSrc types.VirtualMachineConfigInfo
		switch s := src.(type) {
		case *types.VirtualMachineConfigInfo:
			tSrc = *s
		case types.VirtualMachineConfigInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineConfigInfo")
		}
		if err := convertVIMToCRD_ConfigInfoSpec(tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineConfigInfo: %w", err)
		}

	case *vimv1.VirtualMachineDefaultPowerOpInfo:
		var tSrc types.VirtualMachineDefaultPowerOpInfo
		switch s := src.(type) {
		case *types.VirtualMachineDefaultPowerOpInfo:
			tSrc = *s
		case types.VirtualMachineDefaultPowerOpInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineDefaultPowerOpInfo")
		}
		if err := convertVIMToCRD_VirtualMachineDefaultPowerOpInfo(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineDefaultPowerOpInfo: %w", err)
		}

	case *vimv1.VirtualMachineFileInfo:
		var tSrc types.VirtualMachineFileInfo
		switch s := src.(type) {
		case *types.VirtualMachineFileInfo:
			tSrc = *s
		case types.VirtualMachineFileInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineFileInfo")
		}
		if err := convertVIMToCRD_VirtualMachineFileInfo(tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineFileInfo: %w", err)
		}

	case *vimv1.VirtualMachineSgxInfo:
		var tSrc types.VirtualMachineSgxInfo
		switch s := src.(type) {
		case *types.VirtualMachineSgxInfo:
			tSrc = *s
		case types.VirtualMachineSgxInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineSgxInfo")
		}
		if err := convertVIMToCRD_VirtualMachineSgxInfo(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineSgxInfo: %w", err)
		}

	case *vimv1.VirtualMachineVcpuConfig:
		var tSrc types.VirtualMachineVcpuConfig
		switch s := src.(type) {
		case *types.VirtualMachineVcpuConfig:
			tSrc = *s
		case types.VirtualMachineVcpuConfig:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineVcpuConfig")
		}
		if err := convertVIMToCRD_VirtualMachineVcpuConfig(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineVcpuConfig: %w", err)
		}

	case *vimv1.VirtualMachineVirtualDeviceGroups:
		var tSrc types.VirtualMachineVirtualDeviceGroups
		switch s := src.(type) {
		case *types.VirtualMachineVirtualDeviceGroups:
			tSrc = *s
		case types.VirtualMachineVirtualDeviceGroups:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineVirtualDeviceGroups")
		}
		if err := convertVIMToCRD_VirtualMachineVirtualDeviceGroups(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineVirtualDeviceGroups: %w", err)
		}

	case *vimv1.VirtualMachineVirtualDeviceSwap:
		var tSrc types.VirtualMachineVirtualDeviceSwap
		switch s := src.(type) {
		case *types.VirtualMachineVirtualDeviceSwap:
			tSrc = *s
		case types.VirtualMachineVirtualDeviceSwap:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineVirtualDeviceSwap")
		}
		if err := convertVIMToCRD_VirtualMachineVirtualDeviceSwap(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineVirtualDeviceSwap: %w", err)
		}

	case *vimv1.VirtualMachineVirtualNumaInfo:
		var tSrc types.VirtualMachineVirtualNumaInfo
		switch s := src.(type) {
		case *types.VirtualMachineVirtualNumaInfo:
			tSrc = *s
		case types.VirtualMachineVirtualNumaInfo:
			tSrc = s
		default:
			return fmt.Errorf(
				"src is not a %[1]s or *%[1]s",
				"types.VirtualMachineVirtualNumaInfo")
		}
		if err := convertVIMToCRD_VirtualMachineVirtualNumaInfo(
			tDst, tSrc); err != nil {
			return fmt.Errorf(
				"failed to convert types.VirtualMachineVirtualNumaInfo: %w", err)
		}

	}
	return nil
}

// convertCryptoKeyId is a non-fallible helper since CryptoKeyId has no
// conversion logic that can fail.
func convertCryptoKeyId(src *types.CryptoKeyId) *vimv1.CryptoKeyId {
	if src == nil {
		return nil
	}
	dst := &vimv1.CryptoKeyId{KeyId: src.KeyId}
	if src.ProviderId != nil {
		dst.ProviderId = &vimv1.KeyProviderId{Id: src.ProviderId.Id}
	}
	return dst
}

func init() {
	vim2CRDFuncs["VirtualMachineConfigInfoSpec"] = convertVIMToCRD_ConfigInfoSpec
}

func convertVIMToCRD_ConfigInfoSpec(dstObj, srcObj any) error {

	dst, ok := dstObj.(*vimv1.VirtualMachineConfigInfoSpec)
	if !ok {
		return fmt.Errorf(
			"dst is not a *vimv1.VirtualMachineConfigInfoSpec")
	}
	src, ok := srcObj.(types.VirtualMachineConfigInfo)
	if !ok {
		return fmt.Errorf(
			"src is not a types.VirtualMachineConfigInfo")
	}

	// Identity
	dst.Name = src.Name
	dst.GuestFullName = src.GuestFullName
	dst.HardwareVersion = src.Version
	dst.UUID = src.Uuid
	dst.InstanceUUID = src.InstanceUuid
	dst.LocationID = src.LocationId
	dst.Template = src.Template
	dst.GuestID = src.GuestId
	dst.AlternateGuestName = src.AlternateGuestName
	dst.Annotation = src.Annotation

	// NPIV
	dst.NPIVNodeWorldWideName = src.NpivNodeWorldWideName
	dst.NPIVPortWorldWideName = src.NpivPortWorldWideName
	dst.NPIVWorldWideNameType = src.NpivWorldWideNameType
	dst.NPIVTemporaryDisabled = src.NpivTemporaryDisabled
	dst.NPIVOnNonRdmDisks = src.NpivOnNonRdmDisks
	if v := int32(src.NpivDesiredNodeWwns); v != 0 {
		dst.NPIVDesiredNodeWwns = &v
	}
	if v := int32(src.NpivDesiredPortWwns); v != 0 {
		dst.NPIVDesiredPortWwns = &v
	}

	if err := ConvertVIMToCRD(&dst.Files, src.Files); err != nil {
		return err
	}

	if src.Tools != nil {
		dst.Tools = &vimv1.ToolsConfigInfo{}
		if err := ConvertVIMToCRD(dst.Tools, src.Tools); err != nil {
			return err
		}
	}

	if err := ConvertVIMToCRD(
		&dst.DefaultPowerOps, src.DefaultPowerOps); err != nil {
		return err
	}

	dst.RebootPowerOff = src.RebootPowerOff

	if err := ConvertVIMToCRD(&dst.Hardware, src.Hardware); err != nil {
		return err
	}

	if l := len(src.VcpuConfig); l > 0 {
		dst.VCPUConfig = make([]vimv1.VirtualMachineVcpuConfig, l)
		for i := range src.VcpuConfig {
			if err := ConvertVIMToCRD(
				&dst.VCPUConfig[i], src.VcpuConfig[i]); err != nil {
				return err
			}
		}
	}

	if src.CpuAllocation != nil {
		dst.CPUAllocation = &vimv1.ResourceAllocationInfo{}
		if err := ConvertVIMToCRD(dst.CPUAllocation, src.CpuAllocation); err != nil {
			return err
		}
	}

	if src.MemoryAllocation != nil {
		dst.MemoryAllocation = &vimv1.ResourceAllocationInfo{}
		if err := ConvertVIMToCRD(
			dst.MemoryAllocation, src.MemoryAllocation); err != nil {
			return err
		}
	}

	if src.LatencySensitivity != nil {
		dst.LatencySensitivity = &vimv1.LatencySensitivity{}
		if err := ConvertVIMToCRD(
			dst.LatencySensitivity, src.LatencySensitivity); err != nil {
			return err
		}
	}

	dst.MemoryHotAddEnabled = src.MemoryHotAddEnabled
	dst.CPUHotAddEnabled = src.CpuHotAddEnabled
	dst.CPUHotRemoveEnabled = src.CpuHotRemoveEnabled

	// HotPlugMemoryLimit / HotPlugMemoryIncrementSize (MB → resource.Quantity)
	if v := src.HotPlugMemoryLimit; v != 0 {
		dst.HotPlugMemoryLimit = resource.NewQuantity(v*1024*1024, resource.BinarySI)
	}
	if v := src.HotPlugMemoryIncrementSize; v != 0 {
		dst.HotPlugMemoryIncrementSize = resource.NewQuantity(
			v*1024*1024, resource.BinarySI)
	}

	if a := src.CpuAffinity; a != nil {
		dst.CPUAffinity = &vimv1.VirtualMachineAffinityInfo{
			AffinitySet: a.AffinitySet,
		}
	}
	if a := src.MemoryAffinity; a != nil {
		dst.MemoryAffinity = &vimv1.VirtualMachineAffinityInfo{
			AffinitySet: a.AffinitySet,
		}
	}

	if l := len(src.CpuFeatureMask); l > 0 {
		dst.CPUFeatureMask = make([]vimv1.HostCPUIDInfo, l)
		for i := range src.CpuFeatureMask {
			if err := ConvertVIMToCRD(
				&dst.CPUFeatureMask[i], src.CpuFeatureMask[i]); err != nil {
				return err
			}
		}
	}

	dst.SwapPlacement = src.SwapPlacement

	// TODO(akutz): dst.BootOptions = ...

	if src.RepConfig != nil {
		dst.RepConfig = &vimv1.ReplicationConfigSpec{}
		if err := ConvertVIMToCRD(dst.RepConfig, src.RepConfig); err != nil {
			return err
		}
	}

	dst.VAssertsEnabled = src.VAssertsEnabled
	dst.ChangeTrackingEnabled = src.ChangeTrackingEnabled

	if src.Firmware != "" {
		switch types.GuestOsDescriptorFirmwareType(src.Firmware) {
		case types.GuestOsDescriptorFirmwareTypeBios:
			v := vimv1.FirmwareBIOS
			dst.Firmware = &v
		case types.GuestOsDescriptorFirmwareTypeEfi:
			v := vimv1.FirmwareEFI
			dst.Firmware = &v
		default:
			return fmt.Errorf("unknown firmware type: %s", src.Firmware)
		}
	}

	if v := src.MaxMksConnections; v != 0 {
		dst.MaxMKSConnections = &v
	}

	dst.GuestAutoLockEnabled = src.GuestAutoLockEnabled

	if src.ManagedBy != nil {
		dst.ManagedBy = &vimv1.ManagedByInfo{}
		if err := ConvertVIMToCRD(dst.ManagedBy, src.ManagedBy); err != nil {
			return err
		}
	}

	dst.MemoryReservationLockedToMax = src.MemoryReservationLockedToMax

	if o := src.InitialOverhead; o != nil {
		dst.InitialOverhead = &vimv1.VirtualMachineConfigInfoOverheadInfo{}
		if v := o.InitialMemoryReservation; v != 0 {
			dst.InitialOverhead.InitialMemoryReservation = &v
		}
		if v := o.InitialSwapReservation; v != 0 {
			dst.InitialOverhead.InitialSwapReservation = &v
		}
	}

	dst.NestedHVEnabled = src.NestedHVEnabled
	dst.VPMCEnabled = src.VPMCEnabled

	if src.ScheduledHardwareUpgradeInfo != nil {
		dst.ScheduledHardwareUpgradeInfo = &vimv1.ScheduledHardwareUpgradeInfo{}
		if err := ConvertVIMToCRD(
			dst.ScheduledHardwareUpgradeInfo,
			src.ScheduledHardwareUpgradeInfo); err != nil {
			return err
		}
	}

	if v := src.VFlashCacheReservation; v != 0 {
		dst.VFlashCacheReservation = &v
	}

	dst.VMXConfigChecksum = src.VmxConfigChecksum
	dst.MessageBusTunnelEnabled = src.MessageBusTunnelEnabled
	dst.VMStorageObjectID = src.VmStorageObjectId
	dst.SwapStorageObjectID = src.SwapStorageObjectId

	if src.KeyId != nil {
		dst.KeyId = convertCryptoKeyId(src.KeyId)
	}

	dst.MigrateEncryption = src.MigrateEncryption

	if src.SgxInfo != nil {
		dst.SgxInfo = &vimv1.VirtualMachineSgxInfo{}
		if err := ConvertVIMToCRD(dst.SgxInfo, src.SgxInfo); err != nil {
			return err
		}
	}

	dst.FTEncryptionMode = src.FtEncryptionMode
	dst.SEVEnabled = src.SevEnabled

	if src.NumaInfo != nil {
		dst.NumaInfo = &vimv1.VirtualMachineVirtualNumaInfo{}
		if err := ConvertVIMToCRD(dst.NumaInfo, src.NumaInfo); err != nil {
			return err
		}
	}

	dst.PMEMFailoverEnabled = src.PmemFailoverEnabled
	dst.VMXStatsCollectionEnabled = src.VmxStatsCollectionEnabled
	dst.VMOpNotificationToAppEnabled = src.VmOpNotificationToAppEnabled

	if v := src.VmOpNotificationTimeout; v != 0 {
		dst.VMOpNotificationTimeout = &v
	}

	if src.DeviceSwap != nil {
		dst.DeviceSwap = &vimv1.VirtualMachineVirtualDeviceSwap{}
		if err := ConvertVIMToCRD(dst.DeviceSwap, src.DeviceSwap); err != nil {
			return err
		}
	}

	if src.DeviceGroups != nil {
		dst.DeviceGroups = &vimv1.VirtualMachineVirtualDeviceGroups{}
		if err := ConvertVIMToCRD(dst.DeviceGroups, src.DeviceGroups); err != nil {
			return err
		}
	}

	// ExtraConfig: BaseOptionValue → OptionValue (Value encoded as string)
	if len(src.ExtraConfig) > 0 {
		dst.ExtraConfig = make([]vimv1.OptionValue, 0, len(src.ExtraConfig))
		for _, bov := range src.ExtraConfig {
			if ov := bov.GetOptionValue(); ov != nil {
				dst.ExtraConfig = append(dst.ExtraConfig, vimv1.OptionValue{
					Key:   ov.Key,
					Value: fmt.Sprintf("%v", ov.Value),
				})
			}
		}
	}

	dst.FixedPassthruHotPlugEnabled = src.FixedPassthruHotPlugEnabled
	dst.MetroFTEnabled = src.MetroFtEnabled
	dst.MetroFTHostGroup = src.MetroFtHostGroup
	dst.TDXEnabled = src.TdxEnabled
	dst.SEVSNPEnabled = src.SevSnpEnabled

	return nil
}

func init() {
	vim2CRDFuncs["HostCPUIDInfo"] = convertVIMToCRD_HostCPUIDInfo
}

func convertVIMToCRD_HostCPUIDInfo(dstObj, srcObj any) error {

	dst, ok := dstObj.(*vimv1.HostCPUIDInfo)
	if !ok {
		return fmt.Errorf("dst is not a *vimv1.HostCPUIDInfo")
	}
	src, ok := srcObj.(types.HostCpuIdInfo)
	if !ok {
		return fmt.Errorf("src is not a types.HostCpuIdInfo")
	}

	dst.Level = src.Level
	dst.Vendor = src.Vendor
	dst.EAX = src.Eax
	dst.EBX = src.Ebx
	dst.ECX = src.Ecx
	dst.EDX = src.Edx

	return nil
}

func convertVIMToCRD_LatencySensitivity(
	dst *vimv1.LatencySensitivity,
	src types.LatencySensitivity) error {

	dst.Level = string(src.Level)
	if src.Sensitivity != 0 {
		v := src.Sensitivity
		dst.Sensitivity = &v
	}

	return nil
}

func convertVIMToCRD_ManagedByInfo(
	dst *vimv1.ManagedByInfo,
	src types.ManagedByInfo) error {

	dst.ExtensionKey = src.ExtensionKey
	dst.Type = src.Type

	return nil
}

func convertVIMToCRD_ReplicationConfigSpec(
	dst *vimv1.ReplicationConfigSpec,
	src types.ReplicationConfigSpec) error {

	dst.Generation = src.Generation
	dst.VMReplicationId = src.VmReplicationId
	dst.Destination = src.Destination
	dst.Port = src.Port
	dst.RPO = src.Rpo
	dst.QuiesceGuestEnabled = src.QuiesceGuestEnabled
	dst.Paused = src.Paused
	dst.OppUpdatesEnabled = src.OppUpdatesEnabled
	dst.NetCompressionEnabled = src.NetCompressionEnabled
	dst.NetEncryptionEnabled = src.NetEncryptionEnabled
	dst.EncryptionDestination = src.EncryptionDestination
	dst.RemoteCertificateThumbprint = src.RemoteCertificateThumbprint
	dst.DataSetsReplicationEnabled = src.DataSetsReplicationEnabled
	if v := src.EncryptionPort; v != 0 {
		dst.EncryptionPort = &v
	}
	if l := len(src.Disk); l > 0 {
		dst.Disk = make([]vimv1.ReplicationInfoDiskSettings, l)
		for i, d := range src.Disk {
			dst.Disk[i] = vimv1.ReplicationInfoDiskSettings{
				Key:               d.Key,
				DiskReplicationId: d.DiskReplicationId,
			}
		}
	}

	return nil
}

func convertVIMToCRD_ResourceAllocationInfo(
	dst *vimv1.ResourceAllocationInfo,
	src types.ResourceAllocationInfo) error {

	dst.Reservation = src.Reservation
	dst.ExpandableReservation = src.ExpandableReservation
	dst.Limit = src.Limit
	dst.OverheadLimit = src.OverheadLimit
	if src.Shares != nil {
		dst.Shares = &vimv1.SharesInfo{
			Level:  string(src.Shares.Level),
			Shares: src.Shares.Shares,
		}
	}

	return nil
}

func convertVIMToCRD_ScheduledHardwareUpgradeInfo(
	dst *vimv1.ScheduledHardwareUpgradeInfo,
	src types.ScheduledHardwareUpgradeInfo) error {

	dst.UpgradePolicy = src.UpgradePolicy
	dst.VersionKey = src.VersionKey
	dst.ScheduledHardwareUpgradeStatus = src.ScheduledHardwareUpgradeStatus
	if src.Fault != nil {
		dst.Fault = &vimv1.LocalizedMethodFault{
			LocalizedMessage: src.Fault.LocalizedMessage,
		}
	}

	return nil
}

func convertVIMToCRD_ToolsConfigInfo(
	dst *vimv1.ToolsConfigInfo,
	src types.ToolsConfigInfo) error {

	if src.ToolsVersion != 0 {
		v := src.ToolsVersion
		dst.ToolsVersion = &v
	}
	dst.AfterPowerOn = src.AfterPowerOn
	dst.AfterResume = src.AfterResume
	dst.BeforeGuestReboot = src.BeforeGuestReboot
	dst.BeforeGuestShutdown = src.BeforeGuestShutdown
	dst.BeforeGuestStandby = src.BeforeGuestStandby
	dst.ToolsUpgradePolicy = src.ToolsUpgradePolicy
	dst.PendingCustomization = src.PendingCustomization
	dst.SyncTimeWithHost = src.SyncTimeWithHost
	dst.SyncTimeWithHostAllowed = src.SyncTimeWithHostAllowed
	dst.ToolsInstallType = src.ToolsInstallType
	dst.CustomizationKeyId = convertCryptoKeyId(src.CustomizationKeyId)
	if li := src.LastInstallInfo; li != nil {
		dst.LastInstallInfo = &vimv1.ToolsConfigInfoToolsLastInstallInfo{
			Counter: li.Counter,
		}
		if li.Fault != nil {
			dst.LastInstallInfo.Fault = &vimv1.LocalizedMethodFault{
				LocalizedMessage: li.Fault.LocalizedMessage,
			}
		}
	}

	return nil
}

func convertVIMToCRD_VirtualHardwareSpec(
	dst *vimv1.VirtualHardwareSpec,
	src types.VirtualHardware) error {

	dst.NumCPU = src.NumCPU
	dst.NumCoresPerSocket = src.NumCoresPerSocket
	dst.AutoCoresPerSocket = src.AutoCoresPerSocket
	dst.Memory = *resource.NewQuantity(
		int64(src.MemoryMB)*1024*1024, resource.BinarySI)
	dst.VirtualICH7MPresent = src.VirtualICH7MPresent
	dst.VirtualSMCPresent = src.VirtualSMCPresent
	dst.MotherboardLayout = src.MotherboardLayout
	if src.SimultaneousThreads != 0 {
		v := src.SimultaneousThreads
		dst.SimultaneousThreads = &v
	}
	// TODO(akutz): dst.Devices = ...

	return nil
}

func convertVIMToCRD_VirtualMachineDefaultPowerOpInfo(
	dst *vimv1.VirtualMachineDefaultPowerOpInfo,
	src types.VirtualMachineDefaultPowerOpInfo) error {

	dst.PowerOffType = src.PowerOffType
	dst.SuspendType = src.SuspendType
	dst.ResetType = src.ResetType
	dst.DefaultPowerOffType = src.DefaultPowerOffType
	dst.DefaultSuspendType = src.DefaultSuspendType
	dst.DefaultResetType = src.DefaultResetType
	dst.StandbyAction = src.StandbyAction

	return nil
}

func convertVIMToCRD_VirtualMachineFileInfo(
	dst *vimv1.VirtualMachineFileInfo,
	src types.VirtualMachineFileInfo) error {

	dst.VMPathName = src.VmPathName
	dst.SnapshotDirectory = src.SnapshotDirectory
	dst.SuspendDirectory = src.SuspendDirectory
	dst.LogDirectory = src.LogDirectory
	dst.FTMetadataDirectory = src.FtMetadataDirectory

	return nil
}

func convertVIMToCRD_VirtualMachineSgxInfo(
	dst *vimv1.VirtualMachineSgxInfo,
	src types.VirtualMachineSgxInfo) error {

	dst.EpcSize = src.EpcSize
	dst.FlcMode = src.FlcMode
	dst.LePubKeyHash = src.LePubKeyHash
	dst.RequireAttestation = src.RequireAttestation

	return nil
}

func convertVIMToCRD_VirtualMachineVcpuConfig(
	dst *vimv1.VirtualMachineVcpuConfig,
	src types.VirtualMachineVcpuConfig) error {

	if src.LatencySensitivity != nil {
		dst.LatencySensitivity = &vimv1.LatencySensitivity{}
		if err := ConvertVIMToCRD(
			dst.LatencySensitivity, src.LatencySensitivity); err != nil {
			return err
		}
	}

	return nil
}

func convertVIMToCRD_VirtualMachineVirtualDeviceGroups(
	dst *vimv1.VirtualMachineVirtualDeviceGroups,
	src types.VirtualMachineVirtualDeviceGroups) error {

	for _, bdg := range src.DeviceGroup {
		dg := bdg.GetVirtualMachineVirtualDeviceGroupsDeviceGroup()
		vdg := vimv1.VirtualMachineVirtualDeviceGroupsDeviceGroup{
			GroupInstanceKey: dg.GroupInstanceKey,
		}
		if dg.DeviceInfo != nil {
			di := dg.DeviceInfo.GetDescription()
			vdg.DeviceInfo = &vimv1.VirtualDeviceDescription{
				Label:   di.Label,
				Summary: di.Summary,
			}
		}
		dst.DeviceGroup = append(dst.DeviceGroup, vdg)
	}

	return nil
}

func convertVIMToCRD_VirtualMachineVirtualDeviceSwap(
	dst *vimv1.VirtualMachineVirtualDeviceSwap,
	src types.VirtualMachineVirtualDeviceSwap) error {

	if src.LsiToPvscsi != nil {
		dst.LsiToPvscsi = &vimv1.VirtualMachineVirtualDeviceSwapDeviceSwapInfo{
			Enabled:    src.LsiToPvscsi.Enabled,
			Applicable: src.LsiToPvscsi.Applicable,
			Status:     src.LsiToPvscsi.Status,
		}
	}

	return nil
}

func convertVIMToCRD_VirtualMachineVirtualNumaInfo(
	dst *vimv1.VirtualMachineVirtualNumaInfo,
	src types.VirtualMachineVirtualNumaInfo) error {

	dst.AutoCoresPerNumaNode = src.AutoCoresPerNumaNode
	dst.CoresPerNumaNode = src.CoresPerNumaNode
	dst.VnumaOnCpuHotaddExposed = src.VnumaOnCpuHotaddExposed

	return nil
}
