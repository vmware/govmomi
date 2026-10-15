// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

// QueryConfigOptionEx fills in GuestOSDescriptor on the response only,
// it must not modify what QueryConfigOption returns.
func TestEnvironmentBrowserQueryConfigOptionEx(t *testing.T) {
	tests := []struct {
		name   string
		stored *types.VirtualMachineConfigOption // as set when the browser has a stored config option
	}{
		{"default", nil},
		{"stored", &types.VirtualMachineConfigOption{
			Version:       esx.HardwareVersion,
			DefaultDevice: esx.VirtualDevice,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := VPX()
			if err := m.Create(); err != nil {
				t.Fatal(err)
			}
			defer m.Remove()

			ctx := context.Background()
			c := m.Service.client()

			vm := m.Service.Context.Map.Any("VirtualMachine").(*VirtualMachine)
			ref := vm.EnvironmentBrowser
			eb := m.Service.Context.Map.Get(ref).(*EnvironmentBrowser)
			eb.QueryConfigOptionResponse.Returnval = test.stored

			count := func() int {
				res, err := methods.QueryConfigOption(ctx, c, &types.QueryConfigOption{This: ref})
				if err != nil {
					t.Fatal(err)
				}
				return len(res.Returnval.GuestOSDescriptor)
			}

			if n := count(); n != 0 {
				t.Fatalf("QueryConfigOption: expected 0 guestOSDescriptor, got %d", n)
			}

			for _, spec := range []*types.EnvironmentBrowserConfigOptionQuerySpec{
				{},
				{GuestId: []string{string(GuestID[0])}},
			} {
				res, err := methods.QueryConfigOptionEx(ctx, c, &types.QueryConfigOptionEx{This: ref, Spec: spec})
				if err != nil {
					t.Fatal(err)
				}
				expect := len(GuestID)
				if len(spec.GuestId) != 0 {
					expect = 1
				}
				if n := len(res.Returnval.GuestOSDescriptor); n != expect {
					t.Errorf("QueryConfigOptionEx(%v): expected %d guestOSDescriptor, got %d", spec.GuestId, expect, n)
				}

				if n := count(); n != 0 {
					t.Errorf("QueryConfigOption after QueryConfigOptionEx(%v): expected 0 guestOSDescriptor, got %d", spec.GuestId, n)
				}
			}
		})
	}
}
