// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/nfc"
	"github.com/vmware/govmomi/vim25/soap"
)

type export struct {
	*flags.VirtualMachineFlag

	lease bool
	dest  string
}

func init() {
	cli.Register("snapshot.export", &export{})
}

func (cmd *export) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.lease, "lease", false, "Output NFC Lease only")
	f.StringVar(&cmd.dest, "d", ".", "Destination directory")
}

func (cmd *export) Usage() string {
	return "NAME"
}

func (cmd *export) Description() string {
	return `Export snapshot of VM with given NAME.

NAME can be the snapshot name, tree path, or managed object ID.

Examples:
  govc snapshot.export -vm my-vm my-snapshot
  govc snapshot.export -vm my-vm -lease my-snapshot`
}

func (cmd *export) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 && !cmd.lease {
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	s, err := vm.FindSnapshot(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	lease, err := vm.ExportSnapshot(ctx, s)
	if err != nil {
		return err
	}

	info, err := lease.Wait(ctx, nil)
	if err != nil {
		return err
	}

	if cmd.lease {
		o, err := lease.Properties(ctx)
		if err != nil {
			return err
		}

		return cmd.WriteResult(o)
	}

	u := lease.StartUpdater(ctx, info)
	defer u.Done()

	for _, i := range info.Items {
		err := cmd.Download(ctx, lease, i)
		if err != nil {
			return err
		}
	}

	return lease.Complete(ctx)
}

func (cmd *export) Download(ctx context.Context, lease *nfc.Lease, item nfc.FileItem) error {
	path := filepath.Join(cmd.dest, item.Path)

	logger := cmd.ProgressLogger(fmt.Sprintf("Downloading %s... ", item.Path))
	defer logger.Wait()

	opts := soap.Download{
		Progress: logger,
	}

	return lease.DownloadFile(ctx, path, item, opts)
}
