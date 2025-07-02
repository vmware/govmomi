// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
)

func TestHostFirewallSystem(t *testing.T) {
	ctx := context.Background()

	m := ESX()
	m.Datastore = 0
	m.Machine = 0

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	c := m.Service.client()

	host := object.NewHostSystem(c, esx.HostSystem.Reference())

	hfs, _ := host.ConfigManager().FirewallSystem(ctx)

	err = hfs.DisableRuleset(ctx, "enoent")
	if err == nil {
		t.Error("expected error")
	}

	err = hfs.EnableRuleset(ctx, "enoent")
	if err == nil {
		t.Error("expected error")
	}

	err = hfs.DisableRuleset(ctx, "sshServer")
	if err != nil {
		t.Error(err)
	}

	err = hfs.EnableRuleset(ctx, "sshServer")
	if err != nil {
		t.Error(err)
	}

	_, err = hfs.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
