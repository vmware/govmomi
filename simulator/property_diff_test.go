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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.Len(t, changes, 1, "expected 1 change, got %d: %+v", len(changes), changes)

	// Verify the change details
	change := changes[0]
	assert.Equal(t, "name", change.Name, "expected change name 'name', got %q", change.Name)
	assert.Equal(t, types.PropertyChangeOpAssign, change.Op, "expected Op Assign, got %v", change.Op)
	assert.Equal(t, "new-name", change.Val, "expected Val 'new-name', got %v", change.Val)
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
	require.Len(t, changes, 2, "expected at least 2 changes, got %d: %+v", len(changes), changes)

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

	assert.True(t, foundGuest, "expected change for 'guest' property")
	assert.True(t, foundSummaryGuest, "expected change for 'summary' property")
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
	assert.True(t, foundGuestAdd, "expected Add operation for 'guest' property")

	// Test Remove operation (value -> nil)
	checkpoint2 := Checkpoint(vm)
	vm.Guest = nil

	changes2 := PropertyDiff(checkpoint2, vm)

	for _, c := range changes2 {
		if c.Name == "guest" && c.Op == types.PropertyChangeOpRemove {
			return
		}
	}
	t.Error("expected Remove operation for 'guest' property")
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

	require.Len(t, changes, 0, "expected 0 changes for unmodified object, got %d: %+v", len(changes), changes)
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
	for _, c := range changes {
		if c.Name == "guest" {
			return
		}
	}

	t.Error("expected change for 'guest' property containing network changes")
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
	require.False(t, checkpoint == original, "checkpoint should be a different pointer")

	// Verify values are equal
	assert.Equal(t, original.Name, checkpoint.Name, "checkpoint Name mismatch: %q vs %q", checkpoint.Name, original.Name)
	assert.Equal(t, original.Guest.IpAddress, checkpoint.Guest.IpAddress, "checkpoint Guest.IpAddress mismatch")

	// Modify original, checkpoint should be unchanged
	original.Name = "modified"
	original.Guest.IpAddress = "10.0.0.1"

	assert.Equal(t, "original", checkpoint.Name, "checkpoint should not be affected by changes to original")
	assert.Equal(t, "192.168.1.1", checkpoint.Guest.IpAddress, "checkpoint Guest.IpAddress should not be affected by changes to original")
}

// TestPropertyDiff_WithSimulator tests PropertyDiff in the context of a running simulator
func TestPropertyDiff_WithSimulator(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	defer m.Remove()

	err := m.Create()
	require.NoError(t, err, "expected no error creating simulator")

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	require.NoError(t, err, "expected no error creating client")

	finder := find.NewFinder(c.Client)
	vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
	require.NoError(t, err, "expected no error retrieving VM")

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
	require.Greater(t, len(changes), 0, "expected property changes after modifying VM")

	// Apply the changes via Update
	simCtx.Update(obj, changes)

	// Now verify the changes are visible via the property collector
	pc := property.DefaultCollector(c.Client)
	var mvm mo.VirtualMachine
	err = pc.RetrieveOne(ctx, ref, []string{"guest.ipAddress", "guest.hostName"}, &mvm)
	require.NoError(t, err, "expected no error retrieving VM properties")

	require.NotNil(t, mvm.Guest, "expected Guest to be set")
	assert.Equal(t, "10.20.30.40", mvm.Guest.IpAddress, "expected IpAddress '10.20.30.40', got %q", mvm.Guest.IpAddress)
	assert.Equal(t, "test-hostname", mvm.Guest.HostName, "expected HostName 'test-hostname', got %q", mvm.Guest.HostName)
}

// TestPropertyDiff_MultipleChanges tests that PropertyDiff correctly handles multiple changes
func TestPropertyDiff_MultipleChanges(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	defer m.Remove()

	err := m.Create()
	require.NoError(t, err, "expected no error creating simulator")

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	require.NoError(t, err, "expected no error creating client")

	finder := find.NewFinder(c.Client)
	vm, err := finder.VirtualMachine(ctx, "DC0_H0_VM0")
	require.NoError(t, err, "expected no error retrieving VM")

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
			assert.Equal(t, "renamed-vm", c.Val, "expected name 'renamed-vm', got %v", c.Val)
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
				assert.IsType(t, types.GuestInfo{}, c.Val, "expected GuestInfo type, got %T", c.Val)
				return
			}
			assert.Equal(t, "99.99.99.99", guestIP, "expected IpAddress '99.99.99.99', got %q", guestIP)
			assert.Equal(t, 1, netLen, "expected 1 NIC, got %d", netLen)
		}
	}

	assert.True(t, foundName, "expected change for 'name' property")
	assert.True(t, foundGuest, "expected change for 'guest' property")

	// Apply changes
	simCtx.Update(obj, changes)

	// Verify changes are visible via property collector
	pc := property.DefaultCollector(c.Client)
	var mvm mo.VirtualMachine
	err = pc.RetrieveOne(ctx, ref, []string{"name", "guest"}, &mvm)
	require.NoError(t, err, "expected no error retrieving VM properties")

	assert.Equal(t, "renamed-vm", mvm.Name, "expected Name 'renamed-vm', got %q", mvm.Name)
	require.NotNil(t, mvm.Guest, "expected Guest to be set")
	assert.Equal(t, "99.99.99.99", mvm.Guest.IpAddress, "expected IpAddress '99.99.99.99', got %q", mvm.Guest.IpAddress)
	assert.Equal(t, 1, len(mvm.Guest.Net), "expected 1 NIC, got %d", len(mvm.Guest.Net))
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
			assert.Equal(t, tt.expected, op, "expected %v, got %v", tt.expected, op)
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

	for _, c := range changes {
		if c.Name == "runtime" {
			runtime, ok := c.Val.(types.VirtualMachineRuntimeInfo)
			assert.True(t, ok, "expected types.VirtualMachineRuntimeInfo, got %T", c.Val)
			assert.Equal(t, types.VirtualMachinePowerStatePoweredOn, runtime.PowerState, "expected PowerState PoweredOn")
			return
		}
	}

	t.Error("expected change for 'runtime' property")
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
	require.Greater(t, len(changes), 0, "expected property changes for guest network info")

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

			if assert.NotNil(t, guestInfo, "expected GuestInfo, got %T", c.Val) {
				continue
			}

			require.Equal(t, 1, len(guestInfo.Net), "expected 1 NIC, got %d", len(guestInfo.Net))
			assert.Equal(t, "172.17.0.2", guestInfo.IpAddress, "expected IpAddress '172.17.0.2'")
		}
	}

	assert.True(t, foundGuest, "expected change for 'guest' property")
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

	for _, c := range changes {
		if c.Name == "summary" {
			return
		}
	}

	t.Error("expected change for 'summary' property")
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
		require.NoError(t, err, "expected no error retrieving resource pool")
		dc, err := finder.Datacenter(ctx, "DC0")
		require.NoError(t, err, "expected no error retrieving datacenter")

		// Use busybox with a simple sleep command to keep the container running
		busybox := os.Getenv("VCSIM_BUSYBOX")
		if busybox == "" {
			busybox = "busybox"
		}

		network := os.Getenv("VCSIM_NETWORK")
		if network == "" {
			// podman requires we specify a network to get an IP at all.
			// podman doesn't allow "bridge" as a network name, so we create a custom name
			network = "generic-bridge"
			_, err := createBridge(network)
			require.NoError(t, err, "expected no error creating bridge network")
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
		require.NoError(t, err, "expected no error retrieving folders")

		// Create a new VM
		task, err := f.VmFolder.CreateVM(ctx, spec, pool, nil)
		require.NoError(t, err, "expected no error creating VM")

		info, err := task.WaitForResult(ctx, nil)
		require.NoError(t, err, "expected no error waiting for task result")

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
		require.NoError(t, err, "expected no error retrieving initial VM state")
		assert.Equal(t, types.VirtualMachinePowerStatePoweredOff, initialVM.Runtime.PowerState, "expected initial power state PoweredOff")

		// Verify Guest.Net is empty before power on
		require.NotNil(t, initialVM.Guest, "expected Guest to be set")
		require.Equal(t, 0, len(initialVM.Guest.Net), "expected Guest.Net to be empty before power on")

		// Power on the VM
		task, err = vm.PowerOn(ctx)
		require.NoError(t, err, "expected no error powering on VM")

		err = task.Wait(ctx)
		require.NoError(t, err, "expected no error waiting for task result")

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
		require.NoError(t, err, "expected no error retrieving updated VM state")

		// Verify power state changed
		assert.Equal(t, types.VirtualMachinePowerStatePoweredOn, updatedVM.Runtime.PowerState, "expected power state PoweredOn")

		// If we got an IP, verify detailed guest properties
		if ip == "" {
			t.Log("No IP assigned (rootless podman without bridge network)")
			return
		}

		t.Logf("Container IP: %s", ip)

		require.NotNil(t, updatedVM.Guest, "expected Guest to be populated")

		// Verify Guest.IpAddress
		if assert.NotEqual(t, "", updatedVM.Guest.IpAddress, "expected Guest.IpAddress to be set") {
			t.Logf("Guest.IpAddress: %s", updatedVM.Guest.IpAddress)
		}

		// Verify Guest.HostName
		if assert.NotEqual(t, "", updatedVM.Guest.HostName, "expected Guest.HostName to be set") {
			t.Logf("Guest.HostName: %s", updatedVM.Guest.HostName)
		}

		// Verify Guest.Net is now populated with detailed NIC info
		if assert.Greater(t, len(updatedVM.Guest.Net), 0, "expected Guest.Net to be populated after power on") {
			t.Logf("Guest.Net has %d entries", len(updatedVM.Guest.Net))
			for i, nic := range updatedVM.Guest.Net {
				t.Logf("  NIC %d:", i)
				t.Logf("    Network: %s", nic.Network)
				t.Logf("    MacAddress: %s", nic.MacAddress)
				t.Logf("    IpAddress: %v", nic.IpAddress)
				t.Logf("    Connected: %v", nic.Connected)

				// Verify NIC has expected fields populated
				assert.NotEqual(t, "", nic.Network, "expected Network to be set")
				assert.Greater(t, len(nic.IpAddress), 0, "expected IpAddress to be set")
				assert.NotEqual(t, "", nic.MacAddress, "expected MacAddress to be set")
				assert.True(t, nic.Connected, "expected Connected to be true")

				// Verify IpConfig is populated
				if assert.NotNil(t, nic.IpConfig, "expected IpConfig to be set (nic %d)", i) {
					t.Logf("    IpConfig.IpAddress: %+v", nic.IpConfig.IpAddress)
					if assert.Greater(t, len(nic.IpConfig.IpAddress), 0, "expected IpConfig.IpAddress to be set") {
						ipAddr := nic.IpConfig.IpAddress[0]
						assert.NotEqual(t, "", ipAddr.IpAddress, "expected IpConfig.IpAddress[0].IpAddress to be set")
						assert.NotEqual(t, int32(0), ipAddr.PrefixLength, "expected PrefixLength to be set")
					}
				}
			}
		}

		// Verify Guest.IpStack is populated
		if assert.Greater(t, len(updatedVM.Guest.IpStack), 0, "expected Guest.IpStack to be populated") {
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
			if assert.NotEqual(t, "", updatedVM.Summary.Guest.IpAddress, "expected Summary.Guest.IpAddress to be set") {
				t.Logf("Summary.Guest.IpAddress: %s", updatedVM.Summary.Guest.IpAddress)
			}
			if updatedVM.Summary.Guest.HostName != "" {
				t.Logf("Summary.Guest.HostName: %s", updatedVM.Summary.Guest.HostName)
			}
		}

	})
}
