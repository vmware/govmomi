// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package dcui

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/appliance/access/dcui"
)

type set struct {
	*flags.ClientFlag
	enabled bool
}

func init() {
	cli.Register("vcsa.access.dcui.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.enabled,
		"enabled",
		false,
		"Enable Direct Console User Interface (DCUI TTY2).")
}

func (cmd *set) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *set) Description() string {
	return `Set enabled state of Direct Console User Interface (DCUI TTY2).

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable DCUI
govc vcsa.access.dcui.set -enabled=true

# Disable DCUI
govc vcsa.access.dcui.set -enabled=false`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := dcui.NewManager(c)

	input := dcui.Access{Enabled: cmd.enabled}
	if err = m.Set(ctx, input); err != nil {
		return err
	}

	return nil
}
