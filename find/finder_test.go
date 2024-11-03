/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package find_test // TODO: move ../simulator/finder_test.go tests here

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func TestFindNetwork(t *testing.T) {
	model := simulator.VPX()
	model.PortgroupNSX = 3

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)
		pc := property.DefaultCollector(c)

		pgs, err := finder.NetworkList(ctx, "DC0_NSXPG*")
		if err != nil {
			t.Fatal(err)
		}

		// Rename DC0_NSXPG1 to DC0_NSXPG0 so we have a duplicate name
		task, err := pgs[1].(*object.DistributedVirtualPortgroup).Rename(ctx, pgs[0].(*object.DistributedVirtualPortgroup).Name())
		if err != nil {
			t.Fatal(err)
		}
		err = task.Wait(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// 2 PGs, same switch, same name
		pgs, err = finder.NetworkList(ctx, "DC0_NSXPG0")
		if err != nil {
			t.Fatal(err)
		}

		if len(pgs) != 2 {
			t.Fatalf("expected 2 NSX PGs, got %d", len(pgs))
		}

		for _, pg := range pgs {
			// Using InventoryPath fails as > 1 are found
			_, err = finder.Network(ctx, pg.GetInventoryPath())
			if _, ok := err.(*find.MultipleFoundError); !ok {
				t.Fatalf("expected MultipleFoundError, got %s", err)
			}

			// Find by MOID
			_, err = finder.Network(ctx, pg.Reference().String())
			if err != nil {
				t.Errorf("find by moid: %s", err)
			}

			// Find by Switch UUID
			var props mo.DistributedVirtualPortgroup
			err = pc.RetrieveOne(ctx, pg.Reference(), []string{"config.logicalSwitchUuid", "config.segmentId"}, &props)
			if err != nil {
				t.Fatal(err)
			}

			net, err := finder.Network(ctx, props.Config.LogicalSwitchUuid)
			if err != nil {
				t.Fatal(err)
			}

			if net.Reference() != pg.Reference() {
				t.Errorf("%s vs %s", net.Reference(), pg.Reference())
			}

			net, err = finder.Network(ctx, props.Config.SegmentId)
			if err != nil {
				t.Fatal(err)
			}

			networks, err := finder.NetworkList(ctx, props.Config.SegmentId)
			if err != nil {
				t.Fatal(err)
			}
			if len(networks) != 1 {
				t.Errorf("expected 1 network, found %d", len(networks))
			}
		}
	}, model)
}

func TestFindByID(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		find := find.NewFinder(c)

		vms, err := find.VirtualMachineList(ctx, "*")
		if err != nil {
			t.Fatal(err)
		}

		for _, vm := range vms {
			ref := vm.Reference()
			byRef, err := find.VirtualMachine(ctx, ref.String())
			if err != nil {
				t.Fatal(err)
			}
			if byRef.InventoryPath != vm.InventoryPath {
				t.Errorf("InventoryPath=%q", byRef.InventoryPath)
			}
			if byRef.Reference() != ref {
				t.Error(byRef.Reference())
			}
			_, err = find.VirtualMachine(ctx, ref.String()+"invalid")
			if err == nil {
				t.Error("expected error")
			}

			byID, err := find.VirtualMachine(ctx, ref.Value)
			if err != nil {
				t.Error(err)
			}
			if byID.InventoryPath != vm.InventoryPath {
				t.Errorf("InventoryPath=%q", byID.InventoryPath)
			}
			if byID.Reference() != ref {
				t.Error(byID.Reference())
			}
			_, err = find.VirtualMachine(ctx, ref.Value+"invalid")
			if err == nil {
				t.Error("expected error")
			}

		}
	})
}
