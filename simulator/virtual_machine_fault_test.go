// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestSwitchMembers(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)
		dvs, _ := finder.Network(ctx, "DVS0")
		pg, _ := finder.Network(ctx, "DC0_DVPG0")
		vm, _ := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		host, _ := vm.HostSystem(ctx)

		// The proper way to remove a host from is with dvs.Reconfigure + types.ConfigSpecOperationRemove,
		// but that would fail here with ResourceInUse. And creating a new DVS + PG is cumbersome,
		// so just force the removal of host from the DVS + PG, which has the same effect on vm.AddDevice().
		dswitch := simulator.Map(ctx).Get(dvs.Reference()).(*simulator.DistributedVirtualSwitch)
		portgrp := simulator.Map(ctx).Get(pg.Reference()).(*simulator.DistributedVirtualPortgroup)
		simulator.RemoveReference(&dswitch.Summary.HostMember, host.Reference())
		simulator.RemoveReference(&portgrp.Host, host.Reference())

		backing, _ := pg.EthernetCardBackingInfo(ctx)
		nic, _ := object.EthernetCardTypes().CreateEthernetCard("", backing)

		err := vm.AddDevice(ctx, nic)
		if err == nil {
			t.Fatal("expected error")
		}

		fault := err.(task.Error).Fault()
		invalid, ok := fault.(*types.InvalidArgument)
		if !ok {
			t.Fatalf("unexpected fault=%T", fault)
		}

		if invalid.InvalidProperty != "spec.deviceChange.device.port.switchUuid" {
			t.Errorf("unexpected property=%s", invalid.InvalidProperty)
		}
	})
}

func TestMultiSwitchMembers(t *testing.T) {
	model := simulator.VPX()
	model.PortgroupNSX = 1

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)
		dc, _ := finder.Datacenter(ctx, "DC0")
		f, _ := dc.Folders(ctx)
		vm, _ := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		pg0, _ := finder.Network(ctx, "DC0_NSXPG0")

		var props0 mo.DistributedVirtualPortgroup
		_ = dc.Properties(ctx, pg0.Reference(), []string{"config"}, &props0)

		// Create a 2nd DVS "DVS1" with no hosts attached and 1 PG "DC0_NSXPG1",
		// using the same LogicalSwitchUuid as DC0_NSXPG0
		_, _ = f.NetworkFolder.CreateDVS(ctx, types.DVSCreateSpec{
			ConfigSpec: &types.VMwareDVSConfigSpec{
				DVSConfigSpec: types.DVSConfigSpec{
					Name: "DVS1",
				},
			},
		})
		dvs, _ := finder.Network(ctx, "DVS1")
		dswitch := dvs.(*object.DistributedVirtualSwitch)
		_, _ = dswitch.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{{
			Name:              "DC0_NSXPG1",
			LogicalSwitchUuid: props0.Config.LogicalSwitchUuid,
		}})
		pg1, _ := finder.Network(ctx, "DC0_NSXPG1")

		backing, _ := pg1.EthernetCardBackingInfo(ctx)
		nic, _ := object.EthernetCardTypes().CreateEthernetCard("", backing)

		err := vm.AddDevice(ctx, nic)
		if err == nil {
			t.Fatal("expected error")
		}

		fault := err.(task.Error).Fault()
		invalid, ok := fault.(*types.InvalidArgument)
		if !ok {
			t.Fatalf("unexpected fault=%T", fault)
		}

		if invalid.InvalidProperty != "spec.deviceChange.device.port.switchUuid" {
			t.Errorf("unexpected property=%s", invalid.InvalidProperty)
		}

		var props1 mo.DistributedVirtualPortgroup
		_ = dc.Properties(ctx, pg1.Reference(), []string{"config"}, &props1)

		if props0.Config.LogicalSwitchUuid != props1.Config.LogicalSwitchUuid {
			t.Error("LogicalSwitchUuid should be the same")
		}
	}, model)
}
