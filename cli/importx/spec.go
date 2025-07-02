// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importx

import (
	"context"
	"flag"
	"fmt"
	"io"
	"path"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/ovf/importer"
)

type spec struct {
	*flags.ClientFlag
	*flags.OutputFlag

	Archive importer.Archive

	hidden bool
}

func init() {
	cli.Register("import.spec", &spec{})
}

func (cmd *spec) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.hidden, "hidden", false, "Enable hidden properties")
}

func (cmd *spec) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *spec) Usage() string {
	return "PATH_TO_OVF_OR_OVA"
}

func (cmd *spec) Run(ctx context.Context, f *flag.FlagSet) error {
	fpath := ""
	args := f.Args()
	if len(args) == 1 {
		fpath = f.Arg(0)
	}

	if len(fpath) > 0 {
		switch path.Ext(fpath) {
		case ".ovf":
			cmd.Archive = &importer.FileArchive{Path: fpath}
		case "", ".ova":
			cmd.Archive = &importer.TapeArchive{Path: fpath}
			fpath = "*.ovf"
		default:
			return fmt.Errorf("invalid file extension %s", path.Ext(fpath))
		}

		if importer.IsRemotePath(f.Arg(0)) {
			client, err := cmd.Client()
			if err != nil {
				return err
			}
			switch archive := cmd.Archive.(type) {
			case *importer.FileArchive:
				archive.Client = client
			case *importer.TapeArchive:
				archive.Client = client
			}
		}
	}

	env, err := importer.Spec(fpath, cmd.Archive, cmd.hidden, cmd.Verbose())
	if err != nil {
		return err
	}

	if !cmd.All() {
		cmd.JSON = true
	}
	return cmd.WriteResult(&specResult{env})
}

type specResult struct {
	*importer.Options
}

func (*specResult) Write(w io.Writer) error {
	return nil
}
