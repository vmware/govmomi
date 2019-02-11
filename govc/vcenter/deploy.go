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
	"github.com/vmware/govmomi/vim25"
)

type deploy struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	datastore string
}

func init() {
	cli.Register("vcenter.deploy", &deploy{})
}

func (cmd *deploy) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.StringVar(&cmd.datastore, "D", "", "Datastore for library")

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *deploy) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *deploy) Usage() string {
	return "NAME"
}

func (cmd *deploy) Description() string {
	return `Deploy library OVF template to vcenter.

Examples:
  govc vcenter.deploy library_name ovf_template vm_name`
}

func (cmd *deploy) lookupDatastore(ctx context.Context, c *vim25.Client, name string) (string, error) {
	finder, err := cmd.Finder()
	if err != nil {
		return name, err
	}
	objects, err := finder.DatastoreList(ctx, name)
	if err != nil {
		return name, err
	}

	return objects[0].Reference().Value, nil
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

func (cmd *deploy) getResourcePoolID(ctx context.Context, name string) (string, error) {
	finder, err := cmd.Finder()
	if err != nil {
		return "", err
	}
	o, err := finder.ClusterComputeResource(ctx, name)
	if err != nil {
		return "", err
	}

	p, err := o.ResourcePool(ctx)
	if err != nil {
		return "", err
	}

	return p.Reference().Value, nil
}

func (cmd *deploy) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
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

		/* 		poolid, err := cmd.getResourcePoolID(ctx, clusterName)
		   		if err != nil {
		   			return err
		   		} */
		finder, err := cmd.Finder()

		// Lookup default datastore
		ds, err := finder.DefaultDatastore(ctx)
		if err != nil {
			return err
		}
		datastoreID := ds.Reference().Value
		fmt.Printf("Using datastore ID %s\n", datastoreID)

		rp, err := finder.DefaultResourcePool(ctx)
		if err != nil {
			return err
		}
		poolID := rp.Reference().Value
		fmt.Printf("Using pool ID %s\n", poolID)

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
