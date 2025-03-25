// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package datastore

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag
	*flags.DatacenterFlag

	host bool
}

func init() {
	cli.Register("datastore.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.BoolVar(&cmd.host, "H", false, "Display info for Datastores shared between hosts")
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *info) Usage() string {
	return "[PATH]..."
}

func (cmd *info) Description() string {
	return `Display info for Datastores.

Examples:
  govc datastore.info
  govc datastore.info vsanDatastore
  # info on Datastores shared between cluster hosts:
  govc collect -s -d " " /dc1/host/k8s-cluster host | xargs govc datastore.info -H
  # info on Datastores shared between VM hosts:
  govc ls /dc1/vm/*k8s* | xargs -n1 -I% govc collect -s % summary.runtime.host | xargs govc datastore.info -H`
}

func intersect(common []types.ManagedObjectReference, refs []types.ManagedObjectReference) []types.ManagedObjectReference {
	var shared []types.ManagedObjectReference
	for i := range common {
		for j := range refs {
			if common[i] == refs[j] {
				shared = append(shared, common[i])
				break
			}
		}
	}
	return shared
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}
	pc := property.DefaultCollector(c)

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	args := f.Args()
	if len(args) == 0 {
		args = []string{"*"}
	}

	var res infoResult
	var props []string

	if cmd.OutputFlag.All() {
		props = nil // Load everything
	} else {
		props = []string{"info", "summary"} // Load summary
	}

	if cmd.host {
		if f.NArg() == 0 {
			return flag.ErrHelp
		}
		refs, err := cmd.ManagedObjects(ctx, args)
		if err != nil {
			return err
		}

		var hosts []mo.HostSystem
		err = pc.Retrieve(ctx, refs, []string{"name", "datastore"}, &hosts)
		if err != nil {
			return err
		}

		refs = hosts[0].Datastore
		for _, host := range hosts[1:] {
			refs = intersect(refs, host.Datastore)
			if len(refs) == 0 {
				return fmt.Errorf("host %s (%s) has no shared datastores", host.Name, host.Reference())
			}
		}
		for i := range refs {
			ds, err := finder.ObjectReference(ctx, refs[i])
			if err != nil {
				return err
			}
			res.objects = append(res.objects, ds.(*object.Datastore))
		}
	} else {
		for _, arg := range args {
			objects, err := finder.DatastoreList(ctx, arg)
			if err != nil {
				return err
			}
			res.objects = append(res.objects, objects...)
		}
	}

	if len(res.objects) != 0 {
		refs := make([]types.ManagedObjectReference, 0, len(res.objects))
		for _, o := range res.objects {
			refs = append(refs, o.Reference())
		}

		err = pc.Retrieve(ctx, refs, props, &res.Datastores)
		if err != nil {
			return err
		}
	}

	return cmd.WriteResult(&res)
}

type infoResult struct {
	Datastores []mo.Datastore `json:"datastores"`
	objects    []*object.Datastore
}

func (r *infoResult) Write(w io.Writer) error {
	// Maintain order via r.objects as Property collector does not always return results in order.
	objects := make(map[types.ManagedObjectReference]mo.Datastore, len(r.Datastores))
	for _, o := range r.Datastores {
		objects[o.Reference()] = o
	}

	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	for _, o := range r.objects {
		ds := objects[o.Reference()]
		s := ds.Summary
		fmt.Fprintf(tw, "Name:\t%s\n", s.Name)
		fmt.Fprintf(tw, "  Path:\t%s\n", o.InventoryPath)
		fmt.Fprintf(tw, "  Type:\t%s\n", s.Type)
		fmt.Fprintf(tw, "  URL:\t%s\n", s.Url)
		fmt.Fprintf(tw, "  Capacity:\t%.1f GB\n", float64(s.Capacity)/(1<<30))
		fmt.Fprintf(tw, "  Free:\t%.1f GB\n", float64(s.FreeSpace)/(1<<30))

		switch info := ds.Info.(type) {
		case *types.NasDatastoreInfo:
			fmt.Fprintf(tw, "  Remote:\t%s:%s\n", info.Nas.RemoteHost, info.Nas.RemotePath)
		}
	}

	return tw.Flush()
}
