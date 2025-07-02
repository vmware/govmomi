// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
