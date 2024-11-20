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

package cluster

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type disableCluster struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("namespace.cluster.disable", &disableCluster{})
}

func (cmd *disableCluster) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *disableCluster) Description() string {
	return `Disables vSphere Namespaces on the specified cluster.

Examples:
  govc namespace.cluster.disable -cluster "Workload-Cluster"`
}

func (cmd *disableCluster) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	// Cluster object reference lookup
	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	clusterId := cluster.Reference().Value

	err = m.DisableCluster(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("error disabling cluster: %s", err)
	}

	return nil
}
