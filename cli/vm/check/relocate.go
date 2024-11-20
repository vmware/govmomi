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

package check

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type relocate struct {
	checkFlag
}

func init() {
	cli.Register("vm.check.relocate", &relocate{}, true)
}

func (cmd *relocate) Description() string {
	return `Check if VM can be relocated.

Examples:
  govc vm.migrate -spec my-vm | govc vm.check.relocate -vm my-vm`
}

func (cmd *relocate) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var spec types.VirtualMachineRelocateSpec
	if err := cmd.Spec(&spec); err != nil {
		return err
	}

	checker, err := cmd.provChecker()
	if err != nil {
		return err
	}

	res, err := checker.CheckRelocate(ctx, vm.Reference(), spec, cmd.testTypes...)
	if err != nil {
		return err
	}

	return cmd.result(ctx, res)
}
