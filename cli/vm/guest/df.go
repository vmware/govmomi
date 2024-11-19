/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package guest

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type df struct {
	*flags.VirtualMachineFlag
}

func init() {
	cli.Register("guest.df", &df{})
}

func (cmd *df) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
}

func (cmd *df) Description() string {
	return `Report file system disk space usage.

Examples:
  govc guest.df -vm $name`
}

type dfResult []types.GuestDiskInfo

func (r dfResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", "Filesystem", "Size", "Used", "Avail", "Use%")
	for _, disk := range r {
		used := disk.Capacity - disk.FreeSpace
		use := 100.0 * float32(used) / float32(disk.Capacity)
		_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%.0f%%\n", disk.DiskPath,
			units.ByteSize(disk.Capacity), units.ByteSize(used), units.ByteSize(disk.FreeSpace), use)
	}
	return tw.Flush()
}

func (cmd *df) Run(ctx context.Context, f *flag.FlagSet) error {
	obj, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	var vm mo.VirtualMachine
	err = obj.Properties(ctx, obj.Reference(), []string{"guest.disk"}, &vm)
	if err != nil {
		return err
	}

	return cmd.WriteResult(dfResult(vm.Guest.Disk))
}
