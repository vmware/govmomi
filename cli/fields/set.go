// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package fields

import (
	"context"
	"flag"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
)

type set struct {
	*flags.DatacenterFlag
	add  bool
	kind string
}

func init() {
	cli.Register("fields.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.kind, "type", "", "Managed object type on which to add "+
		"the field if it does not exist. This flag is ignored unless -add=true")
	f.BoolVar(&cmd.add, "add", false, "Adds the field if it does not exist. "+
		"Use the -type flag to specify the managed object type to which the "+
		"field is added. Using -add and omitting -kind causes a new, global "+
		"field to be created if a field with the provided name does not "+
		"already exist.")
}

func (cmd *set) Usage() string {
	return "KEY VALUE PATH..."
}

func (cmd *set) Description() string {
	return `Set custom field values for PATH.

Examples:
  govc fields.set my-field-name field-value vm/my-vm
  govc fields.set -add my-new-global-field-name field-value vm/my-vm
  govc fields.set -add -type VirtualMachine my-new-vm-field-name field-value vm/my-vm`
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() < 3 {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetCustomFieldsManager(c)
	if err != nil {
		return err
	}

	args := f.Args()

	key, err := m.FindKey(ctx, args[0])
	if err != nil {
		if !(cmd.add && strings.Contains(err.Error(), "key name not found")) {
			return err
		}
		// Add the missing field.
		def, err := m.Add(ctx, args[0], cmd.kind, nil, nil)
		if err != nil {
			return err
		}
		// Assign the new field's key to the "key" var used below when
		// setting the key/value pair on the provided list of objects.
		key = def.Key
	}

	val := args[1]

	objs, err := cmd.ManagedObjects(ctx, args[2:])
	if err != nil {
		return err
	}

	for _, ref := range objs {
		err := m.Set(ctx, ref, key, val)
		if err != nil {
			return err
		}
	}

	return nil
}
