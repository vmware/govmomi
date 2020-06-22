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

package library

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type publish struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("library.publish", &publish{})
}

func (cmd *publish) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *publish) Usage() string {
	return "NAME|ITEM [SUBSCRIPTION-ID]..."
}

func (cmd *publish) Description() string {
	return `Publish library NAME or ITEM to subscribers.

If no subscriptions are specified, then publishes the library to all its subscribers.
See 'govc library.subscriber.ls' to get a list of subscription IDs.

Examples:
  govc library.publish /my-library
  govc library.publish /my-library subscription-id1 subscription-id2
  govc library.publish /my-library/my-item
  govc library.publish /my-library/my-item subscription-id1 subscription-id2`
}

func (cmd *publish) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	res, err := flags.ContentLibraryResult(ctx, c, "", f.Arg(0))
	if err != nil {
		return err
	}

	m := library.NewManager(c)

	ids := f.Args()[1:]

	switch t := res.GetResult().(type) {
	case library.Library:
		return m.PublishLibrary(ctx, &t, ids)
	case library.Item:
		return m.PublishLibraryItem(ctx, &t, false, ids)
	default:
		return fmt.Errorf("%q is a %T", res.GetPath(), t)
	}
}
