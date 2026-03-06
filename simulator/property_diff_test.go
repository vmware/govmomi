// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestPropertyDiff_SimpleFields(t *testing.T) {
	// Create a simple folder object
	folder := &mo.Folder{
		ManagedEntity: mo.ManagedEntity{
			Name: "original-name",
		},
	}
	folder.Self = types.ManagedObjectReference{Type: "Folder", Value: "folder-1"}

	// Create checkpoint
	checkpoint := Checkpoint(folder)

	// Modify the folder
	folder.Name = "new-name"

	// Get the diff
	changes := PropertyDiff(checkpoint, folder)

	// Verify we got exactly one change
	if len(changes) != 1 {
		t.Fatalf("expected 1 change, got %d: %+v", len(changes), changes)
	}

	// Verify the change details
	change := changes[0]
	if change.Name != "name" {
		t.Errorf("expected change name 'name', got %q", change.Name)
	}
	if change.Op != types.PropertyChangeOpAssign {
		t.Errorf("expected Op Assign, got %v", change.Op)
	}
	if change.Val != "new-name" {
		t.Errorf("expected Val 'new-name', got %v", change.Val)
	}
}

func TestPropertyDiff_NestedFields(t *testing.T) {
	// Create a VM with guest info
	vm := &mo.VirtualMachine{
		Guest: &types.GuestInfo{
			IpAddress: "192.168.1.100",
			HostName:  "test-host",
		},
		Summary: types.VirtualMachineSummary{
			Guest: &types.VirtualMachineGuestSummary{
				IpAddress: "192.168.1.100",
				HostName:  "test-host",
			},
		},
	}
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}
	vm.Name = "test-vm"

	// Create checkpoint
	checkpoint := Checkpoint(vm)

	// Modify nested fields
	vm.Guest.IpAddress = "10.0.0.50"
	vm.Summary.Guest.IpAddress = "10.0.0.50"

	// Get the diff
	changes := PropertyDiff(checkpoint, vm)

	// We should have changes for guest and summary.guest
	if len(changes) < 2 {
		t.Fatalf("expected at least 2 changes, got %d: %+v", len(changes), changes)
	}

	// Check that we have the expected property paths
	foundGuest := false
	foundSummaryGuest := false
	for _, c := range changes {
		if c.Name == "guest" {
			foundGuest = true
		}
		if c.Name == "summary" {
			foundSummaryGuest = true
		}
	}

	if !foundGuest {
		t.Error("expected change for 'guest' property")
	}
	if !foundSummaryGuest {
		t.Error("expected change for 'summary' property")
	}
}

func TestPropertyDiff_AddRemove(t *testing.T) {
	// Test Add operation (nil -> value)
	vm := &mo.VirtualMachine{}
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}

	checkpoint := Checkpoint(vm)

	// Add a guest info
	vm.Guest = &types.GuestInfo{
		IpAddress: "192.168.1.100",
	}

	changes := PropertyDiff(checkpoint, vm)

	foundGuestAdd := false
	for _, c := range changes {
		if c.Name == "guest" && c.Op == types.PropertyChangeOpAdd {
			foundGuestAdd = true
		}
	}
	if !foundGuestAdd {
		t.Error("expected Add operation for 'guest' property")
	}

	// Test Remove operation (value -> nil)
	checkpoint2 := Checkpoint(vm)
	vm.Guest = nil

	changes2 := PropertyDiff(checkpoint2, vm)

	foundGuestRemove := false
	for _, c := range changes2 {
		if c.Name == "guest" && c.Op == types.PropertyChangeOpRemove {
			foundGuestRemove = true
		}
	}
	if !foundGuestRemove {
		t.Error("expected Remove operation for 'guest' property")
	}
}

func TestPropertyDiff_NoChanges(t *testing.T) {
	folder := &mo.Folder{
		ManagedEntity: mo.ManagedEntity{
			Name: "test-folder",
		},
	}
	folder.Self = types.ManagedObjectReference{Type: "Folder", Value: "folder-1"}

	checkpoint := Checkpoint(folder)

	// No modifications
	changes := PropertyDiff(checkpoint, folder)

	if len(changes) != 0 {
		t.Errorf("expected 0 changes for unmodified object, got %d: %+v", len(changes), changes)
	}
}

func TestPropertyDiff_SliceFields(t *testing.T) {
	vm := &mo.VirtualMachine{
		Guest: &types.GuestInfo{
			Net: []types.GuestNicInfo{
				{
					IpAddress:  []string{"192.168.1.100"},
					MacAddress: "00:50:56:aa:bb:cc",
				},
			},
		},
	}
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}

	checkpoint := Checkpoint(vm)

	// Modify the network info
	vm.Guest.Net[0].IpAddress = []string{"10.0.0.50", "10.0.0.51"}
	vm.Guest.Net[0].MacAddress = "00:50:56:dd:ee:ff"

	changes := PropertyDiff(checkpoint, vm)

	// Should detect change in guest
	foundGuest := false
	for _, c := range changes {
		if c.Name == "guest" {
			foundGuest = true
		}
	}
	if !foundGuest {
		t.Error("expected change for 'guest' property containing network changes")
	}
}

func TestCheckpoint(t *testing.T) {
	original := &mo.VirtualMachine{
		ManagedEntity: mo.ManagedEntity{
			Name: "original",
		},
		Guest: &types.GuestInfo{
			IpAddress: "192.168.1.1",
		},
	}
	original.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}

	// Create checkpoint
	checkpoint := Checkpoint(original)

	// Verify it's a different pointer
	if checkpoint == original {
		t.Error("checkpoint should be a different pointer")
	}

	// Verify values are equal
	if checkpoint.Name != original.Name {
		t.Errorf("checkpoint Name mismatch: %q vs %q", checkpoint.Name, original.Name)
	}
	if checkpoint.Guest.IpAddress != original.Guest.IpAddress {
		t.Errorf("checkpoint Guest.IpAddress mismatch")
	}

	// Modify original, checkpoint should be unchanged
	original.Name = "modified"
	original.Guest.IpAddress = "10.0.0.1"

	if checkpoint.Name != "original" {
		t.Error("checkpoint should not be affected by changes to original")
	}
	if checkpoint.Guest.IpAddress != "192.168.1.1" {
		t.Error("checkpoint Guest.IpAddress should not be affected by changes to original")
	}
}

// TestPropertyDiff_WithSimulator tests PropertyDiff in the context of a running simulator
func TestPropertyDiff_WithSimulator(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(c.Client)
	vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
	if err != nil {
		t.Fatal(err)
	}

	// Get the simulator's internal VM object
	simCtx := m.Service.Context
	ref := vm.Reference()
	obj := simCtx.Map.Get(ref).(*VirtualMachine)

	// Create a checkpoint of the VM state
	checkpoint := Checkpoint(&obj.VirtualMachine)

	// Modify the VM's guest info
	if obj.Guest == nil {
		obj.Guest = &types.GuestInfo{}
	}
	obj.Guest.IpAddress = "10.20.30.40"
	obj.Guest.HostName = "test-hostname"

	// Generate property changes
	changes := PropertyDiff(checkpoint, &obj.VirtualMachine)

	// Verify we got changes
	if len(changes) == 0 {
		t.Fatal("expected property changes after modifying VM")
	}

	// Apply the changes via Update
	simCtx.Update(obj, changes)

	// Now verify the changes are visible via the property collector
	pc := property.DefaultCollector(c.Client)
	var mvm mo.VirtualMachine
	err = pc.RetrieveOne(ctx, ref, []string{"guest.ipAddress", "guest.hostName"}, &mvm)
	if err != nil {
		t.Fatal(err)
	}

	if mvm.Guest == nil {
		t.Fatal("expected Guest to be set")
	}
	if mvm.Guest.IpAddress != "10.20.30.40" {
		t.Errorf("expected IpAddress '10.20.30.40', got %q", mvm.Guest.IpAddress)
	}
	if mvm.Guest.HostName != "test-hostname" {
		t.Errorf("expected HostName 'test-hostname', got %q", mvm.Guest.HostName)
	}
}

// TestPropertyDiff_MultipleChanges tests that PropertyDiff correctly handles multiple changes
func TestPropertyDiff_MultipleChanges(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(c.Client)
	vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
	if err != nil {
		t.Fatal(err)
	}

	// Get the simulator's internal VM object
	simCtx := m.Service.Context
	ref := vm.Reference()
	obj := simCtx.Map.Get(ref).(*VirtualMachine)

	// Create a checkpoint of the VM state
	checkpoint := Checkpoint(&obj.VirtualMachine)

	// Make multiple changes
	obj.Name = "renamed-vm"
	if obj.Guest == nil {
		obj.Guest = &types.GuestInfo{}
	}
	obj.Guest.IpAddress = "99.99.99.99"
	obj.Guest.HostName = "test-hostname"
	obj.Guest.Net = []types.GuestNicInfo{
		{
			IpAddress:  []string{"99.99.99.99", "fe80::1"},
			MacAddress: "00:50:56:aa:bb:cc",
		},
	}

	// Generate property changes
	changes := PropertyDiff(checkpoint, &obj.VirtualMachine)

	// Verify we got changes for name and guest
	foundName := false
	foundGuest := false
	for _, c := range changes {
		if c.Name == "name" {
			foundName = true
			if c.Val != "renamed-vm" {
				t.Errorf("expected name 'renamed-vm', got %v", c.Val)
			}
		}
		if c.Name == "guest" {
			foundGuest = true
			// The value can be either *types.GuestInfo or types.GuestInfo depending on wrapping
			var guestIP string
			var netLen int
			switch v := c.Val.(type) {
			case *types.GuestInfo:
				guestIP = v.IpAddress
				netLen = len(v.Net)
			case types.GuestInfo:
				guestIP = v.IpAddress
				netLen = len(v.Net)
			default:
				t.Errorf("expected GuestInfo type, got %T", c.Val)
				continue
			}
			if guestIP != "99.99.99.99" {
				t.Errorf("expected IpAddress '99.99.99.99', got %q", guestIP)
			}
			if netLen != 1 {
				t.Errorf("expected 1 NIC, got %d", netLen)
			}
		}
	}

	if !foundName {
		t.Error("expected change for 'name' property")
	}
	if !foundGuest {
		t.Error("expected change for 'guest' property")
	}

	// Apply changes
	simCtx.Update(obj, changes)

	// Verify changes are visible via property collector
	pc := property.DefaultCollector(c.Client)
	var mvm mo.VirtualMachine
	err = pc.RetrieveOne(ctx, ref, []string{"name", "guest"}, &mvm)
	if err != nil {
		t.Fatal(err)
	}

	if mvm.Name != "renamed-vm" {
		t.Errorf("expected Name 'renamed-vm', got %q", mvm.Name)
	}
	if mvm.Guest == nil {
		t.Fatal("expected Guest to be set")
	}
	if mvm.Guest.IpAddress != "99.99.99.99" {
		t.Errorf("expected IpAddress '99.99.99.99', got %q", mvm.Guest.IpAddress)
	}
	if len(mvm.Guest.Net) != 1 {
		t.Errorf("expected 1 NIC, got %d", len(mvm.Guest.Net))
	}
}

func TestDetermineChangeOp(t *testing.T) {
	tests := []struct {
		name     string
		oldVal   interface{}
		newVal   interface{}
		expected types.PropertyChangeOp
	}{
		{
			name:     "nil to value is Add",
			oldVal:   (*string)(nil),
			newVal:   stringPtr("hello"),
			expected: types.PropertyChangeOpAdd,
		},
		{
			name:     "value to nil is Remove",
			oldVal:   stringPtr("hello"),
			newVal:   (*string)(nil),
			expected: types.PropertyChangeOpRemove,
		},
		{
			name:     "value to value is Assign",
			oldVal:   stringPtr("hello"),
			newVal:   stringPtr("world"),
			expected: types.PropertyChangeOpAssign,
		},
		{
			name:     "empty string to non-empty is Add",
			oldVal:   "",
			newVal:   "hello",
			expected: types.PropertyChangeOpAdd,
		},
		{
			name:     "non-empty string to empty is Remove",
			oldVal:   "hello",
			newVal:   "",
			expected: types.PropertyChangeOpRemove,
		},
		{
			name:     "slice changes are Assign (slices are not considered empty)",
			oldVal:   []string{},
			newVal:   []string{"a", "b"},
			expected: types.PropertyChangeOpAssign,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldVal := reflect.ValueOf(tt.oldVal)
			newVal := reflect.ValueOf(tt.newVal)
			op := determineChangeOp(oldVal, newVal)
			if op != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, op)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

// TestPropertyDiff_RuntimePowerState tests that runtime.powerState changes are properly tracked
func TestPropertyDiff_RuntimePowerState(t *testing.T) {
	vm := &mo.VirtualMachine{
		Runtime: types.VirtualMachineRuntimeInfo{
			PowerState: types.VirtualMachinePowerStatePoweredOff,
		},
	}
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}

	checkpoint := Checkpoint(vm)

	// Change power state
	vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOn

	changes := PropertyDiff(checkpoint, vm)

	foundRuntime := false
	for _, c := range changes {
		if c.Name == "runtime" {
			foundRuntime = true
			runtime, ok := c.Val.(types.VirtualMachineRuntimeInfo)
			if !ok {
				t.Errorf("expected types.VirtualMachineRuntimeInfo, got %T", c.Val)
			} else if runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
				t.Errorf("expected PowerState PoweredOn, got %v", runtime.PowerState)
			}
		}
	}

	if !foundRuntime {
		t.Error("expected change for 'runtime' property")
	}
}

// TestPropertyDiff_GuestNetInfo tests that guest.net changes produce correct property changes
func TestPropertyDiff_GuestNetInfo(t *testing.T) {
	vm := &mo.VirtualMachine{
		Guest: &types.GuestInfo{
			Net: []types.GuestNicInfo{},
		},
	}
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}

	checkpoint := Checkpoint(vm)

	// Add network info (simulating container network detection)
	vm.Guest.Net = []types.GuestNicInfo{
		{
			Network:    "bridge",
			IpAddress:  []string{"172.17.0.2", "fe80::42:acff:fe11:2"},
			MacAddress: "02:42:ac:11:00:02",
			Connected:  true,
		},
	}
	vm.Guest.IpAddress = "172.17.0.2"
	vm.Guest.HostName = "container-hostname"

	changes := PropertyDiff(checkpoint, vm)

	if len(changes) == 0 {
		t.Fatal("expected property changes for guest network info")
	}

	foundGuest := false
	for _, c := range changes {
		if c.Name == "guest" {
			foundGuest = true
			// Verify the guest info contains network data
			var guestInfo *types.GuestInfo
			switch v := c.Val.(type) {
			case *types.GuestInfo:
				guestInfo = v
			case types.GuestInfo:
				guestInfo = &v
			}
			if guestInfo == nil {
				t.Errorf("expected GuestInfo, got %T", c.Val)
				continue
			}
			if len(guestInfo.Net) != 1 {
				t.Errorf("expected 1 NIC, got %d", len(guestInfo.Net))
			}
			if guestInfo.IpAddress != "172.17.0.2" {
				t.Errorf("expected IpAddress '172.17.0.2', got %q", guestInfo.IpAddress)
			}
		}
	}

	if !foundGuest {
		t.Error("expected change for 'guest' property")
	}
}

// TestPropertyDiff_SummaryGuest tests that summary.guest changes are tracked
func TestPropertyDiff_SummaryGuest(t *testing.T) {
	vm := &mo.VirtualMachine{
		Summary: types.VirtualMachineSummary{
			Guest: &types.VirtualMachineGuestSummary{},
		},
	}
	vm.Self = types.ManagedObjectReference{Type: "VirtualMachine", Value: "vm-1"}

	checkpoint := Checkpoint(vm)

	// Update summary guest info
	vm.Summary.Guest.IpAddress = "10.0.0.100"
	vm.Summary.Guest.HostName = "test-host"

	changes := PropertyDiff(checkpoint, vm)

	foundSummary := false
	for _, c := range changes {
		if c.Name == "summary" {
			foundSummary = true
		}
	}

	if !foundSummary {
		t.Error("expected change for 'summary' property")
	}
}

// TestContainerVMNetworkPropertyChanges tests that a container-backed VM produces
// the expected network property changes when powered on, including detailed Guest.Net info.
func TestContainerVMNetworkPropertyChanges(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		if !test.HasDocker() {
			t.Skip("requires docker on linux")
			return
		}

		finder := find.NewFinder(c)
		pool, err := finder.ResourcePool(ctx, "DC0_H0/Resources")
		if err != nil {
			t.Fatal(err)
		}
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			t.Fatal(err)
		}

		// Use busybox with a simple sleep command to keep the container running
		busybox := os.Getenv("VCSIM_BUSYBOX")
		if busybox == "" {
			busybox = "busybox"
		}

		// Use podman network if available for IP assignment
		network := os.Getenv("VCSIM_NETWORK")
		if network == "" {
			network = "podman"
		}

		spec := types.VirtualMachineConfigSpec{
			Name: "busybox-network-test",
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0] busybox-test",
			},
			ExtraConfig: []types.BaseOptionValue{
				&types.OptionValue{Key: ContainerBackingOptionKey, Value: busybox + " sleep 300"},
				&types.OptionValue{Key: "RUN.mountdmi", Value: "false"},
				&types.OptionValue{Key: "RUN.network", Value: network},
			},
		}

		f, err := dc.Folders(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Create a new VM
		task, err := f.VmFolder.CreateVM(ctx, spec, pool, nil)
		if err != nil {
			t.Fatal(err)
		}

		info, err := task.WaitForResult(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}

		vmRef := info.Result.(types.ManagedObjectReference)
		vm := object.NewVirtualMachine(c, vmRef)
		defer func() {
			task, _ = vm.PowerOff(ctx)
			_ = task.Wait(ctx)
			task, _ = vm.Destroy(ctx)
			_ = task.Wait(ctx)
		}()

		// Get initial state before power on
		pc := property.DefaultCollector(c)
		var initialVM mo.VirtualMachine
		err = pc.RetrieveOne(ctx, vmRef, []string{"runtime.powerState", "guest"}, &initialVM)
		if err != nil {
			t.Fatal(err)
		}

		if initialVM.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOff {
			t.Errorf("expected initial power state PoweredOff, got %v", initialVM.Runtime.PowerState)
		}

		// Verify Guest.Net is empty before power on
		if initialVM.Guest != nil && len(initialVM.Guest.Net) != 0 {
			t.Errorf("expected Guest.Net to be empty before power on, got %d entries", len(initialVM.Guest.Net))
		}

		// Power on the VM
		task, err = vm.PowerOn(ctx)
		if err != nil {
			t.Fatal(err)
		}
		err = task.Wait(ctx)
		if err != nil {
			t.Fatal(err)
		}

		// Wait for IP to be assigned with a timeout
		waitCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		ip, err := vm.WaitForIP(waitCtx, false)
		if err != nil {
			t.Logf("WaitForIP error (may be expected with rootless podman): %v", err)
		}

		// Retrieve the updated VM state with all guest properties
		var updatedVM mo.VirtualMachine
		err = pc.RetrieveOne(ctx, vmRef, []string{
			"runtime.powerState",
			"guest.ipAddress",
			"guest.hostName",
			"guest.net",
			"guest.ipStack",
			"summary.guest.ipAddress",
			"summary.guest.hostName",
		}, &updatedVM)
		if err != nil {
			t.Fatal(err)
		}

		// Verify power state changed
		if updatedVM.Runtime.PowerState != types.VirtualMachinePowerStatePoweredOn {
			t.Errorf("expected power state PoweredOn, got %v", updatedVM.Runtime.PowerState)
		}

		// If we got an IP, verify detailed guest properties
		if ip != "" {
			t.Logf("Container IP: %s", ip)

			if updatedVM.Guest == nil {
				t.Fatal("expected Guest to be populated")
			}

			// Verify Guest.IpAddress
			if updatedVM.Guest.IpAddress == "" {
				t.Error("expected Guest.IpAddress to be set")
			} else {
				t.Logf("Guest.IpAddress: %s", updatedVM.Guest.IpAddress)
			}

			// Verify Guest.HostName
			if updatedVM.Guest.HostName != "" {
				t.Logf("Guest.HostName: %s", updatedVM.Guest.HostName)
			}

			// Verify Guest.Net is now populated with detailed NIC info
			if len(updatedVM.Guest.Net) == 0 {
				t.Error("expected Guest.Net to be populated after power on")
			} else {
				t.Logf("Guest.Net has %d entries", len(updatedVM.Guest.Net))
				for i, nic := range updatedVM.Guest.Net {
					t.Logf("  NIC %d:", i)
					t.Logf("    Network: %s", nic.Network)
					t.Logf("    MacAddress: %s", nic.MacAddress)
					t.Logf("    IpAddress: %v", nic.IpAddress)
					t.Logf("    Connected: %v", nic.Connected)

					// Verify NIC has expected fields populated
					if nic.Network == "" {
						t.Errorf("NIC %d: expected Network to be set", i)
					}
					if len(nic.IpAddress) == 0 {
						t.Errorf("NIC %d: expected IpAddress to be set", i)
					}
					if nic.MacAddress == "" {
						t.Errorf("NIC %d: expected MacAddress to be set", i)
					}
					if !nic.Connected {
						t.Errorf("NIC %d: expected Connected to be true", i)
					}

					// Verify IpConfig is populated
					if nic.IpConfig != nil {
						t.Logf("    IpConfig.IpAddress: %+v", nic.IpConfig.IpAddress)
						if len(nic.IpConfig.IpAddress) == 0 {
							t.Errorf("NIC %d: expected IpConfig.IpAddress to be set", i)
						} else {
							ipAddr := nic.IpConfig.IpAddress[0]
							if ipAddr.IpAddress == "" {
								t.Errorf("NIC %d: expected IpConfig.IpAddress[0].IpAddress to be set", i)
							}
							if ipAddr.PrefixLength == 0 {
								t.Logf("NIC %d: PrefixLength is 0 (may be expected)", i)
							}
						}
					} else {
						t.Errorf("NIC %d: expected IpConfig to be set", i)
					}
				}
			}

			// Verify Guest.IpStack is populated
			if len(updatedVM.Guest.IpStack) == 0 {
				t.Error("expected Guest.IpStack to be populated")
			} else {
				t.Logf("Guest.IpStack has %d entries", len(updatedVM.Guest.IpStack))
				for i, stack := range updatedVM.Guest.IpStack {
					if stack.DnsConfig != nil {
						t.Logf("  Stack %d DnsConfig: HostName=%s, DomainName=%s, DNS=%v",
							i, stack.DnsConfig.HostName, stack.DnsConfig.DomainName, stack.DnsConfig.IpAddress)
					}
					if stack.IpRouteConfig != nil && len(stack.IpRouteConfig.IpRoute) > 0 {
						route := stack.IpRouteConfig.IpRoute[0]
						t.Logf("  Stack %d DefaultRoute: Gateway=%s", i, route.Gateway.IpAddress)
					}
				}
			}

			// Verify Summary.Guest
			if updatedVM.Summary.Guest != nil {
				if updatedVM.Summary.Guest.IpAddress != "" {
					t.Logf("Summary.Guest.IpAddress: %s", updatedVM.Summary.Guest.IpAddress)
				} else {
					t.Error("expected Summary.Guest.IpAddress to be set")
				}
				if updatedVM.Summary.Guest.HostName != "" {
					t.Logf("Summary.Guest.HostName: %s", updatedVM.Summary.Guest.HostName)
				}
			}
		} else {
			t.Log("No IP assigned (rootless podman without bridge network)")
		}
	})
}
