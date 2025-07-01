// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestVirtualMachineWaitForIP(t *testing.T) {
	m := simulator.VPX()
	err := m.Run(func(ctx context.Context, c *vim25.Client) error {
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}

		reconfig := func(ip string) error {
			task, err := vm.Reconfigure(ctx, types.VirtualMachineConfigSpec{
				ExtraConfig: []types.BaseOptionValue{
					&types.OptionValue{Key: "SET.guest.ipAddress", Value: ip},
				},
			})
			if err != nil {
				return err
			}
			return task.Wait(ctx)
		}

		if err := reconfig("fe80::250:56ff:fe97:2458"); err != nil {
			return err
		}

		ip, err := vm.WaitForIP(ctx)
		if err != nil {
			return err
		}

		if net.ParseIP(ip).To4() != nil {
			t.Errorf("expected v6 ip, but %q is v4", ip)
		}

		delay := time.Second / 2
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Logf("delaying map update for %v", delay)
			time.Sleep(delay)
			if err := reconfig("10.0.0.1"); err != nil {
				t.Logf("reconfig error: %s", err)
			}
		}()

		ip, err = vm.WaitForIP(ctx, true)
		if err != nil {
			t.Fatal(err)
		}

		if net.ParseIP(ip).To4() == nil {
			t.Errorf("expected v4 ip, but %q is v6", ip)
		}

		wg.Wait()
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
