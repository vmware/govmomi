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
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"path"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.DatacenterFlag

	long bool
	link bool
	url  bool

	names map[string]string
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

	f.BoolVar(&cmd.long, "l", false, "Long listing format")
	f.BoolVar(&cmd.link, "L", false, "List Datastore path only")
	f.BoolVar(&cmd.url, "U", false, "List pub/sub URL(s) only")

	cmd.names = make(map[string]string)
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
  govc library.info -l /lib1 | grep Size:
  govc library.info /lib1/item1
  govc library.info /lib1/item1/
  govc library.info */
  govc device.cdrom.insert -vm $vm -device cdrom-3000 $(govc library.info -L /lib1/item1/file1)
  govc library.info -json | jq .
  govc library.info /lib1/item1 -json | jq .`
}

type infoResultsWriter struct {
	Result []finder.FindResult
	m      *library.Manager
	cmd    *info
}

func (r infoResultsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Result)
}

func (r infoResultsWriter) Dump() interface{} {
	res := make([]interface{}, len(r.Result))
	for i := range r.Result {
		res[i] = r.Result[0].GetResult()
	}
	return res
}

func (r infoResultsWriter) Write(w io.Writer) error {
	if r.cmd.link {
		for _, j := range r.Result {
			p, err := r.cmd.getDatastoreFilePath(j)
			if err != nil {
				return err
			}
			if !r.cmd.long {
				var path object.DatastorePath
				path.FromString(p)
				p = path.Path
			}
			fmt.Fprintln(w, p)
		}
		return nil
	}

	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	defer tw.Flush()
	for _, j := range r.Result {
		switch t := j.GetResult().(type) {
		case library.Library:
			if err := r.writeLibrary(tw, t, j); err != nil {
				return err
			}
		case library.Item:
			if err := r.writeItem(tw, t, j); err != nil {
				return err
			}
		case library.File:
			if err := r.writeFile(tw, t, j); err != nil {
				return err
			}
		}
		tw.Flush()
	}
	return nil
}

func (r infoResultsWriter) writeLibrary(
	w io.Writer, v library.Library, res finder.FindResult) error {

	published := v.Publication != nil && *v.Publication.Published

	if r.cmd.url {
		switch {
		case v.Subscription != nil:
			_, _ = fmt.Fprintf(w, "%s\n", v.Subscription.SubscriptionURL)
		case published:
			_, _ = fmt.Fprintf(w, "%s\n", v.Publication.PublishURL)
		}

		return nil
	}

	fmt.Fprintf(w, "Name:\t%s\n", v.Name)
	fmt.Fprintf(w, "  ID:\t%s\n", v.ID)
	fmt.Fprintf(w, "  Path:\t%s\n", res.GetPath())
	fmt.Fprintf(w, "  Description:\t%s\n", v.Description)
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)
	fmt.Fprintf(w, "  Created:\t%s\n", v.CreationTime.Format(time.ANSIC))
	fmt.Fprintf(w, "  StorageBackings:\t\n")
	for _, d := range v.Storage {
		fmt.Fprintf(w, "    DatastoreID:\t%s\n", d.DatastoreID)
		fmt.Fprintf(w, "    Type:\t%s\n", d.Type)
	}
	if r.cmd.long {
		fmt.Fprintf(w, "  Datastore Path:\t%s\n", r.cmd.getDatastorePath(res))
		items, err := r.m.GetLibraryItems(context.Background(), v.ID)
		if err != nil {
			return err
		}
		var size int64
		for i := range items {
			size += items[i].Size
		}
		fmt.Fprintf(w, "  Size:\t%s\n", units.ByteSize(size))
		fmt.Fprintf(w, "  Items:\t%d\n", len(items))
	}
	if v.Subscription != nil {
		dl := "All"
		if v.Subscription.OnDemand != nil && *v.Subscription.OnDemand {
			dl = "On Demand"
		}

		fmt.Fprintf(w, "  Subscription:\t\n")
		fmt.Fprintf(w, "    AutoSync:\t%t\n", *v.Subscription.AutomaticSyncEnabled)
		fmt.Fprintf(w, "    URL:\t%s\n", v.Subscription.SubscriptionURL)
		fmt.Fprintf(w, "    Auth:\t%s\n", v.Subscription.AuthenticationMethod)
		fmt.Fprintf(w, "    Download:\t%s\n", dl)
	}
	if published {
		fmt.Fprintf(w, "  Publication:\t\n")
		fmt.Fprintf(w, "    URL:\t%s\n", v.Publication.PublishURL)
		fmt.Fprintf(w, "    Auth:\t%s\n", v.Publication.AuthenticationMethod)
	}
	return nil
}

func (r infoResultsWriter) writeItem(
	w io.Writer, v library.Item, res finder.FindResult) error {

	fmt.Fprintf(w, "Name:\t%s\n", v.Name)
	fmt.Fprintf(w, "  ID:\t%s\n", v.ID)
	fmt.Fprintf(w, "  Path:\t%s\n", res.GetPath())
	fmt.Fprintf(w, "  Description:\t%s\n", v.Description)
	fmt.Fprintf(w, "  Type:\t%s\n", v.Type)
	fmt.Fprintf(w, "  Size:\t%s\n", units.ByteSize(v.Size))
	fmt.Fprintf(w, "  Created:\t%s\n", v.CreationTime.Format(time.ANSIC))
	fmt.Fprintf(w, "  Modified:\t%s\n", v.LastModifiedTime.Format(time.ANSIC))
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)

	if r.cmd.long {
		fmt.Fprintf(w, "  Datastore Path:\t%s\n", r.cmd.getDatastorePath(res))
	}

	return nil
}

func (r infoResultsWriter) writeFile(
	w io.Writer, v library.File, res finder.FindResult) error {

	size := "-"
	if v.Size != nil {
		size = units.ByteSize(*v.Size).String()
	}
	fmt.Fprintf(w, "Name:\t%s\n", v.Name)
	fmt.Fprintf(w, "  Path:\t%s\n", res.GetPath())
	fmt.Fprintf(w, "  Size:\t%s\n", size)
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)

	if r.cmd.long {
		fmt.Fprintf(w, "  Datastore Path:\t%s\n", r.cmd.getDatastorePath(res))
	}

	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

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
					t.Storage[j].DatastoreID = cmd.getDatastoreName(t.Storage[j].DatastoreID)
				}
			}
		}
	}
	return cmd.WriteResult(&infoResultsWriter{findResults, m, cmd})
}

func (cmd *info) getDatastoreName(id string) string {
	if name, ok := cmd.names[id]; ok {
		return name
	}

	c, err := cmd.Client()
	if err != nil {
		return id
	}

	obj := types.ManagedObjectReference{
		Type:  "Datastore",
		Value: id,
	}
	pc := property.DefaultCollector(c)
	var me mo.ManagedEntity

	err = pc.RetrieveOne(context.Background(), obj, []string{"name"}, &me)
	if err != nil {
		return id
	}

	cmd.names[id] = me.Name
	return me.Name
}

func (cmd *info) getDatastorePath(r finder.FindResult) string {
	p, _ := cmd.getDatastoreFilePath(r)
	return p
}

func (cmd *info) getDatastoreFilePath(r finder.FindResult) (string, error) {
	switch x := r.GetResult().(type) {
	case library.Library:
		id := ""
		if len(x.Storage) != 0 {
			id = cmd.getDatastoreName(x.Storage[0].DatastoreID)
		}
		return fmt.Sprintf("[%s] contentlib-%s", id, x.ID), nil
	case library.Item:
		return fmt.Sprintf("%s/%s", cmd.getDatastorePath(r.GetParent()), x.ID), nil
	case library.File:
		return cmd.getDatastoreFileItemPath(r)
	default:
		return "", fmt.Errorf("unsupported type=%T", x)
	}
}

// getDatastoreFileItemPath returns the absolute datastore path for a library.File
func (cmd *info) getDatastoreFileItemPath(r finder.FindResult) (string, error) {
	name := r.GetName()
	dir := cmd.getDatastorePath(r.GetParent())
	p := path.Join(dir, name)

	lib := r.GetParent().GetParent().GetResult().(library.Library)
	if len(lib.Storage) == 0 {
		return p, nil
	}

	ctx := context.Background()
	c, err := cmd.Client()
	if err != nil {
		return p, err
	}

	ref := types.ManagedObjectReference{Type: "Datastore", Value: lib.Storage[0].DatastoreID}
	ds := object.NewDatastore(c, ref)

	b, err := ds.Browser(ctx)
	if err != nil {
		return p, err
	}

	// The file ID isn't available via the API, so we use DatastoreBrowser to search
	ext := path.Ext(name)
	pat := strings.Replace(name, ext, "*"+ext, 1)
	spec := types.HostDatastoreBrowserSearchSpec{
		MatchPattern: []string{pat},
	}

	task, err := b.SearchDatastore(ctx, dir, &spec)
	if err != nil {
		return p, err
	}

	info, err := task.WaitForResult(ctx, nil)
	if err != nil {
		return p, err
	}

	res, ok := info.Result.(types.HostDatastoreBrowserSearchResults)
	if !ok {
		return p, fmt.Errorf("search(%s) result type=%T", pat, info.Result)
	}

	if len(res.File) != 1 {
		return p, fmt.Errorf("search(%s) result files=%d", pat, len(res.File))
	}

	return path.Join(dir, res.File[0].GetFileInfo().Path), nil
}
