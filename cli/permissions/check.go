// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package permissions

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25/types"
)

type check struct {
	*PermissionFlag

	principal string
	privs     string
}

func init() {
	cli.Register("permissions.check", &check{})
}

func (cmd *check) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.PermissionFlag, ctx = NewPermissionFlag(ctx)
	cmd.PermissionFlag.Register(ctx, f)

	f.StringVar(&cmd.principal, "principal", "", "User name (defaults to the current session)")
	f.StringVar(&cmd.privs, "priv", "", "Comma separated list of privilege IDs")
}

func (cmd *check) Process(ctx context.Context) error {
	return cmd.PermissionFlag.Process(ctx)
}

func (cmd *check) Usage() string {
	return "PATH..."
}

func (cmd *check) Description() string {
	return `Check whether privileges are granted on entity PATH.

Checks the current session's privileges, or those of PRINCIPAL if specified.

Examples:
  govc permissions.check -priv System.Read,VirtualMachine.Inventory.Create /dc1
  govc permissions.check -principal user@vsphere.local -priv System.Read /dc1/host/cluster1
  govc permissions.check -json -priv Datastore.AllocateSpace /dc1/datastore/ds1`
}

type checkResult struct {
	EntityPrivilege []types.EntityPrivilege `json:"entityPrivilege"`

	paths map[types.ManagedObjectReference]string
}

func (r *checkResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Entity", "Privilege", "Granted")

	for _, e := range r.EntityPrivilege {
		for _, p := range e.PrivAvailability {
			granted := "No"
			if p.IsGranted {
				granted = "Yes"
			}
			fmt.Fprintf(tw, "%s\t%s\t%s\n", r.paths[e.Entity], p.PrivId, granted)
		}
	}

	return tw.Flush()
}

func (cmd *check) entityPaths(ctx context.Context, refs []types.ManagedObjectReference, paths map[types.ManagedObjectReference]string) error {
	if cmd.asRef {
		for _, ref := range refs {
			paths[ref] = ref.String()
		}
		return nil
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	for _, ref := range refs {
		e, err := finder.Element(ctx, ref)
		if err != nil {
			return err
		}
		paths[ref] = e.Path
	}

	return nil
}

func (cmd *check) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() == 0 || cmd.privs == "" {
		return flag.ErrHelp
	}

	privs := strings.Split(cmd.privs, ",")

	refs, err := cmd.ManagedObjects(ctx, f.Args())
	if err != nil {
		return err
	}

	m, err := cmd.Manager(ctx)
	if err != nil {
		return err
	}

	res := checkResult{paths: make(map[types.ManagedObjectReference]string)}
	if err := cmd.entityPaths(ctx, refs, res.paths); err != nil {
		return err
	}

	if cmd.principal != "" {
		res.EntityPrivilege, err = m.HasUserPrivilegeOnEntities(ctx, refs, cmd.principal, privs)
		if err != nil {
			return err
		}
		return cmd.WriteResult(&res)
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	s, err := session.NewManager(c).UserSession(ctx)
	if err != nil {
		return err
	}
	if s == nil {
		return fmt.Errorf("no active session")
	}

	for _, ref := range refs {
		granted, err := m.HasPrivilegeOnEntity(ctx, ref, s.Key, privs)
		if err != nil {
			return err
		}

		e := types.EntityPrivilege{Entity: ref}
		for i, id := range privs {
			e.PrivAvailability = append(e.PrivAvailability, types.PrivilegeAvailability{PrivId: id, IsGranted: granted[i]})
		}
		res.EntityPrivilege = append(res.EntityPrivilege, e)
	}

	return cmd.WriteResult(&res)
}
