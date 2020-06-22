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

type clone struct {
	*flags.ClusterFlag
	*flags.DatastoreFlag
	*flags.FolderFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.VirtualMachineFlag

	profile string
	ovf     bool
}

func init() {
	cli.Register("library.clone", &clone{})
}

func (cmd *clone) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.ovf, "ovf", false, "Clone as OVF (default is VM Template)")
	f.StringVar(&cmd.profile, "profile", "", "Storage profile")
}

func (cmd *clone) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.VirtualMachineFlag.Process(ctx)
}

func (cmd *clone) Usage() string {
	return "PATH NAME"
}

func (cmd *clone) Description() string {
	return `Clone VM to Content Library PATH.

By default, clone as a VM template (requires vCenter version 6.7U1 or higher).
Clone as an OVF when the '-ovf' flag is specified.

Examples:
  govc library.clone -vm template-vm my-content template-vm-item
  govc library.clone -ovf -vm template-vm my-content ovf-item`
}

func (cmd *clone) Run(ctx context.Context, f *flag.FlagSet) error {
	path := f.Arg(0)
	name := f.Arg(1)

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}
	ds, err := cmd.DatastoreIfSpecified()
	if err != nil {
		return err
	}
	folder, err := cmd.FolderOrDefault("vm")
	if err != nil {
		return err
	}
	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}
	cluster, err := cmd.ClusterIfSpecified()
	if err != nil {
		return err
	}
	pool, err := cmd.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}

	dsID := ""
	if ds != nil {
		dsID = ds.Reference().Value
	}

	c, err := cmd.FolderFlag.RestClient()
	if err != nil {
		return err
	}

	l, err := flags.ContentLibrary(ctx, c, path)
	if err != nil {
		return err
	}

	if cmd.ovf {
		ovf := vcenter.OVF{
			Spec: vcenter.CreateSpec{
				Name: name,
			},
			Source: vcenter.ResourceID{
				Value: vm.Reference().Value,
			},
			Target: vcenter.LibraryTarget{
				LibraryID: l.ID,
			},
		}
		id, err := vcenter.NewManager(c).CreateOVF(ctx, ovf)
		if err != nil {
			return err
		}
		fmt.Println(id)
		return nil
	}

	storage := &vcenter.DiskStorage{
		Datastore: dsID,
		StoragePolicy: &vcenter.StoragePolicy{
			Policy: cmd.profile,
			Type:   "USE_SOURCE_POLICY",
		},
	}
	if cmd.profile != "" {
		storage.StoragePolicy.Type = "USE_SPECIFIED_POLICY"
	}

	spec := vcenter.Template{
		Name:          name,
		Library:       l.ID,
		DiskStorage:   storage,
		VMHomeStorage: storage,
		SourceVM:      vm.Reference().Value,
		Placement: &vcenter.Placement{
			Folder: folder.Reference().Value,
		},
	}
	if pool != nil {
		spec.Placement.ResourcePool = pool.Reference().Value
	}
	if host != nil {
		spec.Placement.Host = host.Reference().Value
	}
	if cluster != nil {
		spec.Placement.Cluster = cluster.Reference().Value
	}

	id, err := vcenter.NewManager(c).CreateTemplate(ctx, spec)
	if err != nil {
		return err
	}

	fmt.Println(id)

	return nil
}
