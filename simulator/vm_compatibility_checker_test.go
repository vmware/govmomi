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

func TestVmCompatibilityChecker(t *testing.T) {
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
		vmRef := vm.Reference()

		vmcc := object.NewVmCompatibilityChecker(c)

		t.Run("CheckCompatibility", func(t *testing.T) {
			results, err := vmcc.CheckCompatibility(
				ctx,
				vm.Reference(),
				nil,
				nil)

			if err != nil {
				t.Fatal(err)
			}

			for _, result := range results {
				if len(result.Error) > 0 {
					t.Fatal("result has errors")
				}
				if len(result.Warning) > 0 {
					t.Fatal("result has warnings")
				}
			}
		})

		t.Run("CheckVmConfig", func(t *testing.T) {
			t.Run("for existing VM", func(t *testing.T) {
				results, err := vmcc.CheckVmConfig(
					ctx,
					types.VirtualMachineConfigSpec{
						NumCPUs: 2,
					},
					&vmRef,
					nil,
					nil)

				if err != nil {
					t.Fatal(err)
				}

				for _, result := range results {
					if len(result.Error) > 0 {
						t.Fatal("result has errors")
					}
					if len(result.Warning) > 0 {
						t.Fatal("result has warnings")
					}
				}
			})
			t.Run("for new VM", func(t *testing.T) {
				results, err := vmcc.CheckVmConfig(
					ctx,
					types.VirtualMachineConfigSpec{
						Name:    "my-vm",
						NumCPUs: 2,
					},
					&vmRef,
					nil,
					nil)

				if err != nil {
					t.Fatal(err)
				}

				for _, result := range results {
					if len(result.Error) > 0 {
						t.Fatal("result has errors")
					}
					if len(result.Warning) > 0 {
						t.Fatal("result has warnings")
					}
				}
			})
		})
	})
}
