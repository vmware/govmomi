/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package guest_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestTranferURL(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		vm := simulator.Map.Any("VirtualMachine").(*simulator.VirtualMachine)
		host := simulator.Map.Get(*vm.Runtime.Host).(*simulator.HostSystem)

		ops := guest.NewOperationsManager(c, vm.Reference())
		m, err := ops.FileManager(ctx)
		if err != nil {
			t.Fatal(err)
		}

		turl := "https://esx:443/foo/bar"
		u, err := m.TransferURL(ctx, turl)
		if err != nil {
			t.Fatal(err)
		}
		if u.Hostname() != "127.0.0.1" {
			t.Errorf("hostname=%s", u.Hostname())
		}

		host.Config = nil
		m, err = ops.FileManager(ctx)
		if err != nil {
			t.Fatal(err)
		}

		u, err = m.TransferURL(ctx, turl)
		if err == nil {
			t.Errorf("expected error (url=%s)", u)
		}
		t.Log(err)
	})
}
