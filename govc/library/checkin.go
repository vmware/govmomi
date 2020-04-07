/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package library

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/vcenter"
)

type checkin struct {
	*flags.VirtualMachineFlag
	vcenter.CheckIn
}

func init() {
	cli.Register("library.checkin", &checkin{})
}

func (cmd *checkin) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.Message, "m", "", "Check in message")
}

func (cmd *checkin) Usage() string {
	return "PATH"
}

func (cmd *checkin) Description() string {
	return `Check in VM to Content Library item PATH.

Examples:
  govc library.checkin -vm my-vm my-content/template-vm-item`
}

func (cmd *checkin) Run(ctx context.Context, f *flag.FlagSet) error {
	path := f.Arg(0)
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	l, err := flags.ContentLibraryItem(ctx, c, path)
	if err != nil {
		return err
	}

	version, err := vcenter.NewManager(c).CheckIn(ctx, l.ID, vm, &cmd.CheckIn)
	if err != nil {
		return err
	}

	fmt.Printf("%s (%s) checked in as content version %s", l.Name, l.ID, version)

	return nil
}
