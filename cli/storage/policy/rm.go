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

package policy

import (
	"context"
	"errors"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/pbm/types"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("storage.policy.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Usage() string {
	return "ID"
}

func (cmd *rm) Description() string {
	return `Remove Storage Policy ID.

Examples:
  govc storage.policy.rm "my policy name"
  govc storage.policy.rm af7935ab-466d-4b0e-af3c-4ec6bce2112f`
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	c, err := cmd.PbmClient()
	if err != nil {
		return err
	}

	arg := f.Arg(0)
	id := types.PbmProfileId{UniqueId: arg}

	profile, err := ListProfiles(ctx, c, arg)
	if err != nil {
		return err
	}

	if len(profile) == 1 {
		id = profile[0].GetPbmProfile().ProfileId
	}

	res, err := c.DeleteProfile(ctx, []types.PbmProfileId{id})
	if err != nil {
		return err
	}

	if len(res) != 0 {
		if f := res[0].Fault; f != nil {
			return errors.New(f.LocalizedMessage)
		}
	}

	return nil
}
