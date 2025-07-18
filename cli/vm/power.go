// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type power struct {
	*flags.ClientFlag
	*flags.SearchFlag

	On       bool
	Off      bool
	Reset    bool
	Reboot   bool
	Shutdown bool
	Standby  bool
	Suspend  bool
	Force    bool
	Multi    bool
	Wait     bool
}

func init() {
	cli.Register("vm.power", &power{})
}

func (cmd *power) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	f.BoolVar(&cmd.On, "on", false, "Power on")
	f.BoolVar(&cmd.Off, "off", false, "Power off")
	f.BoolVar(&cmd.Reset, "reset", false, "Power reset")
	f.BoolVar(&cmd.Suspend, "suspend", false, "Power suspend")
	f.BoolVar(&cmd.Reboot, "r", false, "Reboot guest")
	f.BoolVar(&cmd.Shutdown, "s", false, "Shutdown guest")
	f.BoolVar(&cmd.Standby, "standby", false, "Standby guest")
	f.BoolVar(&cmd.Force, "force", false, "Force (ignore state error and hard shutdown/reboot if tools unavailable)")
	f.BoolVar(&cmd.Multi, "M", false, "Use Datacenter.PowerOnMultiVM method instead of VirtualMachine.PowerOnVM")
	f.BoolVar(&cmd.Wait, "wait", true, "Wait for the operation to complete")
}

func (cmd *power) Usage() string {
	return "NAME..."
}

func (cmd *power) Description() string {
	return `Invoke VM power operations.

Examples:
  govc vm.power -on VM1 VM2 VM3
  govc vm.power -on -M VM1 VM2 VM3
  govc vm.power -off -force VM1`
}

func (cmd *power) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.SearchFlag.Process(ctx); err != nil {
		return err
	}
	opts := []bool{cmd.On, cmd.Off, cmd.Reset, cmd.Suspend, cmd.Reboot, cmd.Shutdown, cmd.Standby}
	selected := false

	for _, opt := range opts {
		if opt {
			if selected {
				return flag.ErrHelp
			}
			selected = opt
		}
	}

	if !selected {
		return flag.ErrHelp
	}

	return nil
}

func isToolsUnavailable(err error) bool {
	return fault.Is(err, &types.ToolsUnavailable{})
}

// this is annoying, but the likely use cases for Datacenter.PowerOnVM outside of this command would
// use []types.ManagedObjectReference via ContainerView or field such as ResourcePool.Vm rather than the Finder.
func vmReferences(vms []*object.VirtualMachine) []types.ManagedObjectReference {
	refs := make([]types.ManagedObjectReference, len(vms))
	for i, vm := range vms {
		refs[i] = vm.Reference()
	}
	return refs
}

func (cmd *power) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	if cmd.On && cmd.Multi {
		dc, derr := cmd.Datacenter()
		if derr != nil {
			return derr
		}

		task, derr := dc.PowerOnVM(ctx, vmReferences(vms))
		if derr != nil {
			return derr
		}

		msg := fmt.Sprintf("Powering on %d VMs...", len(vms))
		if task == nil {
			// running against ESX
			fmt.Fprintf(cmd, "%s OK\n", msg)
			return nil
		}

		if cmd.Wait {
			logger := cmd.ProgressLogger(msg)
			defer logger.Wait()

			_, err = task.WaitForResult(ctx, logger)
			return err
		}
	}

	for _, vm := range vms {
		var task *object.Task

		switch {
		case cmd.On:
			fmt.Fprintf(cmd, "Powering on %s... ", vm.Reference())
			task, err = vm.PowerOn(ctx)
		case cmd.Off:
			fmt.Fprintf(cmd, "Powering off %s... ", vm.Reference())
			task, err = vm.PowerOff(ctx)
		case cmd.Reset:
			fmt.Fprintf(cmd, "Reset %s... ", vm.Reference())
			task, err = vm.Reset(ctx)
		case cmd.Suspend:
			fmt.Fprintf(cmd, "Suspend %s... ", vm.Reference())
			task, err = vm.Suspend(ctx)
		case cmd.Reboot:
			fmt.Fprintf(cmd, "Reboot guest %s... ", vm.Reference())
			err = vm.RebootGuest(ctx)

			if err != nil && cmd.Force && isToolsUnavailable(err) {
				task, err = vm.Reset(ctx)
			}
		case cmd.Shutdown:
			fmt.Fprintf(cmd, "Shutdown guest %s... ", vm.Reference())
			err = vm.ShutdownGuest(ctx)

			if err != nil && cmd.Force && isToolsUnavailable(err) {
				task, err = vm.PowerOff(ctx)
			}
		case cmd.Standby:
			fmt.Fprintf(cmd, "Standby guest %s... ", vm.Reference())
			err = vm.StandbyGuest(ctx)

			if err != nil && cmd.Force && isToolsUnavailable(err) {
				task, err = vm.Suspend(ctx)
			}
		}

		if err != nil {
			return err
		}

		if cmd.Wait && task != nil {
			err = task.Wait(ctx)
		}
		if err == nil {
			fmt.Fprintf(cmd, "OK\n")
			continue
		}

		if cmd.Force {
			fmt.Fprintf(cmd, "Error: %s\n", err)
			continue
		}

		return err
	}

	return nil
}
