// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package envbrowser

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type queryConfigTarget struct {
	command
}

func init() {
	cli.Register("envbrowser.query-config-target", &queryConfigTarget{})
}

func (cmd *queryConfigTarget) Description() string {
	return `Query the environment browser for a config target.

Examples:
  govc envbrowser.query-config-target -cluster my-cluster
  govc envbrowser.query-config-target -cluster my-cluster -host my-host`
}

func (cmd *queryConfigTarget) Run(ctx context.Context, f *flag.FlagSet) error {
	if err := cmd.command.Run(ctx, f); err != nil {
		return err
	}

	r, err := cmd.eb.QueryConfigTarget(ctx, cmd.host)
	if err != nil {
		return fmt.Errorf("failed to get config target: %w", err)
	}

	var ref types.ManagedObjectReference
	if cmd.host != nil {
		ref = cmd.host.Reference()
	} else {
		ref = cmd.cluster.Reference()
	}

	return cmd.WriteResult(queryConfigTargetResult{
		r:          r,
		ref:        ref,
		copyToFile: cmd.copyToFile,
	})
}

type queryConfigTargetResult struct {
	ref        types.ManagedObjectReference
	r          *types.ConfigTarget
	copyToFile bool
}

func (r queryConfigTargetResult) Write(w io.Writer) error {
	fmt.Println(r.ref)
	if r.copyToFile {
		fileName := fmt.Sprintf(
			"config-target-%s-%s",
			r.ref.Type,
			r.ref.Value)
		if s := r.ref.ServerGUID; s != "" {
			fileName += "-" + s
		}
		fileName += ".xml"
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		enc := xml.NewEncoder(f)
		enc.Indent("", "  ")
		return enc.Encode(r.r)
	}
	return nil
}
