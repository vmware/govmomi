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

package library_test

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"

	_ "github.com/vmware/govmomi/vapi/simulator"
)

func ExampleManager_CreateLibrary() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c := rest.NewClient(vc)

		err := c.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			return err
		}

		ds, err := find.NewFinder(vc).DefaultDatastore(ctx)
		if err != nil {
			return err
		}

		m := library.NewManager(c)

		id, err := m.CreateLibrary(ctx, library.Library{
			Name: "example",
			Type: "LOCAL",
			Storage: []library.StorageBackings{{
				DatastoreID: ds.Reference().Value,
				Type:        "DATASTORE",
			}},
		})
		if err != nil {
			return err
		}

		l, err := m.GetLibraryByID(ctx, id)
		if err != nil {
			return err
		}

		fmt.Println("created library", l.Name)
		return nil
	})
	// Output: created library example
}
