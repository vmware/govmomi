// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package draft

import (
	"context"
	"flag"
	"io"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type lsResult map[string]clusters.SettingsClustersSoftwareDraftsMetadata

func (r lsResult) Write(w io.Writer) error {
	return nil
}

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	clusterId string
	owners    string
}

func init() {
	cli.Register("cluster.draft.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.owners, "owners", "", "A comma-separated list of owners to filter by.")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *ls) Usage() string {
	return "CLUSTER"
}

func (cmd *ls) Description() string {
	return `Displays the list of software drafts.

Examples:
  govc cluster.draft.ls -cluster-id=domain-c21
  govc cluster.draft.ls -cluster-id=domain-c21 -owners=stoyan1,stoyan2`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	var owners *[]string
	if cmd.owners != "" {
		o := strings.Split(cmd.owners, ",")
		owners = &o
	}

	if d, err := dm.ListSoftwareDrafts(cmd.clusterId, owners); err != nil {
		return err
	} else {
		if !cmd.All() {
			cmd.JSON = true
		}
		return cmd.WriteResult(lsResult(d))
	}
}
