// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/vcenter"
)

type sync struct {
	*flags.FolderFlag
	*flags.ResourcePoolFlag

	force bool
	vmtx  string
}

func init() {
	cli.Register("library.sync", &sync{})
}

func (cmd *sync) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	f.BoolVar(&cmd.force, "f", false, "Forcefully synchronize file content")
	f.StringVar(&cmd.vmtx, "vmtx", "", "Sync subscribed library to local library as VM Templates")
}

func (cmd *sync) Process(ctx context.Context) error {
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.ResourcePoolFlag.Process(ctx)
}

func (cmd *sync) Description() string {
	return `Sync library NAME or ITEM.

Examples:
  govc library.sync subscribed-library
  govc library.sync subscribed-library/item
  govc library.sync -vmtx local-library subscribed-library # convert subscribed OVFs to local VMTX`
}

func (cmd *sync) Usage() string {
	return "NAME|ITEM"
}

func (cmd *sync) syncVMTX(ctx context.Context, m *library.Manager, src library.Library, dst library.Library, items ...library.Item) error {
	if cmd.vmtx == "" {
		return nil
	}

	pool, err := cmd.ResourcePool()
	if err != nil {
		return err
	}

	folder, err := cmd.Folder()
	if err != nil {
		return err
	}

	l := vcenter.TemplateLibrary{
		Source:      src,
		Destination: dst,
		Placement: vcenter.Target{
			FolderID:       folder.Reference().Value,
			ResourcePoolID: pool.Reference().Value,
		},
		Include: func(item library.Item, current *library.Item) bool {
			fmt.Printf("Syncing /%s/%s to /%s/%s...", src.Name, item.Name, dst.Name, item.Name)
			if current == nil {
				fmt.Println()
				return true
			}
			fmt.Println("already exists.")
			return false
		},
	}

	return vcenter.NewManager(m.Client).SyncTemplateLibrary(ctx, l, items...)
}

func (cmd *sync) shouldSync(l library.Library) bool {
	if cmd.vmtx == "" {
		return true

	}
	// Allow library.sync -vmtx of LOCAL or SUBSCRIBED library
	return l.Type == "SUBSCRIBED"
}

func (cmd *sync) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	path := f.Arg(0)

	c, err := cmd.FolderFlag.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)

	var local library.Library
	if cmd.vmtx != "" {
		l, err := flags.ContentLibrary(ctx, c, cmd.vmtx)
		if err != nil {
			return err
		}
		local = *l
	}

	res, err := flags.ContentLibraryResult(ctx, c, "", path)
	if err != nil {
		return err
	}

	fmt.Printf("Syncing %s...\n", path)

	switch t := res.GetResult().(type) {
	case library.Library:
		if cmd.shouldSync(t) {
			if err = m.SyncLibrary(ctx, &t); err != nil {
				return err
			}
		}
		return cmd.syncVMTX(ctx, m, t, local)
	case library.Item:
		lib := res.GetParent().GetResult().(library.Library)
		if cmd.shouldSync(lib) {
			if err = m.SyncLibraryItem(ctx, &t, cmd.force); err != nil {
				return err
			}
		}
		return cmd.syncVMTX(ctx, m, lib, local, t)
	default:
		return fmt.Errorf("%q is a %T", res.GetPath(), t)
	}
}
