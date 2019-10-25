/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
		obj := simulator.Map.Any("HostSystem").(*simulator.HostSystem)
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
