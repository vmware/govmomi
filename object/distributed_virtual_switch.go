// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type DistributedVirtualSwitch struct {
	Common
}

func NewDistributedVirtualSwitch(c *vim25.Client, ref types.ManagedObjectReference) *DistributedVirtualSwitch {
	return &DistributedVirtualSwitch{
		Common: NewCommon(c, ref),
	}
}

func (s DistributedVirtualSwitch) GetInventoryPath() string {
	return s.InventoryPath
}

func (s DistributedVirtualSwitch) EthernetCardBackingInfo(ctx context.Context) (types.BaseVirtualDeviceBackingInfo, error) {
	ref := s.Reference()
	name := s.InventoryPath
	if name == "" {
		name = ref.String()
	}
	return nil, fmt.Errorf("type %s (%s) cannot be used for EthernetCardBackingInfo", ref.Type, name)
}

func (s DistributedVirtualSwitch) Reconfigure(ctx context.Context, spec types.BaseDVSConfigSpec) (*Task, error) {
	req := types.ReconfigureDvs_Task{
		This: s.Reference(),
		Spec: spec,
	}

	res, err := methods.ReconfigureDvs_Task(ctx, s.Client(), &req)
	if err != nil {
		return nil, err
	}

	return NewTask(s.Client(), res.Returnval), nil
}

func (s DistributedVirtualSwitch) AddPortgroup(ctx context.Context, spec []types.DVPortgroupConfigSpec) (*Task, error) {
	req := types.AddDVPortgroup_Task{
		This: s.Reference(),
		Spec: spec,
	}

	res, err := methods.AddDVPortgroup_Task(ctx, s.Client(), &req)
	if err != nil {
		return nil, err
	}

	return NewTask(s.Client(), res.Returnval), nil
}

func (s DistributedVirtualSwitch) FetchDVPorts(ctx context.Context, criteria *types.DistributedVirtualSwitchPortCriteria) ([]types.DistributedVirtualPort, error) {
	req := &types.FetchDVPorts{
		This:     s.Reference(),
		Criteria: criteria,
	}

	res, err := methods.FetchDVPorts(ctx, s.Client(), req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil
}

func (s DistributedVirtualSwitch) ReconfigureDVPort(ctx context.Context, spec []types.DVPortConfigSpec) (*Task, error) {
	req := types.ReconfigureDVPort_Task{
		This: s.Reference(),
		Port: spec,
	}

	res, err := methods.ReconfigureDVPort_Task(ctx, s.Client(), &req)
	if err != nil {
		return nil, err
	}

	return NewTask(s.Client(), res.Returnval), nil
}

func (s DistributedVirtualSwitch) ReconfigureLACP(ctx context.Context, spec []types.VMwareDvsLacpGroupSpec) (*Task, error) {
	req := types.UpdateDVSLacpGroupConfig_Task{
		This:          s.Reference(),
		LacpGroupSpec: spec,
	}

	res, err := methods.UpdateDVSLacpGroupConfig_Task(ctx, s.Client(), &req)
	if err != nil {
		return nil, err
	}

	return NewTask(s.Client(), res.Returnval), nil
}
