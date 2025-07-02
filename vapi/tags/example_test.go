// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tags_test

import (
	"context"
	"fmt"
	"log"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func ExampleManager_CreateTag() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)
		_ = c.Login(ctx, simulator.DefaultLogin)

		m := tags.NewManager(c)

		id, err := m.CreateCategory(ctx, &tags.Category{
			AssociableTypes: []string{"VirtualMachine"},
			Cardinality:     "SINGLE",
			Description:     "This is My Category",
			Name:            "my-category",
		})
		if err != nil {
			return err
		}

		id, err = m.CreateTag(ctx, &tags.Tag{
			CategoryID:  id,
			Description: "This is My Tag",
			Name:        "my-tag",
		})
		if err != nil {
			return err
		}

		tag, err := m.GetTag(ctx, id)
		if err != nil {
			return err
		}

		fmt.Println(tag.Name)
		return nil
	})
	// Output: my-tag
}

func ExampleManager_GetAttachedTagsOnObjects() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)
		_ = c.Login(ctx, simulator.DefaultLogin)

		m := tags.NewManager(c)

		id, err := m.CreateCategory(ctx, &tags.Category{Name: "my-category"})
		if err != nil {
			return err
		}

		id, err = m.CreateTag(ctx, &tags.Tag{CategoryID: id, Name: "my-tag"})
		if err != nil {
			return err
		}

		v, err := view.NewManager(vc).CreateContainerView(ctx, vc.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
		if err != nil {
			log.Fatal(err)
		}

		vms, err := v.Find(ctx, nil, property.Match{}) // List all VMs in the inventory
		if err != nil {
			return err
		}
		refs := make([]mo.Reference, len(vms)) // Convert list type
		for i := range vms {
			refs[i] = vms[i]
		}

		for i := 0; i < len(refs)/2; i++ { // AttachTag to half of the VMs
			if err = m.AttachTag(ctx, id, refs[i]); err != nil {
				return err
			}
		}

		attached, err := m.GetAttachedTagsOnObjects(ctx, refs) // Get AttachedTags for all VMs
		if err != nil {
			return err
		}

		n := 0
		for _, a := range attached { // Count tags attached to all VMs
			n += len(a.Tags)
		}

		fmt.Printf("%d of %d vms are tagged", n, len(vms))
		return nil
	})
	// Output: 2 of 4 vms are tagged
}
