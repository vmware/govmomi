// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package draft

import (
	"context"
	"flag"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type createResult string

func (r createResult) Write(w io.Writer) error {
	return nil
}

type create struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clusterId string
}

func init() {
	cli.Register("cluster.draft.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
}

func (cmd *create) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *create) Usage() string {
	return "CLUSTER"
}

func (cmd *create) Description() string {
	return `Creates a new software draft.

There can be only one active draft at a time on every cluster.

Examples:
  govc cluster.draft.create -cluster-id=domain-c21`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	if !cmd.All() {
		cmd.JSON = true
	}

	if draftId, err := dm.CreateSoftwareDraft(cmd.clusterId); err != nil {
		return err
	} else if err := cmd.WriteResult(createResult(draftId)); err != nil {
		return err
	} else {
		return nil
	}
}
