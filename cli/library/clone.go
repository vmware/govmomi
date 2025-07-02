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
	"github.com/vmware/govmomi/vapi/vcenter"
)

type clone struct {
	*flags.ClusterFlag
	*flags.DatastoreFlag
	*flags.FolderFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.VirtualMachineFlag
	*flags.StorageProfileFlag

	ovf   bool
	extra bool
	mac   bool
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

	cmd.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx)
	cmd.StorageProfileFlag.Register(ctx, f)

	f.BoolVar(&cmd.ovf, "ovf", false, "Clone as OVF (default is VM Template)")
	f.BoolVar(&cmd.extra, "e", false, "Include extra configuration")
	f.BoolVar(&cmd.mac, "m", false, "Preserve MAC-addresses on network adapters")
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
	if err := cmd.StorageProfileFlag.Process(ctx); err != nil {
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
	if vm == nil || name == "" {
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
		if cmd.extra {
			ovf.Spec.Flags = append(ovf.Spec.Flags, "EXTRA_CONFIG")
		}
		if cmd.mac {
			ovf.Spec.Flags = append(ovf.Spec.Flags, "PRESERVE_MAC")
		}
		id, err := vcenter.NewManager(c).CreateOVF(ctx, ovf)
		if err != nil {
			return err
		}
		fmt.Println(id)
		return nil
	}

	profile, err := cmd.StorageProfile(ctx)
	if err != nil {
		return err
	}

	storage := &vcenter.DiskStorage{
		Datastore: dsID,
		StoragePolicy: &vcenter.StoragePolicy{
			Policy: profile,
			Type:   "USE_SOURCE_POLICY",
		},
	}
	if profile != "" {
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
