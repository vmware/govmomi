// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"errors"
	"flag"
	"net/url"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/guest/toolbox"
	"github.com/vmware/govmomi/object"
)

type GuestFlag struct {
	*flags.ClientFlag
	*flags.VirtualMachineFlag

	*AuthFlag
}

func newGuestFlag(ctx context.Context) (*GuestFlag, context.Context) {
	f := &GuestFlag{}
	f.ClientFlag, ctx = flags.NewClientFlag(ctx)
	f.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	f.AuthFlag, ctx = newAuthFlag(ctx)
	return f, ctx
}

func newGuestProcessFlag(ctx context.Context) (*GuestFlag, context.Context) {
	f, gctx := newGuestFlag(ctx)
	f.proc = true
	return f, gctx
}

func (flag *GuestFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.ClientFlag.Register(ctx, f)
	flag.VirtualMachineFlag.Register(ctx, f)
	flag.AuthFlag.Register(ctx, f)
}

func (flag *GuestFlag) Process(ctx context.Context) error {
	if err := flag.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := flag.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := flag.AuthFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (flag *GuestFlag) Toolbox(ctx context.Context) (*toolbox.Client, error) {
	vm, err := flag.VirtualMachine()
	if err != nil {
		return nil, err
	}

	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	return toolbox.NewClient(ctx, c, vm, flag.Auth())
}

func (flag *GuestFlag) FileManager() (*guest.FileManager, error) {
	ctx := context.TODO()
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	vm, err := flag.VirtualMachine()
	if err != nil {
		return nil, err
	}

	o := guest.NewOperationsManager(c, vm.Reference())
	return o.FileManager(ctx)
}

func (flag *GuestFlag) ProcessManager() (*guest.ProcessManager, error) {
	ctx := context.TODO()
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	vm, err := flag.VirtualMachine()
	if err != nil {
		return nil, err
	}

	o := guest.NewOperationsManager(c, vm.Reference())
	return o.ProcessManager(ctx)
}

func (flag *GuestFlag) ParseURL(urlStr string) (*url.URL, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	return c.Client.ParseURL(urlStr)
}

func (flag *GuestFlag) VirtualMachine() (*object.VirtualMachine, error) {
	vm, err := flag.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return nil, err
	}
	if vm == nil {
		return nil, errors.New("no vm specified")
	}
	return vm, nil
}
