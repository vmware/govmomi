// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"bytes"
	"context"
	"encoding/pem"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func TestHostSystemManagementIPs(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host, err := find.NewFinder(c).HostSystem(ctx, "DC0_C0_H0")
		if err != nil {
			t.Fatal(err)
		}

		ips, err := host.ManagementIPs(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if len(ips) != 1 {
			t.Fatal("no mgmt ip found")
		}
		if ips[0].String() != "127.0.0.1" {
			t.Fatalf("Expected management ip %s, got %s", "127.0.0.1", ips[0].String())
		}

		// These fields can be nil while ESX is being upgraded
		hs := simulator.Map(ctx).Get(host.Reference()).(*simulator.HostSystem)
		tests := []func(){
			func() { hs.Config.VirtualNicManagerInfo = nil },
			func() { hs.Config = nil },
		}

		for _, f := range tests {
			f()
			ips, err = host.ManagementIPs(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(ips) != 0 {
				t.Fatal("expected zero ips")
			}
		}
	})
}

func TestHostSystemConfig(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host, err := find.NewFinder(c).HostSystem(ctx, "DC0_C0_H0")
		if err != nil {
			t.Fatal(err)
		}

		var props mo.HostSystem
		if err := host.Properties(ctx, host.Reference(), []string{"config"}, &props); err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(props.Config.Certificate, esx.HostConfigInfo.Certificate) {
			t.Errorf("certificate=%s", string(props.Config.Certificate))
		}

		b, _ := pem.Decode(props.Config.Certificate)
		if b == nil {
			t.Error("failed to parse certificate")
		}
	})
}
