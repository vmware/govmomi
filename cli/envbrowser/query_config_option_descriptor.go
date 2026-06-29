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

type queryConfigOptionDescriptor struct {
	command
}

func init() {
	cli.Register("envbrowser.query-config-option-descriptor", &queryConfigOptionDescriptor{})
}

func (cmd *queryConfigOptionDescriptor) Description() string {
	return `Query the environment browser for the config descriptors.

Examples:
  govc envbrowser.query-config-option-descriptor -cluster my-cluster`
}

func (cmd *queryConfigOptionDescriptor) Run(
	ctx context.Context,
	f *flag.FlagSet) error {

	if err := cmd.command.Run(ctx, f); err != nil {
		return err
	}

	r, err := cmd.eb.QueryConfigOptionDescriptor(ctx)
	if err != nil {
		return fmt.Errorf("failed to get config option descriptor: %w", err)
	}

	return cmd.WriteResult(queryConfigOptionDescriptorResult{
		r:          r,
		copyToFile: cmd.copyToFile,
	})
}

type queryConfigOptionDescriptorResult struct {
	r          []types.VirtualMachineConfigOptionDescriptor
	copyToFile bool
}

func (r queryConfigOptionDescriptorResult) Write(w io.Writer) error {
	for _, r := range r.r {
		fmt.Println(r.Key)
	}
	if r.copyToFile {
		f, err := os.Create("config-option-descriptor.xml")
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
