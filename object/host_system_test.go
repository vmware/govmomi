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
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
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
		hs := simulator.Map.Get(host.Reference()).(*simulator.HostSystem)
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
