// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/list"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.DatacenterFlag

	Long  bool
	kind  kinds
	ToRef bool
	ToID  bool
	DeRef bool
}

func init() {
	cli.Register("ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.BoolVar(&cmd.Long, "l", false, "Long listing format")
	f.BoolVar(&cmd.ToRef, "i", false, "Print the managed object reference")
	f.BoolVar(&cmd.ToID, "I", false, "Print the managed object ID")
	f.BoolVar(&cmd.DeRef, "L", false, "Follow managed object references")
	f.Var(&cmd.kind, "t", "Object type")
}

func (cmd *ls) Description() string {
	return fmt.Sprintf(`List inventory items.

The '-t' flag value can be a managed entity type or one of the following aliases:
%s
Examples:
  govc ls -l '*'
  govc ls -t ClusterComputeResource host
  govc ls -t Datastore host/ClusterA/* | grep -v local | xargs -n1 basename | sort | uniq`, aliasHelp())
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return "[PATH]..."
}

func (cmd *ls) typeMatch(ref types.ManagedObjectReference) bool {
	if len(cmd.kind) == 0 {
		return true
	}

	for _, kind := range cmd.kind {
		if strings.EqualFold(kind, ref.Type) {
			return true
		}
	}

	return false
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	finder, err := cmd.Finder(cmd.All())
	if err != nil {
		return err
	}

	lr := listResult{
		ls:       cmd,
		Elements: nil,
	}

	args := f.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	var ref = new(types.ManagedObjectReference)

	for _, arg := range args {
		if cmd.DeRef && ref.FromString(arg) {
			e, err := finder.Element(ctx, *ref)
			if err == nil {
				if cmd.typeMatch(*ref) {
					if e.Path == "/" && ref.Type != "Folder" {
						// Special case: when given a moref with no ancestors,
						// just echo the moref.
						e.Path = ref.String()
					}
					lr.Elements = append(lr.Elements, *e)
				}
				continue
			}
		}

		es, err := finder.ManagedObjectListChildren(ctx, arg, cmd.kind...)
		if err != nil {
			return err
		}

		for _, e := range es {
			if cmd.typeMatch(e.Object.Reference()) {
				lr.Elements = append(lr.Elements, e)
			}
		}
	}

	return cmd.WriteResult(lr)
}

type listResult struct {
	*ls      `json:"-" xml:"-"`
	Elements []list.Element `json:"elements"`
}

func (l listResult) Write(w io.Writer) error {
	var err error

	for _, e := range l.Elements {
		if l.ToRef || l.ToID {
			ref := e.Object.Reference()
			id := ref.String()
			if l.ToID {
				id = ref.Value
			}
			fmt.Fprint(w, id)
			if l.Long {
				fmt.Fprintf(w, " %s", e.Path)
			}
			fmt.Fprintln(w)
			continue
		}

		if !l.Long {
			fmt.Fprintf(w, "%s\n", e.Path)
			continue
		}

		switch e.Object.(type) {
		case mo.Folder:
			if _, err = fmt.Fprintf(w, "%s/\n", e.Path); err != nil {
				return err
			}
		default:
			if _, err = fmt.Fprintf(w, "%s (%s)\n", e.Path, e.Object.Reference().Type); err != nil {
				return err
			}
		}
	}

	return nil
}
