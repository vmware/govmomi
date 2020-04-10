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

package simulator_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25"
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
		dswitch := simulator.Map.Get(dvs.Reference()).(*simulator.DistributedVirtualSwitch)
		portgrp := simulator.Map.Get(pg.Reference()).(*simulator.DistributedVirtualPortgroup)
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
