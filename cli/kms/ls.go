/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package kms

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("kms.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Usage() string {
	return "NAME"
}

func (cmd *ls) Description() string {
	return `Display KMS info.

Examples:
  govc kms.ls
  govc kms.ls -json
  govc kms.ls - # default provider
  govc kms.ls ProviderName
  govc kms.ls -json ProviderName`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	info, err := m.ListKmipServers(ctx, nil)
	if err != nil {
		return err
	}

	id := f.Arg(0)

	if id == "" {
		status, err := m.GetStatus(ctx, info...)
		if err != nil {
			return err
		}
		return cmd.WriteResult(&clusterResult{Info: info, Status: status})
	}

	if id == "-" {
		for _, s := range info {
			if s.UseAsDefault {
				id = s.ClusterId.Id
				break
			}
		}
	}

	status, err := m.GetClusterStatus(ctx, id)
	if err != nil {
		return err
	}

	format := &serverResult{Status: status}

	for _, s := range info {
		if s.ClusterId.Id == id {
			format.Info = s
		}
	}

	return cmd.WriteResult(format)
}

type serverResult struct {
	Info   types.KmipClusterInfo                 `json:"info"`
	Status *types.CryptoManagerKmipClusterStatus `json:"status"`
}

func (r *serverResult) status(name string) types.ManagedEntityStatus {
	for _, server := range r.Status.Servers {
		if server.Name == name {
			return server.Status
		}
	}
	return types.ManagedEntityStatusGray
}

func (r *serverResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	if r.Info.ManagementType == string(types.KmipClusterInfoKmsManagementTypeNativeProvider) {
		boolVal := func(v *bool) bool {
			if v == nil {
				return false
			}
			return *v
		}

		fmt.Fprintf(tw, "Key ID: %s\tHas Backup: %t\tTPM Required: %t\n",
			r.Info.KeyId, boolVal(r.Info.HasBackup), boolVal(r.Info.TpmRequired))
	} else {
		for _, s := range r.Info.Servers {
			status := r.status(s.Name)
			fmt.Fprintf(tw, "%s\t%s:%d\t%s\n", s.Name, s.Address, s.Port, status)
		}
	}

	return tw.Flush()
}

type clusterResult struct {
	Info   []types.KmipClusterInfo                `json:"info"`
	Status []types.CryptoManagerKmipClusterStatus `json:"status"`
}

func (r *clusterResult) status(id types.KeyProviderId) types.ManagedEntityStatus {
	for _, status := range r.Status {
		if status.ClusterId == id {
			return status.OverallStatus
		}
	}
	return types.ManagedEntityStatusGray
}

func kmsType(kind string) string {
	switch types.KmipClusterInfoKmsManagementType(kind) {
	case types.KmipClusterInfoKmsManagementTypeVCenter:
		return "Standard"
	case types.KmipClusterInfoKmsManagementTypeNativeProvider:
		return "Native"
	default:
		return kind
	}
}

func (r *clusterResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, info := range r.Info {
		name := info.ClusterId.Id
		kind := kmsType(info.ManagementType)
		status := r.status(info.ClusterId)
		use := ""
		if info.UseAsDefault {
			use = "default"
		}
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", name, kind, status, use)
	}

	return tw.Flush()
}
