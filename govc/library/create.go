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
)

type create struct {
	*flags.DatastoreFlag
	library library.Library
}

func init() {
	cli.Register("library.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.StringVar(&cmd.library.Description, "d", "", "Description of library")
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

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	ds, err := cmd.Datastore()
	if err != nil {
		return err
	}

	cmd.library.Name = f.Arg(0)
	cmd.library.Type = "LOCAL"
	cmd.library.Storage = []library.StorageBackings{
		{
			DatastoreID: ds.Reference().Value,
			Type:        "DATASTORE",
		},
	}

	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		id, err := library.NewManager(c).CreateLibrary(ctx, cmd.library)
		if err != nil {
			return err
		}

		fmt.Println(id)
		return nil
	})
}
