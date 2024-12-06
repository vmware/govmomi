/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package ovf

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	ResourceSubTypeSoundCardSB16          = "vmware.soundcard.sb16"
	ResourceSubTypeSoundCardEnsoniq1371   = "vmware.soundcard.ensoniq1371"
	ResourceSubTypeSoundCardHDAudio       = "vmware.soundcard.hdaudio"
	ResourceSubTypePCIController          = "vmware.pcicontroller"
	ResourceSubTypePS2Controller          = "vmware.ps2controller"
	ResourceSubTypeSIOController          = "vmware.siocontroller"
	ResourceSubTypeKeyboard               = "vmware.keyboard"
	ResourceSubTypePointingDevice         = "vmware.pointingdevice"
	ResourceSubTypeVMCI                   = "vmware.vmci"
	ResourceSubTypeUSBEHCI                = "vmware.usb.ehci" /* USB 2.0 */
	ResourceSubTypeUSBXHCI                = "vmware.usb.xhci" /* USB 3.0 */
	ResourceSubTypeCdromISO               = "vmware.cdrom.iso"
	ResourceSubTypeCDROMRemotePassthrough = "vmware.cdrom.remotepassthrough"
	ResourceSubTypeCDROMRemoteATAPI       = "vmware.cdrom.remoteatapi"
	ResourceSubTypeCDROMPassthrough       = "vmware.cdrom.passthrough"
	ResourceSubTypeCDROMATAPI             = "vmware.cdrom.atapi"
	ResourceSubTypeFloppyDevice           = "vmware.floppy.device"
	ResourceSubTypeFloppyImage            = "vmware.floppy.image"
	ResourceSubTypeFloppyRemoveDevice     = "vmware.floppy.remotedevice"
	ResourceSubTypeSCSIPassthrough        = "vmware.scsi.passthrough"
	ResourceSubTypeParallelPortDevice     = "vmware.parallelport.device"
	ResourceSubTypeParallelPortFile       = "vmware.parallelport.file"
	ResourceSubTypeSerialPortDevice       = "vmware.serialport.device"
	ResourceSubTypeSerialPortFile         = "vmware.serialport.file"
	ResourceSubTypeSerialPortPipe         = "vmware.serialport.pipe"
	ResourceSubTypeSerialPortURI          = "vmware.serialport.uri"
	ResourceSubTypeSerialPortThinPrint    = "vmware.serialport.thinprint"
	ResourceSubTypeSATAAHCI               = "vmware.sata.ahci"
	ResourceSubTypeSATAAHCIAlter          = "ahci"
	ResourceSubTypeNVMEController         = "vmware.nvme.controller"
	ResourceSubTypeNVDIMMController       = "vmware.nvdimm.controller"
	ResourceSubTypeNVDIMMDevice           = "vmware.nvdimm.device"
	ResourceSubTypePCIPassthrough         = "vmware.pci.passthrough"
	ResourceSubTypePCIPassthroughDVX      = "vmware.pci.passthrough-dvx"
	ResourceSubTypePCIPassthroughAH       = "vmware.pci.passthrough-ah"
	ResourceSubTypePCIPassthroughVMIOP    = "vmware.pci.passthrough-vmiop"
	ResourceSubTypePrecisionClock         = "vmware.precisionclock"
	ResourceSubTypeWatchdogTimer          = "vmware.watchdogtimer"
	ResourceSubTypeVTPM                   = "vmware.vtpm"
)

// ErrUnsupportedItem is returned by Envelope.ToConfigSpec when there is an
// invalid item configuration.
type ErrUnsupportedItem struct {
	Name             string
	Index            int
	InstanceID       string
	DeviceType       CIMResourceType
	LocalizedMessage string
}

func (e ErrUnsupportedItem) Error() string {
	msg := fmt.Sprintf(
		"unsupported item name=%q, index=%d, instanceID=%q",
		e.Name, e.Index, e.InstanceID)
	if e.DeviceType > 0 {
		msg = fmt.Sprintf("%s, type=%d", msg, e.DeviceType)
	}
	if e.LocalizedMessage != "" {
		msg = fmt.Sprintf("%s, msg=%q", msg, e.LocalizedMessage)
	}
	return msg
}

type configSpec = types.VirtualMachineConfigSpec

// ToConfigSpec transforms the envelope into a ConfigSpec that may be used to
// create a new virtual machine.
// Please note, at this time:
//   - Only a single VirtualSystem is supported. The VirtualSystemCollection
//     section is ignored.
//   - Only the first VirtualHardware section is supported.
//   - Only the default deployment option configuration is considered. Elements
//     part of a non-default configuration are ignored.
//   - Disks must specify zero or one HostResource elements.
//   - Many, many more constraints...
func (e Envelope) ToConfigSpec() (types.VirtualMachineConfigSpec, error) {

	vs := e.VirtualSystem
	if vs == nil {
		return configSpec{}, errors.New("no VirtualSystem")
	}

	// Determine if there is a default configuration.
	var defaultConfigName string
	if do := e.DeploymentOption; do != nil {
		for _, c := range do.Configuration {
			if d := c.Default; d != nil && *d {
				defaultConfigName = c.ID
				break
			}
		}
	}

	dst := configSpec{
		Files: &types.VirtualMachineFileInfo{},
		Name:  vs.ID,
	}

	// Set the guest ID.
	if os := vs.OperatingSystem; os != nil && os.OSType != nil {
		dst.GuestId = *os.OSType
	}

	// Parse the hardware.
	if err := e.toHardware(&dst, defaultConfigName, vs); err != nil {
		return configSpec{}, err
	}

	// Parse the vApp config.
	if err := e.toVAppConfig(&dst, defaultConfigName, vs); err != nil {
		return configSpec{}, err
	}

	return dst, nil
}

func (e Envelope) toHardware(
	dst *configSpec,
	configName string,
	vs *VirtualSystem) error {

	var hw VirtualHardwareSection
	if len(vs.VirtualHardware) == 0 {
		return errors.New("no VirtualHardware")
	}
	hw = vs.VirtualHardware[0]

	// Set the hardware version.
	if vmx := hw.System.VirtualSystemType; vmx != nil {
		dst.Version = *vmx
	}

	// Parse the config
	e.toConfig(dst, hw)

	// Parse the extra config.
	e.toExtraConfig(dst, hw)

	var (
		devices   object.VirtualDeviceList
		resources = map[string]types.BaseVirtualDevice{}
	)

	for index := range hw.Item {
		item := hw.Item[index]

		if c := item.Configuration; c != nil {
			if *c != configName {
				// Skip items that do not belong to the provided config.
				continue
			}
		}

		rt := item.ResourceType
		if rt == nil {
			return e.errUnsupportedItem(
				index, item, nil, "nil ResourceType")
		}

		var (
			d   types.BaseVirtualDevice
			err error
		)

		switch *rt {
		case Other:
			d, err = e.toOther(item, devices, resources)

		case Processor:
			if item.VirtualQuantity == nil {
				return e.errUnsupportedItem(
					index, item, nil, "nil VirtualQuantity")
			}
			dst.NumCPUs = int32(*item.VirtualQuantity)
			if cps := item.CoresPerSocket; cps != nil {
				dst.NumCoresPerSocket = cps.Value
			}

		case Memory:
			if item.VirtualQuantity == nil {
				return e.errUnsupportedItem(
					index, item, nil, "nil VirtualQuantity")
			}
			dst.MemoryMB = int64(*item.VirtualQuantity)

		case IdeController:
			d, err = e.toIDEController(item, devices, resources)

		case ParallelScsiHba:
			d, err = e.toSCSIController(item, devices, resources)

		case EthernetAdapter:
			d, err = e.toNetworkInterface(item, devices, resources)

		case FloppyDrive:
			if devices.PickController((*types.VirtualSIOController)(nil)) == nil {
				c := &types.VirtualSIOController{}
				c.Key = devices.NewKey()
				devices = append(devices, c)
			}
			d, err = e.toFloppyDrive(item, devices, resources)

		case CdDrive, DvdDrive:
			d, err = e.toCDOrDVDDrive(item, devices, resources)

		case DiskDrive:
			d, err = e.toVirtualDisk(item, devices, resources)

		case OtherStorage:
			d, err = e.toOtherStorage(item, devices, resources)

		case UsbController:
			d, err = e.toUSB(item, devices, resources)

		case Graphics:
			d, err = e.toVideoCard(item, devices, resources)

		default:
			return e.errUnsupportedItem(
				index, item, nil, "unsupported resource type")
		}

		if err != nil {
			return e.errUnsupportedItem(index, item, err)
		}

		if d != nil {
			setConnectable(d, item)
			if err := e.setUnitNumber(item, d); err != nil {
				return e.errUnsupportedItem(index, item, err)
			}
			if err := e.setPCISlotNumber(item, d); err != nil {
				return e.errUnsupportedItem(index, item, err)
			}
			devices = append(devices, d)
		}
	}

	// Add the devices to the ConfigSpec.
	dst.DeviceChange, _ = devices.ConfigSpec(types.VirtualDeviceConfigSpecOperationAdd)

	return nil
}

func (e Envelope) errUnsupportedItem(
	index int,
	item ResourceAllocationSettingData,
	inner error,
	args ...any) error {

	err := ErrUnsupportedItem{
		Name:       item.ElementName,
		InstanceID: item.InstanceID,
		Index:      index,
	}

	if item.ResourceType != nil {
		err.DeviceType = *item.ResourceType
	}

	if len(args) == 1 {
		err.LocalizedMessage = args[0].(string)
	} else if len(args) > 1 {
		err.LocalizedMessage = fmt.Sprintf(args[0].(string), args[1:]...)
	}

	if inner != nil {
		return fmt.Errorf("%w, inner=%w", err, inner)
	}

	return err
}

func (e Envelope) setUnitNumber(
	item ResourceAllocationSettingData,
	d types.BaseVirtualDevice) error {

	if item.AddressOnParent == nil || *item.AddressOnParent == "" {
		return nil
	}

	unitNumber, err := strconv.ParseInt(*item.AddressOnParent, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid AddressOnParent=%q", *item.AddressOnParent)
	}

	d.GetVirtualDevice().UnitNumber = types.NewInt32(int32(unitNumber))
	return nil
}

func (e Envelope) setBusNumber(
	item ResourceAllocationSettingData,
	d types.BaseVirtualDevice) error {

	if item.Address == nil || *item.Address == "" {
		return nil
	}

	c, ok := d.(types.BaseVirtualController)
	if !ok {
		return fmt.Errorf("expectedType=%s, actualType=%T",
			"types.BaseVirtualController", d)
	}

	busNumber, err := strconv.ParseInt(*item.Address, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid Address=%q", *item.Address)
	}

	c.GetVirtualController().BusNumber = int32(busNumber)
	return nil
}

func (e Envelope) setPCISlotNumber(
	item ResourceAllocationSettingData,
	d types.BaseVirtualDevice) error {

	var pciSlotNumber int32 = -1

	for i := range item.Config {
		c := item.Config[i]
		if c.Key == "slotInfo.pciSlotNumber" {
			if c.Value != "" {
				v, err := strconv.ParseInt(c.Value, 10, 32)
				if err != nil {
					return fmt.Errorf("invalid pci slot number %s", c.Value)
				}
				pciSlotNumber = int32(v)
			}
			break
		}
	}

	if pciSlotNumber >= 0 {
		vd := d.GetVirtualDevice()
		if vd.SlotInfo == nil {
			vd.SlotInfo = &types.VirtualDevicePciBusSlotInfo{}
		}
		si, ok := vd.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
		if !ok {
			return fmt.Errorf("expectedType=%s, actualType=%T",
				"*types.VirtualDevicePciBusSlotInfo", vd.SlotInfo)
		}
		si.PciSlotNumber = pciSlotNumber
	}

	return nil
}

func (e Envelope) ovfDisk(diskID string) *VirtualDiskDesc {
	for _, disk := range e.Disk.Disks {
		if strings.HasSuffix(diskID, disk.DiskID) {
			return &disk
		}
	}
	return nil
}

func (e Envelope) toVirtualDisk(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	if item.Parent == nil {
		return nil, fmt.Errorf("missing Parent")
	}

	r, ok := resources[*item.Parent]
	if !ok {
		return nil, nil
	}

	c, ok := r.(types.BaseVirtualController)
	if !ok {
		return nil, fmt.Errorf("expectedType=%s, actualType=%T",
			"types.BaseVirtualController", r)
	}

	d := devices.CreateDisk(c, types.ManagedObjectReference{}, "")

	devices.AssignController(d, c)

	d.VirtualDevice.DeviceInfo = &types.Description{
		Label: item.ElementName,
	}

	// Find the disk's capacity.
	var capacityInBytes uint64
	switch len(item.HostResource) {

	case 0:
		var allocUnitsSz string
		if item.AllocationUnits != nil {
			allocUnitsSz = *item.AllocationUnits
		}
		capacityInBytes = uint64(ParseCapacityAllocationUnits(allocUnitsSz))
		if r := item.VirtualQuantity; r != nil {
			capacityInBytes *= uint64(*r)
		}

	case 1:
		diskID := item.HostResource[0]
		dd := e.ovfDisk(diskID)
		if dd == nil {
			return nil, fmt.Errorf("missing diskID %q", diskID)
		}

		var allocUnitsSz string
		if dd.CapacityAllocationUnits != nil {
			allocUnitsSz = *dd.CapacityAllocationUnits
		}
		capacityInBytes = uint64(ParseCapacityAllocationUnits(allocUnitsSz))
		if capSz := dd.Capacity; capSz != "" {
			cap, err := strconv.ParseUint(dd.Capacity, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("disk=%s has invalid capacity=%q",
					diskID, capSz)
			}
			capacityInBytes *= cap
		}

	default:
		return nil, fmt.Errorf("multiple HostResource elements")
	}

	if capacityInBytes > math.MaxInt64 {
		return nil, fmt.Errorf(
			"capacityInBytes=%d exceeds math.MaxInt64", capacityInBytes)
	}

	d.CapacityInBytes = int64(capacityInBytes)

	return d, nil
}

func (e Envelope) toCDOrDVDDrive(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	if item.Parent == nil {
		return nil, fmt.Errorf("missing Parent")
	}

	r, ok := resources[*item.Parent]
	if !ok {
		return nil, nil // Parent is unsupported
	}

	c, ok := r.(*types.VirtualIDEController)
	if !ok {
		return nil, fmt.Errorf("expectedType=%s, actualType=%T",
			"*types.VirtualIDEController", r)
	}

	d, err := devices.CreateCdrom(c)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (e Envelope) toSCSIController(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateSCSIController(item.toResourceSubType())
	if err != nil {
		return nil, err
	}
	if err := e.setBusNumber(item, d); err != nil {
		return nil, err
	}
	resources[item.InstanceID] = d

	return d, nil
}

func (e Envelope) toIDEController(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateIDEController()
	if err != nil {
		return nil, err
	}
	if err := e.setBusNumber(item, d); err != nil {
		return nil, err
	}
	resources[item.InstanceID] = d
	return d, nil
}

func (e Envelope) toOtherStorage(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	rst := item.toResourceSubType()
	switch rst {
	case ResourceSubTypeSATAAHCI, ResourceSubTypeSATAAHCIAlter:
		return e.toSATAController(item, devices, resources)
	case ResourceSubTypeNVMEController:
		return e.toNVMEController(item, devices, resources)
	}
	return nil, fmt.Errorf("unknown subtype %q", rst)
}

func (e Envelope) toSATAController(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateSATAController()
	if err != nil {
		return nil, err
	}
	if err := e.setBusNumber(item, d); err != nil {
		return nil, err
	}
	resources[item.InstanceID] = d
	return d, nil
}

func (e Envelope) toNVMEController(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateNVMEController()
	if err != nil {
		return nil, err
	}
	if err := e.setBusNumber(item, d); err != nil {
		return nil, err
	}
	resources[item.InstanceID] = d
	return d, nil
}

func (e Envelope) toUSB(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	var d types.BaseVirtualDevice

	vc := types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: devices.NewKey(),
		},
	}

	rst := item.toResourceSubType()
	switch rst {
	case ResourceSubTypeUSBEHCI:
		c := &types.VirtualUSBController{VirtualController: vc}
		for i := range item.Config {
			ic := item.Config[i]
			switch ic.Key {
			case "autoConnectDevices":
				c.AutoConnectDevices = szToBoolPtr(ic.Value)
			case "ehciEnabled":
				c.EhciEnabled = szToBoolPtr(ic.Value)
			}
		}
		d = c
	case ResourceSubTypeUSBXHCI:
		c := &types.VirtualUSBXHCIController{VirtualController: vc}
		for i := range item.Config {
			ic := item.Config[i]
			switch ic.Key {
			case "autoConnectDevices":
				c.AutoConnectDevices = szToBoolPtr(ic.Value)
			}
		}
		d = c
	default:
		return nil, fmt.Errorf("invalid subtype %q", rst)
	}

	if err := e.setBusNumber(item, d); err != nil {
		return nil, err
	}

	resources[item.InstanceID] = d
	return d, nil
}

func (e Envelope) toNetworkInterface(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	_ map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateEthernetCard(item.toResourceSubType(), nil)
	if err != nil {
		return nil, err
	}

	nic := d.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()

	for i := range item.Config {
		c := item.Config[i]
		switch c.Key {
		case "wakeOnLanEnabled":
			nic.WakeOnLanEnabled = szToBoolPtr(c.Value)
		case "uptCompatibilityEnabled":
			nic.UptCompatibilityEnabled = szToBoolPtr(c.Value)
		}
	}

	return d, nil
}

func (e Envelope) toFloppyDrive(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateFloppy()
	if err != nil {
		return nil, err
	}
	resources[item.InstanceID] = d

	return d, nil
}

func (e Envelope) toVideoCard(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	_ map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d := &types.VirtualMachineVideoCard{
		VirtualDevice: types.VirtualDevice{
			Key: devices.NewKey(),
		},
	}

	for i := range item.Config {
		c := item.Config[i]
		switch c.Key {
		case "enable3DSupport":
			d.Enable3DSupport = szToBoolPtr(c.Value)
		case "graphicsMemorySizeInKB":
			v, err := strconv.ParseInt(c.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %q=%s", c.Key, c.Value)
			}
			d.GraphicsMemorySizeInKB = v
		case "useAutoDetect":
			d.UseAutoDetect = szToBoolPtr(c.Value)
		case "videoRamSizeInKB":
			v, err := strconv.ParseInt(c.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid %q=%s", c.Key, c.Value)
			}
			d.VideoRamSizeInKB = v
		case "numDisplays":
			v, err := strconv.ParseInt(c.Value, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid %q=%s", c.Key, c.Value)
			}
			d.NumDisplays = int32(v)
		case "use3dRenderer":
			d.Use3dRenderer = c.Value
		}
	}

	return d, nil
}

func (e Envelope) toOther(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	switch item.toResourceSubType() {
	case ResourceSubTypeVMCI:
		return e.toVMCI(item, devices, resources)
	}
	return nil, nil
}

func (e Envelope) toVMCI(
	item ResourceAllocationSettingData,
	devices object.VirtualDeviceList,
	_ map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d := &types.VirtualMachineVMCIDevice{
		VirtualDevice: types.VirtualDevice{
			Key: devices.NewKey(),
		},
	}

	for i := range item.Config {
		c := item.Config[i]
		switch c.Key {
		case "allowUnrestrictedCommunication":
			d.AllowUnrestrictedCommunication = szToBoolPtr(c.Value)
		}
	}

	return d, nil
}

func (r *ResourceAllocationSettingData) toResourceSubType() string {
	rst := "unknown"
	if r != nil {
		if r.ResourceSubType != nil {
			rst = strings.ToLower(*r.ResourceSubType)
		}
	}
	return rst
}

func (e Envelope) toConfig(
	dst *configSpec,
	hw VirtualHardwareSection) {

	for i := range hw.Config {
		c := hw.Config[i]
		switch c.Key {
		case "cpuHotAddEnabled":
			dst.CpuHotAddEnabled = szToBoolPtr(c.Value)
		case "cpuHotRemoveEnabled":
			dst.CpuHotRemoveEnabled = szToBoolPtr(c.Value)
		case "bootOptions.efiSecureBootEnabled":
			initBootOptions(dst)
			dst.BootOptions.EfiSecureBootEnabled = szToBoolPtr(c.Value)
		case "firmware":
			dst.Firmware = c.Value
		case "flags.vbsEnabled":
			initFlags(dst)
			dst.Flags.VbsEnabled = szToBoolPtr(c.Value)
		case "flags.vvtdEnabled":
			initFlags(dst)
			dst.Flags.VvtdEnabled = szToBoolPtr(c.Value)
		case "memoryHotAddEnabled":
			dst.MemoryHotAddEnabled = szToBoolPtr(c.Value)
		case "nestedHVEnabled":
			dst.NestedHVEnabled = szToBoolPtr(c.Value)
		case "virtualICH7MPresent":
			dst.VirtualICH7MPresent = szToBoolPtr(c.Value)
		case "virtualSMCPresent":
			dst.VirtualSMCPresent = szToBoolPtr(c.Value)
		case "cpuAllocation.shares.shares":
			initCPUAllocationShares(dst)
			dst.CpuAllocation.Shares.Shares = szToInt32(c.Value)
		case "cpuAllocation.shares.level":
			initCPUAllocationShares(dst)
			dst.CpuAllocation.Shares.Level = types.SharesLevel(c.Value)
		case "simultaneousThreads":
			dst.SimultaneousThreads = szToInt32(c.Value)
		case "tools.syncTimeWithHost":
			initToolsConfig(dst)
			dst.Tools.SyncTimeWithHost = szToBoolPtr(c.Value)
		case "tools.syncTimeWithHostAllowed":
			initToolsConfig(dst)
			dst.Tools.SyncTimeWithHostAllowed = szToBoolPtr(c.Value)
		case "tools.afterPowerOn":
			initToolsConfig(dst)
			dst.Tools.AfterPowerOn = szToBoolPtr(c.Value)
		case "tools.afterResume":
			initToolsConfig(dst)
			dst.Tools.AfterResume = szToBoolPtr(c.Value)
		case "tools.beforeGuestShutdown":
			initToolsConfig(dst)
			dst.Tools.BeforeGuestShutdown = szToBoolPtr(c.Value)
		case "tools.beforeGuestStandby":
			initToolsConfig(dst)
			dst.Tools.BeforeGuestStandby = szToBoolPtr(c.Value)
		case "tools.toolsUpgradePolicy":
			initToolsConfig(dst)
			dst.Tools.ToolsUpgradePolicy = c.Value
		case "powerOpInfo.powerOffType":
			initPowerOpInfo(dst)
			dst.PowerOpInfo.PowerOffType = c.Value
		case "powerOpInfo.resetType":
			initPowerOpInfo(dst)
			dst.PowerOpInfo.ResetType = c.Value
		case "powerOpInfo.suspendType":
			initPowerOpInfo(dst)
			dst.PowerOpInfo.SuspendType = c.Value
		case "powerOpInfo.standbyAction":
			initPowerOpInfo(dst)
			dst.PowerOpInfo.StandbyAction = c.Value
		case "vPMCEnabled":
			dst.VPMCEnabled = szToBoolPtr(c.Value)
		}
	}
}

func (e Envelope) toExtraConfig(
	dst *configSpec,
	hw VirtualHardwareSection) {

	var newEC object.OptionValueList
	for i := range hw.ExtraConfig {
		newEC = append(newEC, &types.OptionValue{
			Key:   hw.ExtraConfig[i].Key,
			Value: hw.ExtraConfig[i].Value,
		})
	}
	dst.ExtraConfig = newEC.Join(dst.ExtraConfig...)
}

func initToolsConfig(dst *configSpec) {
	if dst.Tools == nil {
		dst.Tools = &types.ToolsConfigInfo{}
	}
}

func initPowerOpInfo(dst *configSpec) {
	if dst.PowerOpInfo == nil {
		dst.PowerOpInfo = &types.VirtualMachineDefaultPowerOpInfo{}
	}
}

func initCPUAllocation(dst *configSpec) {
	if dst.CpuAllocation == nil {
		dst.CpuAllocation = &types.ResourceAllocationInfo{}
	}
}

func initCPUAllocationShares(dst *configSpec) {
	initCPUAllocation(dst)
	if dst.CpuAllocation.Shares == nil {
		dst.CpuAllocation.Shares = &types.SharesInfo{}
	}
}

func initFlags(dst *configSpec) {
	if dst.Flags == nil {
		dst.Flags = &types.VirtualMachineFlagInfo{}
	}
}

func initBootOptions(dst *configSpec) {
	if dst.BootOptions == nil {
		dst.BootOptions = &types.VirtualMachineBootOptions{}
	}
}

func setConnectable(
	dst types.BaseVirtualDevice,
	src ResourceAllocationSettingData) {

	d := dst.GetVirtualDevice()
	for i := range src.Config {
		c := src.Config[i]
		switch c.Key {
		case "connectable.allowGuestControl":
			if d.Connectable == nil {
				d.Connectable = &types.VirtualDeviceConnectInfo{}
			}
			d.Connectable.AllowGuestControl = szToBool(c.Value)
		}
	}
}

func szToBoolPtr(s string) *bool {
	if s == "" {
		return nil
	}
	b, _ := strconv.ParseBool(s)
	return &b
}

func szToBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func szToInt32(s string) int32 {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int32(v)
}

func deref[T any](pT *T) T {
	var t T
	if pT != nil {
		t = *pT
	}
	return t
}

func (e Envelope) toVAppConfig(
	dst *configSpec,
	configName string,
	vs *VirtualSystem) error {

	if len(vs.Product) == 0 {
		return nil
	}

	vapp := &types.VAppConfigSpec{}

	index := 0
	for i, product := range vs.Product {
		vapp.Product = append(vapp.Product, types.VAppProductSpec{
			ArrayUpdateSpec: types.ArrayUpdateSpec{
				Operation: types.ArrayUpdateOperationAdd,
			},
			Info: &types.VAppProductInfo{
				Key:         int32(i),
				ClassId:     deref(product.Class),
				InstanceId:  deref(product.Instance),
				Name:        product.Product,
				Vendor:      product.Vendor,
				Version:     product.Version,
				FullVersion: product.FullVersion,
				ProductUrl:  product.ProductURL,
				VendorUrl:   product.VendorURL,
				AppUrl:      product.AppURL,
			},
		})

		for _, p := range product.Property {
			if p.Configuration != nil && *p.Configuration != configName {
				// Skip properties that are not part of the provided
				// configuration.
				continue
			}
			vapp.Property = append(vapp.Property, types.VAppPropertySpec{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: types.ArrayUpdateOperationAdd,
				},
				Info: &types.VAppPropertyInfo{
					Key:              int32(index),
					ClassId:          deref(product.Class),
					InstanceId:       deref(product.Instance),
					Id:               p.Key,
					Category:         product.Category,
					Label:            deref(p.Label),
					Type:             p.Type,
					UserConfigurable: p.UserConfigurable,
					DefaultValue:     deref(p.Default),
					Value:            "",
					Description:      deref(p.Description),
				},
			})
			index++
		}
	}

	dst.VAppConfig = vapp
	return nil
}
