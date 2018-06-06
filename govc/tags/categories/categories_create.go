package tags

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/tags/tags_helper"
)

type create struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("tags.categories.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Usage() string {
	return `
	Examples:
	  govc tags.categories.create 'name' 'description' 'categoryType' 'multiValue(true/false)'`
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 4 {
		return flag.ErrHelp
	}

	name := f.Arg(0)
	description := f.Arg(1)
	categoryType := f.Arg(2)
	multiValue := f.Arg(3)
	value := false
	if multiValue == "true" {
		value = true
	}

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		id, err := c.CreateCategoryIfNotExist(ctx, name, description, categoryType, value)
		if err != nil {

			fmt.Printf(err.Error())
			return err
		}

		fmt.Printf("Create category success id %s\n", *id)

		return nil

	})
}
