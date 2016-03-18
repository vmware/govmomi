/*
Copyright (c) 2014-2016 VMware, Inc. All Rights Reserved.

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
	"os"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type NetworkFlag struct {
	common

	*DatacenterFlag
	*ClientFlag

	name    string
	net     object.NetworkReference
	adapter string
	address string
	isset   bool

	dvsPath  string
	dvpgPath string
}

var networkFlagKey = flagKey("network")

func NewNetworkFlag(ctx context.Context) (*NetworkFlag, context.Context) {
	if v := ctx.Value(networkFlagKey); v != nil {
		return v.(*NetworkFlag), ctx
	}

	v := &NetworkFlag{}
	v.DatacenterFlag, ctx = NewDatacenterFlag(ctx)

	v.ClientFlag, ctx = NewClientFlag(ctx)

	ctx = context.WithValue(ctx, networkFlagKey, v)
	return v, ctx
}

func (flag *NetworkFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		flag.DatacenterFlag.Register(ctx, f)

		flag.ClientFlag.Register(ctx, f)

		env := "GOVC_NETWORK"
		value := os.Getenv(env)
		flag.name = value
		usage := fmt.Sprintf("Network [%s]", env)
		f.Var(flag, "net", usage)
		f.StringVar(&flag.adapter, "net.adapter", "e1000", "Network adapter type")
		f.StringVar(&flag.address, "net.address", "", "Network hardware address")

		f.StringVar(&flag.dvsPath, "port.dvs", "", "Distributed Virtual Switch inventory path")
		f.StringVar(&flag.dvpgPath, "port.dvpg", "", "Distributed Virtual Portgroup inventory path")
	})
}

func (flag *NetworkFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if err := flag.DatacenterFlag.Process(ctx); err != nil {
			return err
		}
		if err := flag.ClientFlag.Process(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (flag *NetworkFlag) String() string {
	return flag.name
}

func (flag *NetworkFlag) Set(name string) error {
	flag.name = name
	flag.isset = true
	return nil
}

func (flag *NetworkFlag) IsSet() bool {
	return flag.isset
}

func (flag *NetworkFlag) Network() (object.NetworkReference, error) {
	if flag.net != nil {
		return flag.net, nil
	}

	finder, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	if flag.net, err = finder.NetworkOrDefault(context.TODO(), flag.name); err != nil {
		return nil, err
	}

	return flag.net, nil
}

func (flag *NetworkFlag) Device() (types.BaseVirtualDevice, error) {
	net, err := flag.Network()
	if err != nil {
		return nil, err
	}

	//	var backing types.VirtualDeviceBackingInfo
	var device types.BaseVirtualDevice

	if len(flag.dvsPath) > 0 {

		client, err := flag.Client()
		if err != nil {
			return nil, err
		}

		// Retrieve DVS object by inventory path
		dvsInv, err := object.NewSearchIndex(client).FindByInventoryPath(context.TODO(), flag.dvsPath)
		if err != nil {
			return nil, err
		}
		if dvsInv == nil {
			return nil, fmt.Errorf("DistributedVirtualSwitch was not found at %s", flag.dvsPath)
		}

		// Convert DVS object type
		dvs := (*dvsInv.(*object.VmwareDistributedVirtualSwitch))

		// Set base search criteria
		criteria := types.DistributedVirtualSwitchPortCriteria{
			Connected:  types.NewBool(false),
			Active:     types.NewBool(false),
			UplinkPort: types.NewBool(false),
			Inside:     types.NewBool(true),
		}

		// If a distributed virtual portgroup path is set, then add its portgroup key to the base criteria
		if len(flag.dvpgPath) > 0 {

			// Retrieve distributed virtual portgroup object by inventory path
			dvpgInv, err := object.NewSearchIndex(client).FindByInventoryPath(context.TODO(), flag.dvpgPath)
			if err != nil {
				return nil, err
			}
			if dvpgInv == nil {
				return nil, fmt.Errorf("DistributedVirtualPortgroup was not found at %s", flag.dvpgPath)
			}

			// Convert distributed virtual portgroup object type
			dvpg := (*dvpgInv.(*object.DistributedVirtualPortgroup))

			// Obtain portgroup key property
			var dvp mo.DistributedVirtualPortgroup
			if err := dvpg.Properties(context.TODO(), dvpg.Reference(), []string{"key"}, &dvp); err != nil {
				return nil, err
			}
			spew.Dump(dvp.Config.DistributedVirtualSwitch.Value)

			// Add portgroup key to port search criteria
			criteria.PortgroupKey = []string{dvp.Key}
		}

		// Prepare request
		req := types.FetchDVPorts{
			This:     dvs.Reference(),
			Criteria: &criteria,
		}

		// Fetch ports
		res, err := methods.FetchDVPorts(context.TODO(), client, &req)
		if err != nil {
			return nil, err
		}

		if len(res.Returnval) == 0 {
			return nil, fmt.Errorf("No available ports were found")
		}
		matchedDvPort := res.Returnval[0]

		backing := &types.VirtualEthernetCardDistributedVirtualPortBackingInfo{
			Port: types.DistributedVirtualSwitchPortConnection{
				PortgroupKey: matchedDvPort.PortgroupKey,
				SwitchUuid:   matchedDvPort.DvsUuid,
				PortKey:      matchedDvPort.Key,
			},
		}

		device, err = object.EthernetCardTypes().CreateEthernetCard(flag.adapter, backing)
		if err != nil {
			return nil, err
		}

	} else {

		backing, err := net.EthernetCardBackingInfo(context.TODO())
		if err != nil {
			return nil, err
		}

		device, err = object.EthernetCardTypes().CreateEthernetCard(flag.adapter, backing)
		if err != nil {
			return nil, err
		}

	}

	if flag.address != "" {
		card := device.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
		card.AddressType = string(types.VirtualEthernetCardMacTypeManual)
		card.MacAddress = flag.address
	}

	return device, nil
}
