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
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/pbm/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.ClientFlag

	spec types.PbmCapabilityProfileCreateSpec
	tag  string
	cat  string
}

func init() {
	cli.Register("storage.policy.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.spec.Description, "d", "", "Description")
	f.StringVar(&cmd.tag, "tag", "", "Tag")
	f.StringVar(&cmd.cat, "category", "", "Category")
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Create VM Storage Policy.

Examples:
  govc storage.policy.create -category my_cat -tag my_tag MyStoragePolicy # Tag based placement`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	cmd.spec.Name = f.Arg(0)
	cmd.spec.ResourceType.ResourceType = string(types.PbmProfileResourceTypeEnumSTORAGE)

	if cmd.tag == "" {
		return flag.ErrHelp
	}

	id := fmt.Sprintf("com.vmware.storage.tag.%s.property", cmd.cat)
	instance := types.PbmCapabilityInstance{
		Id: types.PbmCapabilityMetadataUniqueId{
			Namespace: "http://www.vmware.com/storage/tag",
			Id:        cmd.cat,
		},
		Constraint: []types.PbmCapabilityConstraintInstance{{
			PropertyInstance: []types.PbmCapabilityPropertyInstance{{
				Id: id,
				Value: types.PbmCapabilityDiscreteSet{
					Values: []vim.AnyType{cmd.tag},
				},
			}},
		}},
	}

	cmd.spec.Constraints = &types.PbmCapabilitySubProfileConstraints{
		SubProfiles: []types.PbmCapabilitySubProfile{{
			Name:       "Tag based placement",
			Capability: []types.PbmCapabilityInstance{instance},
		}},
	}

	c, err := cmd.PbmClient()
	if err != nil {
		return err
	}

	pid, err := c.CreateProfile(ctx, cmd.spec)
	if err != nil {
		return err
	}

	fmt.Println(pid.UniqueId)
	return nil
}
