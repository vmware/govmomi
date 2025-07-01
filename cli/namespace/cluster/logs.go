// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type logs struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("namespace.logs.download", &logs{})
}

func (cmd *logs) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *logs) Usage() string {
	return "[NAME]"
}

func (cmd *logs) Description() string {
	return `Download namespace cluster support bundle.

If NAME name is "-", bundle is written to stdout.

See also: govc logs.download

Examples:
  govc namespace.logs.download -cluster k8s
  govc namespace.logs.download -cluster k8s - | tar -xvf -
  govc namespace.logs.download -cluster k8s logs.tar`
}

func (cmd *logs) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	id := cluster.Reference().Value

	name := f.Arg(0)

	m := namespace.NewManager(c)

	bundle, err := m.CreateSupportBundle(ctx, id)
	if err != nil {
		return err
	}

	req, err := m.SupportBundleRequest(ctx, bundle)
	if err != nil {
		return err
	}

	if id := c.SessionID(); id != "" {
		req.Header.Set("vmware-api-session-id", id)
	}

	return c.DownloadAttachment(ctx, req, name)
}
