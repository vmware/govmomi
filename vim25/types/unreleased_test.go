// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
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

func TestPodVMOverheadInfo(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		host := simulator.Map(ctx).Any("HostSystem").(*simulator.HostSystem)

		host.Capability.PodVMOverheadInfo = &types.PodVMOverheadInfo{
			CrxPageSharingSupported:         true,
			PodVMOverheadWithoutPageSharing: int32(42),
			PodVMOverheadWithPageSharing:    int32(53),
		}

		var props mo.HostSystem
		pc := property.DefaultCollector(c)
		err := pc.RetrieveOne(ctx, host.Self, []string{"capability"}, &props)
		if err != nil {
			t.Fatal(err)
		}

		if *props.Capability.PodVMOverheadInfo != *host.Capability.PodVMOverheadInfo {
			t.Errorf("%#v", props.Capability.PodVMOverheadInfo)
		}
	})
}
