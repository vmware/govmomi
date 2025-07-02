// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.DatacenterFlag

	long bool
	link bool
	url  bool
	stor bool
	Stor bool

	pathFinder *finder.PathFinder
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
	f.BoolVar(&cmd.stor, "s", false, "Include file specific storage details")
	if cli.ShowUnreleased() {
		f.BoolVar(&cmd.Stor, "S", false, "Include file specific storage details (resolved)")
	}
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Description() string {
	return `Display library information.

Note: the '-s' flag only applies to files, not items or the library itself.

Examples:
  govc library.info
  govc library.info /lib1
  govc library.info -l /lib1 | grep Size:
  govc library.info /lib1/item1
  govc library.info /lib1/item1/
  govc library.info */
  govc library.info -L /lib1/item1/file1 # file path relative to Datastore
  govc library.info -L -l /lib1/item1/file1 # file path including Datastore
  govc library.info -json | jq .
  govc library.info -json /lib1/item1 | jq .`
}

type infoResultsWriter struct {
	Result []finder.FindResult `json:"result"`
	m      *library.Manager
	cmd    *info
	ctx    context.Context
}

func (r infoResultsWriter) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Result)
}

func (r infoResultsWriter) Dump() any {
	res := make([]any, len(r.Result))
	for i := range r.Result {
		res[i] = r.Result[0].GetResult()
	}
	return res
}

func (r infoResultsWriter) Write(w io.Writer) error {
	if r.cmd.link {
		for _, j := range r.Result {
			p, err := r.cmd.pathFinder.Path(context.Background(), j)
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
	if v.Description != nil {
		fmt.Fprintf(w, "  Description:\t%s\n", *v.Description)
	}
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)
	fmt.Fprintf(w, "  Created:\t%s\n", v.CreationTime.Format(time.ANSIC))
	fmt.Fprintf(w, "  Security Policy ID\t%s\n", v.SecurityPolicyID)
	fmt.Fprintf(w, "  StorageBackings:\t\n")
	for _, d := range v.Storage {
		fmt.Fprintf(w, "    DatastoreID:\t%s\n", d.DatastoreID)
		fmt.Fprintf(w, "    Type:\t%s\n", d.Type)
	}
	if r.cmd.long {
		p, err := r.cmd.pathFinder.Path(r.ctx, res)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "  Datastore Path:\t%s\n", p)
		items, err := r.m.GetLibraryItems(r.ctx, v.ID)
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
	if v.Description != nil {
		fmt.Fprintf(w, "  Description:\t%s\n", *v.Description)
	}
	fmt.Fprintf(w, "  Type:\t%s\n", v.Type)
	fmt.Fprintf(w, "  Size:\t%s\n", units.ByteSize(v.Size))
	fmt.Fprintf(w, "  Cached:\t%t\n", v.Cached)
	fmt.Fprintf(w, "  Created:\t%s\n", v.CreationTime.Format(time.ANSIC))
	fmt.Fprintf(w, "  Modified:\t%s\n", v.LastModifiedTime.Format(time.ANSIC))
	fmt.Fprintf(w, "  Version:\t%s\n", v.Version)
	if v.SecurityCompliance != nil {
		fmt.Fprintf(w, "  Security Compliance:\t%t\n", *v.SecurityCompliance)
	}
	if v.CertificateVerification != nil {
		fmt.Fprintf(w, "  Certificate Status:\t%s\n", v.CertificateVerification.Status)
	}
	if r.cmd.long {
		p, err := r.cmd.pathFinder.Path(r.ctx, res)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "  Datastore Path:\t%s\n", p)
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
		p, err := r.cmd.pathFinder.Path(r.ctx, res)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "  Datastore Path:\t%s\n", p)
	}
	if r.cmd.stor || r.cmd.Stor {
		label := "Storage URI"
		s, err := r.m.GetLibraryItemStorage(r.ctx, res.GetParent().GetID(), v.Name)
		if err != nil {
			return err
		}
		if r.cmd.Stor {
			label = "Resolved URI"
			err = r.cmd.pathFinder.ResolveLibraryItemStorage(r.ctx, nil, nil, s)
			if err != nil {
				return err
			}
		}
		for i := range s {
			for _, uri := range s[i].StorageURIs {
				fmt.Fprintf(w, "  %s:\t%s\n", label, uri)
			}
		}
	}

	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	m := library.NewManager(c)
	cmd.pathFinder = finder.NewPathFinder(m, vc)
	finder := finder.NewFinder(m)
	findResults, err := finder.Find(ctx, f.Args()...)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&infoResultsWriter{findResults, m, cmd, ctx})
}
