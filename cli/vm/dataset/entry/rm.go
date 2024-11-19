/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package entry

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	govc "github.com/vmware/govmomi/cli/vm/dataset"
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type rm struct {
	*flags.VirtualMachineFlag
	dataSet string
}

func init() {
	cli.Register("vm.dataset.entry.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.StringVar(&cmd.dataSet, "dataset", "", "Data set name or ID")
}

func (cmd *rm) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *rm) Usage() string {
	return "KEY"
}

func (cmd *rm) Description() string {
	return `Delete data set entry.

Examples:
  govc vm.dataset.entry.rm -vm $vm -dataset com.example.project2 somekey`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	entryKey := f.Arg(0)

	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}
	vmId := vm.Reference().Value

	if cmd.dataSet == "" {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	mgr := dataset.NewManager(c)
	dataSetId, err := govc.FindDataSetId(ctx, mgr, vmId, cmd.dataSet)
	if err != nil {
		return err
	}
	return mgr.DeleteEntry(ctx, vmId, dataSetId, entryKey)
}
