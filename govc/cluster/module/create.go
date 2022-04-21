/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package module

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	vapicluster "github.com/vmware/govmomi/vapi/cluster"
)

type create struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("cluster.module.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	return cmd.ClusterFlag.Process(ctx)
}

func (cmd *create) Description() string {
	return `Create cluster module.

This command will output the ID of the new module.

Examples:
  govc cluster.module.create -cluster my_cluster`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 0 {
		return flag.ErrHelp
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	id, err := vapicluster.NewManager(c).CreateModule(ctx, cluster.Reference())
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}
