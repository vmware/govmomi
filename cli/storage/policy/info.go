// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/property"
	vim "github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag

	compliance bool
	storage    bool
	iofilters  bool
}

func init() {
	cli.Register("storage.policy.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.storage, "s", false, "Check Storage Compatibility")
	f.BoolVar(&cmd.compliance, "c", false, "Check VM Compliance")

	if cli.ShowUnreleased() {
		f.BoolVar(&cmd.iofilters, "i", false, "Query IO Filters")
	}
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Usage() string {
	return "[NAME]"
}

func (cmd *info) Description() string {
	return `VM Storage Policy info.

Examples:
  govc storage.policy.info
  govc storage.policy.info "vSAN Default Storage Policy"
  govc storage.policy.info -c -s`
}

type Policy struct {
	Profile              types.BasePbmProfile            `json:"profile"`
	CompliantVM          []string                        `json:"compliantVM"`
	CompatibleDatastores []string                        `json:"compatibleDatastores"`
	FilterMap            []types.PbmProfileToIofilterMap `json:"filterMap,omitempty"`
}

type infoResult struct {
	Policies []Policy `json:"policies"`
	cmd      *info
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, policy := range r.Policies {
		p := policy.Profile.GetPbmProfile()
		_, _ = fmt.Fprintf(tw, "Name:\t%s\n", p.Name)
		_, _ = fmt.Fprintf(tw, "  ID:\t%s\n", p.ProfileId.UniqueId)
		_, _ = fmt.Fprintf(tw, "  Description:\t%s\n", p.Description)
		if r.cmd.compliance {
			_, _ = fmt.Fprintf(tw, "  Compliant VMs:\t%s\n", strings.Join(policy.CompliantVM, ","))
		}
		if r.cmd.storage {
			_, _ = fmt.Fprintf(tw, "  Compatible Datastores:\t%s\n", strings.Join(policy.CompatibleDatastores, ","))
		}
	}

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	c, err := cmd.PbmClient()
	if err != nil {
		return err
	}

	pc := property.DefaultCollector(vc)

	profiles, err := ListProfiles(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}

	ds, err := c.DatastoreMap(ctx, vc, vc.ServiceContent.RootFolder)
	if err != nil {
		return err
	}

	var policies []Policy

	for _, profile := range profiles {
		p := profile.GetPbmProfile()

		policy := Policy{
			Profile: profile,
		}

		if cmd.compliance {
			entities, err := c.QueryAssociatedEntity(ctx, p.ProfileId, string(types.PbmObjectTypeVirtualMachine))
			if err != nil {
				return err
			}

			if len(entities) == 0 {
				continue
			}

			res, err := c.FetchComplianceResult(ctx, entities)
			if err != nil {
				return err
			}

			var refs []vim.ManagedObjectReference
			for _, r := range res {
				if r.ComplianceStatus == string(types.PbmComplianceStatusCompliant) {
					refs = append(refs, vim.ManagedObjectReference{Type: "VirtualMachine", Value: r.Entity.Key})
				}
			}

			var content []vim.ObjectContent
			err = pc.Retrieve(ctx, refs, []string{"name"}, &content)
			if err != nil {
				return err
			}

			for _, c := range content {
				policy.CompliantVM = append(policy.CompliantVM, c.PropSet[0].Val.(string))
			}
		}

		if cmd.storage {
			req := []types.BasePbmPlacementRequirement{
				&types.PbmPlacementCapabilityProfileRequirement{
					ProfileId: p.ProfileId,
				},
			}

			res, err := c.CheckRequirements(ctx, ds.PlacementHub, nil, req)
			if err != nil {
				return err
			}

			for _, hub := range res.CompatibleDatastores() {
				policy.CompatibleDatastores = append(policy.CompatibleDatastores, ds.Name[hub.HubId])
			}
		}

		if cmd.iofilters {
			policy.FilterMap, err = c.QueryIOFiltersFromProfileId(ctx, p.ProfileId.UniqueId)
			if err != nil {
				return err
			}
		}

		policies = append(policies, policy)
	}

	return cmd.WriteResult(&infoResult{policies, cmd})
}
