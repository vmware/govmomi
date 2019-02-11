/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

package library

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

type create struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	library   library.Library
	datastore string
}

func init() {
	cli.Register("library.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.StringVar(&cmd.datastore, "D", "", "Datastore for library")
	f.StringVar(&cmd.library.Description, "d", "", "Description of library")

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Create library.

Examples:
  govc library.create library_name
  govc library.create -json | jq .
  govc library.create library_name -json | jq .`
}

type createResult []library.Library

func (r createResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i].Name)
	}
	return nil
}

func (cmd *create) lookupDatastore(ctx context.Context, c *vim25.Client, name string) (string, error) {
	finder, err := cmd.Finder()
	if err != nil {
		return name, err
	}
	objects, err := finder.DatastoreList(ctx, name)
	if err != nil {
		return name, err
	}

	return objects[0].Reference().Value, nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	cmd.library.Name = f.Arg(0)

	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		id, err := library.NewManager(c).CreateLibrary(ctx, cmd.library)

		if err != nil {
			return err
		}

		fmt.Println(id)
		return nil
	})
}
