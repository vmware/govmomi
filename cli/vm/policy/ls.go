// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/pbm/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("vm.policy.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *ls) Description() string {
	return `List VM storage policies.

Examples:
  govc vm.policy.ls -vm $name
  govc vm.policy.ls -vm $name -json`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	c, err := cmd.PbmClient()
	if err != nil {
		return err
	}

	devices, err := vm.Device(ctx)
	if err != nil {
		return err
	}

	m, err := c.ProfileMap(ctx)
	if err != nil {
		return err
	}

	ref := vm.Reference().Value

	res := []Entity{{
		Name: "VM home",
		ID: types.PbmServerObjectRef{
			ObjectType: string(types.PbmObjectTypeVirtualMachine),
			Key:        ref,
		},
	}}

	for _, disk := range devices.SelectByType((*vim.VirtualDisk)(nil)) {
		key := disk.GetVirtualDevice().Key

		entity := Entity{
			Name: devices.Name(disk),
			ID: types.PbmServerObjectRef{
				ObjectType: string(types.PbmObjectTypeVirtualDiskId),
				Key:        fmt.Sprintf("%s:%d", ref, key),
			},
		}

		res = append(res, entity)
	}

	for i := range res {
		entity := &res[i]
		entity.PolicyName = "None"

		ids, err := c.QueryAssociatedProfile(ctx, entity.ID)
		if err != nil {
			return err
		}

		if len(ids) != 0 {
			entity.PolicyID = ids[0].UniqueId
			entity.PolicyName = m.Name[entity.PolicyID].GetPbmProfile().Name

			res, err := c.FetchComplianceResult(ctx, []types.PbmServerObjectRef{entity.ID})
			if err != nil {
				return err
			}

			entity.Compliance = &res[0]
		}
	}

	return cmd.WriteResult(List(res))
}

type Entity struct {
	Name       string                     `json:"name"`
	ID         types.PbmServerObjectRef   `json:"id"`
	PolicyID   string                     `json:"policyID"`
	PolicyName string                     `json:"policyName"`
	Compliance *types.PbmComplianceResult `json:"compliance"`
}

type List []Entity

func (r List) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, disk := range r {
		status := ""
		checked := ""

		if c := disk.Compliance; c != nil {
			status = disk.Compliance.ComplianceStatus
			checked = c.CheckTime.Format(time.ANSIC)
		}

		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", disk.Name, disk.PolicyName, status, checked)
	}

	return tw.Flush()
}
