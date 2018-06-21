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

package association

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/tags"
)

type ls struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("tags.association.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *ls) Usage() string {
	return "ID or PATH"
}

func (cmd *ls) Description() string {
	return `List all attached tags or objects.

Examples:
  govc tags.association.ls ID
  govc tags.association.ls PATH`
}

func withClient(ctx context.Context, cmd *flags.ClientFlag, f func(*tags.RestClient) error) error {
	vc, err := cmd.Client()
	if err != nil {
		return err
	}
	tagsURL := vc.URL()
	tagsURL.User = cmd.Userinfo()

	c := tags.NewClient(tagsURL, !cmd.IsSecure(), "")
	if err != nil {
		return err
	}

	if err = c.Login(ctx); err != nil {
		return err
	}
	defer c.Logout(ctx)

	return f(c)
}

func isTagID(arg string) bool {
	s := strings.Split(arg, ":")
	if len(s) > 2 {
		return true
	}
	return false
}

type getAssociated []tags.AssociatedObject

func (r getAssociated) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

type getResult []string

func (r getResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i])
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	arg := f.Arg(0)

	return withClient(ctx, cmd.ClientFlag, func(c *tags.RestClient) error {
		if isTagID(arg) {
			objAssociated, err := c.ListAttachedObjects(ctx, arg)
			if err != nil {
				return err
			}

			result := getAssociated(objAssociated)
			cmd.WriteResult(result)
			return nil
		}

		ref, err := convertPath(ctx, cmd.DatacenterFlag, arg)
		if err != nil {
			return err
		}

		tagsAssociated, err := c.ListAttachedTags(ctx, ref)
		if err != nil {
			return err
		}
		result := getResult(tagsAssociated)
		cmd.WriteResult(result)
		return nil

	})

}
