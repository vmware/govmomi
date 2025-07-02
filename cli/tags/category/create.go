// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package category

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type create struct {
	*flags.ClientFlag
	cat   tags.Category
	multi bool
}

func init() {
	cli.Register("tags.category.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	f.StringVar(&cmd.cat.CategoryID, "id", "", "Category ID")
	f.StringVar(&cmd.cat.Description, "d", "", "Description")
	f.Var((*kinds)(&cmd.cat.AssociableTypes), "t", "Object types")
	f.BoolVar(&cmd.multi, "m", false, "Allow multiple tags per object")
}

type kinds []string

func (e *kinds) String() string {
	return fmt.Sprint(*e)
}

func (e *kinds) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func cardinality(multi bool) string {
	if multi {
		return "MULTIPLE"
	}
	return "SINGLE"
}

func (cmd *create) Usage() string {
	return "NAME"
}

func (cmd *create) Description() string {
	return `Create tag category.

This command will output the ID of the new tag category.

Examples:
  govc tags.category.create -d "Kubernetes region" -t Datacenter k8s-region
  govc tags.category.create -d "Kubernetes zone" k8s-zone`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	cmd.cat.Name = f.Arg(0)
	cmd.cat.Cardinality = cardinality(cmd.multi)

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	id, err := tags.NewManager(c).CreateCategory(ctx, &cmd.cat)
	if err != nil {
		return err
	}

	fmt.Println(id)
	return nil
}
