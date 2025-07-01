// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/pbm"
	"github.com/vmware/govmomi/pbm/types"
)

type create struct {
	*flags.ClientFlag

	spec pbm.CapabilityProfileCreateSpec
	tag  string
	cat  string
	zone bool
	enc  bool
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
	f.BoolVar(&cmd.enc, "e", false, "Enable encryption")
	f.BoolVar(&cmd.zone, "z", false, "Enable Zonal topology for multi-zone Supervisor")
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Create VM Storage Policy.

Examples:
  govc storage.policy.create -category my_cat -tag my_tag MyStoragePolicy # Tag based placement
  govc storage.policy.create -z MyZonalPolicy # Zonal topology`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	cmd.spec.Name = f.Arg(0)
	cmd.spec.Category = string(types.PbmProfileCategoryEnumREQUIREMENT)

	if cmd.tag == "" && !cmd.zone && !cmd.enc {
		return flag.ErrHelp
	}

	if cmd.tag != "" {
		cmd.spec.CapabilityList = append(cmd.spec.CapabilityList, pbm.Capability{
			ID:        cmd.cat,
			Namespace: "http://www.vmware.com/storage/tag",
			PropertyList: []pbm.Property{{
				ID:       fmt.Sprintf("com.vmware.storage.tag.%s.property", cmd.cat),
				Value:    cmd.tag,
				DataType: "set",
			}},
		})
	}

	if cmd.zone {
		cmd.spec.CapabilityList = append(cmd.spec.CapabilityList, pbm.Capability{
			ID:        "StorageTopology",
			Namespace: "com.vmware.storage.consumptiondomain",
			PropertyList: []pbm.Property{{
				ID:       "StorageTopologyType",
				Value:    "Zonal",
				DataType: "string",
			}},
		})
	}

	if cmd.enc {
		const encryptionCapabilityID = "ad5a249d-cbc2-43af-9366-694d7664fa52"

		cmd.spec.CapabilityList = append(cmd.spec.CapabilityList, pbm.Capability{
			ID:        encryptionCapabilityID,
			Namespace: "com.vmware.storageprofile.dataservice",
			PropertyList: []pbm.Property{{
				ID:       encryptionCapabilityID,
				Value:    encryptionCapabilityID,
				DataType: "string",
			}},
		})
	}

	c, err := cmd.PbmClient()
	if err != nil {
		return err
	}

	spec, err := pbm.CreateCapabilityProfileSpec(cmd.spec)
	if err != nil {
		return err
	}

	pid, err := c.CreateProfile(ctx, *spec)
	if err != nil {
		return err
	}

	fmt.Println(pid.UniqueId)
	return nil
}
