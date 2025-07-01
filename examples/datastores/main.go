// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

/*
This example program shows how the `view` and `property` packages can
be used to navigate a vSphere inventory structure using govmomi.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func main() {
	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		// Create a view of Datastore objects
		m := view.NewManager(c)

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"Datastore"}, true)
		if err != nil {
			return err
		}

		defer v.Destroy(ctx)

		// Retrieve summary property for all datastores
		// Reference: https://developer.broadcom.com/xapis/vsphere-web-services-api/latest/vim.Datastore.html
		var dss []mo.Datastore
		err = v.Retrieve(ctx, []string{"Datastore"}, []string{"summary"}, &dss)
		if err != nil {
			return err
		}

		// Print summary per datastore (see also: govc/datastore/info.go)

		tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
		fmt.Fprintf(tw, "Name:\tType:\tCapacity:\tFree:\n")

		for _, ds := range dss {
			fmt.Fprintf(tw, "%s\t", ds.Summary.Name)
			fmt.Fprintf(tw, "%s\t", ds.Summary.Type)
			fmt.Fprintf(tw, "%s\t", units.ByteSize(ds.Summary.Capacity))
			fmt.Fprintf(tw, "%s\t", units.ByteSize(ds.Summary.FreeSpace))
			fmt.Fprintf(tw, "\n")
		}

		_ = tw.Flush()

		return nil
	})
}
