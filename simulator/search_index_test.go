// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"strings"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func TestSearchIndex(t *testing.T) {
	ctx := context.Background()

	for _, model := range []*Model{ESX(), VPX()} {
		defer model.Remove()
		err := model.Create()
		if err != nil {
			t.Fatal(err)
		}

		s := model.Service.NewServer()
		defer s.Close()

		c, err := govmomi.NewClient(ctx, s.URL, true)
		if err != nil {
			t.Fatal(err)
		}

		finder := find.NewFinder(c.Client, false)
		dc, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatal(err)
		}

		finder.SetDatacenter(dc)

		vms, err := finder.VirtualMachineList(ctx, "*")
		if err != nil {
			t.Fatal(err)
		}

		vm := model.Map().Get(vms[0].Reference()).(*VirtualMachine)

		si := object.NewSearchIndex(c.Client)

		ref, err := si.FindByDatastorePath(ctx, dc, vm.Config.Files.VmPathName)
		if err != nil {
			t.Fatal(err)
		}

		if ref.Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", ref, vm.Reference())
		}

		ref, err = si.FindByDatastorePath(ctx, dc, vm.Config.Files.VmPathName+"enoent")
		if err != nil {
			t.Fatal(err)
		}

		if ref != nil {
			t.Errorf("ref=%s", ref)
		}

		ref, err = si.FindByUuid(ctx, dc, vm.Config.Uuid, true, nil)
		if err != nil {
			t.Fatal(err)
		}

		if ref.Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", ref, vm.Reference())
		}

		refs, err := si.FindAllByUuid(ctx, dc, vm.Config.Uuid, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(refs) != 1 {
			t.Errorf("len(refs) %d != 1", len(refs))
		}
		if refs[0].Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", refs[0], vm.Reference())
		}

		ref, err = si.FindByUuid(ctx, dc, vm.Config.Uuid, true, types.NewBool(false))
		if err != nil {
			t.Fatal(err)
		}
		if ref.Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", ref, vm.Reference())
		}

		refs, err = si.FindAllByUuid(ctx, dc, vm.Config.Uuid, true, types.NewBool(false))
		if err != nil {
			t.Fatal(err)
		}
		if len(refs) != 1 {
			t.Errorf("len(refs) %d != 1", len(refs))
		}
		if refs[0].Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", refs[0], vm.Reference())
		}

		ref, err = si.FindByUuid(ctx, dc, vm.Config.InstanceUuid, true, types.NewBool(true))
		if err != nil {
			t.Fatal(err)
		}
		if ref.Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", ref, vm.Reference())
		}

		refs, err = si.FindAllByUuid(ctx, dc, vm.Config.InstanceUuid, true, types.NewBool(true))
		if err != nil {
			t.Fatal(err)
		}
		if len(refs) != 1 {
			t.Errorf("len(refs) %d != 1", len(refs))
		}
		if refs[0].Reference() != vm.Reference() {
			t.Errorf("moref mismatch %s != %s", refs[0], vm.Reference())
		}

		ref, err = si.FindByUuid(ctx, dc, vm.Config.Uuid, false, nil)
		if err != nil {
			t.Fatal(err)
		}
		if ref != nil {
			t.Error("expected nil")
		}

		refs, err = si.FindAllByUuid(ctx, dc, vm.Config.Uuid, false, nil)
		if err != nil {
			t.Fatal(err)
		}
		if refs != nil {
			t.Error("refs != nil")
		}

		host := model.Map().Any("HostSystem").(*HostSystem)

		ref, err = si.FindByUuid(ctx, dc, host.Summary.Hardware.Uuid, false, nil)
		if err != nil {
			t.Fatal(err)
		}
		if ref.Reference() != host.Reference() {
			t.Errorf("moref mismatch %s != %s", ref, host.Reference())
		}

		refs, err = si.FindAllByUuid(ctx, dc, host.Summary.Hardware.Uuid, false, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(refs) != 1 {
			t.Errorf("len(refs) %d != 1", len(refs))
		}
		if refs[0].Reference() != host.Reference() {
			t.Errorf("moref mismatch %s != %s", refs[0], host.Reference())
		}

		rootFolder, err := finder.Folder(ctx, "/")
		if err != nil {
			t.Fatal(err)
		}

		ref, err = si.FindByInventoryPath(ctx, "/")
		if err != nil {
			t.Fatal(err)
		}

		if ref.Reference() != rootFolder.Reference() {
			t.Errorf("moref mismatch %s != %s", ref, rootFolder.Reference())
		}

		{
			// Duplicate UUIDs to test multiple results from FindAllByUuid().

			if len(vms) == 1 {
				t.Errorf("len(vms) %d == 1", len(vms))
			}

			task, err := vms[1].Reconfigure(ctx, types.VirtualMachineConfigSpec{
				InstanceUuid: vm.Config.InstanceUuid,
				Uuid:         vm.Config.Uuid,
			})
			if err != nil {
				t.Fatal(err)
			}
			if err := task.Wait(ctx); err != nil {
				t.Fatal(err)
			}

			refs, err = si.FindAllByUuid(ctx, dc, vm.Config.InstanceUuid, true, types.NewBool(true))
			if err != nil {
				t.Fatal(err)
			}
			if len(refs) != 2 {
				t.Errorf("len(refs) %d != 2", len(refs))
			}

			refs, err = si.FindAllByUuid(ctx, dc, vm.Config.Uuid, true, nil)
			if err != nil {
				t.Fatal(err)
			}
			if len(refs) != 2 {
				t.Errorf("len(refs) %d != 2", len(refs))
			}
		}
	}
}

func TestSearchIndexFindChild(t *testing.T) {
	ctx := context.Background()

	model := VPX()
	model.Pool = 3

	defer model.Remove()
	err := model.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	si := object.NewSearchIndex(c.Client)

	tests := [][]string{
		// Datacenter -> host Folder -> Cluster -> HostSystem
		{"DC0", "host", "DC0_C0", "DC0_C0_H0"},
		// Datacenter -> host Folder -> ComputeResource -> HostSystem
		{"DC0", "host", "DC0_H0", "DC0_H0"},
		// Datacenter -> host Folder -> Cluster -> ResourcePool -> ResourcePool
		{"DC0", "host", "DC0_C0", "Resources", "DC0_C0_RP1"},
		// Datacenter -> host Folder -> Cluster -> ResourcePool -> VirtualMachine
		{"DC0", "host", "DC0_C0", "Resources", "DC0_C0_RP1", "DC0_C0_RP1_VM0"},
		// Datacenter -> vm Folder -> VirtualMachine
		{"DC0", "vm", "DC0_C0_RP1_VM0"},
	}

	root := c.ServiceContent.RootFolder

	for _, path := range tests {
		parent := root
		ipath := []string{""}

		for _, name := range path {
			ref, err := si.FindChild(ctx, parent, name)
			if err != nil {
				t.Fatal(err)
			}

			if ref == nil {
				t.Fatalf("failed to match %s using %s", name, parent)
			}

			parent = ref.Reference()

			ipath = append(ipath, name)

			iref, err := si.FindByInventoryPath(ctx, strings.Join(ipath, "/"))
			if err != nil {
				t.Fatal(err)
			}

			if iref.Reference() != ref.Reference() {
				t.Errorf("%s != %s", iref, ref)
			}
		}
	}

	ref, err := si.FindChild(ctx, root, "enoent")
	if err != nil {
		t.Fatal(err)
	}

	if ref != nil {
		t.Error("unexpected match")
	}

	root.Value = "enoent"
	_, err = si.FindChild(ctx, root, "enoent")
	if err == nil {
		t.Error("expected error")
	}

	if _, ok := soap.ToSoapFault(err).VimFault().(types.ManagedObjectNotFound); !ok {
		t.Error("expected ManagedObjectNotFound fault")
	}

	for _, path := range []string{"", "/enoent"} {
		ref, err := si.FindByInventoryPath(ctx, path)
		if err != nil {
			t.Fatal(err)
		}

		if ref != nil {
			t.Error("unexpected match")
		}
	}
}

// TestSearchIndexFindByIpWithEsxcliHelpers asserts that FindByIp/FindAllByIp tolerate
// the non-entity helper objects the registry accumulates.
//
// RetrieveManagedMethodExecuter/RetrieveDynamicTypeManager (esxcli passthrough) lazily
// register ManagedMethodExecuter and DynamicTypeManager objects, keyed with their owning
// host's reference Value but a different Type. Those don't embed an mo type, so
// converting them via asHostSystemMO/asVirtualMachineMO panics in getManagedObject.
// FindAllByIp used to iterate every registry object without checking the type first, so
// any FindByIp call after an esxcli passthrough crashed the request.
func TestSearchIndexFindByIpWithEsxcliHelpers(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	m.Datacenter = 1
	m.Cluster = 1

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}
	defer m.Remove()

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	sctx := m.Service.Context
	host := sctx.Map.Any("HostSystem").(*HostSystem)

	// Give the host a known management IP to search for.
	const ip = "10.20.30.40"
	host.Config.Network.Vnic[0].Spec.Ip.IpAddress = ip

	// Trigger the lazy helper-object registration exactly as an esxcli passthrough
	// caller would.
	host.RetrieveManagedMethodExecuter(sctx, nil)
	host.RetrieveDynamicTypeManager(sctx, nil)

	for _, kind := range []string{"ManagedMethodExecuter", "DynamicTypeManager"} {
		ref := types.ManagedObjectReference{Type: kind, Value: host.Self.Value}
		if sctx.Map.Get(ref) == nil {
			t.Fatalf("%s helper object was not registered", kind)
		}
	}

	si := object.NewSearchIndex(c.Client)

	// Before the fix this panicked while walking the registry, surfacing as
	// "http: panic serving ...: <ref> does not have an embedded mo type".
	ref, err := si.FindByIp(ctx, nil, ip, false)
	if err != nil {
		t.Fatalf("FindByIp: %s", err)
	}
	if ref == nil {
		t.Fatal("FindByIp returned no result for the host IP")
	}
	if ref.Reference() != host.Self {
		t.Errorf("FindByIp = %s, want %s", ref.Reference(), host.Self)
	}

	// vm=true takes the VirtualMachine branch, which had the same problem.
	if _, err = si.FindByIp(ctx, nil, ip, true); err != nil {
		t.Fatalf("FindByIp(vm=true): %s", err)
	}
}
