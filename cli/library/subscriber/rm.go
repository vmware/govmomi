// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package subscriber

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("library.subscriber.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "SUBSCRIPTION-ID"
}

func (cmd *rm) Description() string {
	return `Delete subscription of the published library.

The subscribed library associated with the subscription will not be deleted.

Examples:
  id=$(govc library-subscriber.ls | grep my-library-name | awk '{print $1}')
  govc library.subscriber.rm $id`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	lib, err := flags.ContentLibrary(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}
	m := library.NewManager(c)

	return m.DeleteSubscriber(ctx, lib, f.Arg(0))
}
