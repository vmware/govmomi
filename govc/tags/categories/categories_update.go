package tags

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/tags/tags_helper"
)

type update struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("tags.categories.update", &update{})
}

func (cmd *update) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *update) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *update) Usage() string {
	return `
	Examples:
	  govc tags.categories.create 'name' 'description' 'categoryType' 'multiValue(true/false)'`
}

func (cmd *update) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 5 {
		return flag.ErrHelp
	}
	id := f.Arg(0)

	name := f.Arg(1)
	description := f.Arg(2)
	categoryType := f.Arg(3)
	multiValue := f.Arg(4)

	category := new(tags.CategoryUpdateSpec)

	if categoryType == "" {
		category.UpdateSpec = tags.CategoryUpdate{
			// AssociableTypes: categoryType,
			Cardinality: multiValue,
			Description: description,
			Name:        name,
		}

	} else {
		types := strings.Split(categoryType, ",")
		category.UpdateSpec = tags.CategoryUpdate{
			AssociableTypes: types,
			Cardinality:     multiValue,
			Description:     description,
			Name:            name,
		}

	}

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		err := c.UpdateCategory(ctx, id, category)

		if err != nil {

			fmt.Printf(err.Error())
			return err
		}
		return nil

	})
}
