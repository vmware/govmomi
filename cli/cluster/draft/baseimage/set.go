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

package draft

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/esx/settings/clusters"
)

type set struct {
	*flags.ClientFlag

	clusterId string
	draftId   string
	version   string
}

func init() {
	cli.Register("cluster.draft.baseimage.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.clusterId, "cluster-id", "", "The identifier of the cluster.")
	f.StringVar(&cmd.draftId, "draft-id", "", "The identifier of the software draft.")
	f.StringVar(&cmd.version, "version", "", "The identifier of the ESXi image version.")
}

func (cmd *set) Process(ctx context.Context) error {
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *set) Usage() string {
	return "CLUSTER"
}

func (cmd *set) Description() string {
	return `Sets the ESXi base image on the software draft.

Examples:
  govc cluster.draft.baseimage.set -cluster-id=domain-c21 -draft-id=31`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	rc, err := cmd.RestClient()

	if err != nil {
		return err
	}

	dm := clusters.NewManager(rc)

	return dm.SetSoftwareDraftBaseImage(cmd.clusterId, cmd.draftId, cmd.version)
}
