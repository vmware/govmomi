// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func TestResourcePool(t *testing.T) {
	ctx := context.Background()

	m := &Model{
		ServiceContent: esx.ServiceContent,
		RootFolder:     esx.RootFolder,
	}

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	c := m.Service.client()

	finder := find.NewFinder(c, false)
	finder.SetDatacenter(object.NewDatacenter(c, esx.Datacenter.Reference()))

	spec := types.DefaultResourceConfigSpec()

	parent := object.NewResourcePool(c, esx.ResourcePool.Self)

	spec.CpuAllocation.Reservation = nil
	// missing required field (Reservation) for create
	_, err = parent.Create(ctx, "fail", spec)
	if err == nil {
		t.Error("expected error")
	}
	spec = types.DefaultResourceConfigSpec()

	// can't destroy a root pool
	task, err := parent.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err == nil {
		t.Fatal("expected error destroying a root pool")
	}

	// create a child pool
	childName := uuid.New().String()

	child, err := parent.Create(ctx, childName, spec)
	if err != nil {
		t.Fatal(err)
	}

	if child.Reference() == esx.ResourcePool.Self {
		t.Error("expected new pool Self reference")
	}

	*spec.CpuAllocation.Reservation = -1
	// invalid field value (Reservation) for update
	err = child.UpdateConfig(ctx, "", &spec)
	if err == nil {
		t.Error("expected error")
	}

	// valid config update
	*spec.CpuAllocation.Reservation = 10
	err = child.UpdateConfig(ctx, "", &spec)
	if err != nil {
		t.Error(err)
	}

	var p mo.ResourcePool
	err = child.Properties(ctx, child.Reference(), []string{"config.cpuAllocation"}, &p)
	if err != nil {
		t.Error(err)
	}

	if *p.Config.CpuAllocation.Reservation != 10 {
		t.Error("config not updated")
	}

	// duplicate name
	_, err = parent.Create(ctx, childName, spec)
	if err == nil {
		t.Error("expected error")
	}

	// create a grandchild pool
	grandChildName := uuid.New().String()
	_, err = child.Create(ctx, grandChildName, spec)
	if err != nil {
		t.Fatal(err)
	}

	// create sibling (of the grand child) pool
	siblingName := uuid.New().String()
	_, err = child.Create(ctx, siblingName, spec)
	if err != nil {
		t.Fatal(err)
	}

	// finder should return the 2 grand children
	pools, err := finder.ResourcePoolList(ctx, "*/Resources/"+childName+"/*")
	if err != nil {
		t.Fatal(err)
	}
	if len(pools) != 2 {
		t.Fatalf("len(pools) == %d", len(pools))
	}

	// destroy the child
	task, err = child.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// finder should error not found after destroying the child
	_, err = finder.ResourcePoolList(ctx, "*/Resources/"+childName+"/*")
	if err == nil {
		t.Fatal("expected not found error")
	}

	// since the child was destroyed, grand child pools should now be children of the root pool
	pools, err = finder.ResourcePoolList(ctx, "*/Resources/*")
	if err != nil {
		t.Fatal(err)
	}

	if len(pools) != 2 {
		t.Fatalf("len(pools) == %d", len(pools))
	}
}

func TestCreateVAppESX(t *testing.T) {
	ctx := context.Background()

	m := ESX()
	m.Datastore = 0
	m.Machine = 0

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	c := m.Service.client()

	parent := object.NewResourcePool(c, esx.ResourcePool.Self)

	rspec := types.DefaultResourceConfigSpec()
	vspec := NewVAppConfigSpec()

	_, err = parent.CreateVApp(ctx, "myapp", rspec, vspec, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	fault := soap.ToSoapFault(err).Detail.Fault

	if reflect.TypeOf(fault) != reflect.TypeOf(&types.MethodDisabled{}) {
		t.Errorf("fault=%#v", fault)
	}
}

func TestCreateVAppVPX(t *testing.T) {
	ctx := context.Background()

	m := VPX()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	defer m.Remove()

	c := m.Service.client()

	pool := m.Map().Any("ResourcePool")
	parent := object.NewResourcePool(c, pool.Reference())
	rspec := types.DefaultResourceConfigSpec()
	vspec := NewVAppConfigSpec()

	vapp, err := parent.CreateVApp(ctx, "myapp", rspec, vspec, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = parent.CreateVApp(ctx, "myapp", rspec, vspec, nil)
	if err == nil {
		t.Error("expected error")
	}

	spec := types.VirtualMachineConfigSpec{
		GuestId: string(types.VirtualMachineGuestOsIdentifierOtherGuest),
		Files: &types.VirtualMachineFileInfo{
			VmPathName: "[LocalDS_0]",
		},
	}

	for _, fail := range []bool{true, false} {
		task, cerr := vapp.CreateChildVM(ctx, spec, nil)
		if cerr != nil {
			t.Fatal(err)
		}

		cerr = task.Wait(ctx)
		if fail {
			if cerr == nil {
				t.Error("expected error")
			}
		} else {
			if cerr != nil {
				t.Error(err)
			}
		}

		spec.Name = "test"
	}

	si := object.NewSearchIndex(c)
	vm, err := si.FindChild(ctx, vapp, spec.Name)
	if err != nil {
		t.Fatal(err)
	}

	if vm == nil {
		t.Errorf("FindChild(%s)==nil", spec.Name)
	}

	ref := m.Map().Get(m.Map().getEntityDatacenter(pool).VmFolder).Reference()
	folder, err := object.NewFolder(c, ref).CreateFolder(ctx, "myapp-clone")
	if err != nil {
		t.Fatal(err)
	}

	cspec := types.VAppCloneSpec{
		VmFolder: types.NewReference(folder.Reference()),
	}

	task, err := vapp.Clone(ctx, "myapp-clone", parent.Reference(), cspec)
	if err != nil {
		t.Fatal(err)
	}

	res, err := task.WaitForResult(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	var apps []mo.VirtualApp
	refs := []types.ManagedObjectReference{vapp.Reference(), res.Result.(types.ManagedObjectReference)}
	err = property.DefaultCollector(c).Retrieve(ctx, refs, []string{"vm"}, &apps)
	if err != nil {
		t.Fatal(err)
	}

	if len(apps) != 2 {
		t.Errorf("apps=%d", len(apps))
	}

	for _, app := range apps {
		if len(app.Vm) != 1 {
			t.Errorf("app %s vm=%d", app.Reference(), len(app.Vm))
		}
	}

	task, err = vapp.Destroy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestResourcePoolValidation(t *testing.T) {
	tests := []func() bool{
		func() bool {
			return allResourceFieldsSet(&types.ResourceAllocationInfo{})
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			spec.CpuAllocation.Limit = nil
			return allResourceFieldsSet(&spec.CpuAllocation)
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			spec.CpuAllocation.Reservation = nil
			return allResourceFieldsSet(&spec.CpuAllocation)
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			spec.CpuAllocation.ExpandableReservation = nil
			return allResourceFieldsSet(&spec.CpuAllocation)
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			spec.CpuAllocation.Shares = nil
			return allResourceFieldsSet(&spec.CpuAllocation)
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			spec.CpuAllocation.Reservation = types.NewInt64(-1)
			return allResourceFieldsValid(&spec.CpuAllocation)
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			spec.CpuAllocation.Limit = types.NewInt64(-100)
			return allResourceFieldsValid(&spec.CpuAllocation)
		},
		func() bool {
			spec := types.DefaultResourceConfigSpec()
			shares := spec.CpuAllocation.Shares
			shares.Level = types.SharesLevelCustom
			shares.Shares = -1
			return allResourceFieldsValid(&spec.CpuAllocation)
		},
	}

	for i, test := range tests {
		ok := test()
		if ok {
			t.Errorf("%d: expected false", i)
		}
	}
}
