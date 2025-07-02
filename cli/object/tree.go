// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	gopath "path"
	"time"

	gotree "github.com/a8m/tree"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

type tree struct {
	*flags.DatacenterFlag

	long  bool
	kind  bool
	color bool
	level int
}

func init() {
	cli.Register("tree", &tree{})
}

func (cmd *tree) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.BoolVar(&cmd.color, "C", false, "Colorize output")
	f.BoolVar(&cmd.long, "l", false, "Follow runtime references (e.g. HostSystem VMs)")
	f.BoolVar(&cmd.kind, "p", false, "Print the object type")
	f.IntVar(&cmd.level, "L", 0, "Max display depth of the inventory tree")
}

func (cmd *tree) Description() string {
	return `List contents of the inventory in a tree-like format.

Examples:
  govc tree -C /
  govc tree /datacenter/vm`
}

func (cmd *tree) Usage() string {
	return "[PATH]"
}

func (cmd *tree) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	path := f.Arg(0)
	if path == "" {
		path = "/"
	}

	vfs := &virtualFileSystem{
		ctx:   ctx,
		cmd:   cmd,
		c:     c,
		m:     view.NewManager(c),
		names: make(map[types.ManagedObjectReference]string),
		dvs:   make(map[types.ManagedObjectReference][]types.ManagedObjectReference),
		path:  path,
	}

	treeOpts := &gotree.Options{
		Fs:        vfs,
		OutFile:   cmd.Out,
		Colorize:  cmd.color,
		Color:     color,
		DeepLevel: cmd.level,
	}

	inf := gotree.New(path)
	inf.Visit(treeOpts)
	inf.Print(treeOpts)

	return nil
}

type virtualFileSystem struct {
	ctx   context.Context
	cmd   *tree
	c     *vim25.Client
	m     *view.Manager
	names map[types.ManagedObjectReference]string
	dvs   map[types.ManagedObjectReference][]types.ManagedObjectReference
	root  types.ManagedObjectReference
	path  string
}

func style(kind string) string {
	switch kind {
	case "VirtualMachine":
		return "1;32"
	case "HostSystem":
		return "1;33"
	case "ResourcePool":
		return "1;30"
	case "Network", "OpaqueNetwork", "DistributedVirtualPortgroup":
		return "1;35"
	case "Datastore":
		return "1;36"
	case "Datacenter":
		return "1;37"
	default:
		return ""
	}
}

func color(node *gotree.Node, s string) string {
	ref := pathReference(node.Path())

	switch ref.Type {
	case "ResourcePool":
		return s
	}

	c := style(ref.Type)
	if c == "" {
		return gotree.ANSIColor(node, s)
	}

	return gotree.ANSIColorFormat(c, s)
}

func (vfs *virtualFileSystem) Stat(path string) (os.FileInfo, error) {
	var ref types.ManagedObjectReference

	if len(vfs.names) == 0 {
		// This is the first Stat() call, where path is the initial user input
		if path == "/" {
			ref = vfs.c.ServiceContent.RootFolder
		} else {
			var err error
			ref, err = vfs.cmd.ManagedObject(vfs.ctx, path)
			if err != nil {
				return nil, err
			}
		}
		vfs.names[ref] = path
		vfs.root = ref
	} else {
		// The Node.Path in subsequent calls to Stat() will have a MOR base
		ref = pathReference(path)
	}

	name := vfs.names[ref]

	var mode os.FileMode
	switch ref.Type {
	case "ComputeResource",
		"ClusterComputeResource",
		"Datacenter",
		"Folder",
		"ResourcePool",
		"VirtualApp",
		"StoragePod",
		"DistributedVirtualSwitch",
		"VmwareDistributedVirtualSwitch":
		mode = os.ModeDir
	case "HostSystem":
		if vfs.cmd.long {
			mode = os.ModeDir
		}
	}

	if vfs.cmd.kind {
		name = fmt.Sprintf("[%s] %s", ref.Type, name)
	}

	return fileInfo{name: name, mode: mode}, nil
}

// pathReference converts the base of the given Node.Path to a MOR
func pathReference(s string) types.ManagedObjectReference {
	var ref types.ManagedObjectReference
	r, _ := url.PathUnescape(gopath.Base(s))
	ref.FromString(r)
	return ref
}

func (vfs *virtualFileSystem) ReadDir(path string) ([]string, error) {
	var ref types.ManagedObjectReference

	if path == vfs.path {
		// This path is the initial user input (e.g. "/" or "/dc1")
		ref = vfs.root
	} else {
		// This path will have had 1 or more MORs appended to it, as returned by this func
		ref = pathReference(path)
	}

	var childPaths []string

	switch ref.Type {
	// In the vCenter inventory switches and portgroups are siblings, hack to display them as parent child in the tree
	case "DistributedVirtualSwitch", "VmwareDistributedVirtualSwitch":
		pgs := vfs.dvs[ref]
		for _, pg := range pgs {
			childPaths = append(childPaths, url.PathEscape(pg.String()))
		}
		return childPaths, nil
	}

	v, err := vfs.m.CreateContainerView(vfs.ctx, ref, nil, false)
	if err != nil {
		return nil, err
	}
	defer v.Destroy(vfs.ctx)

	var kind []string
	if !vfs.cmd.long {
		switch ref.Type {
		case "HostSystem":
			return nil, nil
		case "ResourcePool", "VirtualApp":
			kind = []string{"ResourcePool", "VirtualApp"}
		}
	}

	var children []types.ObjectContent

	pspec := []types.PropertySpec{
		{Type: "DistributedVirtualSwitch", PathSet: []string{"portgroup"}},
		{Type: "VmwareDistributedVirtualSwitch", PathSet: []string{"portgroup"}},
	}

	err = v.Retrieve(vfs.ctx, kind, []string{"name"}, &children, pspec...)
	if err != nil {
		return nil, err
	}

	for _, content := range children {
		ref = content.Obj
		for _, p := range content.PropSet {
			switch p.Name {
			case "name":
				vfs.names[ref] = p.Val.(string)
			case "portgroup":
				vfs.dvs[ref] = p.Val.(types.ArrayOfManagedObjectReference).ManagedObjectReference
			}
		}
		if ref.Type == "DistributedVirtualPortgroup" {
			continue // Returned on ReadDir() of the DVS above
		}
		childPaths = append(childPaths, url.PathEscape(ref.String()))
	}

	return childPaths, nil
}

type fileInfo struct {
	name string
	mode os.FileMode
}

func (f fileInfo) Name() string {
	return f.name
}
func (f fileInfo) Size() int64 {
	return 0
}
func (f fileInfo) Mode() os.FileMode {
	return f.mode
}
func (f fileInfo) ModTime() time.Time {
	return time.Now()
}
func (f fileInfo) IsDir() bool {
	return f.mode&os.ModeDir == os.ModeDir
}
func (f fileInfo) Sys() any {
	return nil
}
