// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func TestHostOptionManager(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := m.Service.Context

	hs := NewHostSystem(ctx, esx.HostSystem)

	advOpts, ok := m.Map().Get(hs.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	require.True(t, ok, "Expected to inflate OptionManager from reference")

	option := &types.OptionValue{
		Key:   "TEST.hello",
		Value: "world",
	}

	fault := advOpts.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Fault()
	require.IsType(t, &types.InvalidName{}, fault.VimFault(), "Expected new host from template not to have test option set")

	fault = advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option}}).Fault()
	require.Nil(t, fault, "Expected setting test option to succeed")

	queryRes := advOpts.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Res
	require.Equal(t, 1, len(queryRes.Returnval), "Expected query of set option to succeed")
	require.Equal(t, option.Value, queryRes.Returnval[0].GetOptionValue().Value, "Expected set value")

	option2 := &types.OptionValue{
		Key:   "TEST.hello",
		Value: "goodbye",
	}

	fault = advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option2}}).Fault()
	require.Nil(t, fault, "Expected update of test option to succeed")

	queryRes = advOpts.QueryOptions(&types.QueryOptions{Name: option2.Key}).(*methods.QueryOptionsBody).Res
	require.Equal(t, 1, len(queryRes.Returnval), "Expected query of updated option to succeed")
	require.Equal(t, option2.Value, queryRes.Returnval[0].GetOptionValue().Value, "Expected updated value")

	hs.configure(ctx, types.HostConnectSpec{}, true)
	assert.Nil(t, hs.sh, "Expected not to have container backing if not requested")
}

func TestSyncWithOptionsStruct(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := m.Service.Context
	hs := NewHostSystem(ctx, esx.HostSystem)

	advOpts, ok := m.Map().Get(hs.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	require.True(t, ok, "Expected to inflate OptionManager from reference")

	option := &types.OptionValue{
		Key:   "TEST.hello",
		Value: "world",
	}

	fault := advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option}}).Fault()
	require.Nil(t, fault, "Expected setting test option to succeed")

	assert.Equal(t, option, hs.Config.Option[1], "Expected mirror to reflect changes")
}

func TestPerHostOptionManager(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := m.Service.Context
	hs := NewHostSystem(ctx, esx.HostSystem)
	hs2 := NewHostSystem(ctx, esx.HostSystem)

	advOpts, ok := m.Map().Get(hs.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	require.True(t, ok, "Expected to inflate OptionManager from reference")

	advOpts2 := m.Map().Get(hs2.ConfigManager.AdvancedOption.Reference()).(*OptionManager)

	option := &types.OptionValue{
		Key:   "TEST.hello",
		Value: "world",
	}

	fault := advOpts.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Fault()
	require.IsType(t, &types.InvalidName{}, fault.VimFault(), "Expected host from template not to have test option set")

	fault = advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option}}).Fault()
	require.Nil(t, fault, "Expected setting test option to succeed")

	queryRes := advOpts.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Res
	require.Equal(t, 1, len(queryRes.Returnval), "Expected query of set option to succeed")
	require.Equal(t, option.Value, queryRes.Returnval[0].GetOptionValue().Value, "Expected set value")

	fault = advOpts2.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Fault()
	require.IsType(t, &types.InvalidName{}, fault.VimFault(), "Expected second host to be unchanged")

	option2 := &types.OptionValue{
		Key:   "TEST.hello",
		Value: "goodbye",
	}

	fault = advOpts.UpdateOptions(&types.UpdateOptions{ChangedValue: []types.BaseOptionValue{option2}}).Fault()
	require.Nil(t, fault, "Expected update of test option to succeed")

	queryRes = advOpts.QueryOptions(&types.QueryOptions{Name: option2.Key}).(*methods.QueryOptionsBody).Res
	require.Equal(t, 1, len(queryRes.Returnval), "Expected query of updated option to succeed")
	require.Equal(t, option2.Value, queryRes.Returnval[0].GetOptionValue().Value, "Expected updated value")

	assert.Equal(t, option2, hs.Config.Option[1], "Expected mirror to reflect changes")

	hs.configure(ctx, types.HostConnectSpec{}, true)
	assert.Nil(t, hs.sh, "Expected not to have container backing if not requested")

	hs3 := NewHostSystem(ctx, esx.HostSystem)

	advOpts3 := m.Map().Get(hs3.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	fault = advOpts3.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Fault()
	require.IsType(t, &types.InvalidName{}, fault.VimFault(), "Expected host created after update not to inherit change")

}

func TestHostContainerBacking(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker on linux")
	}

	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := m.Service.Context

	hs := NewHostSystem(ctx, esx.HostSystem)
	hs.configureContainerBacking(ctx, "alpine", defaultSimVolumes, "vcsim-mgmt-underlay")

	details, err := hs.getNetConfigInterface(ctx, "management")
	assert.NoError(t, err, "Expected no error from management netconfig check")
	assert.Equal(t, "0.0.0.0", details.vmk.Spec.Ip.IpAddress, "Expected IP to be empty prior to container creation")

	hs.configure(ctx, types.HostConnectSpec{}, true)

	assert.NoError(t, err, "Expected no error from management netconfig check")
	assert.NotEqual(t, "0.0.0.0", details.vmk.Spec.Ip.IpAddress, "Expected management IP to set after container creation")

	hs.sh.remove(ctx)
}

func TestMultipleSimHost(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker on linux")
	}

	m := ESX()

	defer m.Remove()

	err := m.Create()
	require.Nil(t, err, "expected successful creation of model")

	ctx := m.Service.Context

	hs := NewHostSystem(ctx, esx.HostSystem)
	hs.configureContainerBacking(ctx, "alpine", defaultSimVolumes)

	hs2 := NewHostSystem(ctx, esx.HostSystem)
	hs2.configureContainerBacking(ctx, "alpine", defaultSimVolumes)

	details, err := hs.getNetConfigInterface(ctx, "management")
	assert.NoError(t, err, "Expected no error from management netconfig check")
	assert.Equal(t, "0.0.0.0", details.vmk.Spec.Ip.IpAddress, "Expected IP to be empty prior to container creation")

	hs.configure(ctx, types.HostConnectSpec{}, true)

	details2, err := hs2.getNetConfigInterface(ctx, "management")
	assert.NoError(t, err, "Expected no error from management netconfig check")
	assert.Equal(t, "0.0.0.0", details2.vmk.Spec.Ip.IpAddress, "Expected IP to be empty prior to container creation")

	hs2.configure(ctx, types.HostConnectSpec{}, true)

	assert.NotEqual(t, details.vmk.Spec.Ip.IpAddress, details2.vmk.Spec.Ip.IpAddress, "Expected hosts to get different IPs")

	hs.sh.remove(ctx)

	// TODO: assert one container plus volumes left - need to wait for
	// https://github.com/containers/podman/issues/19219 to be fixed for podman to work - otherwise all volumes get removed
	// with the first host removed
	hs2.sh.remove(ctx)
}
