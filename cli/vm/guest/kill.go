// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
)

type kill struct {
	*GuestFlag

	pids pidSelector
}

func init() {
	cli.Register("guest.kill", &kill{})
}

func (cmd *kill) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestProcessFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.Var(&cmd.pids, "p", "Process ID")
}

func (cmd *kill) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *kill) Description() string {
	return `Kill process ID on VM.

Examples:
  govc guest.kill -vm $name -p 12345`
}

func (cmd *kill) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.ProcessManager()
	if err != nil {
		return err
	}

	for _, pid := range cmd.pids {
		if err := m.TerminateProcess(ctx, cmd.Auth(), pid); err != nil {
			return err
		}
	}

	return nil
}
