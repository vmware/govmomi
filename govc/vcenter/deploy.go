/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

package vcenter

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vapi/vcenter"
)

type deploy struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.FolderFlag
}

func init() {
	cli.Register("vcenter.deploy", &deploy{})
}

func (cmd *deploy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *deploy) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.FolderFlag.Process(ctx)
}

func (cmd *deploy) Usage() string {
	return "LIBRARY TEMPLATE VM_NAME"
}

func (cmd *deploy) Description() string {
	return `Deploy library OVF template to vcenter.

Examples:
  govc vcenter.deploy library_name ovf_template vm_name`
}

func getOVFItemID(ctx context.Context, c *rest.Client, libname string, ovfname string) (string, error) {
	m := library.NewManager(c)

	var library *library.Library
	library, err := m.GetLibraryByName(ctx, libname)
	if err != nil {
		return "", err
	}
	res, err := m.GetLibraryItems(ctx, library.ID)

	for _, r := range res {
		if r.Name == ovfname || r.ID == ovfname {
			return r.ID, nil
		}
	}

	return "", fmt.Errorf("Could not find %s in library %s", ovfname, libname)
}

func (cmd *deploy) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.DatastoreFlag.WithRestClient(ctx, func(c *rest.Client) error {
		m := vcenter.NewManager(c)

		if f.NArg() != 3 {
			return flag.ErrHelp
		}

		libname := f.Arg(0)
		templateName := f.Arg(1)
		vmName := f.Arg(2)

		ovfItemID, err := getOVFItemID(ctx, c, libname, templateName)
		if err != nil {
			return err
		}

		ds, err := cmd.Datastore()
		if err != nil {
			return err
		}
		datastoreID := ds.Reference().Value
		fmt.Printf("Using datastore ID %s\n", datastoreID)

		rp, err := cmd.ResourcePool()
		if err != nil {
			return err
		}
		poolID := rp.Reference().Value
		fmt.Printf("Using pool ID %s\n", poolID)

		folder, err := cmd.Folder()
		if err != nil {
			return err
		}
		folderID := folder.Reference().Value
		fmt.Printf("Using folder ID %s\n", folderID)

		filter := vcenter.FilterRequest{
			Target: vcenter.Target{
				ResourcePoolID: poolID,
			},
		}

		filterResponse, err := m.FilterLibraryItem(ctx, ovfItemID, filter)
		fmt.Printf("Found OVA for deployment: %s\n", filterResponse.Name)

		deploy := vcenter.Deploy{
			DeploymentSpec: vcenter.DeploymentSpec{
				Name:               vmName,
				DefaultDatastoreID: datastoreID,
				AcceptAllEULA:      true,
			},
			Target: vcenter.Target{
				ResourcePoolID: poolID,
				FolderID:       folderID,
			},
		}

		resp, err := m.DeployLibraryItem(ctx, ovfItemID, deploy)
		if err != nil {
			return err
		}

		if resp.Succeeded {
			fmt.Printf("Deploy succeeded: %s\n", resp.ResourceID.ID)
		} else {
			fmt.Printf("%+v", resp)
		}

		return nil
	})
}
