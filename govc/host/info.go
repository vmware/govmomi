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

package host

import (
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.SearchFlag
}

func init() {
	i := info{
		SearchFlag: flags.NewSearchFlag(flags.SearchHosts),
	}

	cli.Register(&i)
}

func (c *info) Register(f *flag.FlagSet) {}

func (c *info) Process() error { return nil }

func (c *info) Run(f *flag.FlagSet) error {
	client, err := c.Client()
	if err != nil {
		return err
	}

	host, err := c.HostSystem()
	if err != nil {
		return err
	}

	var res infoResult
	var props []string

	if c.OutputFlag.JSON {
		props = nil // Load everything
	} else {
		props = []string{"summary"} // Load summary
	}

	err = client.Properties(host.Reference(), props, &res.HostSystem)
	if err != nil {
		return err
	}

	return c.WriteResult(&res)
}

type infoResult struct {
	HostSystem mo.HostSystem
}

func (r *infoResult) WriteTo(w io.Writer) error {
	s := r.HostSystem.Summary
	h := s.Hardware

	tw := tabwriter.NewWriter(os.Stderr, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Name:\t%s\n", s.Config.Name)
	fmt.Fprintf(tw, "Manufacturer:\t%s\n", h.Vendor)
	fmt.Fprintf(tw, "Logical CPUs:\t%d CPUs @ %dMHz\n", h.NumCpuPkgs*h.NumCpuCores*h.NumCpuThreads, h.CpuMhz)
	fmt.Fprintf(tw, "Processor type:\t%s\n", h.CpuModel)
	fmt.Fprintf(tw, "Memory:\t%dMB\n", h.MemorySize/(1024*1024))
	fmt.Fprintf(tw, "Boot time:\t%s\n", s.Runtime.BootTime)
	return tw.Flush()
}
