/*
Copyright (c) 2014-2024 VMware, Inc. All Rights Reserved.

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

package importx

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf/importer"
)

type ovfx struct {
	*flags.DatastoreFlag
	*flags.HostSystemFlag
	*flags.OutputFlag
	*flags.ResourcePoolFlag
	*flags.FolderFlag

	*OptionsFlag

	Importer importer.Importer
}

func init() {
	cli.Register("import.ovf", &ovfx{})
}

func (cmd *ovfx) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.OptionsFlag, ctx = newOptionsFlag(ctx)
	cmd.OptionsFlag.Register(ctx, f)

	f.StringVar(&cmd.Importer.Name, "name", "", "Name to use for new entity")
	f.BoolVar(&cmd.Importer.VerifyManifest, "m", false, "Verify checksum of uploaded files against manifest (.mf)")
	f.BoolVar(&cmd.Importer.Hidden, "hidden", false, "Enable hidden properties")
}

func (cmd *ovfx) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OptionsFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ovfx) Usage() string {
	return "PATH_TO_OVF"
}

func (cmd *ovfx) Run(ctx context.Context, f *flag.FlagSet) error {
	fpath, err := cmd.Prepare(f)
	if err != nil {
		return err
	}

	archive := &importer.FileArchive{Path: fpath}
	archive.Client = cmd.Importer.Client

	cmd.Importer.Archive = archive

	moref, err := cmd.Importer.Import(context.TODO(), fpath, cmd.Options)
	if err != nil {
		return err
	}

	vm := object.NewVirtualMachine(cmd.Importer.Client, *moref)
	return cmd.Deploy(vm, cmd.OutputFlag)
}

func (cmd *ovfx) Prepare(f *flag.FlagSet) (string, error) {
	var err error

	args := f.Args()
	if len(args) != 1 {
		return "", errors.New("no file specified")
	}

	cmd.Importer.Log = cmd.OutputFlag.Log
	cmd.Importer.Client, err = cmd.DatastoreFlag.Client()
	if err != nil {
		return "", err
	}

	cmd.Importer.Datacenter, err = cmd.DatastoreFlag.Datacenter()
	if err != nil {
		return "", err
	}

	cmd.Importer.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return "", err
	}

	cmd.Importer.ResourcePool, err = cmd.ResourcePoolIfSpecified()
	if err != nil {
		return "", err
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return "", err
	}

	if cmd.Importer.ResourcePool == nil {
		if host == nil {
			cmd.Importer.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool()
		} else {
			cmd.Importer.ResourcePool, err = host.ResourcePool(context.TODO())
		}
		if err != nil {
			return "", err
		}
	}

	cmd.Importer.Finder, err = cmd.DatastoreFlag.Finder()
	if err != nil {
		return "", err
	}

	cmd.Importer.Host, err = cmd.HostSystemIfSpecified()
	if err != nil {
		return "", err
	}

	// The folder argument must not be set on a VM in a vApp, otherwise causes
	// InvalidArgument fault: A specified parameter was not correct: pool
	if cmd.Importer.ResourcePool.Reference().Type != "VirtualApp" {
		cmd.Importer.Folder, err = cmd.FolderOrDefault("vm")
		if err != nil {
			return "", err
		}
	}

	if cmd.Importer.Name == "" {
		// Override name from options if specified
		if cmd.Options.Name != nil {
			cmd.Importer.Name = *cmd.Options.Name
		}
	} else {
		cmd.Options.Name = &cmd.Importer.Name
	}

	return f.Arg(0), nil
}
