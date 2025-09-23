// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func TestManagerCreateLibrary(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatal(err)
		}

		ds, err := find.NewFinder(vc).DefaultDatastore(ctx)
		if err != nil {
			t.Fatal(err)
		}

		m := library.NewManager(c)

		libName := "example"
		libType := "LOCAL"
		id, err := m.CreateLibrary(ctx, library.Library{
			Name: libName,
			Type: libType,
			Storage: []library.StorageBacking{{
				DatastoreID: ds.Reference().Value,
				Type:        "DATASTORE",
			}},
		})
		if err != nil {
			t.Fatal(err)
		}

		l, err := m.GetLibraryByID(ctx, id)
		if err != nil {
			t.Fatal(err)
		}

		if l.ID == "" {
			t.Fatal("library ID should be generated")
		}
		if l.ServerGUID == "" {
			t.Fatal("library server GUID should be generated")
		}
		if l.Name != libName {
			t.Fatalf("expected library name %s, got %s", libName, l.Name)
		}
		if l.Type != libType {
			t.Fatalf("expected library type %s, got %s", libType, l.Type)
		}
		if len(l.Storage) == 0 {
			t.Fatal("library should have a storage backing")
		}
		if l.Storage[0].Type != "DATASTORE" {
			t.Fatalf("expected library storage type DATASTORE, got %s", l.Storage[0].Type)
		}
		if l.Storage[0].DatastoreID != ds.Reference().Value {
			t.Fatalf("expected library datastore ref %s, got %s", ds.Reference().Value, l.Storage[0].DatastoreID)
		}
		if l.StateInfo == nil || l.StateInfo.State != "ACTIVE" {
			t.Fatal("library should have state ACTIVE")
		}
	})
}

func TestManagerLibraryUsage(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatal(err)
		}

		ds, err := find.NewFinder(vc).DefaultDatastore(ctx)
		if err != nil {
			t.Fatal(err)
		}

		m := library.NewManager(c)

		libName := "example"
		libType := "LOCAL"
		id, err := m.CreateLibrary(ctx, library.Library{
			Name: libName,
			Type: libType,
			Storage: []library.StorageBacking{{
				DatastoreID: ds.Reference().Value,
				Type:        "DATASTORE",
			}},
		})
		if err != nil {
			t.Fatal(err)
		}

		// Add library usage
		resourceUrn := "vmomi:service:wcp"
		addUsage := library.AddUsage{ResourceUrn: resourceUrn}
		usageID, err := m.AddLibraryUsage(ctx, id, addUsage)
		if err != nil {
			t.Fatal(err)
		}
		if usageID == "" {
			t.Fatal("library usage ID should be generated")
		}

		// Get library usage
		usage, err := m.GetLibraryUsage(ctx, id, usageID)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("returned usage: %+v", usage)
		if usage.ID != usageID {
			t.Fatalf("library usage ID should be matched: '%s', '%s'", usageID, usage.ID)
		}
		if usage.ResourceUrn != resourceUrn {
			t.Fatalf("library usage URN should be matched: '%s', '%s'", resourceUrn, usage.ResourceUrn)
		}
		if usage.AdditionTime == nil {
			t.Fatalf("library usage addition time should be set: '%s', '%s'", usageID, usage.ID)
		}

		// List library usages
		usageList, err := m.ListLibraryUsage(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		if len(usageList.LibraryUsageList) == 0 {
			t.Fatalf("Library usage should not be empty: %s", id)
		}

		l, err := m.GetLibraryByID(ctx, id)
		if err != nil {
			t.Fatal(err)
		}

		err = m.DeleteLibrary(ctx, l)
		if err == nil {
			t.Fatalf("Library %s in use should not allowed for delete", id)
		}

		// Apply library usage to items.
		trueValue := true
		updateSpec := library.Library{ID: id, Name: libName, Configuration: &library.Configuration{ApplyLibraryUsageToItems: &trueValue}}
		err = m.UpdateLibrary(ctx, &updateSpec)
		if err != nil {
			t.Fatalf("updating library %s with new configuration should not fail. err: %v", id, err)
		}
		// Add library item and deletion of the item should not be allowed.
		itemDesc := "test item description"
		itemID, err := m.CreateLibraryItem(ctx, library.Item{Name: "test-item", Description: &itemDesc, LibraryID: id})
		if err != nil {
			t.Fatalf("library item creation should not fail. err: %v", err)
		}
		err = m.DeleteLibraryItem(ctx, &library.Item{ID: itemID})
		if err == nil {
			t.Fatalf("library item %s deletion should fail. err: %v", itemID, err)

		}

		// Relax library usage to items.
		falseValue := false
		updateSpec = library.Library{ID: id, Name: libName, Configuration: &library.Configuration{ApplyLibraryUsageToItems: &falseValue}}
		err = m.UpdateLibrary(ctx, &updateSpec)
		if err != nil {
			t.Fatalf("updating library %s with new configuration should not fail. err: %v", id, err)
		}

		// Remove library item should be allowed.
		err = m.DeleteLibraryItem(ctx, &library.Item{ID: itemID})
		if err != nil {
			t.Fatalf("library item %s deletion should not fail. err: %v", itemID, err)

		}

		// Remove library usage
		err = m.RemoveLibraryUsage(ctx, id, usageID)
		if err != nil {
			t.Fatal(err)
		}

		// Force delete library
		err = m.ForceDeleteLibrary(ctx, l)
		if err != nil {
			t.Fatal(err)
		}
	})
}
