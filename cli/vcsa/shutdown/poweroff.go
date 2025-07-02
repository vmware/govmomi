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

type powerOff struct {
	*flags.ClientFlag

	reason string
	delay  int // in minutes
}

func init() {
	cli.Register("vcsa.shutdown.poweroff", &powerOff{})
}

func (cmd *powerOff) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.IntVar(&cmd.delay,
		"delay",
		0,
		"Minutes after which poweroff should start.")
}

func (cmd *powerOff) Usage() string {
	return "REASON"
}

func (cmd *powerOff) Description() string {
	return `Power off the appliance.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.poweroff -delay 10 "powering off for maintenance"`
}

func (cmd *powerOff) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := shutdown.NewManager(c)

	if err = m.PowerOff(ctx, f.Arg(0), cmd.delay); err != nil {
		return err
	}

	return nil
}
