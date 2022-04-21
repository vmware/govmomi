/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package module

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/cluster"
	"github.com/vmware/govmomi/vim25/mo"
)

type rmVMs struct {
	*flags.SearchFlag
	moduleID string
}

func init() {
	cli.Register("cluster.module.vm.rm", &rmVMs{})
}

func (cmd *rmVMs) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.SearchFlag, ctx = flags.NewSearchFlag(ctx, flags.SearchVirtualMachines)
	cmd.SearchFlag.Register(ctx, f)

	f.StringVar(&cmd.moduleID, "id", "", "Module ID")
}

func (cmd *rmVMs) Process(ctx context.Context) error {
	return cmd.SearchFlag.Process(ctx)
}

func (cmd *rmVMs) Usage() string {
	return `VM...`
}

func (cmd *rmVMs) Description() string {
	return `Remove VM(s) from a cluster module.

Examples:
  govc cluster.module.vm.rm -id module_id $vm...`
}

func (cmd *rmVMs) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() < 1 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		return err
	}

	refs := make([]mo.Reference, 0, len(vms))
	for _, vm := range vms {
		refs = append(refs, vm.Reference())
	}

	allRemoved, err := cluster.NewManager(c).RemoveModuleMembers(ctx, cmd.moduleID, refs...)
	if err != nil {
		return err
	}

	if !allRemoved {
		return fmt.Errorf("a VM was not a member of the module")
	}

	return nil
}
