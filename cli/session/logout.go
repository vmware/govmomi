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

type logout struct {
	*flags.ClientFlag

	vapi bool
}

func init() {
	cli.Register("session.logout", &logout{})
}

func (cmd *logout) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.vapi, "r", false, "REST logout")
}

func (cmd *logout) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *logout) Description() string {
	return `Logout the current session.

By default, govc commands persist sessions and do not logout unless '-persist-session=false' is set.
The session.logout command can be used to end the current persisted session.
The session.rm command can be used to remove sessions other than the current session.

Examples:
  govc session.logout`
}

func (cmd *logout) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	err = session.NewManager(c).Logout(ctx)

	if cmd.vapi {
		rc, err := cmd.RestClient()
		if err != nil {
			return err
		}
		return rc.Logout(ctx)
	}

	return err
}
