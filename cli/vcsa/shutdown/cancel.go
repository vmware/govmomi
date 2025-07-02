// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package shutdown

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/appliance/shutdown"
)

type cancel struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("vcsa.shutdown.cancel", &cancel{})
}

func (cmd *cancel) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *cancel) Description() string {
	return `Cancel pending shutdown action.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.cancel`
}

func (cmd *cancel) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := shutdown.NewManager(c)

	if err = m.Cancel(ctx); err != nil {
		return err
	}

	return nil
}
