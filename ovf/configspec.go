// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"errors"
	"fmt"
	"math"
	"path"
	"regexp"
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

var errUnsupportedResourceSubtype = errors.New("unsupported resource subtype")

// ErrUnsupportedItem is returned by Envelope.ToConfigSpec when there is an
// invalid item configuration.
type ErrUnsupportedItem struct {
	Name             string
	Index            int
	InstanceID       string
	ResourceType     CIMResourceType
	ResourceSubType  string
	LocalizedMessage string
}

func (e ErrUnsupportedItem) Error() string {
	msg := fmt.Sprintf(
		"unsupported item name=%q, index=%d, instanceID=%q",
		e.Name, e.Index, e.InstanceID)
	if e.ResourceType > 0 {
		msg = fmt.Sprintf("%s, resourceType=%d", msg, e.ResourceType)
	}
	if e.ResourceSubType != "" {
		msg = fmt.Sprintf("%s, resourceSubType=%s", msg, e.ResourceSubType)
	}
	if e.LocalizedMessage != "" {
		msg = fmt.Sprintf("%s, msg=%q", msg, e.LocalizedMessage)
	}
	return msg
}

// AsErrUnsupportedItem returns any possible wrapped ErrUnsupportedItem error
// from the provided error.
func AsErrUnsupportedItem(in error) (ErrUnsupportedItem, bool) {
	var out ErrUnsupportedItem
	if errors.As(in, &out) {
		return out, true
	}
	return ErrUnsupportedItem{}, false
}

func errUnsupportedItem(
	index int,
	item itemElement,
	inner error,
	args ...any) error {

	err := ErrUnsupportedItem{
		Name:            item.ElementName,
		InstanceID:      item.InstanceID,
		Index:           index,
		ResourceSubType: item.resourceSubType,
	}

	if item.ResourceType != nil {
		err.ResourceType = *item.ResourceType
	}

	if len(args) == 1 {
		err.LocalizedMessage = args[0].(string)
	} else if len(args) > 1 {
		err.LocalizedMessage = fmt.Sprintf(args[0].(string), args[1:]...)
	}

	if inner != nil {
		return fmt.Errorf("%w, %w", err, inner)
	}

	return err
}

type itemElement struct {
	ResourceAllocationSettingData

	resourceSubType string
}

type configSpec = types.VirtualMachineConfigSpec

// ToConfigSpecOptions influence the behavior of the ToConfigSpecWithOptions
// function.
type ToConfigSpecOptions struct {

	// Strict indicates that an error should be returned on Item elements in
	// a VirtualHardware section that have an unknown ResourceType, i.e. a value
	// that falls outside the range of the enum CIMResourceType.
	Strict bool

	// VirtualSystemCollectionIndex specifies the index of the VirtualSystem in
	// the OVF's VirtualSystemCollection to transform into the ConfigSpec.
	VirtualSystemCollectionIndex int

	// DeploymentConfiguration specifies the deployment configuration name (see
	// DeploymentOptionSection in DSP0243). If empty, the default configuration
	// is used; if none is marked default, the first configuration is used.
	// Only elements (e.g. VirtualHardware Item, ProductSection Property)
	// that match this configuration or have no configuration are included.
	DeploymentConfiguration string
}

// ToConfigSpec calls ToConfigSpecWithOptions with an empty ToConfigSpecOptions
// object.
func (e Envelope) ToConfigSpec() (types.VirtualMachineConfigSpec, error) {
	return e.ToConfigSpecWithOptions(ToConfigSpecOptions{})
}

// resolveDeploymentConfiguration returns the deployment configuration name to
// use per DSP0243: opts.DeploymentConfiguration if set (and valid), else the
// default configuration, else the first configuration.
func (e Envelope) resolveDeploymentConfiguration(
	opts ToConfigSpecOptions) (string, error) {
	if opts.DeploymentConfiguration != "" {
		if do := e.DeploymentOption; do == nil || len(do.Configuration) == 0 {
			return "", fmt.Errorf(
				"deployment configuration %q specified but no DeploymentOptionSection",
				opts.DeploymentConfiguration)
		}
		for _, c := range e.DeploymentOption.Configuration {
			if c.ID == opts.DeploymentConfiguration {
				return c.ID, nil
			}
		}
		return "", fmt.Errorf(
			"deployment configuration %q not found in DeploymentOptionSection",
			opts.DeploymentConfiguration)
	}
	if do := e.DeploymentOption; do != nil && len(do.Configuration) > 0 {
		for _, c := range do.Configuration {
			if d := c.Default; d != nil && *d {
				return c.ID, nil
			}
		}
		return do.Configuration[0].ID, nil
	}
	return "", nil
}

// ToConfigSpecWithOptions transforms the envelope into a ConfigSpec that may be
// used to create a new virtual machine.
// Please note, at this time:
//   - Only a single VirtualSystem is supported. The VirtualSystemCollection
//     section is ignored.
//   - Only the first VirtualHardware section is supported.
//   - Deployment configuration is selected via opts.DeploymentConfiguration
//     (default or first if empty). Elements that are part of another
//     configuration are excluded.
//   - Disks must specify zero or one HostResource elements.
//   - Many, many more constraints...
func (e Envelope) ToConfigSpecWithOptions(
	opts ToConfigSpecOptions) (types.VirtualMachineConfigSpec, error) {

	fileRefs := make(map[string]File, len(e.References))
	for _, ref := range e.References {
		fileRefs[ref.ID] = ref
	}

	vs := e.VirtualSystem
	if vs == nil {
		if vsc := e.VirtualSystemCollection; vsc != nil {
			i := opts.VirtualSystemCollectionIndex
			if len(vsc.VirtualSystem) > i {
				vs = &vsc.VirtualSystem[i]

				// If the VirtualSystem from the VirtualSystemCollection
				// has no Product, use the Product from the
				// VirtualSystemCollection.
				if vs.Product == nil {
					vs.Product = vsc.Product
				}
			}
		}
	}

	if vs == nil {
		return configSpec{}, nil
	}

	configName, err := e.resolveDeploymentConfiguration(opts)
	if err != nil {
		return configSpec{}, err
	}

	dst := configSpec{
		Files: &types.VirtualMachineFileInfo{},
		Name:  vs.ID,
	}

	// Set the guest ID from vmw:osType if present, otherwise derive from the
	// DMTF-standard CIM OS type integer (ovf:id) and optional version string.
	if os := vs.OperatingSystem; os != nil {
		if os.OSType != nil {
			dst.GuestId = *os.OSType
		} else {
			ver := ""
			if os.Version != nil {
				ver = *os.Version
			}
			if guestID := CIMOSTypeToGuestID(CIMOSType(os.ID), ver); guestID != "" {
				dst.GuestId = string(guestID)
			}
		}
	}

	// Parse the hardware.
	if err := e.toHardware(
		&dst, configName, vs, fileRefs, opts); err != nil {

		return configSpec{}, err
	}

	// Parse the vApp config.
	if err := e.toVAppConfig(&dst, configName, vs); err != nil {
		return configSpec{}, err
	}

	return dst, nil
}

func (e Envelope) toHardware(
	dst *configSpec,
	configName string,
	vs *VirtualSystem,
	fileRefs map[string]File,
	opts ToConfigSpecOptions) error {

	var hw VirtualHardwareSection
	if len(vs.VirtualHardware) == 0 {
		return nil
	}

	hw = vs.VirtualHardware[0]

	// Set the hardware version.
	if hw.System != nil {
		if vmx := hw.System.VirtualSystemType; vmx != nil {
			dst.Version = *vmx
		}
	}

	// Parse the config
	e.toConfig(dst, hw)

	// Parse the extra config.
	e.toExtraConfig(dst, hw)

	// Property defaults for the selected config (used e.g. for disk capacity "${key}").
	propertyValues := e.propertyDefaultsForConfig(vs, configName)

	var (
		devices   object.VirtualDeviceList
		resources = map[string]types.BaseVirtualDevice{}
	)

	for index := range hw.Item {
		item := itemElement{
			ResourceAllocationSettingData: hw.Item[index],
		}

		if c := item.Configuration; c != nil {
			if *c != configName {
				// Skip items that do not belong to the provided config.
				continue
			}
		}

		if item.ResourceType == nil {
			return errUnsupportedItem(index, item, nil, "nil ResourceType")
		}

		// Get the resource sub type, if any.
		if rst := item.ResourceSubType; rst != nil {
			item.resourceSubType = strings.ToLower(*rst)
		}

		var (
			d   types.BaseVirtualDevice
			err error
		)

		switch *item.ResourceType {

		case Other: // 1
			d, err = e.toOther(item, devices, resources)

		case ComputerSystem: // 2
			// TODO(akutz)

		case Processor: // 3
			if item.VirtualQuantity != nil {
				dst.NumCPUs = int32(*item.VirtualQuantity)
			} else if item.Reservation != nil {
				dst.NumCPUs = int32(*item.Reservation)
			} else {
				return errUnsupportedItem(
					index, item, nil, "nil VirtualQuantity and Reservation")
			}
			if item.CoresPerSocket != nil {
				dst.NumCoresPerSocket = &item.CoresPerSocket.Value
			}

		case Memory: // 4
			if item.VirtualQuantity != nil {
				dst.MemoryMB = int64(*item.VirtualQuantity)
			} else if item.Reservation != nil {
				dst.MemoryMB = int64(*item.Reservation)
			} else {
				return errUnsupportedItem(
					index, item, nil, "nil VirtualQuantity and Reservation")
			}

		case IdeController: // 5
			d, err = e.toIDEController(item, devices, resources)

		case ParallelScsiHba: // 6
			d, err = e.toSCSIController(item, devices, resources)

		case FcHba: // 7
			// TODO(akutz)

		case IScsiHba: // 8
			// TODO(akutz)

		case IbHba: // 9
			// TODO(akutz)

		case EthernetAdapter: // 10
			d, err = e.toNetworkInterface(item, devices, resources)

		case OtherNetwork: // 11
			// TODO(akutz)

		case IoSlot: // 12
			// TODO(akutz)

		case IoDevice: // 13
			// TODO(akutz)

		case FloppyDrive: // 14
			if devices.PickController((*types.VirtualSIOController)(nil)) == nil {
				c := &types.VirtualSIOController{}
				c.Key = devices.NewKey()
				devices = append(devices, c)
			}
			d, err = e.toFloppyDrive(item, devices, resources)

		case CdDrive, DvdDrive: // 15, 16
			d, err = e.toCDOrDVDDrive(item, devices, resources)

		case DiskDrive: // 17
			d, err = e.toVirtualDisk(item, devices, resources, fileRefs, propertyValues)

		case TapeDrive: // 18
			// TODO(akutz)

		case StorageExtent: // 19
			// TODO(akutz)

		case OtherStorage: // 20
			d, err = e.toOtherStorage(item, devices, resources)

		case SerialPort: // 21
			// TODO(akutz)

		case ParallelPort: // 22
			// TODO(akutz)

		case UsbController: // 23
			d, err = e.toUSB(item, devices, resources)

		case Graphics: // 24
			d, err = e.toVideoCard(item, devices, resources)

		case Ieee1394: // 25
			// TODO(akutz)

		case PartitionableUnit: // 26
			// TODO(akutz)

		case BasePartitionable: // 27
			// TODO(akutz)

		case PowerSupply: // 28
			// TODO(akutz)

		case CoolingDevice: // 29
			// TODO(akutz)

		case EthernetSwitchPort: // 30
			// TODO(akutz)

		case LogicalDisk: // 31
			// TODO(akutz)

		case StorageVolume: // 32
			// TODO(akutz)

		case EthernetConnection: // 33
			// TODO(akutz)

		default:
			if opts.Strict {
				return errUnsupportedItem(
					index, item, nil, "unsupported resource type")
			}
		}

		if err != nil {
			if err == errUnsupportedResourceSubtype {
				if !opts.Strict {
					continue
				}
			}
			return errUnsupportedItem(index, item, err)
		}

		if d != nil {
			setConnectable(d, item)
			if err := e.setUnitNumber(item, d); err != nil {
				return errUnsupportedItem(index, item, err)
			}
			if err := e.setPCISlotNumber(item, d); err != nil {
				return errUnsupportedItem(index, item, err)
			}
			devices = append(devices, d)
		}
	}

	// Add the devices to the ConfigSpec.
	dst.DeviceChange = make([]types.BaseVirtualDeviceConfigSpec, len(devices))
	for i, d := range devices {
		vdcs := &types.VirtualDeviceConfigSpec{
			Device:    devices[i],
			Operation: types.VirtualDeviceConfigSpecOperationAdd,
		}
		if disk, ok := d.(*types.VirtualDisk); ok {
			vdcs.FileOperation = types.VirtualDeviceConfigSpecFileOperationCreate
			if b, ok := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
				if b.FileName != "" {
					// Set the file operation to empty for existing disks.
					vdcs.FileOperation = ""
				}
			}
		}
		dst.DeviceChange[i] = vdcs
	}

	return nil
}

func (e Envelope) setUnitNumber(
	item itemElement,
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
	item itemElement,
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
	item itemElement,
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

// propertyDefaultsForConfig returns a map of property key to default
// value for the given deployment configuration, from all
// ProductSections of vs (DSP0243 9.5.1). Used for resolving
// ovf:capacity="${key}" on Disk elements (DSP0243 9.1).
func (e Envelope) propertyDefaultsForConfig(vs *VirtualSystem, configName string) map[string]string {
	out := make(map[string]string)
	for _, product := range vs.Product {
		for _, pair := range product.PropertiesWithCategory() {
			p := pair.Property
			if p.Configuration != nil && *p.Configuration != configName {
				continue
			}
			var value string
			if p.Default != nil {
				value = *p.Default
			}
			for _, v := range p.Values {
				if v.Configuration != nil && *v.Configuration == configName {
					value = v.Value
					break
				}
			}
			out[p.Key] = strings.TrimSpace(value)
		}
	}
	return out
}

func (e Envelope) toVirtualDisk(
	item itemElement,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice,
	fileRefs map[string]File,
	propertyValues map[string]string) (types.BaseVirtualDevice, error) {

	var c types.BaseVirtualController
	if item.Parent != nil {
		if r, ok := resources[*item.Parent]; ok {
			if c1, ok := r.(types.BaseVirtualController); ok {
				c = c1
			}
		}
	}

	// The diskName variable is used to store the name of the disk from
	// disk from the OVF that are backed by a file.
	// This is used to set the disk name in the ConfigSpec to distinguish
	// between empty disks and disks that are backed by a file.
	diskName := ""

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

		if dd.FileRef != nil && *dd.FileRef != "" {
			if f, ok := fileRefs[*dd.FileRef]; ok {
				diskName = path.Base(f.Href)
			}
		}

		var allocUnitsSz string
		if dd.CapacityAllocationUnits != nil {
			allocUnitsSz = *dd.CapacityAllocationUnits
		}
		capacityInBytes = uint64(ParseCapacityAllocationUnits(allocUnitsSz))
		if capSz := dd.Capacity; capSz != "" {
			var capVal uint64
			if strings.HasPrefix(capSz, "${") && strings.HasSuffix(capSz, "}") {
				// DSP0243 9.1: capacity may reference a Property, e.g. ovf:capacity="${disk.size}"
				key := strings.TrimSpace(capSz[2 : len(capSz)-1])
				if key == "" {
					return nil, fmt.Errorf("disk=%s has empty property reference in capacity=%q", diskID, capSz)
				}
				val, ok := propertyValues[key]
				if !ok || val == "" {
					return nil, fmt.Errorf("disk=%s capacity references property %q which is not defined or has no value for this configuration", diskID, key)
				}
				var err error
				capVal, err = strconv.ParseUint(val, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("disk=%s capacity property %q value %q is not a valid integer: %w", diskID, key, val, err)
				}
				capacityInBytes *= capVal
			} else {
				var err error
				capVal, err = strconv.ParseUint(dd.Capacity, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("disk=%s has invalid capacity=%q",
						diskID, capSz)
				}
				capacityInBytes *= capVal
			}
		}

	default:
		return nil, fmt.Errorf("multiple HostResource elements")
	}

	if capacityInBytes > math.MaxInt64 {
		return nil, fmt.Errorf(
			"capacityInBytes=%d exceeds math.MaxInt64", capacityInBytes)
	}

	d := devices.CreateDisk(c, types.ManagedObjectReference{}, diskName)

	d.VirtualDevice.DeviceInfo = &types.Description{
		Label: item.ElementName,
	}

	d.CapacityInBytes = int64(capacityInBytes)

	return d, nil
}

func (e Envelope) toCDOrDVDDrive(
	item itemElement,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	if item.Parent == nil {
		return nil, fmt.Errorf("missing Parent")
	}

	r, ok := resources[*item.Parent]
	if !ok {
		return nil, nil // Parent is unsupported
	}

	c, ok := r.(types.BaseVirtualController)
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

var validSCSIControllerTypes = map[string]struct{}{
	"buslogic":     {},
	"lsilogic":     {},
	"lsilogic-sas": {},
	"pvscsi":       {},
	"scsi":         {},
	"virtualscsi":  {},
}

// resolveSCSIControllerType parses resourceSubType by splitting on
// comma or space, trims whitespace from each element, and returns the
// first element that matches a key in validSCSIControllerTypes (match
// is case-insensitive). If no match, returns the original string.
func resolveSCSIControllerType(resourceSubType string) string {
	trimmed := strings.TrimSpace(resourceSubType)
	// Split on comma or space (one or more of either)
	tokens := strings.Fields(strings.ReplaceAll(trimmed, ",", " "))
	for _, tok := range tokens {
		t := strings.TrimSpace(tok)
		lower := strings.ToLower(t)
		if _, ok := validSCSIControllerTypes[lower]; ok {
			return lower
		}
	}
	return resourceSubType
}

func (e Envelope) toSCSIController(
	item itemElement,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	controllerType := resolveSCSIControllerType(item.resourceSubType)

	d, err := devices.CreateSCSIController(controllerType)
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
	item itemElement,
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
	item itemElement,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	switch item.resourceSubType {
	case ResourceSubTypeSATAAHCI, ResourceSubTypeSATAAHCIAlter:
		return e.toSATAController(item, devices, resources)
	case ResourceSubTypeNVMEController:
		return e.toNVMEController(item, devices, resources)
	}
	return nil, errUnsupportedResourceSubtype
}

func (e Envelope) toSATAController(
	item itemElement,
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
	item itemElement,
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
	item itemElement,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	var d types.BaseVirtualDevice

	vc := types.VirtualController{
		VirtualDevice: types.VirtualDevice{
			Key: devices.NewKey(),
		},
	}

	switch item.resourceSubType {
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
		return nil, errUnsupportedResourceSubtype
	}

	if err := e.setBusNumber(item, d); err != nil {
		return nil, err
	}

	resources[item.InstanceID] = d
	return d, nil
}

func (e Envelope) toNetworkInterface(
	item itemElement,
	devices object.VirtualDeviceList,
	_ map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	d, err := devices.CreateEthernetCard(item.resourceSubType, nil)
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
	item itemElement,
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
	item itemElement,
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
	item itemElement,
	devices object.VirtualDeviceList,
	resources map[string]types.BaseVirtualDevice) (types.BaseVirtualDevice, error) {

	switch item.resourceSubType {
	case ResourceSubTypeVMCI:
		return e.toVMCI(item, devices, resources)
	}
	return nil, errUnsupportedResourceSubtype
}

func (e Envelope) toVMCI(
	item itemElement,
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

func setConnectable(dst types.BaseVirtualDevice, src itemElement) {

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

// vApp PropertyInfo base type constants (values repeated in ovfToVAppBaseType).
const vAppBaseTypeInt = "int"

// OVF property type to vApp PropertyInfo.Type (base type) per DSP0243 Table 6
// and vSphere API vim.vApp.PropertyInfo. Integer OVF types map to "int",
// real32/real64 to "real".
var ovfToVAppBaseType = map[string]string{
	"uint8":   vAppBaseTypeInt,
	"sint8":   vAppBaseTypeInt,
	"uint16":  vAppBaseTypeInt,
	"sint16":  vAppBaseTypeInt,
	"uint32":  vAppBaseTypeInt,
	"sint32":  vAppBaseTypeInt,
	"uint64":  vAppBaseTypeInt,
	"sint64":  vAppBaseTypeInt,
	"string":  "string",
	"boolean": "boolean",
	"real32":  "real",
	"real64":  "real",
	"int":     vAppBaseTypeInt, // common in OVF although not in DSP0243 Table 6
}

// Regexes for OVF qualifiers per DSP0243 section 9.5.1 Table 7 and vSphere QualifierMap.
var (
	ovfMinLenRx   = regexp.MustCompile(`MinLen\((\d+)\)`)
	ovfMaxLenRx   = regexp.MustCompile(`MaxLen\((\d+)\)`)
	ovfValueMapRx = regexp.MustCompile(`ValueMap\{(.*)\}`)
	ovfMinValueRx = regexp.MustCompile(`MinValue\(([-]?\d+(?:\.\d*)?(?:[eE][-+]?\d+)?)\)`)
	ovfMaxValueRx = regexp.MustCompile(`MaxValue\(([-]?\d+(?:\.\d*)?(?:[eE][-+]?\d+)?)\)`)
	// VMware qualifiers: Ip, Ip(), or Ip("network") / Ip(network) -> ip / ip:network
	ovfIpNoArgRx = regexp.MustCompile(`Ip\s*\(\s*\)`)
	ovfIpArgRx   = regexp.MustCompile(`Ip\s*\(\s*["']?([^"')]*)["']?\s*\)`) // arg may be quoted or unquoted
	ovfIpBareRx  = regexp.MustCompile(`\bIp\b`)                             // bare "Ip" (no parens)
)

// vSphere expression qualifier pattern (VmwOvf) -> vApp type "expression".
// With parens: VimIp(), Net("Management"). Without parens: bare "VimIp" etc.
var (
	ovfExpressionQualifierRx = regexp.MustCompile(
		`(AutoIp|VimIp|Net|Netmask|Gateway|DomainName|HostPrefix|Dns|Subnet|SearchPath|HttpProxy)\s*\(\s*["']?([^"')]*)["']?\s*\)`)
	ovfExpressionBareRx = regexp.MustCompile(
		`\b(AutoIp|VimIp|Net|Netmask|Gateway|DomainName|HostPrefix|Dns|Subnet|SearchPath|HttpProxy)\b`)
)

// ovfTypeToVAppBaseType returns the vApp PropertyInfo base type for an OVF
// property type.
func ovfTypeToVAppBaseType(ovfType string) string {
	key := strings.TrimSpace(strings.ToLower(ovfType))
	if t, ok := ovfToVAppBaseType[key]; ok {
		return t
	}
	return ovfType
}

// ovfQualifiers holds parsed OVF qualifiers (DSP0243 9.5.1 Table 7 and vSphere QualifierMap).
type ovfQualifiers struct {
	minLen int // MinLen(min) for string
	maxLen int // MaxLen(max) for string; -1 means not set
	// ValueMap{...} (parsed as strings for string and int)
	valueMap []string
	// MinValue(n), MaxValue(n) for int/real
	minValue, maxValue *float64
	// VMware Ip() or Ip("network") -> ip or ip:network
	ipArg *string // nil = not Ip, "" = Ip(), "x" = Ip("x")
	// VMware expression qualifier (VimIp, Net, etc.) -> type "expression"
	expressionName string // e.g. "VimIp"
	expressionArg  string // e.g. "" or "Management"
}

func parseOVFQualifiers(qualifiers string) (q ovfQualifiers) {
	q.maxLen = -1
	qualifiers = strings.TrimSpace(qualifiers)
	if qualifiers == "" {
		return q
	}
	if m := ovfMinLenRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		q.minLen, _ = strconv.Atoi(m[1])
	}
	if m := ovfMaxLenRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		q.maxLen, _ = strconv.Atoi(m[1])
	}
	if m := ovfValueMapRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		q.valueMap = parseValueMapContent(m[1])
	}
	if m := ovfMinValueRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		if v, err := strconv.ParseFloat(m[1], 64); err == nil {
			q.minValue = &v
		}
	}
	if m := ovfMaxValueRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		if v, err := strconv.ParseFloat(m[1], 64); err == nil {
			q.maxValue = &v
		}
	}
	if m := ovfIpArgRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		arg := strings.TrimSpace(m[1])
		q.ipArg = &arg
	} else if ovfIpNoArgRx.MatchString(qualifiers) {
		s := ""
		q.ipArg = &s
	} else if ovfIpBareRx.MatchString(qualifiers) {
		s := ""
		q.ipArg = &s
	}
	if m := ovfExpressionQualifierRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		q.expressionName = m[1]
		q.expressionArg = strings.TrimSpace(m[2])
	} else if m := ovfExpressionBareRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		q.expressionName = m[1]
		q.expressionArg = ""
	}
	return q
}

// expressionDefaultFromQualifiers returns the default value for an expression-typed
// property (e.g. "${vimIp:}" or "${net:Management}") per vSphere QualifierMap.
// Returns "" if qualifiers do not contain a vSphere expression qualifier.
func expressionDefaultFromQualifiers(qualifiers string) string {
	if qualifiers == "" {
		return ""
	}
	qualifiers = strings.TrimSpace(qualifiers)
	var name, arg string
	if m := ovfExpressionQualifierRx.FindStringSubmatch(qualifiers); len(m) >= 3 {
		name = m[1]
		arg = strings.TrimSpace(m[2])
	} else if m := ovfExpressionBareRx.FindStringSubmatch(qualifiers); len(m) > 0 {
		name = m[1]
		arg = ""
	} else {
		return ""
	}
	// Lowercase first letter to match vSphere (e.g. VimIp -> vimIp)
	lcName := name
	if len(name) > 0 {
		lcName = strings.ToLower(name[:1]) + name[1:]
	}
	if arg != "" {
		return "${" + lcName + ":" + arg + "}"
	}
	return "${" + lcName + ":}"
}

// parseValueMapContent parses the content inside ValueMap{...}:
// comma-separated values, each a quoted string or number (per CIM/DSP0004).
func parseValueMapContent(content string) []string {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil
	}
	var values []string
	i := 0
	for i < len(content) {
		for i < len(content) && (content[i] == ',' || content[i] == ' ') {
			i++
		}
		if i >= len(content) {
			break
		}
		if content[i] == '"' {
			// Quoted string: find closing " (handle escaped \" and \')
			i++
			start := i
			for i < len(content) {
				if content[i] == '\\' && i+1 < len(content) {
					i += 2
					continue
				}
				if content[i] == '"' {
					values = append(values, content[start:i])
					i++
					break
				}
				i++
			}
			continue
		}
		// Number or unquoted token
		start := i
		for i < len(content) && content[i] != ',' {
			i++
		}
		values = append(values, strings.TrimSpace(content[start:i]))
	}
	return values
}

// toVAppPropertyType returns the vApp PropertyInfo.Type string for an OVF
// Property, applying the OVF-to-vApp type mapping and embedding qualifiers
// per DSP0243 9.5.1, vim.vApp.PropertyInfo, and vSphere QualifierMap.
func toVAppPropertyType(p Property) string {
	baseType := ovfTypeToVAppBaseType(p.Type)
	if p.Password != nil && *p.Password && baseType == "string" {
		baseType = "password"
	}
	qualStr := ""
	if p.Qualifiers != nil {
		qualStr = strings.TrimSpace(*p.Qualifiers)
	}
	if qualStr == "" {
		return baseType
	}
	q := parseOVFQualifiers(qualStr)

	// VMware expression qualifiers (VimIp, Net, etc.) -> "expression"
	if q.expressionName != "" {
		return "expression"
	}
	// VMware Ip() or Ip("network") -> "ip" or "ip:network" (string type only)
	if q.ipArg != nil && (baseType == "string" || baseType == "password") {
		if *q.ipArg == "" {
			return "ip"
		}
		return "ip:" + *q.ipArg
	}

	switch baseType {
	case "string", "password":
		if len(q.valueMap) > 0 {
			// string["choice1", "choice2", ...] — escape " and \ in choices
			parts := make([]string, 0, len(q.valueMap))
			for _, v := range q.valueMap {
				esc := strings.ReplaceAll(strings.ReplaceAll(v, `\`, `\\`), `"`, `\"`)
				parts = append(parts, `"`+esc+`"`)
			}
			return baseType + "[" + strings.Join(parts, ", ") + "]"
		}
		// When both MinLen and MaxLen are set (including MinLen(0)), use
		// (min..max) so the parser's strWithMinMaxLenRx applies correctly.
		// (min..) and (..max) are parsed by strWithMinLenRx and strWithMaxLenRx.
		if q.maxLen >= 0 && q.minLen >= 0 {
			return fmt.Sprintf("%s(%d..%d)", baseType, q.minLen, q.maxLen)
		}
		if q.maxLen >= 0 {
			return fmt.Sprintf("%s(..%d)", baseType, q.maxLen)
		}
		if q.minLen > 0 {
			return fmt.Sprintf("%s(%d..)", baseType, q.minLen)
		}
	case "int":
		if len(q.valueMap) > 0 {
			minVal, maxVal, ok := valueMapIntRange(q.valueMap)
			if ok {
				return fmt.Sprintf("int(%d..%d)", minVal, maxVal)
			}
		}
		if q.minValue != nil || q.maxValue != nil {
			minV, maxV := int64(-2147483648), int64(2147483647)
			if q.minValue != nil {
				minV = int64(*q.minValue)
			}
			if q.maxValue != nil {
				maxV = int64(*q.maxValue)
			}
			return fmt.Sprintf("int(%d..%d)", minV, maxV)
		}
	case "real":
		if q.minValue != nil || q.maxValue != nil {
			minV, maxV := -1e308, 1e308
			if q.minValue != nil {
				minV = *q.minValue
			}
			if q.maxValue != nil {
				maxV = *q.maxValue
			}
			return fmt.Sprintf("real(%v..%v)", minV, maxV)
		}
	}
	return baseType
}

// valueMapIntRange parses valueMap as integers and returns min, max and true;
// otherwise 0, 0, false.
func valueMapIntRange(valueMap []string) (minVal, maxVal int64, ok bool) {
	if len(valueMap) == 0 {
		return 0, 0, false
	}
	first, err := strconv.ParseInt(strings.TrimSpace(valueMap[0]), 10, 64)
	if err != nil {
		return 0, 0, false
	}
	minVal, maxVal = first, first
	for i := 1; i < len(valueMap); i++ {
		v, err := strconv.ParseInt(strings.TrimSpace(valueMap[i]), 10, 64)
		if err != nil {
			return 0, 0, false
		}
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return minVal, maxVal, true
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

		for _, pair := range product.PropertiesWithCategory() {
			p := pair.Property
			if p.Configuration != nil && *p.Configuration != configName {
				// Skip properties that are not part of the provided
				// configuration.
				continue
			}
			// DSP0243 9.5.1: ovf:key shall not contain period or colon.
			// if strings.ContainsAny(p.Key, ".:") {
			// 	return fmt.Errorf(
			// 		"property key %q contains invalid character "+
			// 			"(period or colon) per DSP0243 9.5.1", p.Key)
			// }

			// Get the default values for the current configuration from the
			// list of default values that are per-config.
			var value string
			if p.Default != nil {
				value = *p.Default
			}
			for _, v := range p.Values {
				if v.Configuration == nil || *v.Configuration == configName {
					value = v.Value
					break
				}
			}

			if p.UserConfigurable == nil {
				p.UserConfigurable = types.NewBool(false)
			}
			vAppType := toVAppPropertyType(p)
			// Parse the value using the vApp config parser with the computed
			// type so that qualifier-derived constraints (e.g. string(1..65535))
			// are applied.
			value = strings.TrimSpace(value)
			parsedValue := value
			if value != "" {
				pForParse := Property{Key: p.Key, Type: vAppType}
				var err error
				parsedValue, err = parseVAppConfigValue(pForParse, value)
				if err != nil {
					return err
				}
			} else if vAppType == "expression" && p.Qualifiers != nil {
				// vSphere QualifierMap: expression qualifiers get default like "${vimIp:}".
				if d := expressionDefaultFromQualifiers(*p.Qualifiers); d != "" {
					parsedValue = d
				}
			}
			// Use per-property category from DSP0243 9.5.1 grouping.
			category := pair.Category
			np := types.VAppPropertySpec{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: types.ArrayUpdateOperationAdd,
				},
				Info: &types.VAppPropertyInfo{
					Key:              int32(index),
					ClassId:          deref(product.Class),
					InstanceId:       deref(product.Instance),
					Id:               p.Key,
					Category:         category,
					Label:            deref(p.Label),
					Type:             vAppType,
					UserConfigurable: p.UserConfigurable,
					DefaultValue:     parsedValue,
					Value:            "",
					Description:      deref(p.Description),
				},
			}

			vapp.Property = append(vapp.Property, np)
			index++
		}
	}

	dst.VAppConfig = vapp
	return nil
}
