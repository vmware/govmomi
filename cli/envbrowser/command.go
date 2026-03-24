// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package envbrowser

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type command struct {
	*flags.ClusterFlag
	*flags.HostSystemFlag
	*flags.OutputFlag

	hardwareVersion     string
	guestIDs            string
	allHardwareVersions bool
	copyToFile          bool

	host    *object.HostSystem
	cluster *object.ClusterComputeResource

	eb *object.EnvironmentBrowser
}

func (cmd *command) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.BoolVar(
		&cmd.copyToFile,
		"copy-to-file",
		false,
		"True to marshal the result's XML to a file.")
}

func (cmd *command) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *command) Run(ctx context.Context, f *flag.FlagSet) error {
	cluster, err := cmd.ClusterIfSpecified()
	if err != nil {
		return err
	}
	cmd.cluster = cluster

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}
	cmd.host = host

	if cluster != nil {
		eb, err := cluster.EnvironmentBrowser(ctx)
		if err != nil {
			return err
		}
		cmd.eb = eb
	} else if host != nil {
		pool, err := host.ResourcePool(ctx)
		if err != nil {
			return err
		}

		var cr *object.ComputeResource

		crRef, err := pool.Owner(ctx)
		if err != nil {
			return err
		}

		switch tCR := crRef.(type) {
		case *object.ComputeResource:
			cr = tCR
		case *object.ClusterComputeResource:
			cr = &tCR.ComputeResource
		}

		eb, err := cr.EnvironmentBrowser(ctx)
		if err != nil {
			return err
		}
		cmd.eb = eb
	}

	return cmd.run(ctx, cmd.eb)
}

func (cmd *command) run(
	_ context.Context,
	_ *object.EnvironmentBrowser) error {

	return nil
}
