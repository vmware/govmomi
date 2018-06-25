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

package association

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25/types"
)

type attach struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("tags.attach", &attach{})
}

func (cmd *attach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *attach) Usage() string {
	return "ID PATH"
}

func (cmd *attach) Description() string {
	return `Attach tag to object.

Examples:
  govc tags.attach ID /dc1/host/cluster1/hostname
  govc tags.attach ID HostSystem:host`
}

func convertPath(ctx context.Context, cmd *flags.DatacenterFlag, managedObj string) (*types.ManagedObjectReference, error) {
	client, err := cmd.ClientFlag.Client()
	if err != nil {
		return nil, err
	}
	finder, err := cmd.Finder()
	if err != nil {
		return nil, err
	}

	ref := client.ServiceContent.RootFolder

	switch managedObj {
	case "", "-":
	default:
		if !ref.FromString(managedObj) {
			l, ferr := finder.ManagedObjectList(ctx, managedObj)
			if ferr != nil {
				return nil, ferr
			}

			switch len(l) {
			case 0:
				return nil, fmt.Errorf("%s not found", managedObj)
			case 1:
				ref = l[0].Object.Reference()
			default:
				return nil, flag.ErrHelp
			}
		}
	}
	return &ref, nil
}

func (cmd *attach) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	tagID := f.Arg(0)
	managedObj := f.Arg(1)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {
		ref, err := convertPath(ctx, cmd.DatacenterFlag, managedObj)
		if err != nil {
			return err
		}
		return c.AttachTagToObject(ctx, tagID, ref)
	})
}
