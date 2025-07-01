// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ssh

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/appliance/access/ssh"
)

type set struct {
	*flags.ClientFlag
	enabled bool
}

func init() {
	cli.Register("vcsa.access.ssh.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.enabled,
		"enabled",
		false,
		"Enable SSH-based controlled CLI.")
}

func (cmd *set) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *set) Description() string {
	return `Set enabled state of the SSH-based controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable SSH
govc vcsa.access.ssh.set -enabled=true

# Disable SSH
govc vcsa.access.ssh.set -enabled=false`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := ssh.NewManager(c)

	input := ssh.Access{Enabled: cmd.enabled}
	if err = m.Set(ctx, input); err != nil {
		return err
	}

	return nil
}
