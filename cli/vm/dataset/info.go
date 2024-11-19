/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package dataset

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type info struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("vm.dataset.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *info) Usage() string {
	return "NAME"
}

func (cmd *info) Description() string {
	return `Display data set information.

Examples:
  govc vm.dataset.info -vm $vm
  govc vm.dataset.info -vm $vm com.example.project2`
}

type infoResult []*dataset.Info

func (r infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, item := range r {
		fmt.Fprintf(tw, "Name:\t%s\n", item.Name)
		fmt.Fprintf(tw, "  Description:\t%s\n", item.Description)
		fmt.Fprintf(tw, "  Host:\t%s\n", item.Host)
		fmt.Fprintf(tw, "  Guest:\t%s\n", item.Guest)
		fmt.Fprintf(tw, "  Used: \t%d\n", item.Used)
		fmt.Fprintf(tw, "  OmitFromSnapshotAndClone: \t%t\n", item.OmitFromSnapshotAndClone)
	}

	return tw.Flush()
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() > 1 {
		return errors.New("please specify at most one data set")
	}

	vm, err := cmd.VirtualMachineFlag.VirtualMachine()
	if err != nil {
		return err
	}
	if vm == nil {
		return flag.ErrHelp
	}
	vmId := vm.Reference().Value

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	mgr := dataset.NewManager(c)

	// Discover the relevant data set id(s)
	var ids []string
	if f.NArg() == 1 {
		// single data set
		id, err := FindDataSetId(ctx, mgr, vmId, f.Arg(0))
		if err != nil {
			return err
		}
		ids = []string{id}
	} else {
		// all data sets
		l, err := mgr.ListDataSets(ctx, vmId)
		if err != nil {
			return err
		}
		for _, summary := range l {
			ids = append(ids, summary.DataSet)
		}
	}

	// Fetch the information for each data set id
	var res []*dataset.Info
	for _, id := range ids {
		inf, err := mgr.GetDataSet(ctx, vmId, id)
		if err != nil {
			return err
		}
		res = append(res, inf)
	}

	return cmd.WriteResult(infoResult(res))
}
