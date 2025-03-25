// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
)

type getenv struct {
	*GuestFlag
}

func init() {
	cli.Register("guest.getenv", &getenv{})
}

func (cmd *getenv) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestProcessFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)
}

func (cmd *getenv) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *getenv) Usage() string {
	return "[NAME]..."
}

func (cmd *getenv) Description() string {
	return `Read NAME environment variables from VM.

Examples:
  govc guest.getenv -vm $name
  govc guest.getenv -vm $name HOME`
}

func (cmd *getenv) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.ProcessManager()
	if err != nil {
		return err
	}

	vars, err := m.ReadEnvironmentVariable(ctx, cmd.Auth(), f.Args())
	if err != nil {
		return err
	}

	for _, v := range vars {
		fmt.Printf("%s\n", v)
	}

	return nil
}
