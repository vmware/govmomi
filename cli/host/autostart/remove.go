// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package autostart

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type remove struct {
	*AutostartFlag
}

func init() {
	cli.Register("host.autostart.remove", &remove{})
}

func (cmd *remove) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.AutostartFlag, ctx = newAutostartFlag(ctx)
	cmd.AutostartFlag.Register(ctx, f)
}

func (cmd *remove) Process(ctx context.Context) error {
	if err := cmd.AutostartFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *remove) Usage() string {
	return "VM..."
}

func (cmd *remove) Run(ctx context.Context, f *flag.FlagSet) error {
	var powerInfo = types.AutoStartPowerInfo{
		StartAction:      "none",
		StartDelay:       -1,
		StartOrder:       -1,
		StopAction:       "none",
		StopDelay:        -1,
		WaitForHeartbeat: types.AutoStartWaitHeartbeatSettingSystemDefault,
	}

	return cmd.ReconfigureVMs(f.Args(), powerInfo)
}
