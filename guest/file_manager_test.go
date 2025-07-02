// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestTranferURL(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		vm := simulator.Map(ctx).Any("VirtualMachine").(*simulator.VirtualMachine)
		host := simulator.Map(ctx).Get(*vm.Runtime.Host).(*simulator.HostSystem)

		ops := guest.NewOperationsManager(c, vm.Reference())
		m, err := ops.FileManager(ctx)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 2; i++ { // 2nd time is to validate the cached value
			turl := "https://esx:443/foo/bar"
			u, err := m.TransferURL(ctx, turl)
			if err != nil {
				t.Fatal(err)
			}
			if u.Hostname() != "127.0.0.1" {
				t.Errorf("hostname=%s", u.Hostname())
			}

			ipv6 := "[fd01:1:3:1528::10bf]:443"
			turl = fmt.Sprintf("https://%s/foo/bar", ipv6)
			u, err = m.TransferURL(ctx, turl)
			if err != nil {
				t.Fatal(err)
			}
			if u.Host != ipv6 {
				t.Errorf("host=%s", u.Host)
			}
		}

		// hostname should be returned by TransferURL when there are multiple management IPs
		for _, nc := range host.Config.VirtualNicManagerInfo.NetConfig {
			if nc.NicType == string(types.HostVirtualNicManagerNicTypeManagement) {
				host.Config.VirtualNicManagerInfo.NetConfig = append(host.Config.VirtualNicManagerInfo.NetConfig, nc)
				break
			}
		}

		for i := 0; i < 2; i++ { // 2nd time is to validate the cached value
			turl := "https://esx2:443/foo/bar"
			u, err := m.TransferURL(ctx, turl)
			if err != nil {
				t.Fatal(err)
			}
			if u.Hostname() != "esx2" {
				t.Errorf("hostname=%s", u.Hostname())
			}
		}

		// error should be returned if HostSystem.Config is nil
		host.Config = nil
		m, err = ops.FileManager(ctx)
		if err != nil {
			t.Fatal(err)
		}

		turl := "https://noconfig:443/foo/bar"
		u, err := m.TransferURL(ctx, turl)
		if err == nil {
			t.Errorf("expected error (url=%s)", u)
		}
	})
}
