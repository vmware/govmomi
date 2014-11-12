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

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

// Type values for use in BootOrder
const (
	DeviceTypeCdrom    = "cdrom"
	DeviceTypeDisk     = "disk"
	DeviceTypeEthernet = "ethernet"
	DeviceTypeFloppy   = "floppy"
)

// VirtualDeviceList provides helper methods for working with a list of virtual devices.
type VirtualDeviceList []types.BaseVirtualDevice

// Select returns a new list containing all elements of the list for which the given func returns true.
func (l VirtualDeviceList) Select(f func(device types.BaseVirtualDevice) bool) VirtualDeviceList {
	var found VirtualDeviceList

	for _, device := range l {
		if f(device) {
			found = append(found, device)
		}
	}

	return found
}

// SelectByType returns a new list with devices that are equal to or extend the given type.
func (l VirtualDeviceList) SelectByType(deviceType types.BaseVirtualDevice) VirtualDeviceList {
	dtype := reflect.TypeOf(deviceType)
	dname := dtype.Elem().Name()

	return l.Select(func(device types.BaseVirtualDevice) bool {
		t := reflect.TypeOf(device)

		if t == dtype {
			return true
		}

		_, ok := t.Elem().FieldByName(dname)

		return ok
	})
}

// SelectByBackingInfo returns a new list with devices matching the given backing info.
func (l VirtualDeviceList) SelectByBackingInfo(backing types.BaseVirtualDeviceBackingInfo) VirtualDeviceList {
	t := reflect.TypeOf(backing)

	return l.Select(func(device types.BaseVirtualDevice) bool {
		db := device.GetVirtualDevice().Backing
		if db == nil {
			return false
		}

		if reflect.TypeOf(db) != t {
			return false
		}

		switch a := db.(type) {
		case *types.VirtualEthernetCardNetworkBackingInfo:
			b := backing.(*types.VirtualEthernetCardNetworkBackingInfo)
			return a.DeviceName == b.DeviceName
		case *types.VirtualEthernetCardDistributedVirtualPortBackingInfo:
			b := backing.(*types.VirtualEthernetCardDistributedVirtualPortBackingInfo)
			return a.Port.SwitchUuid == b.Port.SwitchUuid
		case *types.VirtualDiskFlatVer2BackingInfo:
			b := backing.(*types.VirtualDiskFlatVer2BackingInfo)
			if a.Parent != nil && b.Parent != nil {
				return a.Parent.FileName == b.Parent.FileName
			}
			return a.FileName == b.FileName
		case types.BaseVirtualDeviceFileBackingInfo:
			b := backing.(types.BaseVirtualDeviceFileBackingInfo)
			return a.GetVirtualDeviceFileBackingInfo().FileName == b.GetVirtualDeviceFileBackingInfo().FileName
		default:
			return false
		}
	})
}

// Find returns the device matching the given name.
func (l VirtualDeviceList) Find(name string) types.BaseVirtualDevice {
	for _, device := range l {
		if l.Name(device) == name {
			return device
		}
	}
	return nil
}

// FindByKey returns the device matching the given key.
func (l VirtualDeviceList) FindByKey(key int) types.BaseVirtualDevice {
	for _, device := range l {
		if device.GetVirtualDevice().Key == key {
			return device
		}
	}
	return nil
}

// FindIDEController will find the named IDE controller if given, otherwise will pick an available controller.
// An error is returned if the named controller is not found or not an IDE controller.  Or, if name is not
// given and no available controller can be found.
func (l VirtualDeviceList) FindIDEController(name string) (*types.VirtualIDEController, error) {
	if name != "" {
		d := l.Find(name)
		if d == nil {
			return nil, fmt.Errorf("device '%s' not found", name)
		}
		if c, ok := d.(*types.VirtualIDEController); ok {
			return c, nil
		}
		return nil, fmt.Errorf("%s is not an IDE controller", name)
	}

	c := l.PickController((*types.VirtualIDEController)(nil))
	if c == nil {
		return nil, errors.New("no available IDE controller")
	}

	return c.(*types.VirtualIDEController), nil
}

// PickController returns a controller of the given type(s).
// If no controllers are found or have no available slots, then nil is returned.
func (l VirtualDeviceList) PickController(kind types.BaseVirtualController) types.BaseVirtualController {
	l = l.SelectByType(kind.(types.BaseVirtualDevice)).Select(func(device types.BaseVirtualDevice) bool {
		num := len(device.(types.BaseVirtualController).GetVirtualController().Device)

		switch device.(type) {
		case types.BaseVirtualSCSIController:
			return num < 15
		case *types.VirtualIDEController:
			return num < 2
		default:
			return true
		}
	})

	if len(l) == 0 {
		return nil
	}

	return l[0].(types.BaseVirtualController)
}

// newUnitNumber returns the unit number to use for attaching a new device to the given controller.
func (l VirtualDeviceList) newUnitNumber(c types.BaseVirtualController) int {
	key := c.GetVirtualController().Key
	max := -1

	for _, device := range l {
		d := device.GetVirtualDevice()

		if d.ControllerKey == key {
			if d.UnitNumber > max {
				max = d.UnitNumber
			}
		}
	}

	return max + 1
}

// AssignController assigns a device to a controller.
func (l VirtualDeviceList) AssignController(device types.BaseVirtualDevice, c types.BaseVirtualController) {
	d := device.GetVirtualDevice()
	d.ControllerKey = c.GetVirtualController().Key
	d.UnitNumber = l.newUnitNumber(c)
	d.Key = -1
}

func (l VirtualDeviceList) connectivity(device types.BaseVirtualDevice, v bool) error {
	c := device.GetVirtualDevice().Connectable
	if c == nil {
		return fmt.Errorf("%s is not connectable", l.Name(device))
	}

	c.Connected = v
	c.StartConnected = v

	return nil
}

// Connect changes the device to connected, returns an error if the device is not connectable.
func (l VirtualDeviceList) Connect(device types.BaseVirtualDevice) error {
	return l.connectivity(device, true)
}

// Disconnect changes the device to disconnected, returns an error if the device is not connectable.
func (l VirtualDeviceList) Disconnect(device types.BaseVirtualDevice) error {
	return l.connectivity(device, false)
}

// FindCdrom finds a cdrom device with the given name, defaulting to the first cdrom device if any.
func (l VirtualDeviceList) FindCdrom(name string) (*types.VirtualCdrom, error) {
	if name != "" {
		d := l.Find(name)
		if d == nil {
			return nil, fmt.Errorf("device '%s' not found", name)
		}
		if c, ok := d.(*types.VirtualCdrom); ok {
			return c, nil
		}
		return nil, fmt.Errorf("%s is not a cdrom device", name)
	}

	c := l.SelectByType((*types.VirtualCdrom)(nil))
	if len(c) == 0 {
		return nil, errors.New("no cdrom device found")
	}

	return c[0].(*types.VirtualCdrom), nil
}

// CreateCdrom creates a new VirtualCdrom device which can be added to a VM.
func (l VirtualDeviceList) CreateCdrom(c *types.VirtualIDEController) (*types.VirtualCdrom, error) {
	device := &types.VirtualCdrom{}

	l.AssignController(device, c)

	l.setDefaultCdromBacking(device)

	device.Connectable = &types.VirtualDeviceConnectInfo{
		AllowGuestControl: true,
		Connected:         true,
		StartConnected:    true,
	}

	return device, nil
}

// InsertIso changes the cdrom device backing to use the given iso file.
func (l VirtualDeviceList) InsertIso(device *types.VirtualCdrom, iso string) *types.VirtualCdrom {
	device.Backing = &types.VirtualCdromIsoBackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: iso,
		},
	}

	return device
}

// EjectIso removes the iso file based backing and replaces with the default cdrom backing.
func (l VirtualDeviceList) EjectIso(device *types.VirtualCdrom) *types.VirtualCdrom {
	l.setDefaultCdromBacking(device)
	return device
}

func (l VirtualDeviceList) setDefaultCdromBacking(device *types.VirtualCdrom) {
	device.Backing = &types.VirtualCdromAtapiBackingInfo{
		VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
			DeviceName:    fmt.Sprintf("%s-%d-%d", DeviceTypeCdrom, device.ControllerKey, device.UnitNumber),
			UseAutoDetect: false,
		},
	}
}

// FindFloppy finds a floppy device with the given name, defaulting to the first floppy device if any.
func (l VirtualDeviceList) FindFloppy(name string) (*types.VirtualFloppy, error) {
	if name != "" {
		d := l.Find(name)
		if d == nil {
			return nil, fmt.Errorf("device '%s' not found", name)
		}
		if c, ok := d.(*types.VirtualFloppy); ok {
			return c, nil
		}
		return nil, fmt.Errorf("%s is not a floppy device", name)
	}

	c := l.SelectByType((*types.VirtualFloppy)(nil))
	if len(c) == 0 {
		return nil, errors.New("no floppy device found")
	}

	return c[0].(*types.VirtualFloppy), nil
}

// CreateFloppy creates a new VirtualFloppy device which can be added to a VM.
func (l VirtualDeviceList) CreateFloppy() (*types.VirtualFloppy, error) {
	device := &types.VirtualFloppy{}

	c := l.PickController((*types.VirtualSIOController)(nil))
	if c == nil {
		return nil, errors.New("no available SIO controller")
	}

	l.AssignController(device, c)

	l.setDefaultFloppyBacking(device)

	device.Connectable = &types.VirtualDeviceConnectInfo{
		AllowGuestControl: true,
		Connected:         true,
		StartConnected:    true,
	}

	return device, nil
}

// InsertImg changes the floppy device backing to use the given img file.
func (l VirtualDeviceList) InsertImg(device *types.VirtualFloppy, img string) *types.VirtualFloppy {
	device.Backing = &types.VirtualFloppyImageBackingInfo{
		VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
			FileName: img,
		},
	}

	return device
}

// EjectImg removes the img file based backing and replaces with the default floppy backing.
func (l VirtualDeviceList) EjectImg(device *types.VirtualFloppy) *types.VirtualFloppy {
	l.setDefaultFloppyBacking(device)
	return device
}

func (l VirtualDeviceList) setDefaultFloppyBacking(device *types.VirtualFloppy) {
	device.Backing = &types.VirtualFloppyDeviceBackingInfo{
		VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
			DeviceName:    fmt.Sprintf("%s-%d", DeviceTypeFloppy, device.UnitNumber),
			UseAutoDetect: false,
		},
	}
}

// FindSerialPort finds a serial port device with the given name, defaulting to the first serial port device if any.
func (l VirtualDeviceList) FindSerialPort(name string) (*types.VirtualSerialPort, error) {
	if name != "" {
		d := l.Find(name)
		if d == nil {
			return nil, fmt.Errorf("device '%s' not found", name)
		}
		if c, ok := d.(*types.VirtualSerialPort); ok {
			return c, nil
		}
		return nil, fmt.Errorf("%s is not a serial port device", name)
	}

	c := l.SelectByType((*types.VirtualSerialPort)(nil))
	if len(c) == 0 {
		return nil, errors.New("no serial port device found")
	}

	return c[0].(*types.VirtualSerialPort), nil
}

// CreateSerialPort creates a new VirtualSerialPort device which can be added to a VM.
func (l VirtualDeviceList) CreateSerialPort() (*types.VirtualSerialPort, error) {
	device := &types.VirtualSerialPort{
		YieldOnPoll: true,
	}

	c := l.PickController((*types.VirtualSIOController)(nil))
	if c == nil {
		return nil, errors.New("no available SIO controller")
	}

	l.AssignController(device, c)

	l.setDefaultSerialPortBacking(device)

	return device, nil
}

// ConnectSerialPort connects a serial port to a server or client uri.
func (l VirtualDeviceList) ConnectSerialPort(device *types.VirtualSerialPort, uri string, client bool) *types.VirtualSerialPort {
	direction := types.VirtualDeviceURIBackingOptionDirectionServer
	if client {
		direction = types.VirtualDeviceURIBackingOptionDirectionClient
	}

	device.Backing = &types.VirtualSerialPortURIBackingInfo{
		VirtualDeviceURIBackingInfo: types.VirtualDeviceURIBackingInfo{
			Direction:  string(direction),
			ServiceURI: uri,
		},
	}

	return device
}

// DisconnectSerialPort disconnects the serial port backing.
func (l VirtualDeviceList) DisconnectSerialPort(device *types.VirtualSerialPort) *types.VirtualSerialPort {
	l.setDefaultSerialPortBacking(device)
	return device
}

func (l VirtualDeviceList) setDefaultSerialPortBacking(device *types.VirtualSerialPort) {
	device.Backing = &types.VirtualSerialPortURIBackingInfo{
		VirtualDeviceURIBackingInfo: types.VirtualDeviceURIBackingInfo{
			Direction:  "client",
			ServiceURI: "localhost:0",
		},
	}
}

// PrimaryMacAddress returns the MacAddress field of the primary VirtualEthernetCard
func (l VirtualDeviceList) PrimaryMacAddress() string {
	eth0 := l.Find("ethernet-0")

	if eth0 == nil {
		return ""
	}

	return eth0.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard().MacAddress
}

// convert a BaseVirtualDevice to a BaseVirtualMachineBootOptionsBootableDevice
var bootableDevices = map[string]func(device types.BaseVirtualDevice) types.BaseVirtualMachineBootOptionsBootableDevice{
	DeviceTypeCdrom: func(types.BaseVirtualDevice) types.BaseVirtualMachineBootOptionsBootableDevice {
		return &types.VirtualMachineBootOptionsBootableCdromDevice{}
	},
	DeviceTypeDisk: func(d types.BaseVirtualDevice) types.BaseVirtualMachineBootOptionsBootableDevice {
		return &types.VirtualMachineBootOptionsBootableDiskDevice{
			DeviceKey: d.GetVirtualDevice().Key,
		}
	},
	DeviceTypeEthernet: func(d types.BaseVirtualDevice) types.BaseVirtualMachineBootOptionsBootableDevice {
		return &types.VirtualMachineBootOptionsBootableEthernetDevice{
			DeviceKey: d.GetVirtualDevice().Key,
		}
	},
	DeviceTypeFloppy: func(types.BaseVirtualDevice) types.BaseVirtualMachineBootOptionsBootableDevice {
		return &types.VirtualMachineBootOptionsBootableFloppyDevice{}
	},
}

// BootOrder returns a list of devices which can be used to set boot order via VirtualMachine.SetBootOptions.
// The order can any of "ethernet", "cdrom", "floppy" or "disk" or by specific device name.
func (l VirtualDeviceList) BootOrder(order []string) []types.BaseVirtualMachineBootOptionsBootableDevice {
	var devices []types.BaseVirtualMachineBootOptionsBootableDevice

	for _, name := range order {
		if kind, ok := bootableDevices[name]; ok {
			for _, device := range l {
				if l.Type(device) == name {
					devices = append(devices, kind(device))
				}

			}
			continue
		}

		if d := l.Find(name); d != nil {
			if kind, ok := bootableDevices[l.Type(d)]; ok {
				devices = append(devices, kind(d))
			}
		}
	}

	return devices
}

// SelectBootOrder returns an ordered list of devices matching the given bootable device order
func (l VirtualDeviceList) SelectBootOrder(order []types.BaseVirtualMachineBootOptionsBootableDevice) VirtualDeviceList {
	var devices VirtualDeviceList

	for _, bd := range order {
		for _, device := range l {
			if kind, ok := bootableDevices[l.Type(device)]; ok {
				if reflect.DeepEqual(kind(device), bd) {
					devices = append(devices, device)
				}
			}
		}
	}

	return devices
}

// TypeName returns the vmodl type name of the device
func (l VirtualDeviceList) TypeName(device types.BaseVirtualDevice) string {
	return reflect.TypeOf(device).Elem().Name()
}

var deviceNameRegexp = regexp.MustCompile(`(?:Virtual)?(?:Machine)?(\w+?)(?:Card|Device|Controller)?$`)

// Type returns a human-readable name for the given device
func (l VirtualDeviceList) Type(device types.BaseVirtualDevice) string {
	switch device.(type) {
	case types.BaseVirtualEthernetCard:
		return DeviceTypeEthernet
	case *types.ParaVirtualSCSIController:
		return "pvscsi"
	default:
		name := "device"
		typeName := l.TypeName(device)
		m := deviceNameRegexp.FindStringSubmatch(typeName)
		if len(m) == 2 {
			name = strings.ToLower(m[1])
		}
		return name
	}
}

// Name returns a stable, human-readable name for the given device
func (l VirtualDeviceList) Name(device types.BaseVirtualDevice) string {
	var key string
	d := device.GetVirtualDevice()
	dtype := l.Type(device)

	switch dtype {
	case DeviceTypeEthernet:
		key = fmt.Sprintf("%d", d.UnitNumber-7)
	case DeviceTypeDisk:
		key = fmt.Sprintf("%d-%d", d.ControllerKey, d.UnitNumber)
	default:
		key = fmt.Sprintf("%d", d.Key)
	}

	return fmt.Sprintf("%s-%s", dtype, key)
}
