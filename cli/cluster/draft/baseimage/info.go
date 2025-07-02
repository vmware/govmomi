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

type infoResult clusters.SettingsBaseImageInfo

func (r infoResult) Write(w io.Writer) error {
	return nil
}

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clusterId string
	draftId   string
}

func init() {
	cli.Register("cluster.draft.baseimage.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.draftId, "draft-id", "", "The identifier of the software draft.")
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *info) Usage() string {
	return "CLUSTER"
}

func (cmd *info) Description() string {
	return `Displays the base image version of a software draft.

Examples:
  govc cluster.draft.baseimage.info -cluster-id=domain-c21 -draft-id=13`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	if d, err := dm.GetSoftwareDraftBaseImage(cmd.clusterId, cmd.draftId); err != nil {
		return err
	} else {
		if !cmd.All() {
			cmd.JSON = true
		}
		return cmd.WriteResult(infoResult(d))
	}
}
