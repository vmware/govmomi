// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/types"
)

// newPropertyFilterTestVM starts a VPX model and returns its client, the default
// VM, and the simulator Context that mutates it.
func newPropertyFilterTestVM(t *testing.T) (*govmomi.Client, *VirtualMachine, *Context, func()) {
	t.Helper()
	goCtx := context.Background()

	m := VPX()
	require.NoError(t, m.Create())
	s := m.Service.NewServer()

	c, err := govmomi.NewClient(goCtx, s.URL, true)
	require.NoError(t, err)

	vm, err := find.NewFinder(c.Client).VirtualMachine(goCtx, "DC0_H0_VM0")
	require.NoError(t, err)

	simCtx := m.Service.Context
	obj := simCtx.Map.Get(vm.Reference()).(*VirtualMachine)

	return c, obj, simCtx, func() {
		s.Close()
		m.Remove()
	}
}

// TestPropertyFilter_NestedPathParentChange checks that a filter on a nested
// property reports a change that PropertyDiff, and so ctx.AutoUpdate, reports
// at its top-level parent.
func TestPropertyFilter_NestedPathParentChange(t *testing.T) {
	c, vm, simCtx, done := newPropertyFilterTestVM(t)
	defer done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		time.Sleep(200 * time.Millisecond)
		simCtx.AutoUpdate(vm, func() {
			if vm.Guest == nil {
				vm.Guest = &types.GuestInfo{}
			}
			vm.Guest.HostName = "nested-change"
		})
	}()

	var got any
	err := property.Wait(ctx, property.DefaultCollector(c.Client), vm.Reference(), []string{"guest.hostName"}, func(changes []types.PropertyChange) bool {
		for _, c := range changes {
			require.Equal(t, "guest.hostName", c.Name)
			got = c.Val
		}
		return got == "nested-change"
	})
	require.NoError(t, err, "got %v", got)
}

// TestPropertyFilter_MultiPathParentChange checks that a change to a parent
// property, as PropertyDiff reports it, is reported for every requested path
// under that parent, not just the first one listed. This is the guest NIC
// report: a watcher of guest.net and guest.ipAddress must see both.
func TestPropertyFilter_MultiPathParentChange(t *testing.T) {
	c, vm, simCtx, done := newPropertyFilterTestVM(t)
	defer done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const ip = "10.0.2.102"

	go func() {
		time.Sleep(200 * time.Millisecond)
		simCtx.AutoUpdate(vm, func() {
			if vm.Guest == nil {
				vm.Guest = &types.GuestInfo{}
			}
			vm.Guest.Net = []types.GuestNicInfo{{Network: "VM Network", IpAddress: []string{ip}}}
			vm.Guest.IpAddress = ip
		})
	}()

	got := make(map[string]any)
	err := property.Wait(ctx, property.DefaultCollector(c.Client), vm.Reference(), []string{"guest.net", "guest.ipAddress"}, func(changes []types.PropertyChange) bool {
		for _, c := range changes {
			got[c.Name] = c.Val
		}
		nics, _ := got["guest.net"].(types.ArrayOfGuestNicInfo)
		return got["guest.ipAddress"] == ip && len(nics.GuestNicInfo) == 1 && len(nics.GuestNicInfo[0].IpAddress) == 1 && nics.GuestNicInfo[0].IpAddress[0] == ip
	})
	require.NoError(t, err, "got %v", got)
}
