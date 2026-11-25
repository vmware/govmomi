// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tagging_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	_ "github.com/vmware/govmomi/vapi/simulator"
	"github.com/vmware/govmomi/vapi/tagging"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25"
)

func TestManager_List(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)
		if err := rc.Login(ctx, simulator.DefaultLogin); err != nil {
			t.Fatal(err)
		}
		m := tagging.NewManager(rc)
		old := tags.NewManager(rc) // used to create fixtures, there is no create API in tagging yet

		const n = 25 // > tagging.DefaultListPageSize
		var cats []string
		for i := range 2 {
			id, err := old.CreateCategory(ctx, &tags.Category{
				Name:            fmt.Sprintf("cat-%d", i),
				Cardinality:     "MULTIPLE",
				AssociableTypes: []string{"VirtualMachine"},
			})
			if err != nil {
				t.Fatal(err)
			}
			cats = append(cats, id)
		}
		for i := range n {
			if _, err := old.CreateTag(ctx, &tags.Tag{Name: fmt.Sprintf("tag-%d", i), CategoryID: cats[0]}); err != nil {
				t.Fatal(err)
			}
		}

		// More tags than one server page, all must be returned.
		all, err := m.ListTags(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if len(all) != n {
			t.Errorf("expected %d tags, got %d", n, len(all))
		}
		if all[0].Info.Category != cats[0] || all[0].Tag == "" {
			t.Errorf("unexpected item: %+v", all[0])
		}

		// Name filter.
		named, err := m.ListTags(ctx, "tag-1", "tag-2", "nope")
		if err != nil {
			t.Fatal(err)
		}
		if len(named) != 2 {
			t.Errorf("expected 2 tags, got %d", len(named))
		}

		// Same tag name in another category, names are only unique per category.
		if _, err = old.CreateTag(ctx, &tags.Tag{Name: "tag-1", CategoryID: cats[1]}); err != nil {
			t.Fatal(err)
		}
		dups, err := m.ListTags(ctx, "tag-1")
		if err != nil {
			t.Fatal(err)
		}
		if len(dups) != 2 {
			t.Errorf("expected 2 tags named tag-1, got %d", len(dups))
		}

		// Tags for a category, by category name.
		forCat, err := m.ListTagsForCategory(ctx, "cat-1")
		if err != nil {
			t.Fatal(err)
		}
		if len(forCat) != 1 || forCat[0].Info.Category != cats[1] {
			t.Errorf("unexpected tags for cat-1: %+v", forCat)
		}
		forCat, err = m.ListTagsForCategory(ctx, "cat-0")
		if err != nil {
			t.Fatal(err)
		}
		if len(forCat) != n {
			t.Errorf("expected %d tags for cat-0, got %d", n, len(forCat))
		}

		// Category and tag name together identify one tag.
		forCat, err = m.ListTagsForCategory(ctx, "cat-1", "tag-1")
		if err != nil {
			t.Fatal(err)
		}
		if len(forCat) != 1 || forCat[0].Info.Category != cats[1] || forCat[0].Info.Name != "tag-1" {
			t.Errorf("unexpected tags for cat-1/tag-1: %+v", forCat)
		}

		// Name that is not in the category.
		forCat, err = m.ListTagsForCategory(ctx, "cat-1", "tag-2")
		if err != nil {
			t.Fatal(err)
		}
		if len(forCat) != 0 {
			t.Errorf("expected no tags, got %+v", forCat)
		}

		// Unknown category.
		if _, err = m.ListTagsForCategory(ctx, "nope"); err == nil {
			t.Error("expected error for unknown category")
		}

		// Categories.
		citems, err := m.ListCategories(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if len(citems) != 2 {
			t.Errorf("expected 2 categories, got %d", len(citems))
		}
		citems, err = m.ListCategories(ctx, "cat-1")
		if err != nil {
			t.Fatal(err)
		}
		if len(citems) != 1 || citems[0].CategoryID != cats[1] || citems[0].Info.Name != "cat-1" ||
			citems[0].Info.Cardinality != "MULTIPLE" {
			t.Errorf("unexpected categories: %+v", citems)
		}
	})
}
