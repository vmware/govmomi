// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func TestEnvironmentBrowser(t *testing.T) {
	for _, model := range []*Model{ESX(), VPX()} {
		require.NoError(t, model.Run(func(ctx context.Context, c *vim25.Client) error {
			for _, entity := range Map(ctx).AllReference("EnvironmentBrowser") {
				b := entity.(*EnvironmentBrowser)
				hosts := b.hosts(model.Service.Context)
				var expected int32
				for i, ref := range hosts {
					host := Map(ctx).Get(ref).(*HostSystem)
					host.Summary.Hardware.MemorySize = int64(4096+i*1024) * 1024 * 1024
					expected = int32(host.Summary.Hardware.MemorySize / (1024 * 1024))
				}
				selections := []*types.ManagedObjectReference{nil}
				if len(hosts) > 0 {
					selections = append(selections, &hosts[0])
				}
				for _, host := range selections {
					res, err := methods.QueryConfigTarget(ctx, c, &types.QueryConfigTarget{This: b.Self, Host: host})
					require.NoError(t, err)
					want := expected
					if host != nil {
						want = 4096
					}
					require.Equal(t, want, res.Returnval.MaxMemMBOptimalPerf)
				}
				if len(hosts) > 1 {
					cluster := Map(ctx).Any("ClusterComputeResource").(*ClusterComputeResource)
					saved := cluster.Host
					cluster.Host = nil
					res, err := methods.QueryConfigTarget(ctx, c, &types.QueryConfigTarget{This: b.Self})
					cluster.Host = saved
					require.NoError(t, err)
					require.Zero(t, res.Returnval.MaxMemMBOptimalPerf)
					require.Empty(t, res.Returnval.Datastore)
				}
				var env mo.EnvironmentBrowser
				require.NoError(t, property.DefaultCollector(c).RetrieveOne(ctx, b.Self, []string{"datastoreBrowser"}, &env))
				require.NotNil(t, env.DatastoreBrowser)
				var browser mo.HostDatastoreBrowser
				require.NoError(t, property.DefaultCollector(c).RetrieveOne(ctx, *env.DatastoreBrowser, nil, &browser))
				require.Equal(t, "HostDatastoreBrowser", browser.Self.Type)
				for i := 0; i < 2; i++ {
					basic, err := methods.QueryConfigOption(ctx, c, &types.QueryConfigOption{This: b.Self})
					require.NoError(t, err)
					extended, err := methods.QueryConfigOptionEx(ctx, c, &types.QueryConfigOptionEx{This: b.Self, Spec: &types.EnvironmentBrowserConfigOptionQuerySpec{GuestId: []string{"ubuntu64Guest"}}})
					require.NoError(t, err)
					require.Equal(t, "ubuntu64Guest", extended.Returnval.GuestOSDescriptor[0].Id)
					for _, opt := range []*types.VirtualMachineConfigOption{basic.Returnval, extended.Returnval} {
						require.Len(t, opt.HardwareOptions.VirtualDeviceOption, 1)
						require.IsType(t, &types.VirtualLsiLogicControllerOption{}, opt.HardwareOptions.VirtualDeviceOption[0])
						require.Len(t, opt.DefaultDevice, len(esx.VirtualDevice))
						data, err := xml.Marshal(opt)
						require.NoError(t, err)
						require.Equal(t, 2, strings.Count(string(data), "<defaultLevel>normal</defaultLevel>"))
						require.Contains(t, string(data), "virtualDeviceOption")
						require.NotContains(t, string(data), "<defaultLevel></defaultLevel>")
					}
				}
				recorded := &types.VirtualMachineConfigOption{
					GuestOSDescriptor: []types.GuestOsDescriptor{{Id: "recorded"}},
					HardwareOptions: types.VirtualHardwareOption{
						VirtualDeviceOption: []types.BaseVirtualDeviceOption{&types.VirtualLsiLogicControllerOption{}},
					},
				}
				b.QueryConfigOptionResponse.Returnval = recorded
				b.QueryConfigTargetResponse.Returnval = &types.ConfigTarget{MaxMemMBOptimalPerf: 1234}
				for i := 0; i < 2; i++ {
					res, err := methods.QueryConfigTarget(ctx, c, &types.QueryConfigTarget{This: b.Self})
					require.NoError(t, err)
					require.EqualValues(t, 1234, res.Returnval.MaxMemMBOptimalPerf)
					_, err = methods.QueryConfigOptionEx(ctx, c, &types.QueryConfigOptionEx{This: b.Self, Spec: &types.EnvironmentBrowserConfigOptionQuerySpec{GuestId: []string{"ubuntu64Guest"}}})
					require.NoError(t, err)
					basic, err := methods.QueryConfigOption(ctx, c, &types.QueryConfigOption{This: b.Self})
					require.NoError(t, err)
					require.Equal(t, "recorded", basic.Returnval.GuestOSDescriptor[0].Id)
					require.Len(t, basic.Returnval.HardwareOptions.VirtualDeviceOption, 1)
				}
			}
			finder := find.NewFinder(c)
			dc, err := finder.DefaultDatacenter(ctx)
			require.NoError(t, err)
			finder.SetDatacenter(dc)
			folders, err := dc.Folders(ctx)
			require.NoError(t, err)
			pools, err := finder.ResourcePoolList(ctx, "*")
			require.NoError(t, err)
			ds, err := finder.DefaultDatastore(ctx)
			require.NoError(t, err)
			task, err := folders.VmFolder.CreateVM(ctx, types.VirtualMachineConfigSpec{
				Name: "environment-browser-vm", GuestId: "ubuntu64Guest", NumCPUs: 1, MemoryMB: 1024,
				Files: &types.VirtualMachineFileInfo{VmPathName: "[" + ds.Name() + "]"},
			}, pools[0], nil)
			require.NoError(t, err)
			require.NoError(t, task.Wait(ctx))
			return nil
		}))
	}
}
