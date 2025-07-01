// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

func DrsBehaviorUsage() string {
	drsModes := types.DrsBehavior("").Strings()
	return "DRS behavior for virtual machines: " + strings.Join(drsModes, ", ")
}

type change struct {
	*flags.DatacenterFlag

	types.ClusterConfigSpecEx
}

func init() {
	cli.Register("cluster.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	cmd.DrsConfig = new(types.ClusterDrsConfigInfo)
	cmd.DasConfig = new(types.ClusterDasConfigInfo)
	cmd.VsanConfig = new(types.VsanClusterConfigInfo)
	cmd.VsanConfig.DefaultConfig = new(types.VsanClusterConfigInfoHostDefaultInfo)

	// DRS
	f.Var(flags.NewOptionalBool(&cmd.DrsConfig.Enabled), "drs-enabled", "Enable DRS")

	f.StringVar((*string)(&cmd.DrsConfig.DefaultVmBehavior), "drs-mode", "", DrsBehaviorUsage())

	f.Var(flags.NewInt32(&cmd.DrsConfig.VmotionRate), "drs-vmotion-rate", "Aggressiveness of vMotions (1-5)")

	// HA
	f.Var(flags.NewOptionalBool(&cmd.DasConfig.Enabled), "ha-enabled", "Enable HA")
	f.Var(flags.NewOptionalBool(&cmd.DasConfig.AdmissionControlEnabled), "ha-admission-control-enabled", "Enable HA admission control")

	// vSAN
	f.Var(flags.NewOptionalBool(&cmd.VsanConfig.Enabled), "vsan-enabled", "Enable vSAN")
	f.Var(flags.NewOptionalBool(&cmd.VsanConfig.DefaultConfig.AutoClaimStorage), "vsan-autoclaim", "Autoclaim storage on cluster hosts")
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) Usage() string {
	return "CLUSTER..."
}

func (cmd *change) Description() string {
	return `Change configuration of the given clusters.

Examples:
  govc cluster.change -drs-enabled -vsan-enabled -vsan-autoclaim ClusterA
  govc cluster.change -drs-enabled=false ClusterB
  govc cluster.change -drs-vmotion-rate=4 ClusterC`
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 {
		return flag.ErrHelp
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	for _, path := range f.Args() {
		clusters, err := finder.ClusterComputeResourceList(ctx, path)
		if err != nil {
			return err
		}

		for _, cluster := range clusters {
			task, err := cluster.Reconfigure(ctx, &cmd.ClusterConfigSpecEx, true)
			if err != nil {
				return err
			}

			_, err = task.WaitForResult(ctx, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
