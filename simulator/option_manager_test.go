// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/types"
)

func TestOptionManagerESX(t *testing.T) {
	ctx := context.Background()

	model := ESX()
	model.Datastore = 0
	model.Machine = 0

	err := model.Create()
	if err != nil {
		t.Fatal(err)
	}

	c := model.Service.client()

	m := object.NewOptionManager(c, *c.ServiceContent.Setting)
	_, err = m.Query(ctx, "config.vpxd.")
	if err == nil {
		t.Error("expected error")
	}

	host := object.NewHostSystem(c, esx.HostSystem.Reference())
	m, err = host.ConfigManager().OptionManager(ctx)
	if err != nil {
		t.Fatal(err)
	}

	res, err := m.Query(ctx, "Config.HostAgent.")
	if err != nil {
		t.Error(err)
	}

	if len(res) == 0 {
		t.Error("no results")
	}

	err = m.Update(ctx, []types.BaseOptionValue{&types.OptionValue{
		Key:   "Config.HostAgent.log.level",
		Value: "verbose",
	}})

	if err != nil {
		t.Error(err)
	}
}

func TestOptionManagerVPX(t *testing.T) {
	ctx := context.Background()

	model := VPX()
	model.Datastore = 0
	model.Machine = 0

	err := model.Create()
	if err != nil {
		t.Fatal(err)
	}

	c := model.Service.client()

	m := object.NewOptionManager(c, *c.ServiceContent.Setting)
	_, err = m.Query(ctx, "enoent")
	if err == nil {
		t.Error("expected error")
	}

	res, err := m.Query(ctx, "event.")
	if err != nil {
		t.Error(err)
	}

	if len(res) == 0 {
		t.Error("no results")
	}

	val := &types.OptionValue{
		Key: "event.maxAge",
	}

	// Get the existing maxAge value
	for _, r := range res {
		opt := r.GetOptionValue()
		if opt.Key == val.Key {
			val.Value = opt.Value
		}
	}

	// Increase maxAge * 2
	val.Value = val.Value.(int32) * 2
	err = m.Update(ctx, []types.BaseOptionValue{val})
	if err != nil {
		t.Error(err)
	}

	// Verify maxAge was updated
	res, err = m.Query(ctx, val.Key)
	if err != nil {
		t.Error(err)
	}
	if res[0].GetOptionValue().Value != val.Value {
		t.Errorf("%s was not updated", val.Key)
	}

	// Expected to throw InvalidName fault
	err = m.Update(ctx, []types.BaseOptionValue{&types.OptionValue{
		Key: "ENOENT.anything",
	}})
	if err == nil {
		t.Error("expected error")
	}

	// Add a new option
	err = m.Update(ctx, []types.BaseOptionValue{&types.OptionValue{
		Key:   "config.anything",
		Value: "OK",
	}})
	if err != nil {
		t.Error(err)
	}
}
