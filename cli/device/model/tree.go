// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"sort"
	"unsafe"

	"github.com/xlab/treeprint"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type tree struct {
	backings bool
}

func init() {
	cli.Register("device.model.tree", &tree{})
}

func (cmd *tree) Register(ctx context.Context, f *flag.FlagSet) {
	f.BoolVar(&cmd.backings, "backings", false, "Print the devices backings")
}

func (cmd *tree) Description() string {
	return `Print the device model as a tree.

Examples:
  govc device.model.tree
  govc device.model.tree VirtualEthernetCard
  govc device.model.tree -backings
  govc device.model.tree -backings VirtualDiskRawDiskVer2BackingInfo`
}

func (cmd *tree) Process(ctx context.Context) error {
	return nil
}

func (cmd *tree) Run(ctx context.Context, f *flag.FlagSet) error {

	typeName := f.Arg(0)

	var node treeprint.Tree
	if cmd.backings {
		node = getTree[
			types.BaseVirtualDeviceBackingInfo,
			types.VirtualDeviceBackingInfo,
		]()
	} else {
		node = getTree[
			types.BaseVirtualDevice,
			types.VirtualDevice,
		]()
	}

	if typeName != "" {
		var found treeprint.Tree
		node.VisitAll(func(n *treeprint.Node) {
			if n.Value == typeName {
				found = n
			}
		})
		if found == nil {
			return fmt.Errorf("%q not found", typeName)
		}
		node = found
		node = node.Branch()
	}

	fmt.Print(node.String())

	return nil
}

//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname add reflect.add
func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer

// typeEmbeds returns true if a embeds any of the b's.
func typeEmbeds(a reflect.Type, b ...reflect.Type) bool {
	for i := range b {
		if _, ok := a.FieldByName(b[i].Name()); ok {
			return true
		}
	}
	return false
}

// typeEmbeddedBy returns true if a is embedded by any of the b's.
func typeEmbeddedBy(a reflect.Type, b ...reflect.Type) bool {
	for i := range b {
		if _, ok := b[i].FieldByName(a.Name()); ok {
			return true
		}
	}
	return false
}

func getTree[T, K any]() treeprint.Tree {
	var (
		rootObj       K
		rootType      = reflect.TypeOf(rootObj)
		allTypes      = []reflect.Type{}
		embedsTypeMap = map[reflect.Type][]reflect.Type{}
		rootIfaceType = reflect.TypeOf((*T)(nil)).Elem()
	)

	sections, offsets := typelinks()
	for i := range sections {
		base := sections[i]
		for _, offset := range offsets[i] {
			typeAddr := add(base, uintptr(offset), "")
			typ3 := reflect.TypeOf(*(*any)(unsafe.Pointer(&typeAddr)))
			if typ3.Implements(rootIfaceType) {
				realType := reflect.Zero(typ3.Elem()).Type()
				allTypes = append(allTypes, realType)
				embedsTypeMap[realType] = []reflect.Type{}
			}
		}
	}

	// Create the child->parents map.
	for i := range allTypes {
		a := allTypes[i]
		for b := range embedsTypeMap {
			if typeEmbeds(a, b) {
				embedsTypeMap[a] = append(embedsTypeMap[a], b)
			}
		}
	}

	// Each child should have a single parent.
	for child, parents := range embedsTypeMap {
		notAncestors := []reflect.Type{}
		for i := range parents {
			p := parents[i]
			if !typeEmbeddedBy(p, parents...) {
				notAncestors = append(notAncestors, p)
			}
		}
		embedsTypeMap[child] = notAncestors
	}

	// Create the parent->children map.
	typeMap := map[string][]string{}
	for child, parents := range embedsTypeMap {
		for i := range parents {
			p := parents[i]
			typeMap[p.Name()] = append(typeMap[p.Name()], child.Name())
		}
	}

	// Sort the children for each parent by name.
	for _, children := range typeMap {
		sort.Strings(children)
	}

	var buildTree func(parent string, tree treeprint.Tree) treeprint.Tree
	buildTree = func(parent string, tree treeprint.Tree) treeprint.Tree {
		children := typeMap[parent]
		for i := range children {
			child := children[i]
			if _, childIsParentToo := typeMap[child]; childIsParentToo {
				buildTree(child, tree.AddBranch(child))
			} else {
				tree.AddNode(child)
			}
		}
		return tree
	}

	return buildTree(rootType.Name(), treeprint.NewWithRoot(rootType.Name()))
}
