// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func TestReconfigureVMTagSpecs(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		// Create REST client for VAPI tags
		rc := rest.NewClient(c)
		err := rc.Login(ctx, simulator.DefaultLogin)
		require.NoError(t, err)

		tagManager := tags.NewManager(rc)

		// Create a category
		categoryID, err := tagManager.CreateCategory(ctx, &tags.Category{
			Name:            "test-category",
			Description:     "Test category for TagSpec",
			Cardinality:     "MULTIPLE",
			AssociableTypes: []string{"VirtualMachine"},
		})
		require.NoError(t, err, "CreateCategory")

		// Create two tags
		tag1ID, err := tagManager.CreateTag(ctx, &tags.Tag{
			Name:        "test-tag-1",
			Description: "Test tag 1",
			CategoryID:  categoryID,
		})
		require.NoError(t, err, "CreateTag 1")

		tag2ID, err := tagManager.CreateTag(ctx, &tags.Tag{
			Name:        "test-tag-2",
			Description: "Test tag 2",
			CategoryID:  categoryID,
		})
		require.NoError(t, err, "CreateTag 2")

		// Get the category to retrieve its name
		category, err := tagManager.GetCategory(ctx, categoryID)
		require.NoError(t, err, "GetCategory")

		// Find a VM to reconfigure
		finder := find.NewFinder(c, false)
		dc, err := finder.DefaultDatacenter(ctx)
		require.NoError(t, err)
		finder.SetDatacenter(dc)

		vms, err := finder.VirtualMachineList(ctx, "*")
		require.NoError(t, err)
		require.NotEmpty(t, vms, "no VMs found")
		vm := vms[0]

		// Verify no tags are attached initially
		attachedTags, err := tagManager.ListAttachedTags(ctx, vm.Reference())
		require.NoError(t, err, "ListAttachedTags (initial)")
		require.Empty(t, attachedTags, "expected 0 tags initially")

		// Reconfigure VM to attach tag1 using TagSpecs
		spec := types.VirtualMachineConfigSpec{
			TagSpecs: []types.TagSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationAdd,
					},
					Id: types.TagId{
						NameId: &types.TagIdNameId{
							Tag:      "test-tag-1",
							Category: category.Name,
						},
					},
				},
			},
		}

		task, err := vm.Reconfigure(ctx, spec)
		require.NoError(t, err, "Reconfigure (attach tag1)")
		err = task.Wait(ctx)
		require.NoError(t, err, "Reconfigure task (attach tag1)")

		// Verify tag1 is attached
		attachedTags, err = tagManager.ListAttachedTags(ctx, vm.Reference())
		require.NoError(t, err, "ListAttachedTags (after attach tag1)")
		require.Len(t, attachedTags, 1, "expected 1 tag after attach")
		require.Equal(t, tag1ID, attachedTags[0], "expected tag1")

		// Reconfigure VM to attach tag2 as well
		spec = types.VirtualMachineConfigSpec{
			TagSpecs: []types.TagSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationAdd,
					},
					Id: types.TagId{
						NameId: &types.TagIdNameId{
							Tag:      "test-tag-2",
							Category: category.Name,
						},
					},
				},
			},
		}

		task, err = vm.Reconfigure(ctx, spec)
		require.NoError(t, err, "Reconfigure (attach tag2)")
		err = task.Wait(ctx)
		require.NoError(t, err, "Reconfigure task (attach tag2)")

		// Verify both tags are attached
		attachedTags, err = tagManager.ListAttachedTags(ctx, vm.Reference())
		require.NoError(t, err, "ListAttachedTags (after attach tag2)")
		require.Len(t, attachedTags, 2, "expected 2 tags after attaching both")
		require.ElementsMatch(t, []string{tag1ID, tag2ID}, attachedTags, "both tags should be attached")

		// Reconfigure VM to detach tag1
		spec = types.VirtualMachineConfigSpec{
			TagSpecs: []types.TagSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationRemove,
					},
					Id: types.TagId{
						NameId: &types.TagIdNameId{
							Tag:      "test-tag-1",
							Category: category.Name,
						},
					},
				},
			},
		}

		task, err = vm.Reconfigure(ctx, spec)
		require.NoError(t, err, "Reconfigure (detach tag1)")
		err = task.Wait(ctx)
		require.NoError(t, err, "Reconfigure task (detach tag1)")

		// Verify only tag2 remains
		attachedTags, err = tagManager.ListAttachedTags(ctx, vm.Reference())
		require.NoError(t, err, "ListAttachedTags (after detach tag1)")
		require.Len(t, attachedTags, 1, "expected 1 tag after detach")
		require.Equal(t, tag2ID, attachedTags[0], "expected tag2")

		// Reconfigure VM to detach tag2
		spec = types.VirtualMachineConfigSpec{
			TagSpecs: []types.TagSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationRemove,
					},
					Id: types.TagId{
						NameId: &types.TagIdNameId{
							Tag:      "test-tag-2",
							Category: category.Name,
						},
					},
				},
			},
		}

		task, err = vm.Reconfigure(ctx, spec)
		require.NoError(t, err, "Reconfigure (detach tag2)")
		err = task.Wait(ctx)
		require.NoError(t, err, "Reconfigure task (detach tag2)")

		// Verify no tags remain
		attachedTags, err = tagManager.ListAttachedTags(ctx, vm.Reference())
		require.NoError(t, err, "ListAttachedTags (after detach all)")
		require.Empty(t, attachedTags, "expected 0 tags after detaching all")
	})
}
