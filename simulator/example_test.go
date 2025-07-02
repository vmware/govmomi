// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"context"
	"fmt"
	"log"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// Example boilerplate for starting a simulator initialized with an ESX model.
func ExampleESX() {
	ctx := context.Background()

	// ESXi model + initial set of objects (VMs, network, datastore)
	model := simulator.ESX()

	defer model.Remove()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, _ := govmomi.NewClient(ctx, s.URL, true)

	fmt.Printf("%s with %d host", c.Client.ServiceContent.About.ApiType, model.Count().Host)
	// Output: HostAgent with 1 host
}

// Example for starting a simulator with empty inventory, similar to a fresh install of vCenter.
func ExampleModel() {
	ctx := context.Background()

	model := simulator.VPX()
	model.Datacenter = 0 // No DC == no inventory

	defer model.Remove()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, _ := govmomi.NewClient(ctx, s.URL, true)

	fmt.Printf("%s with %d hosts", c.Client.ServiceContent.About.ApiType, model.Count().Host)
	// Output: VirtualCenter with 0 hosts
}

// Example boilerplate for starting a simulator initialized with a vCenter model.
func ExampleVPX() {
	ctx := context.Background()

	// vCenter model + initial set of objects (cluster, hosts, VMs, network, datastore, etc)
	model := simulator.VPX()

	defer model.Remove()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, _ := govmomi.NewClient(ctx, s.URL, true)

	fmt.Printf("%s with %d hosts", c.Client.ServiceContent.About.ApiType, model.Count().Host)
	// Output: VirtualCenter with 4 hosts
}

// Run simplifies startup/cleanup of a simulator instance for example or testing purposes.
func ExampleModel_Run() {
	err := simulator.VPX().Run(func(ctx context.Context, c *vim25.Client) error {
		// Client has connected and logged in to a new simulator instance.
		// Server.Close and Model.Remove are called when this func returns.
		s, err := session.NewManager(c).UserSession(ctx)
		if err != nil {
			return err
		}
		fmt.Print(s.UserName)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// Output: user
}

// Test simplifies startup/cleanup of a simulator instance for testing purposes.
func ExampleTest() {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		// Client has connected and logged in to a new simulator instance.
		// Server.Close and Model.Remove are called when this func returns.
		s, err := session.NewManager(c).UserSession(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(s.UserName)
	})
	// Output: user
}

// Folder.AddOpaqueNetwork can be used to create an NSX backed OpaqueNetwork.
func ExampleFolder_AddOpaqueNetwork() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		sctx := ctx.(*simulator.Context)
		finder := find.NewFinder(c)

		// Find the network folder via vSphere API
		obj, err := finder.Folder(ctx, "network")
		if err != nil {
			return err
		}

		// Get vcsim's network Folder object
		folder := sctx.Map.Get(obj.Reference()).(*simulator.Folder)

		spec := types.OpaqueNetworkSummary{
			NetworkSummary: types.NetworkSummary{
				Name: "my-nsx-network",
			},
			OpaqueNetworkId:   "my-nsx-id",
			OpaqueNetworkType: "nsx.LogicalSwitch",
		}

		// Add NSX backed OpaqueNetwork, calling the simulator.Folder method directly.
		err = folder.AddOpaqueNetwork(sctx, spec)
		if err != nil {
			return err
		}

		// Find the OpaqueNetwork via vSphere API
		net, err := finder.Network(ctx, spec.Name)
		if err != nil {
			return err
		}

		nsx := net.(*object.OpaqueNetwork)
		summary, err := nsx.Summary(ctx)
		if err != nil {
			return err
		}

		// The summary fields should match those of the spec used to create it
		fmt.Printf("%s: %s", nsx.Name(), summary.OpaqueNetworkId)
		return nil
	})
	// Output: my-nsx-network: my-nsx-id
}

// AddDVPortgroup against vcsim can create both standard and nsx backed DistributedVirtualPortgroup networks
func ExampleDistributedVirtualSwitch_AddDVPortgroupTask() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)

		dvs0, err := finder.Network(ctx, "DVS0")
		if err != nil {
			return err
		}

		spec := types.DVPortgroupConfigSpec{
			Name:              "my-nsx-dvpg",
			LogicalSwitchUuid: "my-nsx-id",
		}

		dvs := dvs0.(*object.DistributedVirtualSwitch)
		task, err := dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{spec})
		if err != nil {
			return err
		}
		if err = task.Wait(ctx); err != nil {
			return err
		}

		pg0, err := finder.Network(ctx, spec.Name)
		if err != nil {
			return err
		}

		pg := pg0.(*object.DistributedVirtualPortgroup)

		var props mo.DistributedVirtualPortgroup
		err = pg.Properties(ctx, pg.Reference(), []string{"config"}, &props)
		if err != nil {
			return err
		}

		fmt.Printf("%s: %s %s", pg.Name(), props.Config.BackingType, props.Config.LogicalSwitchUuid)

		return nil
	})
	// Output: my-nsx-dvpg: nsx my-nsx-id
}
