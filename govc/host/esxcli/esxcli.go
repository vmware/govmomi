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

package esxcli

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type esxcli struct {
	*flags.HostSystemFlag
}

func init() {
	cli.Register("host.esxcli", &esxcli{})
}

func (cmd *esxcli) Register(f *flag.FlagSet) {}

func (cmd *esxcli) Process() error { return nil }

func (cmd *esxcli) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return nil
	}

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	e, err := NewExecutor(c, host)
	if err != nil {
		return err
	}

	res, err := e.Run(f.Args())
	if err != nil {
		return err
	}

	if len(res.Values) == 0 {
		return nil
	}

	var keys []string
	for key := range res.Values[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	// TODO: OutputFlag / format options
	for i, rv := range res.Values {
		if i > 0 {
			fmt.Fprintln(tw)
			_ = tw.Flush()
		}
		for _, key := range keys {
			fmt.Fprintf(tw, "%s:\t%s\n", key, strings.Join(rv[key], ", "))
		}
	}

	return tw.Flush()
}
