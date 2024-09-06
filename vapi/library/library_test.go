/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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
