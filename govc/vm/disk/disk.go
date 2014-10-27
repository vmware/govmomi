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

package disk

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

var dsPathRegexp = regexp.MustCompile(`^\[.*\] (?:.*/)?([^/]+)\.vmdk$`)

func FindController(mvm mo.VirtualMachine) (int, error) {
	for _, dev := range mvm.Config.Hardware.Device {
		switch disk := dev.(type) {
		case *types.VirtualLsiLogicController:
			vdev := disk.GetVirtualDevice()
			return vdev.Key, nil
		}
	}

	return -1, nil
}

func FindDisk(dname string, mvm mo.VirtualMachine) (*types.VirtualDisk, error) {
	for _, dev := range mvm.Config.Hardware.Device {
		switch disk := dev.(type) {
		case *types.VirtualDisk:
			switch backing := disk.Backing.(type) {
			case *types.VirtualDiskFlatVer2BackingInfo:
				m := dsPathRegexp.FindStringSubmatch(backing.FileName)
				if len(m) >= 2 && m[1] == dname {
					return disk, nil
				}
			default:
				name := reflect.TypeOf(disk.Backing).String()
				panic(fmt.Sprintf("unsupported backing: %s", name))
			}
		}
	}

	return nil, nil
}
