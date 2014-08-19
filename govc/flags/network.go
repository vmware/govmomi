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
	"flag"
	"fmt"
	"sync"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type NetworkFlag struct {
	*DatacenterFlag

	register sync.Once
	name     string
	net      *govmomi.Network
}

func (f *NetworkFlag) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.name, "net", "", "Network")
}

func (f *NetworkFlag) Process() error {
	return nil
}

func (f *NetworkFlag) Network() (*govmomi.Network, error) {
	if f.net != nil {
		return f.net, nil
	}

	dc, err := f.Datacenter()
	if err != nil {
		return nil, err
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	folders, err := dc.Folders(c)
	if err != nil {
		return nil, err
	}
	nf := folders.NetworkFolder

	if f.name != "" {
		ref, err := c.SearchIndex().FindChild(nf, f.name)
		if err == nil {
			return nil, err
		}
		f.net = ref.(*govmomi.Network)
		return f.net, nil
	}

	cs, err := nf.Children(c)
	if err != nil {
		return nil, err
	}
	// Default to using the only network if there is only one.
	if len(cs) != 1 {
		return nil, fmt.Errorf("more than one network, please specify one")
	}

	f.net = cs[0].(*govmomi.Network)
	return f.net, nil
}

func (f *NetworkFlag) Properties(p []string) (*mo.Network, error) {
	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	_, err = f.Network()
	if err != nil {
		return nil, err
	}

	var net mo.Network
	if err := c.Properties(f.net.Reference(), p, &net); err != nil {
		return nil, err
	}
	return &net, nil
}

func (f *NetworkFlag) Name() (string, error) {
	if f.name == "" {
		net, err := f.Properties([]string{"name"})
		if err != nil {
			return "", nil
		}
		f.name = net.Name
	}
	return f.name, nil
}

func (f *NetworkFlag) Device() (types.BaseVirtualDevice, error) {
	name, err := f.Name()
	if err != nil {
		return nil, err
	}

	// TODO: adapter type should be an option, default to e1000 for now.
	return &types.VirtualE1000{types.VirtualEthernetCard{
		VirtualDevice: types.VirtualDevice{
			Key: -1,
			DeviceInfo: &types.Description{
				Label:   "Network Adapter 1",
				Summary: name,
			},
			Backing: &types.VirtualEthernetCardNetworkBackingInfo{
				VirtualDeviceDeviceBackingInfo: types.VirtualDeviceDeviceBackingInfo{
					DeviceName: name,
				},
			},
		},
		AddressType: string(types.VirtualEthernetCardMacTypeGenerated),
	}}, nil
}
