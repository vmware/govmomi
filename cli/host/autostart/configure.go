// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package autostart

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type configure struct {
	*AutostartFlag

	types.AutoStartDefaults
}

func init() {
	cli.Register("host.autostart.configure", &configure{})
}

func (cmd *configure) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.AutostartFlag, ctx = newAutostartFlag(ctx)
	cmd.AutostartFlag.Register(ctx, f)

	f.Var(flags.NewOptionalBool(&cmd.Enabled), "enabled", "Enable autostart")
	f.Var(flags.NewInt32(&cmd.StartDelay), "start-delay", "Start delay")
	f.StringVar(&cmd.StopAction, "stop-action", "", "Stop action")
	f.Var(flags.NewInt32(&cmd.StopDelay), "stop-delay", "Stop delay")
	f.Var(flags.NewOptionalBool(&cmd.WaitForHeartbeat), "wait-for-heartbeat", "Wait for hearbeat")
}

func (cmd *configure) Process(ctx context.Context) error {
	if err := cmd.AutostartFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *configure) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.ReconfigureDefaults(cmd.AutoStartDefaults)
}
