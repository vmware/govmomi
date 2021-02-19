/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package snapshot

import (
	"context"
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type tree struct {
	*flags.VirtualMachineFlag

	current     bool
	currentName bool
	date        bool
	description bool
	fullPath    bool
	id          bool
	size        bool

	info   *types.VirtualMachineSnapshotInfo
	layout *types.VirtualMachineFileLayoutEx
}

func init() {
	cli.Register("snapshot.tree", &tree{})
}

func (cmd *tree) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.BoolVar(&cmd.current, "c", true, "Print the current snapshot")
	f.BoolVar(&cmd.currentName, "C", false,
		"Print the current snapshot name only")
	f.BoolVar(&cmd.date, "D", false, "Print the snapshot creation date")
	f.BoolVar(&cmd.description, "d", false,
		"Print the snapshot description")
	f.BoolVar(&cmd.fullPath, "f", false,
		"Print the full path prefix for snapshot")
	f.BoolVar(&cmd.id, "i", false, "Print the snapshot id")
	f.BoolVar(&cmd.size, "s", false, "Print the snapshot size")
}

func (cmd *tree) Description() string {
	return `List VM snapshots in a tree-like format.

The command will exit 0 with no output if VM does not have any snapshots.

Examples:
  govc snapshot.tree -vm my-vm
  govc snapshot.tree -vm my-vm -D -i -d`
}

func (cmd *tree) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *tree) write(level int, parent string, pref *types.ManagedObjectReference, st []types.VirtualMachineSnapshotTree) {
	for _, s := range st {
		sname := s.Name

		if cmd.fullPath && parent != "" {
			sname = path.Join(parent, sname)
		}

		var names []string

		if !cmd.currentName {
			names = append(names, sname)
		}

		isCurrent := false

		if s.Snapshot == *cmd.info.CurrentSnapshot {
			isCurrent = true
			if cmd.current {
				names = append(names, ".")
			} else if cmd.currentName {
				fmt.Println(sname)
				return
			}
		}

		for _, name := range names {
			var attr []string
			var meta string

			if cmd.size {
				size := object.SnapshotSize(s.Snapshot, pref, cmd.layout, isCurrent)

				attr = append(attr, units.ByteSize(size).String())
			}

			if cmd.id {
				attr = append(attr, s.Snapshot.Value)
			}

			if cmd.date {
				attr = append(attr, s.CreateTime.Format("Jan 2 15:04"))
			}

			if len(attr) > 0 {
				meta = fmt.Sprintf("[%s]  ", strings.Join(attr, " "))
			}

			if cmd.description {
				fmt.Printf("%s%s%s - %4s\n",
					strings.Repeat(" ", level), meta, name,
					s.Description)
			} else {
				fmt.Printf("%s%s%s\n",
					strings.Repeat(" ", level), meta, name)
			}
		}

		cmd.write(level+2, sname, &s.Snapshot, s.ChildSnapshotList)
	}
}

func (cmd *tree) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 0 {
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	var o mo.VirtualMachine

	err = vm.Properties(ctx, vm.Reference(), []string{"snapshot", "layoutEx"}, &o)
	if err != nil {
		return err
	}

	if o.Snapshot == nil {
		return nil
	}

	if o.Snapshot.CurrentSnapshot == nil || cmd.currentName {
		cmd.current = false
	}

	cmd.info = o.Snapshot
	cmd.layout = o.LayoutEx

	cmd.write(0, "", nil, o.Snapshot.RootSnapshotList)

	return nil
}
