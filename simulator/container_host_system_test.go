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

package simulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vmware/govmomi/simulator/esx"
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

	hs := NewHostSystem(esx.HostSystem)

	advOpts, ok := Map.Get(hs.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
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

	hs.configure(SpoofContext(), types.HostConnectSpec{}, true)
	assert.Nil(t, hs.sh, "Expected not to have container backing if not requested")
}

func TestSyncWithOptionsStruct(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	hs := NewHostSystem(esx.HostSystem)

	advOpts, ok := Map.Get(hs.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
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

	hs := NewHostSystem(esx.HostSystem)
	hs2 := NewHostSystem(esx.HostSystem)

	advOpts, ok := Map.Get(hs.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	require.True(t, ok, "Expected to inflate OptionManager from reference")

	advOpts2 := Map.Get(hs2.ConfigManager.AdvancedOption.Reference()).(*OptionManager)

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

	hs.configure(SpoofContext(), types.HostConnectSpec{}, true)
	assert.Nil(t, hs.sh, "Expected not to have container backing if not requested")

	hs3 := NewHostSystem(esx.HostSystem)

	advOpts3 := Map.Get(hs3.ConfigManager.AdvancedOption.Reference()).(*OptionManager)
	fault = advOpts3.QueryOptions(&types.QueryOptions{Name: option.Key}).(*methods.QueryOptionsBody).Fault()
	require.IsType(t, &types.InvalidName{}, fault.VimFault(), "Expected host created after update not to inherit change")

}

func TestHostContainerBacking(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ctx := SpoofContext()

	hs := NewHostSystem(esx.HostSystem)
	hs.configureContainerBacking(ctx, "alpine", defaultSimVolumes)

	hs.configure(ctx, types.HostConnectSpec{}, true)

	//TODO: assert there's a container representing the host (consider a separate test for matching datastores and networks)

	hs.sh.remove(ctx)
}

func TestMultipleSimHost(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	require.Nil(t, err, "expected successful creation of model")

	ctx := SpoofContext()

	hs := NewHostSystem(esx.HostSystem)
	hs.configureContainerBacking(ctx, "alpine", defaultSimVolumes)

	hs.configure(ctx, types.HostConnectSpec{}, true)
	// TODO: assert container present

	hs2 := NewHostSystem(esx.HostSystem)
	hs2.configureContainerBacking(ctx, "alpine", defaultSimVolumes)

	hs2.configure(ctx, types.HostConnectSpec{}, true)
	// TODO: assert 2nd container present

	hs.sh.remove(ctx)

	// TODO: assert one container plus volumes left
	hs2.sh.remove(ctx)
}
