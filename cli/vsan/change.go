// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vsan

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vsan"
	"github.com/vmware/govmomi/vsan/types"
)

type change struct {
	*flags.DatacenterFlag

	unmap *bool
	fs    *bool
}

func init() {
	cli.Register("vsan.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.Var(flags.NewOptionalBool(&cmd.unmap), "unmap-enabled", "Enable Unmap")
	f.Var(flags.NewOptionalBool(&cmd.fs), "file-service-enabled", "Enable FileService")
}

func (cmd *change) Usage() string {
	return "CLUSTER"
}

func (cmd *change) Description() string {
	return `Change vSAN configuration.

Examples:
  govc vsan.change -unmap-enabled ClusterA # enable unmap
  govc vsan.change -unmap-enabled=false ClusterA # disable unmap`
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	cluster, err := finder.ClusterComputeResourceOrDefault(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	c, err := vsan.NewClient(ctx, vc)
	if err != nil {
		return err
	}

	c.RoundTripper = cmd.RoundTripper(c.Client)

	var spec types.VimVsanReconfigSpec

	if cmd.unmap == nil && cmd.fs == nil {
		return flag.ErrHelp
	}
	if cmd.unmap != nil {
		spec.UnmapConfig = &types.VsanUnmapConfig{Enable: *cmd.unmap}
	}
	if cmd.fs != nil {
		spec.FileServiceConfig = &types.VsanFileServiceConfig{Enabled: *cmd.fs}
	}

	task, err := c.VsanClusterReconfig(ctx, cluster.Reference(), spec)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger("Updating vSAN...")
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
