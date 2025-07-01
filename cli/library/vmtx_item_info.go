// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/vcenter"
)

type vmtxItemInfo struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

type vmtxItemInfoResultsWriter struct {
	Result *vcenter.TemplateInfo `json:"result"`
	m      *vcenter.Manager
	cmd    *vmtxItemInfo
}

func (r vmtxItemInfoResultsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Result)
}

func (r vmtxItemInfoResultsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	defer tw.Flush()
	if err := r.writeLibraryTemplateDetails(tw, *r.Result); err != nil {
		return err
	}
	tw.Flush()
	return nil
}

func (r vmtxItemInfoResultsWriter) writeLibraryTemplateDetails(
	w io.Writer, v vcenter.TemplateInfo) error {

	fmt.Fprintf(w, "  VM Template:\t%s\n", v.VmTemplate)
	fmt.Fprintf(w, "  Guest OS:\t%s\n", v.GuestOS)
	fmt.Fprintf(w, "  CPU:\t\n")
	fmt.Fprintf(w, "    Count:\t%d\n", v.CPU.Count)
	fmt.Fprintf(w, "    Cores Per Socket:\t%d\n", v.CPU.CoresPerSocket)
	fmt.Fprintf(w, "  Memory:\t\n")
	fmt.Fprintf(w, "    Size in MB:\t%d\n", v.Memory.SizeMB)

	fmt.Fprintf(w, "  Disks:\t\n")
	for _, d := range v.Disks {
		fmt.Fprintf(w, "    Key:\t%s\n", d.Key)
		fmt.Fprintf(w, "    Capacity:\t%d\n", d.Value.Capacity)
		fmt.Fprintf(w, "    Datastore:\t%s\n\n", d.Value.DiskStorage.Datastore)
	}

	fmt.Fprintf(w, "  Nics:\t\n")
	for _, d := range v.Nics {
		fmt.Fprintf(w, "    Key:\t%s\n", d.Key)
		fmt.Fprintf(w, "    Backing Type:\t%s\n", d.Value.BackingType)
		fmt.Fprintf(w, "    Mac Type:\t%s\n", d.Value.MacType)
		fmt.Fprintf(w, "    Network:\t%s\n\n", d.Value.Network)
	}
	return nil
}

func init() {
	cli.Register("library.vmtx.info", &vmtxItemInfo{})
}

func (cmd *vmtxItemInfo) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *vmtxItemInfo) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *vmtxItemInfo) Description() string {
	return `Display VMTX template details

Examples:
  govc library.vmtx.info /library_name/vmtx_template_name`
}

func (cmd *vmtxItemInfo) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	path := f.Arg(0)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := vcenter.NewManager(c)

	// Fetch library item
	item, err := flags.ContentLibraryItem(ctx, c, path)
	if err != nil {
		return err
	}

	var res *vcenter.TemplateInfo
	if item.Type != library.ItemTypeVMTX {
		return fmt.Errorf("library item type should be 'vm-template' instead of '%s'", item.Type)
	}

	res, err = m.GetLibraryTemplateInfo(ctx, item.ID)
	if err != nil {
		return fmt.Errorf("error fetching library item details:  %s", err.Error())
	}

	return cmd.WriteResult(&vmtxItemInfoResultsWriter{res, m, cmd})
}
