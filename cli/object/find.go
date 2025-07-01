// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type find struct {
	*flags.DatacenterFlag

	ref      bool
	id       bool
	long     bool
	parent   bool
	kind     kinds
	name     string
	maxdepth int
}

var alias = []struct {
	name string
	kind string
}{
	{"a", "VirtualApp"},
	{"c", "ClusterComputeResource"},
	{"d", "Datacenter"},
	{"f", "Folder"},
	{"g", "DistributedVirtualPortgroup"},
	{"h", "HostSystem"},
	{"m", "VirtualMachine"},
	{"n", "Network"},
	{"o", "OpaqueNetwork"},
	{"p", "ResourcePool"},
	{"r", "ComputeResource"},
	{"s", "Datastore"},
	{"w", "DistributedVirtualSwitch"},
}

func aliasHelp() string {
	var help bytes.Buffer

	for _, a := range alias {
		fmt.Fprintf(&help, "  %s    %s\n", a.name, a.kind)
	}

	return help.String()
}

type kinds []string

func (e *kinds) String() string {
	return fmt.Sprint(*e)
}

func (e *kinds) Set(value string) error {
	*e = append(*e, e.alias(value))
	return nil
}

func (e *kinds) alias(value string) string {
	if len(value) != 1 {
		return value
	}

	for _, a := range alias {
		if a.name == value {
			return a.kind
		}
	}

	return value
}

func (e *kinds) wanted(kind string) bool {
	if len(*e) == 0 {
		return true
	}

	for _, k := range *e {
		if kind == k {
			return true
		}
	}

	return false
}

func init() {
	cli.Register("find", &find{})
}

func (cmd *find) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.Var(&cmd.kind, "type", "Resource type")
	f.StringVar(&cmd.name, "name", "*", "Resource name")
	f.IntVar(&cmd.maxdepth, "maxdepth", -1, "Max depth")
	f.BoolVar(&cmd.ref, "i", false, "Print the managed object reference")
	f.BoolVar(&cmd.id, "I", false, "Print the managed object ID")
	f.BoolVar(&cmd.long, "l", false, "Long listing format")
	f.BoolVar(&cmd.parent, "p", false, "Find parent objects")
}

func (cmd *find) Usage() string {
	return "[ROOT] [KEY VAL]..."
}

func (cmd *find) Description() string {
	atable := aliasHelp()

	return fmt.Sprintf(`Find managed objects.

ROOT can be an inventory path or ManagedObjectReference.
ROOT defaults to '.', an alias for the root folder or DC if set.

Optional KEY VAL pairs can be used to filter results against object instance properties.
Use the govc 'collect' command to view possible object property keys.

The '-type' flag value can be a managed entity type or one of the following aliases:

%s
Examples:
  govc find
  govc find -l / # include object type in output
  govc find -l -I / # include MOID in output
  govc find /dc1 -type c
  govc find vm -name my-vm-*
  govc find . -type n
  govc find -p /folder-a/dc-1/host/folder-b/cluster-a -type Datacenter # prints /folder-a/dc-1
  govc find . -type m -runtime.powerState poweredOn
  govc find . -type m -datastore $(govc find -i datastore -name vsanDatastore)
  govc find . -type s -summary.type vsan
  govc find . -type s -customValue *:prod # Key:Value
  govc find . -type h -hardware.cpuInfo.numCpuCores 16`, atable)
}

// rootMatch returns true if the root object path should be printed
func (cmd *find) rootMatch(ctx context.Context, root object.Reference, client *vim25.Client, filter property.Match) bool {
	ref := root.Reference()

	if !cmd.kind.wanted(ref.Type) {
		return false
	}

	if len(filter) == 1 && filter["name"] == "*" {
		return true
	}

	var content []types.ObjectContent

	pc := property.DefaultCollector(client)
	_ = pc.RetrieveWithFilter(ctx, []types.ManagedObjectReference{ref}, filter.Keys(), &content, filter)

	return content != nil
}

type findResult []string

func (r findResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

func (r findResult) Dump() any {
	return []string(r)
}

type findResultLong []string

func (r findResultLong) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for i := range r {
		fmt.Fprintln(tw, r[i])
	}
	return tw.Flush()
}

func (cmd *find) writeResult(paths []string) error {
	if cmd.long {
		return cmd.WriteResult(findResultLong(paths))
	}
	return cmd.WriteResult(findResult(paths))
}

func (cmd *find) mo(o types.ManagedObjectReference) string {
	if cmd.id {
		return o.Value
	}
	return o.String()
}

func (cmd *find) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	if cmd.id {
		cmd.ref = true
	}

	root := client.ServiceContent.RootFolder
	rootPath := "/"

	arg := f.Arg(0)
	props := f.Args()

	if len(props) > 0 {
		if strings.HasPrefix(arg, "-") {
			arg = "."
		} else {
			props = props[1:]
		}
	}

	if len(props)%2 != 0 {
		return flag.ErrHelp
	}

	dc, err := cmd.DatacenterIfSpecified()
	if err != nil {
		return err
	}

	switch arg {
	case rootPath:
	case "", ".":
		if dc == nil {
			arg = rootPath
		} else {
			arg = "."
			root = dc.Reference()
			rootPath = dc.InventoryPath
		}
	default:
		path := arg
		if !strings.Contains(arg, "/") {
			// Force list mode
			p := "."
			if dc != nil {
				p = dc.InventoryPath
			}
			path = strings.Join([]string{p, arg}, "/")
		}

		l, ferr := finder.ManagedObjectList(ctx, path)
		if ferr != nil {
			return err
		}

		switch len(l) {
		case 0:
			return fmt.Errorf("%s not found", arg)
		case 1:
			root = l[0].Object.Reference()
			rootPath = l[0].Path
		default:
			return fmt.Errorf("%q matches %d objects", arg, len(l))
		}
	}

	filter := property.Match{}

	if len(props)%2 != 0 {
		return flag.ErrHelp
	}

	for i := 0; i < len(props); i++ {
		key := props[i]
		if !strings.HasPrefix(key, "-") {
			return flag.ErrHelp
		}

		key = key[1:]
		i++
		val := props[i]

		if xf := f.Lookup(key); xf != nil {
			// Support use of -flag following the ROOT arg (flag package does not do this)
			if err = xf.Value.Set(val); err != nil {
				return err
			}
		} else {
			filter[key] = val
		}
	}

	filter["name"] = cmd.name
	var paths []string

	printPath := func(o types.ManagedObjectReference, p string) {
		if cmd.ref && !cmd.long {
			paths = append(paths, cmd.mo(o))
			return
		}

		path := strings.Replace(p, rootPath, arg, 1)
		if cmd.long {
			id := strings.TrimPrefix(o.Type, "Vmware")
			if cmd.ref {
				id = cmd.mo(o)
			}

			path = id + "\t" + path
		}
		paths = append(paths, path)
	}

	recurse := false

	switch cmd.maxdepth {
	case -1:
		recurse = true
	case 0:
	case 1:
	default:
		return flag.ErrHelp // TODO: ?
	}

	if cmd.parent {
		entities, err := mo.Ancestors(ctx, client, client.ServiceContent.PropertyCollector, root)
		if err != nil {
			return err
		}

		for i := len(entities) - 1; i >= 0; i-- {
			if cmd.rootMatch(ctx, entities[i], client, filter) {
				printPath(entities[i].Reference(), internal.InventoryPath(entities[:i+1]))
			}
		}

		return cmd.writeResult(paths)
	}

	if cmd.rootMatch(ctx, root, client, filter) {
		printPath(root, arg)
	}

	if cmd.maxdepth == 0 {
		return cmd.writeResult(paths)
	}

	m := view.NewManager(client)

	v, err := m.CreateContainerView(ctx, root, cmd.kind, recurse)
	if err != nil {
		return err
	}

	defer func() {
		_ = v.Destroy(ctx)
	}()

	objs, err := v.Find(ctx, cmd.kind, filter)
	if err != nil {
		return err
	}

	for _, o := range objs {
		var path string

		if cmd.long || !cmd.ref {
			e, err := finder.Element(ctx, o)
			if err != nil {
				if fault.Is(err, &types.ManagedObjectNotFound{}) {
					continue // object was deleted after v.Find() returned
				}
				return err
			}
			path = e.Path
		}

		printPath(o, path)
	}

	return cmd.writeResult(paths)
}
