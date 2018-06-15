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

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type detach struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
}

func init() {
	cli.Register("tags.detach", &detach{})
}

func (cmd *detach) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *detach) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *detach) Usage() string {
	return "ID MANAGEDOBJECTREFERENCE"
}

func (cmd *detach) Description() string {
	return `Detach tag from object.

Examples:
  govc tags.detach ID MANAGEDOBJECTREFERENCE`
}

func (cmd *detach) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	tagID := f.Arg(0)
	managedObj := f.Arg(1)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		objType, objID, err := convertPath(ctx, cmd.DatacenterFlag, managedObj)
		if err != nil {
			return err
		}
		return c.DetachTagFromObject(ctx, tagID, objID, objType)

	})

}
