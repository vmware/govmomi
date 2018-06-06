package tags

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/tags/tags_helper"
)

type getbyname struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("tags.categories.getbyname", &getbyname{})
}

func (cmd *getbyname) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *getbyname) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *getbyname) Usage() string {
	return `
	Examples:
	  govc tags.categories.getbyname 'name'`
}

func (cmd *getbyname) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	name := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {

		categories, err := c.GetCategoriesByName(context.Background(), name)
		if err != nil {
			fmt.Printf(err.Error())
			return nil
		}
		fmt.Println(categories)
		return nil

	})
}
