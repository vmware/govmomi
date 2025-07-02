// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package override

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*flags.ClusterFlag
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("cluster.override.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *remove) Description() string {
	return `Remove cluster VM overrides.

Examples:
  govc cluster.override.remove -cluster cluster_1 -vm vm_1`
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
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
	ref := vm.Reference()

	for _, c := range config.DrsVmConfig {
		if c.Key == ref {
			spec.DrsVmConfigSpec = []types.ClusterDrsVmConfigSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationRemove,
						RemoveKey: ref,
					},
				},
			}
			break
		}
	}

	for _, c := range config.DasVmConfig {
		if c.Key == ref {
			spec.DasVmConfigSpec = []types.ClusterDasVmConfigSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationRemove,
						RemoveKey: ref,
					},
				},
			}
			break
		}
	}

	for _, c := range config.VmOrchestration {
		if c.Vm == ref {
			spec.VmOrchestrationSpec = []types.ClusterVmOrchestrationSpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{
						Operation: types.ArrayUpdateOperationRemove,
						RemoveKey: ref,
					},
				},
			}
			break
		}
	}

	return cmd.Reconfigure(ctx, spec)
}
