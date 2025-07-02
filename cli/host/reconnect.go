// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package host

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type reconnect struct {
	*flags.HostSystemFlag
	*flags.HostConnectFlag

	types.HostSystemReconnectSpec
}

func init() {
	cli.Register("host.reconnect", &reconnect{})
}

func (cmd *reconnect) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.HostConnectFlag, ctx = flags.NewHostConnectFlag(ctx)
	cmd.HostConnectFlag.Register(ctx, f)

	cmd.HostSystemReconnectSpec.SyncState = types.NewBool(false)
	f.BoolVar(cmd.HostSystemReconnectSpec.SyncState, "sync-state", false, "Sync state")
}

func (cmd *reconnect) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostConnectFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *reconnect) Description() string {
	return `Reconnect HOST to vCenter.

This command can also be used to change connection properties (hostname, fingerprint, username, password),
without disconnecting the host.`
}

func (cmd *reconnect) Reconnect(ctx context.Context, host *object.HostSystem) error {
	task, err := host.Reconnect(ctx, &cmd.HostConnectSpec, &cmd.HostSystemReconnectSpec)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s reconnecting... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *reconnect) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.Reconnect(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
