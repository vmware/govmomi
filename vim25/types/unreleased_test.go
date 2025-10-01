// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types_test

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestPodVMInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host := simulator.Map(ctx).Any("HostSystem").(*simulator.HostSystem)

		host.Runtime.PodVMInfo = &types.PodVMInfo{
			HasPageSharingPodVM: true,
			PodVMOverheadInfo: types.PodVMOverheadInfo{
				PodVMOverheadWithoutPageSharing: int32(50),
				PodVMOverheadWithPageSharing:    int32(25),
			},
		}

		var props mo.HostSystem
		pc := property.DefaultCollector(c)
		err := pc.RetrieveOne(ctx, host.Self, []string{"runtime"}, &props)
		if err != nil {
			t.Fatal(err)
		}

		if *props.Runtime.PodVMInfo != *host.Runtime.PodVMInfo {
			t.Errorf("%#v", props.Runtime.PodVMInfo)
		}
	})
}
