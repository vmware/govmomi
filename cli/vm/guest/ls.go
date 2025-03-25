// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*GuestFlag

	simple bool
}

func init() {
	cli.Register("guest.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.BoolVar(&cmd.simple, "s", false, "Simple path only listing") // sadly we used '-l' for guest login
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return "PATH"
}

func (cmd *ls) Description() string {
	return `List PATH files in VM.

Examples:
  govc guest.ls -vm $name /tmp`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := cmd.FileManager()
	if err != nil {
		return err
	}

	var offset int32
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for {
		info, err := m.ListFiles(ctx, cmd.Auth(), f.Arg(0), offset, 0, f.Arg(1))
		if err != nil {
			return err
		}

		for _, f := range info.Files {
			if cmd.simple {
				fmt.Fprintln(tw, f.Path)
				continue
			}

			var kind byte

			switch types.GuestFileType(f.Type) {
			case types.GuestFileTypeDirectory:
				kind = 'd'
				if f.Size == 0 {
					f.Size = 4092
				}
			case types.GuestFileTypeSymlink:
				kind = 'l'
			case types.GuestFileTypeFile:
				kind = '-'
			}

			switch x := f.Attributes.(type) {
			case *types.GuestPosixFileAttributes:
				perm := os.FileMode(x.Permissions).Perm().String()[1:]
				fmt.Fprintf(tw, "%c%s\t%d\t%d\t", kind, perm, *x.OwnerId, *x.GroupId)
			}

			attr := f.Attributes.GetGuestFileAttributes()

			fmt.Fprintf(tw, "%s\t%s\t%s", units.FileSize(f.Size), attr.ModificationTime.Format("Jan 2 15:04 2006"), f.Path)
			if attr.SymlinkTarget != "" {
				fmt.Fprintf(tw, " -> %s", attr.SymlinkTarget)
			}
			fmt.Fprintln(tw)
		}

		err = tw.Flush()
		if err != nil {
			return err
		}

		if info.Remaining == 0 {
			break
		}
		offset += int32(len(info.Files))
	}

	return nil
}
