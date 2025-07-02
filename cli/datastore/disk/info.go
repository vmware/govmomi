// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type info struct {
	*flags.DatastoreFlag

	c bool
	d bool
	p bool

	uuid bool
}

func init() {
	cli.Register("datastore.disk.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.BoolVar(&cmd.c, "c", false, "Chain format")
	f.BoolVar(&cmd.d, "d", false, "Include datastore in output")
	f.BoolVar(&cmd.p, "p", true, "Include parents")
	f.BoolVar(&cmd.uuid, "uuid", false, "Include disk UUID")
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Usage() string {
	return "VMDK"
}

func (cmd *info) Description() string {
	return `Query VMDK info on DS.

Examples:
  govc datastore.disk.info disks/disk1.vmdk`
}

func fullPath(s string) string {
	return s
}

func dsPath(s string) string {
	var p object.DatastorePath

	if p.FromString(s) {
		return p.Path
	}

	return s
}

var infoPath = dsPath

var queryUUID func(string) string

type infoResult []object.VirtualDiskInfo

func (r infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, info := range r {
		fmt.Fprintf(tw, "Name:\t%s\n", infoPath(info.Name))
		if queryUUID != nil {
			fmt.Fprintf(tw, "  UUID:\t%s\n", queryUUID(info.Name))
		}
		fmt.Fprintf(tw, "  Type:\t%s\n", info.DiskType)
		fmt.Fprintf(tw, "  Parent:\t%s\n", infoPath(info.Parent))
	}

	return tw.Flush()
}

type chainResult []object.VirtualDiskInfo

func (r chainResult) Write(w io.Writer) error {
	for i, info := range r {
		fmt.Fprint(w, strings.Repeat(" ", i*2))
		fmt.Fprintln(w, infoPath(info.Name))
	}

	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	dc, err := cmd.Datacenter()
	if err != nil {
		return err
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	m := object.NewVirtualDiskManager(ds.Client())

	if cmd.uuid {
		queryUUID = func(name string) string {
			id, _ := m.QueryVirtualDiskUuid(ctx, name, dc)
			return id
		}
	}

	info, err := m.QueryVirtualDiskInfo(ctx, ds.Path(f.Arg(0)), dc, cmd.p)
	if err != nil {
		return err
	}

	if cmd.d {
		infoPath = fullPath
	}

	var r flags.OutputWriter = infoResult(info)

	if cmd.c {
		r = chainResult(info)
	}

	return cmd.WriteResult(r)
}
