// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package maintenance

import (
	"context"
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

var vsanModes = types.VsanHostDecommissionModeObjectAction("").Strings()

type enter struct {
	*flags.HostSystemFlag

	timeout  int32
	evacuate bool
	vsanMode string
}

func init() {
	cli.Register("host.maintenance.enter", &enter{})
}

func (cmd *enter) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.Var(flags.NewInt32(&cmd.timeout), "timeout", "Timeout")
	f.BoolVar(&cmd.evacuate, "evacuate", false, "Evacuate powered off VMs")
	f.StringVar(&cmd.vsanMode, "vsan-mode", "", "vSAN data migration mode: "+strings.Join(vsanModes, ", "))
}

func (cmd *enter) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}

	if cmd.vsanMode != "" && !slices.Contains(vsanModes, cmd.vsanMode) {
		return fmt.Errorf("invalid -vsan-mode %q, must be one of: %s",
			cmd.vsanMode, strings.Join(vsanModes, ", "))
	}

	return nil
}

func (cmd *enter) Usage() string {
	return "HOST..."
}

func (cmd *enter) Description() string {
	return `Put HOST in maintenance mode.

While this task is running and when the host is in maintenance mode,
no VMs can be powered on and no provisioning operations can be performed on the host.

The -vsan-mode flag controls vSAN data migration for hosts in a vSAN cluster.
The values match the vSphere Client options:
  evacuateAllData:           Full data migration
  ensureObjectAccessibility: Ensure accessibility (default when unset)
  noAction:                  No data migration

Examples:
  govc host.maintenance.enter -vsan-mode evacuateAllData host.example.com`
}

func (cmd *enter) EnterMaintenanceMode(ctx context.Context, host *object.HostSystem) error {
	var spec *types.HostMaintenanceSpec
	if cmd.vsanMode != "" {
		spec = &types.HostMaintenanceSpec{
			VsanMode: &types.VsanHostDecommissionMode{ObjectAction: cmd.vsanMode},
		}
	}

	task, err := host.EnterMaintenanceMode(ctx, cmd.timeout, cmd.evacuate, spec)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("%s entering maintenance mode... ", host.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *enter) Run(ctx context.Context, f *flag.FlagSet) error {
	hosts, err := cmd.HostSystems(f.Args())
	if err != nil {
		return err
	}

	for _, host := range hosts {
		err = cmd.EnterMaintenanceMode(ctx, host)
		if err != nil {
			return err
		}
	}

	return nil
}
