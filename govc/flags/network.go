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

package flags

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type NetworkFlag struct {
	*DatacenterFlag

	name string
	net  *types.ManagedObjectReference
}

func NewNetworkFlag() *NetworkFlag {
	f := &NetworkFlag{}
	_ = f.Set(os.Getenv("GOVC_NETWORK"))
	return f
}

func (flag *NetworkFlag) Register(f *flag.FlagSet) {}

func (flag *NetworkFlag) Process() error {
	return nil
}

func (flag *NetworkFlag) String() string {
	return flag.name
}

func (flag *NetworkFlag) Set(name string) error {
	flag.name = name
	return nil
}

func (flag *NetworkFlag) findNetwork(path string) ([]*types.ManagedObjectReference, error) {
	relativeFunc := func() (govmomi.Reference, error) {
		dc, err := flag.Datacenter()
		if err != nil {
			return nil, err
		}

		c, err := flag.Client()
		if err != nil {
			return nil, err
		}

		f, err := dc.Folders(c)
		if err != nil {
			return nil, err
		}

		return f.NetworkFolder, nil
	}

	es, err := flag.List(path, false, relativeFunc)
	if err != nil {
		return nil, err
	}

	var ns []*types.ManagedObjectReference
	for _, e := range es {
		ref := e.Object.Reference()
		ns = append(ns, &ref)
	}

	return ns, nil
}

func (flag *NetworkFlag) findSpecifiedNetwork(path string) (*types.ManagedObjectReference, error) {
	networks, err := flag.findNetwork(path)
	if err != nil {
		return nil, err
	}

	if len(networks) == 0 {
		return nil, errors.New("no such network")
	}

	if len(networks) > 1 {
		return nil, errors.New("path resolves to multiple networks")
	}

	flag.net = networks[0]
	return flag.net, nil
}

func (flag *NetworkFlag) findDefaultNetwork() (*types.ManagedObjectReference, error) {
	networks, err := flag.findNetwork("*")
	if err != nil {
		return nil, err
	}

	if len(networks) == 0 {
		panic("no networks") // Should never happen
	}

	if len(networks) > 1 {
		return nil, errors.New("please specify a network")
	}

	flag.net = networks[0]
	return flag.net, nil
}

func (flag *NetworkFlag) Network() (*types.ManagedObjectReference, error) {
	if flag.net != nil {
		return flag.net, nil
	}

	if flag.name == "" {
		return flag.findDefaultNetwork()
	}

	return flag.findSpecifiedNetwork(flag.name)
}

func (flag *NetworkFlag) Device() (types.BaseVirtualDevice, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	net, err := flag.Network()
	if err != nil {
		return nil, err
	}

	var backing types.BaseVirtualDeviceBackingInfo
	name := flag.name

	switch net.Type {
	case "Network":
		backing = &types.VirtualEthernetCardNetworkBackingInfo{
			VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
				DeviceName: name,
			},
		}
	case "DistributedVirtualPortgroup":
		var dvp mo.DistributedVirtualPortgroup
		var dvs mo.VmwareDistributedVirtualSwitch // TODO: should be mo.BaseDistributedVirtualSwitch

		if err := c.Properties(*net, []string{"key", "config.distributedVirtualSwitch"}, &dvp); err != nil {
			return nil, err
		}

		if err := c.Properties(*dvp.Config.DistributedVirtualSwitch, []string{"uuid"}, &dvs); err != nil {
			return nil, err
		}

		backing = &types.VirtualEthernetCardDistributedVirtualPortBackingInfo{
			Port: types.DistributedVirtualSwitchPortConnection{
				PortgroupKey: dvp.Key,
				SwitchUuid:   dvs.Uuid,
			},
		}
	default:
		return nil, fmt.Errorf("%s not supported", net.Type)
	}

	// TODO: adapter type should be an option, default to e1000 for now.
	device := &types.VirtualE1000{
		VirtualEthernetCard: types.VirtualEthernetCard{
			VirtualDevice: types.VirtualDevice{
				Key: -1,
				DeviceInfo: &types.Description{
					Label:   "", // Label will be chosen for us
					Summary: name,
				},
				Backing: backing,
			},
			AddressType: string(types.VirtualEthernetCardMacTypeGenerated),
		},
	}

	return device, nil
}
