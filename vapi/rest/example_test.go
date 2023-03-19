/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package rest_test

import (
	"context"
	"fmt"
	"net/url"

	_ "github.com/vmware/govmomi/lookup/simulator"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/sts"
	_ "github.com/vmware/govmomi/sts/simulator"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
)

func ExampleClient_LoginByToken() {
	simulator.Run(func(ctx context.Context, vc *vim25.Client) error {
		c, err := sts.NewClient(ctx, vc)
		if err != nil {
			return err
		}

		// Issue a bearer token
		req := sts.TokenRequest{
			Userinfo: url.UserPassword("Administrator@VSPHERE.LOCAL", "password"),
		}

		signer, err := c.Issue(ctx, req)
		if err != nil {
			return err
		}

		rc := rest.NewClient(vc)

		err = rc.LoginByToken(rc.WithSigner(ctx, signer))
		if err != nil {
			return err
		}

		session, err := rc.Session(ctx)
		if err != nil {
			return err
		}

		// Note: vcsim does not currently parse the token NameID for rest as it does for soap
		fmt.Println(session.User)

		return nil
	})
	// Output: TODO
}
