/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package consumptiondomains_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/vcenter/consumptiondomains/associations"
	"github.com/vmware/govmomi/vapi/vcenter/consumptiondomains/zones"
	"github.com/vmware/govmomi/vim25"

	_ "github.com/vmware/govmomi/vapi/simulator"
	_ "github.com/vmware/govmomi/vapi/vcenter/consumptiondomains/simulator"
)

func TestConsumptionDomains(t *testing.T) {
	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		rc := rest.NewClient(vc)

		err := rc.Login(ctx, simulator.DefaultLogin)
		if err != nil {
			t.Fatal(err)
		}

		zm := zones.NewManager(rc)

		// Create a zone
		zoneId, err := zm.CreateZone(zones.CreateSpec{
			Zone:        "test-zone-1",
			Description: "placeholder description",
		})

		if err != nil {
			t.Error(err)
		}

		// List all zones and find the one created earlier
		zonesList, err := zm.ListZones()

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 1, len(zonesList))
		assert.Equal(t, "test-zone-1", zonesList[0].Zone)
		assert.Equal(t, "placeholder description", zonesList[0].Info.Description)

		// Query zone by ID
		zone, err := zm.GetZone(zoneId)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "placeholder description", zone.Description)

		am := associations.NewManager(rc)

		// Create a cluster association
		err = am.AddAssociations(zoneId, "domain-c9")

		if err != nil {
			t.Error(err)
		}

		// Query the associations for the test zone
		assc, err := am.GetAssociations(zoneId)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 1, len(assc))
		assert.Equal(t, "domain-c9", assc[0])

		// Delete the cluster association
		err = am.RemoveAssociations(zoneId, "domain-c9")

		if err != nil {
			t.Error(err)
		}

		// Verify that the association is removed
		assc, err = am.GetAssociations(zoneId)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 0, len(assc))

		// Delete the zone
		err = zm.DeleteZone(zoneId)

		if err != nil {
			t.Error(err)
		}

		// Verify the zone is removed
		zone, err = zm.GetZone(zoneId)

		if err == nil {
			t.Error(err)
		}

		zonesList, err = zm.ListZones()

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 0, len(zonesList))
	})
}
