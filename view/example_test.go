// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package view_test

import (
	"context"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
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

func ExampleContainerView_retrieveClusters() {
	model := simulator.VPX()
	model.Cluster = 3

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m := view.NewManager(c)
		kind := []string{"ClusterComputeResource"}

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			log.Fatal(err)
		}

		var clusters []mo.ClusterComputeResource
		var names []string

		err = v.Retrieve(ctx, kind, []string{"name"}, &clusters)
		if err != nil {
			return err
		}

		for _, cluster := range clusters {
			names = append(names, cluster.Name)
		}

		sort.Strings(names)
		fmt.Println(names)

		return v.Destroy(ctx)
	}, model)
	// Output: [DC0_C0 DC0_C1 DC0_C2]
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

		err = v.RetrieveWithFilter(ctx, kind, []string{"name"}, &vms, property.Match{"name": "*_VM1"})
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

		vms, err := v.Find(ctx, kind, property.Match{})
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

// This example uses a single PropertyCollector with ListView for waiting on updates to N tasks
func ExampleListView_tasks() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		list, err := view.NewManager(c).CreateListView(ctx, nil)
		if err != nil {
			return err
		}

		defer list.Destroy(ctx)

		vms, err := find.NewFinder(c).VirtualMachineList(ctx, "*")
		if err != nil {
			return err
		}

		result := map[types.TaskInfoState]int{}
		n := len(vms)
		p := property.DefaultCollector(c)

		// wait for any updates to tasks in our list view
		filter := new(property.WaitFilter).Add(list.Reference(), "Task", []string{"info"}, list.TraversalSpec())

		var werr error
		var wg sync.WaitGroup
		wg.Add(1)
		go func() { // WaitForUpdates blocks until func returns true
			defer wg.Done()
			werr = property.WaitForUpdates(ctx, p, filter, func(updates []types.ObjectUpdate) bool {
				for _, update := range updates {
					for _, change := range update.ChangeSet {
						info := change.Val.(types.TaskInfo)

						switch info.State {
						case types.TaskInfoStateSuccess, types.TaskInfoStateError:
							_, _ = list.Remove(ctx, []types.ManagedObjectReference{update.Obj})
							result[info.State]++
							n--
							if n == 0 {
								return true
							}
						}
					}
				}

				return false
			})
		}()

		for _, vm := range vms {
			task, err := vm.PowerOff(ctx)
			if err != nil {
				return err
			}
			_, err = list.Add(ctx, []types.ManagedObjectReference{task.Reference()})
			if err != nil {
				return err
			}
		}

		wg.Wait() // wait until all tasks complete and WaitForUpdates returns

		for state, n := range result {
			fmt.Printf("%s=%d", state, n)
		}

		return werr
	})
	// Output: success=4
}

// viewOfVM returns a spec to traverse a ListView of ContainerView, for tracking VM updates.
func viewOfVM(ref types.ManagedObjectReference, pathSet []string) types.CreateFilter {
	return types.CreateFilter{
		Spec: types.PropertyFilterSpec{
			ObjectSet: []types.ObjectSpec{{
				Obj:  ref,
				Skip: types.NewBool(true),
				SelectSet: []types.BaseSelectionSpec{
					// ListView --> ContainerView
					&types.TraversalSpec{
						Type: "ListView",
						Path: "view",
						SelectSet: []types.BaseSelectionSpec{
							&types.SelectionSpec{
								Name: "visitViews",
							},
						},
					},
					// ContainerView --> VM
					&types.TraversalSpec{
						SelectionSpec: types.SelectionSpec{
							Name: "visitViews",
						},
						Type: "ContainerView",
						Path: "view",
					},
				},
			}},
			PropSet: []types.PropertySpec{{
				Type:    "VirtualMachine",
				PathSet: pathSet,
			}},
		},
	}
}

// Example of using WaitForUpdates with a ListView of ContainerView.
// Each container root is a Cluster view of VirtualMachines.
// Modifying the ListView changes which VirtualMachine updates are returned by WaitForUpdates.
func ExampleListView_ofContainerView() {
	model := simulator.VPX()
	model.Cluster = 3

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m := view.NewManager(c)

		kind := []string{"ClusterComputeResource"}
		root, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			return err
		}

		clusters, err := root.Find(ctx, kind, property.Match{})
		if err != nil {
			return err
		}

		list, err := m.CreateListView(ctx, nil)
		if err != nil {
			return err
		}

		kind = []string{"VirtualMachine"}
		views := make(map[types.ManagedObjectReference]*view.ContainerView)

		for _, cluster := range clusters {
			cv, err := m.CreateContainerView(ctx, cluster, kind, true)
			if err != nil {
				return err
			}

			_, err = list.Add(ctx, []types.ManagedObjectReference{cv.Reference()})

			views[cluster] = cv
		}

		pc := property.DefaultCollector(c)

		pathSet := []string{"config.extraConfig"}
		pf, err := pc.CreateFilter(ctx, viewOfVM(list.Reference(), pathSet))
		if err != nil {
			return err
		}
		defer pf.Destroy(ctx)

		// MaxWaitSeconds value of 0 causes WaitForUpdatesEx to do one update calculation and return any results.
		req := types.WaitForUpdatesEx{
			This: pc.Reference(),
			Options: &types.WaitOptions{
				MaxWaitSeconds: types.NewInt32(0),
			},
		}

		updates := func() ([]types.ObjectUpdate, error) {
			res, err := methods.WaitForUpdatesEx(ctx, c.Client, &req)
			if err != nil {
				return nil, err
			}

			set := res.Returnval

			if set == nil {
				return nil, nil
			}

			req.Version = set.Version

			if len(set.FilterSet) == 0 {
				return nil, nil
			}

			return set.FilterSet[0].ObjectSet, nil
		}

		set, err := updates()
		if err != nil {
			return err
		}

		all := make([]types.ManagedObjectReference, len(set))
		for i := range set {
			all[i] = set[i].Obj
		}

		fmt.Printf("Initial updates=%d\n", len(all))

		reconfig := func() {
			spec := types.VirtualMachineConfigSpec{
				ExtraConfig: []types.BaseOptionValue{
					&types.OptionValue{Key: "time", Value: time.Now().String()},
				},
			}
			// Change state of all VMs in all Clusters
			for _, ref := range all {
				task, err := object.NewVirtualMachine(c, ref).Reconfigure(ctx, spec)
				if err != nil {
					panic(err)
				}
				if err = task.Wait(ctx); err != nil {
					panic(err)
				}
			}

			set, err := updates()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Reconfig updates=%d\n", len(set))
		}

		reconfig()

		cluster := views[clusters[0]].Reference()
		_, err = list.Remove(ctx, []types.ManagedObjectReference{cluster})
		if err != nil {
			return err
		}

		reconfig()

		_, err = list.Add(ctx, []types.ManagedObjectReference{cluster})
		if err != nil {
			return err
		}

		reconfig()

		return nil
	}, model)

	// Output:
	// Initial updates=6
	// Reconfig updates=6
	// Reconfig updates=4
	// Reconfig updates=6
}
