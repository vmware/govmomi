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

type set struct {
	*flags.VirtualMachineFlag
	dataSet string
}

func init() {
	cli.Register("vm.dataset.entry.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	f.StringVar(&cmd.dataSet, "dataset", "", "Data set name or ID")
}

func (cmd *set) Process(ctx context.Context) error {
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *set) Usage() string {
	return "KEY VALUE"
}

func (cmd *set) Description() string {
	return `Set the value of a data set entry.

Examples:
  govc vm.dataset.entry.set -vm $vm -dataset com.example.project2 somekey somevalue`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}
	entryKey := f.Arg(0)
	entryValue := f.Arg(1)

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
	return mgr.SetEntry(ctx, vmId, dataSetId, entryKey, entryValue)
}
