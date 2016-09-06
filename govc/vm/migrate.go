/*
Copyright (c) 2016 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vm

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type migrate struct {
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.SearchFlag

	priority types.VirtualMachineMovePriority
	state    types.VirtualMachinePowerState
}

func init() {
	cli.Register("vm.migrate", &migrate{})
}

func (cmd *migrate) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.StringVar((*string)(&cmd.priority), "priority", string(types.VirtualMachineMovePriorityDefaultPriority), "The task priority")
	f.StringVar((*string)(&cmd.state), "state", "", "If specified, the VM migrates only if its state matches")
}

func (cmd *migrate) Process(ctx context.Context) error {
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *migrate) Usage() string {
	return "VM..."
}

func (cmd *migrate) Description() string {
	return `Migrates VM execution to a specific resource pool or host.
Example:
govc vm.migrate -host another-host vm-1 vm-2 vm-3
`
}

func (cmd *migrate) Run(ctx context.Context, f *flag.FlagSet) error {
	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemFlag.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	pool, err := cmd.ResourcePoolFlag.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}

	for _, vm := range vms {
		task, err := vm.Migrate(ctx, pool, host, cmd.priority, cmd.state)
		if err != nil {
			return err
		}

		err = task.Wait(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
