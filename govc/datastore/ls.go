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

package datastore

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.DatastoreFlag

	force bool
}

func init() {
	cli.Register("datastore.ls", &ls{})
}

func (cmd *ls) Register(f *flag.FlagSet) {}

func (cmd *ls) Process() error { return nil }

func (cmd *ls) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	b, err := ds.Browser()
	if err != nil {
		return err
	}

	path, err := cmd.DatastorePath(f.Arg(0))
	if err != nil {
		return err
	}

	spec := types.HostDatastoreBrowserSearchSpec{
		Details: &types.FileQueryFlags{
			FileType:     true,
			FileSize:     true,
			FileOwner:    true, // TODO: omitempty is generated, but seems to be required
			Modification: true,
		},
	}

	task, err := b.SearchDatastore(c, path, &spec)
	if err != nil {
		return err
	}

	info, err := task.WaitForResult(nil)
	if err != nil {
		return err
	}

	res := info.Result.(types.HostDatastoreBrowserSearchResults)

	tw := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)

	for _, file := range res.File {
		info := file.GetFileInfo()
		fmt.Fprintf(tw, "%d\t%s\t%s\n", info.FileSize, info.Modification.Format("Mon Jan 2 15:04:05 2006"), info.Path)
	}

	return tw.Flush()
}
