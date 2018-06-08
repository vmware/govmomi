/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tags

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/sts"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25/soap"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("tags.category.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Usage() string {
	return `List all categories
Examples:
  govc tags.category.ls`
}

func withClient(ctx context.Context, cmd *flags.ClientFlag, f func(*tags.RestClient) error) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}
	usrDecode, err := url.QueryUnescape(cmd.Userinfo().String())
	if err != nil {
		return err
	}

	govcUrl := "https://" + usrDecode + "@" + vc.URL().Hostname()

	URL, err := url.Parse(govcUrl)
	if err != nil {
		return err
	}

	c := tags.NewClient(URL, true, "")
	if err != nil {
		return err
	}

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

	if err = c.Login(ctx); err != nil {
		return err
	}

	return f(c)
}

type getResult []string

func (r getResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {
		categories, err := c.ListCategories(ctx)
		if err != nil {
			return err
		}

		result := getResult(categories)
		cmd.WriteResult(result)
		return nil

	})
}
