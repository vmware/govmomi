package tags

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/tags/tags_helper"
)

type delete struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("tags.categories.delete", &delete{})
}

func (cmd *delete) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *delete) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *delete) Usage() string {
	return `
	Examples:
	  govc tags.categories.delete 'id'`
}

func (cmd *delete) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	categoryID := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		err := c.DeleteCategory(ctx, categoryID)
		if err != nil {

			fmt.Printf(err.Error())
			return err
		}
		return nil

	})
}
