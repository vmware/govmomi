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
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.DatacenterFlag
}

func init() {
	cli.Register("library.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Description() string {
	return `Display library information.

Examples:
  govc library.info
  govc library.info library_name
  govc library.info -json | jq .
  govc library.info library_name -json | jq .`
}

type infoResult []library.Library

func (t infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, item := range t {
		fmt.Fprintf(tw, "Name:\t%s\n", item.Name)
		fmt.Fprintf(tw, "  ID:\t%s\n", item.ID)
		fmt.Fprintf(tw, "  Description:\t%s\n", item.Description)
		fmt.Fprintf(tw, "  CategoryID:\t%s\n", item.Version)
		fmt.Fprintf(tw, "  StorageBackings:\n")
		fmt.Fprintf(tw, "    DatastoreID:\t%s\n", item.Storage[0].DatastoreID)
		fmt.Fprintf(tw, "    Type:\t%s\n", item.Storage[0].Type)
	}

	return tw.Flush()
}

// entityExists returns true if obj exists, false otherwise.  error is non-nil if an unexpected error occurs.
func convertDatastore(ctx context.Context, c *vim25.Client, managedObject string) string {
	obj := types.ManagedObjectReference{
		Type:  "Datastore",
		Value: managedObject,
	}
	pc := property.DefaultCollector(c)
	var me mo.ManagedEntity

	err := pc.RetrieveOne(ctx, obj, []string{"name"}, &me)
	//fmt.Printf("%+v\n", me)

	if err != nil {
		if soap.IsSoapFault(err) {
			_, notFound := soap.ToSoapFault(err).VimFault().(types.ManagedObjectNotFound)
			if notFound {
				return managedObject
			}
		}
		return managedObject
	}

	return me.Name
}

func (cmd *info) lookupDatastore(ctx context.Context, c *vim25.Client, name string) (string, error) {
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

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {

	client, err := cmd.Client()
	if err != nil {
		return err
	}

	cmd.lookupDatastore(ctx, client, "vsanDatastore")

	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var res infoResult
		var err error

		if f.NArg() == 1 {
			var result *library.Library
			arg := f.Arg(0)
			result, err = m.GetLibraryByName(ctx, arg)
			res = append(res, *result)
		} else {
			res, err = m.GetLibraries(ctx)
		}

		if err != nil {
			return err
		}

		for _, r := range res {
			r.Storage[0].DatastoreID = convertDatastore(ctx, client, r.Storage[0].DatastoreID)
		}
		return cmd.WriteResult(res)
	})
}
