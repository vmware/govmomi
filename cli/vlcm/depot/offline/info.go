// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package offline

import (
	"context"
	"flag"
	"io"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/depots"
)

type infoResult depots.SettingsDepotsOfflineContentInfo

func (r infoResult) Write(w io.Writer) error {
	return nil
}

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag

	depotId string
}

func init() {
	cli.Register("vlcm.depot.offline.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)

	f.StringVar(&cmd.depotId, "depot-id", "", "The identifier of the depot. Use the 'ls' command to see the list of depots.")
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
	return "VLCM"
}

func (cmd *info) Description() string {
	return `Displays the contents of an offline image depot.

Examples:
  govc vlcm.depot.offline.info -depot-id=<your depot's identifier>`
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := depots.NewManager(rc)

	if d, err := dm.GetOfflineDepotContent(cmd.depotId); err != nil {
		return err
	} else {
		if !cmd.All() {
			cmd.JSON = true
		}

		return cmd.WriteResult(infoResult(d))
	}
}
