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

// SelectByType returns a new list with devices that are equal to or extend the given types.
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
		default:
			return num < 2
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
			DeviceName:    fmt.Sprintf("cdrom-%d-%d", device.ControllerKey, device.UnitNumber),
			UseAutoDetect: false,
		},
	}
}

// TypeName returns the vmodl type name of the device
func (l VirtualDeviceList) TypeName(device types.BaseVirtualDevice) string {
	return reflect.TypeOf(device).Elem().Name()
}

var deviceNameRegexp = regexp.MustCompile(`(?:Virtual)?(?:Machine)?(\w+?)(?:Card|Device|Controller)?$`)

// Name returns a stable, human-readable name for the given device
func (l VirtualDeviceList) Name(device types.BaseVirtualDevice) string {
	typeName := l.TypeName(device)
	d := device.GetVirtualDevice()

	switch device.(type) {
	case types.BaseVirtualEthernetCard:
		return fmt.Sprintf("ethernet-%d", d.UnitNumber-7)
	case *types.ParaVirtualSCSIController:
		return fmt.Sprintf("pvscsi-%d", d.Key)
	case *types.VirtualDisk:
		return fmt.Sprintf("disk-%d-%d", d.ControllerKey, d.UnitNumber)
	default:
		name := "device"
		m := deviceNameRegexp.FindStringSubmatch(typeName)
		if len(m) == 2 {
			name = strings.ToLower(m[1])
		}
		return fmt.Sprintf("%s-%d", name, d.Key)
	}
}
