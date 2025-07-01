// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package snapshot

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
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

type SnapshotRecord struct {
	CreateTime        *time.Time       `json:"createTime,omitempty"`
	Id                *string          `json:"id,omitempty"`
	Size              *int             `json:"size,omitempty"`
	Name              string           `json:"name"`
	Description       *string          `json:"description,omitempty"`
	IsCurrent         bool             `json:"current"`
	ChildSnapshotList []SnapshotRecord `json:"childSnapshotList"`
}

func (cmd *tree) IsCurrent(vm mo.VirtualMachine, moref types.ManagedObjectReference) bool {
	return vm.Snapshot.CurrentSnapshot.Value == moref.Value
}

func (cmd *tree) CreateTime(snapshot types.VirtualMachineSnapshotTree) *time.Time {
	if cmd.date {
		return &snapshot.CreateTime
	}
	return nil

}

func (cmd *tree) SnapshotId(snapshot types.VirtualMachineSnapshotTree) *string {
	if cmd.id {
		return &snapshot.Snapshot.Value
	}
	return nil
}

func (cmd *tree) SnapshotDescription(snapshot types.VirtualMachineSnapshotTree) *string {
	if cmd.description {
		return &snapshot.Description
	}
	return nil
}

func (cmd *tree) SnapshotSize(vm mo.VirtualMachine, snapshot types.ManagedObjectReference, parent *types.ManagedObjectReference) *int {
	if cmd.size {
		size := object.SnapshotSize(snapshot, parent, vm.LayoutEx, cmd.IsCurrent(vm, snapshot))
		return &size
	}
	return nil
}

func (cmd *tree) makeSnapshotRecord(vm mo.VirtualMachine, node types.VirtualMachineSnapshotTree, parent *types.ManagedObjectReference) SnapshotRecord {

	var SnapshotRecords []SnapshotRecord
	for _, snapshot := range node.ChildSnapshotList {
		SnapshotRecords = append(SnapshotRecords, cmd.makeSnapshotRecord(vm, snapshot, &node.Snapshot))
	}
	return SnapshotRecord{Name: node.Name,
		Id:                cmd.SnapshotId(node),
		CreateTime:        cmd.CreateTime(node),
		Description:       cmd.SnapshotDescription(node),
		ChildSnapshotList: SnapshotRecords,
		Size:              cmd.SnapshotSize(vm, node.Snapshot, parent),
		IsCurrent:         cmd.IsCurrent(vm, node.Snapshot)}

}

func (cmd *tree) writeJson(vm mo.VirtualMachine) {
	var SnapshotRecords []SnapshotRecord
	for _, rootSnapshot := range vm.Snapshot.RootSnapshotList {
		SnapshotRecords = append(SnapshotRecords, cmd.makeSnapshotRecord(vm, rootSnapshot, nil))
	}
	b, _ := json.MarshalIndent(SnapshotRecords, "", "  ")
	fmt.Println(string(b))

}
func (cmd *tree) write(level int, parent string, pref *types.ManagedObjectReference, st []types.VirtualMachineSnapshotTree) {
	for _, s := range st {
		s := s // avoid implicit memory aliasing

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
				attr = append(attr, s.CreateTime.Format("2006-01-02T15:04:05Z07:00"))
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
	if cmd.JSON {
		cmd.writeJson(o)
	} else {
		cmd.write(0, "", nil, o.Snapshot.RootSnapshotList)
	}

	return nil
}
