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
	"github.com/vmware/govmomi/vapi/library/finder"
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
  govc library.info /lib1
  govc library.info /lib1/item1
  govc library.info /lib1/item1/
  govc library.info */
  govc library.info -json | jq .
  govc library.info /lib1/item1 -json | jq .`
}

type infoResultsWriter []finder.FindResult

func (r infoResultsWriter) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	defer tw.Flush()
	for _, j := range r {
		switch t := j.GetResult().(type) {
		case library.Library:
			if err := r.writeLibrary(tw, t, j.GetPath()); err != nil {
				return err
			}
		case library.Item:
			if err := r.writeItem(tw, t, j.GetPath()); err != nil {
				return err
			}
		case library.File:
			if err := r.writeFile(tw, t, j.GetPath()); err != nil {
				return err
			}
		}
		tw.Flush()
	}
	return nil
}

func (r infoResultsWriter) writeLibrary(
	w io.Writer, v library.Library, ipath string) error {

	fmt.Fprintf(w, "Name:\t%s\n", v.Name)
	fmt.Fprintf(w, "  ID:\t%s\n", v.ID)
	fmt.Fprintf(w, "  Path:\t%s\n", ipath)
	fmt.Fprintf(w, "  Description:\t%s\n", v.Description)
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)
	fmt.Fprintf(w, "  StorageBackings:\n")
	for _, d := range v.Storage {
		fmt.Fprintf(w, "    DatastoreID:\t%s\n", d.DatastoreID)
		fmt.Fprintf(w, "    Type:\t%s\n", d.Type)
	}
	return nil
}
func (r infoResultsWriter) writeItem(
	w io.Writer, v library.Item, ipath string) error {

	fmt.Fprintf(w, "Name:\t%s\n", v.Name)
	fmt.Fprintf(w, "  ID:\t%s\n", v.ID)
	fmt.Fprintf(w, "  Path:\t%s\n", ipath)
	fmt.Fprintf(w, "  Description:\t%s\n", v.Description)
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)
	return nil
}
func (r infoResultsWriter) writeFile(
	w io.Writer, v library.File, ipath string) error {

	fmt.Fprintf(w, "Name:\t%s\n", v.Name)
	fmt.Fprintf(w, "  Path:\t%s\n", ipath)
	fmt.Fprintf(w, "  Size:\t%d\n", v.Size)
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)
	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {

	client, err := cmd.Client()
	if err != nil {
		return err
	}

	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		finder := finder.NewFinder(m)
		findResults, err := finder.Find(ctx, f.Args()...)
		if err != nil {
			return err
		}
		// Lookup the names(s) of the library's datastore(s).
		for i := range findResults {
			if t, ok := findResults[i].GetResult().(library.Library); ok {
				for j := range t.Storage {
					if t.Storage[j].Type == "DATASTORE" {
						t.Storage[j].DatastoreID = getDatastoreName(
							ctx, client, t.Storage[j].DatastoreID)
					}
				}
			}
		}
		return cmd.WriteResult(infoResultsWriter(findResults))
	})
}

func getDatastoreName(ctx context.Context, c *vim25.Client, managedObject string) string {
	obj := types.ManagedObjectReference{
		Type:  "Datastore",
		Value: managedObject,
	}
	pc := property.DefaultCollector(c)
	var me mo.ManagedEntity

	err := pc.RetrieveOne(ctx, obj, []string{"name"}, &me)
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
