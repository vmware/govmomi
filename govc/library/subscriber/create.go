/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package subscriber

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type create struct {
	*flags.ClusterFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.NetworkFlag
	*flags.FolderFlag
}

func init() {
	cli.Register("library.subscriber.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.FolderFlag.Process(ctx)
}

func (cmd *create) Usage() string {
	return "PUBLISHED-LIBRARY SUBSCRIPTION-LIBRARY"
}

func (cmd *create) Description() string {
	return `Create library subscriber.

Examples:
  govc library.subscriber.create -cluster my-cluster published-library subscription-library`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	lib, err := flags.ContentLibrary(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}
	m := library.NewManager(c)

	sub, err := flags.ContentLibrary(ctx, c, f.Arg(1))
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
	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}
	folder, err := cmd.Folder()
	if err != nil {
		return err
	}
	network, err := cmd.Network()
	if err != nil {
		return err
	}

	spec := library.SubscriberLibrary{
		Target:    "USE_EXISTING",
		Location:  "LOCAL",
		LibraryID: sub.ID,
		Placement: &library.Placement{
			Folder:  folder.Reference().Value,
			Network: network.Reference().Value,
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

	id, err := m.CreateSubscriber(ctx, lib, spec)
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}
