/*
Copyright (c) 2016-2017 VMware, Inc. All Rights Reserved.

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

package pci

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type ls struct {
	*flags.ClientFlag
	*flags.HostSystemFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("host.hardware.pci.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Description() string {
	return `
Examples:
  govc host.hardware.pci.ls`
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	// We could do without the -host flag, leaving it for compat
	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	} else if host == nil {
		host, err = cmd.HostSystem()
		if err != nil {
			return err
		}
	}

	res := []mo.HostSystem{}
	refs := []types.ManagedObjectReference{host.Reference()}
	pc := property.DefaultCollector(c)
	props := []string{"summary", "hardware.pciDevice", "config.pciPassthruInfo"}
	err = pc.Retrieve(ctx, refs, props, &res)
	if err != nil {
		return err
	}

	printInfo := passthruInfo{}
	for _, obj := range res {
		infos := map[string]*types.HostPciPassthruInfo{}
		for _, o := range obj.Config.PciPassthruInfo {
			info := o.GetHostPciPassthruInfo()
			infos[info.Id] = info
		}
		for _, o := range obj.Hardware.PciDevice {
			status := "Not Capable"
			info, ok := infos[o.Id]
			if ok {
				if info.PassthruActive {
					status = "Active"
				} else if info.PassthruEnabled {
					status = "Enabled"
				} else if info.PassthruCapable {
					status = "Disabled"
				}
			}
			printInfo.Passthru = append(printInfo.Passthru, passthru{
				Id: o.Id, VendorName: o.VendorName, ParentBridge: o.ParentBridge, Status: status,
			})
		}
	}

	return cmd.WriteResult(&printInfo)
}

type passthru struct {
	Id           string
	VendorName   string
	ParentBridge string
	Status       string
}

type passthruInfo struct {
	Passthru []passthru
}

func (info *passthruInfo) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Address\tDescription\tParent\tPassthrough\n")
	for _, o := range info.Passthru {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", o.Id, o.VendorName, o.ParentBridge, o.Status)
	}
	tw.Flush()
	return nil
}
