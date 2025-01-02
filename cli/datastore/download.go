// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datastore

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vmdk"
)

type download struct {
	*flags.DatastoreFlag
	*flags.HostSystemFlag
}

func init() {
	cli.Register("datastore.download", &download{})
}

func (cmd *download) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
}

func (cmd *download) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *download) Usage() string {
	return "SOURCE DEST"
}

func (cmd *download) Description() string {
	return `Copy SOURCE from DS to DEST on the local system.

If DEST name is "-", source is written to stdout.

Examples:
  govc datastore.download vm-name/vmware.log ./local.log
  govc datastore.download vm-name/vmware.log - | grep -i error
  govc datastore.download -json vm-name/vm-name.vmdk - | jq .ddb
  ovf=$(govc library.info -l -L vmservice/photon-5.0/*.ovf)
  govc datastore.download -json "$ovf" - | jq -r .diskSection.disk[].capacity`
}

func (cmd *download) Run(ctx context.Context, f *flag.FlagSet) error {
	args := f.Args()
	if len(args) != 2 {
		return errors.New("invalid arguments")
	}

	src := args[0]
	dst := args[1]

	var dp object.DatastorePath
	if dp.FromString(src) {
		// e.g. `govc library.info -l -L ...`
		cmd.DatastoreFlag.Name = dp.Datastore
		src = dp.Path
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	h, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	var via string

	if h != nil {
		via = fmt.Sprintf(" via %s", h.InventoryPath)
		ctx = ds.HostContext(ctx, h)
	}

	p := soap.DefaultDownload

	if dst == "-" {
		f, _, err := ds.Download(ctx, src, &p)
		if err != nil {
			return err
		}

		if cmd.DatastoreFlag.All() {
			switch path.Ext(src) {
			case ".vmdk":
				data, err := vmdk.ParseDescriptor(f)
				if err != nil {
					return err
				}
				return cmd.DatastoreFlag.WriteResult(data)
			case ".ovf":
				data, err := ovf.Unmarshal(f)
				if err != nil {
					return err
				}
				return cmd.DatastoreFlag.WriteResult(data)
			}
		}

		_, err = io.Copy(os.Stdout, f)
		return err
	}

	if cmd.DatastoreFlag.OutputFlag.TTY {
		logger := cmd.DatastoreFlag.ProgressLogger(fmt.Sprintf("Downloading%s... ", via))
		p.Progress = logger
		defer logger.Wait()
	}

	return ds.DownloadFile(ctx, src, dst, &p)
}
