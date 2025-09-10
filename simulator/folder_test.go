// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func addStandaloneHostTask(folder *object.Folder, spec types.HostConnectSpec) (*object.Task, error) {
	// TODO: add govmomi wrapper
	req := types.AddStandaloneHost_Task{
		This:         folder.Reference(),
		Spec:         spec,
		AddConnected: true,
	}

	res, err := methods.AddStandaloneHost_Task(context.TODO(), folder.Client(), &req)
	if err != nil {
		return nil, err
	}

	task := object.NewTask(folder.Client(), res.Returnval)
	return task, nil
}

func TestFolderESX(t *testing.T) {
	content := esx.ServiceContent
	s := New(NewServiceInstance(NewContext(), content, esx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	f := object.NewRootFolder(c.Client)

	_, err = f.CreateFolder(ctx, "foo")
	if err == nil {
		t.Error("expected error")
	}

	_, err = f.CreateDatacenter(ctx, "foo")
	if err == nil {
		t.Fatal("expected error")
	}

	finder := find.NewFinder(c.Client, false)
	dc, err := finder.DatacenterOrDefault(ctx, "")
	if err != nil {
		t.Fatal(err)
	}

	folders, err := dc.Folders(ctx)
	if err != nil {
		t.Fatal(err)
	}

	spec := types.HostConnectSpec{}
	_, err = addStandaloneHostTask(folders.HostFolder, spec)
	if err == nil {
		t.Fatal("expected error")
	}

	_, err = folders.DatastoreFolder.CreateStoragePod(ctx, "pod")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFolderVC(t *testing.T) {
	content := vpx.ServiceContent
	ctx := NewContext()
	s := New(NewServiceInstance(ctx, content, vpx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	f := object.NewRootFolder(c.Client)

	ff, err := f.CreateFolder(ctx, "foo")
	if err != nil {
		t.Error(err)
	}

	_, err = f.CreateFolder(ctx, "foo")
	if err == nil {
		t.Error("expected error")
	}

	var dup *types.DuplicateName
	_, ok := fault.As(err, &dup)
	if !ok {
		t.Fatal("expected DuplicateName type")
	}
	if dup.Object != ff.Reference() {
		t.Fatal("Duplicate object not matched")
	}

	dc, err := f.CreateDatacenter(ctx, "bar")
	if err != nil {
		t.Error(err)
	}

	for _, ref := range []object.Reference{ff, dc} {
		o := ctx.Map.Get(ref.Reference())
		if o == nil {
			t.Fatalf("failed to find %#v", ref)
		}

		e := o.(mo.Entity).Entity()
		if *e.Parent != f.Reference() {
			t.Fail()
		}
	}

	dc, err = ff.CreateDatacenter(ctx, "biz")
	if err != nil {
		t.Error(err)
	}

	folders, err := dc.Folders(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = folders.VmFolder.CreateStoragePod(ctx, "pod")
	if err == nil {
		t.Error("expected error")
	}

	pod, err := folders.DatastoreFolder.CreateStoragePod(ctx, "pod")
	if err != nil {
		t.Error(err)
	}

	_, err = folders.DatastoreFolder.CreateStoragePod(ctx, "pod")
	if err == nil {
		t.Error("expected error")
	}
	_, ok = fault.As(err, &dup)
	if !ok {
		t.Fatal("expected DuplicateName type")
	}
	if dup.Object != pod.Reference() {
		t.Fatal("Duplicate object not matched")
	}

	tests := []struct {
		name  string
		state types.TaskInfoState
	}{
		{"", types.TaskInfoStateError},
		{"foo.local", types.TaskInfoStateSuccess},
	}

	for _, test := range tests {
		spec := types.HostConnectSpec{
			HostName: test.name,
		}

		task, err := addStandaloneHostTask(folders.HostFolder, spec)
		if err != nil {
			t.Fatal(err)
		}

		res, err := task.WaitForResult(ctx, nil)
		if test.state == types.TaskInfoStateError {
			if err == nil {
				t.Error("expected error")
			}

			if res.Result != nil {
				t.Error("expected nil")
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}

			ref, ok := res.Result.(types.ManagedObjectReference)
			if !ok {
				t.Errorf("expected moref, got type=%T", res.Result)
			}
			host := ctx.Map.Get(ref).(*HostSystem)
			if host.Name != test.name {
				t.Fail()
			}

			if ref == esx.HostSystem.Self {
				t.Error("expected new host Self reference")
			}
			if *host.Summary.Host == esx.HostSystem.Self {
				t.Error("expected new host summary Self reference")
			}

			pool := ctx.Map.Get(*host.Parent).(*mo.ComputeResource).ResourcePool
			if *pool == esx.ResourcePool.Self {
				t.Error("expected new pool Self reference")
			}
		}

		if res.State != test.state {
			t.Fatalf("%s", res.State)
		}
	}
}

func TestFolderSpecialCharaters(t *testing.T) {
	content := vpx.ServiceContent
	ctx := NewContext()
	s := New(NewServiceInstance(ctx, content, vpx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	f := object.NewRootFolder(c.Client)

	tests := []struct {
		name     string
		expected string
	}{
		{`/`, `%2f`},
		{`\`, `%5c`},
		{`%`, `%25`},
		// multiple special characters
		{`%%`, `%25%25`},
	}

	for _, test := range tests {
		ff, err := f.CreateFolder(ctx, test.name)
		if err != nil {
			t.Fatal(err)
		}

		o := ctx.Map.Get(ff.Reference())
		if o == nil {
			t.Fatalf("failed to find %#v", ff)
		}

		e := o.(mo.Entity).Entity()
		if e.Name != test.expected {
			t.Errorf("expected %s, got %s", test.expected, e.Name)
		}
	}
}

func TestFolderFaults(t *testing.T) {
	f := Folder{}
	f.ChildType = []string{"VirtualMachine"}

	if f.CreateFolder(nil, nil).Fault() == nil {
		t.Error("expected fault")
	}

	if f.CreateDatacenter(nil, nil).Fault() == nil {
		t.Error("expected fault")
	}
}

func TestRegisterVm(t *testing.T) {
	for i, model := range []*Model{ESX(), VPX()} {
		match := "*"
		if i == 1 {
			model.App = 1
			match = "*APP*"
		}
		defer model.Remove()
		err := model.Create()
		if err != nil {
			t.Fatal(err)
		}

		s := model.Service.NewServer()
		defer s.Close()

		ctx := model.Service.Context

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

		folders, err := dc.Folders(ctx)
		if err != nil {
			t.Fatal(err)
		}

		vmFolder := folders.VmFolder

		vms, err := finder.VirtualMachineList(ctx, match)
		if err != nil {
			t.Fatal(err)
		}

		vm := ctx.Map.Get(vms[0].Reference()).(*VirtualMachine)

		req := types.RegisterVM_Task{
			This:       vmFolder.Reference(),
			AsTemplate: true,
		}

		steps := []struct {
			e any
			f func()
		}{
			{
				new(types.InvalidArgument), func() { req.AsTemplate = false },
			},
			{
				new(types.InvalidArgument), func() { req.Pool = vm.ResourcePool },
			},
			{
				new(types.InvalidArgument), func() { req.Path = "enoent" },
			},
			{
				new(types.InvalidDatastorePath), func() { req.Path = vm.Config.Files.VmPathName + "-enoent" },
			},
			{
				new(types.NotFound), func() { req.Path = vm.Config.Files.VmPathName },
			},
			{
				new(types.AlreadyExists), func() { ctx.Map.Remove(ctx, vm.Reference()) },
			},
			{
				nil, func() {},
			},
		}

		for _, step := range steps {
			res, err := methods.RegisterVM_Task(ctx, c.Client, &req)
			if err != nil {
				t.Fatal(err)
			}

			ct := object.NewTask(c.Client, res.Returnval)
			_ = ct.Wait(ctx)

			rt := ctx.Map.Get(res.Returnval).(*Task)

			if step.e != nil {
				fault := rt.Info.Error.Fault
				if reflect.TypeOf(fault) != reflect.TypeOf(step.e) {
					t.Errorf("%T != %T", fault, step.e)
				}
			} else {
				if rt.Info.Error != nil {
					t.Errorf("unexpected error: %#v", rt.Info.Error)
				}
			}

			step.f()
		}

		nvm, err := finder.VirtualMachine(ctx, vm.Name)
		if err != nil {
			t.Fatal(err)
		}

		if nvm.Reference() == vm.Reference() {
			t.Error("expected new moref")
		}

		onTask, _ := nvm.PowerOn(ctx)
		_ = onTask.Wait(ctx)

		steps = []struct {
			e any
			f func()
		}{
			{
				types.InvalidPowerState{}, func() { offTask, _ := nvm.PowerOff(ctx); _ = offTask.Wait(ctx) },
			},
			{
				nil, func() {},
			},
			{
				types.ManagedObjectNotFound{}, func() {},
			},
		}

		for _, step := range steps {
			err = nvm.Unregister(ctx)

			if step.e != nil {
				fault := soap.ToSoapFault(err).VimFault()
				if reflect.TypeOf(fault) != reflect.TypeOf(step.e) {
					t.Errorf("%T != %T", fault, step.e)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %#v", err)
				}
			}

			step.f()
		}
	}
}

func TestFolderMoveInto(t *testing.T) {
	ctx := context.Background()
	model := VPX()
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

	folders, err := dc.Folders(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ds, err := finder.DefaultDatastore(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Move Datastore into a vm folder should fail
	task, err := folders.VmFolder.MoveInto(ctx, []types.ManagedObjectReference{ds.Reference()})
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err == nil {
		t.Errorf("expected error")
	}

	// Move Datacenter into a sub folder should pass
	f, err := object.NewRootFolder(c.Client).CreateFolder(ctx, "foo")
	if err != nil {
		t.Error(err)
	}

	task, _ = f.MoveInto(ctx, []types.ManagedObjectReference{dc.Reference()})
	err = task.Wait(ctx)
	if err != nil {
		t.Error(err)
	}

	pod, err := folders.DatastoreFolder.CreateStoragePod(ctx, "pod")
	if err != nil {
		t.Error(err)
	}

	// Moving any type other than Datastore into a StoragePod should fail
	task, _ = pod.MoveInto(ctx, []types.ManagedObjectReference{dc.Reference()})
	err = task.Wait(ctx)
	if err == nil {
		t.Error("expected error")
	}

	// Move DS into a StoragePod
	task, _ = pod.MoveInto(ctx, []types.ManagedObjectReference{ds.Reference()})
	err = task.Wait(ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestFolderCreateDVS(t *testing.T) {
	ctx := context.Background()
	model := VPX()
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

	folders, err := dc.Folders(ctx)
	if err != nil {
		t.Fatal(err)
	}

	var spec types.DVSCreateSpec
	spec.ConfigSpec = &types.VMwareDVSConfigSpec{}
	spec.ConfigSpec.GetDVSConfigSpec().Name = "foo"

	task, err := folders.NetworkFolder.CreateDVS(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Error(err)
	}

	net, err := finder.Network(ctx, "foo")
	if err != nil {
		t.Error(err)
	}

	dvs, ok := net.(*object.DistributedVirtualSwitch)
	if !ok {
		t.Fatalf("%T is not of type %T", net, dvs)
	}

	task, err = folders.NetworkFolder.CreateDVS(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err == nil {
		t.Error("expected error")
	}

	pspec := types.DVPortgroupConfigSpec{Name: "xnet"}
	task, err = dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{pspec})
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Error(err)
	}

	net, err = finder.Network(ctx, "xnet")
	if err != nil {
		t.Error(err)
	}

	pg, ok := net.(*object.DistributedVirtualPortgroup)
	if !ok {
		t.Fatalf("%T is not of type %T", net, pg)
	}

	backing, err := net.EthernetCardBackingInfo(ctx)
	if err != nil {
		t.Fatal(err)
	}

	info, ok := backing.(*types.VirtualEthernetCardDistributedVirtualPortBackingInfo)
	if ok {
		if info.Port.SwitchUuid == "" || info.Port.PortgroupKey == "" {
			t.Errorf("invalid port: %#v", info.Port)
		}
	} else {
		t.Fatalf("%T is not of type %T", net, info)
	}

	task, err = dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{pspec})
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err == nil {
		t.Error("expected error")
	}
}

func TestPlaceVmsXClusterCreateAndPowerOn(t *testing.T) {
	vpx := VPX()
	vpx.Cluster = 3

	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, false)
		spec := types.PlaceVmsXClusterSpec{}

		pools, err := finder.ResourcePoolList(ctx, "/DC0/host/DC0_C*/*")
		if err != nil {
			t.Fatal(err)
		}

		for _, pool := range pools {
			spec.ResourcePools = append(spec.ResourcePools, pool.Reference())
		}

		spec.VmPlacementSpecs = []types.PlaceVmsXClusterSpecVmPlacementSpec{{
			ConfigSpec: types.VirtualMachineConfigSpec{
				Name: "test-vm",
			},
		}}

		folder := object.NewRootFolder(c)
		res, err := folder.PlaceVmsXCluster(ctx, spec)
		if err != nil {
			t.Fatal(err)
		}

		if len(res.PlacementInfos) != len(spec.VmPlacementSpecs) {
			t.Errorf("%d PlacementInfos vs %d VmPlacementSpecs", len(res.PlacementInfos), len(spec.VmPlacementSpecs))
		}
	}, vpx)
}

func TestPlaceVmsXClusterCreateAndPowerOnWithCandidateNetworks(t *testing.T) {
	vpx := VPX()
	vpx.Cluster = 3

	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, false)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("failed to get default datacenter: %v", err)
		}
		finder.SetDatacenter(datacenter)

		netA, err := finder.Network(ctx, "VM Network")
		if err != nil {
			t.Fatalf("unexpected error while getting network reference: %v", err)
		}
		netB, err := finder.Network(ctx, "DC0_DVPG0")
		if err != nil {
			t.Fatalf("unexpected error while getting network reference: %v", err)
		}

		candidateNetworks := []types.PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks{
			{
				Networks: []types.ManagedObjectReference{
					netA.Reference(), netB.Reference(),
				},
			},
			{
				Networks: []types.ManagedObjectReference{
					netA.Reference(),
				},
			},
		}
		spec := types.PlaceVmsXClusterSpec{}

		pools, err := finder.ResourcePoolList(ctx, "/DC0/host/DC0_C*/*")
		if err != nil {
			t.Fatal(err)
		}

		for _, pool := range pools {
			spec.ResourcePools = append(spec.ResourcePools, pool.Reference())
		}

		spec.VmPlacementSpecs = []types.PlaceVmsXClusterSpecVmPlacementSpec{{
			ConfigSpec: types.VirtualMachineConfigSpec{
				Name: "test-vm",
			},
			CandidateNetworks: candidateNetworks,
		}}

		folder := object.NewRootFolder(c)
		res, err := folder.PlaceVmsXCluster(ctx, spec)
		if err != nil {
			t.Fatal(err)
		}

		if len(res.PlacementInfos) != len(spec.VmPlacementSpecs) {
			t.Errorf("%d PlacementInfos vs %d VmPlacementSpecs", len(res.PlacementInfos), len(spec.VmPlacementSpecs))
		}

		// Validate AvailableNetworks returned in placement recommendations.
		for _, pinfo := range res.PlacementInfos {
			for i, action := range pinfo.Recommendation.Action {
				// Ensure the action is of expected extended type.
				initPlaceAction, ok := action.(*types.ClusterClusterInitialPlacementAction)
				if !ok {
					t.Errorf("Action[%d] is not ClusterClusterInitialPlacementActionEx, got %T", i, action)
					continue
				}
				if len(initPlaceAction.AvailableNetworks) == 0 {
					t.Errorf("AvailableNetworks is empty for VM %v", pinfo.Vm)
				} else {
					t.Logf("AvailableNetworks for VM %v:", pinfo.Vm)
					for _, net := range initPlaceAction.AvailableNetworks {
						t.Logf("- %s", net.Value)
					}
				}
				// Define the expected networks that should be present in AvailableNetworks.
				expected := map[string]bool{
					netA.Reference().Value: true,
					netB.Reference().Value: true,
				}

				// Verify that all returned networks are part of the expected set.
				for _, actual := range initPlaceAction.AvailableNetworks {
					if !expected[actual.Value] {
						t.Errorf("unexpected network in availableNetworks: %s", actual.Value)
					}
				}
			}
		}
	}, vpx)
}

func TestPlaceVmsXClusterRelocate(t *testing.T) {
	vpx := VPX()
	vpx.Cluster = 3

	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("failed to get default datacenter: %v", err)
		}
		finder.SetDatacenter(datacenter)

		pools, err := finder.ResourcePoolList(ctx, "/DC0/host/DC0_C*/*")
		if err != nil {
			t.Fatal(err)
		}

		var poolMoRefs []types.ManagedObjectReference
		for _, pool := range pools {
			poolMoRefs = append(poolMoRefs, pool.Reference())
		}

		vmMoRef := Map(ctx).Any("VirtualMachine").(*VirtualMachine).Reference()

		netA, err := finder.Network(context.Background(), "VM Network")
		if err != nil {
			t.Fatalf("unexpected error while getting network reference: %v", err)
		}
		netB, err := finder.Network(context.Background(), "DC0_DVPG0")
		if err != nil {
			t.Fatalf("unexpected error while getting network reference: %v", err)
		}

		candidateNetworks := []types.PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks{
			{Networks: []types.ManagedObjectReference{netA.Reference(), netB.Reference()}}, // NIC 0
			{Networks: []types.ManagedObjectReference{netA.Reference()}},                   // NIC 1
		}
		cfgSpec := types.VirtualMachineConfigSpec{}

		tests := []struct {
			name         string
			poolMoRefs   []types.ManagedObjectReference
			configSpec   types.VirtualMachineConfigSpec
			relocateSpec *types.VirtualMachineRelocateSpec
			vmMoRef      *types.ManagedObjectReference
			expectedErr  string
		}{
			{
				"relocate without any resource pools",
				nil,
				cfgSpec,
				&types.VirtualMachineRelocateSpec{},
				&vmMoRef,
				"InvalidArgument",
			},
			{
				"relocate without a relocate spec",
				poolMoRefs,
				cfgSpec,
				nil,
				&vmMoRef,
				"InvalidArgument",
			},
			{
				"relocate without a vm in the placement spec",
				poolMoRefs,
				cfgSpec,
				&types.VirtualMachineRelocateSpec{},
				nil,
				"InvalidArgument",
			},
			{
				"relocate with a non-existing vm in the placement spec",
				poolMoRefs,
				cfgSpec,
				&types.VirtualMachineRelocateSpec{},
				&types.ManagedObjectReference{
					Type:  "VirtualMachine",
					Value: "fake-vm-999",
				},
				"InvalidArgument",
			},
			{
				"relocate with an empty relocate spec",
				poolMoRefs,
				cfgSpec,
				&types.VirtualMachineRelocateSpec{},
				&vmMoRef,
				"",
			},
		}

		for testNo, test := range tests {
			test := test // assign to local var since loop var is reused

			truebool := true

			placeVmsXClusterSpec := types.PlaceVmsXClusterSpec{
				ResourcePools:           test.poolMoRefs,
				PlacementType:           string(types.PlaceVmsXClusterSpecPlacementTypeRelocate),
				HostRecommRequired:      &truebool,
				DatastoreRecommRequired: &truebool,
			}

			placeVmsXClusterSpec.VmPlacementSpecs = []types.PlaceVmsXClusterSpecVmPlacementSpec{{
				ConfigSpec:        test.configSpec,
				Vm:                test.vmMoRef,
				RelocateSpec:      test.relocateSpec,
				CandidateNetworks: candidateNetworks,
			}}

			folder := object.NewRootFolder(c)
			res, err := folder.PlaceVmsXCluster(ctx, placeVmsXClusterSpec)

			if err == nil && test.expectedErr != "" {
				t.Fatalf("Test %v: expected error %q, received nil", testNo, test.expectedErr)
			} else if err != nil &&
				(test.expectedErr == "" || !strings.Contains(err.Error(), test.expectedErr)) {
				t.Fatalf("Test %v: expected error %q, received %v", testNo, test.expectedErr, err)
			}

			if err == nil {
				if len(res.PlacementInfos) != len(placeVmsXClusterSpec.VmPlacementSpecs) {
					t.Errorf("Test %v: %d PlacementInfos vs %d VmPlacementSpecs", testNo, len(res.PlacementInfos), len(placeVmsXClusterSpec.VmPlacementSpecs))
				}

				// Validate AvailableNetworks returned in placement recommendations.
				for _, pinfo := range res.PlacementInfos {
					for _, action := range pinfo.Recommendation.Action {
						relocateAction, ok := action.(*types.ClusterClusterRelocatePlacementAction)
						if !ok {
							t.Errorf("Test %v: received wrong action type in recommendation", testNo)
							continue
						}
						if relocateAction.TargetHost == nil {
							t.Errorf("Test %v: received nil host recommendation", testNo)
						}
						// Check if AvailableNetworks field is populated.
						if len(relocateAction.AvailableNetworks) == 0 {
							t.Errorf("AvailableNetworks is empty for VM %v", pinfo.Vm)
						} else {
							t.Logf("AvailableNetworks for VM %v:", pinfo.Vm)
							for _, net := range relocateAction.AvailableNetworks {
								t.Logf("- %s", net.Value)
							}
						}
						// Define the expected networks that should be present in AvailableNetworks.
						expected := map[string]bool{
							netA.Reference().Value: true,
							netB.Reference().Value: true,
						}
						// Verify that all returned networks are part of the expected set.
						for _, actual := range relocateAction.AvailableNetworks {
							if !expected[actual.Value] {
								t.Errorf("unexpected network in availableNetworks: %s", actual.Value)
							}
						}
					}
				}
			}
		}
	}, vpx)
}

func TestPlaceVmsXClusterReconfigure(t *testing.T) {
	vpx := VPX()
	// All hosts are cluster hosts
	vpx.Host = 0

	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, true)
		datacenter, err := finder.DefaultDatacenter(ctx)
		if err != nil {
			t.Fatalf("failed to get default datacenter: %v", err)
		}
		finder.SetDatacenter(datacenter)

		vm := Map(ctx).Any("VirtualMachine").(*VirtualMachine)
		host := Map(ctx).Get(vm.Runtime.Host.Reference()).(*HostSystem)
		cluster := Map(ctx).Get(*host.Parent).(*ClusterComputeResource)
		pool := Map(ctx).Get(*cluster.ResourcePool).(*ResourcePool)

		var poolMoRefs []types.ManagedObjectReference
		poolMoRefs = append(poolMoRefs, pool.Reference())

		cfgSpec := types.VirtualMachineConfigSpec{}

		tests := []struct {
			name         string
			poolMoRefs   []types.ManagedObjectReference
			configSpec   types.VirtualMachineConfigSpec
			relocateSpec *types.VirtualMachineRelocateSpec
			vmMoRef      *types.ManagedObjectReference
			expectedErr  string
		}{
			{
				"reconfigure without any resource pools",
				nil,
				cfgSpec,
				nil,
				&vm.Self,
				"InvalidArgument",
			},
			{
				"reconfigure with a relocate spec",
				poolMoRefs,
				cfgSpec,
				&types.VirtualMachineRelocateSpec{},
				&vm.Self,
				"InvalidArgument",
			},
			{
				"reconfigure without a vm in the placement spec",
				poolMoRefs,
				cfgSpec,
				nil,
				nil,
				"InvalidArgument",
			},
			{
				"reconfigure with a non-existing vm in the placement spec",
				poolMoRefs,
				cfgSpec,
				nil,
				&types.ManagedObjectReference{
					Type:  "VirtualMachine",
					Value: "fake-vm-999",
				},
				"InvalidArgument",
			},
			{
				"reconfigure with an empty config spec",
				poolMoRefs,
				cfgSpec,
				nil,
				&vm.Self,
				"",
			},
		}

		for testNo, test := range tests {
			test := test // assign to local var since loop var is reused

			placeVmsXClusterSpec := types.PlaceVmsXClusterSpec{
				ResourcePools:           test.poolMoRefs,
				PlacementType:           string(types.PlaceVmsXClusterSpecPlacementTypeReconfigure),
				HostRecommRequired:      types.NewBool(true),
				DatastoreRecommRequired: types.NewBool(true),
			}

			placeVmsXClusterSpec.VmPlacementSpecs = []types.PlaceVmsXClusterSpecVmPlacementSpec{{
				ConfigSpec:   test.configSpec,
				Vm:           test.vmMoRef,
				RelocateSpec: test.relocateSpec,
			}}

			folder := object.NewRootFolder(c)
			res, err := folder.PlaceVmsXCluster(ctx, placeVmsXClusterSpec)

			if err == nil && test.expectedErr != "" {
				t.Fatalf("Test %v: expected error %q, received nil", testNo, test.expectedErr)
			} else if err != nil &&
				(test.expectedErr == "" || !strings.Contains(err.Error(), test.expectedErr)) {
				t.Fatalf("Test %v: expected error %q, received %v", testNo, test.expectedErr, err)
			}

			if err == nil {
				if len(res.PlacementInfos) != len(placeVmsXClusterSpec.VmPlacementSpecs) {
					t.Errorf("%d PlacementInfos vs %d VmPlacementSpecs", len(res.PlacementInfos), len(placeVmsXClusterSpec.VmPlacementSpecs))
				}

				for _, pinfo := range res.PlacementInfos {
					for _, action := range pinfo.Recommendation.Action {
						if reconfigureAction, ok := action.(*types.ClusterClusterReconfigurePlacementAction); ok {
							if reconfigureAction.TargetHost == nil {
								t.Errorf("Test %v: received nil host recommendation", testNo)
							}
						} else {
							t.Errorf("Test %v: received wrong action type in recommendation", testNo)
						}
					}
				}
			}
		}
	}, vpx)
}

func TestPlaceVmsXClusterCreateAndPowerOnWithMultipleVms(t *testing.T) {
	vpx := VPX()
	vpx.Cluster = 3

	Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c, false)
		spec := types.PlaceVmsXClusterSpec{}

		pools, err := finder.ResourcePoolList(ctx, "/DC0/host/DC0_C*/*")
		if err != nil {
			t.Fatal(err)
		}

		for _, pool := range pools {
			spec.ResourcePools = append(spec.ResourcePools, pool.Reference())
		}

		spec.VmPlacementSpecs = []types.PlaceVmsXClusterSpecVmPlacementSpec{{
			ConfigSpec: types.VirtualMachineConfigSpec{
				Name: "test-vm0",
			},
		}, {
			ConfigSpec: types.VirtualMachineConfigSpec{
				Name: "test-vm1",
			},
		}}

		folder := object.NewRootFolder(c)
		res, err := folder.PlaceVmsXCluster(ctx, spec)
		if err != nil {
			t.Fatal(err)
		}

		if len(res.PlacementInfos) != 2 {
			t.Errorf("Expected 2 PlacementInfos. Received %d", len(res.PlacementInfos))
		}

		if len(res.PlacementInfos) != len(spec.VmPlacementSpecs) {
			t.Errorf("%d PlacementInfos vs %d VmPlacementSpecs", len(res.PlacementInfos), len(spec.VmPlacementSpecs))
		}
	}, vpx)
}
