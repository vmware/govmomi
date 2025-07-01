// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type NetworkFlag struct {
	common

	*DatacenterFlag

	name    string
	net     object.NetworkReference
	adapter string
	address string
	isset   bool
	proto   string
}

var networkFlagKey = flagKey("network")

func NewNetworkFlag(ctx context.Context) (*NetworkFlag, context.Context) {
	if v := ctx.Value(networkFlagKey); v != nil {
		return v.(*NetworkFlag), ctx
	}

	v := &NetworkFlag{}
	v.DatacenterFlag, ctx = NewDatacenterFlag(ctx)
	ctx = context.WithValue(ctx, networkFlagKey, v)
	return v, ctx
}

func (flag *NetworkFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		flag.DatacenterFlag.Register(ctx, f)

		env := "GOVC_NETWORK"
		value := os.Getenv(env)
		flag.name = value
		usage := fmt.Sprintf("Network [%s]", env)
		f.Var(flag, "net", usage)
		f.StringVar(&flag.adapter, "net.adapter", "e1000", "Network adapter type")
		f.StringVar(&flag.address, "net.address", "", "Network hardware address")
		f.StringVar(&flag.proto, "net.protocol", "", fmt.Sprintf("Network device protocol. Applicable to vmxnet3vrdma. Default to '%s'", string(types.VirtualVmxnet3VrdmaOptionDeviceProtocolsRocev2)))
	})
}

func (flag *NetworkFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if err := flag.DatacenterFlag.Process(ctx); err != nil {
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

	backing, err := net.EthernetCardBackingInfo(context.TODO())
	if err != nil {
		return nil, err
	}

	device, err := object.EthernetCardTypes().CreateEthernetCard(flag.adapter, backing)
	if err != nil {
		return nil, err
	}

	if a, ok := device.(*types.VirtualVmxnet3Vrdma); ok {
		if flag.proto != "" {
			if flag.proto != string(types.VirtualVmxnet3VrdmaOptionDeviceProtocolsRocev2) &&
				flag.proto != string(types.VirtualVmxnet3VrdmaOptionDeviceProtocolsRocev1) {
				return nil, fmt.Errorf("invalid device protocol '%s'", flag.proto)
			}
			a.DeviceProtocol = flag.proto
		}
	} else if flag.proto != "" {
		return nil, fmt.Errorf("device protocol is only supported for vmxnet3vrdma at the moment")
	}

	if flag.address == "-" {
		card := device.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
		card.AddressType = string(types.VirtualEthernetCardMacTypeGenerated)
		card.MacAddress = ""
	} else if flag.address != "" {
		card := device.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
		card.AddressType = string(types.VirtualEthernetCardMacTypeManual)
		card.MacAddress = flag.address
	}

	return device, nil
}

// Change applies update backing and hardware address changes to the given network device.
func (flag *NetworkFlag) Change(device types.BaseVirtualDevice, update types.BaseVirtualDevice) {
	current := device.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()
	changed := update.(types.BaseVirtualEthernetCard).GetVirtualEthernetCard()

	current.Backing = changed.Backing

	if changed.MacAddress != "" {
		current.MacAddress = changed.MacAddress
	}

	if changed.AddressType != "" {
		current.AddressType = changed.AddressType
	}
}
