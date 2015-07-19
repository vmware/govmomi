/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package vm

import (
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"golang.org/x/net/context"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.SearchFlag

	WaitForIP   bool
	ExtraConfig bool
}

func init() {
	cli.Register("vm.info", &info{})
}

func (cmd *info) Register(f *flag.FlagSet) {
	cmd.SearchFlag = flags.NewSearchFlag(flags.SearchVirtualMachines)

	f.BoolVar(&cmd.WaitForIP, "waitip", false, "Wait for VM to acquire IP address")
	f.BoolVar(&cmd.ExtraConfig, "e", false, "Show ExtraConfig")
}

func (cmd *info) Process() error { return nil }

func (cmd *info) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vms, err := cmd.VirtualMachines(f.Args())
	if err != nil {
		if _, ok := err.(*find.NotFoundError); ok {
			// Continue with empty VM slice
		} else {
			return err
		}
	}

	var res infoResult
	var props []string

	if cmd.OutputFlag.JSON {
		props = nil // Load everything
	} else {
		props = []string{"summary", "guest.ipAddress"} // Load summary
		if cmd.ExtraConfig {
			props = append(props, "config.extraConfig")
		}
	}

	ctx := context.TODO()

	for _, vm := range vms {
		for {
			var mvm mo.VirtualMachine

			pc := property.DefaultCollector(c)
			err = pc.RetrieveOne(ctx, vm.Reference(), props, &mvm)
			if err != nil {
				return err
			}

			if cmd.WaitForIP && mvm.Guest.IpAddress == "" {
				_, err = vm.WaitForIP(ctx)
				if err != nil {
					return err
				}

				// Reload virtual machine object
				continue
			}

			var hostName string
			hostRef := mvm.Summary.Runtime.Host
			if hostRef == nil {
				hostName = "<unavailable>"
			} else {
				host := object.NewHostSystem(c, *hostRef)
				hostName, err = host.Name(ctx)
				if err != nil {
					return err
				}
			}

			res.VmInfos = append(res.VmInfos, vmInfo{mvm, hostName})
			break
		}
	}

	return cmd.WriteResult(&res)
}

type vmInfo struct {
	mo.VirtualMachine
	hostName string
}

type infoResult struct {
	VmInfos []vmInfo
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, vmInfo := range r.VmInfos {
		vm := vmInfo.VirtualMachine
		s := vm.Summary

		fmt.Fprintf(tw, "Name:\t%s\n", s.Config.Name)
		fmt.Fprintf(tw, "  UUID:\t%s\n", s.Config.Uuid)
		fmt.Fprintf(tw, "  Guest name:\t%s\n", s.Config.GuestFullName)
		fmt.Fprintf(tw, "  Memory:\t%dMB\n", s.Config.MemorySizeMB)
		fmt.Fprintf(tw, "  CPU:\t%d vCPU(s)\n", s.Config.NumCpu)
		fmt.Fprintf(tw, "  Power state:\t%s\n", s.Runtime.PowerState)
		fmt.Fprintf(tw, "  Boot time:\t%s\n", s.Runtime.BootTime)
		fmt.Fprintf(tw, "  IP address:\t%s\n", s.Guest.IpAddress)
		fmt.Fprintf(tw, "  Host:\t%s\n", vmInfo.hostName)
		if vm.Config != nil && vm.Config.ExtraConfig != nil {
			fmt.Fprintf(tw, "  ExtraConfig:\n")
			for _, v := range vm.Config.ExtraConfig {
				fmt.Fprintf(tw, "    %s:\t%s\n", v.GetOptionValue().Key, v.GetOptionValue().Value)
			}
		}
	}

	return tw.Flush()
}
