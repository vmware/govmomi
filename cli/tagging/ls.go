// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tagging

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/tagging"
)

type names []string

func (n *names) String() string     { return strings.Join(*n, ",") }
func (n *names) Set(v string) error { *n = append(*n, v); return nil }

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	names    names
	c        string // category name filter, tags only
	category bool
}

func init() {
	cli.Register("tagging.tag.ls", &ls{})
	cli.Register("tagging.category.ls", &ls{category: true})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)

	if !cmd.category {
		f.StringVar(&cmd.c, "c", "", "Category name")
	}
	f.Var(&cmd.names, "n", "Only list items with this name (can be specified multiple times)")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Usage() string {
	return "[-n NAME]..."
}

func (cmd *ls) Description() string {
	if cmd.category {
		return `List categories using the vSphere 9.0+ /api/vcenter/tagging/categories API.

Examples:
  govc tagging.category.ls
  govc tagging.category.ls -n my-category
  govc tagging.category.ls -json | jq .`
	}
	return `List tags using the vSphere 9.0+ /api/vcenter/tagging/tags API.

Examples:
  govc tagging.tag.ls
  govc tagging.tag.ls -n my-tag -n other-tag
  govc tagging.tag.ls -c k8s-zone
  govc tagging.tag.ls -json | jq .`
}

// tagResult is the tag list, with category ids resolved to category names for text output.
type tagResult struct {
	items []tagging.TagListItem
	cats  map[string]string // category id -> name
}

func (r tagResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	for _, i := range r.items {
		cat := i.Info.Category
		if name, ok := r.cats[cat]; ok {
			cat = name
		}
		fmt.Fprintf(tw, "%s\t%s\n", i.Info.Name, cat)
	}
	return tw.Flush()
}

func (r tagResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Items []tagging.TagListItem `json:"items"`
	}{Items: append([]tagging.TagListItem{}, r.items...)})
}

type categoryResult []tagging.CategoryListItem

func (r categoryResult) Write(w io.Writer) error {
	for _, i := range r {
		fmt.Fprintln(w, i.Info.Name)
	}
	return nil
}

func (r categoryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Items []tagging.CategoryListItem `json:"items"`
	}{Items: append([]tagging.CategoryListItem{}, r...)})
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	m := tagging.NewManager(c)

	if cmd.category {
		items, err := m.ListCategories(ctx, cmd.names...)
		if err != nil {
			return err
		}
		return cmd.WriteResult(categoryResult(items))
	}

	cats := make(map[string]string) // category id -> name

	var items []tagging.TagListItem
	if cmd.c == "" {
		// Category ids are resolved to names for the output.
		categories, err := m.ListCategories(ctx)
		if err != nil {
			return err
		}
		for _, category := range categories {
			cats[category.CategoryID] = category.Info.Name
		}

		items, err = m.ListTags(ctx, cmd.names...)
		if err != nil {
			return err
		}
	} else {
		// ListTagsForCategory looks up the category by name, so no need to list all categories,
		// every item is in the category named by -c.
		items, err = m.ListTagsForCategory(ctx, cmd.c, cmd.names...)
		if err != nil {
			return err
		}
		for _, item := range items {
			cats[item.Info.Category] = cmd.c
		}
	}

	return cmd.WriteResult(tagResult{items, cats})
}
