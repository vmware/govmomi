/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package namespace

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi"
	"github.com/vmware/govmomi/vapi/namespace"
)

type registervm struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("namespace.registervm", &registervm{})
}

func (cmd *registervm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *registervm) Usage() string {
	return "NAME"
}

func (cmd *registervm) Description() string {
	return `Register an existing virtual machine as VM Service managed VM.

Examples:
  govc namespace.registervm -vm my-vm my-namespace`
}

func (cmd *registervm) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil || f.NArg() != 1 {
		return flag.ErrHelp
	}

	rc, err := cmd.RestClient()
	if err != nil {
		return err
	}

	spec := namespace.RegisterVMSpec{VM: vm.Reference().Value}

	id, err := namespace.NewManager(rc).RegisterVM(ctx, f.Arg(0), spec)
	if err != nil {
		return err
	}

	task := object.NewTask(vm.Client(), vapi.Task(id))

	logger := cmd.ProgressLogger(fmt.Sprintf("registervm %s... ", vm.InventoryPath))
	_, err = task.WaitForResult(ctx, logger)
	logger.Wait()

	return err
}
