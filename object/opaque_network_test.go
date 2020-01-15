/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package object_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestOpaqueNetworkEthernetCardBackingInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		summary := &types.OpaqueNetworkSummary{
			NetworkSummary: types.NetworkSummary{
				Network:    &types.ManagedObjectReference{Type: "OpaqueNetwork", Value: "network-o1196"},
				Name:       "vnet-gc-nsx-t-vnet-0",
				Accessible: true,
			},
			OpaqueNetworkId:   "6153f069-dafe-4b43-ae0a-5d2f3ab26911",
			OpaqueNetworkType: "nsx.LogicalSwitch",
		}

		finder := find.NewFinder(c)

		_, err := finder.Network(ctx, summary.Name)
		if err == nil {
			t.Fatal("expected error") // network does not exist yet
		}

		fobj, err := finder.Folder(ctx, "network")
		if err != nil {
			t.Fatal(err)
		}

		// Inject the nsx network into the inventory
		// TODO: the vCenter API does not have a way to add nsx/OpaqueNetwork types,
		// but vcsim should provide a simple built-in way
		folder := simulator.Map.Get(fobj.Reference()).(*simulator.Folder)
		nsx := new(mo.OpaqueNetwork)
		nsx.Self = *summary.Network
		nsx.Network.Name = summary.Name
		nsx.Summary = summary
		folder.ChildEntity = append(folder.ChildEntity, nsx.Self)
		simulator.Map.PutEntity(folder, nsx)

		// Now expect success
		net, err := finder.Network(ctx, summary.Name)
		if err != nil {
			t.Fatal(err)
		}
		_, err = net.EthernetCardBackingInfo(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})
}
