// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package autostart

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*AutostartFlag
	// from types.AutoStartPowerInfo
	StartOrder       int32
	StartDelay       int32
	WaitForHeartbeat string
	StartAction      string
	StopDelay        int32
	StopAction       string
}

func init() {
	cli.Register("host.autostart.add", &add{})
}

var waitHeartbeatTypes = types.AutoStartWaitHeartbeatSetting("").Strings()

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.AutostartFlag, ctx = newAutostartFlag(ctx)
	cmd.AutostartFlag.Register(ctx, f)

	cmd.StartOrder = -1
	cmd.StartDelay = -1
	cmd.StopDelay = -1
	f.Var(flags.NewInt32(&cmd.StartOrder), "start-order", "Start Order")
	f.Var(flags.NewInt32(&cmd.StartDelay), "start-delay", "Start Delay")
	f.Var(flags.NewInt32(&cmd.StopDelay), "stop-delay", "Stop Delay")
	f.StringVar(&cmd.StartAction, "start-action", "powerOn", "Start Action")
	f.StringVar(&cmd.StopAction, "stop-action", "systemDefault", "Stop Action")
	f.StringVar(&cmd.WaitForHeartbeat, "wait", string(types.AutoStartWaitHeartbeatSettingSystemDefault),
		fmt.Sprintf("Wait for Hearbeat Setting (%s)", strings.Join(waitHeartbeatTypes, "|")))
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.AutostartFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Usage() string {
	return "VM..."
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	powerInfo := types.AutoStartPowerInfo{
		StartOrder:       cmd.StartOrder,
		StartDelay:       cmd.StartDelay,
		WaitForHeartbeat: types.AutoStartWaitHeartbeatSetting(cmd.WaitForHeartbeat),
		StartAction:      cmd.StartAction,
		StopDelay:        cmd.StopDelay,
		StopAction:       cmd.StopAction,
	}
	return cmd.ReconfigureVMs(f.Args(), powerInfo)
}
