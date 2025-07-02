// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package disk

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/types"
	vslm "github.com/vmware/govmomi/vslm/types"
)

type ls struct {
	*flags.DatastoreFlag
	all      bool
	long     bool
	path     bool
	r        bool
	category string
	tag      string
	tags     bool
	query    flags.StringList
}

func init() {
	cli.Register("disk.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	f.BoolVar(&cmd.all, "a", false, "List IDs with missing file backing")
	f.BoolVar(&cmd.long, "l", false, "Long listing format")
	f.BoolVar(&cmd.path, "L", false, "Print disk backing path instead of disk name")
	f.BoolVar(&cmd.r, "R", false, "Reconcile the datastore inventory info")
	f.StringVar(&cmd.category, "c", "", "Query tag category")
	f.StringVar(&cmd.tag, "t", "", "Query tag name")
	f.BoolVar(&cmd.tags, "T", false, "List attached tags")
	f.Var(&cmd.query, "q", "Query spec")
}

func (cmd *ls) Usage() string {
	return "[ID]..."
}

func (cmd *ls) Description() string {
	var fields vslm.VslmVsoVStorageObjectQuerySpecQueryFieldEnum

	return fmt.Sprintf(`List disk IDs on DS.

The '-q' flag can be used to match disk fields.
Each query must be in the form of:
  FIELD.OP=VAL

Where FIELD can be one of:
  %s

And OP can be one of:
%s
Examples:
  govc disk.ls
  govc disk.ls -l -T
  govc disk.ls -l e9b06a8b-d047-4d3c-b15b-43ea9608b1a6
  govc disk.ls -c k8s-region -t us-west-2
  govc disk.ls -q capacity.ge=100 # capacity in MB
  govc disk.ls -q name.sw=my-disk
  govc disk.ls -q metadataKey.eq=cns.k8s.pvc.namespace -q metadataValue.eq=dev`,
		strings.Join(fields.Strings(), "\n  "),
		aliasHelp())
}

type VStorageObject struct {
	types.VStorageObject
	Tags []types.VslmTagEntry `json:"tags"`
}

func (o *VStorageObject) tags() string {
	var tags []string
	for _, tag := range o.Tags {
		tags = append(tags, tag.ParentCategoryName+":"+tag.TagName)
	}
	return strings.Join(tags, ",")
}

type lsResult struct {
	cmd     *ls
	Objects []VStorageObject `json:"objects"`
}

var alias = []struct {
	name string
	kind vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnum
}{
	{"eq", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals},
	{"ne", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumNotEquals},
	{"lt", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThan},
	{"le", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThanOrEqual},
	{"gt", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan},
	{"ge", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThanOrEqual},
	{"ct", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumContains},
	{"sw", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith},
	{"ew", vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEndsWith},
}

func opAlias(value string) string {
	if len(value) != 2 {
		return value
	}

	for _, a := range alias {
		if a.name == value {
			return string(a.kind)
		}
	}

	return value
}

func aliasHelp() string {
	var help bytes.Buffer

	for _, a := range alias {
		fmt.Fprintf(&help, "  %s    %s\n", a.name, a.kind)
	}

	return help.String()
}

func (cmd *ls) querySpec() ([]vslm.VslmVsoVStorageObjectQuerySpec, error) {
	q := make([]vslm.VslmVsoVStorageObjectQuerySpec, len(cmd.query))

	for i, s := range cmd.query {
		val := strings.SplitN(s, "=", 2)
		if len(val) != 2 {
			return nil, fmt.Errorf("invalid query: %s", s)
		}

		op := string(vslm.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals)
		field := strings.SplitN(val[0], ".", 2)
		if len(field) == 2 {
			op = field[1]
		}

		q[i] = vslm.VslmVsoVStorageObjectQuerySpec{
			QueryField:    field[0],
			QueryOperator: opAlias(op),
			QueryValue:    []string{val[1]},
		}
	}

	return q, nil
}

func (r *lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(r.cmd.Out, 2, 0, 2, ' ', 0)

	for _, o := range r.Objects {
		name := o.Config.Name
		if r.cmd.path {
			if file, ok := o.Config.Backing.(*types.BaseConfigInfoDiskFileBackingInfo); ok {
				name = file.FilePath
			}
		}
		_, _ = fmt.Fprintf(tw, "%s\t%s", o.Config.Id.Id, name)
		if r.cmd.long {
			created := o.Config.CreateTime.Format(time.Stamp)
			size := units.FileSize(o.Config.CapacityInMB * 1024 * 1024)
			_, _ = fmt.Fprintf(tw, "\t%s\t%s", size, created)
		}
		if r.cmd.tags {
			_, _ = fmt.Fprintf(tw, "\t%s", o.tags())
		}
		_, _ = fmt.Fprintln(tw)
	}

	return tw.Flush()
}

func (r *lsResult) Dump() any {
	return r.Objects
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	m, err := NewManagerFromFlag(ctx, cmd.DatastoreFlag)
	if err != nil {
		return err
	}

	if cmd.r {
		if err = m.ReconcileDatastoreInventory(ctx); err != nil {
			return err
		}
	}
	res := lsResult{cmd: cmd}

	filterNotFound := false
	ids := f.Args()
	q, err := cmd.querySpec()
	if err != nil {
		return err
	}

	if len(ids) == 0 {
		filterNotFound = true
		var oids []types.ID
		if cmd.category == "" {
			oids, err = m.List(ctx, q...)
		} else {
			oids, err = m.ListAttachedObjects(ctx, cmd.category, cmd.tag)
		}

		if err != nil {
			return err
		}
		for _, id := range oids {
			ids = append(ids, id.Id)
		}
	}

	for _, id := range ids {
		o, err := m.Retrieve(ctx, id)
		if err != nil {
			if filterNotFound && fault.Is(err, &types.NotFound{}) {
				// The case when an FCD is deleted by something other than DeleteVStorageObject_Task, such as VM destroy
				if cmd.all {
					obj := VStorageObject{VStorageObject: types.VStorageObject{
						Config: types.VStorageObjectConfigInfo{
							BaseConfigInfo: types.BaseConfigInfo{
								Id:   types.ID{Id: id},
								Name: "not found: use 'disk.ls -R' to reconcile datastore inventory",
							},
						},
					}}
					res.Objects = append(res.Objects, obj)
				}
				continue
			}
			return fmt.Errorf("retrieve %q: %s", id, err)
		}

		obj := VStorageObject{VStorageObject: *o}
		if cmd.tags {
			obj.Tags, err = m.ListAttachedTags(ctx, id)
			if err != nil {
				return err
			}
		}
		res.Objects = append(res.Objects, obj)
	}

	return cmd.WriteResult(&res)
}
