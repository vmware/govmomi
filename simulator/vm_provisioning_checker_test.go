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
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestVmProvisioningChecker(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("default datacenter not found: %s", err)
		}
		finder.SetDatacenter(datacenter)
		vmList, err := finder.VirtualMachineList(ctx, "*")
		if len(vmList) == 0 {
			t.Fatal("vmList == 0")
		}
		vm := vmList[0]

		vmpc := object.NewVmProvisioningChecker(c)

		t.Run("CheckRelocate", func(t *testing.T) {
			results, err := vmpc.CheckRelocate(
				ctx,
				vm.Reference(),
				types.VirtualMachineRelocateSpec{})

			for _, result := range results {
				if err != nil {
					t.Fatal(err)
				}
				if len(result.Error) > 0 {
					t.Fatal("result has errors")
				}
				if len(result.Warning) > 0 {
					t.Fatal("result has warnings")
				}
			}
		})
	})
}
