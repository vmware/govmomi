// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestHostConfigManager(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		obj := simulator.Map(ctx).Any("HostSystem").(*simulator.HostSystem)
		host := object.NewHostSystem(c, obj.Self)

		m := host.ConfigManager()
		// All should succeed
		funcs := []string{
			"DatastoreSystem",
			"NetworkSystem",
			"FirewallSystem",
			"StorageSystem",
			"VirtualNicManager",
			"VsanSystem",
			"VsanInternalSystem",
			"AccountManager",
			"OptionManager",
			"ServiceSystem",
			"CertificateManager",
			"DateTimeSystem",
		}

		rm := reflect.ValueOf(m)
		rctx := reflect.ValueOf(ctx)

		for _, name := range funcs {
			method := rm.MethodByName(name)
			ret := method.Call([]reflect.Value{rctx})
			err := ret[1]
			if !err.IsNil() {
				t.Errorf("%s: %s", name, err.Interface().(error))
			}
		}

		// Force some errors

		obj.ConfigManager.NetworkSystem = nil

		_, err := host.ConfigManager().NetworkSystem(ctx)
		if err == nil {
			t.Error("expected error")
		}

		obj.ConfigManager.VsanSystem = nil
		_, err = host.ConfigManager().VsanSystem(ctx)
		if err != object.ErrNotSupported {
			t.Errorf("expected %s", object.ErrNotSupported)
		}
	})
}
