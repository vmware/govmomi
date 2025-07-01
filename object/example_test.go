// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"fmt"
	"log"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func ExampleResourcePool_Owner() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)

		for _, name := range []string{"DC0_H0_VM0", "DC0_C0_RP0_VM0"} {
			vm, err := finder.VirtualMachine(ctx, name)
			if err != nil {
				return err
			}

			pool, err := vm.ResourcePool(ctx)
			if err != nil {
				return err
			}

			owner, err := pool.Owner(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("%s owner is a %T\n", name, owner)
		}

		return nil
	})
	// Output:
	// DC0_H0_VM0 owner is a *object.ComputeResource
	// DC0_C0_RP0_VM0 owner is a *object.ClusterComputeResource
}

func ExampleVirtualMachine_CreateSnapshot() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		task, err := vm.CreateSnapshot(ctx, "backup", "Backup", false, false)
		if err != nil {
			return err
		}
		if err = task.Wait(ctx); err != nil {
			return err
		}

		id, err := vm.FindSnapshot(ctx, "backup")
		if err != nil {
			return err
		}

		var snapshot mo.VirtualMachineSnapshot
		err = vm.Properties(ctx, *id, []string{"config.hardware.device"}, &snapshot)
		if err != nil {
			return err
		}

		fmt.Printf("%d devices", len(snapshot.Config.Hardware.Device))

		return nil
	})
	// Output: 13 devices
}

func ExampleVirtualMachine_Customize() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}
		task, err := vm.PowerOff(ctx)
		if err != nil {
			return err
		}
		if err = task.Wait(ctx); err != nil {
			return err
		}

		spec := types.CustomizationSpec{
			NicSettingMap: []types.CustomizationAdapterMapping{
				types.CustomizationAdapterMapping{
					Adapter: types.CustomizationIPSettings{
						Ip: &types.CustomizationFixedIp{
							IpAddress: "192.168.1.100",
						},
						SubnetMask:    "255.255.255.0",
						Gateway:       []string{"192.168.1.1"},
						DnsServerList: []string{"192.168.1.1"},
						DnsDomain:     "ad.domain",
					},
				},
			},
			Identity: &types.CustomizationLinuxPrep{
				HostName: &types.CustomizationFixedName{
					Name: "hostname",
				},
				Domain:     "ad.domain",
				TimeZone:   "Etc/UTC",
				HwClockUTC: types.NewBool(true),
			},
			GlobalIPSettings: types.CustomizationGlobalIPSettings{
				DnsSuffixList: []string{"ad.domain"},
				DnsServerList: []string{"192.168.1.1"},
			},
		}

		task, err = vm.Customize(ctx, spec)
		if err != nil {
			return err
		}
		if err = task.Wait(ctx); err != nil {
			return err
		}

		task, err = vm.PowerOn(ctx)
		if err != nil {
			return err
		}
		if err = task.Wait(ctx); err != nil {
			return err
		}

		ip, err := vm.WaitForIP(ctx)
		if err != nil {
			return err
		}

		fmt.Println(ip)

		return nil
	})
	// Output: 192.168.1.100
}

func ExampleVirtualMachine_HostSystem() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		host, err := vm.HostSystem(ctx)
		if err != nil {
			return err
		}

		name, err := host.ObjectName(ctx)
		if err != nil {
			return err
		}

		fmt.Println(name)

		return nil
	})
	// Output: DC0_H0
}

func ExampleVirtualMachine_ResourcePool() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		if err != nil {
			return err
		}

		pool, err := vm.ResourcePool(ctx)
		if err != nil {
			return err
		}

		name, err := pool.ObjectName(ctx)
		if err != nil {
			return err
		}

		fmt.Println(name)

		// The InventoryPath field not populated unless Finder.ResourcePool() or
		// Finder.ResourcePoolList() was used. But we can populate it explicity.
		pool.InventoryPath, err = find.InventoryPath(ctx, c, pool.Reference())
		if err != nil {
			return err
		}

		fmt.Println(pool.InventoryPath)

		return nil
	})
	// Output:
	// Resources
	// /DC0/host/DC0_C0/Resources
}

func ExampleVirtualMachine_Clone() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			return err
		}

		finder.SetDatacenter(dc)

		vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		folders, err := dc.Folders(ctx)
		if err != nil {
			return err
		}

		spec := types.VirtualMachineCloneSpec{
			PowerOn: false,
		}

		task, err := vm.Clone(ctx, folders.VmFolder, "example-clone", spec)
		if err != nil {
			return err
		}

		info, err := task.WaitForResult(ctx)
		if err != nil {
			return err
		}

		clone := object.NewVirtualMachine(c, info.Result.(types.ManagedObjectReference))
		name, err := clone.ObjectName(ctx)
		if err != nil {
			return err
		}

		fmt.Println(name)

		return nil
	})
	// Output: example-clone
}

func ExampleFolder_CreateVM() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			return err
		}

		finder.SetDatacenter(dc)

		folders, err := dc.Folders(ctx)
		if err != nil {
			return err
		}

		pool, err := finder.ResourcePool(ctx, "DC0_C0/Resources")
		if err != nil {
			return err
		}

		spec := types.VirtualMachineConfigSpec{
			Name:    "example-vm",
			GuestId: string(types.VirtualMachineGuestOsIdentifierOtherGuest),
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0]",
			},
		}

		task, err := folders.VmFolder.CreateVM(ctx, spec, pool, nil)
		if err != nil {
			return err
		}

		info, err := task.WaitForResult(ctx)
		if err != nil {
			return err
		}

		vm := object.NewVirtualMachine(c, info.Result.(types.ManagedObjectReference))
		name, err := vm.ObjectName(ctx)
		if err != nil {
			return err
		}

		fmt.Println(name)

		return nil
	})
	// Output: example-vm
}

func ExampleVirtualMachine_Reconfigure() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		spec := types.VirtualMachineConfigSpec{Annotation: "example reconfig"}

		task, err := vm.Reconfigure(ctx, spec)
		if err != nil {
			return err
		}

		err = task.Wait(ctx)
		if err != nil {
			return err
		}

		var obj mo.VirtualMachine
		err = vm.Properties(ctx, vm.Reference(), []string{"config.annotation"}, &obj)
		if err != nil {
			return err
		}

		fmt.Println(obj.Config.Annotation)

		return nil
	})
	// Output: example reconfig
}

func ExampleCommon_Destroy() {
	model := simulator.VPX()
	model.Datastore = 2

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		// Change to "LocalDS_0" will cause ResourceInUse error,
		// as simulator VMs created by the VPX model use "LocalDS_0".
		ds, err := find.NewFinder(c).Datastore(ctx, "LocalDS_1")
		if err != nil {
			return err
		}

		task, err := ds.Destroy(ctx)
		if err != nil {
			return err
		}

		if err = task.Wait(ctx); err != nil {
			return err
		}

		fmt.Println("destroyed", ds.InventoryPath)
		return nil
	}, model)
	// Output: destroyed /DC0/datastore/LocalDS_1
}

func ExampleCommon_Rename() {
	model := simulator.VPX()

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		dc, err := find.NewFinder(c).Datacenter(ctx, "DC0")
		if err != nil {
			return err
		}

		task, err := dc.Rename(ctx, "MyDC")
		if err != nil {
			return err
		}

		if err = task.Wait(ctx); err != nil {
			return err
		}

		name, err := dc.ObjectName(ctx)
		if err != nil {
			return err
		}

		fmt.Println(name)
		return nil
	}, model)
	// Output: MyDC
}

func ExampleCustomFieldsManager_Set() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m, err := object.GetCustomFieldsManager(c)
		if err != nil {
			return err
		}

		any := []string{"ManagedEntity"}
		field, err := m.Add(ctx, "backup", any[0], nil, nil) // adds the custom field "backup" to all types
		if err != nil {
			return err
		}

		v, err := view.NewManager(c).CreateContainerView(ctx, c.ServiceContent.RootFolder, any, true)
		if err != nil {
			log.Fatal(err)
		}

		all, err := v.Find(ctx, any, nil) // gives us the count of all objects in the inventory
		if err != nil {
			return err
		}

		refs, err := v.Find(ctx, []string{"VirtualMachine", "Datastore"}, nil)
		if err != nil {
			return err
		}

		for _, ref := range refs {
			err = m.Set(ctx, ref, field.Key, "true") // sets the custom value "backup=true" on specific types
			if err != nil {
				return err
			}
		}

		// filter used to find objects with "backup=true"
		filter := property.Match{"customValue": &types.CustomFieldStringValue{
			CustomFieldValue: types.CustomFieldValue{Key: field.Key},
			Value:            "true",
		}}

		var objs []mo.ManagedEntity
		err = v.RetrieveWithFilter(ctx, any, []string{"name", "customValue"}, &objs, filter)
		if err != nil {
			return err
		}

		fmt.Printf("backup %d of %d objects", len(objs), len(all))
		return v.Destroy(ctx)
	})
	// Output: backup 5 of 22 objects
}

func ExampleCustomizationSpecManager_Info() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m := object.NewCustomizationSpecManager(c)
		info, err := m.Info(ctx)
		if err != nil {
			return err
		}

		for i := range info {
			item, err := m.GetCustomizationSpec(ctx, info[i].Name)
			if err != nil {
				return err
			}
			fmt.Printf("%s=%T\n", item.Info.Name, item.Spec.Identity)
		}
		return nil
	})
	// Output:
	// vcsim-linux=*types.CustomizationLinuxPrep
	// vcsim-linux-static=*types.CustomizationLinuxPrep
	// vcsim-windows-static=*types.CustomizationSysprep
	// vcsim-windows-domain=*types.CustomizationSysprep
}

func ExampleNetworkReference_EthernetCardBackingInfo() {
	model := simulator.VPX()
	model.OpaqueNetwork = 1 // Create 1 NSX backed OpaqueNetwork per DC

	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		// finder.Network returns an object.NetworkReference
		net, err := finder.Network(ctx, "DC0_NSX0")
		if err != nil {
			return err
		}

		// EthernetCardBackingInfo creates the backing for any network type:
		// "Network", "DistributedVirtualPortgroup" or "OpaqueNetwork"
		backing, err := net.EthernetCardBackingInfo(ctx)
		if err != nil {
			return err
		}

		device, err := object.EthernetCardTypes().CreateEthernetCard("e1000", backing)
		if err != nil {
			return err
		}

		err = vm.AddDevice(ctx, device)
		if err != nil {
			return err
		}

		list, err := vm.Device(ctx)
		if err != nil {
			return err
		}

		// The NSX OpaqueNetwork backing should replaced with the DVPG.
		dvpgNet, err := finder.Network(ctx, "DC0_NSXPG0")
		if err != nil {
			return err
		}

		dvpgBacking, err := dvpgNet.EthernetCardBackingInfo(ctx)
		if err != nil {
			return err
		}

		nics := list.SelectByType((*types.VirtualEthernetCard)(nil)) // All VM NICs (DC0_DVPG0 + DC0_NSX0)
		match := list.SelectByBackingInfo(dvpgBacking)               // VM NIC with DC0_NSX0 backing

		fmt.Printf("%d of %d NICs match backing\n", len(match), len(nics))

		return nil
	}, model)
	// Output: 1 of 2 NICs match backing
}

func ExampleVirtualDeviceList_SelectByBackingInfo() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		finder := find.NewFinder(c)
		vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		backing := &types.VirtualPCIPassthroughVmiopBackingInfo{
			Vgpu: "grid_v100-4q",
		}

		gpu := &types.VirtualPCIPassthrough{
			VirtualDevice: types.VirtualDevice{Backing: backing},
		}

		err = vm.AddDevice(ctx, gpu) // add a GPU to this VM
		if err != nil {
			return err
		}

		device, err := vm.Device(ctx) // get the VM's virtual device list
		if err != nil {
			return err
		}

		// find the device with the given backing info
		gpus := device.SelectByBackingInfo(backing)

		name := gpus[0].(*types.VirtualPCIPassthrough).Backing.(*types.VirtualPCIPassthroughVmiopBackingInfo).Vgpu

		fmt.Println(name)

		// example alternative to using SelectByBackingInfo for the use-case above:
		for i := range device {
			switch d := device[i].(type) {
			case *types.VirtualPCIPassthrough:
				switch b := d.Backing.(type) {
				case *types.VirtualPCIPassthroughVmiopBackingInfo:
					if b.Vgpu == backing.Vgpu {
						fmt.Println(b.Vgpu)
					}
				}
			}
		}

		return nil
	})

	// Output:
	// grid_v100-4q
	// grid_v100-4q
}

// Find a VirtualMachine's Cluster
func ExampleVirtualMachine_resourcePoolOwner() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		obj, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_C0_RP0_VM0")
		if err != nil {
			return err
		}

		pool, err := obj.ResourcePool(ctx)
		if err != nil {
			return err
		}

		// ResourcePool owner will be a ComputeResource or ClusterComputeResource
		cluster, err := pool.Owner(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("%s", cluster.Reference().Type)

		return nil
	})
	// Output: ClusterComputeResource
}

func ExampleHostConfigManager_OptionManager() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		m := view.NewManager(c)
		kind := []string{"HostSystem"}

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, kind, true)
		if err != nil {
			log.Fatal(err)
		}

		refs, err := v.Find(ctx, kind, nil)
		if err != nil {
			return err
		}

		setting := ""

		for _, ref := range refs {
			host := object.NewHostSystem(c, ref)
			m, err := host.ConfigManager().OptionManager(ctx)
			if err != nil {
				return err
			}

			opt := []types.BaseOptionValue{&types.OptionValue{
				Key:   "vcrun",
				Value: "Config.HostAgent.plugins.hostsvc.esxAdminsGroup",
			}}

			err = m.Update(ctx, opt)
			if err != nil {
				return err
			}

			opt, err = m.Query(ctx, "vcrun")
			if err != nil {
				return err
			}
			setting = opt[0].GetOptionValue().Value.(string)
		}

		fmt.Println(setting)

		return v.Destroy(ctx)
	})
	// Output: Config.HostAgent.plugins.hostsvc.esxAdminsGroup
}
