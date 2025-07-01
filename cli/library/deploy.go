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
	"github.com/vmware/govmomi/cli/importx"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/vcenter"
	"github.com/vmware/govmomi/vim25/types"
)

type deploy struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.FolderFlag
	*flags.StorageProfileFlag
	*importx.OptionsFlag

	config string
}

func init() {
	cli.Register("library.deploy", &deploy{})
}

func (cmd *deploy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx)
	cmd.StorageProfileFlag.Register(ctx, f)

	cmd.OptionsFlag = new(importx.OptionsFlag)
	cmd.OptionsFlag.Register(ctx, f)

	if cli.ShowUnreleased() {
		f.StringVar(&cmd.config, "config", "", "VM config spec")
	}
}

func (cmd *deploy) Process(ctx context.Context) error {
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
	return cmd.OptionsFlag.Process(ctx)
}

func (cmd *deploy) Usage() string {
	return "TEMPLATE [NAME]"
}

func (cmd *deploy) Description() string {
	return `Deploy library OVF template.

Examples:
  govc library.deploy /library_name/ovf_template vm_name
  govc library.export /library_name/ovf_template/*.ovf # save local copy of .ovf
  govc import.spec *.ovf > deploy.json # generate options from .ovf
  # edit deploy.json as needed
  govc library.deploy -options deploy.json /library_name/ovf_template`
}

func (cmd *deploy) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() > 2 {
		return flag.ErrHelp
	}
	path := f.Arg(0)
	name := f.Arg(1)

	if name == "" && cmd.Options.Name != nil {
		name = *cmd.Options.Name
	}

	vc, err := cmd.DatastoreFlag.Client()
	if err != nil {
		return err
	}
	cmd.FolderFlag.KeepAlive(vc)

	c, err := cmd.DatastoreFlag.RestClient()
	if err != nil {
		return err
	}
	cmd.FolderFlag.KeepAlive(c)

	m := vcenter.NewManager(c)

	item, err := flags.ContentLibraryItem(ctx, c, path)
	if err != nil {
		return err
	}

	ds, err := cmd.DatastoreIfSpecified()
	if err != nil {
		return err
	}
	rp, err := cmd.ResourcePoolIfSpecified()
	if err != nil {
		return err
	}
	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}
	hostID := ""
	if rp == nil {
		if host == nil {
			rp, err = cmd.ResourcePoolFlag.ResourcePool()
		} else {
			rp, err = host.ResourcePool(ctx)
			hostID = host.Reference().Value
		}
		if err != nil {
			return err
		}
	}
	folder, err := cmd.Folder()
	if err != nil {
		return err
	}
	finder, err := cmd.FolderFlag.Finder(false)
	if err != nil {
		return err
	}

	var networks []vcenter.NetworkMapping
	for _, net := range cmd.Options.NetworkMapping {
		if net.Network == "" {
			continue
		}
		obj, err := finder.Network(ctx, net.Network)
		if err != nil {
			return err
		}
		networks = append(networks, vcenter.NetworkMapping{
			Key:   net.Name,
			Value: obj.Reference().Value,
		})
	}

	var properties []vcenter.Property
	for _, prop := range cmd.Options.PropertyMapping {
		properties = append(properties, vcenter.Property{
			ID:    prop.Key,
			Value: prop.Value,
		})
	}

	dsID := ""
	if ds != nil {
		dsID = ds.Reference().Value
	}

	cmd.FolderFlag.Log("Deploying library item...\n")

	var ref *types.ManagedObjectReference

	profile, err := cmd.StorageProfile(ctx)
	if err != nil {
		return err
	}

	switch item.Type {
	case library.ItemTypeOVF:
		deploy := vcenter.Deploy{
			DeploymentSpec: vcenter.DeploymentSpec{
				Name:               name,
				DefaultDatastoreID: dsID,
				AcceptAllEULA:      true,
				Annotation:         cmd.Options.Annotation,
				AdditionalParams: []vcenter.AdditionalParams{
					{
						Class:       vcenter.ClassDeploymentOptionParams,
						Type:        vcenter.TypeDeploymentOptionParams,
						SelectedKey: cmd.Options.Deployment,
					},
					{
						Class:       vcenter.ClassPropertyParams,
						Type:        vcenter.TypePropertyParams,
						SelectedKey: cmd.Options.Deployment,
						Properties:  properties,
					},
				},
				NetworkMappings:     networks,
				StorageProvisioning: cmd.Options.DiskProvisioning,
				StorageProfileID:    profile,
			},
			Target: vcenter.Target{
				ResourcePoolID: rp.Reference().Value,
				HostID:         hostID,
				FolderID:       folder.Reference().Value,
			},
		}

		if cmd.config != "" {
			deploy.VmConfigSpec = &vcenter.VmConfigSpec{
				Provider: "XML",
				XML:      cmd.config,
			}
		}

		ref, err = m.DeployLibraryItem(ctx, item.ID, deploy)
		if err != nil {
			return err
		}
	case library.ItemTypeVMTX:
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

		deploy := vcenter.DeployTemplate{
			Name:          name,
			Description:   cmd.Options.Annotation,
			DiskStorage:   storage,
			VMHomeStorage: storage,
			Placement: &vcenter.Placement{
				ResourcePool: rp.Reference().Value,
				Host:         hostID,
				Folder:       folder.Reference().Value,
			},
		}
		ref, err = m.DeployTemplateLibraryItem(ctx, item.ID, deploy)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported library item type: %s", item.Type)
	}

	obj, err := finder.ObjectReference(ctx, *ref)
	if err != nil {
		return err
	}

	vm := obj.(*object.VirtualMachine)

	return cmd.Deploy(vm, cmd.FolderFlag.OutputFlag)
}
