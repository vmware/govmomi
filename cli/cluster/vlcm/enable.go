/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package vlcm

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/vapi/cis/tasks"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type enable struct {
	*flags.ClientFlag

	clusterId string
	skipCheck bool
}

func init() {
	cli.Register("cluster.vlcm.enable", &enable{})
}

func (cmd *enable) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.BoolVar(&cmd.skipCheck, "skip-check", false, "Whether to skip the software check after enabling vLCM")
}

func (cmd *enable) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *enable) Usage() string {
	return "CLUSTER"
}

func (cmd *enable) Description() string {
	return `Enables vLCM on the provided cluster

This operation is irreversible

Examples:
  govc cluster.vlcm.enable -cluster-id=domain-c21`
}

func (cmd *enable) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	if taskId, err := dm.EnableSoftwareManagement(cmd.clusterId, cmd.skipCheck); err != nil {
		return err
	} else if _, err := tasks.NewManager(rc).WaitForCompletion(ctx, taskId); err != nil {
		return err
	} else {
		return nil
	}
}
