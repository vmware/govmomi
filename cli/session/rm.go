// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package session

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/session"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("session.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "KEY..."
}

func (cmd *rm) Description() string {
	return `Remove active sessions.

Examples:
  govc session.ls | grep root
  govc session.rm 5279e245-e6f1-4533-4455-eb94353b213a`
}

func (cmd *rm) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	return session.NewManager(c).TerminateSession(ctx, f.Args())
}
