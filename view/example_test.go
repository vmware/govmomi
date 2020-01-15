/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package view_test

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

// Create a view of all hosts in the inventory, printing host names that belong to a cluster and excluding standalone hosts.
func ExampleContainerView_Retrieve() {
	model := simulator.VPX()
	model.Datacenter = 2

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m := view.NewManager(c)
		kind := []string{"HostSystem"}

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			log.Fatal(err)
		}

		var hosts []mo.HostSystem
		var names []string

		err = v.Retrieve(ctx, kind, []string{"summary.config.name", "parent"}, &hosts)
		if err != nil {
			return err
		}

		for _, host := range hosts {
			if host.Parent.Type != "ClusterComputeResource" {
				continue
			}
			names = append(names, host.Summary.Config.Name)
		}

		sort.Strings(names)
		fmt.Println(names)

		return v.Destroy(ctx)
	}, model)
	// Output: [DC0_C0_H0 DC0_C0_H1 DC0_C0_H2 DC1_C0_H0 DC1_C0_H1 DC1_C0_H2]
}

// Create a view of all VMs in the inventory, printing VM names that end with "_VM1".
func ExampleContainerView_RetrieveWithFilter() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m := view.NewManager(c)
		kind := []string{"VirtualMachine"}

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			log.Fatal(err)
		}

		var vms []mo.VirtualMachine
		var names []string

		err = v.RetrieveWithFilter(ctx, kind, []string{"name"}, &vms, property.Filter{"name": "*_VM1"})
		if err != nil {
			return err
		}

		for _, vm := range vms {
			names = append(names, vm.Name)
		}

		sort.Strings(names)
		fmt.Println(names)

		return v.Destroy(ctx)
	})
	// Output: [DC0_C0_RP0_VM1 DC0_H0_VM1]
}

// Create a view of all VMs in a specific subfolder, powering off all VMs within
func ExampleContainerView_Find() {
	model := simulator.VPX()
	model.Folder = 1 // put everything inside subfolders

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		folder, err := object.NewSearchIndex(c).FindByInventoryPath(ctx, "/F0/DC0/vm/F0")
		if err != nil {
			return err
		}

		m := view.NewManager(c)
		kind := []string{"VirtualMachine"} // include VMs only, ignoring other object types

		// Root of the view is the subfolder moid (true == recurse into any subfolders of the root)
		v, err := m.CreateContainerView(ctx, folder.Reference(), kind, true)
		if err != nil {
			log.Fatal(err)
		}

		vms, err := v.Find(ctx, kind, property.Filter{})
		if err != nil {
			return err
		}

		for _, id := range vms {
			vm := object.NewVirtualMachine(c, id)
			task, err := vm.PowerOff(ctx)
			if err != nil {
				return err
			}

			if err = task.Wait(ctx); err != nil {
				return err
			}
		}

		fmt.Println(len(vms))

		return v.Destroy(ctx)
	}, model)
	// Output: 4
}
