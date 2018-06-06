package tags

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/tags/tags_helper"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/vim25/soap"
)

type ls struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("tags.categories.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return `
	Examples:
	  govc tags.categories.ls`
}

func withClient(ctx context.Context, cmd *flags.ClientFlag, f func(*tags.RestClient) error) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	govcUrl := "https://" + os.Getenv("GOVC_URL")

	URL, err := url.Parse(govcUrl)
	if err != nil {
		return err
	}

	c := tags.NewClient(URL, true, "")
	if err != nil {
		return err
	}

	// SSO admin server has its own session manager, so the govc persisted session cookies cannot
	// be used to authenticate.  There is no SSO token persistence in govc yet, so just use an env
	// var for now.  If no GOVC_LOGIN_TOKEN is set, issue a new token.
	token := os.Getenv("GOVC_LOGIN_TOKEN")
	header := soap.Header{
		Security: &sts.Signer{
			Certificate: vc.Certificate(),
			Token:       token,
		},
	}

	if token == "" {
		tokens, cerr := sts.NewClient(ctx, vc)
		if cerr != nil {
			return cerr
		}

		req := sts.TokenRequest{
			Certificate: vc.Certificate(),
			Userinfo:    cmd.Userinfo(),
		}

		header.Security, cerr = tokens.Issue(ctx, req)
		if cerr != nil {
			return cerr
		}
	}

	if err = c.Login(context.TODO()); err != nil {
		return err
	}

	return f(c)
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {
		categories, err := c.ListCategories(ctx)
		if err != nil {
			fmt.Printf(err.Error())
			return err
		}

		fmt.Println(categories)
		return nil

	})
}
