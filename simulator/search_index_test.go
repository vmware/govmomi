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
