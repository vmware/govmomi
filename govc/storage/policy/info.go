/*
Copyright (c) 2020-2023 VMware, Inc. All Rights Reserved.

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
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	vim "github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag

	compliance bool
	storage    bool
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
	Profile              types.BasePbmProfile `json:"profile"`
	CompliantVM          []string             `json:"compliantVM"`
	CompatibleDatastores []string             `json:"compatibleDatastores"`
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

	profiles, err := ListProfiles(ctx, c, f.Arg(0))
	if err != nil {
		return err
	}

	pc := property.DefaultCollector(vc)
	kind := []string{"Datastore"}
	m := view.NewManager(vc)

	v, err := m.CreateContainerView(ctx, vc.ServiceContent.RootFolder, kind, true)
	if err != nil {
		return err
	}

	var content []vim.ObjectContent
	err = v.Retrieve(ctx, kind, []string{"name"}, &content)
	_ = v.Destroy(ctx)
	if err != nil {
		return err
	}

	dsmap := make(map[string]string)
	var hubs []types.PbmPlacementHub

	for _, ds := range content {
		hubs = append(hubs, types.PbmPlacementHub{
			HubType: ds.Obj.Type,
			HubId:   ds.Obj.Value,
		})
		dsmap[ds.Obj.Value] = ds.PropSet[0].Val.(string)
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

			res, err := c.CheckRequirements(ctx, hubs, nil, req)
			if err != nil {
				return err
			}

			for _, hub := range res.CompatibleDatastores() {
				policy.CompatibleDatastores = append(policy.CompatibleDatastores, dsmap[hub.HubId])
			}
		}

		policies = append(policies, policy)
	}

	return cmd.WriteResult(&infoResult{policies, cmd})
}
