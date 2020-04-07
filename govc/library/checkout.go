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

type checkout struct {
	*flags.ClusterFlag
	*flags.FolderFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
}

func init() {
	cli.Register("library.checkout", &checkout{})
}

func (cmd *checkout) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *checkout) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.FolderFlag.Process(ctx)
}

func (cmd *checkout) Usage() string {
	return "PATH NAME"
}

func (cmd *checkout) Description() string {
	return `Check out Content Library item PATH to vm NAME.

Examples:
  govc library.checkout -cluster my-cluster my-content/template-vm-item my-vm`
}

func (cmd *checkout) Run(ctx context.Context, f *flag.FlagSet) error {
	path := f.Arg(0)
	name := f.Arg(1)

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

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	l, err := flags.ContentLibraryItem(ctx, c, path)
	if err != nil {
		return err
	}

	spec := vcenter.CheckOut{
		Name: name,
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

	id, err := vcenter.NewManager(c).CheckOut(ctx, l.ID, &spec)
	if err != nil {
		return err
	}

	fmt.Println(id)

	return nil
}
