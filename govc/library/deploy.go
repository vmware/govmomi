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
	"github.com/vmware/govmomi/govc/importx"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/vcenter"
	"github.com/vmware/govmomi/vim25/types"
)

type deploy struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.FolderFlag
	*importx.OptionsFlag
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

	cmd.OptionsFlag = new(importx.OptionsFlag)
	cmd.OptionsFlag.Register(ctx, f)
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
	return cmd.OptionsFlag.Process(ctx)
}

func (cmd *deploy) Usage() string {
	return "TEMPLATE [NAME]"
}

func (cmd *deploy) Description() string {
	return `Deploy library OVF template.

Examples:
  govc library.deploy /library_name/ovf_template vm_name
  govc library.deploy /library_name/ovf_template -options deploy.json`
}

func (cmd *deploy) Run(ctx context.Context, f *flag.FlagSet) error {
	path := f.Arg(0)
	name := f.Arg(1)

	if name == "" && cmd.Options.Name != nil {
		name = *cmd.Options.Name
	}

	return cmd.DatastoreFlag.WithRestClient(ctx, func(c *rest.Client) error {
		m := vcenter.NewManager(c)

		res, err := finder.NewFinder(library.NewManager(c)).Find(ctx, path)
		if err != nil {
			return err
		}
		if len(res) != 1 {
			return fmt.Errorf("%q matches %d items", path, len(res))
		}
		item, ok := res[0].GetResult().(library.Item)
		if !ok {
			return fmt.Errorf("%q is a %T", path, item)
		}

		ds, err := cmd.Datastore()
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

		deploy := vcenter.Deploy{
			DeploymentSpec: vcenter.DeploymentSpec{
				Name:               name,
				DefaultDatastoreID: ds.Reference().Value,
				AcceptAllEULA:      true,
				Annotation:         cmd.Options.Annotation,
				AdditionalParams: []vcenter.AdditionalParams{
					{
						Class:       vcenter.ClassOvfParams,
						Type:        vcenter.TypeDeploymentOptionParams,
						SelectedKey: cmd.Options.Deployment,
					},
				},
				NetworkMappings:     networks,
				StorageProvisioning: cmd.Options.DiskProvisioning,
			},
			Target: vcenter.Target{
				ResourcePoolID: rp.Reference().Value,
				HostID:         hostID,
				FolderID:       folder.Reference().Value,
			},
		}

		cmd.FolderFlag.Log("Deploying library item...\n")

		d, err := m.DeployLibraryItem(ctx, item.ID, deploy)
		if err != nil {
			return err
		}

		if !d.Succeeded {
			return d.Error
		}

		ref, err := finder.ObjectReference(ctx, types.ManagedObjectReference(*d.ResourceID))
		if err != nil {
			return err
		}

		vm := ref.(*object.VirtualMachine)

		return cmd.Deploy(vm, cmd.FolderFlag.OutputFlag)
	})
}
