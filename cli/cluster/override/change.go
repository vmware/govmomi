// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package override

import (
	"context"
	"flag"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/cluster"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*flags.ClusterFlag
	*flags.VirtualMachineFlag

	drs types.ClusterDrsVmConfigInfo
	das types.ClusterDasVmConfigInfo
	orc types.ClusterVmOrchestrationInfo
}

func init() {
	cli.Register("cluster.override.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	// DRS
	f.Var(flags.NewOptionalBool(&cmd.drs.Enabled), "drs-enabled", "Enable DRS")

	f.StringVar((*string)(&cmd.drs.Behavior), "drs-mode", "", cluster.DrsBehaviorUsage())

	// HA
	rp := types.ClusterDasVmSettingsRestartPriority("").Strings()
	cmd.das.DasSettings = new(types.ClusterDasVmSettings)

	f.StringVar((*string)(&cmd.das.DasSettings.RestartPriority), "ha-restart-priority", "", "HA restart priority: "+strings.Join(rp, ", "))

	f.Var(flags.NewInt32(&cmd.orc.VmReadiness.PostReadyDelay), "ha-additional-delay", "HA Additional Delay")

	rc := types.ClusterVmReadinessReadyCondition("").Strings()
	f.StringVar((*string)(&cmd.orc.VmReadiness.ReadyCondition), "ha-ready-condition", "", "HA VM Ready Condition (Start next priority VMs when): "+strings.Join(rc, ", "))
}

func (cmd *change) Description() string {
	return `Change cluster VM overrides.

Examples:
  govc cluster.override.change -cluster cluster_1 -vm vm_1 -drs-enabled=false
  govc cluster.override.change -cluster cluster_1 -vm vm_2 -drs-enabled -drs-mode fullyAutomated
  govc cluster.override.change -cluster cluster_1 -vm vm_3 -ha-restart-priority high
  govc cluster.override.change -cluster cluster_1 -vm vm_4 -ha-additional-delay 30
  govc cluster.override.change -cluster cluster_1 -vm vm_5 -ha-ready-condition poweredOn`
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	config, err := cluster.Configuration(ctx)
	if err != nil {
		return err
	}

	spec := &types.ClusterConfigSpecEx{}
	cmd.drs.Key = vm.Reference()
	cmd.das.Key = vm.Reference()
	cmd.orc.Vm = vm.Reference()

	if cmd.drs.Behavior != "" || cmd.drs.Enabled != nil {
		op := types.ArrayUpdateOperationAdd
		for _, c := range config.DrsVmConfig {
			if c.Key == cmd.drs.Key {
				op = types.ArrayUpdateOperationEdit
				break
			}
		}

		spec.DrsVmConfigSpec = []types.ClusterDrsVmConfigSpec{
			{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: op,
				},
				Info: &cmd.drs,
			},
		}
	}

	if cmd.das.DasSettings.RestartPriority != "" {
		op := types.ArrayUpdateOperationAdd
		for _, c := range config.DasVmConfig {
			if c.Key == cmd.das.Key {
				op = types.ArrayUpdateOperationEdit
				break
			}
		}

		spec.DasVmConfigSpec = []types.ClusterDasVmConfigSpec{
			{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: op,
				},
				Info: &cmd.das,
			},
		}
	}

	if cmd.orc.VmReadiness.PostReadyDelay > 0 || cmd.orc.VmReadiness.ReadyCondition != "" {
		op := types.ArrayUpdateOperationAdd
		for _, c := range config.VmOrchestration {
			if c.Vm == cmd.orc.Vm {
				op = types.ArrayUpdateOperationEdit
				break
			}
		}
		spec.VmOrchestrationSpec = []types.ClusterVmOrchestrationSpec{
			{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: op,
				},
				Info: &cmd.orc,
			},
		}
	}

	if spec.DrsVmConfigSpec == nil && spec.DasVmConfigSpec == nil && spec.VmOrchestrationSpec == nil {
		return flag.ErrHelp
	}

	return cmd.Reconfigure(ctx, spec)
}
