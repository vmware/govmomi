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

package object_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
)

func TestOpaqueNetworkEthernetCardBackingInfo(t *testing.T) {
	model := simulator.VPX()
	model.OpaqueNetwork = 1 // Create 1 NSX backed OpaqueNetwork per DC

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		finder := find.NewFinder(c)

		net, err := finder.Network(ctx, "DC0_NSX0")
		if err != nil {
			t.Fatal(err)
		}

		nsx, ok := net.(*object.OpaqueNetwork)
		if !ok {
			t.Fatalf("network type=%T", net)
		}

		summary, err := nsx.Summary(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if net.Reference() != *summary.Network {
			t.Fatal()
		}

		_, err = uuid.Parse(summary.OpaqueNetworkId)
		if err != nil {
			t.Errorf("parsing %q: %s", summary.OpaqueNetworkId, err)
		}

		_, err = net.EthernetCardBackingInfo(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}, model)
}
